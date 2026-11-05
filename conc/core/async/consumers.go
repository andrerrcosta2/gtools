// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

func ConsumeAfterSignal[T any](blocker syncs.Blocker, fn function.Consumer[T], data T) {
	go func() {
		<-blocker.Done()
		fn(data)
	}()
}

// CountedConsume consumes a data asynchronously using a function.Consumer increasing a CounterCompleter
// and releasing it when the task is finished
func CountedConsume[T any](counter syncs.CountCompleter, fn function.Consumer[T], data T) {
	counter.Add(1)
	go func() {
		defer func() {
			counter.Done()
		}()

		fn(data)
	}()
}

// SemCountedConsume consumes a data asynchronously using a function.Consumer increasing a CounterCompleter
// using a semaphore to control access and releasing both when the task is finished
func SemCountedConsume[T any](counter syncs.CountCompleter, sem syncs.Semaphore, fn function.Consumer[T], data T) {
	counter.Add(1)
	sem.Acq()
	go func() {
		defer func() {
			sem.Rls()
			counter.Done()
		}()

		fn(data)
	}()
}
