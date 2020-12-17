package logger

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

//Init add logrus hook
func Init(loglevel log.Level, writefile bool, LogName string, path string) {
	log.SetLevel(loglevel)
	log.AddHook(NewContextHook(log.ErrorLevel, log.WarnLevel, log.DebugLevel, log.TraceLevel))

	if writefile {
		log.AddHook(fileHook(fmt.Sprintf("%v/%v_%v.log", path, LogName, "%Y%m%d%H%M")))
	}

}
