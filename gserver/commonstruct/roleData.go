package commonstruct

import (
	"server/db"
	"sync"

	"github.com/sirupsen/logrus"
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
		rolebase := &RoleBaseInfo{}
		if err := db.FindOneBson(db.RoleBaseTable, rolebase, bson.D{primitive.E{Key: "roleid", Value: roleid}}); err == nil {
			roledata.RoleBase = rolebase
			roledata.RoleBase.DirtyData = false
		}
	}

	if roledata.RoleItems == nil {
		roleitem := &RoleItemlist{}
		if err := db.FindOneBson(db.RoleItemsTable, roleitem, bson.D{primitive.E{Key: "roleid", Value: roleid}}); err == nil {
			roledata.RoleItems = roleitem
			roledata.RoleItems.DirtyData = false
		}
	}

	return roledata
}

func StoreRoleAllData(roledata *RoleData) {
	if roledata.RoleBase != nil {
		roleDataMap.Store(roledata.RoleBase.RoleID, *roledata)
	}
}

func RangeAllData(fc func(*RoleData) bool) {
	roleDataMap.Range(func(key, value interface{}) bool {
		data := value.(RoleData)
		if fc(&data) {
			roleDataMap.Store(key, data)
		}
		return true
	})
}

func SaveRoleData(rd *RoleData) (updateDB bool) {
	findfield := bson.D{primitive.E{Key: "roleid", Value: rd.RoleBase.RoleID}}
	if rd.RoleBase.DirtyData {
		if i, err := db.ReplaceOne(db.RoleBaseTable, findfield, rd.RoleBase); err == nil {
			logrus.Debugf("保存数据 %v   %v  %v", i, rd.RoleBase.Name, rd.RoleBase.RoleID)
			rd.RoleBase.DirtyData = false
			updateDB = true
		}
	}

	if rd.RoleItems.DirtyData {
		if i, err := db.ReplaceOne(db.RoleItemsTable, findfield, rd.RoleItems); err == nil {
			logrus.Debugf("保存数据 %v   %v  %v", i, rd.RoleItems, rd.RoleItems.RoleID)
			rd.RoleItems.DirtyData = false
			updateDB = true
		}
	}

	return updateDB
}
