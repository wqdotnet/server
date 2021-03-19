package bigmapmanage

import (
	"slgserver/db"
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/gserver/process"
	"slgserver/msgproto/bigmap"
	"slgserver/msgproto/common"
	"slgserver/msgproto/fight"
	"slgserver/msgproto/troops"
	"slgserver/tool"

	log "github.com/sirupsen/logrus"
)

//-----------------------------地图区域-----------------------------

//AreasInfo 区域信息
type AreasInfo struct {
	AreasIndex int32 //区域ID
	Type       int32 //0 中立区域 1-3:国家首都
	State      int32 //0 正常 1 战斗
	Occupy     int32 //占领信息  0 :无人占领  1-3国家KEY

	AttackQueue  []int32 //攻方队列
	DefenseQueue []int32 //守方队列

	AttackQueueNum  int32 //队列编号
	DefenseQueueNum int32

	TroopsNum map[int32]int32 //各方部队数量

	MsgSubscribe map[int32]int32 //消息订阅  roleid

	PushNum int //队列数据推送数量

	BaseStamp int64 //下轮战斗基础等待时长
	TimeStamp int64 //下轮战斗时间戳

}

func newAreasInfo(index, arestype int32) AreasInfo {
	return AreasInfo{AreasIndex: index,
		Type:         arestype,
		State:        0,
		Occupy:       arestype,
		TroopsNum:    make(map[int32]int32),
		MsgSubscribe: make(map[int32]int32),
		DefenseQueue: []int32{},
		AttackQueue:  []int32{},
		PushNum:      2,
	}
}

//计算下一回合战斗的时间戳
// 1.是否自动战斗 ->  2.是否有主动技能 -> 3.战术选择等待时长
func (areas *AreasInfo) refreshFightTimeStamp(attackTroops, defenseTroops *commonstruct.TroopsStruct) int64 {
	attacktime := refreshFightTimeStamp(attackTroops)
	defensetime := refreshFightTimeStamp(defenseTroops)

	if attacktime > defensetime {
		return attacktime
	}

	return defensetime
}

func refreshFightTimeStamp(attackTroops *commonstruct.TroopsStruct) int64 {
	// // ===========选择时长 ===============
	//自动战斗 等待时间
	if attackTroops.FightSet.AutoSelect {
		return cfg.GetGlobalInt64("autoWaitTime")
	}

	if attackTroops.FightSet.SkillID != 0 || attackTroops.FightSet.TacticsID != 0 {
		return cfg.GetGlobalInt64("autoWaitTime")
	}

	if !attackTroops.RoundWins && attackTroops.SkillCD() == 0 {
		return cfg.GetGlobalInt64("skillWaitTime")
	}

	//战术选择等待时间
	return cfg.GetGlobalInt64("tacticsWaitTime")
}

func (areas *AreasInfo) refreshbaseWaitTime(record *fight.S2C_FightRecordPush, attackTroops, defenseTroops *commonstruct.TroopsStruct) {
	//固定时长
	// 移动播放时长			moveTime	0.5
	basetime := cfg.GetGlobalfloat("moveTime")

	// 平砍每刀时长			cutTime	0.3
	attacktime := cfg.GetGlobalfloat("cutTime")
	//放技能时长
	// 施法ui过场动画时长	uiTime	1.5      放技能
	// 主动释放动画时长		activeTime	1.5  技能时长
	// 被动释放动画时长		passiveTime	1.5
	skilltime := cfg.GetGlobalfloat("uiTime") + cfg.GetGlobalfloat("activeTime")
	// 斩杀全队庆祝时长		victoryTime	1  //下队入场时长
	// 下队入场等待时长		waitTime	5
	nexttime := cfg.GetGlobalfloat("victoryTime") + cfg.GetGlobalfloat("waitTime")

	if record != nil {
		for _, v := range record.Record {
			//0 普攻 1 战法（技能） 2 战术
			switch v.Type {
			case 0:
				basetime += attacktime * float32(len(v.Defense))
			case 1:
				basetime += skilltime
			case 2:
				basetime += skilltime
			default:
			}
		}
	}

	//有队伍死亡 下队进场
	_, atthp := getHPfun(attackTroops.RowHP)
	_, defhp := getHPfun(defenseTroops.RowHP)
	if atthp == 0 || defhp == 0 {
		basetime += nexttime
	}

	areas.BaseStamp = int64(tool.Round(float64(basetime)))
}

//部队经过该区域 触发战斗
//1.中立区/敌国区域  触发区域npc战斗
//2.有敌国部队  触发国战
func (areas *AreasInfo) troopsTriggerBattle(troops *commonstruct.TroopsStruct) bool {
	//中立区域 并且 不属于自己国家地盘的 触发战斗 || 区域正在战斗
	if areas.Type == 0 && areas.Occupy != troops.Country || areas.State == 1 {
		troops.State = common.TroopsState_fight //进入战斗状态
		troops.MoveNum = 0
		//清除之前战斗记录
		troops.SkillUseNumber = 0
		troops.KillNumber = 0
		troops.RoundWins = false

		process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgOnFitht, Data: *troops})
		if areas.State == 0 {
			areas.State = 1
			//无守方部队时 创建守方npc 或者中立npc
			if len(areas.DefenseQueue) == 0 {

				npc := commonstruct.NewTroops("守军NPC", db.GetAutoID(db.TroopsTable), 3, areas.Occupy, 1)
				npc.AreasIndex = areas.AreasIndex
				npc.Attack = 80
				npc.State = common.TroopsState_fight
				npc.FitghtState = 1
				npc.Level = 3
				npc.FightSet.AutoSelect = true
				npc.CalculationAttribute()
				areas.DefenseQueue = append(areas.DefenseQueue, npc.TroopsID)
				saveMapTroops(npc)

			}

			log.Info("区域状态变化通知所有人:", areas.AreasIndex, areas.Occupy, 1)
			//区域状态变化通知所有人
			sendAllRolesAreasStateChange(areas.AreasIndex, areas.Occupy, 1)
		}
		return true
	}

	return false
}

//部队正常移动进入区域
//进入区域战斗队列
func (areas *AreasInfo) entryAreasInfo(troops *commonstruct.TroopsStruct) {
	//log.Infof("部队进入区域：%v  %v", areas.AreasIndex, areas.State)

	isrepeat := false
	selectitem := func(list []int32, indexnum int32) []int32 {
		for _, v := range list {
			if v == troops.TroopsID {
				log.Warn("entryAreasInfo:[重复进入] ", troops)
				isrepeat = true
				return list
			}
		}
		troops.QueueNum = indexnum
		return append(list, troops.TroopsID)
	}

	//战斗队列
	if areas.Occupy == troops.Country {
		areas.AttackQueueNum++
		areas.DefenseQueue = selectitem(areas.DefenseQueue, areas.AttackQueueNum)
	} else {
		areas.DefenseQueueNum++
		areas.AttackQueue = selectitem(areas.AttackQueue, areas.DefenseQueueNum)
	}

	//重复进入
	if isrepeat {
		return
	}

	if num, ok := areas.TroopsNum[troops.Country]; ok {
		areas.TroopsNum[troops.Country] = num + 1
	} else {
		areas.TroopsNum[troops.Country] = 1
	}

	//战斗状态
	if areas.State == 1 {
		//部队进入队列 发送消息给订阅用户
		ptroops := troops.ConvertTroopsProto()
		for _, v := range areas.MsgSubscribe {
			process.SendSocketMsg(v,
				int32(bigmap.MSG_BIGMAP_Module_BIGMAP),
				int32(bigmap.MSG_BIGMAP_S2C_EntryQueue),
				&bigmap.S2C_EntryQueue{AreasIndex: areas.AreasIndex, Troops: ptroops})
		}
		areas.refreshTop2TroopsFightState()
	}
	//log.Infof("区域：%v", areas)
}

//部队离开该区域
func (areas *AreasInfo) leaveArea(troopsid, country int32) {
	deleteitem := func(list []int32) []int32 {
		for index, v := range list {
			if v == troopsid {
				return append(list[:index], list[index+1:]...)
			}
		}
		return list
	}

	if num, ok := areas.TroopsNum[country]; ok {
		areas.TroopsNum[country] = num - 1
	}

	areas.AttackQueue = deleteitem(areas.AttackQueue)
	areas.AttackQueueNum = int32(len(areas.AttackQueue))
	areas.DefenseQueue = deleteitem(areas.DefenseQueue)
	areas.DefenseQueueNum = int32(len(areas.DefenseQueue))

	//部队离开队列 发送消息给订阅用户
	for _, v := range areas.MsgSubscribe {
		process.SendSocketMsg(v,
			int32(bigmap.MSG_BIGMAP_Module_BIGMAP),
			int32(bigmap.MSG_BIGMAP_S2C_LeaveQueue), &bigmap.S2C_LeaveQueue{AreasIndex: areas.AreasIndex, TroopsID: troopsid})
	}
	//log.Infof("部队离开区域：%v", areas)
}

//刷新队列战斗状态
func (areas *AreasInfo) refreshTop2TroopsFightState() {

	var State common.TroopsState = common.TroopsState_Stationed
	if areas.State == 1 {
		State = common.TroopsState_fight
	}

	refreshQueueFun := func(list []int32) []*troops.P_Troops {
		troopslist := []*troops.P_Troops{}
		for k, troopsid := range list {

			troops, ok := GetMapTroopsInfo(troopsid)
			if !ok {
				log.Warn("未找到该支部队:", areas.AreasIndex, troopsid)
				continue
			}

			oldState := troops.State
			oldFightState := troops.FitghtState

			troops.State = State
			troops.FitghtState = 0
			if k < 1 && areas.State == 1 {
				troops.FitghtState = 1
			}

			//log.Info("部队状态对比：", troops.Roleid, oldState, troops.State, oldFightState, troops.FitghtState)

			//状态变化通知客户端
			if oldState != troops.State || oldFightState != troops.FitghtState {
				saveMapTroops(&troops)
				troopslist = append(troopslist, troops.ConvertTroopsProto())
				//npc 不发送
				if troops.Attribute == 0 {
					process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgUpdateTroopsInfo, Data: troops})
				}
			}
			saveMapTroops(&troops)
		}
		return troopslist
	}

	attacktroopslist := refreshQueueFun(areas.AttackQueue)
	defensetroopslist := refreshQueueFun(areas.DefenseQueue)

	//新上阵队伍，把队伍详细信息发给所有订阅用户
	if len(attacktroopslist) != 0 && len(defensetroopslist) != 0 {
		for _, v := range areas.MsgSubscribe {
			process.SendSocketMsg(v,
				int32(bigmap.MSG_BIGMAP_Module_BIGMAP),
				int32(bigmap.MSG_BIGMAP_S2C_QueueTroopsInfo), &bigmap.S2C_QueueTroopsInfo{AttackTroopsList: attacktroopslist, DefenseTroopsList: defensetroopslist})
		}
	}
}

//整理队列  区分出阵营进行战斗
func (areas *AreasInfo) queueSplit() {
	var defenseQueue, attackQueue []int32
	for _, troopsid := range areas.AttackQueue {
		if troops, ok := GetMapTroopsInfo(troopsid); ok {
			if troops.Country == areas.Occupy {
				defenseQueue = append(defenseQueue, troops.TroopsID)
			} else {
				attackQueue = append(attackQueue, troops.TroopsID)
			}
		}
	}

	for _, troopsid := range areas.DefenseQueue {
		if troops, ok := GetMapTroopsInfo(troopsid); ok {
			if troops.Country == areas.Occupy {
				defenseQueue = append(defenseQueue, troops.TroopsID)
			} else {
				attackQueue = append(attackQueue, troops.TroopsID)
			}
		}
	}

	areas.DefenseQueue = defenseQueue
	areas.AttackQueue = attackQueue

}

// 将消息发送给订阅用户
func (areas *AreasInfo) sendQueueInfo(roleid int32) {
	var attacklist, defenselist []*troops.P_Troops
	for k, troopsid := range areas.AttackQueue {
		if troopsinfo, ok := GetMapTroopsInfo(troopsid); ok {
			if k > (areas.PushNum - 1) {
				attacklist = append(attacklist,
					&troops.P_Troops{Name: troopsinfo.Name,
						Level:    troopsinfo.Level,
						Country:  troopsinfo.Country,
						Type:     troopsinfo.Type,
						TroopsID: troopsinfo.TroopsID})
			} else {
				attacklist = append(attacklist, troopsinfo.ConvertTroopsProto())
			}
		}
	}

	for k, troopsid := range areas.DefenseQueue {
		if troopsinfo, ok := GetMapTroopsInfo(troopsid); ok {
			if k > (areas.PushNum - 1) {
				defenselist = append(defenselist,
					&troops.P_Troops{Name: troopsinfo.Name,
						Level:    troopsinfo.Level,
						Country:  troopsinfo.Country,
						Type:     troopsinfo.Type,
						TroopsID: troopsinfo.TroopsID})
			} else {
				defenselist = append(defenselist, troopsinfo.ConvertTroopsProto())
			}
		}
	}
	pbdata := &bigmap.S2C_FightSubscribe{AttackTroopsList: attacklist,
		DefenseTroopsList: defenselist,
		FightStamp:        areas.TimeStamp,
	}

	process.SendSocketMsg(roleid,
		int32(bigmap.MSG_BIGMAP_Module_BIGMAP),
		int32(bigmap.MSG_BIGMAP_S2C_FightSubscribe), pbdata)

	//战斗设置信息
	areas.pushSkillSelect()
}

//===========================================================================================
//用户是否订阅消息
func (areas *AreasInfo) isSubscribe(roleid int32) bool {
	_, ok := areas.MsgSubscribe[roleid]
	return ok
}

//saveAreasInfo save
func saveAreasInfo(info AreasInfo) {
	areasSMap.Store(info.AreasIndex, info)
}

//getAreasInfo 获取区域信息
func getAreasInfo(key int32) *AreasInfo {
	if info, ok := areasSMap.Load(key); ok {
		areas := info.(AreasInfo)
		return &areas
	}
	return nil
}

//---------------------------------------------------------
// 将区域状态变化发送给所有用户
func sendAllRolesAreasStateChange(areasindex, occupy, state int32) {
	//区域状态
	s2cAreasinfo := &bigmap.S2C_AreasInfo{}
	s2cAreasinfo.AreasInfoList = append(s2cAreasinfo.AreasInfoList,
		&bigmap.P_AreasInfo{AreasIndex: areasindex,
			Type:  occupy,
			State: state})
	process.SendAllSocketMsg(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_AreasInfo), s2cAreasinfo)
}

//区域消息订阅/取消
func areasMsgSubscribe(sub bool, areasindex, roleid int32) {
	if areas := getAreasInfo(areasindex); areas != nil {
		//战斗已结束
		if areas.State == 0 {
			process.SendAllSocketMsg(int32(bigmap.MSG_BIGMAP_Module_BIGMAP),
				int32(bigmap.MSG_BIGMAP_S2C_FightSubscribe),
				&bigmap.S2C_FightSubscribe{Msg: cfg.GetErrorCodeNumber("FIGHT_OVER")})
			return
		}

		if sub {
			areas.MsgSubscribe[roleid] = roleid
			areas.sendQueueInfo(roleid)
		} else {
			delete(areas.MsgSubscribe, roleid)
		}
	}
}

//------------------------------外部接口------------------

//AreasRange func
func AreasRange(exefunc func(value AreasInfo) bool) {
	areasSMap.Range(func(key, value interface{}) bool {
		areas := value.(AreasInfo)
		return exefunc(areas)
	})
}

//GetAreasSimple 获取区域简迅
func GetAreasSimple(areasindex int32) map[int32]int32 {
	if areas := getAreasInfo(areasindex); areas != nil {
		return areas.TroopsNum
	}
	return map[int32]int32{}
}
