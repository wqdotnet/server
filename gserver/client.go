package gserver

import (
	"fmt"
	msg "server/proto"

	"github.com/golang/protobuf/proto"
)

type client struct {
	sendchan chan []byte

	//用户id
	userid int32
	//用户名
	username string
}

func (c *client) OnConnect(sendc chan []byte) {
	//sendmsg <-
	c.sendchan = sendc
	fmt.Println("client OnConnect")
}

func (c *client) OnClose() {
	fmt.Println("client OnClose")
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

func (c *client) OnMessage(module int32, method int32, buf []byte) {
	switch method {
	case 1:
		c.UserLogin(buf)
	}

}

//UserLogin 用户登陆
func (c *client) UserLogin(buf []byte) {

	c.Send(1, 1, &msg.SearchRequest{
		Query:         "asdf",
		PageNumber:    3,
		ResultPerPage: 2,
	})
}
