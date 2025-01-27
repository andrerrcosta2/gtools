// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package channels

import (
	"context"
	"github.com/andrerrcosta2/gtools/conc/contexts"
	"github.com/andrerrcosta2/gtools/conc/contexts/cancelers"
	"github.com/andrerrcosta2/gtools/core/durations"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"testing"
)

func TestSignal(t *testing.T) {
	// helper
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)
	flag := NewFlag()

	main, exit := contexts.WithConditionalCancel(context.Background())

	// Create a blocking context to wait for the signal
	tt.RunBlocking(func() {
		flagCtx, flagCancel := context.WithTimeout(main, durations.Ms(150))
		// Condition signal to be sent after the delay
		tt.AfterAsync(func() {
			tt.StackLogf("sending flag signal")
			flag.Signal()
			tt.Flag(true, "signal")
			// call cancel to ensure no context leaks
			flagCancel()
		}, flagCtx)

		consCtx, consCancel := contexts.WithConditionalCancel(main)
		tt.AsyncBefore(func() {
			for {
				select {
				case <-flag.Receiver():
					tt.StackLogf("signal received, closing context")
					tt.Flag(true, "signal")
					consCancel.Now()
					return
				default:
					tt.RegisterCalls(1, "consumer")
				}
			}
		}, consCtx)
		exit.When(cancelers.All())
	}, main)

	tt.AssertFlag(true, "signal")
	tt.Condition(tt.CallsTo("consumer") > 0, "At least one consumptions should be executed, got %d", tt.CallsTo("consumer"))
	tt.PrintLogStack()
}
