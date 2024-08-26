package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// 设置日志输出到控制台
	logrus.SetOutput(os.Stdout)
	// 设置日志格式
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		// FullTimestamp:true,
		// DisableLevelTruncation:true,
	})

	// 设置日志级别，默认为 Info
	logrus.SetLevel(logrus.InfoLevel)
}

// Info 支持变参数的日志记录
func Info(args ...interface{}) {
	logrus.Infof(fmt.Sprintf("%v", args...))
}

// Error 支持变参数的日志记录
func Error(args ...interface{}) {
	logrus.Errorf(fmt.Sprintf("%v", args...))
}

// Infof 支持格式化的日志记录
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Errorf 支持格式化的日志记录
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
