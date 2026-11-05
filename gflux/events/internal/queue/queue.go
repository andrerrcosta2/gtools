// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package queue

import "sync"

type backpressured[T any] struct {
	ring     ring[T]
	mu       sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
}

func (q *backpressured[T]) Push(v T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for q.ring.size == len(q.ring.buf) {
		q.notFull.Wait()
	}

	q.ring.Push(v)
	q.notEmpty.Signal()
}

func (q *backpressured[T]) Pop() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	for q.ring.size == 0 {
		q.notEmpty.Wait()
	}

	v, _ := q.ring.Pop()
	q.notFull.Signal()
	return v
}

type ring[T any] struct {
	buf  []T
	head int // next read
	tail int // next write
	size int
}

func (r *ring[T]) Push(v T) bool {
	if r.size == len(r.buf) {
		return false // full
	}

	r.buf[r.tail] = v
	r.tail = (r.tail + 1) % len(r.buf)
	r.size++
	return true
}

func (r *ring[T]) Pop() (T, bool) {
	var zero T

	if r.size == 0 {
		return zero, false
	}

	v := r.buf[r.head]
	r.head = (r.head + 1) % len(r.buf)
	r.size--

	return v, true
}
