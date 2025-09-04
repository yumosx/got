package cache

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yumosx/got/internal/errs"
)

type Cache struct {
	client redis.Cmdable
}

func NewCache(client redis.Cmdable) *Cache {
	return &Cache{client: client}
}

// Set 设置一个键值对, 并且设置过期时间, expire 为 0的 时候表示不过期
func (c *Cache) Set(ctx context.Context, key string, value any, expire time.Duration) error {
	return c.client.Set(ctx, key, value, expire).Err()
}

// SetNX 设置一个键值对
func (c *Cache) SetNX(ctx context.Context, key string, value any, expire time.Duration) (bool, error) {
	result, err := c.client.SetNX(ctx, key, value, expire).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	result, err := c.client.Get(ctx, key).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		return "", errs.ErrKeyNotExists
	}
	return result, err
}

func (c *Cache) Delete(ctx context.Context, key string) (int64, error) {
	result, err := c.client.Del(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

// LPush 将所指定的值插入存储列表的头部
func (c *Cache) LPush(ctx context.Context, key string, val ...any) (int64, error) {
	res, err := c.client.LPush(ctx, key, val...).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}

// LPop 从列表中删除对应的元素
func (c *Cache) LPop(ctx context.Context, key string) (string, error) {
	val, err := c.client.LPop(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errs.ErrKeyNotExists
		}
		return "", err
	}
	return val, nil
}

// LLen 用来计算当前列表的长度
func (c *Cache) LLen(ctx context.Context, key string) (int64, error) {
	result, err := c.client.LLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}
