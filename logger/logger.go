package logger

import (
	"bytes"
	"fmt"
	"server/tools"

	"github.com/sirupsen/logrus"
)

//Init add logrus hook
func Init(loglevel logrus.Level, writefile bool, LogName string, path string) {
	logrus.SetLevel(loglevel)

	// if loglevel == logrus.DebugLevel || loglevel == logrus.TraceLevel {
	// 	logrus.SetFormatter(new(MyFormatter))
	// }
	logrus.SetFormatter(&logrus.TextFormatter{
		//DisableTimestamp: false,
		FullTimestamp: true,
		// 定义时间戳格式
		TimestampFormat: tools.DateTimeFormat,
		DisableSorting:  true,
	})
	logrus.AddHook(NewContextHook(logrus.ErrorLevel, logrus.WarnLevel, logrus.DebugLevel, logrus.TraceLevel, logrus.FatalLevel))
	if writefile {
		logrus.Infof("log path: [%v]", path)
		logrus.AddHook(fileHook(fmt.Sprintf("%v/%v_%v.log", path, LogName, "%Y%m%d%H%M")))
	}

}

// MyFormatter 自定义 formatter
type MyFormatter struct {
	Prefix string
	Suffix string
}

// Format implement the Formatter interface
func (mf *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
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
