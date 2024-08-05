package sbj

import (
	"github.com/andrerrcosta2/gtools/pkg/obs"
	"sync"
)

type Dsbj[T any] struct {
	mtx  sync.Mutex
	obvs map[string]*obs.Obv[T]
	opn  bool
}

func NewDsbj[T any]() *Dsbj[T] {
	return &Dsbj[T]{
		mtx:  sync.Mutex{},
		obvs: make(map[string]*obs.Obv[T]),
		opn:  true,
	}
}

func (s *Dsbj[T]) Sub(obv *obs.Obv[T]) (*obs.Sub[T], error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	sct := obs.NewSub[T](s)
	s.obvs[sct.Gid()] = obv

	return sct, nil
}

func (s *Dsbj[T]) Nxt(d T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, o := range s.obvs {
		o.Nxt(d)
	}
}

func (s *Dsbj[T]) Err(e error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, o := range s.obvs {
		o.Err(e)
	}
}

func (s *Dsbj[T]) Cpt() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, o := range s.obvs {
		o.Cpt()
	}
}

func (s *Dsbj[T]) Rmo(sub *obs.Sub[T]) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	delete(s.obvs, sub.Gid())
}

func (s *Dsbj[T]) Uns() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.obvs = make(map[string]*obs.Obv[T])
}

func (s *Dsbj[T]) Cld() bool {
	return !s.opn
}

var _ obs.Sbj[string] = (*Dsbj[string])(nil)
var _ obs.Obs[string] = (*Dsbj[string])(nil)
var _ obs.Uns[string] = (*Dsbj[string])(nil)
var _ obs.Sct[string] = (*Dsbj[string])(nil)
