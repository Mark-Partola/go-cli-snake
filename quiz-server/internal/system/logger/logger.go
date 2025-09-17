package logger

import (
	"fmt"
	"sync"
)

type loggerInstance interface {
	Error(string, ...any)
	Warn(string, ...any)
	Info(string, ...any)
	Debug(string, ...any)
}

type logger struct {
	tag    string
	logger loggerInstance
}

var once sync.Once
var instance loggerInstance

func Setup(env string) Logger {
	instantiate := func() loggerInstance {
		switch env {
		case "dev":
			return newTextLogger()
		default:
			return newJsonLogger()
		}
	}

	once.Do(func() {
		instance = instantiate()
	})

	return Get("setup")
}

func Get(tag string) Logger {
	return logger{
		tag:    tag,
		logger: instance,
	}
}

func (l logger) Error(message string, err error) {
	l.logger.Error(l.format(message, err.Error()))
}

func (l logger) Warn(message string, args ...any) {
	l.logger.Warn(l.format(message, args...))
}

func (l logger) Info(message string, args ...any) {
	l.logger.Info(l.format(message, args...))
}

func (l logger) Debug(message string, args ...any) {
	l.logger.Debug(l.format(message, args...))
}

func (l logger) format(message string, args ...any) string {
	message = fmt.Sprintf(message, args...)

	return fmt.Sprintf("[%s]: %s", l.tag, message)
}
