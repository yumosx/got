package saramax

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func NewConsumer(consumer sarama.Consumer) *KafkaConsumer {
	return &KafkaConsumer{
		consumer: consumer,
	}
}

func (c *KafkaConsumer) ConsumePartition(
	ctx context.Context,
	topic string,
	partition int32,
	offset int64,
	msgHandle func(msg *sarama.ConsumerMessage),
	errHandle func(err error),
) error {
	partitionConsumer, err := c.consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		return fmt.Errorf("failed to consume partition: %w", err)
	}

	go func() {
		defer partitionConsumer.AsyncClose()

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-partitionConsumer.Messages():
				if msgHandle != nil {
					msgHandle(msg)
				}
			case err := <-partitionConsumer.Errors():
				if errHandle != nil {
					errHandle(err)
				}
			}
		}
	}()

	return nil
}

func (c *KafkaConsumer) Close() error {
	return c.consumer.Close()
}
