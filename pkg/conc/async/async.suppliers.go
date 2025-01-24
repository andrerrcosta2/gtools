// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"context"
	"errors"
	"github.com/andrerrcosta2/gtools/conc/channels"
	"github.com/andrerrcosta2/gtools/conc/streams"
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"sync"
	"sync/atomic"
)

// BufferedSupplier creates a gtools.AsyncCloseableSupplier that uses a channel with the given buffer size
// to store the values to be produced.
//
// Parameters:
// - bufferSize: The buffer size of the channel to be used.
//
// Returns:
//   - A Supplier that uses a channel with the given buffer size to store the values to be produced.
func BufferedSupplier[T any](capacity int) gtools.AsyncCloseableSupplier[T] {
	return &bufferedChannelSupplier[T]{
		capacity: capacity,
	}
}

type bufferedChannelSupplier[T any] struct {
	mtx      sync.RWMutex
	supplier functions.Supplier[[]T]
	capacity int
	closed   atomic.Bool
}

// Supply supplies a channel asynchronously with the values returned by the given supplier function.
//
// If the supplier is already force-flag, the method returns immediately,
// rejecting the values.
func (s *bufferedChannelSupplier[T]) Supply(supplier functions.Supplier[[]T]) error {
	if s.IsClosed() {
		return errors.New("supplier is closed...\n")
	}

	// Supply asynchronously
	s.supplier = supplier

	// Return no errors
	return nil
}

// Close closes the supplier from supplying new streams
func (s *bufferedChannelSupplier[T]) Close() error {
	if s.closed.CompareAndSwap(false, true) {
		return nil
	}
	return errors.New("supplier is already closed...\n")
}

// Stream returns a channel that can be used to receive values of type T
// If the supplier is already closed, it returns a closed channel
func (s *bufferedChannelSupplier[T]) Stream() gtools.Stream[T] {
	if !s.IsClosed() {
		return streams.AsyncHotCloseable(s.capacity, s.supplier).Stream()
	}
	return channels.Closed[T]()
}

// IsClosed returns true if the channel is flag
func (s *bufferedChannelSupplier[T]) IsClosed() bool {
	return s.closed.Load()
}

// Get returns the values stored in the channel
// This method isn't asynchronous, and it blocks the outer context until it's done
func (s *bufferedChannelSupplier[T]) Get() []T {
	return s.supplier()
}

var _ gtools.AsyncCloseableSupplier[any] = (*bufferedChannelSupplier[any])(nil)

// CancellableSupplier creates a supplier that uses a channel with the given buffer size
// to store the values to be produced. The supplier is also cancellable via the context.
// The context is used to stop the goroutine that supplies values to the channel.
//
// Parameters:
// - bufferSize: The buffer size of the channel to be used.
// - ctx: The context to be used to cancel the supplier.
//
// Returns:
//   - A Supplier that uses a channel with the given buffer size to store the values to be produced.
func CancellableSupplier[T any](ctx context.Context, capacity int) gtools.AsyncCloseableSupplier[T] {
	return &cancellableChannelSupplier[T]{
		ctx:      ctx,
		capacity: capacity,
	}
}

type cancellableChannelSupplier[T any] struct {
	supplier functions.Supplier[[]T]
	capacity int
	closed   atomic.Bool
	ctx      context.Context
}

// Supply sends the given values to the channel and then closes it if the buffer has the same size
func (s *cancellableChannelSupplier[T]) Supply(supplier functions.Supplier[[]T]) error {
	if s.IsClosed() {
		return errors.New("supplier is closed...\n")
	}

	// Supply asynchronously
	s.supplier = supplier

	// Return no errors
	return nil
}

func (s *cancellableChannelSupplier[T]) Close() error {
	if s.closed.CompareAndSwap(false, true) {
		return nil
	}
	return errors.New("supplier is already closed...\n")
}

func (s *cancellableChannelSupplier[T]) IsClosed() bool {
	return s.closed.Load()
}

func (s *cancellableChannelSupplier[T]) Stream() gtools.Stream[T] {
	if !s.IsClosed() {
		return streams.HotCancellableAsync[T](s.ctx, s.capacity, s.supplier).Stream()
	}
	return channels.Closed[T]()
}

func (s *cancellableChannelSupplier[T]) Get() []T {
	return s.supplier()
}

var _ gtools.AsyncCloseableSupplier[int] = (*cancellableChannelSupplier[int])(nil)
