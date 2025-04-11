package cache

import "context"

type Cache interface {
	Get(ctx context.Context, key string) (any, error)
	GetMulti(ctx context.Context, keys []string) ([]any, error)
	Delete(ctx context.Context) error
	LoadAndDelete(ctx context.Context) (any, error)
	Incr(ctx context.Context) error
	Decr(ctx context.Context) error
	IsExist(ctx context.Context, key string) (bool, error)
}
