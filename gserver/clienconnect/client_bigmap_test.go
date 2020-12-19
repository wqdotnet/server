package clienconnect

import (
	"server/gserver/cfg"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestMap(t *testing.T) {

	log.Info(cfg.GlobalCfg.MapInfo)
}
