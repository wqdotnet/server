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
	AccountID          int32 //账号id
	Password           string
	Equipment          string //设备信息
	RegistrationSource string //注册来源(平台)
	RegistrationTime   time.Time
	RoleID             int32 //角色id

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
	case account.MSG_ACCOUNT_C2S_UpdateRoleName:
		upName := &account.C2S_UpdateRoleName{}
		e := proto.Unmarshal(buf, upName)
		if e != nil {
			log.Error(e)
			return
		}
		c.updateRole(upName)

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

	if userlogin.GetAccount() == "" || userlogin.Password == "" {
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login),
			&account.S2C_Login{Success: false, Msg: cfg.ERROR_PARAMETER_EMPTY})
	}

	log.Debugf("login %v %v", userlogin.Account, userlogin.Password)
	var accountinfo accountInfo
	if err := db.FindOneBson(db.AccountTable, &accountinfo, filter); err != nil {
		accountid := db.GetAutoID(db.AccountTable)
		c.accountid = accountid
		//创建账号
		db.InsertOne(db.AccountTable, &accountInfo{
			AccountID:        accountid,
			Account:          userlogin.Account,
			Password:         userlogin.Password,
			RegistrationTime: time.Now(),
		})
	} else {
		c.accountid = accountinfo.AccountID
	}
	//连接状态
	c.setLoginStatus()

	var userinfo account.P_RoleInfo
	filter = bson.D{{"roleid", accountinfo.RoleID}}
	if err := db.FindOneBson(db.UserTable, &userinfo, filter); err != nil {
		log.Debug("未找到 角色ID:", accountinfo.RoleID)
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login), &account.S2C_Login{Success: true})
		return
	}

	//登陆成功后注册进程
	cservice.Register(userinfo.RoleName, c)
	c.rolename = userinfo.RoleName
	c.roleid = userinfo.RoleID
	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_Login), &account.S2C_Login{Success: true, RoleInfo: &userinfo})
}

func (c *Client) createRole(userlogin *account.C2S_CreateRole) {
	returnmsg := &account.S2C_CreateRole{Success: false}

	if userlogin.GetRoleName() == "" {
		returnmsg.Msg = cfg.ERROR_PARAMETER_EMPTY
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
		return
	}

	var roleinfo account.P_RoleInfo
	if err := db.FindOneBson(db.UserTable, &roleinfo, bson.D{{"rolename", userlogin.GetRoleName()}}); err == nil &&
		roleinfo.GetRoleName() == userlogin.GetRoleName() {
		returnmsg.Msg = cfg.ERROR_RoleNameExists
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
	}

	roleid := db.GetAutoID(db.UserTable)
	//角色游戏信息
	db.InsertOne(db.UserTable, &account.P_RoleInfo{
		RoleID:   roleid,
		RoleName: userlogin.GetRoleName(),
		Country:  userlogin.GetCountry(),
		Level:    0,
	})

	//更新账号信息里角色id
	db.Update(db.AccountTable,
		bson.D{{"accountid", c.accountid}},
		bson.D{{"$set", bson.D{{"roleid", roleid}}}},
	)

	log.Info("update [%v] [%v]", c.accountid, roleid)

	//登陆成功后注册进程
	cservice.Register(userlogin.GetRoleName(), c)
	c.rolename = userlogin.GetRoleName()
	c.roleid = roleid

	returnmsg.Success = true
	returnmsg.Roleid = roleid
	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_CreateRole), returnmsg)
}

func (c *Client) updateRole(userlogin *account.C2S_UpdateRoleName) {
	returnmsg := &account.S2C_UpdateRoleName{Success: false}
	filter := bson.D{{"rolename", c.rolename}}
	updatefilter := bson.D{{"$set", bson.D{{"rolename", userlogin.UpdateName}}}}
	if _, err := db.Update(db.UserTable, filter, updatefilter); err != nil {
		log.Error(err)
		c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_UpdateRoleName), returnmsg)
	}

	returnmsg.Success = true
	returnmsg.Msg = userlogin.UpdateName
	c.Send(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_S2C_UpdateRoleName), returnmsg)
}
