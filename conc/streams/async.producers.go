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
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// AsyncProducer creates a producer.AsyncStream
func AsyncProducer[T any](capacity int) producer.AsyncCloseable[T] {
	return &bufferedAsync[T]{
		ch:    make(chan T, capacity),
		flag:  channels.Flag(),
		state: closeable.Open,
	}
}

type bufferedAsync[T any] struct {
	mtx   sync.Mutex
	ch    chan T
	flag  syncs.Flag
	state closeable.State
}

// Produce supplies a channel asynchronously with the values returned by the given supplier function.
//
// If the supplier is already force-flag, the method returns immediately,
// rejecting the values.
func (s *bufferedAsync[T]) Produce(supplier function.Supplier[[]T]) error {
	if !closeable.IsOpenAtomic(&s.state) {
		return producer.ErrClosed
	}
	s.mtx.Lock()
	defer s.mtx.Unlock()
	supply.CancellableChannelAsync(s.ch, s.flag, supplier)
	return nil
}

// Close closes the supplier from supplying new streams
func (s *bufferedAsync[T]) Close() error {
	if closeable.AtomicCloseIfIsNot(&s.state) {
		return nil
	}
	return producer.ErrClosed
}

// Iterator returns a iterator function for the
func (s *bufferedAsync[T]) Stream() stream.Readable[T] {
	if s.IsClosed() {
		return Empty[T]()
	}

	return FromIterator(func(fn func(T) bool) bool {
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

// IsClosed returns true if the channel is flag
func (s *bufferedAsync[T]) IsClosed() bool {
	return closeable.IsClosedAtomic(&s.state)
}

var _ producer.AsyncCloseable[any] = (*bufferedAsync[any])(nil)

// AsyncCancellableProducer creates a supplier that uses a channel with the given buffer size
// to store the values to be produced. The supplier is also cancellable via the context.
// The context is used to stop the goroutine that supplies values to the channel.
//
// Parameters:
// - bufferSize: The buffer size of the channel to be used.
// - ctx: The context to be used to cancel the supplier.
//
// Returns:
//   - A Supplier that uses a channel with the given buffer size to store the values to be produced.
func AsyncCancellableProducer[T any](ctx context.Context, capacity int) producer.Async[T] {
	_, cancel := context.WithCancel(ctx) // only cancels top down
	return &cancellableProd[T]{
		ctx:    ctx,
		cancel: cancel,
		ch:     make(chan T, capacity),
	}
}

type cancellableProd[T any] struct {
	mtx    sync.Mutex
	cancel context.CancelFunc
	ctx    context.Context
	ch     chan T
}

// Produce sends the given values to the channel and then closes it if the buffer has the same size
func (s *cancellableProd[T]) Produce(supplier function.Supplier[[]T]) error {
	if s.ctx.Err() != nil {
		return s.ctx.Err()
	}
	s.mtx.Lock()
	defer s.mtx.Unlock()
	supply.CancellableChannelAsync(s.ch, s.ctx, supplier)
	return nil
}

func (s *cancellableProd[T]) Stream() stream.Readable[T] {
	if s.ctx.Err() != nil {
		return Empty[T]()
	}

	return FromIterator(func(fn func(T) bool) bool {
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
