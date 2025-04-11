package mq

import (
	"github.com/IBM/sarama"
	"time"
)

type Config struct {
	flushTime time.Duration
	brokers   []string
}

type KafkaConfigOption interface {
	Option(config *Config)
}

type KafkaConfig func(config *Config)

func (fn KafkaConfig) Option(config *Config) {
	fn(config)
}

func NewKafkaProducer(config *Config) error {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Flush.Frequency = config.flushTime
	// 设置消息确认模式
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	version, err := sarama.ParseKafkaVersion("2.4.1")
	if err != nil {
		return err
	}
	kafkaConfig.Version = version
	kafkaConfig.Admin.Timeout = 10 * time.Second

	return nil
}
