package bigmapmanage

import (
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/gserver/process"
	"server/msgproto/common"

	log "github.com/sirupsen/logrus"
)

//GetBigMapTroopsInfo 获取大地图中的部队信息
func GetBigMapTroopsInfo(troopsID int32) (commonstruct.TroopsStruct, bool) {
	if value, ok := troopsSMap.Load(troopsID); ok {
		troops := value.(commonstruct.TroopsStruct)
		return troops, true
	}
	return commonstruct.TroopsStruct{}, false
}

// 部队进入大地图
func handleTroopsEnterBigMap(troops commonstruct.TroopsStruct) {
	if data, ok := troopsSMap.Load(troops.TroopsID); ok {
		//log.Debugf(" 大地图中已有部队 (改变部队移动方向) TroopsID:[%v]", troops)
		oldtroops := data.(commonstruct.TroopsStruct)

		//战斗状态下不可以接受移动命令
		if oldtroops.State != common.TroopsState_fight {
			oldtroops.ArrivalTime = troops.ArrivalTime
			oldtroops.AreasList = troops.AreasList
			oldtroops.MoveNum = 0
			oldtroops.State = common.TroopsState_Move
			troopsSMap.Store(oldtroops.TroopsID, oldtroops)
		}
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
		troops.MoveNum = 0
		troopsSMap.Store(key, troops)
		return
	}

	//查看上次移动时间，判断本次是否可移动
	var stepsecond int64 = 3
	//log.Infof("[%v]  [%v]  [%v]", troops.MoveStamp, unix, troops.MoveStamp+stepsecond > unix)
	if unix < (troops.MoveStamp + stepsecond) {
		return
	}

	if troops.AreasList == nil || len(troops.AreasList) == 0 {
		log.Warn("bigmap move AreasList 空列表 roleid:", troops.Roleid)
		troops.State = common.TroopsState_Stationed
		troopsSMap.Store(key, troops)
		return
	}

	//移动判定
	if int(troops.MoveNum) >= len(troops.AreasList) {
		log.Infof("移动超过路径(已达终点) [%v]  [%v]", troops.MoveNum, troops.AreasList)
	} else {
		//前进一格
		nextAreas := troops.AreasList[troops.MoveNum]
		//验证是否合法
		if cfg.AreasIsBeside(troops.AreasIndex, nextAreas) {
			troops.AreasIndex = nextAreas
			troops.MoveNum++
		} else {
			troops.State = common.TroopsState_Pause
			log.Warnf("路径不合法  [%v]  ", troops)
		}
	}

	// var num int = 0
	// for i, v := range troops.AreasList {
	// 	if v == troops.AreasIndex {
	// 		//log.Infof("Index [%v]  [%v] ", troops.AreasIndex, i)
	// 		num = i + 1
	// 		break
	// 	}
	// }
	// if num >= len(troops.AreasList) {
	// 	log.Warnf("格子计算有误: %v    num:%v", troops, num)
	// } else {
	// 	//前进一格
	// 	troops.AreasIndex = troops.AreasList[num]
	// }
	//if num +2 >= len

	//走到终点了 或者 暂停
	if int(troops.MoveNum) >= len(troops.AreasList) || troops.State == common.TroopsState_Pause {
		troops.State = common.TroopsState_Stationed
		//troops.MoveStamp = 0
		troops.AreasList = troops.AreasList[0:0]
		troops.MoveNum = 0
		process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "OverMove", Data: troops})
	} else {
		//正常移动中 发送至客户端
		process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "TroopsMove", Data: troops})
	}

	troops.MoveStamp = unix

	//此点是否需要战斗判定
	if areasinfo := getAreasInfo(troops.AreasIndex); areasinfo != nil {
		//中立区域 并且 不属于自己国家地盘的 触发战斗
		if areasinfo.Type == 0 && areasinfo.Occupy != troops.Country {
			troops.State = common.TroopsState_fight //进入战斗状态
			troops.MoveNum = 0
			process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "OnFitht", Data: troops})
			fightAreas(*areasinfo, troops)
		}
	} else {
		log.Warn("bigmap nil areas  index:", troops.AreasIndex)
	}
	troopsSMap.Store(key, troops)
}
