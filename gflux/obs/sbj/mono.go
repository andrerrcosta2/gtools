// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sbj

import (
	"errors"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"sync"
)

// Mono creates a mono subject
//
// A mono subject is a subject that can have only one subscriber
func Mono[T any](obv *obs.Observer[T]) obs.Subject[T] {
	return &mono[T]{
		obv: obv,
		mtx: sync.Mutex{},
	}
}

type mono[T any] struct {
	mtx sync.Mutex
	obv *obs.Observer[T]
	sct *obs.Sub[T]
}

func (s *mono[T]) Emit(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.obv.Next(t)
}

func (s *mono[T]) OnComplete() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.obv.Complete()
}

func (s *mono[T]) OnError(err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.obv.Error(err)
}

func (s *mono[T]) RemoveSub(sub *obs.Sub[T]) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.obv = nil
	if sub == s.sct {
		s.sct.Close()
	}
}

func (s *mono[T]) Sub(obv *obs.Observer[T]) (*obs.Sub[T], error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.obv != nil {
		return nil, errors.New("this subject was already subscribed")
	}
	s.obv = obv
	s.sct = obs.NewSub[T](s)
	return s.sct, nil
}

func (s *mono[T]) Unsub() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.obv = nil
}

func (s *mono[T]) UnsubscribeAll() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.obv = nil
	s.sct = nil
}

func (s *mono[T]) Closed() bool {
	return s.sct.Closed()
}

var _ obs.Subject[string] = (*mono[string])(nil)
var _ obs.Obs[string] = (*mono[string])(nil)
var _ obs.Unsubscribable[string] = (*mono[string])(nil)
var _ obs.Subscribable[string] = (*mono[string])(nil)
