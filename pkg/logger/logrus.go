package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
	fields Fields
}

func NewLogrusLogger() Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	defaultFields := Fields{
		"service":  "mochi-api",
		"hostname": hostname,
	}

	return &LogrusLogger{
		logger: l,
		fields: defaultFields,
	}
}

func (l *LogrusLogger) Fields(data Fields) Logger {
	return &LogrusLogger{
		logger: l.logger,
		fields: data,
	}
}

func (l *LogrusLogger) Debug(msg string) {
	l.logger.WithFields(logrus.Fields(l.fields)).Debug(msg)
}

func (l *LogrusLogger) Debugf(msg string, args ...interface{}) {
	l.logger.WithFields(logrus.Fields(l.fields)).Debugf(msg, args...)
}

func (l *LogrusLogger) Info(msg string) {
	l.logger.WithFields(logrus.Fields(l.fields)).Info(msg)
}

func (l *LogrusLogger) Infof(msg string, args ...interface{}) {
	l.logger.WithFields(logrus.Fields(l.fields)).Infof(msg, args...)
}

func (l *LogrusLogger) Warn(msg string) {
	l.logger.WithFields(logrus.Fields(l.fields)).Warn(msg)
}

func (l *LogrusLogger) Warnf(msg string, args ...interface{}) {
	l.logger.WithFields(logrus.Fields(l.fields)).Warnf(msg, args...)
}

func (l *LogrusLogger) Error(err error, msg string) {
	l.logger.WithFields(logrus.Fields(l.fields)).WithError(err).Error(msg)
}

func (l *LogrusLogger) Errorf(err error, msg string, args ...interface{}) {
	l.logger.WithFields(logrus.Fields(l.fields)).WithError(err).Errorf(msg, args...)
}

func (l *LogrusLogger) Fatal(err error, msg string) {
	l.logger.WithFields(logrus.Fields(l.fields)).WithError(err).Fatal(msg)
}

func (l *LogrusLogger) Fatalf(err error, msg string, args ...interface{}) {
	l.logger.WithFields(logrus.Fields(l.fields)).WithError(err).Fatalf(msg, args...)
}
