package lib

import "go.uber.org/zap"

type Filed struct {
	key   string
	value any
}

type Logger interface {
	Debug(msg string, args ...Filed)
	Info(msg string, args ...Filed)
	Warn(key string, args ...Filed)
	Error(key string, args ...Filed)
}

type ZapLogger struct {
	log *zap.Logger
}

func (z *ZapLogger) Debug(msg string, args ...Filed) {
	z.log.Debug(msg, z.toArgs(args)...)
}

func (z *ZapLogger) Info(msg string, args ...Filed) {
	z.log.Info(msg, z.toArgs(args)...)
}

func (z *ZapLogger) Warn(key string, args ...Filed) {
	z.log.Info(key, z.toArgs(args)...)
}

func (z *ZapLogger) Error(key string, args ...Filed) {
	z.log.Error(key, z.toArgs(args)...)
}

func (z *ZapLogger) toArgs(args []Filed) []zap.Field {
	ans := Map[Filed, zap.Field](args, func(idx int, src Filed) zap.Field {
		return zap.Any(args[idx].key, args[idx].value)
	})
	return ans
}

func NewZapLogger(l *zap.Logger) *ZapLogger {
	return &ZapLogger{}
}
