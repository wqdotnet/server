package bigmapmanage

import (
	"slgserver/gserver/commonstruct"
	"slgserver/gserver/process"
	"slgserver/msgproto/common"
)

func skillExit(Queue []int32, exitAreasIndex int32) []int32 {
	var Queuelist, exittroopslist []int32
	num := len(Queue)
	if num >= 52 {
		exittroopslist = Queue[num-50:]
		Queuelist = Queue[:num-50]
	} else if num != 1 {
		exittroopslist = Queue[2:]
		Queuelist = Queue[:2]
	} else {
		return Queue
	}

	//退出区域
	for _, troopsid := range exittroopslist {
		if troops, ok := GetMapTroopsInfo(troopsid); ok {
			troops.State = common.TroopsState_Stationed
			troops.FitghtState = 0
			troops.AreasIndex = exitAreasIndex

			//1.区域触发战斗通知 2.进入区域部队队列
			areasinfo := getAreasInfo(exitAreasIndex)

			areasinfo.troopsTriggerBattle(&troops)
			areasinfo.entryAreasInfo(&troops)
			saveAreasInfo(*areasinfo)

			//npc 不发送
			if troops.Attribute == 0 {
				process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgUpdateTroopsInfo, Data: troops})
			}

			saveMapTroops(&troops)
		}
	}

	return Queuelist
}
