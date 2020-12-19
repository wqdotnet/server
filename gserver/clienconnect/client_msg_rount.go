package clienconnect

import (
	"server/msgproto/account"
	"server/msgproto/bigmap"
	"server/msgproto/protocol_base"

	"time"

	log "github.com/sirupsen/logrus"
)

//消息路由
func (c *Client) rount(module int32, method int32, buf []byte) {
	//模块过滤....

	switch module {
	case int32(protocol_base.MSG_BASE_PROTOCOL_BASE):
		switch protocol_base.MSG_BASE(method) {
		case protocol_base.MSG_BASE_C2SHeartBeat:
			c.heartbeat(buf)
		default:
			c.EmptyMsg(module, method)
		}
	case int32(account.MSG_ACCOUNT_Module):
		c.loginModule(method, buf)
	case int32(bigmap.MSG_BIGMAP_Module_BIGMAP):
		c.bigmapModule(method, buf)
	default:
		c.EmptyMsg(module, method)
	}
}

//心跳
func (c *Client) heartbeat(buf []byte) {
	// hearbeat := &protocol_base.C2S_HeartBeat{}
	// e := proto.Unmarshal(buf, hearbeat)
	// if e != nil {
	// 	log.Error(e)
	// }

	c.Send(int32(protocol_base.MSG_BASE_PROTOCOL_BASE),
		int32(protocol_base.MSG_BASE_S2CHeartBeat),
		&protocol_base.S2C_HeartBeat{
			Servertime: time.Now().Unix(),
		})
}

//EmptyMsg 接收到未识别的消息号
func (c *Client) EmptyMsg(module int32, method int32) {
	log.Warnf("Receive Empty Msg => [%v][%v] rolename:[%v]", module, method, c.rolename)
}
