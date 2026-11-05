// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package parallel

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// Consume consumes each item from a slice unbounded
func Consume[S ~[]E, E any](arr S, f function.Consumer[E]) {
	for _, v := range arr {
		go func(v E) {
			f(v)
		}(v)
	}
}

// ConsumeAndWait calls a consumer for each element in arr using goroutines.
// It uses a wait group to wait for all goroutines to finish before release
// the main thread.
func ConsumeAndWait[S ~[]E, E any](arr S, f function.Consumer[E]) {
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, v := range arr {
		go func(v E) {
			defer wg.Done()
			f(v)
		}(v)
	}
	wg.Wait()
}

// ConsumeSemAndWait calls a consumer for each element in arr using goroutines.
// It uses a semaphore to limit the number of concurrent ops
func ConsumeAcqRlsAndWait[S ~[]E, E any](arr S, f function.Consumer[E], sem syncs.Limiter) {
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, v := range arr {
		go func(v E) {
			sem.Acq()
			defer func() {
				sem.Rls()
				wg.Done()
			}()
			f(v)
		}(v)
	}
	wg.Wait()
}
