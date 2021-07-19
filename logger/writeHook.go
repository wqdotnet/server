package logger

import (
	"server/tool"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func fileHook(path string) *lfshook.LfsHook {
	log.Debug("log path:", path)
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
			log.PanicLevel: writer,
			log.FatalLevel: writer,
			log.ErrorLevel: writer,
			log.WarnLevel:  writer,
			log.InfoLevel:  writer,
			log.DebugLevel: writer,
			log.TraceLevel: writer,
		},
		&log.TextFormatter{
			ForceColors:     true,
			TimestampFormat: tool.DateTimeFormat,
		},
	)
}
