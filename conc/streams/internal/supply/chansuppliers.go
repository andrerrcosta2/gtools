// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package supply

import (
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

func CancellableChannelSync[T any](values []T, channel chan T, leave syncs.Blocker) {
	go func() {
		defer close(channel)
		for _, value := range values {
			select {
			case channel <- value:
			case <-leave.Done():
				// Stop sending if an early close signal is received
				return
			}
		}
	}()
}

// CancellableChannelAsync fetches a channel with the provided values from the supplier function
// and closes when it is exhausted, or when it receives a signal to unblock
func CancellableChannelAsync[T any](channel chan T, blocker syncs.Blocker, supplier function.Supplier[[]T]) {
	go func() {
		defer close(channel)

		for {
			values := supplier()

			for _, value := range values {
				select {
				case channel <- value:
				case <-blocker.Done():
					return
				}
			}
		}
	}()
}

// BatchCancellableChannelAsync keep fetching a channel with the provided values from the supplier function
// until it returns false, and closes when it is exhausted, or when it receives a signal to unblock
func BatchCancellableChannelAsync[T any](channel chan T, blocker syncs.Blocker, supplier function.BatchSupplier[[]T, T]) {
	go func() {
		defer close(channel)

		for {
			batch, ok := supplier()
			if !ok {
				return
			}

			for _, value := range batch {
				select {
				case channel <- value:
				case <-blocker.Done():
					return
				}
			}
		}
	}()
}
