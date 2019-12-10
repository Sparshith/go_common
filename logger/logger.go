package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Initialize() *logrus.Logger{
	log := logrus.New()
	if os.Getenv("GO_ENV") == "production" {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	return log
}