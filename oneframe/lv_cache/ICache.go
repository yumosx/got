package lv_cache

import (
	"context"
	"time"

	"github.com/yumosx/oneframe/lv_cache/lv_ram"
	"github.com/yumosx/oneframe/lv_cache/lv_redis"
	"github.com/yumosx/oneframe/lv_global"
)

type ICache interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (value string, err error)
	Del(ctx context.Context, key ...string) error

	HSet(ctx context.Context, key string, values ...any) error
	HMSet(ctx context.Context, key string, mp map[string]any, duration time.Duration) error
	HGet(ctx context.Context, key, field string) (string, error)
	HDel(ctx context.Context, key string, fields ...string) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Exists(ctx context.Context, key string) (int64, error)
	Close()
	Expire(ctx context.Context, key string, duration time.Duration) error
}

var cacheClient ICache //主数据库

func GetCacheClient() ICache {
	if cacheClient == nil {
		var config = lv_global.Config()
		cacheType := config.GetVipperCfg().GetString("go.cache")
		if cacheType == "redis" {
			cacheClient = lv_redis.GetInstance(0)
		} else {
			cacheClient = lv_ram.GetRamCacheClient()
		}
	}
	return cacheClient
}
