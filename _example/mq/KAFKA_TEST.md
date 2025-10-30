1. `docker-compose.yml` - Kafka环境配置文件，包含Zookeeper、Kafka和Kafka UI服务

## 快速开始

直接运行:
```bash
docker-compose up -d
```

等待约30秒，确保Kafka完全启动。

### 2. 验证Kafka环境

访问Kafka UI: http://localhost:8080

### 3. 运行测试程序

**启动消费者 (在新终端中):**
```bash
go run test/consumer.go
```

**启动生产者 (在另一个新终端中):**
```bash
go run test/producer.go
```

### 4. 观察测试结果

消费者将开始接收生产者发送的消息，并在控制台输出消息内容。

## 测试说明

### 消费者测试

消费者程序使用`pkg/config/mq`包中的`NewConsumer`函数创建消费者实例，并配置以下选项:
- `WithConsumerOffsetsInitial(sarama.OffsetOldest)` - 从最早的消息开始消费

消费者会订阅`test-topic`主题，并打印接收到的消息内容。

### 生产者测试

生产者程序使用`pkg/config/mq`包中的`NewKafkaSyncProducer`函数创建同步生产者实例，并配置以下选项:
- `WithProducerPartitioner(sarama.NewRandomPartitioner)` - 使用随机分区策略
- `WithProducerRetryMax(3)` - 设置最大重试次数为3

生产者每2秒向`test-topic`主题发送一条消息。

## 停止Kafka环境

```bash
docker-compose down
```

## 注意事项

1. 确保Docker已安装并正在运行
2. 确保端口9092(Kafka)、2181(Zookeeper)和8080(Kafka UI)未被占用
3. 测试程序需要Go环境支持
4. 如果遇到连接问题，请等待Kafka完全启动后再运行测试程序