package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func New(logPath string) *Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Warn("Failed to log to file, using default stderr")
	}

	return &Logger{log}
}
