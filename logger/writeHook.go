package logger

import (
	"slgserver/tool"
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
		rotatelogs.WithMaxAge(time.Duration(24*7)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Minute),
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
