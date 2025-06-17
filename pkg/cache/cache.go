package cache

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yumosx/got/internal/errs"
	"github.com/yumosx/got/pkg/errx"
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
func (c *Cache) SetNX(ctx context.Context, key string, value any, expire time.Duration) errx.Option[bool] {
	result, err := c.client.SetNX(ctx, key, value, expire).Result()
	if err != nil {
		return errx.Err[bool](err)
	}
	return errx.Ok(result)
}

func (c *Cache) Get(ctx context.Context, key string) errx.Option[string] {
	result, err := c.client.Get(ctx, key).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		return errx.Err[string](errs.ErrKeyNotExists)
	}
	return errx.Ok(result)
}

func (c *Cache) Delete(ctx context.Context, key string) errx.Option[int64] {
	result, err := c.client.Del(ctx, key).Result()
	if err != nil {
		return errx.Err[int64](err)
	}
	return errx.Ok(result)
}

// LPush 将所指定的值插入存储列表的头部
func (c *Cache) LPush(ctx context.Context, key string, val ...any) errx.Option[int64] {
	res, err := c.client.LPush(ctx, key, val...).Result()
	if err != nil {
		return errx.Err[int64](err)
	}
	return errx.Ok(res)
}

// LPop 从列表中删除对应的元素
func (c *Cache) LPop(ctx context.Context, key string) errx.Option[string] {
	val, err := c.client.LPop(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return errx.Err[string](errs.ErrKeyNotExists)
		}
		return errx.Err[string](err)
	}
	return errx.Ok(val)
}

// LLen 用来计算当前列表的长度
func (c *Cache) LLen(ctx context.Context, key string) errx.Option[int64] {
	result, err := c.client.LLen(ctx, key).Result()
	if err != nil {
		return errx.Err[int64](err)
	}
	return errx.Ok(result)
}
