package clienconnect

import (
	"fmt"
	"server/gserver/cfg"
	"server/gserver/nodeManange"
	"server/proto/account"
	pbrole "server/proto/role"

	"github.com/ergo-services/ergo/etf"
	"github.com/sirupsen/logrus"
)

func (c *Client) accountLogin(msg *account.C2S_Login) {
	logrus.Info("C2S_Login: ", msg.Account, msg.Password)
	retmsg := &account.S2C_Login{
		Retcode:  0,
		RoleInfo: &pbrole.Pb_RoleInfo{},
	}

	ok, accountinfo := GetAccountinfo(msg.Account, msg.Password)
	//未注册
	if !ok {
		retmsg.Retcode = cfg.GetErrorCodeNumber("")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Login), retmsg)
		return
	}

	//未创建账号
	if accountinfo.RoleID == 0 {
		retmsg.Retcode = cfg.GetErrorCodeNumber("")
		c.SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Login), retmsg)
		return
	}

	//账号已登陆
	rolePIDName := fmt.Sprintf("role_%v", accountinfo.RoleID)
	node := nodeManange.GetNode(nodeManange.GateNode)
	if registerProcess := node.ProcessByName(rolePIDName); registerProcess != nil {
		node.UnregisterName(rolePIDName)
		roleData, err := c.process.Call(registerProcess.Self(), etf.Atom("Extrusionline"))
		if err != nil {
			logrus.Errorf("重复登陆挤下线 未预料错误: [%v] [%v]", err, roleData)
		}

		//获取了角色数据  roleData
		//retmsg.RoleInfo = roleData.RoleBaseInfo.ToPB()
	} else {
		info := GetRoleAllData(accountinfo.RoleID)
		retmsg.RoleInfo = info.RoleBaseInfo.ToPB()
	}

	//绑定genserver name
	if error := c.process.RegisterName(rolePIDName); error != nil {
		logrus.Error("绑定genserver name 失败: ", error)
	}

	c.SendToClient(int32(account.MSG_ACCOUNT_Module),
		int32(account.MSG_ACCOUNT_Login),
		retmsg)
}

func (c *Client) registerAccount(msg *account.C2S_Register) {
	logrus.Info("C2S_Register: ", msg.Account, msg.Password)

	c.SendToClient(int32(account.MSG_ACCOUNT_Module),
		int32(account.MSG_ACCOUNT_Register),
		&account.S2C_Register{
			Retcode: 0,
		})
}

func (c *Client) accountCreateRole(msg *account.C2S_CreateRole) {
	logrus.Info("C2S_CreateRole: ", msg.RoleName)

	c.SendToClient(int32(account.MSG_ACCOUNT_Module),
		int32(account.MSG_ACCOUNT_CreateRole),
		&account.S2C_CreateRole{
			Retcode:  0,
			RoleInfo: &pbrole.Pb_RoleInfo{},
		})
}

// //用户登陆
// func (c *Client) userLogin(userlogin *account.C2S_Login) {
// 	filter := bson.D{
// 		primitive.E{Key: "account", Value: userlogin.Account},
// 		//primitive.E{Key: "password", Value: userlogin.Password},
// 	}

// 	if userlogin.GetAccount() == "" || userlogin.GetPassword() == "" {
// 		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login),
// 			&account.S2C_Login{Success: false, Msg: cfg.GetErrorCodeNumber("PARAMETER_EMPTY")})
// 		return
// 	}

// 	logrus.Debugf("login %v %v  status:[%v]", userlogin.Account, userlogin.Password, c.status)
// 	if c.status != StatusSockert {
// 		logrus.Warn(c.account, " 已登陆")
// 		return
// 	}
// 	accountinfo := &commonstruct.AccountInfoStruct{}
// 	if err := db.FindOneBson(db.AccountTable, accountinfo, filter); err != nil {
// 		accountid := db.GetAutoID(db.AccountTable)
// 		roleid := db.GetAutoID(db.UserTable)
// 		c.account = userlogin.Account
// 		c.accountid = accountid
// 		c.roleid = roleid

// 		//创建账号
// 		accountinfo.AccountID = accountid
// 		accountinfo.Account = userlogin.Account
// 		accountinfo.Password = userlogin.Password
// 		accountinfo.RoleID = roleid
// 		accountinfo.RegistrationTime = time.Now()
// 		db.InsertOne(db.AccountTable, accountinfo)

// 		//创建角色
// 		rand.Seed(time.Now().Unix())
// 		country := int32(rand.Intn(3) + 1)
// 		createRoleInfo(userlogin.Account, country, roleid)

// 	} else {
// 		if accountinfo.Password != userlogin.Password {
// 			c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login),
// 				&account.S2C_Login{Success: false, Msg: cfg.GetErrorCodeNumber("PASSWORD_ERROR")})
// 			return
// 		}

// 		c.accountid = accountinfo.AccountID
// 		c.account = accountinfo.Account
// 		c.roleid = accountinfo.RoleID
// 	}
// 	//设置连接状态为已登陆
// 	c.setLoginStatus()

// 	var userinfo account.P_RoleInfo
// 	filter = bson.D{primitive.E{Key: "roleid", Value: accountinfo.RoleID}}
// 	if err := db.FindOneBson(db.UserTable, &userinfo, filter); err != nil {
// 		logrus.Debug("未找到 角色ID:", accountinfo.RoleID)
// 		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login), &account.S2C_Login{Success: true})
// 		return
// 	}

// 	//登陆成功
// 	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login), &account.S2C_Login{Success: true, RoleInfo: &userinfo})
// 	c.hookLogin(userinfo.RoleName, userinfo.Country, &userinfo)
// }

// //创建账号  及  角色信息
// func createRoleInfo(rolename string, country, roleid int32) {
// 	//角色游戏信息
// 	db.InsertOne(db.UserTable, &account.P_RoleInfo{
// 		RoleID:        roleid,
// 		RoleName:      rolename,
// 		Country:       country,
// 		Level:         0,
// 		TesourcesList: map[int32]int32{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0},
// 		Settings: &account.RoleSettings{
// 			AutoSelectTactics: true,
// 		},
// 	})

// 	//创建部队信息
// 	for i := 1; i < 6; i++ {
// 		troops := commonstruct.NewTroops(rolename, db.GetAutoID(db.TroopsTable), 0, country, int32(i))
// 		troops.AreasIndex = cfg.GetCountryAreasIndex(country)
// 		troops.StageNumber = int32(i)
// 		troops.Roleid = roleid
// 		db.InsertOne(db.TroopsTable, troops)
// 	}

// }

// func (c *Client) createRole(userlogin *account.C2S_CreateRole) {
// 	returnmsg := &account.S2C_CreateRole{Success: false}
// 	if userlogin.GetRoleName() == "" || userlogin.GetCountry() == 0 || userlogin.GetCountry() > 3 {
// 		returnmsg.Msg = cfg.GetErrorCodeNumber("PARAMETER_EMPTY")
// 		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
// 		return
// 	}

// 	var roleinfo account.P_RoleInfo
// 	if err := db.FindOneBson(db.UserTable, &roleinfo, bson.D{primitive.E{Key: "rolename", Value: userlogin.GetRoleName()}}); err == nil &&
// 		roleinfo.GetRoleName() == userlogin.GetRoleName() {
// 		returnmsg.Msg = cfg.GetErrorCodeNumber("RoleNameExists")
// 		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
// 		return
// 	}

// 	createRoleInfo(userlogin.GetRoleName(), userlogin.GetCountry(), c.roleid)
// 	// //角色游戏信息
// 	// db.InsertOne(db.UserTable, &account.P_RoleInfo{
// 	// 	RoleID:        roleid,
// 	// 	RoleName:      userlogin.GetRoleName(),
// 	// 	Country:       userlogin.GetCountry(),
// 	// 	Level:         0,
// 	// 	TesourcesList: map[int32]int32{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0},
// 	// })
// 	// //创建部队信息
// 	// for i := 0; i < 5; i++ {
// 	// 	db.InsertOne(db.TroopsTable, commonstruct.TroopsStruct{
// 	// 		TroopsID:   db.GetAutoID(db.TroopsTable),
// 	// 		Country:    userlogin.GetCountry(),
// 	// 		AreasIndex: cfg.GetCountryAreasIndex(userlogin.GetCountry()),
// 	// 		Type:       1,
// 	// 		Number:     100,
// 	// 		Roleid:     roleid})
// 	// }
// 	// //更新账号信息里角色id
// 	// db.Update(db.AccountTable,
// 	// 	bson.D{primitive.E{Key: "accountid", Value: c.accountid}},
// 	// 	bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "roleid", Value: roleid}}}},
// 	// )

// 	//登陆成功
// 	c.hookLogin(userlogin.GetRoleName(), userlogin.GetCountry(), &roleinfo)

// 	returnmsg.Success = true
// 	returnmsg.Roleid = c.roleid
// 	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)

// }

// //登陆成功后需要发给客户端的信息
// func (c *Client) hookLogin(rolename string, country int32, roleinfo *account.P_RoleInfo) {
// 	//登陆成功后注册进程
// 	if process.IsRegister(c.roleid) {
// 		roleLoginchan := make(chan string)
// 		process.SendMsg(c.roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgRoleLogin, Data: roleLoginchan})
// 		process.UnRegister(c.roleid)

// 		select {
// 		case <-roleLoginchan:
// 		case <-time.After(2 * time.Second):
// 		}
// 	}

// 	process.Register(c.roleid, c.msgChan)
// 	c.rolename = rolename
// 	c.country = country
// 	c.roleinfo = roleinfo

// 	//地图区域信息
// 	c.sendAllArease()
// 	//部队信息
// 	c.sendTroopsList()
// 	//战斗基础设置发送至大地图
// 	//bigmapmanage.SendFightSetting(c.roleid, 0, 0, 0, c.roleinfo.Settings.AutoSelectTactics)
// }

// func (c *Client) updateRole(userlogin *account.C2S_UpdateRoleName) {
// 	returnmsg := &account.S2C_UpdateRoleName{Success: false}
// 	filter := bson.D{primitive.E{Key: "rolename", Value: c.rolename}}
// 	updatefilter := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "rolename", Value: userlogin.UpdateName}}}}
// 	if _, err := db.Update(db.UserTable, filter, updatefilter); err != nil {
// 		logrus.Error(err)
// 		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_UpdateRoleName), returnmsg)
// 		return
// 	}

// 	returnmsg.Success = true
// 	returnmsg.Msg = userlogin.UpdateName
// 	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_UpdateRoleName), returnmsg)
// }

// //账号相关
// func (c *Client) roleAddExp(exp int64) {
// 	Nlevel, Nexp := cfg.AddRoleExp(int64(c.roleinfo.Level), c.roleinfo.Exp, exp)
// 	c.roleinfo.Level = int32(Nlevel)
// 	c.roleinfo.Exp = Nexp
// 	returnmsg := &account.S2C_RoleAddExp{AddExp: exp, NewExp: Nexp, NewLevel: Nlevel}
// 	//获取经验通知
// 	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_RoleAddExp), returnmsg)
// }
