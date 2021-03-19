package commonstruct

//FightSetting 战斗设置记录
type FightSetting struct {
	AutoSelect bool
	//技能（战法）选择
	SkillID int32
	//战术选择
	TacticsID  int32
	RoleID     int32
	TroopsID   int32
	SelectTime int64
}

// //GetTroopsSet 部队本轮是否设置过 0 没设置过 >0 设置时间
// func (set *FightSetting) GetTroopsSet(troopsid int32) int64 {
// 	if set.TroopsID == troopsid && set.SelectTime  {

// 	}

// 	return 0
// }

//TroopsExitBigmap 部队退出地图chan结构
type TroopsExitBigmap struct {
	TroopsID int32
	Type     int32 //退出类型 1 下阵
	Value    int32 //需要更新的值
}
