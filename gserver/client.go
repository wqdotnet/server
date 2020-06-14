package gserver

import "fmt"

type client struct {
	Connect chan interface{}
	Close   chan interface{}
	MsgByte chan map[int][]byte
	//OnMessage func(module int, method int, buf []byte)
	//handlers map[int32]interface{}
}

//NewClient
// func NewClient() *client {
// 	return &client{
// 		Connect: make(chan interface{}),
// 	}
// }

func (c *client) OnConnect() {
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
