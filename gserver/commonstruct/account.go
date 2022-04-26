package commonstruct

import (
	"server/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AccountInfo 账号信息
type AccountInfo struct {
	Account            string //账号
	Password           string //密码
	Equipment          string //设备信息
	RegistrationSource string //注册来源(平台)
	RegistrationTime   int64
	RoleID             int32  //角色id
	RoleUUID           string //角色uuid
	CDK                string
	Settings           map[uint32]string //设置

}

func (accountinfo *AccountInfo) GetAccountinfo() bool {
	filter := bson.D{
		primitive.E{Key: "account", Value: accountinfo.Account},
	}

	if accountinfo.Password != "" {
		filter = append(filter, primitive.E{Key: "password", Value: accountinfo.Password})
	}

	if err := db.FindOneBson(db.AccountTable, accountinfo, filter); err != nil {
		return false
	}

	return true
}

func GetMaxRoleID(value int32) int32 {
	var obj AccountInfo
	//区号
	db.FindFieldMax(db.AccountTable, "roleid", &obj, bson.D{primitive.E{Key: "zoneid", Value: value}})
	return obj.RoleID
}

func GetNewRoleID() int32 {
	return db.RedisINCRBY("MaxRoleID", 1)
}

func GetCDKinfo(cdk string) bool {
	filter := bson.D{
		primitive.E{Key: "cdk", Value: cdk},
	}

	accountinfo := &AccountInfo{}
	if err := db.FindOneBson(db.AccountTable, accountinfo, filter); err != nil {
		return false
	}
	return true
}
