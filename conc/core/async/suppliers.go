// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"github.com/andrerrcosta2/gtools/conc/core/future"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// Supply runs a supplier asynchronously and returns a future.Future to receive its values.
func Supply[T any](supplier future.Supplier[T]) future.Future[T] {
	fut, prom := newFuture[T]()
	go func() {
		prom.accept(supplier())
	}()
	return fut
}

func SupplyAfterSignal[T any](supplier future.Supplier[T], blocker syncs.Blocker) future.Future[T] {
	fut, prom := newFuture[T]()
	go func() {
		<-blocker.Done()
		prom.accept(supplier())
	}()
	return fut
}

func SupplyInBatches[S ~[]E, E any](supplier function.BatchSupplier[S, E]) <-chan E {
	result := make(chan E)

	go func() {
		defer close(result)

		for {
			values, ok := supplier()
			if !ok {
				return
			}

			for _, value := range values {
				result <- value
			}
		}
	}()

	return result
}
