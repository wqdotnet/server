package fightmod

import (
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
)

//技能效果  2 buff
func skillEffectBuff(bufCfg *cfg.BuffCfg, attackTroops, defenseTroops *commonstruct.TroopsStruct) {

	switch bufCfg.EffectType {
	case 2:
		attackTroops.AddBuff(bufCfg.ID)
		attackTroops.CalculationAttribute()
	}

}
