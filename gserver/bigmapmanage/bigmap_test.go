package bigmapmanage

import (
	"server/gserver/cfg"
	"server/gserver/timedtasks"
	"server/msgproto/bigmap"
	"testing"

	log "github.com/sirupsen/logrus"
)

func init() {
	//ctx, cancelFunc := context.WithCancel(context.Background())

	StartBigmapGoroutine()
	//启动定时器
	timedtasks.StartCronTasks()
	//大地图loop
	timedtasks.AddTasks("bigmaploop", "* * * * * ?", func() {
		SendMsgBigMap("BigMapLoop_OneSecond")
	})

	cfg.InitViperConfig("../../config", "json")
	initBigmapAreasInfo()
}

// func TestMove(t *testing.T) {

// 	fc := func(troopsid int32, roleid int32, path []int32) {
// 		ch := make(chan commonstruct.ProcessMsg)
// 		process.Register(roleid, ch)
// 		SendTroopsMove(commonstruct.TroopsStruct{TroopsID: troopsid, State: 1, Roleid: roleid, AreasList: path})
// 		for troops := range ch {
// 			switch troops.MsgType {
// 			case "TroopsMove":
// 				datat := troops.Data.(commonstruct.TroopsStruct)
// 				log.Infof("[%v]role receive:[%v]", tool.GoID(), datat)

// 				if datat.AreasIndex == 25 {
// 					SendStopMove(datat.TroopsID)
// 				}
// 			case "OverMove":
// 				log.Info("OverMove2:", time.Now().Unix())
// 				return

// 			}
// 		}
// 	}

// 	unix := time.Now().Unix()
// 	arlist := []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}
// 	go fc(102, 111, []int32{21, 22, 23, 24, 25, 26, 27, 28, 29})

// 	log.Info("OverMove1:", int64(len(arlist)*3)+unix)

// 	time.Sleep(time.Second * 50)
// }

func TestInit(t *testing.T) {
	SetAreasInfo(1048, AreasInfo{AreasIndex: 234})
	log.Info("1048:", GetAreasInfo(1048))
	log.Info("748:", GetAreasInfo(748))
}

func TestBigmapConfigInit(t *testing.T) {
	areaslist := &bigmap.S2C_AreasInfo{}

	Range(func(value AreasInfo) bool {
		log.Info("value.Type:", value.Type, "    value.State:", value.State)
		//只发送被占领的，正在发生战斗的
		if value.occupy > 0 || value.State > 0 {
			areaslist.AreasInfoList = append(areaslist.AreasInfoList,
				&bigmap.P_AreasInfo{AreasIndex: value.AreasIndex,
					Type:  value.Type,
					State: value.State})
		}
		return true
	})

	log.Info(len(areaslist.AreasInfoList))

}
