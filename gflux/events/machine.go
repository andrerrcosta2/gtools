// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package events

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
)

type Handler[E any] func(e E)

// Machine is a single-threaded states processor with serialized event handling that guarantees:
//
//  1. Only one handler execution at a time.
//  2. State is only mutated by that handler.
//  3. Send is safe concurrently.
//  4. Stop is idempotent.
//  5. After Stop, no new commands are accepted.
type Machine[E any] interface {
	Send(E) bool
	Stop()
}

func HardMachine[E any](scheduler Schedulable, mailbox mailbox.Closeable[E], handler Handler[E]) Machine[E] {
	m := &machine[E]{
		mb:    mailbox,
		stop:  make(chan struct{}),
		drain: true,
	}

	scheduler.Schedule(func() {
		m.run(handler)
	})

	return m
}

func WeakMachine[E any](scheduler Schedulable, mailbox mailbox.Closeable[E], handler Handler[E]) Machine[E] {
	m := &machine[E]{
		mb:   mailbox,
		stop: make(chan struct{}),
	}

	scheduler.Schedule(func() {
		m.run(handler)
	})

	return m
}

type machine[E any] struct {
	mb    mailbox.Closeable[E]
	stop  chan struct{}
	once  sync.Once
	drain bool
}

func (m *machine[E]) Send(e E) bool {
	return m.mb.Send(e)
}

func (m *machine[E]) Stop() {
	m.once.Do(func() {
		close(m.stop)
		m.mb.Close()
	})
}

func (m *machine[E]) run(handler Handler[E]) {
	for {
		e, ok := m.mb.Receive()
		if !ok {
			return
		}

		handler(e)

		select {
		case <-m.stop:
			if !m.drain {
				return
			}
		default:
		}
	}
}
