package commonstruct

import (
	"time"
)

//AccountInfoStruct 账号信息
type AccountInfoStruct struct {
	Account            string
	AccountID          int32 //账号id
	Password           string
	Equipment          string //设备信息
	RegistrationSource string //注册来源(平台)
	RegistrationTime   time.Time
	RoleID             int32 //角色id

}
