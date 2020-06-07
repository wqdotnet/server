package client

// import (
// 	msg "server/proto"
// )

type clienthandlers struct {
	//OnMessage func(module int, method int, buf []byte)
	handlers map[int32]interface{}
}

var Clienthandlers *clienthandlers

//RegisteredMethod 注册方法类型
func init() {
	Clienthandlers = &clienthandlers{}

	//Clienthandlers.Add(1, msg.SearchRequest)
}

func (c *clienthandlers) Add(messageID int32, handler interface{}) {
	c.handlers[messageID] = handler
}

func (c *clienthandlers) OnConnect() {

}

func (c *clienthandlers) OnMessage(module int, method int, buf []byte) {

}

func (c *clienthandlers) OnClose() {

}
