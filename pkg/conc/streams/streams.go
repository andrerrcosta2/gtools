// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"context"
	"errors"
	"github.com/andrerrcosta2/gtools/conc/channels"
	"github.com/andrerrcosta2/gtools/conc/producers/streamables"
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"sync/atomic"
)

// HotCloseable returns a gtools.CloseableStreamable that is filled with the given values and starts streaming immediately.
// It implements a gtools.CloseableStreamable so once it is closed it cannot be opened or supply any more values.
// It's thread-safe and can be used concurrently.
func HotCloseable[T any](bufferSize int, values ...T) gtools.CloseableStreamable[T] {
	s := &hotStreamableChannel[T]{
		// The channel is used to stream values.
		ch: make(chan T, bufferSize),
		// The flag is used to signal the channel must close.
		flag: channels.NewFlag(),
	}

	// Open the channel with the given values.
	streamables.OpenCloseableSync(values, s.ch, s.flag.Receiver())

	// Start streaming immediately.
	return s
}

// AsyncHotCloseable returns a gtools.CloseableStreamable that is filled with the given supplier function asynchronously
// and starts streaming immediately. It's thread-safe and can be used concurrently.
func AsyncHotCloseable[T any](bufferSize int, supplier functions.Supplier[[]T]) gtools.CloseableStreamable[T] {
	s := &hotStreamableChannel[T]{
		// The channel is used to stream values.
		ch: make(chan T, bufferSize),
		// The flag is used to signal the channel must close.
		flag: channels.NewFlag(),
	}

	// Open the channel with the given supplier function asynchronously.
	streamables.OpenCloseableAsync(s.ch, s.flag.Receiver(), supplier)

	// Start streaming immediately.
	return s
}

type hotStreamableChannel[T any] struct {
	ch     chan T
	flag   *channels.Flag
	closed atomic.Bool
}

// Close closes the stream and the channel.
// It's thread-safe and can be used concurrently.
func (c *hotStreamableChannel[T]) Close() error {
	if c.closed.CompareAndSwap(false, true) {
		c.flag.Signal()
		return nil
	}
	return errors.New("stream already closed")
}

// IsClosed returns true if the channel is closed, false otherwise.
func (c *hotStreamableChannel[T]) IsClosed() bool {
	return c.closed.Load()
}

// Stream returns a channel that can be used to receive values of type T.
// This is an immediate-closed channel.
//
// It's thread-safe and can be used concurrently.
func (c *hotStreamableChannel[T]) Stream() gtools.Stream[T] {
	if !c.IsClosed() {
		return c.ch
	}
	return channels.Closed[T]()
}

var _ gtools.CloseableStreamable[any] = (*hotStreamableChannel[any])(nil)

// HotCancellable creates a new gtools.Streamable with the specified buffer size and given values.
// It's thread-safe and can be used concurrently.
//
// The given values are sent to the channel when the channel is opened.
// The channel is closed when all values have been sent or when the context is canceled.
func HotCancellable[T any](ctx context.Context, bufferSize int, values ...T) gtools.Streamable[T] {
	s := &hotContextStreamable[T]{
		ch:  make(chan T, bufferSize),
		ctx: ctx,
	}
	// Open the channel with the given values and context.
	streamables.OpenCloseableSync(values, s.ch, s.ctx.Done())
	return s
}

// HotCancellableAsync creates a new gtools.Streamable with the specified buffer size and given supplier function.
// It's thread-safe and can be used concurrently.
//
// The supplier function is called asynchronously to generate values.
// The channel is closed when all values have been sent or when the context is canceled.
func HotCancellableAsync[T any](ctx context.Context, bufferSize int, supplier functions.Supplier[[]T]) gtools.Streamable[T] {
	s := &hotContextStreamable[T]{
		ch:  make(chan T, bufferSize),
		ctx: ctx,
	}
	// Open the channel with the given supplier function and context.
	streamables.OpenCloseableAsync(s.ch, s.ctx.Done(), supplier)
	return s
}

type hotContextStreamable[T any] struct {
	ch  chan T
	ctx context.Context
}

// Stream returns a channel that can be used to receive values of type T.
// The channel is closed when all values have been sent or when the context is canceled.
func (c *hotContextStreamable[T]) Stream() gtools.Stream[T] {
	return c.ch
}

var _ gtools.Streamable[any] = (*hotContextStreamable[any])(nil)
