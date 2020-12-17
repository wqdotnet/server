package clienconnect

import (
	"net"
	"server/gserver/cservice"
	"server/network"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

//Client 客户端连接
type Client struct {
	addr     net.Addr
	sendchan chan []byte
	//用户id
	userid int32
	//用户名
	username string
	//用户 连接状态 [0:连接] [1:已登陆] [2:下线]
	userstatus int32
}

//OnConnect 连接接入
func (c *Client) OnConnect(addr net.Addr, sendc chan []byte) {
	//sendmsg <-
	c.sendchan = sendc
	c.addr = addr

	log.Debugf("client OnConnect [%s][%s]", addr.Network(), addr.String())
}

//OnClose 连接关闭
func (c *Client) OnClose() {
	cservice.UnRegister(c.username)
	log.Debug("client OnClose")
}

//Send 发送消息
func (c *Client) Send(module int32, method int32, pb proto.Message) {
	log.Debugf("client send msg [%s] [%s] [%s]", module, method, pb)
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorf("proto encode error[%s]\n", err.Error())
		return
	}
	mldulebuf := network.IntToBytes(int(module), 2)
	methodbuf := network.IntToBytes(int(method), 2)
	c.sendchan <- network.BytesCombine(mldulebuf, methodbuf, data)

	// msginfo := &common.NetworkMsg{}
	// msginfo.Module = module
	// msginfo.Method = method
	// msginfo.MsgBytes = data
	// msgdata, err := proto.Marshal(msginfo)
	// if err != nil {
	// 	log.Errorf("msg encode error[%s]\n", err.Error())
	// }

	// c.sendchan <- msgdata
}

//OnMessage 接受消息
func (c *Client) OnMessage(module int32, method int32, buf []byte) {
	//module 过滤模块
	log.Debugf("c2s : [%s] [%s] buf:[%s]", module, method, len(buf))
	c.rount(module, method, buf)
}

//GetSPType SPInterface
func (c *Client) GetSPType() cservice.CSType {
	return cservice.ClientConnect
}
