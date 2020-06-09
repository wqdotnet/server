package network

import (
	"fmt"
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
	Stop()
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

func (n *NetWorkx) onConnect() {
	n.UserNumber++
}
func (n *NetWorkx) onClose() {
	n.UserNumber--
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
