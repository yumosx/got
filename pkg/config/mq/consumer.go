package mq

import (
	"github.com/IBM/sarama"
	"github.com/yumosx/got/pkg/errx"
)

func NewConsumer(addrs []string, options ...ConfigOption) errx.Option[sarama.Consumer] {
	config := sarama.NewConfig()
	for _, opt := range options {
		opt.Option(config)
	}
	consumer, err := sarama.NewConsumer(addrs, config)
	if err != nil {
		return errx.Err[sarama.Consumer](err)
	}
	return errx.Ok(consumer)
}
