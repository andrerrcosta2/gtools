// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

func PredicateAfterSignal[T any](blocker syncs.Blocker, fn function.Predicate[T], data T) {
	go func() {
		<-blocker.Done()
		fn(data)
	}()
}

func CountedPredicate[T any](counter syncs.CountCompleter, fn function.Predicate[T], data T) {
	counter.Add(1)
	go func() {
		defer func() {
			counter.Done()
		}()

		fn(data)
	}()
}

func SemCountedPredicate[T any](counter syncs.CountCompleter, sem syncs.Semaphore, fn function.Predicate[T], data T) {
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
