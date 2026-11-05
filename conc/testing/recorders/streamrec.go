// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package recorders

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/testlite"
)

// Stream is a test helper that records all values sent to it and allows
// controlled iteration over them. It implements stream.Readable and can
// be used as a drop-in replacement for real streams in tests.
type Stream[T any] struct {
	mu      sync.RWMutex
	values  []T
	closed  atomic.Bool
	nextIdx int
	t       testlite.HelperTesting
}

// StreamOf creates a new Stream with the given initial values.
// The recorder starts in an open state.
func StreamOf[T any](t testlite.LoggerTesting, values ...T) *Stream[T] {
	return &Stream[T]{
		values:  values,
		closed:  atomic.Bool{},
		nextIdx: 0,
		t:       t,
	}
}

// Record appends one or more values to the recorder.
// Returns error if the recorder is closed.
func (r *Stream[T]) Record(values ...T) error {
	if r.closed.Load() {
		return stream.ErrClosed
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.values = append(r.values, values...)
	return nil
}

// Close closes the recorder, preventing further records.
// It's thread-safe and can be used concurrently.
func (r *Stream[T]) Close() error {
	if r.closed.CompareAndSwap(false, true) {
		return nil
	}
	return stream.ErrClosed
}

// IsClosed returns true if the recorder is closed, false otherwise.
func (r *Stream[T]) IsClosed() bool {
	return r.closed.Load()
}

// Iterator returns an iterator over the recorded values.
// The iterator will iterate over all values currently recorded.
func (r *Stream[T]) Iterator() stream.Iterator[T] {
	return func(fn func(e T) bool) bool {
		r.mu.RLock()
		values := make([]T, len(r.values))
		copy(values, r.values)
		r.mu.RUnlock()

		for _, v := range values {
			if !fn(v) {
				return false
			}
		}
		return true
	}
}

// Reset clears all recorded values and resets the iteration state.
// The recorder remains open unless explicitly closed.
func (r *Stream[T]) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.values = nil
	r.nextIdx = 0
}

// Values returns a copy of all recorded values.
func (r *Stream[T]) Values() []T {
	r.mu.RLock()
	defer r.mu.RUnlock()
	values := make([]T, len(r.values))
	copy(values, r.values)
	return values
}

// Len returns the number of recorded values.
func (r *Stream[T]) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.values)
}

// RecordAsync is a helper that returns a function which records values.
// Useful for passing as a supplier function in tests.
func RecordAsync[T any](r *Stream[T]) function.Supplier[[]T] {
	return func() []T {
		// Return some test data or record from external source
		return r.Values()
	}
}

// var _ stream.Readable[int] = (*Stream[int])(nil)

// ColdRecorder is a test helper that mimics the behavior of ColdCloseable
// and ColdCancellable streams but with full control over the data.
type ColdRecorder[T any] struct {
	*Stream[T]
	ctx context.Context
}

// NewColdRecorder creates a new ColdRecorder with the given initial values.
func NewColdRecorder[T any](t testlite.LoggerTesting, values ...T) *ColdRecorder[T] {
	return &ColdRecorder[T]{
		Stream: StreamOf(t, values...),
	}
}

// NewColdRecorderWithContext creates a new ColdRecorder that respects context cancellation.
func NewColdRecorderWithContext[T any](t testlite.LoggerTesting, ctx context.Context, values ...T) *ColdRecorder[T] {
	return &ColdRecorder[T]{
		Stream: StreamOf(t, values...),
		ctx:    ctx,
	}
}

// Iterator returns an iterator that respects context cancellation.
func (c *ColdRecorder[T]) Iterator() stream.Iterator[T] {
	if c.ctx != nil && c.ctx.Err() != nil {
		return func(fn func(e T) bool) bool { return false }
	}

	return func(fn func(e T) bool) bool {
		if c.ctx != nil {
			select {
			case <-c.ctx.Done():
				return false
			default:
			}
		}

		c.mu.RLock()
		values := make([]T, len(c.values))
		copy(values, c.values)
		c.mu.RUnlock()

		for _, v := range values {
			if c.ctx != nil {
				select {
				case <-c.ctx.Done():
					return false
				default:
				}
			}

			if !fn(v) {
				return false
			}
		}
		return true
	}
}

// RecordAsyncForContext records values asynchronously with context support.
func (c *ColdRecorder[T]) RecordAsyncForContext(supplier function.Supplier[[]T]) {
	go func() {
		select {
		case <-c.ctx.Done():
			return
		default:
			values := supplier()
			err := c.Record(values...)
			if err != nil {
				c.t.Error(err.Error())
			}
		}
	}()
}
