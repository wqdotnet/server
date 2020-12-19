package clienconnect

import (
	"net"
	"server/db"
	"server/gserver/cservice"
	"server/msgproto/account"
	"server/network"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

//Client 客户端连接
type Client struct {
	addr     *net.Addr
	sendchan chan []byte
	//用户id
	accountid int32
	//userid int32
	account string

	//角色ID
	roleid int32
	//角色名
	rolename string
	//用户 连接状态 [0:连接] [1:已登陆] [2:下线]
	status userStatus
}

type userStatus int32

const (
	//StatusSockert socker 连接状态
	StatusSockert userStatus = 0
	//StatusLogin 已登陆成功
	StatusLogin userStatus = 1
)

//------------------------------------------------------------------------

//OnConnect 连接接入
func (c *Client) OnConnect(addr net.Addr, sendc chan []byte) {
	//sendmsg <-
	c.sendchan = sendc
	c.addr = &addr

	log.Debugf("client OnConnect [%s][%s]", addr.Network(), addr.String())
}

//OnClose 连接关闭
func (c *Client) OnClose() {
	cservice.UnRegister(c.rolename)
	log.Debugf("client OnClose  add:%s   %v", *c.addr, c)
}

//Send 发送消息
func (c *Client) Send(module int32, method int32, pb proto.Message) {
	//log.Debugf("client send msg [%v] [%v] [%v]", module, method, pb)
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorf("proto encode error[%v]\n", err.Error())
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
	//log.Debugf("c2s : [%v] [%v] buf:[%v]", module, method, len(buf))

	if c.status == StatusSockert && module > int32(account.MSG_ACCOUNT_C2S_CreateRole) {
		log.Warn("用户未登陆  调用模块id:[%s][%s]", module, method)
	}

	c.rount(module, method, buf)
}

//GetSPType SPInterface
func (c *Client) GetSPType() cservice.CSType {
	return cservice.ClientConnect
}

func (c *Client) setLoginStatus() {
	c.status = StatusLogin
}

// //protobuf 解码
// func decode[T proto.Message](s []T) {
// 	hearbeat := &T{}
// 	if err:= proto.Unmarshal(buf, hearbeat); err!=nil{
// 		log.Error("decode error")
// 	}
// }

//InitAutoID 初始化自增id
func InitAutoID() {
	//账号表
	var accountinfo accountInfo
	db.FindFieldMax(db.AccountTable, "accountid", &accountinfo)
	db.RedisExec("SET", db.AccountTable, accountinfo.AccountID+1)

	log.Infof("initAutoID  table:[%v] autoid:[%v]", db.AccountTable, accountinfo.AccountID+1)

	//用户表
	var userinfo account.P_RoleInfo
	db.FindFieldMax(db.UserTable, "roleid", &userinfo)
	db.RedisExec("SET", db.UserTable, userinfo.RoleID+1)

	log.Infof("initAutoID  table:[%v] autoid:[%v]", db.UserTable, userinfo.RoleID+1)
}
