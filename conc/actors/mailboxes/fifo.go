// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/latches"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/waiters"
	"github.com/andrerrcosta2/gtools/core/data/closeable"
)

// FIFO creates a lossy, lock-free mailbox that drops messages when the buffer is full.
//
// BEHAVIOR:
//   - Send: Non-blocking, returns false immediately if buffer is full
//   - Receive: Blocks until a message is available or mailbox is closed
//   - Order: FIFO (First-In-First-Out) for messages that aren't dropped
//   - Concurrency: Lock-free (uses buffered channel internally)
//   - Memory: Fixed buffer size, no growth
//
// USE WHEN:
//   - Message loss is acceptable under backpressure
//   - You need maximum throughput with minimal latency
//   - The operator has "cancellable" semantics where only the latest message matters
//   - Examples: Debounce, Throttle, Sample, Latest operators
//
// AVOID WHEN:
//   - Every message must be processed (no loss allowed)
//   - You need to block producers to slow down the system
//   - Messages have non-idempotent effects
//
// THEORETICAL FOUNDATION:
//
//	FIFO mailbox implements a "lossy mailbox" model. It's ideal when the system
//	can tolerate data loss because the semantics of the operator make older
//	messages obsolete (like debounce where only the latest value matters).
func FIFO[E any](bufferSize int) mailbox.Closeable[E] {
	return &fifo[E]{
		buffered: make(chan E, bufferSize),
		latch:    latches.Once(),
	}
}

type fifo[E any] struct {
	buffered chan E
	state    closeable.State
	latch    syncs.Latch
}

func (m *fifo[E]) Send(e E) bool {
	if !closeable.IsOpenAtomic(&m.state) {
		return false
	}

	select {
	case m.buffered <- e:
		return true
	default:
		return false
	}
}

func (m *fifo[E]) Receive() (E, bool) {
	e, ok := <-m.buffered
	if closeable.IsSealedAtomic(&m.state) && len(m.buffered) == 0 {
		m.state = closeable.Closed
		m.latch.Signal()
	}
	return e, ok
}

func (m *fifo[E]) Shutdown() syncs.Waiter {
	if closeable.AtomicSealIfOpen(&m.state) {
		if len(m.buffered) == 0 {
			m.state = closeable.Closed
			return waiters.Complete()
		}
		return m.latch
	}
	return waiters.Complete()
}

func (m *fifo[E]) Close() error {
	if !closeable.AtomicCloseIfIsNot(&m.state) {
		close(m.buffered)
		m.state = closeable.Closed
		return nil
	}
	return mailbox.ErrClosedTwice
}
