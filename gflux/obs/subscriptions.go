// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package obs

import (
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

type (
	// Subscription represents cooperative cancellation.
	//
	// For synchronous cold Observables (e.g. Of, Range),
	// emissions may occur before Sub returns.
	// In such cases, Unsub does not interrupt in-flight emissions.
	//
	// For asynchronous Observables, Unsub actively stops
	// future emissions and releases resources.
	Subscription = subscr.Subscription
)
