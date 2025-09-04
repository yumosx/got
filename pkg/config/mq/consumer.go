package mq

import (
	"github.com/IBM/sarama"
)

func NewConsumer(addrs []string, options ...ConfigOption) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	for _, opt := range options {
		opt.Option(config)
	}
	consumer, err := sarama.NewConsumer(addrs, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
