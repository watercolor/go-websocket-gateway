package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func Info(msg string) {
	logrus.Infof(msg)
}

func Error(msg string) {
	logrus.Errorf(msg)
}
