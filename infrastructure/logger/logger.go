package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Log interface {
	Println(...interface{})
	Printf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

var (
	Logger Log = NewLoggerLogrus()
)

type LoggerLogrus struct {
	logger *logrus.Logger
}

func setupLogger(l *LoggerLogrus) {
	env := os.Getenv("ENV")

	if env == "debug" {
		l.logger.SetLevel(logrus.DebugLevel)
	} else if env == "prod" {
		l.logger.SetLevel(logrus.InfoLevel)
	}
}

func NewLoggerLogrus() *LoggerLogrus {
	logger := logrus.New()

	logr := &LoggerLogrus{
		logger: logger,
	}

	setupLogger(logr)

	return logr
}

func (l *LoggerLogrus) Println(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *LoggerLogrus) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *LoggerLogrus) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l *LoggerLogrus) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *LoggerLogrus) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}
