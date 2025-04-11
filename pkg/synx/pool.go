package synx

import "sync"

type Pool[T any] struct {
	p sync.Pool
}

func NewPool[T any](factory func() T) *Pool[T] {
	return &Pool[T]{
		p: sync.Pool{
			New: func() any {
				return factory()
			},
		},
	}
}

func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}

func (p *Pool[T]) Put(t T) {
	p.p.Put(t)
}
