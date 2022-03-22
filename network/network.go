package network

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"server/gserver/genServer"
	"server/tools"
	"sync/atomic"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/ergo-services/ergo/node"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

//ClientInterface client hander
// type ClientInterface interface {
// 	//OnConnect(sendchan chan []byte, packet int32, msgchan chan commonstruct.ProcessMsg, addr net.Addr)
// 	OnClose()
// 	OnMessage(module int32, method int32, buf []byte)
// 	Send(module int32, method int32, pb proto.Message)
// 	ProcessMessage(msg []byte) bool
// }

//NetInterface network
type NetInterface interface {
	Start(n *NetWorkx)
	Close()
}

//NetWorkx 网络管理
type NetWorkx struct {
	//tcp/udp/kcp
	src NetInterface

	//包长度0 2 4
	Packet int32
	//读取超时时间(秒)
	Readtimeout int32

	MsgTime int32
	MsgNum  int32

	//tcp kcp
	NetType string
	//监听端口.
	Port int32
	//用户对象池  //nw.UserPool.Get().(*client).OnConnect()
	//UserPool *sync.Pool
	CreateGenServerObj func() genServer.GateGenHanderInterface

	//启动成功后回调
	StartHook func()

	//新连接回调
	connectHook func()
	//连接关闭回调
	closedConnectHook func()
	//socket 关闭回调
	closeHook func()

	ConnectCount int64
	gateNode     node.Node
}

//NewNetWorkX    instance
func NewNetWorkX(createObj func() genServer.GateGenHanderInterface, port, packet, readtimeout int32, nettype string, msgtime, msgnum int32,
	startHook, closeHook, connectHook, closedConnectHook func()) *NetWorkx {

	netWorkx := &NetWorkx{
		Packet:  packet,
		NetType: nettype,
		Port:    port,
		//UserPool: pool,
		//userlist:    make(map[string]ClientInterface),
		CreateGenServerObj: createObj,
		Readtimeout:        readtimeout,
		MsgTime:            msgtime,
		MsgNum:             msgnum,
		StartHook:          startHook,
		closeHook:          closeHook,
		connectHook:        connectHook,
		closedConnectHook:  closedConnectHook,
	}
	atomic.StoreInt64(&netWorkx.ConnectCount, 0)
	return netWorkx
}

//Start 启动网络服务
func (n *NetWorkx) Start(gateNode node.Node) {
	n.gateNode = gateNode
	switch n.NetType {
	case "kcp":
		logrus.Info("NetWorkx [kcp] port:", n.Port)
		n.src = &KCPNetwork{}
	case "tcp":
		logrus.Info("NetWorkx [tcp] port:", n.Port)
		n.src = &TCPNetwork{}
	default:
		logrus.Info("NetWorkx default [tcp] port:", n.Port)
		n.src = new(TCPNetwork)
	}

	//start socket
	go n.src.Start(n)

}

func (n *NetWorkx) createProcess() (gen.Process, chan []byte, error) {
	//genserver := n.UserPool.Get().(ergo.GenServerBehaviour)
	clientHander := n.CreateGenServerObj()

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, nil, err
	}

	sendchan := make(chan []byte, 1)

	process, err := n.gateNode.Spawn(uid.String(), gen.ProcessOptions{}, &genServer.GateGenServer{}, sendchan, clientHander)
	if err != nil {
		return nil, nil, err
	}

	return process, sendchan, nil
}

//HandleClient 消息处理
func (n *NetWorkx) HandleClient(conn net.Conn) {
	process, sendchan, err := n.createProcess()
	if err != nil {
		logrus.Error("createProcess err: [%v]", err)
		return
	}

	n.onConnect()
	atomic.AddInt64(&n.ConnectCount, 1)
	defer atomic.AddInt64(&n.ConnectCount, -1)

	//defer n.UserPool.Put(p)
	defer n.onClosedConnect()
	defer conn.Close()

	defer process.Send(process.Self(), etf.Term(etf.Tuple{etf.Atom("$gen_cast"), etf.Atom("SocketStop")}))
	//defer process.Send(process.Self(), etf.Atom("SocketStop"))

	// sendc := make(chan []byte, 1)
	//c.OnConnect(conn.RemoteAddr(), sendc)
	// go func(conn net.Conn) {
	// 	for {
	// 		select {
	// 		case buf := <-sendc:
	// 			le := IntToBytes(len(buf), n.Packet)
	// 			conn.Write(BytesCombine(le, buf))
	// 		case <-ctx.Done():
	// 			logrus.Debug("exit role sendGO")
	// 			return
	// 		}
	// 	}
	// }(conn)

	rootContext := context.Background()
	sendctx, sendcancelFunc := context.WithCancel(rootContext)
	defer sendcancelFunc()

	//readchan := make(chan []byte, 1)
	//sendchan := make(chan []byte, 1)
	// gamechan := make(chan commonstruct.ProcessMsg)
	// c.OnConnect(sendchan, n.Packet, gamechan, conn.RemoteAddr())

	// for {
	// 	_, buf, e := UnpackToBlockFromReader(conn, n.Packet)
	// 	if e != nil {
	// 		logrus.Error("socket error:", e.Error())
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
	// 		logrus.Errorf("msg decode error[%s]", e.Error())
	// 		msgdata, _ := proto.Marshal(&common.NetworkMsg{
	// 			Module: 0,
	// 			Method: 1,
	// 		})
	// 		conn.Write(msgdata)
	// 	} else {
	// 		c.OnMessage(msginfo.Module, msginfo.Method, msginfo.MsgBytes)
	// 	}
	// }

	go func(conn net.Conn) {
		for {
			select {
			case buf := <-sendchan:
				le := tools.IntToBytes(int32(len(buf)), n.Packet)
				conn.Write(tools.BytesCombine(le, buf))
			case <-sendctx.Done():
				//logrus.Debug("exit role sendGO")
				return
			}
		}
	}(conn)

	//go func(conn net.Conn, sendcancelFunc context.CancelFunc) {
	unix := time.Now().Unix()
	msgNum := 0
	for {
		//超时
		if n.Readtimeout != 0 {
			readtimeout := time.Second * time.Duration(n.Readtimeout)
			conn.SetReadDeadline(time.Now().Add(readtimeout))
		}

		_, buf, e := UnpackToBlockFromReader(conn, n.Packet)
		if e != nil {
			switch e {
			case io.EOF:
				logrus.Debug("socket closed:", e.Error())
			default:
				logrus.Warn("socket closed:", e.Error())
			}
			return
		}
		//readchan <- buf

		module := int32(binary.BigEndian.Uint16(buf[n.Packet : n.Packet+2]))
		method := int32(binary.BigEndian.Uint16(buf[n.Packet+2 : n.Packet+4]))
		//process.Send(process.Self(), etf.Tuple{module, method, buf[n.Packet+4:]})
		process.Send(process.Self(), etf.Term(etf.Tuple{etf.Atom("$gen_cast"), etf.Tuple{module, method, buf[n.Packet+4:]}}))

		//间隔时间大于 N 分钟后 或者 接收到500条消息后 给连接送条信息
		now := time.Now().Unix()
		msgNum++

		if now > unix+int64(n.MsgTime) || msgNum >= int(n.MsgNum) {
			//logrus.Infof("time:=======>[%v] [%v]", time.Now().Format("15:04:05"), msgNum)

			process.Send(process.Self(), etf.Term(etf.Tuple{etf.Atom("$gen_cast"), "timeloop"}))
			//process.Send(process.Self(), "timeloop")

			//gamechan <- commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgTimeInterval, Data: msgNum}
			unix = now
			msgNum = 0
		}
	}
	//}(conn, sendcancelFunc)

	// for {
	// 	select {
	// 	case buf := <-readchan:
	// 		module := int32(binary.BigEndian.Uint16(buf[n.Packet : n.Packet+2]))
	// 		method := int32(binary.BigEndian.Uint16(buf[n.Packet+2 : n.Packet+4]))
	// 		c.OnMessage(module, method, buf[n.Packet+4:])
	// 	// case msg := <-gamechan:
	// 	// 	if !c.ProcessMessage(msg) {
	// 	// 		return
	// 	// 	}
	// 	case <-ctx.Done():
	// 		//logrus.Debug("exit role goroutine")
	// 		return
	// 	}
	// }

}

func (n *NetWorkx) onConnect() {
	if n.connectHook != nil {
		n.connectHook()
	}
}

func (n *NetWorkx) onClosedConnect() {
	if n.closedConnectHook != nil {
		n.closedConnectHook()
	}
}

//Close 关闭
func (n *NetWorkx) Close() {
	if n.closeHook != nil {
		n.closeHook()
	}
	n.src.Close()
}

func checkError(err error) {
	if err != nil {
		logrus.Errorf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// UnpackToBlockFromReader -> unpack the first block from the reader.
func UnpackToBlockFromReader(reader io.Reader, packet int32) (int32, []byte, error) {
	if reader == nil {
		return 0, nil, errors.New("reader is nil")
	}
	var info = make([]byte, packet)
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
	var content = make([]byte, length)
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
		//logrus.Debugf("offset:[%s]  buf[%s]", offset, len(buf))
		offset += n
		if offset >= len(buf) {
			break
		}
	}
	return nil
}
