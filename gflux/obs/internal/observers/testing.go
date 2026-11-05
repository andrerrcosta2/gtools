// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observers

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/gflux/core/lifecycle/gates"
)

type Recorder[T any] struct {
	mu         sync.Mutex
	values     []T
	err        error
	completed  atomic.Bool
	terminated atomic.Bool
	t          *testing.T

	nextCh     chan struct{}
	errorCh    chan struct{}
	completeCh chan struct{}
}

func (o *Recorder[T]) Values() []T {
	o.mu.Lock()
	defer o.mu.Unlock()
	return append([]T(nil), o.values...)
}

func (o *Recorder[T]) Err() error {
	return o.err
}

func (o *Recorder[T]) Completed() bool {
	return o.completed.Load()
}

func (o *Recorder[T]) Terminated() bool {
	return o.terminated.Load()
}

func (o *Recorder[T]) AwaitTermination(timeout time.Duration) {
	select {
	case <-o.errorCh:
		return
	case <-o.completeCh:
		return
	case <-time.After(timeout):
		o.t.Helper()
		o.t.Fatal("❌ Recorder: timeout waiting for termination")
	}
}

func (o *Recorder[T]) AwaitNext(timeout time.Duration) {
	select {
	case <-o.nextCh:
		return
	case <-time.After(timeout):
		o.t.Helper()
		o.t.Fatal("❌ Recorder: timeout waiting for Next")
	}
}

func (o *Recorder[T]) AwaitError(timeout time.Duration) error {
	select {
	case <-o.errorCh:
		return o.err
	case <-time.After(timeout):
		o.t.Helper()
		o.t.Fatal("❌ Recorder: timeout waiting for Error")
	}
	return nil
}

func (o *Recorder[T]) AwaitCompleted(timeout time.Duration) {
	select {
	case <-o.completeCh:
		return
	case <-time.After(timeout):
		o.t.Helper()
		o.t.Fatal("❌ Recorder: timeout waiting for Complete")
	}
}

func NewRecorder[T any](t *testing.T) (*Recorder[T], *Observer[T]) {
	r := &Recorder[T]{
		t:          t,
		nextCh:     make(chan struct{}, 100), // allow multiple signals
		errorCh:    make(chan struct{}, 1),
		completeCh: make(chan struct{}, 1),
	}

	o := &Observer[T]{
		Next: func(v T) {
			r.mu.Lock()
			defer r.mu.Unlock()
			if r.terminated.Load() {
				t.Helper()
				t.Errorf("❌ Recorder: 'next' after termination")
				return
			}
			r.values = append(r.values, v)
			// non-blocking send but buffer big enough to avoid losing signals
			select {
			case r.nextCh <- struct{}{}:
			default:
			}
		},
		Error: func(err error) {
			if r.terminated.CompareAndSwap(false, true) {
				r.err = err
				select {
				case r.errorCh <- struct{}{}:
				default:
				}
			} else {
				t.Helper()
				t.Errorf("❌ Recorder: 'error' after termination")
			}
		},
		Complete: func() {
			if r.terminated.CompareAndSwap(false, true) {
				r.completed.Store(true)
				select {
				case r.completeCh <- struct{}{}:
				default:
				}
			} else {
				t.Helper()
				t.Errorf("❌ Recorder: 'complete' after termination")
			}
		},
	}
	return r, o
}

func NewRecorderFor[T any](t *testing.T, custom *Observer[T]) (*Recorder[T], *Observer[T]) {
	if custom == nil {
		panic("custom observer is nil. use NewRecorder[T] to create a default Observer[T]")
	}
	custom = Create(custom, gates.Atomic())

	r := &Recorder[T]{
		t:          t,
		nextCh:     make(chan struct{}, 100), // allow multiple signals
		errorCh:    make(chan struct{}, 1),
		completeCh: make(chan struct{}, 1),
	}

	o := &Observer[T]{
		Next: func(v T) {
			r.mu.Lock()
			defer r.mu.Unlock()
			if r.terminated.Load() {
				t.Helper()
				t.Errorf("❌ Recorder: 'next' after termination")
				return
			}
			custom.Next(v)
			r.values = append(r.values, v)
			// non-blocking send but buffer big enough to avoid losing signals
			select {
			case r.nextCh <- struct{}{}:
			default:
			}
		},
		Error: func(err error) {
			if r.terminated.CompareAndSwap(false, true) {
				r.err = err
				custom.Error(err)
				select {
				case r.errorCh <- struct{}{}:
				default:
				}
			} else {
				t.Helper()
				t.Errorf("❌ Recorder: 'error' after termination")
			}
		},
		Complete: func() {
			if r.terminated.CompareAndSwap(false, true) {
				r.completed.Store(true)
				custom.Complete()
				select {
				case r.completeCh <- struct{}{}:
				default:
				}
			} else {
				t.Helper()
				t.Errorf("❌ Recorder: 'complete' after termination")
			}
		},
	}
	return r, o
}
