// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sbj

import (
	"errors"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"sync"
)

func Many[T any]() obs.Subject[T] {
	return &many[T]{
		observers: make(map[string]*obs.Observer[T]),
		open:      true,
	}
}

// concMany is a obs.Subject that allows multiple observers
type many[T any] struct {
	observers map[string]*obs.Observer[T]
	open      bool
}

func (s *many[T]) Closed() bool {
	return !s.open
}

func (s *many[T]) Emit(d T) {
	if !s.open {
		return
	}

	for _, o := range s.observers {
		o.Next(d)
	}
}

func (s *many[T]) OnComplete() {
	if !s.open {
		return
	}

	for _, o := range s.observers {
		o.Complete()
	}
}

func (s *many[T]) OnError(e error) {
	if !s.open {
		return
	}

	for _, o := range s.observers {
		o.Error(e)
	}
}

func (s *many[T]) Sub(observer *obs.Observer[T]) (*obs.Sub[T], error) {
	if !s.open {
		return nil, errors.New("subject is closed")
	}
	sct := obs.NewSub[T](s)
	s.observers[sct.Gid()] = observer
	return sct, nil
}

func (s *many[T]) RemoveSub(sub *obs.Sub[T]) {
	delete(s.observers, sub.Gid())
}

func (s *many[T]) Close() {
	s.open = false
}

func (s *many[T]) UnsubAll() {
	if !s.open {
		return
	}
	s.open = false
	for _, o := range s.observers {
		o.Complete()
	}
	s.observers = nil
}

var _ obs.Subject[string] = (*many[string])(nil)
var _ obs.Obs[string] = (*many[string])(nil)
var _ obs.Closable = (*many[string])(nil)
var _ obs.Subscribable[string] = (*many[string])(nil)

func ConcMany[T any]() obs.Subject[T] {
	return &concMany[T]{
		mtx:       sync.RWMutex{},
		observers: make(map[string]*obs.Observer[T]),
		open:      true,
	}
}

// concMany is a sbj that allows multiple observers
type concMany[T any] struct {
	mtx       sync.RWMutex
	observers map[string]*obs.Observer[T]
	open      bool
}

func (s *concMany[T]) Closed() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return !s.open
}

func (s *concMany[T]) Emit(d T) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if !s.open {
		return
	}

	for _, o := range s.observers {
		o.Next(d)
	}
}

func (s *concMany[T]) OnComplete() {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if !s.open {
		return
	}

	for _, o := range s.observers {
		o.Complete()
	}
}

func (s *concMany[T]) OnError(e error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if !s.open {
		return
	}

	for _, o := range s.observers {
		o.Error(e)
	}
}

func (s *concMany[T]) Sub(observer *obs.Observer[T]) (*obs.Sub[T], error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	sct := obs.NewSub[T](s)
	s.observers[sct.Gid()] = observer

	return sct, nil
}

func (s *concMany[T]) RemoveSub(sub *obs.Sub[T]) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	delete(s.observers, sub.Gid())
}

func (s *concMany[T]) Close() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.open = false
}

func (s *concMany[T]) UnsubAll() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.observers = make(map[string]*obs.Observer[T])
}

var _ obs.Subject[string] = (*concMany[string])(nil)
var _ obs.Obs[string] = (*concMany[string])(nil)
var _ obs.Closable = (*concMany[string])(nil)
var _ obs.Subscribable[string] = (*concMany[string])(nil)
