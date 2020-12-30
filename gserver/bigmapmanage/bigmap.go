package bigmapmanage

import (
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
	//加载地图配置
	//redis -> mongodb -> config
	initBigmapAreasInfo()
	// for _, ares := range cfg.GlobalCfg.MapInfo.Areas {
	// 	if ares.Type > 0 {
	// 		SetAreasInfo(int32(ares.Setindex), AreasInfo{AreasIndex: int32(ares.Setindex), Type: int32(ares.Type)})
	// 	}
	// }

	//从缓存初始化地图部队信息、区域信息

	go func() {
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
		loop()
	default:
		log.Warnf("nil command: [%v]", command)
	}
}

func loop() {

	troopsSMap.Range(func(key, value interface{}) bool {
		troops := value.(commonstruct.TroopsStruct)

		switch troops.State {
		case common.TroopsState_Move: //移动状态
			loopTroopsMove(key, troops, time.Now().Unix())
		case common.TroopsState_Pause: //暂停命令
			loopTroopsMove(key, troops, time.Now().Unix())
		case common.TroopsState_Stationed: //原地驻扎
		case common.TroopsState_fight: //战斗状态
			//模拟停留十秒后 让其回城或者留在该区域
			if troops.MoveStamp+10 > time.Now().Unix() {

				fightOver(troops, true)
				//战胜->1.更新部队状态为驻扎 2.改变区域状态
				//战败->1.更新部队状态为0   2.部队从大地图移除 3.部队角色数据更新至role 用户不在线则保存到db
			}
		}
		return true
	})
}

//CloneBigmap 关闭大地图
func CloneBigmap() {
	//保存信息 大地图上所有队伍信息
	log.Info("clone Bigmap save data")
}
