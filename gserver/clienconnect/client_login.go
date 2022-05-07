package clienconnect

import (
	"fmt"
	"server/db"
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/gserver/nodeManange"
	pbaccount "server/proto/account"
	"server/proto/protocol_base"
	pbrole "server/proto/role"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Client) accountLogin(msg *pbaccount.C2S_Login) {
	retmsg := &pbaccount.S2C_Login{
		Retcode:  0,
		RoleInfo: &pbrole.Pb_RoleInfo{},
	}

	//已登陆
	if c.connectState != StatusSockert {
		retmsg.Retcode = cfg.GetErrorCodeNumber("LOGIN")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_Login), retmsg)
	}

	accountinfo := &commonstruct.AccountInfo{
		Account: msg.Account,
	}
	ok := accountinfo.GetAccountinfo()
	c.roleID = accountinfo.RoleID
	//未找到账号
	if !ok {
		retmsg.Retcode = cfg.GetErrorCodeNumber("AccountNull")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_Login), retmsg)
		return
	}
	if accountinfo.Password != msg.Password {
		retmsg.Retcode = cfg.GetErrorCodeNumber("PASSWORD_ERROR")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_Login), retmsg)
		return
	}

	c.connectState = StatusLogin
	roledata := commonstruct.GetRoleAllData(accountinfo.RoleID)
	//未创建账号 角色为空
	if roledata.RoleBase.Name == "" {
		retmsg.Retcode = cfg.GetErrorCodeNumber("RoleNull")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_Login), retmsg)
		return
	}

	//账号已登陆
	c.registerName = fmt.Sprintf("role_%v", accountinfo.RoleID)
	node := nodeManange.GetNode(nodeManange.GateNode)
	if registerProcess := node.ProcessByName(c.registerName); registerProcess != nil {
		//if c.process.Self() == registerProcess.Self()
		node.UnregisterName(c.registerName)
		c.process.Send(registerProcess.Self(), etf.Atom("Extrusionline"))
	}

	retmsg.RoleInfo = roledata.RoleBase.ToPB()

	//绑定genserver name
	if error := c.process.RegisterName(c.registerName); error != nil {
		logrus.Errorf("绑定genserver name 失败: [%v]  [%v]  [%v]", error, c.registerName, accountinfo)
	}

	c.connectState = StatusGame
	roledata.RoleBase.Online = true
	commonstruct.StoreRoleData(roledata)
	c.roleData = roledata

	c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module),
		int32(pbaccount.MSG_ACCOUNT_Login),
		retmsg)
}

func (c *Client) registerAccount(msg *pbaccount.C2S_Register) {
	retmsg := &pbaccount.S2C_Register{
		Retcode: 0,
	}
	//已登陆
	if c.connectState != StatusSockert {
		retmsg.Retcode = cfg.GetErrorCodeNumber("LOGIN")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_Register), retmsg)
	}

	accountinfo := &commonstruct.AccountInfo{
		Account: msg.Account,
	}
	//已注册
	if ok := accountinfo.GetAccountinfo(); ok {
		retmsg.Retcode = cfg.GetErrorCodeNumber("AccountExists")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_Register), retmsg)
		return
	}

	//cdk已注册
	if msg.CDK != "" {
		if ok := commonstruct.GetCDKinfo(msg.CDK); ok {
			retmsg.Retcode = cfg.GetErrorCodeNumber("CDK")
			c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_Register), retmsg)
			return
		}
	}

	c.roleID = commonstruct.GetNewRoleID()
	if c.roleID == 0 {
		logrus.Error("系统错误  获取角色ID失败")
		retmsg.Retcode = cfg.GetErrorCodeNumber("SystemError")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_Register), retmsg)
		return
	}

	accountinfo = &commonstruct.AccountInfo{
		Account:            msg.Account,
		Password:           msg.Password,
		RegistrationSource: msg.Source,
		Equipment:          msg.Equipment,
		CDK:                msg.CDK,
		RoleUUID:           c.process.Name(),
		RoleID:             c.roleID,
		RegistrationTime:   time.Now().Unix(),
		Settings:           make(map[uint32]string),
	}

	db.InsertOne(db.AccountTable, accountinfo)

	//角色信息
	roleBase := &commonstruct.RoleBaseInfo{
		ZoneID:            commonstruct.ServerCfg.ServerID,
		RoleID:            c.roleID,
		Name:              "",
		HeadID:            0,
		Sex:               0,
		Level:             1,
		Exp:               0,
		PracticeTimestamp: time.Now().Unix(),
		AttributeValue:    map[uint32]int64{},
		BodyList:          map[uint32]*commonstruct.RoleBodyInfo{},
		CE:                0,
		ItemList:          map[string]*commonstruct.ItemInfo{},
		Online:            true,
		State:             0,

		DirtyDataRecord: commonstruct.DirtyDataRecord{TableName: db.RoleBaseTable, DirtyDataList: map[string]primitive.E{}},
		// DirtyDataList:     map[string]primitive.E{},
	}
	roleBase.CalculationProperties()
	db.InsertOne(db.RoleBaseTable, roleBase)

	//初始道具
	roleItems := &commonstruct.RoleItemlist{
		RoleID:          c.roleID,
		ItemList:        map[string]*commonstruct.ItemInfo{},
		DirtyDataRecord: commonstruct.DirtyDataRecord{TableName: db.RoleItemsTable, DirtyDataList: map[string]primitive.E{}},
	}
	db.InsertOne(db.RoleItemsTable, roleItems)

	roledata := &commonstruct.RoleData{
		Acconut:   accountinfo,
		RoleBase:  *roleBase,
		RoleItems: *roleItems,
	}
	c.connectState = StatusLogin
	commonstruct.StoreRoleData(roledata)

	c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module),
		int32(pbaccount.MSG_ACCOUNT_Register), retmsg)
}

func (c *Client) accountCreateRole(msg *pbaccount.C2S_CreateRole) {
	retmsg := &pbaccount.S2C_Login{
		Retcode:  0,
		RoleInfo: &pbrole.Pb_RoleInfo{},
	}
	//未登陆
	if c.connectState == StatusSockert {
		retmsg.Retcode = cfg.GetErrorCodeNumber("NOT_LOGIN")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_CreateRole), retmsg)
		return
	}

	roledata := commonstruct.GetRoleAllData(c.roleID)
	//已创建账号
	if roledata.RoleBase.Name != "" {
		retmsg.Retcode = cfg.GetErrorCodeNumber("AccountExists")
		c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module), int32(pbaccount.MSG_ACCOUNT_CreateRole), retmsg)
		return
	}

	roledata.RoleBase.Name = msg.RoleName
	roledata.RoleBase.HeadID = msg.HeadID
	roledata.RoleBase.Sex = msg.Sex
	roledata.RoleBase.Exp = 0
	roledata.RoleBase.PracticeTimestamp = time.Now().Unix()
	roledata.RoleBase.Online = true
	roledata.RoleBase.SetDirtyData(primitive.E{Key: "name", Value: msg.RoleName})
	roledata.RoleBase.SetDirtyData(primitive.E{Key: "headid", Value: msg.HeadID})
	roledata.RoleBase.SetDirtyData(primitive.E{Key: "sex", Value: msg.Sex})
	roledata.RoleBase.SetDirtyData(primitive.E{Key: "exp", Value: 0})
	roledata.RoleBase.SetDirtyData(primitive.E{Key: "practicetimestamp", Value: time.Now().Unix()})
	commonstruct.StoreRoleData(roledata)
	commonstruct.SaveRoleData(roledata)
	c.connectState = StatusGame
	c.roleData = roledata

	c.SendToClient(int32(pbaccount.MSG_ACCOUNT_Module),
		int32(pbaccount.MSG_ACCOUNT_CreateRole),
		&pbaccount.S2C_CreateRole{
			Retcode:  0,
			RoleInfo: roledata.RoleBase.ToPB(),
		})
}

//心跳
func (c *Client) heartBeat(msg *protocol_base.C2S_HeartBeat) {
	c.SendToClient(int32(protocol_base.MSG_BASE_Module),
		int32(protocol_base.MSG_BASE_HeartBeat),
		&protocol_base.S2C_HeartBeat{
			Timestamp: time.Now().Unix(),
		})
}
