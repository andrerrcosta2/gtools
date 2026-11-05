// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/actor/mailbox"
	"github.com/andrerrcosta2/gtools/conc/core/sealable"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/latches"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/waiters"
	"github.com/andrerrcosta2/gtools/core/data/closeable"
	"github.com/andrerrcosta2/gtools/core/data/str/buffer"
)

// Buffered creates a simple mailbox that protects a buffer with a single mutex
func Buffered[E any](storage buffer.Mutable[E]) mailbox.Buffered[E] {
	return &buffered[E]{
		storage: storage,
		done:    latches.Once(),
	}
}

type buffered[E any] struct {
	mtx     sync.Mutex
	storage buffer.Mutable[E]
	done    syncs.Latch
	state   closeable.State
}

func (m *buffered[E]) Send(e E) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if closeable.IsOpen(m.state) {
		return m.storage.Write(e)
	}
	return false
}

func (m *buffered[E]) Receive() (E, bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	switch m.state {
	case sealable.Open:
		return m.storage.Read()

	case sealable.Sealed:
		data, ok := m.storage.Read()
		if !ok || m.storage.Len() == 0 {
			m.close()
		}
		return data, true

	default:
		{
		}
	}
	var zero E
	return zero, false
}

func (m *buffered[E]) Shutdown() syncs.Waiter {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if sealable.SealIfOpen(&m.state) {
		if m.storage.Len() == 0 {
			m.close()
			return waiters.Complete()
		}
	}
	return m.done
}

func (m *buffered[E]) Close() error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if closeable.CloseIfIsNot(&m.state) {
		m.close()
		return nil
	}
	return mailbox.ErrClosedTwice
}

func (m *buffered[E]) close() {
	m.state = closeable.Closed
	m.done.Signal()
}

func (m *buffered[E]) Cap() int {
	// Len doesn't change
	return m.storage.Cap()
}

func (m *buffered[E]) Len() int {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	return m.storage.Len()
}
