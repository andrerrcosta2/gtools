// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testschedulers

import (
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/conc/timing/engines/steps"
)

type cancellableStep struct {
	st        steps.Timed
	stopFunc  func()
	cancelled atomic.Bool
}

func (c *cancellableStep) Cancel() {
	if c.cancelled.CompareAndSwap(false, true) {
		if c.stopFunc != nil {
			c.stopFunc()
		}
		if c.st != nil {
			c.st.Stop()
		}
	}
}
