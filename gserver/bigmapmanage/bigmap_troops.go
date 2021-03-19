package bigmapmanage

import (
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/gserver/process"
	"slgserver/msgproto/common"

	log "github.com/sirupsen/logrus"
)

//GetMapTroopsInfo 获取大地图中的部队信息
func GetMapTroopsInfo(troopsID int32) (commonstruct.TroopsStruct, bool) {
	if value, ok := troopsSMap.Load(troopsID); ok {
		troops := value.(commonstruct.TroopsStruct)
		return troops, true
	}
	return commonstruct.TroopsStruct{}, false
}

func saveMapTroops(troops *commonstruct.TroopsStruct) {
	troopsSMap.Store(troops.TroopsID, *troops)
}

// 部队进入大地图
func handleTroopsEnterBigMap(troops commonstruct.TroopsStruct) {
	log.Infof("部队进入大地图: roleid:[%v] troopsid:[%v]", troops.Roleid, troops.TroopsID)

	if data, ok := troopsSMap.Load(troops.TroopsID); ok {
		//log.Debugf(" 大地图中已有部队 (改变部队移动方向) TroopsID:[%v]", troops)
		oldtroops := data.(commonstruct.TroopsStruct)

		//战斗状态下不可以接受移动命令
		if oldtroops.State != common.TroopsState_fight {
			oldtroops.ArrivalTime = troops.ArrivalTime
			oldtroops.AreasList = troops.AreasList
			oldtroops.MoveNum = 0
			oldtroops.State = common.TroopsState_Move
			saveMapTroops(&oldtroops)
		}
		return
	}
	//log.Info(troops)
	saveMapTroops(&troops)
}

//部队退出大地图
func troopsExitBigmap(troopsexit *commonstruct.TroopsExitBigmap) {
	if data, ok := troopsSMap.Load(troopsexit.TroopsID); ok {
		troops := data.(commonstruct.TroopsStruct)
		if troops.State == common.TroopsState_fight {
			log.Warn("部队正在战斗，无法退出大地图")
			return
		}
		cleanBigmapTroops(&troops)
		switch troopsexit.Type {
		case 1:
			troops.StageNumber = 0
		}

		ok := process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgUpdateTroopsInfo, Data: troops})
		if !ok {
			log.Info("玩家没在线:", troops)
			//2.玩家没在线则直接修改数据库里数据
		}
	}
}

//updateTroopsInfo 大地图中部队数据更新
//只更新属性类数据
func updateTroopsInfo(troops *commonstruct.TroopsStruct) {

}

//接收到部队暂停命令
func handleTroopsStopMove(troopid int32) {
	if value, ok := troopsSMap.Load(troopid); ok {
		troops := value.(commonstruct.TroopsStruct)
		troops.State = common.TroopsState_Pause
		saveMapTroops(&troops)
	}
}

//部队离开所在的格子
func troopsLeaveArea(troops *commonstruct.TroopsStruct) {
	if oldAreas := getAreasInfo(troops.AreasIndex); oldAreas != nil {
		oldAreas.leaveArea(troops.TroopsID, troops.Country)
		saveAreasInfo(*oldAreas)
	}
}

//部队移动
//bigmap进程 loop 每秒执行一次,进行部队移动判断
//v2 考虑是否每个移动部队 单起协程处理
func loopTroopsMove(key interface{}, troops commonstruct.TroopsStruct, unix int64) {
	if troops.FitghtState == 1 {
		log.Warn("部队已上阵 ", troops)
		return
	}

	//第一步
	if troops.MoveStamp == 0 {
		troops.MoveStamp = unix + 2
		troops.MoveNum = 0
		saveMapTroops(&troops)
		//离开原来的格子
		troopsLeaveArea(&troops)
		return
	}

	//查看上次移动时间，判断本次是否可移动
	//var stepsecond int64 = 3
	//log.Infof("[%v]  [%v]  [%v]", troops.MoveStamp, unix, troops.MoveStamp+stepsecond > unix)
	if unix < (troops.MoveStamp + 3) {
		return
	}

	if troops.AreasList == nil || len(troops.AreasList) == 0 {
		log.Warn("bigmap move AreasList 空列表 roleid:", troops.Roleid)
		troops.State = common.TroopsState_Stationed
		saveMapTroops(&troops)
		return
	}

	//leaveindex := troops.AreasIndex
	//移动判定
	if int(troops.MoveNum) >= len(troops.AreasList) {
		log.Infof("移动超过路径(已达终点) [%v]  [%v]", troops.MoveNum, troops.AreasList)
	} else {
		//前进一格
		nextAreas := troops.AreasList[troops.MoveNum]
		//验证是否合法
		if cfg.AreasIsBeside(troops.AreasIndex, nextAreas) {

			//离开原来的格子
			troopsLeaveArea(&troops)

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
		process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgOverMove, Data: troops})
	} else {
		//正常移动中 发送至客户端
		process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgTroopsMove, Data: troops})
	}

	troops.MoveStamp = unix

	//1.区域触发战斗通知 2.进入区域部队队列
	if areasinfo := getAreasInfo(troops.AreasIndex); areasinfo != nil {
		if areasinfo.troopsTriggerBattle(&troops) || troops.MoveNum == 0 {
			areasinfo.entryAreasInfo(&troops)
		}
		saveAreasInfo(*areasinfo)
	} else {
		log.Warn("bigmap nil areas  index:", troops.AreasIndex)
	}
	saveMapTroops(&troops)
}

func cleanBigmapTroops(troops *commonstruct.TroopsStruct) {
	//清除出大地图
	troops.FitghtState = 0
	troops.FightType = 0
	troops.SkillUseNumber = 0
	troops.Scene = commonstruct.SceneNULL
	troops.Number = troops.MaxNumber * 4
	troops.RowHP = []int32{troops.MaxNumber, troops.MaxNumber, troops.MaxNumber, troops.MaxNumber}
	troops.State = common.TroopsState_StandBy
	troops.AreasIndex = cfg.GetCountryAreasIndex(troops.Country) //回归主城
	troops.CleanBuf(0)
	troops.CalculationAttribute()
	troopsSMap.Delete(troops.TroopsID)
	troopsLeaveArea(troops)
}
