package clienconnect

import (
	"server/gserver/cservice"
	"server/msgproto/common"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

//Client 客户端连接
type Client struct {
	sendchan chan []byte
	//用户id
	userid int32
	//用户名
	username string
	//用户 连接状态 [0:连接] [1:已登陆] [2:下线]
	userstatus int32
}

//OnConnect 连接接入
func (c *Client) OnConnect(sendc chan []byte) {
	//sendmsg <-
	c.sendchan = sendc

	log.Debug("client OnConnect")
}

//OnClose 连接关闭
func (c *Client) OnClose() {
	cservice.UnRegister(c.username)
	log.Debug("client OnClose")
}

//Send 发送消息
func (c *Client) Send(module int32, method int32, pb proto.Message) {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorf("proto encode error[%s]\n", err.Error())
		return
	}

	msginfo := &common.NetworkMsg{}
	msginfo.Module = module
	msginfo.Method = method
	msginfo.MsgBytes = data
	msgdata, err := proto.Marshal(msginfo)
	if err != nil {
		log.Errorf("msg encode error[%s]\n", err.Error())
	}

	c.sendchan <- msgdata
}

//OnMessage 接受消息
func (c *Client) OnMessage(module int32, method int32, buf []byte) {
	c.rount(module, method, buf)
}

//GetSPType SPInterface
func (c *Client) GetSPType() cservice.CSType {
	return cservice.ClientConnect
}
