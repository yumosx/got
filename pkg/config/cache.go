package config

import (
	"github.com/redis/go-redis/v9"
)

type CacheConfig struct {
	Addr     string
	Password string
	DB       int
}

type CacheConfigOption interface {
	Option(cache *CacheConfig)
}

type CacheConfigOptionFunc func(cache *CacheConfig)

func (fn CacheConfigOptionFunc) Option(cache *CacheConfig) {
	fn(cache)
}

func WithAddr(addr string) CacheConfigOptionFunc {
	return CacheConfigOptionFunc(func(cache *CacheConfig) {
		cache.Addr = addr
	})
}

func WithCachePassword(password string) CacheConfigOptionFunc {
	return CacheConfigOptionFunc(func(cache *CacheConfig) {
		cache.Password = password
	})
}

func WithDB(db int) CacheConfigOptionFunc {
	return CacheConfigOptionFunc(func(cache *CacheConfig) {
		cache.DB = db
	})
}

func NewCacheConfig(options ...CacheConfigOption) *CacheConfig {
	config := &CacheConfig{}
	for _, opt := range options {
		opt.Option(config)
	}
	return config
}

// NewCache 初始化对应的cache
func NewCache(config *CacheConfig) redis.Cmdable {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	return rdb
}
