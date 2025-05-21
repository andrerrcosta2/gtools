package sbj

import (
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"sync"
)

type many[T any] struct {
	mtx       sync.RWMutex
	observers map[string]*obs.Observer[T]
	open      bool
}

func Many[T any]() obs.Subject[T] {
	return &many[T]{
		mtx:       sync.Mutex{},
		observers: make(map[string]*obs.Observer[T]),
		open:      true,
	}
}

func (s *many[T]) Closed() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return !s.open
}

func (s *many[T]) Emit(d T) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if !s.open {
		return
	}

	for _, o := range s.observers {
		o.Next(d)
	}
}

func (s *many[T]) OnComplete() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, o := range s.observers {
		o.Complete()
	}
}

func (s *many[T]) OnError(e error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, o := range s.observers {
		o.Error(e)
	}
}

func (s *many[T]) Sub(observer *obs.Observer[T]) (*obs.Sub[T], error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	sct := obs.NewSub[T](s)
	s.observers[sct.Gid()] = observer

	return sct, nil
}

func (s *many[T]) RemoveSub(sub *obs.Sub[T]) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	delete(s.observers, sub.Gid())
}

func (s *many[T]) UnsubAll() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.observers = make(map[string]*obs.Observer[T])
}

func (s *many[T]) Cld() bool {
	return !s.open
}

var _ obs.Subject[string] = (*many[string])(nil)
var _ obs.Obs[string] = (*many[string])(nil)
var _ obs.Unsub[string] = (*many[string])(nil)
var _ obs.Subscriptable[string] = (*many[string])(nil)
