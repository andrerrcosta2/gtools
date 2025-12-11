// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"context"
	"github.com/andrerrcosta2/gtools/conc/streams"
	"github.com/andrerrcosta2/gtools/conc/syncs/semaph"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/durations"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"testing"
	"time"
)

// TestSemaphoredConsumer_AsynchronousBehaviour tests the asynchronous behavior of the
// Consumer. It checks that the consumer is fully asynchronous by waiting for the
// runnable to finish and verifying that the consumption is done after the outer
// thread finishes.
func TestSemaphoredConsumer_AsynchronousBehaviour(t *testing.T) {
	// std
	tt := testingtools.AsyncLite(t, testlogs.OnFailure)
	// Create a cancelable context for the runnable
	ctx, exit := context.WithCancel(context.Background())

	// SetWaitingPoint for the runnable finishing
	tt.RunBlocking(func() {
		// Register runnable
		tt.RegisterCalls(1, "runs")

		// Create a stream with 10 values
		stream := streams.HotCloseable(5, smallSupply...)

		// register to exit the context based on the test caller value
		tt.RegisterCallback(functions.Runnable(exit), len(smallSupply), "consumed-values")

		// Expecting consume completely asynchronously
		err := Consumer[string](stream, semaph.Channel(3)).
			Consume(func(i int, s string) {
				tt.RegisterCalls(1, "consumed-values", "goroutine")

				// simulate light processing
				time.Sleep(durations.Ms(10))
				tt.StackLogf("consumer(%d, %v)", i, s)

				// flag asynchronous behaviour
				tt.Flag(true, "asynchronous")
			})

		// Assert async behaviour
		tt.AssertCalls(0, "consumed-values", "goroutine")

		// Assert no error
		if err != nil {
			t.Errorf("no error was expected while consuming asynchronously, got %s", err)
		}

		tt.StackLog("end of context\n+++++++++++++++")

		// Flag bottom of context
		tt.Flag(false, "asynchronous")
	}, ctx)

	tt.AssertFlag(true, "asynchronous")
	tt.AssertCalls(1, "runs")
	tt.AssertCalls(10, "consumed-values", "goroutine")

	tt.PrintLogStack()
}

// TestSemaphoredConsumer_CloseableBehaviour_CloseEarly tests the Consumer.Close() method and asserts that no values are consumed if the consumer is flag early.
// It also tests the Consumer.Consume() method and asserts that an error is returned if the consumer is flag before the Consume() method is called.
func TestSemaphoredConsumer_CloseableBehaviour_CloseEarly(t *testing.T) {
	// std
	tt := testingtools.AsyncLite(t, testlogs.OnFailure)

	ctx, exitRunnable := context.WithCancel(context.Background())

	// SetWaitingPoint for the runnable finishing
	tt.RunBlocking(func() {

		tt.RegisterCalls(1, "runs")

		stream := streams.HotCloseable(10, "a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
		consumer := Consumer[string](stream, semaph.Channel(3))

		// flag the consumer before consuming values
		err := consumer.Close()

		// expect no error
		if err != nil {
			t.Errorf("no error was expected while closing consumer, got %s", err)
		}

		// Expect no consume
		err = consumer.Consume(func(i int, s string) {
			t.Errorf("expected no consume, received value '%s'\n", s)
			_ = stream.Close()
			exitRunnable()
		})

		if err == nil {
			t.Error("expected supplier to return an error, got nil\n")
		}

		_ = stream.Close()
		exitRunnable()
	}, ctx)

	// Assert full run
	tt.AssertCalls(1, "runs")

	tt.PrintLogStack()
}

// TestSemaphoredConsumer_CancelableSupplier_CloseOnFly tests the Consumer.Close() on fly.
// This method is just asserting there's no deadlock when the async consumer is flag on fly.
func TestSemaphoredConsumer_CloseableBehaviour_ClosedOnTheFly(t *testing.T) {
	// Helper
	tt := testingtools.AsyncLite(t, testlogs.OnFailure)

	// Create a cancellable context
	mainCtx, exitMain := context.WithCancel(context.Background())
	tt.RunBlocking(func() {

		tt.StackLog("starting context")
		tt.RegisterCalls(1, "runs")

		stream := streams.HotCloseable(4, bigSupply...)

		consumer := Consumer[string](stream, semaph.Channel(3))

		tt.RegisterCallback(func() {

			err := consumer.Close()
			if err != nil {
				t.Errorf("expected no error while closing consumer, got %s", err)
			}
			_ = stream.Close()
			exitMain()

		}, 5, "consumed-values")

		err := consumer.Consume(func(i int, s string) {
			tt.RegisterCalls(1, "consumed-values", "goroutine")

			time.Sleep(durations.Ms(100))

			tt.StackLogf("consumer(%d, %v)", i, s)

			tt.Flag(true, "asynchronous")
		})

		// Assert async behaviour
		tt.AssertCalls(0, "consumed-values", "goroutine")

		if err != nil {
			t.Errorf("no error was expected while consuming asynchronously, got %s", err)
		}

		tt.Flag(false, "asynchronous")
	}, mainCtx)

	tt.AssertFlag(true, "asynchronous")
	tt.AssertCalls(1, "runs")

	goroutinesCalls := tt.CallsTo("goroutine")
	consumedValues := tt.CallsTo("consumed-values")

	if goroutinesCalls != consumedValues {
		t.Errorf("expected %d goroutines calls, got %d", consumedValues, goroutinesCalls)
	}

	if goroutinesCalls >= len(bigSupply) {
		t.Errorf("expected less than %d goroutines calls, got %d", len(bigSupply), goroutinesCalls)
	}

	tt.PrintLogStack()
}
