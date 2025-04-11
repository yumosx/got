package cache

import (
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

type ConfigOption interface {
	Option(cache *Config)
}

type ConfigOptionFunc func(cache *Config)

func (fn ConfigOptionFunc) Option(cache *Config) {
	fn(cache)
}

func WithAddr(addr string) ConfigOptionFunc {
	return ConfigOptionFunc(func(cache *Config) {
		cache.Addr = addr
	})
}

func WithPassword(password string) ConfigOptionFunc {
	return ConfigOptionFunc(func(cache *Config) {
		cache.Password = password
	})
}

func WithDB(db int) ConfigOptionFunc {
	return ConfigOptionFunc(func(cache *Config) {
		cache.DB = db
	})
}

func NewCacheConfig(options ...ConfigOption) *Config {
	config := &Config{}
	for _, opt := range options {
		opt.Option(config)
	}
	return config
}

// NewCache 初始化对应的cache
func NewCache(config *Config) redis.Cmdable {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	return rdb
}
