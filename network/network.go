package network

import (
	"errors"
	"fmt"

	msg "server/proto"

	"github.com/golang/protobuf/proto"
)

//NetInterface network
type NetInterface interface {
	Start()
	Stop()
	Send(msg []byte)
}

//Send network sendmsg
// func Send(c *NetworkInterface, msg *msg.NetworkMsg) {
// 	fmt.Printf("NetworkMsg send ")
// }

//StartNetWorkServer 启动网络服务
func StartNetWorkServer(nettype string, port int) {
	fmt.Println("network start")
	var network NetInterface

	switch nettype {
	case "kcp":
		fmt.Println("start kcp port:", port)
	case "tcp":
		network = TCPNetwork{}
	default:
		network = new(TCPNetwork) // TCPNetwork{}
	}

	network.Start()
}

//EncodeSend send msg
func EncodeSend(network NetInterface, module int32, method int32, pb proto.Message) {
	// encode
	data, err := proto.Marshal(pb)
	if err != nil {
		fmt.Printf("proto encode error[%s]\n", err.Error())
		return
	}

	msg := &msg.NetworkMsg{}
	msg.MsgBytes = data
	msg.Module = module
	msg.Method = method
	msgdata, err := proto.Marshal(msg)
	if err != nil {
		fmt.Printf("NetworkMsg encode error[%s]\n", err.Error())
		return
	}
	network.Send(msgdata)
}

//Decode  decode  msg
func Decode(msgdata []byte, outpb proto.Message) (int32, int32, error) {
	msginfo := &msg.NetworkMsg{}

	err := proto.Unmarshal(msgdata, msginfo)
	if err != nil {
		fmt.Printf("msg decode error[%s]\n", err.Error())
		return 0, 0, errors.New("proto: msg.NetworkMsg decode error")
	}

	protoerr := proto.Unmarshal(msginfo.MsgBytes, outpb)
	if err != nil {
		fmt.Printf("proto decode error[%s]\n", protoerr.Error())
		return 0, 0, err
	}

	return msginfo.Module, msginfo.Method, nil
}
