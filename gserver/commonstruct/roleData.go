package commonstruct

import (
	"server/db"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	roleDataMap sync.Map
)

type RoleData struct {
	Acconut   *AccountInfo  //账号信息
	RoleBase  *RoleBaseInfo //角色基础数据
	RoleItems *RoleItemlist //角色道具
	//宗门
	//好友
}

func GetRoleAllData(roleid int32) *RoleData {
	data, ok := roleDataMap.Load(roleid)
	roledata := &RoleData{}
	if ok {
		info := data.(RoleData)
		roledata = &info
	}

	if roledata.RoleBase == nil {
		db.FindOneBson(db.RoleBaseTable, roledata.RoleBase, bson.D{primitive.E{Key: "roleid", Value: roleid}})
		if roledata.RoleBase != nil {
			roledata.RoleBase.DirtyData = false
		}
	}

	if roledata.RoleItems == nil {
		db.FindOneBson(db.RoleItemsTable, roledata.RoleItems, bson.D{primitive.E{Key: "roleid", Value: roleid}})
		if roledata.RoleItems != nil {
			roledata.RoleItems.DirtyData = false
		}
	}

	return roledata
}

func SaveRoleAllData(roledata *RoleData) {
	if roledata.RoleBase != nil {
		roleDataMap.Store(roledata.RoleBase.RoleID, *roledata)
	}
}

func RangeAllData(fc func(*RoleData) (issave bool)) {
	roleDataMap.Range(func(key, value interface{}) bool {
		data := value.(RoleData)
		if fc(&data) {
			roleDataMap.Store(key, data)
		}
		return true
	})
}
