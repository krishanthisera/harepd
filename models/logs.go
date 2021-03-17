package models

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

//NewLogger to store logs
func NewLogger(conf *Config) *logrus.Logger {
	if Log != nil {
		return Log
	}

	path := conf.Harepd.Logs.FilePath

	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithMaxAge(time.Duration(conf.Harepd.Logs.MaxAge)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(conf.Harepd.Logs.RotationTime)*time.Second),
	)
	if err != nil {
		panic(err)
	}

	logrus.New().Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.DebugLevel: writer,
			logrus.WarnLevel:  writer,
		},
		&logrus.JSONFormatter{},
	))

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		path,
		&logrus.JSONFormatter{},
	))

	return Log
}
