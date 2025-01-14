package src

import (
	"log"
	"os"
	"sync"

	"euro-exchange/config"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	StdLogger  *logrus.Logger
	FileLogger *logrus.Logger
	sync.Mutex
}

func NewLogger(cfg *config.Config) *Logger {
	fileLogger := logrus.New()
	stdLogger := logrus.New()

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02T15:04:05.999999999Z07:00"

	fileLogger.Formatter = formatter
	stdLogger.Formatter = formatter

	file, err := os.OpenFile(cfg.LogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0666)
	if err == nil {
		fileLogger.Out = file
	} else {
		log.Fatalf("unable to set logger, %s", err)
	}

	return &Logger{
		StdLogger:  stdLogger,
		FileLogger: fileLogger,
	}
}

func (l *Logger) Log(level logrus.Level, args interface{}) {
	l.Lock()
	l.StdLogger.Log(level, args)
	l.FileLogger.Log(level, args)
	l.Unlock()
}
