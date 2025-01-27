// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package syncs

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
)

type ContextRunner interface {
	Before(runnable functions.Runnable, ctx context.Context)
}
