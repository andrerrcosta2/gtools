// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sbj

import (
	"errors"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"sync"
)

// Mono creates a Mono sbj
//
// A Mono sbj is a one-to-one obs.Subject that can have only one subscriber
func Mono[T any](obv *obs.Observer[T]) obs.Subject[T] {
	return &mono[T]{
		obv: obv,
	}
}

type mono[T any] struct {
	obv    *obs.Observer[T]
	closed bool
	sub    *obs.Sub[T]
}

func (s *mono[T]) Emit(t T) {
	if s.closed {
		return
	}
	s.obv.Next(t)
}

func (s *mono[T]) OnComplete() {
	if s.closed {
		return
	}
	s.closed = true

	s.obv.Complete()
}

func (s *mono[T]) OnError(err error) {
	if s.closed {
		return
	}

	s.obv.Error(err)
}

func (s *mono[T]) RemoveSub(sub *obs.Sub[T]) {
	s.obv = nil
	if sub == s.sub {
		s.sub.Close()
	}
}

func (s *mono[T]) Sub(obv *obs.Observer[T]) (*obs.Sub[T], error) {
	if s.obv != nil {
		return nil, errors.New("this sbj was already subscribed")
	}
	s.obv = obv
	s.sub = obs.NewSub[T](s)
	return s.sub, nil
}

func (s *mono[T]) Unsub() {
	s.obv = nil
}

func (s *mono[T]) UnsubAll() {
	s.obv = nil
	s.sub = nil
}

func (s *mono[T]) Closed() bool {
	return s.sub.Closed()
}

// ConcMono creates a concurrent Mono sbj
//
// A Mono sbj is a sbj that can have only one subscriber
func ConcMono[T any](obv *obs.Observer[T]) obs.Subject[T] {
	return &concMono[T]{
		obv: obv,
		mtx: sync.RWMutex{},
	}
}

type concMono[T any] struct {
	mtx    sync.RWMutex
	obv    *obs.Observer[T]
	closed bool
	sct    *obs.Sub[T]
}

func (s *concMono[T]) Emit(t T) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if s.closed {
		return
	}
	s.obv.Next(t)
}

func (s *concMono[T]) OnComplete() {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if s.closed {
		return
	}
	s.closed = true

	s.obv.Complete()
}

func (s *concMono[T]) OnError(err error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if s.closed {
		return
	}

	s.obv.Error(err)
}

func (s *concMono[T]) RemoveSub(sub *obs.Sub[T]) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.obv = nil
	if sub == s.sct {
		s.sct.Close()
	}
}

func (s *concMono[T]) Sub(obv *obs.Observer[T]) (*obs.Sub[T], error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.obv != nil {
		return nil, errors.New("this sbj was already subscribed")
	}
	s.obv = obv
	s.sct = obs.NewSub[T](s)
	return s.sct, nil
}

func (s *concMono[T]) Unsub() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.obv = nil
}

func (s *concMono[T]) UnsubAll() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.obv = nil
	s.sct = nil
}

func (s *concMono[T]) Closed() bool {
	return s.sct.Closed()
}
