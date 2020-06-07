package client

type client struct {
	//OnMessage func(module int, method int, buf []byte)
	//handlers map[int32]interface{}
}

func (c *client) OnConnect() {

}

func (c *client) OnMessage(module int, method int, buf []byte) {

}

func (c *client) OnClose() {

}
