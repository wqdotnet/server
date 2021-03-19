package mapserver

import (
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/gserver/timedtasks"
	"slgserver/tool"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestMap(t *testing.T) {
	timedtasks.StartCronTasks()
	cfg.InitViperConfig("../../config", "json")

	log.Infof("[%v] start go:", tool.GoID())

	for i := 0; i < 2; i++ {
		go func(k int) {
			//log.Infof(" start go:[%v][%v]", tool.GoID(), i)

			switch k {
			case 1000:
				log.Info("start 1000")
			case 5000:
				log.Info("start 5000")
			case 8000:
				log.Info("start 8000")
			case 9000:
				log.Info("start 9000")
			case 9999:
				log.Info("start 9999")
			}

			mapInterface := startbase("asdf", CopyMap, &mapbase{}, cfg.GameCfg.MapInfo)

			for j := 0; j < 50000; j++ {
				mapInterface.SendTroopsMove(commonstruct.TroopsStruct{TroopsID: int32(100000 + j), State: 1, Roleid: 111, AreasList: []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}})
			}

			// time.Sleep(time.Second * 3)
			// mapInterface.SendTroopsMove(commonstruct.TroopsStruct{TroopsID: int32(100000 + k), State: 1, Roleid: 111, AreasList: []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}})
			// mapInterface.SendTroopsMove(commonstruct.TroopsStruct{TroopsID: int32(200000 + k), State: 1, Roleid: 111, AreasList: []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}})
			// mapInterface.SendTroopsMove(commonstruct.TroopsStruct{TroopsID: int32(300000 + k), State: 1, Roleid: 111, AreasList: []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}})
			// mapInterface.SendTroopsMove(commonstruct.TroopsStruct{TroopsID: int32(400000 + k), State: 1, Roleid: 111, AreasList: []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}})
			// mapInterface.SendTroopsMove(commonstruct.TroopsStruct{TroopsID: int32(500000 + k), State: 1, Roleid: 111, AreasList: []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}})
			// mapInterface.SendMsg("show me the money")
			//time.Sleep(time.Second * 30)
			//mapInterface.Clone()
			//time.Sleep(time.Second * 2)
		}(i)
	}

	log.Infof("start over ")
	time.Sleep(time.Second * 40)

}

//============================================================================================
type mapbase struct {
	// Name string
	// Type mapType

}

func (m *mapbase) init() {
	//log.Infof("[%v]Init:", tool.GoID())

}

func (m *mapbase) handleInfo(msg interface{}) {
	//log.Infof("[%v]handle:[%v]", tool.GoID(), msg)
}

func (m *mapbase) loop(troopsSMap map[int32]commonstruct.TroopsStruct, areasSMap map[int32]AreasInfo, unix int64) {
	//log.Infof("[%v]loop  time:[%v]", tool.GoID(), unix)
}

func (m *mapbase) command(command mapCommand) {
	log.Infof("[%v]command  [%v]", tool.GoID(), command)

}

//部队移动
func (m *mapbase) troopsMove(troops commonstruct.TroopsStruct) {
	//process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "TroopsMove", Data: troops})
	switch troops.TroopsID {

	case 149999:
		log.Infof("[%v]TroopsMove [%v] ", tool.GoID(), troops)
	}
	//log.Infof("[%v]TroopsMove [%v] ", tool.GoID(), troops)
}

//部队暂停移动
func (m *mapbase) overMove(troops commonstruct.TroopsStruct) {
	//process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "OverMove", Data: troops})

	switch troops.TroopsID {
	case 100001:
		log.Infof("[%v] %v", tool.GoID(), "OverMove 100001")
	case 149999:
		log.Infof("[%v] %v", tool.GoID(), "OverMove 149999")
	case 140000:
		log.Infof("[%v] %v", tool.GoID(), "OverMove 140000")
	case 130000:
		log.Infof("[%v] %v", tool.GoID(), "OverMove 130000")
	case 120000:
		log.Infof("[%v] %v", tool.GoID(), "OverMove 120000")
	case 110000:
		log.Infof("[%v] %v", tool.GoID(), "OverMove 110000")
	}

	//log.Infof("[%v]OverMove [%v] ", tool.GoID(), troops)
}

func (m *mapbase) terminate() {
	log.Infof("[%v]terminate ", tool.GoID())
}
