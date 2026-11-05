// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package buffers

import "github.com/andrerrcosta2/gtools/core/data/str/buffer"

// OverwritableRing creates a fixed-size ring buffer that overwrites the
// oldest elements when full.
//
// The capacity must be a positive power of two.
//
// Restricting the capacity to powers of two allows index wrapping to be
// implemented with a bit mask instead of a modulo operation, simplifying
// the implementation and matching the requirements of other low-level
// ring buffer implementations.
func OverwritableRing[T any](size int) buffer.Mutable[T] {
	if size <= 0 || (size&(size-1)) != 0 {
		panic("ring size must be a positive power of two")
	}
	return &overwrRing[T]{
		buffer: make([]T, size),
		mask:   uint64(size - 1),
	}
}

// overwrRing
type overwrRing[E any] struct {
	buffer []E

	mask uint64

	head uint64 // next write
	tail uint64 // oldest element
}

func (r *overwrRing[E]) Write(e E) bool {
	if r.isFull() {
		r.tail++
	}

	r.buffer[r.head&r.mask] = e
	r.head++
	return true
}

func (r *overwrRing[E]) isFull() bool {
	return r.head-r.tail == uint64(len(r.buffer))
}

func (r *overwrRing[E]) Read() (E, bool) {
	if r.head == r.tail {
		var zero E
		return zero, false
	}

	e := r.buffer[r.tail&r.mask]
	r.tail++

	return e, true
}

func (r *overwrRing[E]) Len() int {
	return int(r.head - r.tail)
}

func (r *overwrRing[E]) Cap() int {
	return len(r.buffer)
}
