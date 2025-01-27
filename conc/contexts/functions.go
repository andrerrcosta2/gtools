// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package contexts

import (
	"context"
)

func AsContext(ctx any) (context.Context, bool) {
	cont, ok := ctx.(context.Context)
	return cont, ok
}

func AsSynchronizable(ctx any) (Synchronizable, bool) {
	cont, ok := ctx.(Synchronizable)
	return cont, ok
}

func AsCancellable(ctx any) (Cancellable, bool) {
	if cont, ok := ctx.(Cancellable); ok {
		if cont.Done() != nil {
			return cont, ok
		}
	}
	return nil, false
}
