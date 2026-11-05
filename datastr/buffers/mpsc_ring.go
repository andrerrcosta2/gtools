// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package buffers

import (
	"runtime"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/core/data/str/buffer"
)

// Disruptor creates a MPSC (Multiple Producer Single Consumer) Ring Buffer
// lock-free optimized for concurrency using the LMAX Disruptor Pattern.
//
// Its size must be power of two.
func Disruptor[T any](size uint64) buffer.MPSC[T] {
	if (size & (size - 1)) != 0 {
		panic("disruptor size must be a positive power of two")
	}
	return newLmaxRing[T](size)
}

func newLmaxRing[T any](size uint64) *lmaxRing[T] {
	ring := &lmaxRing[T]{
		buffer: make([]T, size),
		seq:    make([]uint64, size),
		mask:   size - 1,
	}
	for i := range ring.seq {
		ring.seq[i] = uint64(i)
	}
	return ring
}

type cachePad struct {
	_ [64]byte
}

// open CPU caches load data in cache lines of 64 bytes. without padding:
//  1. writeSeq and readSeq might share a cache line
//  2. qhen consumer updates readSeq, it invalidates producer's cache
//  3. this causes false sharing: cache lines bouncing between CPUs
//     https://en.wikipedia.org/wiki/False_sharing
type lmaxRing[T any] struct {
	_        cachePad // isolate writeSeq in its own cache line
	writeSeq uint64   // Written ONLY by producers

	_       cachePad // isolate readSeq in its own cache line
	readSeq uint64   // Written ONLY by consumer

	_      cachePad
	buffer []T
	// Range slot has its own sequence number to avoid backpressure contention
	// on scheduling (that doesn't mean it hasn't backpressure, but it avoids
	// backpressure race on reservation)
	// this is a core disruptor trick that allow us to know the slot states
	// (ready, empty, overwritten ...)
	// Range slot in seq[] goes through states:
	// Initial state: seq[i] = i: Empty, available for write
	// After write:   seq[i] = i+1: Contains data, available for read
	// After read:    seq[i] = i+size (far ahead)
	seq  []uint64
	mask uint64
}

func (b *lmaxRing[T]) Cap() int {
	return len(b.buffer)
}
func (b *lmaxRing[T]) Len() int {
	return int(b.writeSeq) - int(b.readSeq)
}

func (b *lmaxRing[T]) Read() (value T, exists bool) {
	// Get current read position
	seq, index := b.nextReadSeq()
	v := b.buffer[index]

	// Mark as empty (far ahead)
	atomic.StoreUint64(&b.seq[index], seq+b.mask+1)

	// there is no concurrency on the listener side
	b.readSeq++
	return v, true
}

func (b *lmaxRing[T]) nextReadSeq() (seq, index uint64) {
	seq = b.readSeq
	index = seq & b.mask

	// Wait for data to be published
	// Data ready when seq[index] == seq+1
	for atomic.LoadUint64(&b.seq[index]) != seq+1 {
		runtime.Gosched()
	}

	return
}

// Write never returns false. this mpsc doesn't work without a
// backpressure (reserve) policy
// Key rule: write → publish (never publish before writing)
func (b *lmaxRing[T]) Write(e T) bool {
	//  Reserve a sequence number (atomic)
	seq := b.reserve()

	// Wait if buffer is full (backpressure)
	b.backpressure(seq)

	// Calculate slot index
	index := seq & b.mask

	// Wait for slot to be "available for write"
	// that happens by a buffer-relative sequence (seq[index] = ~index
	for atomic.LoadUint64(&b.seq[index]) != seq {
		runtime.Gosched()
	}

	// Now it knows it can be overwritten (it was already consumed
	b.buffer[index] = e

	// Publish: mark as readable
	// Consumer waits for seq[index] == seq+1
	atomic.StoreUint64(&b.seq[index], seq+1)
	return true
}

func (b *lmaxRing[T]) reserve() uint64 {
	return atomic.AddUint64(&b.writeSeq, 1) - 1
}

func (b *lmaxRing[T]) backpressure(reserve uint64) {
	for reserve-b.readSeq >= b.mask {
		runtime.Gosched() // yield
	}
}

func (b *lmaxRing[T]) Empty() bool { return b.readSeq == atomic.LoadUint64(&b.writeSeq) }

func (b *lmaxRing[T]) Full() bool {
	return atomic.LoadUint64(&b.writeSeq)-b.readSeq > b.mask
}
