package mq

import (
	"github.com/IBM/sarama"
)

// NewKafkaSyncProducer 初始化kafka 配置
// 初始化的配置 max(发送者总速度/单一分区写入速度, 发送者总速度/单一消费者速度)+buffer
func NewKafkaSyncProducer(addrs []string, options ...ConfigOption) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	for _, opt := range options {
		opt.Option(config)
	}

	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// NewKafkaASyncProducer 异步生产者, 无需等待
func NewKafkaASyncProducer(addrs []string, options ...ConfigOption) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	for _, opt := range options {
		opt.Option(config)
	}
	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		return nil, err
	}
	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
