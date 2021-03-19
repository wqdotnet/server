package mapserver

import (
	"context"
	"fmt"
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/gserver/timedtasks"
	"slgserver/msgproto/common"
	"time"

	log "github.com/sirupsen/logrus"
)

//消息号
type mapType int32

const (
	//BigMap 大地图
	BigMap mapType = 0
	//CopyMap 副本地图
	CopyMap mapType = 1
)

type mapCommand int32

const (
	commandInit mapCommand = 0
	commandLoop mapCommand = 1
)

type mapBaseInterface interface {
	init()
	handleInfo(msg interface{})
	loop(troopsSMap map[int32]commonstruct.TroopsStruct, areasSMap map[int32]AreasInfo, unix int64)
	command(command mapCommand)

	//部队走到终点
	overMove(troops commonstruct.TroopsStruct)
	//部队暂停移动
	troopsMove(troops commonstruct.TroopsStruct)
	//地图关闭
	terminate()

	//获取区域信息
	//getAllAreasInfoList()
	//getAreasInfo(areasid int32)
	//onAreasInfoChanage(areasid int32)
}

//AreasInfo 区域信息
type AreasInfo struct {
	AreasIndex int32   //区域ID
	Type       int32   //0 中立1-3:国家
	State      int32   //0 正常 1 战斗
	Occupy     int32   //占领信息  0 :无人占领  1-3国家KEY
	troopsA    []int32 //部队A
	troopsB    []int32 //部队B
	troopsC    []int32 //部队C

}

func startbase(mapname string, maptype mapType, base mapBaseInterface, mapcfg cfg.MapCfgStruct) mapMsgInterface {
	//log.Info("start map ", mapname)

	ctx, cancelFunc := context.WithCancel(context.Background())

	mapfunc := &mapMsgFunc{
		name:           mapname,
		maptype:        maptype,
		CancelFunc:     cancelFunc,
		commandChan:    make(chan mapCommand),
		mapmsgchan:     make(chan interface{}),
		troopsMoveChan: make(chan commonstruct.TroopsStruct),
		stopMoveChan:   make(chan int32),
	}

	timedtasks.AddTasks(fmt.Sprintf("%v_loop", mapname), "* * * * * ?", func() {
		mapfunc.commandChan <- commandLoop
	})
	go func() {
		// defer func() {
		// 	if err := recover(); err != nil {
		// 		log.Errorf("recover: %v", err)
		// 	}
		// }()

		var areasSMap = make(map[int32]AreasInfo)
		var troopsSMap = make(map[int32]commonstruct.TroopsStruct)

		//配置信息
		for _, ares := range mapcfg.Areas {
			index := int32(ares.Setindex)
			if _, ok := areasSMap[index]; ok {
				log.Warnf("Areasindex key:[%v] is existx", index)
			} else {
				areasSMap[index] = AreasInfo{AreasIndex: index, Type: int32(ares.Type), State: 0, Occupy: int32(ares.Type)}
			}
		}

		for {
			select {
			case command := <-mapfunc.commandChan:
				switch command {
				case commandInit:
					base.init()
				case commandLoop:
					unix := time.Now().Unix()
					loop(troopsSMap, areasSMap, unix, base)
					base.loop(troopsSMap, areasSMap, unix)
				default:
					base.command(command)
				}
			case msg := <-mapfunc.mapmsgchan:
				base.handleInfo(msg)

			case troopsmove := <-mapfunc.troopsMoveChan:
				handleTroopsEnterBigMap(troopsSMap, areasSMap, troopsmove)
			case troopsid := <-mapfunc.stopMoveChan:
				handleTroopsStopMove(troopsSMap, troopsid)
			case <-ctx.Done():
				base.terminate()
				return
			}
		}
	}()
	mapfunc.status = 1
	mapfunc.commandChan <- commandInit
	return mapfunc
}

//=====================================消息接口====================================================
type mapMsgInterface interface {
	GetName() string
	SendCommand(command mapCommand)
	SendMsg(msg interface{})

	SendTroopsMove(troops commonstruct.TroopsStruct)
	SendStopMove(troopsid int32)
	//部队退出地图
	//SendTroopsExit()
	Clone()
}

type mapMsgFunc struct {
	name    string
	maptype mapType
	status  int32

	commandChan chan mapCommand
	mapmsgchan  chan interface{}
	CancelFunc  context.CancelFunc
	//chan
	troopsMoveChan chan commonstruct.TroopsStruct
	stopMoveChan   chan int32
}

func (m *mapMsgFunc) GetName() string {
	return m.name
}

func (m *mapMsgFunc) SendCommand(command mapCommand) {
	m.commandChan <- command
}
func (m *mapMsgFunc) SendMsg(msg interface{}) {
	m.mapmsgchan <- msg
}

//SendTroopsMove 部队进入大地图移动
func (m *mapMsgFunc) SendTroopsMove(troops commonstruct.TroopsStruct) {
	if troops.AreasList == nil || len(troops.AreasList) == 0 {
		log.Warn("部队移动路径为空")
		return
	}
	m.troopsMoveChan <- troops
}

//SendStopMove 部队停止移动
func (m *mapMsgFunc) SendStopMove(troopsid int32) {
	m.stopMoveChan <- troopsid
}

//Clone 关闭
func (m *mapMsgFunc) Clone() {
	m.status = 0
	m.CancelFunc()
}

//=========================================================================================

// 部队进入大地图
func handleTroopsEnterBigMap(troopsSMap map[int32]commonstruct.TroopsStruct, areasSMap map[int32]AreasInfo, troops commonstruct.TroopsStruct) {
	if oldtroops, ok := troopsSMap[troops.TroopsID]; ok {
		//log.Debugf(" 大地图中已有部队 (改变部队移动方向) TroopsID:[%v]", troops)
		//oldtroops := data.(commonstruct.TroopsStruct)

		//战斗状态下不可以接受移动命令
		if oldtroops.State != common.TroopsState_fight {
			oldtroops.ArrivalTime = troops.ArrivalTime
			oldtroops.AreasList = troops.AreasList
			oldtroops.MoveNum = 0
			oldtroops.State = common.TroopsState_Move
			troopsSMap[oldtroops.TroopsID] = oldtroops
			//troopsSMap.Store(oldtroops.TroopsID, oldtroops)
		}
		return
	}
	//log.Info(troops)
	troopsSMap[troops.TroopsID] = troops
	//troopsSMap.Store(troops.TroopsID, troops)
}

//接收到部队暂停命令
func handleTroopsStopMove(troopsSMap map[int32]commonstruct.TroopsStruct, troopid int32) {

	if troops, ok := troopsSMap[troopid]; ok {
		troops.State = common.TroopsState_Pause
		troopsSMap[troopid] = troops
	}

}

func loop(troopsSMap map[int32]commonstruct.TroopsStruct, areasSMap map[int32]AreasInfo, unix int64, base mapBaseInterface) {

	for key, troops := range troopsSMap {
		switch troops.State {
		case common.TroopsState_Move: //移动状态
			loopTroopsMove(troopsSMap, key, troops, unix, base)
		case common.TroopsState_Pause: //暂停命令
			loopTroopsMove(troopsSMap, key, troops, unix, base)
		case common.TroopsState_Stationed: //原地驻扎
		case common.TroopsState_fight: //战斗状态
			//log.Debugf("[%v] [%v] [%v]", troops.MoveStamp, unix, troops.TroopsID)
			//模拟停留十秒后 让其回城或者留在该区域
			if unix > troops.MoveStamp+30 {

				//战胜->1.更新部队状态为驻扎 2.改变区域状态
				//战败->1.更新部队状态为0   2.部队从大地图移除 3.部队角色数据更新至role 用户不在线则保存到db
				//fightOver(key, troops, true)
			}
		}
	}

	// for areasindex, areasinfo := range areasSMap {

	// 	if areasinfo.State > 0 {

	// 	}
	// }

}

//部队移动
//bigmap进程 loop 每秒执行一次,进行部队移动判断
//v2 考虑是否每个移动部队 单起协程处理
func loopTroopsMove(troopsSMap map[int32]commonstruct.TroopsStruct, key int32, troops commonstruct.TroopsStruct, unix int64, base mapBaseInterface) {
	//第一步
	if troops.MoveStamp == 0 {
		troops.MoveStamp = unix
		troops.MoveNum = 0
		troopsSMap[key] = troops
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
		troopsSMap[key] = troops
		return
	}

	//移动判定
	if int(troops.MoveNum) >= len(troops.AreasList) {
		log.Infof("移动超过路径(已达终点) [%v]  [%v]", troops.MoveNum, troops.AreasList)
	} else {
		//前进一格
		nextAreas := troops.AreasList[troops.MoveNum]
		//验证是否合法
		//if cfg.AreasIsBeside(troops.AreasIndex, nextAreas) {
		troops.AreasIndex = nextAreas
		troops.MoveNum++
		//} else {
		//	troops.State = common.TroopsState_Pause
		//	log.Warnf("路径不合法  [%v]  ", troops)
		//}
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
		base.overMove(troops)
		//process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "OverMove", Data: troops})
	} else {
		//正常移动中 发送至客户端
		base.troopsMove(troops)
		//process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "TroopsMove", Data: troops})
	}

	troops.MoveStamp = unix

	// //此点是否需要战斗判定
	// if areasinfo := getAreasInfo(troops.AreasIndex); areasinfo != nil {
	// 	//中立区域 并且 不属于自己国家地盘的 触发战斗
	// 	if areasinfo.Type == 0 && areasinfo.Occupy != troops.Country {
	// 		troops.State = common.TroopsState_fight //进入战斗状态
	// 		troops.MoveNum = 0
	// 		process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "OnFitht", Data: troops})
	// 		fightAreas(*areasinfo, troops)
	// 	}
	// } else {
	// 	log.Warn("bigmap nil areas  index:", troops.AreasIndex)
	// }
	troopsSMap[key] = troops
}
