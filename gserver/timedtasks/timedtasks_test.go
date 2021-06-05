package timedtasks

import (
	"server/tool"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("start task:", tool.GoID())
	StartCronTasks()
}
func TestInit(t *testing.T) {
	AddTasks("st", "* * * * * ?", func() {
		log.Info("st loop")
	})

	time.Sleep(5000000000)
	log.Info("RemoveTasks st")
	RemoveTasks("st")
	time.Sleep(5000000000)
}

// //go test -bench=Chan -run=XXX -benchtime=10s
// func BenchmarkChan(b *testing.B) {
// 	tasks.Range(func(key, value interface{}) bool {
// 		value.(chan interface{}) <- key
// 		return true
// 	})

// }
