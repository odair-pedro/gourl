package logging

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type logger struct {
	*logrus.Logger
}

var loggerInstance *logger
var once sync.Once

func GetLoggerInstance() *logger {
	if loggerInstance == nil {
		once.Do(func() {
			loggerInstance = newLogger()
		})
	}
	return loggerInstance
}

func newLogger() *logger {
	customLogger := &logger{logrus.New()}
	customLogger.SetFormatter(&logrus.TextFormatter{})
	return customLogger
}
