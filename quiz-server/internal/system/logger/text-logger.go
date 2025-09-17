package logger

import (
	"fmt"
)

type textLogger struct{}

func newTextLogger() *textLogger {
	return &textLogger{}
}

func (l *textLogger) Error(message string, _ ...any) {
	fmt.Println(l.format("[error]: ", message))
}

func (l *textLogger) Warn(message string, _ ...any) {
	fmt.Println(l.format("[warn]:  ", message))
}

func (l *textLogger) Info(message string, _ ...any) {
	fmt.Println(l.format("[info]:  ", message))
}

func (l *textLogger) Debug(message string, _ ...any) {
	fmt.Println(l.format("[debug]: ", message))
}

func (l *textLogger) format(tag string, message string) string {
	return fmt.Sprintf("%s%s", tag, message)
}
