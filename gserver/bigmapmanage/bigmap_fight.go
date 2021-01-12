package bigmapmanage

import (
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/gserver/process"
	"server/msgproto/common"
)

//战胜->1.更新部队状态为驻扎 2.改变区域状态
//战败->1.更新部队状态为0   2.部队从大地图移除 3.部队角色数据更新至role 用户不在线则保存到db
func fightOver(key interface{}, troops commonstruct.TroopsStruct, victory bool) {
	//log.Info("战斗结束")

	if victory {
		troops.State = common.TroopsState_Stationed
		troopsSMap.Store(key, troops)
		fightoverAreas(troops.AreasIndex, troops.Country, troops.Roleid)

	} else {
		//战败回归主城 从大地图移除
		troops.AreasIndex = cfg.GetCountryAreasIndex(troops.Country)
		troops.State = common.TroopsState_StandBy
		troopsSMap.Delete(key)
	}

	//
	process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "OverFitht", Data: troops})

}

//--------------------------------------------------------------------------
//部队进入区域触发战斗
func fightAreas(areas AreasInfo, troops commonstruct.TroopsStruct) {
	if areas.State == 0 {
		areas.State = 1
		process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: "AreasState", Data: areas})
	}
	//部队信息加入列表

	saveAreasInfo(areas)
}

//区域战斗结束
func fightoverAreas(index, occupy, roleid int32) {
	areas := getAreasInfo(index)
	areas.Occupy = occupy
	areas.State = 0
	saveAreasInfo(*areas)
	process.SendMsg(roleid, commonstruct.ProcessMsg{MsgType: "AreasState", Data: *areas})
}
