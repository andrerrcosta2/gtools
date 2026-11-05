// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package agendas

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/queue"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/datastr/queues"
)

// newBucketWheel is a circular bucketed priority queue inspired by timing wheels
// used in systems such as Kafka, Netty, Akka and the Linux kernel.
//
// Instead of maintaining a single globally ordered priority queue,
// elements are first grouped into buckets according to their key,
// and only elements within the same bucket are ordered.
//
//				[0]
//
//	[511]				[1]
//
// [510]					[2]
//
// ...
//
//	[256]
//
// Each bucket represents a contiguous interval of the key space,
// called the bucket span.
//
// Example (time domain):
//
//	bucketSpan = 10ms
//	buckets    = 512
//
// Total covered interval:
//
//	512 × 10ms = 5.12s
//
// All values whose keys fall within the same 10ms interval are placed
// into the same bucket and ordered locally.
//
// Smaller bucket spans reduce collisions at the cost of advancing the
// wheel more frequently. Larger spans reduce advancement frequency but
// increase the number of elements sharing each bucket.
//
// The wheel itself is domain-agnostic. It only requires an ordered key
// for each element
func newBucketWheel[T any, K nums.Ordered](cmp comparator.CmpKeyer[T, K], origin, bucketSpan K, buckets int) *bucketWheel[T, K] {
	if buckets <= 0 || (buckets&(buckets-1)) != 0 {
		panic("buckets must be a positive power of two")
	}

	wheel := make([]queue.MutablePriority[bucketNode[T]], buckets)
	for i := range wheel {
		wheel[i] = queues.PriorityOf[bucketNode[T]]()
	}

	return &bucketWheel[T, K]{
		cmp:     cmp,
		wheel:   wheel,
		origin:  origin,
		span:    bucketSpan,
		buckets: buckets,
	}
}

type bucketNode[T any] struct {
	value   T
	lap     int
	compare func(T, T) int
}

func (b bucketNode[T]) Less(than bucketNode[T]) bool {
	return b.compare(b.value, than.value) < 0
}

type bucketWheel[T any, K nums.Ordered] struct {
	cmp     comparator.CmpKeyer[T, K]
	wheel   []queue.MutablePriority[bucketNode[T]]
	origin  K
	span    K // span of ordered keys for each bucket
	buckets int
	len     int // elements inside the wheel
	curr    int // current bucket
	lap     int // current lap around the wheel
}

func (q *bucketWheel[T, K]) Peek() (T, bool) {
	slot := q.wheel[q.curr]
	node, ok := slot.Peek()
	return node.value, ok
}

func (q *bucketWheel[T, K]) PeekBucket() ([]T, bool) {
	slot := q.wheel[q.curr]
	out := make([]T, 0, slot.Len())
	cut := 0
	for cut < slot.Len() {
		next, ok := slot.Peek()
		if !ok || next.lap != q.lap {
			break
		}
		out = append(out, next.value)
		cut++
	}
	return out, cut > 0
}

func (q *bucketWheel[T, K]) Len() int {
	return q.len
}

func (q *bucketWheel[T, K]) Push(t T) {
	key := q.cmp.Key(t) - q.origin

	spans := int(key / q.span)
	slot := q.wheel[(q.curr+spans)%q.buckets]

	slot.Push(bucketNode[T]{
		value:   t,
		lap:     spans / q.buckets,
		compare: q.cmp.Compare,
	})
	q.len++
}

func (q *bucketWheel[T, K]) Pop() (T, bool) {
	slot := q.wheel[q.curr]
	node, ok := slot.Pop()
	if !ok {
		return node.value, false
	}
	q.curr--
	return node.value, true
}

func (q *bucketWheel[T, K]) PopBucket() ([]T, bool) {
	slot := q.wheel[q.curr]
	out := make([]T, 0, slot.Len())
	cut := 0
	for cut < slot.Len() {
		next, ok := q.popIfDue()
		if !ok {
			break
		}
		out = append(out, next)
		cut++
	}
	q.advanceCurr()
	return out, cut > 0
}

func (q *bucketWheel[T, K]) advanceCurr() {
	q.curr = (q.curr + 1) % q.buckets
}

// PopRange invokes fn for each due element in the current bucket, in order.
// It returns false if fn requested early termination — in that case the
// bucket is left partially drained and curr is not advanced, so remaining
// due elements are retried on the next call.
func (q *bucketWheel[T, K]) PopRange(fn func(T) bool) bool {
	for {
		v, ok := q.popIfDue()
		if !ok {
			q.advanceCurr()
			return true
		}
		if !fn(v) {
			return false
		}
	}
}

// PeekRange invokes fn for each due element in the current bucket, in order.
// It returns false if fn requested early termination — in that case the
// bucket is left partially drained and curr is not advanced, so remaining
// due elements are retried on the next call.
func (q *bucketWheel[T, K]) PeekRange(fn func(T) bool) bool {
	for {
		v, ok := q.peekIfDue()
		if !ok {
			q.advanceCurr()
			return true
		}
		if !fn(v) {
			return false
		}
	}
}

func (q *bucketWheel[T, K]) popIfDue() (T, bool) {
	slot := q.wheel[q.curr]
	next, ok := slot.Peek()
	// only works because the priority is ordered
	if !ok || next.lap != q.lap {
		var zero T
		return zero, false
	}
	_, _ = slot.Pop()
	q.len--
	return next.value, true
}

func (q *bucketWheel[T, K]) peekIfDue() (T, bool) {
	slot := q.wheel[q.curr]
	// only works because the priority is ordered
	next, ok := slot.Peek()
	if !ok || next.lap != q.lap {
		var zero T
		return zero, false
	}
	return next.value, true
}

func (q *bucketWheel[T, K]) Clear() {
	q.wheel = make([]queue.MutablePriority[bucketNode[T]], 0, q.buckets)
	q.len = 0
}
