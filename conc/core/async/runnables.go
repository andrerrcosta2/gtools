// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

func RunAfterSignal(fn function.Runnable, blocker syncs.Blocker) {
	go func() {
		<-blocker.Done()
		fn()
	}()
}

func ContextReleaseRun(fn function.Runnable, ctx syncs.Blocker) {
	go func() {
		fn()
		<-ctx.Done()
	}()
}

func CountedRun(counter syncs.CountCompleter, sem syncs.Semaphore, fn function.Runnable) {
	counter.Add(1)
	sem.Acq()
	go func() {
		defer func() {
			sem.Rls()
			counter.Done()
		}()
		fn()
	}()
}
