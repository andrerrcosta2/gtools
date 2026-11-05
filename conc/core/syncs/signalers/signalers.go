// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package signalers

import "github.com/andrerrcosta2/gtools/conc/core/syncs"

func Chan(ch chan struct{}) syncs.Signaler {
	return chanSignaler{ch: ch}
}

type chanSignaler struct {
	ch chan struct{}
}

func (s chanSignaler) Signal() {
	close(s.ch)
}

func Func(f func()) syncs.Signaler {
	return funcSignaler(f)
}

type funcSignaler func()

func (f funcSignaler) Signal() {
	f()
}
