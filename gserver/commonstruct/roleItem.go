package commonstruct

import (
	"server/gserver/cfg"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleItemlist struct {
	RoleID   int32
	ItemList map[string]*ItemInfo

	DirtyDataRecord
}

func (r *RoleItemlist) SetDirtyData(fieldNames ...primitive.E) {
	if len(fieldNames) == 0 {
		r.DirtyData = true
	} else {
		for _, v := range fieldNames {
			r.DirtyDataList[v.Key] = v
		}
	}
}

//道具
type ItemInfo struct {
	ItemUUID  string //uuid
	ID        uint32
	Name      string //名称
	Num       uint32 //数量
	Level     uint32 //等级
	Star      uint32 //星级
	Type      uint32 //类型
	ChildType uint32 //子类型

	UseLocation uint32           //装备位置
	Attribute   map[uint32]int64 //属性
	Lock        bool             //是否锁定
}

func CreateItem(id, num uint32) *ItemInfo {
	itemcfg := cfg.GetItemCfg(id)
	uid, _ := uuid.NewRandom()
	return &ItemInfo{
		ItemUUID:    uid.String(),
		ID:          itemcfg.ID,
		Num:         num,
		Name:        itemcfg.Name,
		Level:       itemcfg.Level,
		Star:        0,
		Type:        itemcfg.Type,
		ChildType:   0,
		UseLocation: 0,
		Attribute:   map[uint32]int64{},
		Lock:        false,
	}
}

//新增道具
func (r *RoleItemlist) AddItem(id, num uint32) {

}

//是否存在N个道具
func (r *RoleItemlist) CheckItem(id, num uint32) bool {
	itemnum := 0
	for _, ii := range r.ItemList {
		if ii.ID == id {
			itemnum += int(ii.Num)
			if itemnum >= int(num) {
				return true
			}
		}
	}
	return false
}

func (r *RoleItemlist) CheckItemList(ids, nums []uint32) bool {
	for i, v := range ids {
		if !r.CheckItem(v, nums[i]) {
			return false
		}
	}
	return true
}

func (r *RoleItemlist) CheckItemUUID(uuid string, num uint32) bool {
	if info, ok := r.ItemList[uuid]; ok {
		return info.Num >= num
	}
	return false
}

func (r *RoleItemlist) CheckItemListUUID(uuids []string, nums []uint32) bool {
	for i, uuid := range uuids {
		if !r.CheckItemUUID(uuid, nums[i]) {
			return false
		}
	}
	return true
}

//扣除道具
func (r *RoleItemlist) DeleteItem(id, num uint32) bool {
	for k, ii := range r.ItemList {
		if ii.ID == id {
			if ii.Num >= num {
				ii.Num -= num
				if ii.Num == 0 {
					delete(r.ItemList, k)
				}
				return true
			} else {
				num -= ii.Num
				delete(r.ItemList, k)
			}

			if num == 0 {
				return true
			}
		}
	}
	return false
}

func (r *RoleItemlist) DeleteItemList(ids, nums []uint32) bool {
	if !r.CheckItemList(ids, nums) {
		return false
	}

	for i, v := range ids {
		r.DeleteItem(v, nums[i])
	}
	return true
}

func (r *RoleItemlist) DeleteItemUUID(uuid string, num uint32) bool {
	if info, ok := r.ItemList[uuid]; ok {
		if info.Num >= num {
			info.Num -= num
			if info.Num == 0 {
				delete(r.ItemList, uuid)
			}
			return true
		}
	}
	return false
}

func (r *RoleItemlist) DeleteItemListUUID(uuids []string, nums []uint32) bool {
	if !r.CheckItemListUUID(uuids, nums) {
		return false
	}

	for i, v := range uuids {
		r.DeleteItemUUID(v, nums[i])
	}
	return true
}
