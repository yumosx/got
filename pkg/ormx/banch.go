package ormx

import (
	"github.com/yumosx/got/pkg/errx"
	"gorm.io/gorm"
)

type BatchConfig struct {
	batchSize  int
	maxRecords int
	total      int
	db         *gorm.DB
}

func NewBatchConfig(total int, db *gorm.DB, options ...BatchConfigOption) BatchConfig {
	var config BatchConfig
	config.db = db
	config.total = total

	for _, option := range options {
		option.Option(&config)
	}
	return config
}

type BatchConfigOption interface {
	Option(option *BatchConfig)
}

type BatchConfigOptionFunc func(option *BatchConfig)

func (fn BatchConfigOptionFunc) Option(config *BatchConfig) {
	fn(config)
}

// WithBatch 表示一批导出多少条数据
func WithBatch(batch int) BatchConfigOption {
	return BatchConfigOptionFunc(func(option *BatchConfig) {
		option.batchSize = batch
	})
}

// WithMaxRecords 表示最大的导出次数
func WithMaxRecords(maxRecords int) BatchConfigOption {
	return BatchConfigOptionFunc(func(option *BatchConfig) {
		option.maxRecords = maxRecords
	})
}

// BatchExport 分批次导出对应的数据
func BatchExport[Result any](config BatchConfig, query func(offset, currentLimit int, db *gorm.DB) errx.Option[[]Result]) errx.Option[[]Result] {
	total := min(config.total, config.maxRecords)

	list := make([]Result, 0, total)
	for offset := 0; offset < total; offset += config.batchSize {
		currentLimit := config.batchSize
		if offset+config.batchSize > total {
			currentLimit = total - offset
		}
		result := query(offset, currentLimit, config.db)
		if result.NoNil() {
			return errx.Err[[]Result](result.Error())
		}
		list = append(list, result.Val...)
	}

	return errx.Ok(list)
}
