package bigmapmanage

import (
	"encoding/json"
	"server/db"
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/msgproto/common"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	areasSMap  sync.Map
	troopsSMap sync.Map

	//chan
	bigmapmsgchan  = make(chan string)
	troopsMoveChan = make(chan commonstruct.TroopsStruct)
	stopMoveChan   = make(chan int32)
)

//StartBigmapGoroutine init
func StartBigmapGoroutine() {
	log.Info("start bigmap goroutine")
	//加载地图配置 缓存
	initBigmapAreasInfo()

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
				// case <-ctx.Done():
				// 	cloneBigmap()
			}
		}

	}()
}

//初始化大地图所有区域信息
func initBigmapAreasInfo() {
	//配置信息
	for _, ares := range cfg.GlobalCfg.MapInfo.Areas {
		index := int32(ares.Setindex)
		if _, ok := areasSMap.Load(index); ok {
			log.Warnf("Areasindex key:[%v] is existx", index)
		} else {
			areasSMap.Store(index, AreasInfo{AreasIndex: index, Type: int32(ares.Type), State: 0, Occupy: int32(ares.Type)})
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
		troopsSMap.Store(troops.TroopsID, *troops)
	}

	db.RedisExec("del", "areasSMap")
	db.RedisExec("del", "troopsSMap")
}

//----------------------------------------发送至大地图接口------------------------------------------------------

//SendMsgBigMap msg
func SendMsgBigMap(msg string) {
	bigmapmsgchan <- msg
}

//SendTroopsMove 部队进入大地图移动
func SendTroopsMove(troops commonstruct.TroopsStruct) {
	if troops.AreasList == nil || len(troops.AreasList) == 0 {
		log.Warn("部队移动路径为空")
		return
	}
	troopsMoveChan <- troops
}

//SendStopMove 部队停止移动
func SendStopMove(troopsid int32) {
	stopMoveChan <- troopsid
}

//------------------------------------------------------------------------------------------------

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
		case common.TroopsState_fight: //战斗状态
			//log.Debugf("[%v] [%v] [%v]", troops.MoveStamp, unix, troops.TroopsID)
			//模拟停留十秒后 让其回城或者留在该区域
			if unix > troops.MoveStamp+30 {

				//战胜->1.更新部队状态为驻扎 2.改变区域状态
				//战败->1.更新部队状态为0   2.部队从大地图移除 3.部队角色数据更新至role 用户不在线则保存到db
				fightOver(key, troops, true)
			}
		}
		return true
	})

}

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
