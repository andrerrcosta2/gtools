// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package contexts

import (
	"context"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
)

type childrenMap = map[context.Context]struct{}

type ctxStream []context.Context

func (s ctxStream) Iterator() stream.Iterator[context.Context] {
	return func(fn func(context.Context) bool) bool {
		for _, ctx := range s {
			if !fn(ctx) {
				return false
			}
		}
		return true
	}
}
