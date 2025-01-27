// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package contexts

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
)

func Runnable(f functions.Runnable, ctx context.Context) {
	// Before the provided runnable
	f()

	// SetWaitingPoint for the context to be canceled
	<-ctx.Done()
}

// BeforeAsync executes the given function in a new goroutine and returns immediately holding its context.
// That function is useful for contexts that require synchronized cancellations.
func BeforeAsync(f functions.Runnable, ctx context.Context) {
	go func() {
		// Before the provided runnable
		f()
		<-ctx.Done()
	}()
}

// AsyncCallback executes the provided function inside a new goroutine after the context is canceled.
// That function is useful for contexts that require synchronized cancellations.
func AsyncCallback(f functions.Runnable, ctx context.Context) {
	go func() {
		// SetWaitingPoint for the context to be canceled
		<-ctx.Done()
		// Before the provided runnable
		f()
	}()
}
