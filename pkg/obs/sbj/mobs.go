// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sbj

import (
	"errors"
	"github.com/andrerrcosta2/gtools/pkg/obs"
	"sync"
)

type Mono[T any] struct {
	mtx sync.Mutex
	obv *obs.Obv[T]
	sct *obs.Sub[T]
}

func NewMono[T any](obv *obs.Obv[T]) *Mono[T] {
	return &Mono[T]{
		obv: obv,
		mtx: sync.Mutex{},
	}
}

func (s *Mono[T]) Sub(obv *obs.Obv[T]) (*obs.Sub[T], error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.obv != nil {
		return nil, errors.New("this subject was already subscribed")
	}
	s.obv = obv
	s.sct = obs.NewSub[T](s)
	return s.sct, nil
}

func (s *Mono[T]) Nxt(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.obv.Nxt(t)
}

func (s *Mono[T]) Rmo(sub *obs.Sub[T]) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.obv = nil
	if sub == s.sct {
		s.sct.Close()
	}
}

func (s *Mono[T]) Err(err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.obv.Err(err)
}

func (s *Mono[T]) Cpt() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.obv.Cpt()
}

func (s *Mono[T]) Uns() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.obv = nil
}

func (s *Mono[T]) Cld() bool {
	return s.sct.Closed()
}

var _ obs.Sbj[string] = (*Mono[string])(nil)
var _ obs.Obs[string] = (*Mono[string])(nil)
var _ obs.Uns[string] = (*Mono[string])(nil)
var _ obs.Sct[string] = (*Mono[string])(nil)
