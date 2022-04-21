package commonstruct

import (
	pbrole "server/proto/role"
)

//角色基础数据
type RoleBaseInfo struct {
	ZoneID            int32  //所属区号
	RoleID            int32  //角色id
	Name              string //角色名
	HeadID            uint32 //头像id
	Sex               uint32 //性别
	Level             int32  //等级
	Exp               int64  //经验
	PracticeTimestamp int64  //练功时间戳

	AttributeValue map[uint32]int64         //属性值
	BodyList       map[uint32]*RoleBodyInfo //体质信息
	CE             int64                    //战斗力 combat effectiveness

	EquipList map[int32]*ItemInfo //道具

	OfflineTimestamp int64 //离线时间戳
	Online           bool  //是否在线

	//好友
	//宗门
	//任务
	//成就记录

	DirtyData bool
}

//体质
type RoleBodyInfo struct {
	BodyID         uint32 //体质类型
	BodyLevel      uint32 //体质等级
	PropertiesId   uint32 //属性id
	AttributeValue int64  //属性值
}

//新增道具
func (r *RoleBaseInfo) AddItem() {
}

//是否存在N个道具
func (r *RoleBaseInfo) ExistsItem(itemid, num uint32) bool {
	return false
}

func (r *RoleBaseInfo) ExistsItemList(itemid, num []uint32) bool {
	return false
}

//扣除道具
func (r *RoleBaseInfo) DeleteItem(itemid, num uint32) bool {
	return false
}

func (r *RoleBaseInfo) DeleteItemList(itemid, num []uint32) bool {
	return false
}

func (r *RoleBaseInfo) ToPB() *pbrole.Pb_RoleInfo {
	return &pbrole.Pb_RoleInfo{
		RoleID:         r.RoleID,
		RoleName:       r.Name,
		Level:          r.Level,
		Exp:            r.Exp,
		Sex:            r.Sex,
		AttributeValue: r.AttributeValue,
		CE:             r.CE,
		//BodyList:       r.BodyList,
	}
}

func (r *RoleBaseInfo) SetDirty() {
	r.DirtyData = true
}