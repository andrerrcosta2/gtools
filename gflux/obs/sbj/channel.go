// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sbj

import (
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"sync/atomic"
)

type Stt int

const (
	LOT Stt = iota
	DET
)

type channel[R any] struct {
	obv    concMono[R]
	closed atomic.Bool
}

func Channel[R any](arr ...R) obs.Subject[R] {
	o := &channel[R]{
		obv: concMono[R]{},
	}
	return o
}

func (s *channel[R]) Closed() bool {
	return s.closed
}

func (s *channel[R]) OnComplete() {
	s.obv.Cpt()
}

func (s *channel[R]) OnError(err error) {
	s.obv.Err(err)
}

func (s *channel[R]) Sub(obv *obs.Observer[R]) (*obs.Sub[R], error) {
	sct := obs.NewSub[R](s)
	return sct, nil
}

func (s *channel[R]) Next(v R) {
	s.obv.Next(v)
}

func (s *channel[R]) Rmo(s *obs.Sub[R]) {
	//TODO implement me
	panic("implement me")
}

func (s *channel[R]) Unsub() {
	//TODO implement me
	panic("implement me")
}

var _ obs.Subject[string] = (*channel[string])(nil)
var _ obs.Obs[string] = (*channel[string])(nil)
var _ obs.Unsub[string] = (*channel[string])(nil)
var _ obs.Subscriptable[string] = (*channel[string])(nil)
