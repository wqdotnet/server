package commonstruct

//道具
type ItemInfo struct {
	EquipID   uint32 //id
	EquipUUID string //uuid
	Name      string //名称
	Num       uint32 //数量
	Level     uint32 //等级
	Star      uint32 //星级
	TypeID    uint32 //类型
	ChildType uint32 //子类型

	UseLocation uint32           //装备位置
	Attribute   map[uint32]int64 //属性
	Lock        bool             //是否锁定
}

func CreateItem(itemid, num uint32) *ItemInfo {
	return &ItemInfo{
		EquipID: itemid,
		Num:     num,
	}
}
