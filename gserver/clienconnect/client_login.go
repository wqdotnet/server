package clienconnect

import (
	"math/rand"
	"server/db"
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/gserver/process"
	"server/msgproto/account"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//module 用户登陆模块
func (c *Client) loginModule(method int32, buf []byte) {
	switch account.MSG_ACCOUNT(method) {
	case account.MSG_ACCOUNT_C2S_Login:
		userlogin := &account.C2S_Login{}
		if e := proto.Unmarshal(buf, userlogin); e != nil {
			log.Error(e)
			return
		}
		c.userLogin(userlogin)
	case account.MSG_ACCOUNT_C2S_CreateRole:
		createMsg := &account.C2S_CreateRole{}
		if e := proto.Unmarshal(buf, createMsg); e != nil {
			log.Error(e)
			return
		}
		c.createRole(createMsg)
	case account.MSG_ACCOUNT_C2S_UpdateRoleName:
		upName := &account.C2S_UpdateRoleName{}
		if e := proto.Unmarshal(buf, upName); e != nil {
			log.Error(e)
			return
		}
		c.updateRole(upName)

	default:
		log.Info("loginModule null methodID:", method)
	}
}

func (c *Client) unmarshalExec(b []byte, m protoreflect.ProtoMessage, exec func(m protoreflect.ProtoMessage)) {
	e := proto.Unmarshal(b, m)
	if e != nil {
		log.Error(e)
		return
	}
	exec(m)
}

//用户登陆
func (c *Client) userLogin(userlogin *account.C2S_Login) {
	filter := bson.D{
		primitive.E{Key: "account", Value: userlogin.Account},
		primitive.E{Key: "password", Value: userlogin.Password},
	}

	if userlogin.GetAccount() == "" || userlogin.GetPassword() == "" {
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login),
			&account.S2C_Login{Success: false, Msg: cfg.ERROR_PARAMETER_EMPTY})
	}

	log.Debugf("login %v %v", userlogin.Account, userlogin.Password)
	accountinfo := &commonstruct.AccountInfoStruct{}
	if err := db.FindOneBson(db.AccountTable, accountinfo, filter); err != nil {
		accountid := db.GetAutoID(db.AccountTable)
		roleid := db.GetAutoID(db.UserTable)
		c.account = userlogin.Account
		c.accountid = accountid
		c.roleid = roleid

		//创建账号
		accountinfo.AccountID = accountid
		accountinfo.Account = userlogin.Account
		accountinfo.Password = userlogin.Password
		accountinfo.RoleID = roleid
		accountinfo.RegistrationTime = time.Now()
		db.InsertOne(db.AccountTable, accountinfo)

		//创建角色
		country := int32(rand.Intn(3) + 1)
		createRoleInfo(userlogin.Account, country, roleid)

	} else {
		c.accountid = accountinfo.AccountID
		c.account = accountinfo.Account
		c.roleid = accountinfo.RoleID
	}
	//设置连接状态为已登陆
	c.setLoginStatus()

	var userinfo account.P_RoleInfo
	filter = bson.D{primitive.E{Key: "roleid", Value: accountinfo.RoleID}}
	if err := db.FindOneBson(db.UserTable, &userinfo, filter); err != nil {
		log.Debug("未找到 角色ID:", accountinfo.RoleID)
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login), &account.S2C_Login{Success: true})
		return
	}

	//登陆成功
	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login), &account.S2C_Login{Success: true, RoleInfo: &userinfo})
	c.hookLogin(userinfo.RoleName, userinfo.Country)
}

//创建账号  及  角色信息
func createRoleInfo(rolename string, country, roleid int32) {
	//角色游戏信息
	db.InsertOne(db.UserTable, &account.P_RoleInfo{
		RoleID:        roleid,
		RoleName:      rolename,
		Country:       country,
		Level:         0,
		TesourcesList: map[int32]int32{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0},
	})

	//创建部队信息
	for i := 0; i < 5; i++ {
		db.InsertOne(db.TroopsTable, commonstruct.TroopsStruct{
			TroopsID:   db.GetAutoID(db.TroopsTable),
			Country:    country,
			AreasIndex: cfg.GetCountryAreasIndex(country),
			Type:       1,
			Number:     100,
			Roleid:     roleid})
	}
}

func (c *Client) createRole(userlogin *account.C2S_CreateRole) {
	returnmsg := &account.S2C_CreateRole{Success: false}
	if userlogin.GetRoleName() == "" || userlogin.GetCountry() == 0 || userlogin.GetCountry() > 3 {
		returnmsg.Msg = cfg.ERROR_PARAMETER_EMPTY
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
		return
	}

	var roleinfo account.P_RoleInfo
	if err := db.FindOneBson(db.UserTable, &roleinfo, bson.D{primitive.E{Key: "rolename", Value: userlogin.GetRoleName()}}); err == nil &&
		roleinfo.GetRoleName() == userlogin.GetRoleName() {
		returnmsg.Msg = cfg.ERROR_RoleNameExists
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
		return
	}

	createRoleInfo(userlogin.GetRoleName(), userlogin.GetCountry(), c.roleid)
	// //角色游戏信息
	// db.InsertOne(db.UserTable, &account.P_RoleInfo{
	// 	RoleID:        roleid,
	// 	RoleName:      userlogin.GetRoleName(),
	// 	Country:       userlogin.GetCountry(),
	// 	Level:         0,
	// 	TesourcesList: map[int32]int32{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0},
	// })
	// //创建部队信息
	// for i := 0; i < 5; i++ {
	// 	db.InsertOne(db.TroopsTable, commonstruct.TroopsStruct{
	// 		TroopsID:   db.GetAutoID(db.TroopsTable),
	// 		Country:    userlogin.GetCountry(),
	// 		AreasIndex: cfg.GetCountryAreasIndex(userlogin.GetCountry()),
	// 		Type:       1,
	// 		Number:     100,
	// 		Roleid:     roleid})
	// }
	// //更新账号信息里角色id
	// db.Update(db.AccountTable,
	// 	bson.D{primitive.E{Key: "accountid", Value: c.accountid}},
	// 	bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "roleid", Value: roleid}}}},
	// )

	//登陆成功
	c.hookLogin(userlogin.GetRoleName(), userlogin.GetCountry())

	returnmsg.Success = true
	returnmsg.Roleid = c.roleid
	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)

}

//登陆成功后需要发给客户端的信息
func (c *Client) hookLogin(rolename string, country int32) {
	//登陆成功后注册进程
	if process.IsRegister(c.roleid) {
		roleLoginchan := make(chan string)
		process.SendMsg(c.roleid, commonstruct.ProcessMsg{MsgType: "RoleLogin", Data: roleLoginchan})
		process.UnRegister(c.roleid)

		select {
		case <-roleLoginchan:
		case <-time.After(2 * time.Second):
		}
	}

	process.Register(c.roleid, c.msgChan)
	c.rolename = rolename
	c.country = country

	//地图区域信息
	c.sendAllArease()
	//部队信息
	c.sendTroopsList()
}

func (c *Client) updateRole(userlogin *account.C2S_UpdateRoleName) {
	returnmsg := &account.S2C_UpdateRoleName{Success: false}
	filter := bson.D{primitive.E{Key: "rolename", Value: c.rolename}}
	updatefilter := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "rolename", Value: userlogin.UpdateName}}}}
	if _, err := db.Update(db.UserTable, filter, updatefilter); err != nil {
		log.Error(err)
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_UpdateRoleName), returnmsg)
	}

	returnmsg.Success = true
	returnmsg.Msg = userlogin.UpdateName
	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_UpdateRoleName), returnmsg)
}
