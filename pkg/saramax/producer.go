package saramax

import (
	"context"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Producer sarama.AsyncProducer
}

func NewProducer(producer sarama.AsyncProducer) *KafkaProducer {
	return &KafkaProducer{Producer: producer}
}

func (producer *KafkaProducer) PushEntry(ctx context.Context, topic string, entry sarama.Encoder) {
	select {
	case <-ctx.Done():
		return
	default:
		producer.Producer.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Value: entry,
		}
	}
}

func (producer *KafkaProducer) Error(ctx context.Context, errHandle func(err error)) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				errHandle(ctx.Err())
				return
			case err, ok := <-producer.Producer.Errors():
				if !ok {
					return
				}
				errHandle(err)
			}
		}
	}()
}

func (producer *KafkaProducer) Close() error {
	return producer.Producer.Close()
}
