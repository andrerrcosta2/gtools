// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package waiters

import (
	"runtime"
	"time"

	"github.com/andrerrcosta2/gtools/conc/core/syncs"
)

func Complete() syncs.Waiter {
	return complete{}
}

type complete struct{}

func (complete) Wait() {}

func Chan(ch chan struct{}) syncs.Waiter {
	return channel{ch: ch}
}

type channel struct {
	ch chan struct{}
}

func (c channel) Wait() {
	<-c.ch
}

type Func func()

func (f Func) Wait() {
	f()
}

func YieldUntil(fn func() bool) syncs.Waiter {
	return Func(func() {
		for !fn() {
			runtime.Gosched()
		}
	})
}

func SpinUntil(fn func() bool) syncs.Waiter {
	return Func(func() {
		for !fn() {
		}
	})
}

func SleepUntil(interval time.Duration, fn func() bool) syncs.Waiter {
	return Func(func() {
		for !fn() {
			time.Sleep(interval)
		}
	})
}
