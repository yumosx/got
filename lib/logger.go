package lib

type Filed struct {
	key   string
	value any
}

type Logger interface {
	Error(key string, value ...Filed)
	Warn(key string, value ...Filed)
}
