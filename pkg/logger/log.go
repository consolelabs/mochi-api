package logger

import "github.com/sirupsen/logrus"

// Log interface for log
type Log interface {
	Log(vals ...interface{})
	Debug(vals ...interface{})
	Info(vals ...interface{})
	Infof(str string, vals ...interface{})
	Warn(vals ...interface{})
	Error(vals ...interface{})
	Errorf(str string, vals ...interface{})
	Fatal(vals ...interface{})
	Fatalf(str string, vals ...interface{})
	GetLogger() *logrus.Logger
}
