package fightmod

import (
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestMove(t *testing.T) {
	cfg.InitViperConfig("../../config", "json")
	log.Info()
	troops1 := commonstruct.NewTroops("test111", 111, 1, 1, 4)
	troops2 := commonstruct.NewTroops("test122", 222, 1, 1, 3)

	list := FightCalculation(troops1, troops2, nil, &commonstruct.FightSetting{AutoSelect: true}, &commonstruct.FightSetting{AutoSelect: true})
	for _, v := range list {
		log.Info("=====>", v)
		if v.Type == 1 && v.Value == 202 {
			log.Info("50%几率吓退敌方后50个部队")
		}
	}

}
