// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package recorders

import "sync"

// Buffer is a test-only buffer.Mutable implementation that records
// operations and allows precise control over Read/Write behavior.
type Buffer[T any] struct {
	mu sync.Mutex

	// Configurable behavior
	writeFn func(T) bool
	readFn  func() (T, bool)
	lenFn   func() int
	capFn   func() int

	// Operation recording
	Writes     []T
	Reads      []T
	WriteCalls int
	ReadCalls  int

	// Internal storage for simple scenarios
	items []T
}

// NewBuffer creates a Buffer with simple FIFO queue behavior.
func NewBuffer[T any](capacity int) *Buffer[T] {
	r := &Buffer[T]{
		items: make([]T, 0, capacity),
	}

	r.writeFn = func(v T) bool {
		if len(r.items) < capacity {
			r.items = append(r.items, v)
			return true
		}
		return false
	}

	r.readFn = func() (T, bool) {
		if len(r.items) == 0 {
			var zero T
			return zero, false
		}
		v := r.items[0]
		r.items = r.items[1:]
		return v, true
	}

	r.lenFn = func() int {
		return len(r.items)
	}

	r.capFn = func() int {
		return capacity
	}

	return r
}

// Write implements buffer.Writable.
func (r *Buffer[T]) Write(v T) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.WriteCalls++
	r.Writes = append(r.Writes, v)
	if r.writeFn != nil {
		return r.writeFn(v)
	}
	return false
}

// Read implements buffer.Readable.
func (r *Buffer[T]) Read() (T, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ReadCalls++
	if r.readFn != nil {
		v, ok := r.readFn()
		if ok {
			r.Reads = append(r.Reads, v)
		}
		return v, ok
	}
	var zero T
	return zero, false
}

// Len implements buffer.Typed.
func (r *Buffer[T]) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.lenFn != nil {
		return r.lenFn()
	}
	return 0
}

// Cap implements buffer.Typed.
func (r *Buffer[T]) Cap() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.capFn != nil {
		return r.capFn()
	}
	return 0
}

// SetWriteFn overrides the Write behavior.
func (r *Buffer[T]) SetWriteFn(fn func(T) bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.writeFn = fn
}

// SetReadFn overrides the Read behavior.
func (r *Buffer[T]) SetReadFn(fn func() (T, bool)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.readFn = fn
}

// SetLenFn overrides the Len behavior.
func (r *Buffer[T]) SetLenFn(fn func() int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.lenFn = fn
}

// Reset clears all recorded operations while preserving behavior fns.
func (r *Buffer[T]) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Writes = nil
	r.WriteCalls = 0
	r.ReadCalls = 0
	r.items = nil
}
