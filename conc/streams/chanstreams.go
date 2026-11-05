// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"context"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/conc/core/channel/channels"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/streams/internal/supply"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// ColdCloseable returns a cold closeable stream backed by a buffered channel.
//
// The producer starts immediately after the stream is created, but the stream
// remains cold: values are never discarded due to the absence of a consumer.
// If the consumer is slower than the producer, production is naturally
// backpressured by the underlying channel.
//
// The stream can be closed early by calling Close().
func ColdCloseable[T any](bufferSize int, values ...T) stream.Closeable[T] {
	s := &coldChannelStream[T]{
		ch:   make(chan T, bufferSize),
		flag: channels.Flag(),
	}

	go supply.CancellableChannelSync(values, s.ch, s.flag)
	return s
}

// AsyncColdCloseable returns a cold closeable stream whose values are produced
// asynchronously by the supplied function.
//
// The producer starts immediately after the stream is created. Despite this,
// the stream remains cold: produced values are preserved through channel
// backpressure until they are consumed or the stream is closed.
//
// The stream can be closed early by calling Close().
func AsyncColdCloseable[T any](bufferSize int, supplier function.BatchSupplier[[]T, T]) stream.Closeable[T] {
	s := &coldChannelStream[T]{
		ch:   make(chan T, bufferSize),
		flag: channels.Flag(),
	}

	// Open the channel with the given supplier function asynchronously.
	go supply.BatchCancellableChannelAsync(s.ch, s.flag, supplier)
	return s
}

type coldChannelStream[T any] struct {
	ch     chan T
	flag   syncs.Flag
	closed atomic.Bool
}

// Close closes the stream and the channel.
// It's thread-safe and can be used concurrently.
func (c *coldChannelStream[T]) Close() error {
	if c.closed.CompareAndSwap(false, true) {
		c.flag.Signal()
		return nil
	}
	return stream.ErrClosed
}

// IsClosed returns true if the channel is closed, false otherwise.
func (c *coldChannelStream[T]) IsClosed() bool {
	return c.closed.Load()
}

func (c *coldChannelStream[T]) Iterator() stream.Iterator[T] {
	return func(fn func(e T) bool) bool {
		defer close(c.ch)
		for {
			value, ok := <-c.ch
			if !ok {
				return false
			}

			if !fn(value) {
				return false
			}
		}
	}
}

var _ stream.Readable[int] = (*coldChannelStream[int])(nil)

// ColdCancellable returns a cold readable stream backed by a buffered channel.
//
// The producer starts immediately after the stream is created, but the stream
// remains cold: values are preserved through channel backpressure until they
// are consumed or the provided context is canceled.
//
// When ctx is canceled, production stops and the stream is closed.
func ColdCancellable[T any](ctx context.Context, bufferSize int, values ...T) stream.Readable[T] {
	s := &coldContextStream[T]{
		ch:  make(chan T, bufferSize),
		ctx: ctx,
	}
	go supply.CancellableChannelSync(values, s.ch, s.ctx)
	return s
}

// ColdCancellableAsync returns a cold readable stream whose values are produced
// asynchronously by the supplied function.
//
// The producer starts immediately after the stream is created. Produced values
// are preserved through channel backpressure until they are consumed or the
// provided context is canceled.
//
// When ctx is canceled, production stops and the stream is closed.
func ColdCancellableAsync[T any](ctx context.Context, bufferSize int, supplier function.BatchSupplier[[]T, T]) stream.Readable[T] {
	s := &coldContextStream[T]{
		ch:  make(chan T, bufferSize),
		ctx: ctx,
	}
	go supply.BatchCancellableChannelAsync(s.ch, s.ctx, supplier)
	return s
}

type coldContextStream[T any] struct {
	ch  chan T
	ctx context.Context
}

func (c *coldContextStream[T]) Iterator() stream.Iterator[T] {
	if c.ctx.Err() != nil {
		return func(fn func(e T) bool) bool { return false }
	}

	return func(fn func(e T) bool) bool {
		defer close(c.ch)
		for {
			select {
			case <-c.ctx.Done():
				return false

			case value, ok := <-c.ch:
				if !ok {
					return false
				}

				if !fn(value) {
					return false
				}
			}
		}
	}
}
