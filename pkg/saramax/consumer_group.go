package saramax

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

type KafkaConsumerGroup struct {
	consumerGroup sarama.ConsumerGroup
}

func NewConsumerGroup(brokers []string, groupID string, config *sarama.Config) (*KafkaConsumerGroup, error) {
	if config == nil {
		config = sarama.NewConfig()
	}

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer group: %w", err)
	}

	return &KafkaConsumerGroup{
		consumerGroup: consumerGroup,
	}, nil
}

func (cg *KafkaConsumerGroup) Consume(
	ctx context.Context,
	topics []string,
	handler sarama.ConsumerGroupHandler,
) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := cg.consumerGroup.Consume(ctx, topics, handler); err != nil {
				return fmt.Errorf("error from consumer: %w", err)
			}
		}
	}
}

func (cg *KafkaConsumerGroup) Close() error {
	return cg.consumerGroup.Close()
}
