// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package contexts

import (
	"context"
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"reflect"
)

func synchronizeContext(ctx context.Context, fn functions.Runnable) error {
	if sc, ok := AsSynchronizable(ctx); ok {
		sc.Add(1)
		go fn()
		return nil
	}
	return errors.New("cannot synchronize context")
}

func runSynchronized(ctx context.Context, fn functions.Runnable) error {
	if sc, ok := AsSynchronizable(ctx); ok {
		go fn()
		sc.Wait()
		return nil
	}
	return errors.New("cannot run synchronized context")
}

func contextName(ctx context.Context) string {
	if s, ok := ctx.(fmt.Stringer); ok {
		return s.String()
	}
	return reflect.TypeOf(ctx).String()
}
