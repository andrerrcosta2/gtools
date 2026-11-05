// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package latches

import (
	"sync"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/signalers"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/waiters"
)

// Unchecked returns a latch that panics on a second signal
func Unchecked() syncs.Latch {
	return &unchecked{
		ch: make(chan struct{}),
	}
}

type unchecked struct {
	ch chan struct{}
}

func (s *unchecked) Signal() {
	close(s.ch)
}

func (s *unchecked) Wait() {
	<-s.ch
}

// UncheckedPair returns a latch/waiter pair that panics on a second signal
func UncheckedPair() syncs.PairLatch {
	ch := make(chan struct{})
	return &uncheckedPair{
		ch:     ch,
		waiter: waiters.Chan(ch),
	}
}

type uncheckedPair struct {
	ch     chan struct{}
	waiter syncs.Waiter
}

func (s *uncheckedPair) Signal() {
	close(s.ch)
}

func (s *uncheckedPair) Wait() {
	<-s.ch
}

func (s *uncheckedPair) Waiter() syncs.Waiter {
	return waiters.Chan(s.ch)
}

func (s *uncheckedPair) Signaler() syncs.Signaler {
	return signalers.Chan(s.ch)
}

// Once is a latch that ensures it signals only once and never blocks again on wait
func Once() syncs.Latch {
	return &once{
		ch: make(chan struct{}),
	}
}

type once struct {
	ch   chan struct{}
	once sync.Once
}

func (l *once) Signal() {
	l.once.Do(func() {
		close(l.ch)
	})
}

func (l *once) Wait() {
	<-l.ch
}

// OncePair is a latch that ensures it signals only once and never blocks again on wait
func OncePair() syncs.PairLatch {
	ch := make(chan struct{})
	return &oncePair{
		ch:     ch,
		waiter: waiters.Chan(ch),
	}
}

type oncePair struct {
	ch     chan struct{}
	waiter syncs.Waiter
	once   sync.Once
}

func (l *oncePair) Signal() {
	l.once.Do(func() {
		close(l.ch)
	})
}

func (l *oncePair) Wait() {
	<-l.ch
}

func (l *oncePair) Signaler() syncs.Signaler {
	return signalers.Func(l.Signal)
}

func (l *oncePair) Waiter() syncs.Waiter {
	return waiters.Chan(l.ch)
}

func Countdown(n int64) syncs.Latch {
	return &countdown{
		n:  n,
		ch: make(chan struct{}),
	}
}

type countdown struct {
	n  int64
	ch chan struct{}
}

func (l *countdown) Signal() {
	if atomic.AddInt64(&l.n, -1) == 0 {
		close(l.ch)
	}
}

func (l *countdown) Wait() {
	<-l.ch
}

// Serial is a reusable latch where each Signal releases at most one waiting goroutine.
// After a successful Wait, the latch automatically returns to the closed state.
func Serial() syncs.Latch {
	return &serial{
		ch: make(chan struct{}, 1),
	}
}

type serial struct {
	ch chan struct{}
}

func (l *serial) Signal() {
	select {
	case l.ch <- struct{}{}:
	default:
	}
}

func (l *serial) Wait() {
	<-l.ch
}
