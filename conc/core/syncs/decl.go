// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package syncs

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type ContextRunner interface {
	Before(runnable function.Runnable, ctx context.Context)
}
