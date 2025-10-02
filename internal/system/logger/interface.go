package logger

type Logger interface {
	Error(string, error)
	Warn(string, ...any)
	Info(string, ...any)
	Debug(string, ...any)
}
