package commonstruct

import (
	"time"
)

//AccountInfo 账号信息
type AccountInfo struct {
	Account            string //账号
	Password           string //密码
	Equipment          string //设备信息
	RegistrationSource string //注册来源(平台)
	RegistrationTime   time.Time
	RoleID             int32  //角色id
	RoleUUID           string //角色uuid
	CDK                string
	Settings           map[uint32]string //设置

}
