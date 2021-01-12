package network

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"server/gserver/commonstruct"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

//ClientInterface client hander
type ClientInterface interface {
	OnConnect(netconn net.Conn, packet int32, msgchan chan commonstruct.ProcessMsg, cancelFunc context.CancelFunc)
	OnClose()
	OnMessage(module int32, method int32, buf []byte)
	Send(module int32, method int32, pb proto.Message)
	ProcessMessage(msg commonstruct.ProcessMsg)
	//SendMsg(msg commonstruct.ProcessMsg)
}

//NetInterface network
type NetInterface interface {
	Start(n *NetWorkx)
}

//NetWorkx 网络管理
type NetWorkx struct {
	//tcp/udp/kcp
	src NetInterface
	//ClientHander ClientInterface
	//包长度0 2 4
	Packet int32
	//读取超时时间
	Readtimeout int32

	MsgTime int32
	MsgNum  int32

	//tcp kcp
	NetType string
	//监听端口.
	Port int32
	//当前连接用户数量
	UserNumber int32
	//用户对象池  //nw.UserPool.Get().(*client).OnConnect()
	UserPool *sync.Pool
	//连接用户
	userlist map[string]ClientInterface
	//启动成功后回调
	StartHook func()
}

//NewNetWorkX    instance
func NewNetWorkX(pool *sync.Pool, port, packet, readtimeout int32, nettype string, msgtime, msgnum int32, funcHook func()) *NetWorkx {
	return &NetWorkx{
		Packet:      packet,
		NetType:     nettype,
		Port:        port,
		UserPool:    pool,
		userlist:    make(map[string]ClientInterface),
		Readtimeout: readtimeout,
		MsgTime:     msgtime,
		MsgNum:      msgnum,
		StartHook:   funcHook,
	}
}

//Start 启动网络服务
func (n *NetWorkx) Start() {
	switch n.NetType {
	case "kcp":
		log.Info("NetWorkx [kcp] port:", n.Port)
		n.src = &KCPNetwork{}
	case "tcp":
		log.Info("NetWorkx [tcp] port:", n.Port)
		n.src = &TCPNetwork{}
	default:
		log.Info("NetWorkx default [tcp] port:", n.Port)
		n.src = new(TCPNetwork)
	}

	//start socket
	go n.src.Start(n)

}

//HandleClient 消息处理
func (n *NetWorkx) HandleClient(conn net.Conn) {
	log.Infof("sockert connect RemoteAddr:[%v]", conn.RemoteAddr().String())

	c := n.UserPool.Get().(ClientInterface)
	n.onConnect()

	defer n.UserPool.Put(c)
	defer n.onClose()
	defer conn.Close()
	defer c.OnClose()

	// sendc := make(chan []byte, 1)
	//c.OnConnect(conn.RemoteAddr(), sendc)// go func(conn net.Conn) {
	// 	for {
	// 		select {
	// 		case buf := <-sendc:
	// 			le := IntToBytes(len(buf), n.Packet)
	// 			conn.Write(BytesCombine(le, buf))
	// 		case <-ctx.Done():
	// 			log.Debug("exit role sendGO")
	// 			return
	// 		}
	// 	}
	// }(conn)

	rootContext := context.Background()
	ctx, cancelFunc := context.WithCancel(rootContext)

	readc := make(chan []byte, 1)
	gamechan := make(chan commonstruct.ProcessMsg)
	c.OnConnect(conn, n.Packet, gamechan, cancelFunc)

	// for {
	// 	_, buf, e := UnpackToBlockFromReader(conn, n.Packet)
	// 	if e != nil {
	// 		log.Error("socket error:", e.Error())
	// 		return
	// 	}
	// 	module := int32(binary.BigEndian.Uint16(buf[n.Packet : n.Packet+2]))
	// 	method := int32(binary.BigEndian.Uint16(buf[n.Packet+2 : n.Packet+4]))
	// 	c.OnMessage(module, method, buf[n.Packet+4:])
	// 	//pb 消息拆包
	// 	// Decode protobuf -> buf[n.Packet:]
	// 	msginfo := &common.NetworkMsg{}
	// 	e = proto.Unmarshal(buf[n.Packet:], msginfo)
	// 	if e != nil {
	// 		log.Errorf("msg decode error[%s]", e.Error())
	// 		msgdata, _ := proto.Marshal(&common.NetworkMsg{
	// 			Module: 0,
	// 			Method: 1,
	// 		})
	// 		conn.Write(msgdata)
	// 	} else {
	// 		c.OnMessage(msginfo.Module, msginfo.Method, msginfo.MsgBytes)
	// 	}
	// }

	go func(conn net.Conn, cancelfunc context.CancelFunc) {
		unix := time.Now().Unix()
		msgNum := 0
		for {
			//超时
			if n.Readtimeout != 0 {
				readtimeout := time.Minute * time.Duration(n.Readtimeout)
				conn.SetReadDeadline(time.Now().Add(readtimeout))
			}

			_, buf, e := UnpackToBlockFromReader(conn, n.Packet)
			if e != nil {
				switch e {
				case io.EOF:
					cancelfunc()
				default:
					cancelfunc()
					log.Warn("socket error:", e.Error())
				}
				return
			}
			readc <- buf

			//间隔时间大于 N 分钟后 或者 接收到500条消息后 给连接送条信息
			now := time.Now().Unix()
			msgNum++

			if now > unix+int64(n.MsgTime) || msgNum >= int(n.MsgNum) {
				//log.Infof("time:=======>[%v] [%v]", time.Now().Format("15:04:05"), msgNum)
				gamechan <- commonstruct.ProcessMsg{MsgType: "TimeInterval", Data: msgNum}
				unix = now
				msgNum = 0
			}
		}
	}(conn, cancelFunc)

	for {
		select {
		case buf := <-readc:
			module := int32(binary.BigEndian.Uint16(buf[n.Packet : n.Packet+2]))
			method := int32(binary.BigEndian.Uint16(buf[n.Packet+2 : n.Packet+4]))
			c.OnMessage(module, method, buf[n.Packet+4:])
		case msg := <-gamechan:
			c.ProcessMessage(msg)
		case <-ctx.Done():
			log.Debug("exit role goroutine")
			return
		}
	}
}

func (n *NetWorkx) onConnect() {
	n.UserNumber++
	log.Debug("user number:", n.UserNumber)
}
func (n *NetWorkx) onClose() {
	n.UserNumber--
	log.Debug("user number:", n.UserNumber)
}

func checkError(err error) {
	if err != nil {
		log.Errorf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//IntToBytes int 转换为[]byte
func IntToBytes(i int, packet int32) []byte {
	var buf = make([]byte, 2)
	if packet == 2 {
		binary.BigEndian.PutUint16(buf, uint16(i))
	} else {
		binary.BigEndian.PutUint32(buf, uint32(i))
	}
	return buf
}

//BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

// UnpackToBlockFromReader -> unpack the first block from the reader.
func UnpackToBlockFromReader(reader io.Reader, packet int32) (int32, []byte, error) {
	if reader == nil {
		return 0, nil, errors.New("reader is nil")
	}
	var info = make([]byte, packet, packet)
	if e := readUntil(reader, info); e != nil {
		if e == io.EOF {
			return 0, nil, e
		}
		return 0, nil, e //errorx.Wrap(e)
	}

	length, e := LengthOf(info, packet)
	if e != nil {
		return 0, nil, e
	}
	var content = make([]byte, length, length)
	if e := readUntil(reader, content); e != nil {
		if e == io.EOF {
			return 0, nil, e
		}
		return 0, nil, e //errorx.Wrap(e)
	}

	return length, append(info, content...), nil
}

//LengthOf Length of the stream starting validly.
//Length doesn't include length flag itself, it refers to a valid message length after it.
func LengthOf(stream []byte, packet int32) (int32, error) {
	if len(stream) < int(packet) {
		return 0, errors.New(fmt.Sprint("stream lenth should be bigger than ", packet))
	}

	switch packet {
	case 2:
		return int32(binary.BigEndian.Uint16(stream[0:2])), nil
	case 4:
		return int32(binary.BigEndian.Uint32(stream[0:4])), nil
	default:
		errstr := fmt.Sprintf("stream lenth seting error  [packet: %v]", packet)
		return 0, errors.New(errstr)
	}

}

func readUntil(reader io.Reader, buf []byte) error {
	if len(buf) == 0 {
		return nil
	}
	var offset int
	for {
		n, e := reader.Read(buf[offset:])
		if e != nil {
			if e == io.EOF {
				return e
			}
			return e //errorx.Wrap(e)
		}
		//log.Debugf("offset:[%s]  buf[%s]", offset, len(buf))
		offset += n
		if offset >= len(buf) {
			break
		}
	}
	return nil
}
