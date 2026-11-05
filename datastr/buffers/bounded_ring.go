// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package buffers

import "github.com/andrerrcosta2/gtools/core/data/str/buffer"

// BoundedRing returns a buffer.Mutable ring that stops allowing
// new values when it is full
func BoundedRing[T any](size int) buffer.Mutable[T] {
	return &boundedRing[T]{
		buffer: make([]T, size),
		mask:   uint64(size - 1),
	}
}

// boundedRing
type boundedRing[E any] struct {
	buffer []E
	mask   uint64

	head uint64 // next write
	tail uint64 // oldest element
}

func (r *boundedRing[E]) Write(e E) bool {
	if r.isFull() {
		return false
	}

	r.buffer[r.head&r.mask] = e
	r.head++
	return true
}

func (r *boundedRing[E]) isFull() bool {
	return r.head-r.tail == uint64(len(r.buffer))
}

func (r *boundedRing[E]) Read() (E, bool) {
	if r.head == r.tail {
		var zero E
		return zero, false
	}

	e := r.buffer[r.tail&r.mask]
	r.tail++

	return e, true
}

func (r *boundedRing[E]) Len() int {
	return int(r.head - r.tail)
}

func (r *boundedRing[E]) Cap() int {
	return len(r.buffer)
}
