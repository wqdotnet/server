package clienconnect

import (
	"slgserver/db"
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"testing"

	"github.com/go-playground/assert/v2"
	log "github.com/sirupsen/logrus"
)

func init() {
	db.StartMongodb("slggame", "mongodb://localhost:27017")
}

func TestMap(t *testing.T) {
	cfg.InitViperConfig("../../config", "json")

	var areasindex int32 = 0
	for _, arecfg := range cfg.GameCfg.MapInfo.Areas {
		if arecfg.Type == 1 {
			areasindex = int32(arecfg.Setindex)
		}
	}

	log.Info("出发点:", areasindex)
}

func TestGetdata(t *testing.T) {
	var troopslist = make(map[int32]*commonstruct.TroopsStruct)
	for _, v := range getRoleAlltroops(2) {
		troopslist[v.TroopsID] = v
	}

	assert.Equal(t, troopslist[1], nil)
	log.Info(troopslist[2])
	log.Info(troopslist[3])

}
