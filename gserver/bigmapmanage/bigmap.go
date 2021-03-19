package bigmapmanage

import (
	"encoding/json"
	"slgserver/db"
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/gserver/process"
	"slgserver/msgproto/common"
	"slgserver/msgproto/fight"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	areasSMap  sync.Map
	troopsSMap sync.Map
	//战斗设置
	fightSetMap  = make(map[int32]commonstruct.FightSetting)
	fightSetAuto = make(map[int32]bool)

	// ================== chan ============================
	bigmapmsgchan        = make(chan string, 1)
	troopsMoveChan       = make(chan commonstruct.TroopsStruct, 1)
	stopMoveChan         = make(chan int32, 1)
	msgSubscribeChan     = make(chan areasMsgChan, 1)                  //区域消息订阅/取消
	fightSettingChan     = make(chan commonstruct.FightSetting, 1)     //战斗设置
	troopsExitBigmapChan = make(chan commonstruct.TroopsExitBigmap, 1) //部队退出大地图

	updateTroopsChan = make(chan commonstruct.TroopsStruct, 1) //更新部队单个数据（未完成）
)

//StartBigmapGoroutine init
func StartBigmapGoroutine() {
	log.Info("start bigmap goroutine")
	//加载地图配置 缓存
	initBigmapAreasInfo()
	fightSetAuto[0] = true
	go func() {

		// defer func() {
		// 	if err := recover(); err != nil {
		// 		log.Errorf("recover: %v", err)
		// 	}
		// }()

		for {
			select {
			case command := <-bigmapmsgchan:
				handleMapCommand(command)
			case troopsmove := <-troopsMoveChan:
				handleTroopsEnterBigMap(troopsmove)
			case troopsid := <-stopMoveChan:
				handleTroopsStopMove(troopsid)
			case msg := <-msgSubscribeChan:
				areasMsgSubscribe(msg.subscribe, msg.areasIndex, msg.roleid)
			case setting := <-fightSettingChan:
				bigmapSetFightSet(&setting)
			case troopsexit := <-troopsExitBigmapChan:
				troopsExitBigmap(&troopsexit)
			case troopsupdate := <-updateTroopsChan:
				updateTroopsInfo(&troopsupdate)
				// case <-ctx.Done():
				// 	cloneBigmap()
			}
		}

	}()
}

func bigmapSetFightSet(fset *commonstruct.FightSetting) {
	fightSetAuto[fset.RoleID] = fset.AutoSelect

	if troops, ok := GetMapTroopsInfo(fset.TroopsID); ok {
		troops.FightSet = fset
		//推送给下轮战斗双方 技能施放选择情况
		process.SendSocketMsg(fset.RoleID,
			int32(fight.MSG_FIGHT_Module_FIGHT),
			int32(fight.MSG_FIGHT_S2C_Select_Tactics), &fight.S2C_SelectTactics{SelectSkill: 2})
	}
	log.Info("战斗设置：", fightSetAuto[fset.RoleID])

}

//=================================================================================================
//初始化大地图所有区域信息
func initBigmapAreasInfo() {
	//配置信息
	for _, ares := range cfg.GameCfg.MapInfo.Areas {
		index := int32(ares.Setindex)
		if _, ok := areasSMap.Load(index); ok {
			log.Warnf("Areasindex key:[%v] is existx", index)
		} else {
			areasSMap.Store(index, newAreasInfo(index, int32(ares.Type)))
		}
	}

	//缓存信息覆盖
	value, _ := db.HVALS("areasSMap")
	for _, v := range value {
		areas := &AreasInfo{}
		json.Unmarshal(v, areas)
		areasSMap.Store(areas.AreasIndex, *areas)
	}

	//从缓存初始化地图部队信息、区域信息
	value, _ = db.HVALS("troopsSMap")
	for _, v := range value {
		troops := &commonstruct.TroopsStruct{}
		json.Unmarshal(v, troops)
		saveMapTroops(troops)

		//处理部队驻扎
		if troops.State == common.TroopsState_Stationed || troops.State == common.TroopsState_fight {
			log.Infof("驻扎部队区域信息：[%v - %v] [%v]", troops.Roleid, troops.TroopsID, *getAreasInfo(troops.AreasIndex))
			// troops.AreasIndex
		}
	}

}

// func getfightSetMap(roleid int32) *commonstruct.FightSetting {
// 	fset := fightSetMap[roleid]
// 	return &fset
// }

// func setfightSetMap(fset *commonstruct.FightSetting) {
// 	fightSetMap[fset.RoleID] = *fset
// }

//===============================================外部接口==================================================

//CloneBigmap 关闭大地图
func CloneBigmap() {
	//保存信息 大地图上所有队伍信息
	log.Info("clone Bigmap save data")

	AreasRange(func(areas AreasInfo) bool {
		//只保存 中立地区
		if areas.Type > 0 {
			return true
		}

		//已占领的，正在发生战斗的
		if areas.Occupy > 0 || areas.State > 0 {
			b, err := json.Marshal(areas)
			if err != nil {
				return true
			}
			db.HMSET("areasSMap", areas.AreasIndex, b)
			log.Trace("save areasSMap: ", areas)
		}

		return true
	})

	troopsSMap.Range(func(key, value interface{}) bool {
		troops := value.(commonstruct.TroopsStruct)
		b, err := json.Marshal(troops)
		if err != nil {
			return true
		}
		db.HMSET("troopsSMap", troops.TroopsID, b)
		log.Trace("save troopsSMap:", troops)
		return true
	})
}

//----数据发送至大地图接口----

//SendMsgBigMap msg
func SendMsgBigMap(msg string) {
	bigmapmsgchan <- msg
}

//SendTroopsMove 部队进入大地图移动
func SendTroopsMove(troops *commonstruct.TroopsStruct) {
	if troops.AreasList == nil || len(troops.AreasList) == 0 {
		log.Warn("部队移动路径为空")
		return
	}
	troops.Scene = commonstruct.SceneBigMap
	troopsMoveChan <- *troops
}

//SendStopMove 部队停止移动
func SendStopMove(troopsid int32) {
	stopMoveChan <- troopsid
}

type areasMsgChan struct {
	roleid     int32
	areasIndex int32
	subscribe  bool
}

//SendAreasSubscribe 消息订阅
func SendAreasSubscribe(areas, roleid int32) {
	msgSubscribeChan <- areasMsgChan{roleid: roleid, areasIndex: areas, subscribe: true}
}

//SendAreasCancelSubscribe 取消订阅
func SendAreasCancelSubscribe(areas, roleid int32) {
	msgSubscribeChan <- areasMsgChan{roleid: roleid, areasIndex: areas, subscribe: false}
}

//SendFightSetting 战斗设置
//tacticsID 技能ID
//autoSelect 自动选择
func SendFightSetting(roleid, troopsid, skillid, tacticsID int32, autoSelect bool) {
	fightSettingChan <- commonstruct.FightSetting{AutoSelect: autoSelect,
		SkillID:    skillid,
		TacticsID:  tacticsID,
		RoleID:     roleid,
		TroopsID:   troopsid,
		SelectTime: time.Now().Unix(),
	}
}

//SendUpdateTroopsInfo 更新大地图中部队数据
func SendUpdateTroopsInfo(troops *commonstruct.TroopsStruct) {
	updateTroopsChan <- *troops
}

//SendTroopsExitBigmap 部队退出大地图
func SendTroopsExitBigmap(troopsexit *commonstruct.TroopsExitBigmap) {
	troopsExitBigmapChan <- *troopsexit
}

//==================================================接收数据处理==================================================

func handleMapCommand(command string) {
	//log.Infof("[%v] :[%v]", tool.GoID(), command)
	switch command {
	case "BigMapLoop_OneSecond":
		loop(time.Now().Unix())
	default:
		log.Warnf("nil command: [%v]", command)
	}
}

func loop(unix int64) {
	//部队移动
	troopsSMap.Range(func(key, value interface{}) bool {
		troops := value.(commonstruct.TroopsStruct)

		switch troops.State {
		case common.TroopsState_Move: //移动状态
			loopTroopsMove(key, troops, unix)
		case common.TroopsState_Pause: //暂停命令
			loopTroopsMove(key, troops, unix)
		case common.TroopsState_Stationed: //原地驻扎

		}
		return true
	})

	//区域战斗处理
	AreasRange(func(areas AreasInfo) bool {
		if areas.State == 0 {
			return true
		}

		//战斗处理
		areas.areasFighting(unix)
		saveAreasInfo(areas)

		return true
	})
}
