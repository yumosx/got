package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/yumosx/got/pkg/config/mq"
)

func main() {
	brokers := []string{"localhost:9092"}

	producer, err := mq.NewKafkaSyncProducer(
		brokers,
		mq.WithProducerPartitioner(sarama.NewRandomPartitioner), // 使用随机分区策略
		mq.WithProducerRetryMax(3),                              // 设置最大重试次数为3
	)
	if err != nil {
		log.Fatalf("创建生产者失败: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Printf("关闭生产者失败: %v", err)
		}
	}()

	topic := "test-topic"

	fmt.Printf("开始向主题 %s 发送消息...\n", topic)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-sigchan
		fmt.Println("接收到退出信号，正在停止发送消息...")
		cancel()
	}()

	messageCount := 0

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("总共发送了 %d 条消息\n", messageCount)
			fmt.Println("生产者已关闭")
			return

		case <-ticker.C:
			message := &sarama.ProducerMessage{
				Topic: topic,
				Key:   sarama.StringEncoder(fmt.Sprintf("key-%d", messageCount)),
				Value: sarama.StringEncoder(fmt.Sprintf("测试消息 #%d, 发送时间: %s", messageCount, time.Now().Format("2006-01-02 15:04:05"))),
			}

			partition, offset, err := producer.SendMessage(message)
			if err != nil {
				log.Printf("发送消息失败: %v", err)
				continue
			}

			messageCount++
			fmt.Printf("发送成功: 分区=%d, 偏移量=%d, 消息计数=%d\n", partition, offset, messageCount)
		}
	}
}
