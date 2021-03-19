package process

import (
	"slgserver/gserver/commonstruct"
	"testing"

	log "github.com/sirupsen/logrus"
)

func init() {
	//log.Info(cfg.GlobalCfg.MapInfo)

	go func() {
		ch := make(chan commonstruct.ProcessMsg)
		Register("asd", ch)
		for info := range ch {
			log.Info(info.Data.(string))
		}
	}()
	go func() {
		ch := make(chan commonstruct.ProcessMsg)
		Register("bbc", ch)
		for info := range ch {
			log.Info(info)
		}
	}()

}

//go test -bench=Chan -run=XXX -benchtime=10s
func BenchmarkChan(b *testing.B) {
	SendMsg("asd", commonstruct.ProcessMsg{MsgType: "sb", Data: "sb250"})
	SendMsg("bbc", commonstruct.ProcessMsg{MsgType: "sb", Data: 13141515})
}
