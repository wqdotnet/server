package bigmapmanage

import (
	"slgserver/db"
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/gserver/process"
	"slgserver/gserver/timedtasks"
	"slgserver/tool"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	cfg.InitViperConfig("../../config", "json")
	//ctx, cancelFunc := context.WithCancel(context.Background())
	db.StartRedis("127.0.0.1:6379", 0)
	StartBigmapGoroutine()
	//启动定时器
	timedtasks.StartCronTasks()
	//大地图loop
	timedtasks.AddTasks("bigmaploop", "* * * * * ?", func() {
		SendMsgBigMap("BigMapLoop_OneSecond")
	})
}

// func TestBigmLoad(t *testing.T) {
// 	//CloneBigmap()
// 	AreasRange(func(areas AreasInfo) bool {
// 		//只保存 中立地区
// 		if areas.Type > 0 {
// 			return true
// 		}
// 		areas.AttackQueue = append(areas.AttackQueue, 3)
// 		//已占领的，正在发生战斗的
// 		//if areas.Occupy > 0 || areas.State > 0 {
// 		b, _ := json.Marshal(areas)
// 		// if err != nil {
// 		// 	return true
// 		// }
// 		// db.HMSET("areasSMap", areas.AreasIndex, b)
// 		log.Infof("save areasSMap: %v", string(b))
// 		//}
// 		return true
// 	})
// }

// func TestEentryAreasInfo(t *testing.T) {
// 	log.Info("TestEentryAreasInfo")
// 	areas := newAreasInfo(1111, 1) //&AreasInfo{}
// 	areas.State = 1
// 	areas.entryAreasInfo(&commonstruct.TroopsStruct{TroopsID: 101, Country: 1})
// 	areas.entryAreasInfo(&commonstruct.TroopsStruct{TroopsID: 102, Country: 1})
// 	areas.entryAreasInfo(&commonstruct.TroopsStruct{TroopsID: 103, Country: 1})
// 	areas.entryAreasInfo(&commonstruct.TroopsStruct{TroopsID: 101, Country: 1})
// 	areas.entryAreasInfo(&commonstruct.TroopsStruct{TroopsID: 104, Country: 2})
// 	areas.entryAreasInfo(&commonstruct.TroopsStruct{TroopsID: 105, Country: 3})
// 	//areas.refreshTop2TroopsFightState()
// 	log.Info("areas:", areas)
// }

// func TestFight(t *testing.T) {
// 	troops1 := commonstruct.NewTroops("test1", 111, 1, 1, 1, 1)
// 	troops2 := commonstruct.NewTroops("test2", 222, 1, 1, 1, 2)
// 	log.Info("hp:", attackCalculation(2000, troops1, troops2))
// 	log.Info("hp:", attackCalculation(2000, troops2, troops1))
// }

func TestMove(t *testing.T) {

	log.Info()
	log.Info("====================部队移动===============")
	fc := func(troopsid int32, roleid int32, path []int32) {
		ch := make(chan commonstruct.ProcessMsg)
		process.Register(roleid, ch)

		fightSetAuto[roleid] = true
		troops := commonstruct.NewTroops("test1", troopsid, 1, 1, 5)
		log.Info("troops:", troops)
		troops.Roleid = roleid
		troops.State = 1
		troops.AreasList = path
		troops.AreasIndex = 9994
		//troops.Level = 40
		//troops.CalculationAttribute()
		SendTroopsMove(troops)

		// troops2 := commonstruct.NewTroops("test2", troopsid+1, 1, 1, 1, 1)
		// log.Info("troops:", troops)
		// troops2.Roleid = roleid
		// troops2.State = 1
		// troops2.AreasList = path
		// troops2.AreasIndex = 9994
		// SendTroopsMove(troops2)

		for {
			select {
			case troops := <-ch:
				switch troops.MsgType {
				case commonstruct.ProcessMsgSocket:
					log.Info("ProcessMsgSocket=====>:", troops)
				case commonstruct.ProcessMsgOverFitht:
					troops := troops.Data.(commonstruct.TroopsStruct)
					log.Info("战斗结束=====>：", troops.State, troops.FitghtState)
				case commonstruct.ProcessMsgUpdateTroopsInfo:
					troops := troops.Data.(commonstruct.TroopsStruct)
					log.Info("更新部队状态=====>：", troops.State, troops.FitghtState)
				case commonstruct.ProcessMsgOnFitht:
					log.Info("触发战斗=====>")
				case "TroopsMove":
					datat := troops.Data.(commonstruct.TroopsStruct)
					log.Infof("[%v][%v]role receive=====>:[%v]", time.Now().Format("15:04:05"), tool.GoID(), datat)
				case commonstruct.ProcessMsgAddExp:
					addexpitem := troops.Data.(commonstruct.AddExpItem)
					log.Info("部队加经验=====>：", addexpitem)

				case "OverMove":
					datat := troops.Data.(commonstruct.TroopsStruct)
					log.Infof("[%v] OverMove2=====>: [%v]  [%v]", time.Now().Format("15:04:05"), time.Now().Unix(), datat)
				}
			}
		}

	}

	// unix := time.Now().Unix()
	// arlist := []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}
	go fc(102, 111, []int32{9425})

	// log.Info("OverMove1:", int64(len(arlist)*3)+unix)

	time.Sleep(time.Second * 120)
	tr, _ := GetMapTroopsInfo(102)
	log.Info("查询大地图中部队数据：", tr.State, tr.FitghtState)

}
