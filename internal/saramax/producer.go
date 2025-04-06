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

func (producer *KafkaProducer) PushEntry(topic string, entry sarama.Encoder) {
	producer.Producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: entry,
	}
}

// ErrorHandler 这个函数主要的作用就是用来判断当前的 error 是不是空的
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
