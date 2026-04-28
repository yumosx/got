package lv_redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yumosx/oneframe/lv_conf"
	"github.com/yumosx/oneframe/lv_log"
)

var (
	redisClient *RedisClient
)

type RedisClient struct {
	c *redis.Client
}

// // 获取缓存单例
func GetInstance(indexDb int) *RedisClient {
	if redisClient == nil {
		redisClient = NewRedisClient(indexDb)
	}
	return redisClient
}

// 获取缓存单例
func NewRedisClient(indexDb int) *RedisClient {
	conf := lv_conf.ConfigDefault{}
	addr := conf.GetValueStr("go.redis.host")
	port := conf.GetValueStr("go.redis.port")
	password := conf.GetValueStr("go.redis.password")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password, // 没有密码，默认值
		DB:       indexDb,  // 默认DB 0
	})
	redisClient = new(RedisClient)
	redisClient.c = rdb
	if redisClient.c.Ping(context.Background()).Val() == "" {
		msg := ` 
			  ------------>连接 reids 错误：
			  无法链接到redis!!!! 检查相关配置:
			  host: %v
			  port: %v
			  password: %v
             `
		host := conf.GetValueStr("go.redis.host")
		lv_log.Error(fmt.Sprintf(msg, host, conf.GetValueStr("go.redis.port"), conf.GetValueStr("go.redis.password")))
		panic("redis 错误:" + host + " port:" + port)
	}
	return redisClient
}

func (rcc *RedisClient) HMSet(ctx context.Context, key string, mp map[string]any, expiration time.Duration) error {
	err := rcc.c.HSet(ctx, key, mp).Err()
	err = rcc.c.Expire(ctx, key, expiration).Err()
	return err
}

func (rcc *RedisClient) Expire(ctx context.Context, key string, duration time.Duration) error {
	return rcc.c.Expire(ctx, key, duration).Err()
}

func (rcc *RedisClient) Exists(ctx context.Context, key string) (int64, error) {
	return rcc.c.Exists(ctx, key).Result()
}

func (rcc *RedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	rcc.c.Set(ctx, key, value, expiration)
	return nil
}

func (rcc *RedisClient) Get(ctx context.Context, key string) (data string, err error) {
	data, err = rcc.c.Get(ctx, key).Result()
	return data, err
}

func (rcc *RedisClient) Del(ctx context.Context, keys ...string) error {
	var err error = nil
	for _, key := range keys {
		err = rcc.c.Del(ctx, key).Err()
	}
	return err
}

func (rcc *RedisClient) HSet(ctx context.Context, key string, values ...any) error {
	err := rcc.c.HSet(ctx, key, values...).Err()
	return err
}

func (rcc *RedisClient) HGet(ctx context.Context, key, field string) (string, error) {
	data, err := rcc.c.HGet(ctx, key, field).Result()
	return data, err
}

func (rcc *RedisClient) HDel(ctx context.Context, key string, fields ...string) error {
	return rcc.c.HDel(ctx, key, fields...).Err()
}

func (rcc *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return rcc.c.HGetAll(ctx, key).Result()
}

func (rcc *RedisClient) Close() {
	rcc.c.Close()
}
