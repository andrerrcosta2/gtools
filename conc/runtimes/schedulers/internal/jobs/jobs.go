// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package jobs

import (
	"context"

	"github.com/andrerrcosta2/gtools/conc/timing/schedulers"
	"github.com/andrerrcosta2/gtools/conc/timing/schedulers/internal/cancellation"
)

func WithCancellation(job schedulers.Job, token *cancellation.Token) schedulers.Job {
	return func(ctx context.Context) error {
		if token.Cancelled() || ctx.Err() != nil {
			return context.Canceled
		}

		return job(ctx)
	}
}
