package mq

import "github.com/IBM/sarama"

// NewKafkaProducer 初始化kafka 配置
// 初始化的配置 max(发送者总速度/单一分区写入速度, 发送者总速度/单一消费者速度)+buffer
func NewKafkaProducer() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	return config
}
