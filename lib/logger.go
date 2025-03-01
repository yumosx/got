package lib

type Filed struct {
	key   string
	value any
}

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(key string, value ...Filed)
	Error(key string, value ...Filed)
}
