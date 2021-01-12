package logger

import (
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

	return lfshook.NewHook(
		lfshook.WriterMap{
			log.InfoLevel:  writer,
			log.ErrorLevel: writer,
			log.WarnLevel:  writer,
		},
		&log.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
}
