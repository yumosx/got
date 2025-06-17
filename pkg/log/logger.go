package log

import (
	"fmt"

	"github.com/yumosx/got/pkg/stream"
	"go.uber.org/zap"
)

type Record struct {
	Key   string
	Value any
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

func NewZapLogger(cfg zap.Config) (*ZapLogger, error) {
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{log: logger}, nil
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

func (z *ZapLogger) Errorf(key string, value error, args ...Record) {
	z.log.Error(fmt.Sprintf("%s %s", key, value.Error()), z.toArgs(args)...)
}

func (z *ZapLogger) toArgs(args []Record) []zap.Field {
	ans := stream.Map[Record, zap.Field](args, func(idx int, src Record) zap.Field {
		return zap.Any(args[idx].Key, args[idx].Value)
	})
	return ans
}
