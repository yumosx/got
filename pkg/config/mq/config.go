package mq

import (
	"github.com/IBM/sarama"
)

type ConfigOption interface {
	Option(config *sarama.Config)
}

type OptionFunc func(config *sarama.Config)

func (f OptionFunc) Option(config *sarama.Config) {
	f(config)
}

// WithProducerPartitioner 设置生产者分区策略
func WithProducerPartitioner(partitioner sarama.PartitionerConstructor) ConfigOption {
	return OptionFunc(func(config *sarama.Config) {
		config.Producer.Partitioner = partitioner
	})
}

// WithProducerRetryMax 设置最大重试次数
func WithProducerRetryMax(retryMax int) ConfigOption {
	return OptionFunc(func(config *sarama.Config) {
		config.Producer.Retry.Max = retryMax
	})
}

// WithConsumerOffsetsInitial 设置消费者开始偏移量
func WithConsumerOffsetsInitial(offset int64) ConfigOption {
	return OptionFunc(func(config *sarama.Config) {
		if offset == sarama.OffsetOldest {
			config.Consumer.Offsets.Initial = sarama.OffsetOldest
		} else {
			config.Consumer.Offsets.Initial = sarama.OffsetNewest
		}
	})
}
