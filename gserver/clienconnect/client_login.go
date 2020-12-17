package clienconnect

import (
	"server/db"
	"server/gserver/cfg"
	"server/gserver/cservice"
	"server/msgproto/account"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
)

//账号信息
type accountInfo struct {
	Account            string
	Password           string
	Equipment          string //设备信息
	RegistrationSource string //注册来源(平台)
	RegistrationTime   time.Time
	Rolename           string //角色名
}

//module 用户登陆模块
func (c *Client) loginModule(method int32, buf []byte) {
	switch account.MSG_ACCOUNT(method) {
	case account.MSG_ACCOUNT_C2S_Login:
		userlogin := &account.C2S_Login{}
		e := proto.Unmarshal(buf, userlogin)
		if e != nil {
			log.Error(e)
			return
		}
		c.userLogin(userlogin)
	case account.MSG_ACCOUNT_C2S_CreateRole:
		createMsg := &account.C2S_CreateRole{}
		e := proto.Unmarshal(buf, createMsg)
		if e != nil {
			log.Error(e)
			return
		}
		c.createRole(createMsg)
	default:
		log.Info("loginModule null methodID:", method)
	}
}

//用户登陆
func (c *Client) userLogin(userlogin *account.C2S_Login) {

	filter := bson.D{
		{"account", userlogin.Account},
		{"password", userlogin.Password},
	}

	var accountinfo accountInfo
	if err := db.FindOneBson(&accountinfo, db.AccountTable, filter); err != nil {
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login),
			&account.S2C_Login{
				Success: false,
				Msg:     cfg.ERROR_AccountNull,
			})
	}

	var userinfo account.P_RoleInfo
	filter = bson.D{{"rolename", accountinfo.Rolename}}
	if err := db.FindOneBson(&userinfo, db.UserTable, filter); err != nil {
		//账号下没角色则创建，角色名与账号相同
		userinfo.RoleName = userlogin.Account
		db.InsertOne(db.UserTable, &userinfo)
	}

	//成功登陆
	c.setLoginStatus()
	//登陆成功后注册进程
	cservice.Register(userinfo.RoleName, c)
	c.username = userinfo.RoleName
	c.account = userlogin.GetAccount()

	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login),
		&account.S2C_Login{
			Success:  true,
			RoleInfo: &userinfo})
}

func (c *Client) createRole(userlogin *account.C2S_CreateRole) {
	returnmsg := &account.S2C_CreateRole{Success: false}

	var accountinfo accountInfo
	if err := db.FindOneBson(&accountinfo, db.AccountTable, bson.D{{"account", userlogin.GetAccount()}}); err == nil &&
		accountinfo.Account == userlogin.GetAccount() {
		returnmsg.Msg = cfg.ERROR_AccountExists
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
		return
	}

	var roleinfo account.P_RoleInfo
	if err := db.FindOneBson(&roleinfo, db.UserTable, bson.D{{"rolename", userlogin.GetRoleName()}}); err == nil &&
		roleinfo.GetRoleName() == userlogin.GetRoleName() {
		returnmsg.Msg = cfg.ERROR_RoleNameExists
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
	}

	initRoleData(userlogin.GetAccount(),
		userlogin.GetPassword(),
		userlogin.GetRoleName())

	returnmsg.Success = true
	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
}

func initRoleData(accountName string, password string, rolename string) {

	//账号信息
	db.InsertOne(db.AccountTable, &accountInfo{
		Account:          accountName,
		Password:         password,
		Rolename:         rolename,
		RegistrationTime: time.Now(),
	})

	//角色游戏信息
	db.InsertOne(db.UserTable, &account.P_RoleInfo{RoleName: rolename, Level: 0})
}
