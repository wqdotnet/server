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
	Acconut   *AccountInfo //账号信息
	RoleBase  RoleBaseInfo //角色基础数据
	RoleItems RoleItemlist //角色道具
	//宗门
	//好友
}

//数据更新标注
type DirtyDataRecord struct {
	TableName     string
	DirtyData     bool
	DirtyDataList map[string]primitive.E
}

func (r *DirtyDataRecord) SetDirtyData(fieldNames ...primitive.E) {
	if len(fieldNames) == 0 {
		r.DirtyData = true
	} else {
		for _, v := range fieldNames {
			r.DirtyDataList[v.Key] = v
		}
	}
}

func GetRoleAllData(roleid int32) *RoleData {
	data, ok := roleDataMap.Load(roleid)
	roledata := &RoleData{}
	if ok {
		info := data.(RoleData)
		roledata = &info
	}

	if roledata.RoleBase.RoleID == 0 {
		rolebase := &RoleBaseInfo{}
		if err := db.FindOneBson(db.RoleBaseTable, rolebase, bson.D{primitive.E{Key: "roleid", Value: roleid}}); err == nil {
			roledata.RoleBase = *rolebase
		}
	}

	if roledata.RoleItems.RoleID == 0 {
		roleitem := &RoleItemlist{}
		if err := db.FindOneBson(db.RoleItemsTable, roleitem, bson.D{primitive.E{Key: "roleid", Value: roleid}}); err == nil {
			roledata.RoleItems = *roleitem
		}
	}

	return roledata
}

//缓存数据
func StoreRoleData(roledata *RoleData) {
	if roledata.RoleBase.RoleID != 0 {
		roleDataMap.Store(roledata.RoleBase.RoleID, *roledata)
	}
}

func RangeAllData(fc func(*RoleData) bool) {
	roleDataMap.Range(func(key, value interface{}) bool {
		data := value.(RoleData)
		if fc(&data) {
			logrus.Debug("db 节点定时更新数据: 角色id:[%v]", data.RoleBase.RoleID)
			roleDataMap.Store(key, data)
		}
		return true
	})
}

func saveDirtyData(RoleID int32, data *DirtyDataRecord, replacement interface{}) (updateDB bool) {
	findfield := bson.D{primitive.E{Key: "roleid", Value: RoleID}}
	if len(data.DirtyDataList) != 0 {
		Upfield := bson.D{}
		for _, primitiveE := range data.DirtyDataList {
			Upfield = append(Upfield, primitiveE)
		}
		db.Update(data.TableName, findfield, bson.D{primitive.E{Key: "$set", Value: Upfield}})
		updateDB = true
		data.DirtyDataList = make(map[string]primitive.E)
	} else if data.DirtyData {
		if _, err := db.ReplaceOne(data.TableName, findfield, replacement); err == nil {
			data.DirtyData = false
			updateDB = true
			data.DirtyDataList = make(map[string]primitive.E)
		}
	}
	return false
}

//保存更新到mongo
func SaveRoleData(rd *RoleData) (updateDB bool) {
	if ok := saveDirtyData(rd.RoleBase.RoleID, &rd.RoleBase.DirtyDataRecord, &rd.RoleBase); ok {
		updateDB = ok
	}

	if ok := saveDirtyData(rd.RoleBase.RoleID, &rd.RoleItems.DirtyDataRecord, &rd.RoleItems); ok {
		updateDB = ok
	}
	//findfield := bson.D{primitive.E{Key: "roleid", Value: rd.RoleBase.RoleID}}
	// if len(rd.RoleBase.DirtyDataList) != 0 {
	// 	Upfield := bson.D{}
	// 	for _, primitiveE := range rd.RoleBase.DirtyDataList {
	// 		Upfield = append(Upfield, primitiveE)
	// 	}
	// 	db.Update(db.RoleBaseTable, findfield, bson.D{primitive.E{Key: "$set", Value: Upfield}})
	// 	updateDB = true
	// } else if rd.RoleBase.DirtyData {
	// 	if _, err := db.ReplaceOne(db.RoleBaseTable, findfield, rd.RoleBase); err == nil {
	// 		rd.RoleBase.DirtyData = false
	// 		updateDB = true
	// 	}
	// }
	// rd.RoleBase.DirtyDataList = make(map[string]primitive.E)

	// if len(rd.RoleItems.DirtyDataList) != 0 {
	// 	Upfield := bson.D{}
	// 	for _, primitiveE := range rd.RoleItems.DirtyDataList {
	// 		Upfield = append(Upfield, primitiveE)
	// 	}
	// 	db.Update(db.RoleItemsTable, findfield, bson.D{primitive.E{Key: "$set", Value: Upfield}})
	// 	updateDB = true
	// } else if rd.RoleItems.DirtyData {
	// 	if _, err := db.ReplaceOne(db.RoleItemsTable, findfield, rd.RoleItems); err == nil {
	// 		rd.RoleItems.DirtyData = false
	// 		updateDB = true
	// 	}
	// }
	// rd.RoleItems.DirtyDataList = make(map[string]primitive.E)

	return updateDB
}
