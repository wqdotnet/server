package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	msg "server/proto"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

//ClientInterface client hander
type ClientInterface interface {
	OnConnect(sendmsg chan []byte)
	OnClose()
	OnMessage(module int32, method int32, buf []byte)
}

//NetInterface network
type NetInterface interface {
	Start(n *NetWorkx)
	//Stop()
	//Send(msg []byte)
}

//NetWorkx 网络管理
type NetWorkx struct {
	//tcp/udp/kcp
	src NetInterface

	//ClientHander ClientInterface
	//包长度0 2 4
	Packet int32
	//tcp kcp
	NetType string
	//监听端口.
	Port int32
	//handlers map[int32]func(buf []byte)
	//当前连接用户数量
	UserNumber int32
	//用户对象池  //nw.UserPool.Get().(*client).OnConnect()
	UserPool *sync.Pool
}

//NewNetWorkX    instance
func NewNetWorkX(pool *sync.Pool) *NetWorkx {
	return &NetWorkx{
		Packet:   2,
		NetType:  "TCP",
		Port:     3344,
		UserPool: pool,
	}
}

//Start 启动网络服务
func (n *NetWorkx) Start() {
	fmt.Println("network start")
	switch n.NetType {
	case "kcp":
		fmt.Println("start kcp port:", n.Port)
		n.src = &KCPNetwork{}
	case "tcp":
		fmt.Println("start tcp port:", n.Port)
		n.src = &TCPNetwork{}
	default:
		fmt.Println("start default [tcp] port:", n.Port)
		n.src = new(TCPNetwork) // TCPNetwork{}
	}

	//start socket
	n.src.Start(n)

}

//HandleClient 消息处理
func (n *NetWorkx) HandleClient(conn net.Conn) {
	c := n.UserPool.Get().(ClientInterface)
	n.onConnect()
	defer c.OnClose()
	defer conn.Close()
	defer n.onClose()
	defer n.UserPool.Put(c)
	//超时
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout

	sendc := make(chan []byte, 1)
	c.OnConnect(sendc)
	go func(conn net.Conn) {

		for {
			select {
			case buf := <-sendc:
				le := intToBytes(len(buf), n.Packet)
				conn.Write(bytesCombine(le, buf))
			default:
			}

		}
	}(conn)

	for {
		_, buf, e := UnpackToBlockFromReader(conn, n.Packet)
		if e != nil {
			fmt.Println("socket error:", e.Error())
			return
		}

		// module := int32(binary.BigEndian.Uint16(buf[n.Packet : n.Packet+2]))
		// method := int32(binary.BigEndian.Uint16(buf[n.Packet+2 : n.Packet+4]))
		// c.OnMessage(module, method, buf[n.Packet+4:])

		// pb 消息拆包
		// Decode protobuf -> buf[n.Packet:]
		msginfo := &msg.NetworkMsg{}
		e = proto.Unmarshal(buf[n.Packet:], msginfo)
		if e != nil {
			fmt.Printf("msg decode error[%s]\n", e.Error())
			msgdata, _ := proto.Marshal(&msg.NetworkMsg{
				Module: 0,
				Method: 1,
			})
			conn.Write(msgdata)
		} else {
			c.OnMessage(msginfo.Module, msginfo.Method, msginfo.MsgBytes)
		}

	}
}

func (n *NetWorkx) onConnect() {
	n.UserNumber++
	fmt.Println("user number:", n.UserNumber)
}
func (n *NetWorkx) onClose() {
	n.UserNumber--
	fmt.Println("user number:", n.UserNumber)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// int 转换为[]byte
func intToBytes(i int, packet int32) []byte {
	var buf = make([]byte, 2)
	if packet == 2 {
		binary.BigEndian.PutUint16(buf, uint16(i))
	} else {
		binary.BigEndian.PutUint32(buf, uint32(i))
	}
	return buf
}

//BytesCombine 多个[]byte数组合并成一个[]byte
func bytesCombine(pBytes ...[]byte) []byte {
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
		offset += n
		if offset >= len(buf) {
			break
		}
	}
	return nil
}
