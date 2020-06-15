package gserver

import (
	"fmt"
	msg "server/proto"

	"github.com/golang/protobuf/proto"
)

type client struct {
	sendchan chan []byte
}

//NewClient
// func NewClient() *client {
// 	return &client{
// 		Connect: make(chan interface{}),
// 	}
// }

func (c *client) OnConnect(sendc chan []byte) {
	//sendmsg <-
	c.sendchan = sendc
	fmt.Println("client OnConnect")
}

func (c *client) OnMessage(module int32, method int32, buf []byte) {
	fmt.Println("client msg :", module, method, string(buf))

}

func (c *client) OnClose() {
	fmt.Println("client OnClose")
}

//UserLogin 用户登陆
func (c *client) UserLogin(buf []byte) {

}

func (c *client) Send(module int32, method int32, pb proto.Message) {
	data, err := proto.Marshal(pb)
	if err != nil {
		fmt.Printf("proto encode error[%s]\n", err.Error())
		return
	}

	msginfo := &msg.NetworkMsg{}
	msginfo.Module = module
	msginfo.Method = method
	msginfo.MsgBytes = data
	msgdata, err := proto.Marshal(msginfo)
	if err != nil {
		fmt.Printf("msg encode error[%s]\n", err.Error())
	}

	c.sendchan <- msgdata
}
