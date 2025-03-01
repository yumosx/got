package config

import (
	"github.com/redis/go-redis/v9"
)

func InitCache(addr string, password string, db int) redis.Cmdable {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return rdb
}
