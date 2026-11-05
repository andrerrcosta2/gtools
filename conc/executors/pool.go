package executors

import (
	"sync"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/conc/core/executor"
)

type Handler[T any] func(T)

// NewPool creates and starts a queues-backed worker pool.
func NewPool[T any](workers, queueSize int, handler Handler[T]) executor.QueueExecutor[T] {
	if workers <= 0 {
		workers = 1
	}
	if queueSize < 0 {
		queueSize = 0
	}
	if handler == nil {
		handler = func(T) {}
	}

	p := &pool[T]{
		ch:      make(chan T, queueSize),
		handler: handler,
	}
	p.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go p.run()
	}
	return p
}

type pool[T any] struct {
	ch      chan T
	handler Handler[T]
	closed  atomic.Bool
	once    sync.Once
	wg      sync.WaitGroup
}

func (p *pool[T]) Submit(item T) (err error) {
	defer func() {
		if recover() != nil {
			err = executor.ErrExecutorClosed
		}
	}()

	if p.IsClosed() {
		return executor.ErrExecutorClosed
	}
	p.ch <- item
	return nil
}

func (p *pool[T]) Close() error {
	if p.closed.CompareAndSwap(false, true) {
		p.once.Do(func() { close(p.ch) })
		return nil
	}
	return executor.ErrExecutorClosed
}

func (p *pool[T]) IsClosed() bool {
	return p.closed.Load()
}

func (p *pool[T]) Wait() {
	p.wg.Wait()
}

func (p *pool[T]) run() {
	defer p.wg.Done()
	for item := range p.ch {
		p.handler(item)
	}
}
