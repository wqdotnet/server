package network

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

//ClientInterface client hander
type ClientInterface interface {
	OnConnect()
	OnClose()
	OnMessage(module int32, method int32, buf []byte)
}

//NetInterface network
type NetInterface interface {
	Start(n *NetWorkx)
	//Stop()
	//Send(msg []byte)
}

//ConnInterface 消息处理
// type ConnInterface interface {
// 	Read(b []byte) (n int, err error)
// 	Write(b []byte) (n int, err error)
// 	Close() error
// 	LocalAddr() net.Addr
// 	RemoteAddr() net.Addr
// }

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
	c.OnConnect()
	defer c.OnClose()
	defer conn.Close()
	for {
		_, buf, e := UnpackToBlockFromReader(conn, n.Packet)
		if e != nil {
			fmt.Println("socket error:", e.Error())
			return
		}

		//[:n]
		c.OnMessage(1, 1, buf[n.Packet:])

		// oneRead = buf
		// _, _ = oneRead.readN(int(packet))
		// client.OnMessage(1, 2, oneRead)

		//next 消息处理
		// _, err2 := conn.Write(NewByte(1, 2, 3, 4, 5, 6, 7, 8, 9))
		// if err2 != nil {
		// 	fmt.Println(err2.Error())
		// 	return
		// }
	}
}

func (n *NetWorkx) onConnect() {
	n.UserNumber++
}
func (n *NetWorkx) onClose() {
	n.UserNumber--
}

// UnpackToBlockFromReader -> unpack the first block from the reader.
// protocol is PackWithMarshaller().
// [4]byte -- length             fixed_size,binary big endian encode
// [4]byte -- messageID          fixed_size,binary big endian encode
// [4]byte -- headerLength       fixed_size,binary big endian encode
// [4]byte -- bodyLength         fixed_size,binary big endian encode
// []byte -- header              marshal by json
// []byte -- body                marshal by marshaller
// ussage:
// for {
//     blockBuf, e:= UnpackToBlockFromReader(reader)
// 	   go func(buf []byte){
//         // handle a message block apart
//     }(blockBuf)
//     continue
// }
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

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readOnce(reader io.Reader) ([]byte, error) {
	var buffer = make([]byte, 512, 512)
	var n int
	var e error

	n, e = reader.Read(buffer)
	if e != nil {
		return nil, e
	}

	return buffer[0:n], nil
}

// //RegisteredMethod 方法注册
// func (n *NetWorkx) RegisteredMethod(method int32, handler func(buf []byte)) {
// 	n.handlers[method] = handler
// }

// //OnMessage 消息路由
// func (n *NetWorkx) OnMessage(module int32, method int32, buf []byte) {
// 	handler, ok := n.handlers[method]
// 	if !ok {
// 		fmt.Println(fmt.Sprintf("method %d handler not found", method))
// 		return
// 	}
// 	//module  method 方法合法过滤验证
// 	handler(buf)
// }

// //EncodeSend send msg
// func EncodeSend(network NetInterface, module int32, method int32, pb proto.Message) {
// 	// encode
// 	data, err := proto.Marshal(pb)
// 	if err != nil {
// 		fmt.Printf("proto encode error[%s]\n", err.Error())
// 		return
// 	}

// 	msg := &msg.NetworkMsg{}
// 	msg.MsgBytes = data
// 	msg.Module = module
// 	msg.Method = method
// 	msgdata, err := proto.Marshal(msg)
// 	if err != nil {
// 		fmt.Printf("NetworkMsg encode error[%s]\n", err.Error())
// 		return
// 	}
// 	network.Send(msgdata)
// }

// //Decode  decode  msg
// func Decode(msgdata []byte, outpb proto.Message) (int32, int32, error) {
// 	msginfo := &msg.NetworkMsg{}

// 	err := proto.Unmarshal(msgdata, msginfo)
// 	if err != nil {
// 		fmt.Printf("msg decode error[%s]\n", err.Error())
// 		return 0, 0, errors.New("proto: msg.NetworkMsg decode error")
// 	}

// 	protoerr := proto.Unmarshal(msginfo.MsgBytes, outpb)
// 	if err != nil {
// 		fmt.Printf("proto decode error[%s]\n", protoerr.Error())
// 		return 0, 0, err
// 	}

// 	return msginfo.Module, msginfo.Method, nil
// }
