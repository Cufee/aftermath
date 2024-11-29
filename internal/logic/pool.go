package logic

import "sync"

type pool[T any] struct {
	pool *sync.Pool
}

func (p *pool[T]) Get() *T {
	item, _ := p.pool.Get().(*T)
	if item == nil {
		item = new(T)
	} else {
		// Clear object
		var zero T
		*item = zero
	}
	return item
}

func (p *pool[T]) Put(s *T) {
	p.pool.Put(s)
}

func newPool[T any]() *pool[T] {
	return &pool[T]{
		pool: &sync.Pool{
			New: func() any {
				return new(T)
			},
		},
	}
}
