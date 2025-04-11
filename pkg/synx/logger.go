package synx

import (
	"github.com/yumosx/got/pkg/stream"
	"go.uber.org/zap"
)

type Record struct {
	key   string
	value any
}

type Logger interface {
	Debug(msg string, args ...Record)
	Info(msg string, args ...Record)
	Warn(key string, args ...Record)
	Error(key string, args ...Record)
}

type ZapLogger struct {
	log *zap.Logger
}

func NewZapLogger(cfg zap.Config) *zap.Logger {
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func (z *ZapLogger) Debug(msg string, args ...Record) {
	z.log.Debug(msg, z.toArgs(args)...)
}

func (z *ZapLogger) Info(msg string, args ...Record) {
	z.log.Info(msg, z.toArgs(args)...)
}

func (z *ZapLogger) Warn(key string, args ...Record) {
	z.log.Info(key, z.toArgs(args)...)
}

func (z *ZapLogger) Error(key string, args ...Record) {
	z.log.Error(key, z.toArgs(args)...)
}

func (z *ZapLogger) toArgs(args []Record) []zap.Field {
	ans := stream.Map[Record, zap.Field](args, func(idx int, src Record) zap.Field {
		return zap.Any(args[idx].key, args[idx].value)
	})
	return ans
}
