package logger

import (
	"server/tools"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func fileHook(path string) *lfshook.LfsHook {
	writer, _ := rotatelogs.New(
		path,
		rotatelogs.WithLinkName(path),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Duration(24)*time.Hour),
		//rotatelogs.WithRotationCount(maxRemainCnt),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithRotationSize(200*1024*1024),
	)
	//panic  fatal  error  warn  info  debug  trace
	return lfshook.NewHook(
		lfshook.WriterMap{
			logrus.PanicLevel: writer,
			logrus.FatalLevel: writer,
			logrus.ErrorLevel: writer,
			logrus.WarnLevel:  writer,
			logrus.InfoLevel:  writer,
			logrus.DebugLevel: writer,
			logrus.TraceLevel: writer,
		},
		&logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: tools.DateTimeFormat,
		},
	)
}
