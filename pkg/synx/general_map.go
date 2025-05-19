package synx

import "sync"

type SyncMap[K comparable, V any] struct {
	m sync.Map
}
