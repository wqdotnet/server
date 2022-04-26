package clienconnect

import (
	"fmt"
	"server/db"
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/gserver/nodeManange"
	"server/proto/account"
	pbrole "server/proto/role"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/sirupsen/logrus"
)

func (c *Client) accountLogin(msg *account.C2S_Login) {
	retmsg := &account.S2C_Login{
		Retcode:  0,
		RoleInfo: &pbrole.Pb_RoleInfo{},
	}

	//已登陆
	if c.connectState != StatusSockert {
		retmsg.Retcode = cfg.GetErrorCodeNumber("LOGIN")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Login), retmsg)
	}

	accountinfo := &commonstruct.AccountInfo{
		Account: msg.Account,
	}
	ok := accountinfo.GetAccountinfo()
	c.roleID = accountinfo.RoleID
	//未找到账号
	if !ok {
		retmsg.Retcode = cfg.GetErrorCodeNumber("AccountNull")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Login), retmsg)
		return
	}
	if accountinfo.Password != msg.Password {
		retmsg.Retcode = cfg.GetErrorCodeNumber("PASSWORD_ERROR")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Login), retmsg)
		return
	}

	c.connectState = StatusLogin

	roledata := commonstruct.GetRoleAllData(accountinfo.RoleID)
	//未创建账号 角色为空
	if roledata.RoleBase == nil {
		retmsg.Retcode = cfg.GetErrorCodeNumber("RoleNull")

		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Login), retmsg)
		return
	}
	commonstruct.SaveRoleAllData(roledata)

	//账号已登陆
	c.registerName = fmt.Sprintf("role_%v", accountinfo.RoleID)
	node := nodeManange.GetNode(nodeManange.GateNode)
	if registerProcess := node.ProcessByName(c.registerName); registerProcess != nil {
		//if c.process.Self() == registerProcess.Self()
		node.UnregisterName(c.registerName)
		_, err := c.process.Call(registerProcess.Self(), etf.Atom("Extrusionline"))
		if err != nil {
			logrus.Errorf("重复登陆挤下线 :[%v] [%v] ", c.registerName, err)
		}
	} else {
		//获取了角色数据  roleData
		info := commonstruct.GetRoleAllData(accountinfo.RoleID)
		if info == nil {
			logrus.Warn("获取角色数据失败")
		} else {
			retmsg.RoleInfo = info.RoleBase.ToPB()
		}
	}

	//绑定genserver name
	if error := c.process.RegisterName(c.registerName); error != nil {
		logrus.Error("绑定genserver name 失败: ", error)
	}

	commonstruct.SaveRoleAllData(roledata)
	c.SendToClient(int32(account.MSG_ACCOUNT_Module),
		int32(account.MSG_ACCOUNT_Login),
		retmsg)
}

func (c *Client) registerAccount(msg *account.C2S_Register) {
	retmsg := &account.S2C_Register{
		Retcode: 0,
	}
	//已登陆
	if c.connectState != StatusSockert {
		retmsg.Retcode = cfg.GetErrorCodeNumber("LOGIN")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Register), retmsg)
	}

	accountinfo := &commonstruct.AccountInfo{
		Account: msg.Account,
	}
	//已注册
	if ok := accountinfo.GetAccountinfo(); ok {
		retmsg.Retcode = cfg.GetErrorCodeNumber("AccountExists")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Register), retmsg)
		return
	}

	//cdk已注册
	if msg.CDK != "" {
		if ok := commonstruct.GetCDKinfo(msg.CDK); ok {
			retmsg.Retcode = cfg.GetErrorCodeNumber("CDK")
			c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Register), retmsg)
			return
		}

	}

	accountinfo = &commonstruct.AccountInfo{
		Account:            msg.Account,
		Password:           msg.Password,
		RegistrationSource: msg.Source,
		Equipment:          msg.Equipment,
		CDK:                msg.CDK,
		RoleUUID:           c.process.Name(),
		RoleID:             commonstruct.GetNewRoleID(),
		RegistrationTime:   time.Now().Unix(),
		Settings:           make(map[uint32]string),
	}

	db.InsertOne(db.AccountTable, accountinfo)
	c.connectState = StatusLogin
	c.roleID = accountinfo.RoleID
	c.SendToClient(int32(account.MSG_ACCOUNT_Module),
		int32(account.MSG_ACCOUNT_Register), retmsg)
}

func (c *Client) accountCreateRole(msg *account.C2S_CreateRole) {
	retmsg := &account.S2C_Login{
		Retcode:  0,
		RoleInfo: &pbrole.Pb_RoleInfo{},
	}
	//未登陆
	if c.connectState == StatusSockert {
		retmsg.Retcode = cfg.GetErrorCodeNumber("NOT_LOGIN")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_CreateRole), retmsg)
		return
	}

	roledata := commonstruct.GetRoleAllData(c.roleID)
	//已创建账号
	if roledata.RoleBase != nil {
		retmsg.Retcode = cfg.GetErrorCodeNumber("AccountExists")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_CreateRole), retmsg)
		return
	}

	//角色信息
	roledata.RoleBase = &commonstruct.RoleBaseInfo{
		ZoneID:            commonstruct.ServerCfg.ServerID,
		RoleID:            c.roleID,
		Name:              msg.RoleName,
		HeadID:            msg.HeadID,
		Sex:               msg.Sex,
		Level:             1,
		Exp:               0,
		PracticeTimestamp: time.Now().Unix(),
		AttributeValue:    map[uint32]int64{},
		BodyList:          map[uint32]*commonstruct.RoleBodyInfo{},
		CE:                0,
		ItemList:          map[string]*commonstruct.ItemInfo{},
		Online:            true,
		State:             0,
		DirtyData:         false,
	}
	db.InsertOne(db.RoleBaseTable, roledata.RoleBase)

	//初始道具
	roledata.RoleItems = &commonstruct.RoleItemlist{
		RoleID:    c.roleID,
		ItemList:  map[string]*commonstruct.ItemInfo{},
		DirtyData: false,
	}
	db.InsertOne(db.RoleItemsTable, roledata.RoleItems)

	commonstruct.SaveRoleAllData(roledata)
	c.SendToClient(int32(account.MSG_ACCOUNT_Module),
		int32(account.MSG_ACCOUNT_CreateRole),
		&account.S2C_CreateRole{
			Retcode:  0,
			RoleInfo: roledata.RoleBase.ToPB(),
		})
}
