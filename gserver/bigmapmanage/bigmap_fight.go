package bigmapmanage

import (
	"server/gserver/commonstruct"
	"server/msgproto/common"

	log "github.com/sirupsen/logrus"
)

//战胜->1.更新部队状态为驻扎 2.改变区域状态
//战败->1.更新部队状态为0   2.部队从大地图移除 3.部队角色数据更新至role 用户不在线则保存到db
func fightOver(troops commonstruct.TroopsStruct, victory bool) {
	log.Info("战斗结束")

	if victory {
		troops.State = common.TroopsState_Stationed
	} else {
		troops.State = common.TroopsState_StandBy
	}

}
