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
	"github.com/yumosx/got/pkg/saramax"
)

type ConsumerGroupHandler struct {
	msgHandle func(msg *sarama.ConsumerMessage)
	errHandle func(err error)
}

func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg := <-claim.Messages():
			if msg == nil {
				return nil
			}
			if h.msgHandle != nil {
				h.msgHandle(msg)
			}
			// 标记消息已处理，Kafka 会自动提交偏移量
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func main() {
	brokers := []string{"localhost:9092"}
	groupID := "test-consumer-group" // 消费者组ID
	topic := "test-topic"

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Group.Session.Timeout = 10 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 3 * time.Second

	consumerGroup, err := saramax.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("创建消费者组失败: %v", err)
	}
	defer func() {
		if err := consumerGroup.Close(); err != nil {
			log.Printf("关闭消费者组失败: %v", err)
		}
	}()

	// 创建上下文，用于控制消费者生命周期
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 设置信号处理
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 启动信号监听 goroutine
	go func() {
		<-sigchan
		fmt.Println("接收到退出信号，正在关闭消费者...")
		cancel()
	}()

	// 定义消息处理函数
	msgHandler := func(msg *sarama.ConsumerMessage) {
		fmt.Printf("接收到消息: 主题=%s 分区=%d 偏移量=%d 键=%s 值=%s 时间戳=%v\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value), msg.Timestamp)
		time.Sleep(100 * time.Millisecond)
	}

	// 定义错误处理函数
	errHandler := func(err error) {
		log.Printf("消费者错误: %v", err)
	}

	// 创建消费者组处理器
	handler := &ConsumerGroupHandler{
		msgHandle: msgHandler,
		errHandle: errHandler,
	}

	// 开始消费消息
	fmt.Printf("开始消费主题: %s, 消费者组: %s\n", topic, groupID)

	// 启动消费者组
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := consumerGroup.Consume(ctx, []string{topic}, handler); err != nil {
					log.Printf("消费错误: %v", err)
					time.Sleep(5 * time.Second) // 错误后等待5秒再重试
				}
			}
		}
	}()

	// 等待上下文取消
	<-ctx.Done()
	fmt.Println("消费者已关闭")
}
