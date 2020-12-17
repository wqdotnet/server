package clienconnect

import (
	"server/msgproto/account"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

//module:100 用户登陆模块
func (c *Client) loginModule(method int32, buf []byte) {
	switch account.MSG_ACCOUNT(method) {
	case account.MSG_ACCOUNT_Login:

		userlogin := &account.C2S_Login{}
		e := proto.Unmarshal(buf, userlogin)
		if e != nil {
			log.Info(e)
			return
		}
		c.userLogin(userlogin)
	default:
		log.Info("loginModule null methodID:", method)
	}
}

//用户登陆
func (c *Client) userLogin(userlogin *account.C2S_Login) {
	// c.username = userlogin.UserName

	//登陆成功后注册进程
	// cservice.Register(userlogin.UserName, c)

	// c.Send(1, 1, &msg.UserLoginToc{
	// 	Success: true,
	// 	Msg:     "ok",
	// })
}
