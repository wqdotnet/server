package clienconnect

import (
	"server/msgproto/account"
	"server/msgproto/protocol_base"

	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

//消息路由
func (c *Client) rount(module int32, method int32, buf []byte) {
	//模块过滤....

	switch module {
	case int32(protocol_base.MSG_BASE_PROTOCOL_BASE):
		switch protocol_base.MSG_BASE(method) {
		case protocol_base.MSG_BASE_HeartBeat:
			c.heartbeat(buf)
		default:
			c.EmptyMsg(module, method)
		}
	case int32(account.MSG_ACCOUNT_Module):
		c.loginModule(method, buf)
	default:
		c.EmptyMsg(module, method)
	}
}

//心跳
func (c *Client) heartbeat(buf []byte) {
	hearbeat := &protocol_base.C2S_HeartBeat{}
	e := proto.Unmarshal(buf, hearbeat)
	if e != nil {
		log.Error(e)
	}

	c.Send(int32(protocol_base.MSG_BASE_PROTOCOL_BASE),
		int32(protocol_base.MSG_BASE_HeartBeat),
		&protocol_base.S2C_HeartBeat{
			Servertime: time.Now().Unix(),
		})
}

//防火墙
func (c *Client) breakWall(buf []byte) {

}

//EmptyMsg 接收到未识别的消息号
func (c *Client) EmptyMsg(module int32, method int32) {
	log.Warnf("Receive Empty Msg => [%v][%v] username:[%v]", module, method, c.username)
}
