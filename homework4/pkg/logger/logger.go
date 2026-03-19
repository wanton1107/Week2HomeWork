package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(level string) *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})

	lv, err := logrus.ParseLevel(level)
	if err != nil {
		lv = logrus.InfoLevel
	}
	log.SetLevel(lv)
	return log
}
