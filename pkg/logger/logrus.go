package logger

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	log      *logrus.Logger
	name     string
	hostName string
}

// Option to config log
type Option func(*logrusLogger)

// WithServiceName option for service name attribute
func WithServiceName(serviceName string) func(*logrusLogger) {
	return func(lg *logrusLogger) {
		lg.name = serviceName
	}
}

// WithHostName options for host name attribute
func WithHostName(hostName string) func(*logrusLogger) {
	return func(lg *logrusLogger) {
		lg.hostName = hostName
	}
}

// Operation attribute log operation name
type Operation string

// Op generate operation with name
func Op(val string) Operation {
	return Operation(val)
}

// Event attribute help log can make different when print log
type Event string

// E generate event with name
func E(op string) Event {
	return Event(op)
}

// NewJSONLogger initializes the json standard logger by logrus lib
func NewJSONLogger(opts ...Option) Log {
	var baseLogger = logrus.New()

	baseLogger.SetFormatter(&logrus.JSONFormatter{})
	baseLogger.SetReportCaller(true)

	var logrusLogger = &logrusLogger{
		log: baseLogger,
	}
	for idx := range opts {
		opt := opts[idx]
		opt(logrusLogger)
	}
	return logrusLogger
}

func (l *logrusLogger) GetLogger() *logrus.Logger {
	return l.log
}

func (l *logrusLogger) prepareFields() *logrus.Entry {
	return l.log.WithFields(
		logrus.Fields{
			"name":     l.name,
			"hostname": l.hostName,
		},
	)
}
func (l *logrusLogger) Log(vals ...interface{}) {
	fields, vals := detachFields(vals)
	l.prepareFields().WithFields(fields).Info(vals...)
}

func (l *logrusLogger) Debug(vals ...interface{}) {
	fields, vals := detachFields(vals)
	l.prepareFields().WithFields(fields).Debug(vals...)
}

func (l *logrusLogger) Info(vals ...interface{}) {
	fields, vals := detachFields(vals)
	l.prepareFields().WithFields(fields).Info(vals...)
}

func (l *logrusLogger) Infof(str string, vals ...interface{}) {
	l.prepareFields().Infof(str, vals...)
}

func (l *logrusLogger) Warn(vals ...interface{}) {
	fields, vals := detachFields(vals)
	l.prepareFields().WithFields(fields).Warning(vals...)
}

func (l *logrusLogger) Error(vals ...interface{}) {
	l.prepareFields().
		WithField("stack", string(debug.Stack())).
		Error(vals...)
}

func (l *logrusLogger) Errorf(str string, vals ...interface{}) {
	l.prepareFields().
		WithField("stack", string(debug.Stack())).
		Errorf(str, vals...)
}

func (l *logrusLogger) Fatal(vals ...interface{}) {
	fields, vals := detachFields(vals)
	l.prepareFields().WithFields(fields).Fatal(vals...)
}

func (l *logrusLogger) Fatalf(str string, vals ...interface{}) {
	l.prepareFields().Fatalf(str, vals...)
}
func detachFields(vals []interface{}) (logrus.Fields, []interface{}) {
	fields := logrus.Fields{}
	others := []interface{}{}
	for idx := range vals {
		arg := vals[idx]
		switch arg := arg.(type) {
		case Event:
			fields["event"] = arg
		case Operation:
			fields["op"] = arg
		default:
			others = append(others, arg)
		}
	}
	return fields, others
}
