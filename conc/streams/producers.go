// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"context"
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/channel/channels"
	"github.com/andrerrcosta2/gtools/conc/core/producer"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/streams/internal/supply"
	"github.com/andrerrcosta2/gtools/core/data/closeable"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
)

// Producer creates a producer that uses a buffered channel
func Producer[T any](capacity int) producer.SyncCloseable[T] {
	return &buffered[T]{
		ch:    make(chan T, capacity),
		flag:  channels.Flag(),
		state: closeable.Open,
	}
}

// buffered creates a supplier that uses a buffered channel
type buffered[T any] struct {
	mtx   sync.Mutex
	ch    chan T
	flag  syncs.Flag
	state closeable.State
}

// Produce sends the given values to the channel and then closes it.
// The size of the slice is limited by the capacity of the channel.
// The method returns the supplier itself to allow method chaining.
func (s *buffered[T]) Produce(values ...T) error {
	if !closeable.IsOpenAtomic(&s.state) {
		return producer.ErrClosed
	}
	s.mtx.Lock()
	defer s.mtx.Unlock()
	supply.CancellableChannelSync(values, s.ch, s.flag)
	return nil
}

// Close closes the supplier.
func (s *buffered[T]) Close() error {
	if closeable.AtomicCloseIfIsNot(&s.state) {
		return nil
	}
	return producer.ErrClosed
}

func (s *buffered[T]) Stream() stream.Readable[T] {
	if s.IsClosed() {
		return Empty[T]()
	}

	return FromIterator(func(fn func(e T) bool) bool {
		for value := range s.ch {
			if !fn(value) {
				s.flag.Signal()
				return false
			}
		}
		s.flag.Signal()
		return false
	})
}

// IsClosed returns true if the supplier is closed
func (s *buffered[T]) IsClosed() bool {
	return closeable.IsClosedAtomic(&s.state)
}

// CancellableProducer creates a buffered producer that can be cancelled by a context
func CancellableProducer[T any](ctx context.Context, capacity int) producer.Sync[T] {
	_, cancel := context.WithCancel(ctx) // only cancels top down
	return &cancellable[T]{
		ctx:    ctx,
		cancel: cancel,
		ch:     make(chan T, capacity),
	}
}

type cancellable[T any] struct {
	mtx    sync.Mutex
	cancel context.CancelFunc
	ctx    context.Context
	ch     chan T
}

func (s *cancellable[T]) Produce(values ...T) error {
	if s.ctx.Err() != nil {
		return s.ctx.Err()
	}
	s.mtx.Lock()
	defer s.mtx.Unlock()
	supply.CancellableChannelSync(values, s.ch, s.ctx)
	return nil
}

func (s *cancellable[T]) Stream() stream.Readable[T] {
	if s.ctx.Err() != nil {
		return Empty[T]()
	}
	return FromIterator(func(fn func(e T) bool) bool {
		for value := range s.ch {
			if !fn(value) {
				s.cancel()
				return false
			}
		}
		s.cancel()
		return false
	})
}
