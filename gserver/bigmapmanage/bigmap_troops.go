package bigmapmanage

import (
	"server/gserver/commonstruct"
	"server/msgproto/common"

	"server/gserver/process"

	log "github.com/sirupsen/logrus"
)

//GetTroopsInfo 获取大地图中的部队信息
func GetTroopsInfo(troopsID int32) (*commonstruct.TroopsStruct, bool) {
	if value, ok := troopsSMap.Load(troopsID); ok {
		troops := value.(commonstruct.TroopsStruct)
		return &troops, true
	}
	return nil, false
}

// 部队进入大地图 (v2 考虑是否直接调用 sync.map 进行存储，而不是通过chan传递数据)
func handleTroopsEnterBigMap(troops commonstruct.TroopsStruct) {
	if _, ok := troopsSMap.Load(troops.TroopsID); ok {
		log.Warnf(" 大地图中已有此部队 TroopsID:[%v]", troops.TroopsID)
		return
	}
	//log.Info(troops)
	troopsSMap.Store(troops.TroopsID, troops)
}

//接收到部队暂停命令
func handleTroopsStopMove(troopid int32) {
	if value, ok := troopsSMap.Load(troopid); ok {
		troops := value.(commonstruct.TroopsStruct)
		troops.State = common.TroopsState_Pause
		troopsSMap.Store(troopid, troops)
	}
}

//部队移动
//bigmap进程 loop 每秒执行一次,进行部队移动判断
//v2 考虑是否每个移动部队 单起协程处理
func loopTroopsMove(key interface{}, troops commonstruct.TroopsStruct, unix int64) {
	//第一步
	if troops.MoveStamp == 0 {
		troops.MoveStamp = unix
		troopsSMap.Store(key, troops)
		return
	}

	//查看上次移动时间，判断本次是否可移动
	var stepsecond int64 = 3
	//log.Infof("[%v]  [%v]  [%v]", troops.MoveStamp, unix, troops.MoveStamp+stepsecond > unix)
	if troops.MoveStamp+stepsecond > unix {
		return
	}

	if troops.AreasList == nil || len(troops.AreasList) == 0 {
		log.Warn("bigmap move AreasList 空列表 roleid:", troops.Roleid)
		troops.State = common.TroopsState_Stationed
		troopsSMap.Store(key, troops)
		return
	}

	var num int = 0
	for i, v := range troops.AreasList {
		if v == troops.AreasIndex {
			//log.Infof("Index [%v]  [%v] ", troops.AreasIndex, i)
			num = i + 1
			break
		}
	}
	//前进一格
	troops.AreasIndex = troops.AreasList[num]
	//发送至客户端
	process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "TroopsMove", Data: troops})

	//走到终点了 或者 暂停
	if num+2 > len(troops.AreasList) || troops.State == common.TroopsState_Pause {
		troops.State = common.TroopsState_Stationed
		troops.MoveStamp = 0
		troops.AreasList = troops.AreasList[0:0]
		//troopsSMap.Store(key, troops)
		process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "OverMove", Data: troops})
	} else {
		troops.MoveStamp = unix
		//troopsSMap.Store(key, troops)
	}

	//此点是否需要战斗判定
	if areasinfo := GetAreasInfo(troops.AreasIndex); areasinfo != nil {
		//中立区域 并且 不属于自己国家地盘的 触发战斗
		if areasinfo.Type == 0 && areasinfo.Occupy != troops.Country {
			troops.State = common.TroopsState_fight //进入战斗状态
			process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "OnFitht", Data: troops})
		}
	} else {
		log.Warn("bigmap nil areas  index:", troops.AreasIndex)
	}

	troopsSMap.Store(key, troops)
}
