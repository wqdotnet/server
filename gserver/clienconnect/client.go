package clienconnect

import (
	"context"
	"net"
	"server/db"
	"server/gserver/bigmapmanage"
	"server/gserver/commonstruct"
	"server/gserver/process"
	"server/msgproto/account"
	"server/msgproto/bigmap"
	"server/msgproto/common"
	"server/network"
	"time"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

//Client 客户端连接
type Client struct {
	//addr     *net.Addr
	//sendchan chan []byte
	conn   net.Conn
	packet int32
	//用户 连接状态 [0:连接] [1:已登陆] [2:下线]
	status userStatus
	//接收内部消息
	msgChan    chan commonstruct.ProcessMsg
	cancelFunc context.CancelFunc

	//用户id
	accountid int32
	//userid int32
	account string

	//角色ID
	roleid int32
	//角色名
	rolename string
	//国家
	country int32

	//部队信息
	troopslist map[int32]*commonstruct.TroopsStruct
}

type userStatus int32

const (
	//StatusSockert socker 连接状态
	StatusSockert userStatus = 0
	//StatusLogin 已登陆成功
	StatusLogin userStatus = 1
	//StatusSqueezeOut 重复登陆 挤下线
	StatusSqueezeOut userStatus = 2
)

//NewClient create client
func NewClient() *Client {
	return &Client{
		troopslist: make(map[int32]*commonstruct.TroopsStruct),
		msgChan:    make(chan commonstruct.ProcessMsg),
	}
}

//InitAutoID 初始化自增id
func InitAutoID() {
	//账号表
	var accountinfo commonstruct.AccountInfoStruct
	db.FindFieldMax(db.AccountTable, "accountid", &accountinfo)
	db.RedisExec("SET", db.AccountTable, accountinfo.AccountID+1)
	log.Infof("initAutoID  table:[%v] autoid:[%v]", db.AccountTable, accountinfo.AccountID+1)

	//用户表
	var userinfo account.P_RoleInfo
	db.FindFieldMax(db.UserTable, "roleid", &userinfo)
	db.RedisExec("SET", db.UserTable, userinfo.RoleID+1)
	log.Infof("initAutoID  table:[%v] autoid:[%v]", db.UserTable, userinfo.RoleID+1)

	//部队表
	var troops bigmap.P_Troops
	db.FindFieldMax(db.TroopsTable, "troopsid", &troops)
	db.RedisExec("SET", db.TroopsTable, troops.TroopsID+1)
	log.Infof("initAutoID  table:[%v] autoid:[%v]", db.TroopsTable, troops.TroopsID+1)

}

//------------------------------------------------------------------------

//OnConnect 连接接入
// func (c *Client) OnConnect(addr net.Addr, sendc chan []byte) {
// 	//sendmsg <-
// 	c.sendchan = sendc
// 	c.addr = &addr
// 	log.Debugf("client OnConnect [%s][%s]", addr.Network(), addr.String())
// }
func (c *Client) OnConnect(netconn net.Conn, packet int32, msgchan chan commonstruct.ProcessMsg, cancelFunc context.CancelFunc) {
	c.conn = netconn
	c.packet = packet
	c.msgChan = msgchan
	c.cancelFunc = cancelFunc
}

//Send 发送消息
func (c *Client) Send(module int32, method int32, pb proto.Message) {
	//log.Debugf("client send msg [%v] [%v] [%v]", module, method, pb)
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorf("proto encode error[%v] [%v][%v] [%v]", err.Error(), module, method, pb)
		return
	}
	// msginfo := &common.NetworkMsg{}
	// msginfo.Module = module
	// msginfo.Method = method
	// msginfo.MsgBytes = data
	// msgdata, err := proto.Marshal(msginfo)
	// if err != nil {
	// 	log.Errorf("msg encode error[%s]\n", err.Error())
	// }

	// c.sendchan <- msgdata
	mldulebuf := network.IntToBytes(int(module), 2)
	methodbuf := network.IntToBytes(int(method), 2)
	//c.sendchan <- network.BytesCombine(mldulebuf, methodbuf, data)

	buf := network.BytesCombine(mldulebuf, methodbuf, data)
	le := network.IntToBytes(len(buf), c.packet)
	c.conn.Write(network.BytesCombine(le, buf))
}

//OnMessage  socker接受到的消息
func (c *Client) OnMessage(module int32, method int32, buf []byte) {
	//module 过滤模块
	//log.Debugf("c2s : [%v] [%v] buf:[%v]", module, method, len(buf))

	if c.status == StatusSockert && method > int32(account.MSG_ACCOUNT_C2S_CreateRole) {
		log.Warnf("用户未登陆  调用模块id:[%v][%v]", module, method)
		return
	}

	c.rount(module, method, buf)
}

func (c *Client) setLoginStatus() {
	c.status = StatusLogin
}

// //protobuf 解码
// func decode[T proto.Message](s T,buf []byte) T,err {
// 	hearbeat := &T{}
// 	if err:= proto.Unmarshal(buf, hearbeat); err!=nil{
// 		log.Error("decode error")
// 	}
// }

//ProcessMessage 内部消息
func (c *Client) ProcessMessage(msg commonstruct.ProcessMsg) {
	log.Debugf("[%v][%v][%v] ProcessMessage : %v", time.Now().Format("15:04:05"), c.account, c.rolename, msg)

	switch msg.MsgType {
	case "TimeInterval":
		//功能激活包  N 分钟后 或者收到N 条消息后
	case "RoleLogin":
		//重新登陆成功挤下线
		rlogin := msg.Data.(chan string)
		c.status = StatusSqueezeOut
		c.SaveRoleData()
		rlogin <- "over"
		c.cancelFunc()
	case "TroopsMove":
		troops := msg.Data.(commonstruct.TroopsStruct)
		//部队移动信息
		c.s2cMove(troops.TroopsID, troops.AreasIndex, int32(troops.State), troops.ArrivalTime)
	case "OverMove":
		//结束移动
		troops := msg.Data.(commonstruct.TroopsStruct)
		c.s2cMove(troops.TroopsID, troops.AreasIndex, int32(troops.State), troops.ArrivalTime)
		c.s2cUpdateTroopsInfo(&troops)
	case "OnFitht":
		//触发战斗
		troops := msg.Data.(commonstruct.TroopsStruct)
		c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_Fight), &bigmap.S2C_Fight{TroopsID: troops.TroopsID})
	case "OverFitht":
		troops := msg.Data.(commonstruct.TroopsStruct)
		if oldtroops := c.troopslist[troops.TroopsID]; oldtroops != nil {
			oldtroops.State = troops.State
			oldtroops.Number = troops.Number
			oldtroops.AreasIndex = troops.AreasIndex
			oldtroops.Level = troops.Level
			oldtroops.MoveStamp = 0

			victory := troops.State == common.TroopsState_Stationed
			c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_OverFight),
				&bigmap.S2C_OverFight{TroopsID: troops.TroopsID, Victory: victory})
		} else {
			log.Errorf("未找到部队 [%v] [%v]", oldtroops.TroopsID, c.rolename)
		}

	case "AreasState": //区域状态发生变化 (1.战斗  0.)
		areas := msg.Data.(bigmapmanage.AreasInfo)

		//区域状态
		s2cAreasinfo := &bigmap.S2C_AreasInfo{}
		s2cAreasinfo.AreasInfoList = append(s2cAreasinfo.AreasInfoList,
			&bigmap.P_AreasInfo{AreasIndex: areas.AreasIndex,
				Type:  areas.Occupy,
				State: areas.State})

		c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_AreasInfo), s2cAreasinfo)
	default:
		log.Warn("未知消息:", msg)
	}
}

// //SendMsg 内部消息发送
// func (c *Client) SendMsg(msg commonstruct.ProcessMsg) {
// 	go func() {
// 		c.msgChan <- msg
// 	}()
// }de

//OnClose 连接关闭
func (c *Client) OnClose() {
	//挤下线 或者未登陆 不进行数据处理
	if c.status == StatusSockert {
		return
	}

	if c.status == StatusSqueezeOut {
		log.Warnf("[%v]  [%v] [%v] 被挤下线", c.account, c.rolename, c.roleid)
		return
	}

	c.SaveRoleData()
	process.UnRegister(c.roleid)
	c.cleanData()
	log.Debugf("client OnClose  add:[%s]   account:[%v] ", c.conn.RemoteAddr(), c.account)
}

//SaveRoleData 数据持久
func (c *Client) SaveRoleData() {
	log.Debugf("[%v]  [%v] [%v] 角色数据保存", c.account, c.rolename, c.roleid)
	//角色部队数据保存
	for _, info := range c.troopslist {
		updateTroopsInfo(info)
	}
}

//清理数据
func (c *Client) cleanData() {
	c.accountid = 0
	c.account = ""
	c.roleid = 0
	c.rolename = ""
	c.troopslist = make(map[int32]*commonstruct.TroopsStruct)
	c.status = StatusSockert
}
