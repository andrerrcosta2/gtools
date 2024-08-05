// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// Package sbj: Incomlete.
package sbj

import (
	"github.com/andrerrcosta2/gtools/pkg/obs"
	"sync/atomic"
)

type Stt int

const (
	LOT Stt = iota
	DET
)

type Obschn[R any] struct {
	obv Mono[R]
	opn atomic.Bool
}

func NewObschn[R any](arr ...R) *Obschn[R] {
	o := &Obschn[R]{
		obv: Mono[R]{},
	}
	return o
}

func (s *Obschn[R]) Sub(obv *obs.Obv[R]) (*obs.Sub[R], error) {
	sct := obs.NewSub[R](s)
	return sct, nil
}

func (o *Obschn[R]) Nxt(v R) {
	o.obv.Nxt(v)
}

func (o *Obschn[R]) Err(err error) {
	o.obv.Err(err)
}

func (o *Obschn[R]) Cpt() {
	o.obv.Cpt()
}

func (o *Obschn[R]) Rmo(s *obs.Sub[R]) {
	//TODO implement me
	panic("implement me")
}

func (o *Obschn[R]) Uns() {
	//TODO implement me
	panic("implement me")
}

func (o *Obschn[R]) Cld() bool {
	//TODO implement me
	panic("implement me")
}

var _ obs.Sbj[string] = (*Obschn[string])(nil)
var _ obs.Obs[string] = (*Obschn[string])(nil)
var _ obs.Uns[string] = (*Obschn[string])(nil)
var _ obs.Sct[string] = (*Obschn[string])(nil)
