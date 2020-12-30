package clienconnect

import (
	"server/db"
	"server/gserver/cfg"
	"testing"

	log "github.com/sirupsen/logrus"
)

func init() {
	db.StartMongodb("slggame", "mongodb://localhost:27017")
}

func TestMap(t *testing.T) {
	cfg.InitViperConfig("../../config", "json")

	var areasindex int32 = 0
	for _, arecfg := range cfg.GlobalCfg.MapInfo.Areas {
		if arecfg.Type == 1 {
			areasindex = int32(arecfg.Setindex)
		}
	}

	log.Info("出发点:", areasindex)
}

func TestGetdata(t *testing.T) {
	for k, v := range getRoleAlltroops(2) {
		log.Info(k, ":", v)
	}
}
