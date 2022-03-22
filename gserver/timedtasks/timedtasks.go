package timedtasks

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var (
	crontask *cron.Cron
	taskmap  sync.Map
)

//StartCronTasks spec := "* * * * * ?"
func StartCronTasks() {
	logrus.Infof("Start Cron")
	crontask = cron.New(cron.WithSeconds())
	crontask.Start()
}

//AddTasks add
func AddTasks(key string, spec string, cmd func()) error {
	entryid, err := crontask.AddFunc(spec, cmd)

	if _, ok := taskmap.Load(key); ok {
		return fmt.Errorf("[%v] 已注册", key)
	}

	if err == nil {
		taskmap.Store(key, entryid)
	}
	return err
}

//RemoveTasks Remove Tasks
func RemoveTasks(key string) {
	entryid, ok := taskmap.Load(key)
	if ok {
		crontask.Remove(entryid.(cron.EntryID))
	}
}
