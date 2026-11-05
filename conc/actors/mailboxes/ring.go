// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"runtime"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/latches"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/semaph/limiters"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/waiters"
	"github.com/andrerrcosta2/gtools/core/data/closeable"
)

// Ring creates a bounded, blocking MPSC mailbox backed by a ring buffer.
//
// BEHAVIOR:
//   - Send: Producers reserve and publish slots using atomic operations.
//     When the mailbox is full they cooperatively wait for space.
//   - Receive: Blocks while the mailbox is empty until a message arrives
//     or the mailbox is closed.
//   - Order: Publication order. Messages sent concurrently by different
//     producers may be received in a different order than they reserved slots.
//   - Concurrency: Atomic producer coordination with condition-variable
//     blocking for waiting. A mutex is used exclusively by sync.Cond.
//   - Memory: Fixed-size ring buffer (capacity must be a power of two)
//
// USE WHEN:
//   - Message loss is unacceptable
//   - Producers should naturally slow down under load (backpressure)
//   - High throughput and low producer contention are more important than
//     preserving a global FIFO order across concurrent producers
//   - Examples: Event pipelines, actor mailboxes, work queues, stream processing
//
// AVOID WHEN:
//   - A strict FIFO order across all concurrent producers is required
//   - Message loss is acceptable (prefer a dropping mailbox)
//   - Only the latest value matters (prefer a coalescing mailbox)
//   - The buffer size cannot be a power of two
//
// THEORETICAL FOUNDATION:
//
//	Ring implements a bounded multi-producer, single-consumer mailbox with
//	backpressure. Producers reserve capacity without blocking each other and
//	publish messages independently, allowing maximum throughput under
//	contention. The consumer observes messages in publication order rather
//	than reservation order, avoiding the synchronization costs required to
//	enforce a global FIFO across concurrent producers.
//
//	The receive path uses a condition variable to avoid busy waiting while
//	the mailbox is empty. The associated mutex exists solely to satisfy the
//	synchronization requirements of sync.Cond and does not participate in the
//	producer fast path.
func Ring[E any](size int) mailbox.Closeable[E] {
	if size <= 0 || (size&(size-1)) != 0 {
		panic("size must be power of 2")
	}

	r := &ring[E]{
		buffer:  make([]E, size),
		mask:    uint64(size - 1),
		latch:   latches.Serial(),
		limiter: limiters.Intolerant(1, "ring mailbox is single-consumer; concurrent Receive calls are invalid"),
	}
	return r
}

type ring[E any] struct {
	buffer []E
	mask   uint64

	reserve uint64 // before head
	head    uint64 // atomic (producer)
	tail    uint64 // consumer only

	limiter syncs.Limiter
	latch   syncs.Latch

	state closeable.State
}

func (m *ring[E]) Send(e E) bool {
	if !closeable.IsOpenAtomic(&m.state) {
		return false
	}

	// a posição só controla a entrada
	// o head não é serializado, qualquer ticket que passa o backpressure
	// pode ser o próximo head
	pos := atomic.AddUint64(&m.reserve, 1) - 1

	if m.backpressure(pos) {
		next := atomic.AddUint64(&m.head, 1) - 1
		m.buffer[next&m.mask] = e // head: anyone whose position is within the buffer range
		atomic.AddUint64(&m.head, 1)
		m.latch.Signal()
		return true
	}
	return false
}

// backpressure applies a backpressure with a cooperative yield:
//
// 1. The current Goroutine (G) is removed from the executing thread (M).
// 2. It is placed at the end of the global run mailbox.
// 3. The scheduler picks the next Goroutine from the mailbox and starts running it.
//
// On low thread volumes using gosched may perform similar to busy waiting routines
// but in this case this rescheduling happens exactly when there is too many threads waiting
func (m *ring[E]) backpressure(pos uint64) bool {
	for {
		if pos-m.tail < uint64(len(m.buffer)) {
			return true
		}
		if closeable.IsClosedAtomic(&m.state) {
			return false
		}
		// Yields processor to other goroutines:
		// runtime.Gosched() is a cooperative yield - it tells the scheduler:
		// "I'm waiting, let someone else run.
		runtime.Gosched()
	}
}

func (m *ring[E]) Receive() (E, bool) {
	m.limiter.Acq()
	defer m.limiter.Rls()

	for m.head == m.tail {
		if closeable.IsClosedAtomic(&m.state) {
			var zero E
			return zero, false
		}
		m.latch.Wait()
	}

	pos := m.tail
	e := m.buffer[pos&m.mask]
	m.tail++

	return e, true
}

func (m *ring[E]) Close() error {
	if closeable.AtomicCloseIfIsNot(&m.state) {
		m.latch.Signal() // there is only a single routine calling latch.Wait()
		return nil
	}
	return mailbox.ErrClosedTwice
}

func (m *ring[E]) Shutdown() syncs.Waiter {
	if closeable.AtomicSealIfOpen(&m.state) {
		return waiters.YieldUntil(func() bool {
			if m.head == m.tail {
				m.state = closeable.Closed
				return true
			}
			return false
		})
	}
	return waiters.Complete()
}
