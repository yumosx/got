//go:build go1.24

package synx

import (
	"runtime"
	"sync"
	"weak"
)

type Cache[K comparable, V any] struct {
	create func(K) (*V, error)
	m      sync.Map
}

// NewCache 创建一个 cache
func NewCache[K comparable, V any](create func(K) (*V, error)) *Cache[K, V] {
	return &Cache[K, V]{create: create}
}

func (c *Cache[K, V]) Get(key K) (*V, error) {
	var newValue *V

	for {
		value, ok := c.m.Load(key)
		if !ok {
			if newValue == nil {
				var err error
				newValue, err = c.create(key)
				if err != nil {
					return nil, err
				}
			}
			wp := weak.Make(newValue)
			var loaded bool
			if !loaded {
				runtime.AddCleanup(newValue, func(k K) {
					c.m.CompareAndDelete(key, wp)
				}, key)
			}
			return newValue, nil
		}

		if mf := value.(weak.Pointer[V]).Value(); mf != nil {
			return mf, nil
		}
		c.m.CompareAndDelete(key, value)
	}
}
