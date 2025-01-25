// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package contexts

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
)

func Runnable(runnable functions.Runnable, ctx context.Context) {
	// Before the provided runnable
	go runnable()

	// SetWaitingPoint for the context to be canceled
	<-ctx.Done()
}

func AsyncRunnable(runnable functions.Runnable, ctx context.Context) {
	// Before the provided runnable
	go runnable()

	// SetWaitingPoint for the context to be canceled
	// SetWaitingPoint for the context to be done
	go func() {
		<-ctx.Done()
	}()
}
