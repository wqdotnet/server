package logger

import (
	"bytes"
	"fmt"
	"server/tools"

	log "github.com/sirupsen/logrus"
)

//Init add logrus hook
func Init(loglevel log.Level, writefile bool, LogName string, path string) {
	log.SetLevel(loglevel)

	// if loglevel == log.DebugLevel || loglevel == log.TraceLevel {
	// 	log.SetFormatter(new(MyFormatter))
	// }
	log.SetFormatter(&log.TextFormatter{
		//DisableTimestamp: false,
		FullTimestamp: true,
		// 定义时间戳格式
		TimestampFormat: tools.DateTimeFormat,
		DisableSorting:  true,
	})
	log.AddHook(NewContextHook(log.ErrorLevel, log.WarnLevel, log.DebugLevel, log.TraceLevel, log.FatalLevel))
	if writefile {
		log.Infof("log path: [%v]", path)
		log.AddHook(fileHook(fmt.Sprintf("%v/%v_%v.log", path, LogName, "%Y%m%d%H%M")))
	}

}

// MyFormatter 自定义 formatter
type MyFormatter struct {
	Prefix string
	Suffix string
}

// Format implement the Formatter interface
func (mf *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// entry.Message 就是需要打印的日志
	b.WriteString(fmt.Sprintf("[%s][%v][%s]:  %s\n",
		entry.Level,
		tools.GoID(),
		entry.Time.Format(tools.DateTimeFormat),
		entry.Message))
	return b.Bytes(), nil
}
