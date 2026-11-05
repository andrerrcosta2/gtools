// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/waiters"
	"github.com/andrerrcosta2/gtools/core/data/closeable"
	"github.com/andrerrcosta2/gtools/core/data/str/buffer"
	"github.com/andrerrcosta2/gtools/datastr/buffers"
)

// Coalescing creates a mailbox that only keeps the latest message,
// dropping all previous unprocessed messages.
//
// BEHAVIOR:
//   - Send: If no message is pending, the message becomes pending and
//     increments in-flight count. If a message is already pending, it is
//     replaced (dropped) without changing in-flight count.
//   - Receive: Returns the latest pending message and decrements in-flight count
//   - Send blocks when maxInFlight limit is reached AND no message is pending
//   - Concurrency: Mutex-protected with condition variables
//   - Memory: O(1) - only stores one message at a time
//
// USE WHEN:
//   - Only the most recent value matters
//   - You need to prevent buildup of stale data
//   - Processing is expensive and you want to skip intermediate values
//   - Examples: Latest value cache, UI updates, sensor readings
//
// AVOID WHEN:
//   - Every value must be processed (use Ring)
//   - You need to preserve order of multiple pending items (use FIFO or Ring)
//   - Messages have cumulative effects (each one changes state)
//
// THEORETICAL FOUNDATION:
//
//	Coalescing implements a "last one wins" model. It's ideal for operators
//	like SwitchMap, Latest, or any scenario where intermediate values become
//	obsolete once a newer one arrives. The maxInFlight parameter controls
//	how many "generations" of values can be pending, effectively limiting
//	the staleness of data.
func Coalescing[E any]() mailbox.Closeable[E] {
	m := &coalescing[E]{
		storage: buffers.OverwritableRing[E](1),
	}
	return m
}

type coalescing[E any] struct {
	storage buffer.Mutable[E]
	state   closeable.State
}

func (m *coalescing[E]) Send(e E) bool {
	if !closeable.IsOpen(m.state) {
		return false
	}

	// If already has a value just replace (does not increment)
	m.storage.Write(e)
	return true
}

func (m *coalescing[E]) isEmpty() bool {
	return m.storage.Len() == 0
}

func (m *coalescing[E]) Receive() (E, bool) {
	ev, _ := m.storage.Read()
	return ev, true
}

func (m *coalescing[E]) Close() error {
	if !closeable.IsClosed(m.state) {
		m.state = closeable.Closed
		return nil
	}
	return mailbox.ErrClosedTwice
}

func (m *coalescing[E]) Shutdown() syncs.Waiter {
	if closeable.IsOpen(m.state) {
		m.state = closeable.Sealed
		return waiters.YieldUntil(func() bool {
			if m.isEmpty() {
				m.state = closeable.Closed
				return true
			}
			return false
		})
	}
	return waiters.Complete()
}
