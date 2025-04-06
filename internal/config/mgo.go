package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InitMgo(url string, dbName string, collection string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 我们可以初始化一个监控器
	monitor := &event.CommandMonitor{
		// 当命令开始执行
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			fmt.Println(startedEvent.Command)
		},
		// 当命令执行成功
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
			fmt.Println(succeededEvent.CommandName)
		},
		Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {
			fmt.Println(failedEvent.CommandName)
		},
	}

	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI(url).
			SetMonitor(monitor))
	if err != nil {
		panic(err)
	}

	return client.Database(dbName)
}
