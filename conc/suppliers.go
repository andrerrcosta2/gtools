// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"context"
	"errors"
	"github.com/andrerrcosta2/gtools/conc/channels"
	"github.com/andrerrcosta2/gtools/conc/streams"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"sync"
	"sync/atomic"
)

// BufferedSupplier creates a supplier that uses a buffered channel
// of the given capacity to store the values to be produced.
func BufferedSupplier[T any](capacity int) *BufferedChannelSupplier[T] {
	return &BufferedChannelSupplier[T]{
		stack:    make([]T, 0),
		capacity: capacity,
		done:     make(chan struct{}),
	}
}

// BufferedChannelSupplier creates a supplier that uses a buffered channel
type BufferedChannelSupplier[T any] struct {
	mtx         sync.RWMutex
	stack       []T
	capacity    int
	closed      atomic.Bool
	forceClosed atomic.Bool
	done        chan struct{}
}

// Supply sends the given values to the channel and then closes it.
// The size of the slice is limited by the capacity of the channel.
// The method returns the supplier itself to allow method chaining.
func (s *BufferedChannelSupplier[T]) Supply(values ...T) error {
	if s.closed.Load() {
		return errors.New("supplier is closed...\n")
	}

	// Append values
	s.append(values...)

	// Return no errors
	return nil
}

func (s *BufferedChannelSupplier[T]) append(values ...T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.stack = append(s.stack, values...)
}

// Close closes the supplier.
//
// It's thread-safe and can be used concurrently.
func (s *BufferedChannelSupplier[T]) Close() error {
	if s.closed.CompareAndSwap(false, true) {
		s.stack = nil
		return nil
	}
	return errors.New("supplier is already closed...\n")
}

// Stream returns a channel that can be used to receive values of type T.
//
// It's thread-safe and can be used concurrently.
func (s *BufferedChannelSupplier[T]) Stream() gtools.Stream[T] {
	return streams.HotCloseable(s.capacity, s.stack...).Stream()
}

// IsClosed returns true if the supplier is closed, false otherwise.
//
// It's thread-safe and can be used concurrently.
func (s *BufferedChannelSupplier[T]) IsClosed() bool {
	return s.closed.Load()
}

// Get returns all the values that have been sent to the supplier.
//
// It's thread-safe and can be used concurrently.
func (s *BufferedChannelSupplier[T]) Get() []T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.stack
}

// Clear clears the internal stack of the supplier.
//
// It's thread-safe and can be used concurrently.
func (s *BufferedChannelSupplier[T]) Clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.stack = make([]T, 0)
}

var _ gtools.CloseableSupplier[int] = (*BufferedChannelSupplier[int])(nil)

// CancellableSupplier creates a supplier that uses a channel with the given buffer size
// to store the values to be produced. The supplier is also cancellable via the context.
// The context is used to stop the goroutine that supplies values to the channel.
//
// Parameters:
// - ctx: The context to be used to cancel the supplier.
//
// Returns:
//   - A Supplier that uses a channel with the given buffer size to store the values to be produced.
func CancellableSupplier[T any](ctx context.Context, capacity int) *CancellableChannelSupplier[T] {
	return &CancellableChannelSupplier[T]{ctx: ctx, stack: make([]T, capacity), capacity: capacity}
}

type CancellableChannelSupplier[T any] struct {
	mtx      sync.RWMutex
	stack    []T
	ctx      context.Context
	capacity int
	closed   atomic.Bool
}

func (s *CancellableChannelSupplier[T]) Supply(values ...T) error {
	if s.closed.Load() {
		return errors.New("supplier is closed...\n")
	}

	// Context cancelled
	if s.ctx.Err() != nil {
		return s.ctx.Err()
	}

	// Append values
	go s.append(values...)

	// Return no errors
	return nil
}

func (s *CancellableChannelSupplier[T]) append(values ...T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.stack = append(s.stack, values...)
}

// Close closes the supplier and stops it from accepting new values.
// It's thread-safe and can be used concurrently.
// If the supplier is already closed, the method returns an error.
func (s *CancellableChannelSupplier[T]) Close() error {
	if s.closed.CompareAndSwap(false, true) {
		s.stack = nil
		return nil
	}
	return errors.New("supplier is already closed...\n")
}

// IsClosed returns true if the supplier is closed, false otherwise.
// It's thread-safe and can be used concurrently.
func (s *CancellableChannelSupplier[T]) IsClosed() bool {
	return s.closed.Load()
}

// Stream returns a channel that can be used to receive values of type T.
// It's thread-safe and can be used concurrently.
func (s *CancellableChannelSupplier[T]) Stream() gtools.Stream[T] {
	if s.IsClosed() {
		return channels.Closed[T]()
	}
	return streams.HotCancellable(s.ctx, s.capacity, s.stack...).Stream()
}

func (s *CancellableChannelSupplier[T]) Get() []T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.stack
}

func (s *CancellableChannelSupplier[T]) Clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.stack = make([]T, 0)
}

var _ gtools.CloseableSupplier[int] = (*CancellableChannelSupplier[int])(nil)
