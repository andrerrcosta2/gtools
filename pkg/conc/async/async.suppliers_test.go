// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"context"
	"github.com/andrerrcosta2/gtools/conc/contexts"
	"github.com/andrerrcosta2/gtools/conc/contexts/cancelers"
	"github.com/andrerrcosta2/gtools/core/durations"
	"github.com/andrerrcosta2/gtools/core/funcs/runnables"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"testing"
	"time"
)

// TestBufferedSupplier_Async tests the asynchronous behavior of the
// BufferedSupplier. It asserts that the supplier is called asynchronously
// when the whole stream is being consumed.
func TestBufferedSupplier_AsyncBehaviour(t *testing.T) {
	// Create a supplier with a buffer of 5
	supplier := BufferedSupplier[string](5)
	// Call a helper
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)

	// Create a context to run the test
	ctx, closeRunnable := context.WithCancel(context.Background())

	// Before blocking
	tt.RunBlocking(func() {

		tt.StackLog("Supplying asynchronously")

		// Supply asynchronously
		err := supplier.Supply(func() []string {
			// Increment the counter of asynchronous calls
			tt.RegisterCalls(1, "goroutines", "supplier")
			// Simulate a long-running supplier
			time.Sleep(durations.Ms(300))
			// Log the supplier call
			tt.StackLog("supplier()")
			// Return the values
			return smallSupply
		})

		// Log after supplier call
		tt.StackLog("After supplier call")

		// Assert asynchronous behaviour
		tt.AssertCalls(0, "goroutines", "supplier")

		// Shouldn't fail
		if err != nil {
			t.Error("no error was expected while supplying asynchronously, got", err)
		}

		// Consume asynchronously
		go func() {
			// Register call
			tt.RegisterCalls(1, "goroutines")
			// Consume all values from the stream
			for value := range supplier.Stream() {
				// Log the consumer call
				tt.StackLogf("consumer(%v)", value)
				// Increment the counter of consumed values
				tt.RegisterCalls(1, "consumed-values")
			}
			// Set the flag to assert asynchronous consumption
			tt.Flag(true, "asynchronous-time")
			// Close the context
			closeRunnable()
		}()

		// Log after consumer call
		tt.StackLog("After consumer call")

		// Assert none of the asynchronous components registered their events
		tt.AssertCalls(0, "goroutines", "consumed-values", "supplier")

		// Set the flag as false to assert asynchronous consumption
		tt.Flag(false, "asynchronous-time")

		// Log at the end of the context
		tt.StackLog("End of Main ContextRelease\n++++++++++ Async flow:")
	}, ctx)

	// Assert asynchronous consumption
	tt.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")

	// Assert 2 goroutines calls inside main context
	tt.AssertCalls(2, "goroutines")

	// Assert all values were consumed
	tt.AssertCalls(len(smallSupply), "consumed-values")

	// Assert 1 supplier call
	tt.AssertCalls(1, "supplier")

	tt.PrintLogStack()
}

// TestBufferedChannelSupplier_Closed tests that a supplier with a buffer of 5 can be closed and that it doesn't
// accept new values after it is closed.
func TestBufferedChannelSupplier_Closed(t *testing.T) {
	// Create a supplier with a buffer of 5
	supplier := BufferedSupplier[string](5)
	// Call a helper
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)

	// Close the supplier
	err := supplier.Close()

	// Assert no error
	if err != nil {
		t.Errorf("expected no error, got %s\n", err)
	}

	// Create a context to run the test
	ctx, closeRunnable := context.WithCancel(context.Background())
	tt.RunBlocking(func() {

		// Supply asynchronously after the supplier is closed
		tt.StackLog("Supplying asynchronously after supplier is closed")
		err = supplier.Supply(func() []string {
			// Simulate a long-running supplier
			time.Sleep(durations.Ms(300))
			// Increment the counter of asynchronous calls
			tt.RegisterCalls(1, "goroutines", "supplier")
			// Log the supplier call
			tt.StackLog("supplier()")
			// Return the values
			return smallSupply
		})

		tt.StackLogf("After supplier call: %v", err)

		// Assert no calls were added before the stream is consumed
		tt.AssertCalls(0, "goroutines", "supplier")

		// Should fail supplying after the supplier is closed
		if err == nil {
			t.Error("'closed supplier' error was expected, got nil")
		}

		// Consume asynchronously
		go func() {
			// Consume all values from the stream
			for value := range supplier.Stream() {
				// Increment the counter of asynchronous calls
				tt.RegisterCalls(1, "consumed-values", "goroutines")
				// Log the consumer call
				tt.StackLogf("consumer(%v)", value)
			}
			// Set the flag to assert asynchronous consumption
			tt.Flag(true, "asynchronous-time")
			// Close the context
			closeRunnable()
		}()

		// Log after consumer call
		tt.StackLog("After consumer call")

		// Assert no calls before stream is consumed
		tt.AssertCalls(0, "goroutines", "supplier", "consumed-values")

		// Set the flag to assert asynchronous consumption
		tt.Flag(false, "asynchronous-time")

		// Log at the end of the context
		tt.StackLog("End of Main ContextRelease\n++++++++++ AsyncLite flow:")
	}, ctx)

	// Assert asynchronous consumption
	tt.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")

	// Assert no goroutines ran, no values consumed and no supplier call
	tt.AssertCalls(0, "goroutines", "supplier", "consumed-values")

	tt.PrintLogStack()
}

// TestBufferedSupplier_MultipleAsyncSupplies tests the behavior of a supplier of values when it is supplied asynchronously
// multiple times. It asserts that the supplier and the consumer are run concurrently with different streams.
func TestBufferedSupplier_MultipleAsyncSupplies(t *testing.T) {
	// Create a supplier
	var supplier = BufferedSupplier[string](5)
	// Call a helper
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)
	// Create a context to run the test
	mainCtx, exit := contexts.Synchronized(context.Background())

	// Before the test
	tt.RunBlocking(func() {
		tt.StackLog("Supplying asynchronously")

		// Before the asynchronous supplier
		err := supplier.Supply(func() []string {
			// Simulate a long-running supplier
			time.Sleep(durations.Ms(130))
			// Increment the counter of asynchronous calls
			tt.RegisterCalls(1, "goroutines", "supplier")
			// Log the supplier call
			tt.StackLog("supplier()")
			// Return the values
			return bigSupply
		})

		if err != nil {
			t.Errorf("Expected no error while supplying asynchronously, got %s", err)
		}

		// Assert no calls were added before the context reaches its end
		tt.AssertCalls(0, "goroutines", "supplier")

		// Log after supplier call
		tt.StackLog("After supplier call")

		// Create a new context
		consumer1, exit1 := contexts.Synchronized(mainCtx)

		// Consume asynchronously once
		tt.AsyncBefore(func() {
			defer exit1.Now()
			// Consume all values from the stream
			for value := range supplier.Stream() {
				// consume with some delay
				runnables.Before(func() {
					// Increment the counter of asynchronous calls
					tt.RegisterCalls(1, "consumed-values", "goroutines", "consumer-one")
					// Log the consumer call
					tt.StackLogf("consumer-one(%v)", value)
				}, durations.Ms(30))
			}
			// Assert locally blocking consumption
			tt.AssertCalls(len(bigSupply), "consumer-one")
			// Set the flag to assert asynchronous consumption
			tt.Flag(true, "asynchronous-time")

			tt.StackLog("consumer1 cancelled")
		}, consumer1)

		// Create a new context
		consumer2, exit2 := contexts.Synchronized(mainCtx)

		// Consume asynchronously twice
		tt.AsyncBefore(func() {
			defer exit2.Now()
			// Consume all values from the stream
			for value := range supplier.Stream() {
				runnables.Before(func() {
					// Increment the counter of asynchronous calls
					tt.RegisterCalls(1, "consumed-values", "goroutines", "consumer-two")
					// Log the consumer call
					tt.StackLogf("consumer-two(%v)", value)
				}, durations.Ms(30))
			}
			// Assert locally blocking consumption
			tt.AssertCalls(len(bigSupply), "consumer-two")
			// Set the flag to assert asynchronous consumption
			tt.Flag(true, "asynchronous-time")

			tt.StackLog("consumer2 cancelled")
		}, consumer2)

		// exit main context when all nested contexts are done
		exit.When(cancelers.All())

		// Assert no calls were added before the context reaches its end
		tt.AssertCalls(0, "goroutines", "supplier", "consumed-values", "consumer-one", "consumer-two")

		// Flag asynchronous-time as false to assert asynchronous consumption
		tt.Flag(false, "asynchronous-time")

		// Log end of main context
		tt.StackLog("End of Main ContextRelease\n++++++++ AsyncLite flow:")
	}, mainCtx)

	tt.StackLog("Main context cancelled")

	// Assert asynchronous consumption
	tt.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")
	tt.AssertCalls(len(bigSupply)*2+2, "goroutines")
	tt.AssertCalls(len(bigSupply), "consumer-one", "consumer-two")
	tt.AssertCalls(2, "supplier")
	tt.AssertCalls(len(bigSupply)*2, "consumed-values")

	// Print the log stack
	tt.PrintLogStack()
}

// TestBufferedSupplier_UpdatedValues tests the supplier subsequent calls to a same function which
// delivers different values over time. The consumption should be different on each stream.
func TestBufferedSupplier_UpdatedValues(t *testing.T) {
	// Create a supplier
	supplier := BufferedSupplier[string](15)
	// Call a helper
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)
	// Create a main context
	mainCtx, exit := contexts.WithConditionalCancel(context.Background())

	// Create an asynchronous function to supply values differently
	asyncFunc := func() []string {
		// Simulate a long-running supplier
		time.Sleep(durations.Ms(130))
		// add a call to the supplier after delay
		tt.RegisterCalls(1, "supplier", "goroutines")
		// Return the values
		if tt.CallsTo("supplier")%2 != 0 {
			return mediumSupply
		}
		return bigSupply
	}

	// Before main context
	tt.RunBlocking(func() {

		// Log start of main context
		tt.StackLog("Start of Main ContextRelease")

		// Supply asynchronously
		err := supplier.Supply(asyncFunc)
		if err != nil {
			t.Errorf("no error was expected while supplying asynchronously, got %s", err)
		}

		// Assert no calls were added before the context reaches its end
		tt.AssertCalls(0, "goroutines", "supplier")

		// Create a new context
		consumer1, consumer1Exit := contexts.WithConditionalCancel(mainCtx)

		// Before first consumer
		tt.AsyncBefore(func() {

			// add small delay to ensure fully asynchronous flow
			time.Sleep(durations.Ms(60))

			// Register goroutine
			tt.RegisterCalls(1, "goroutines")

			// Consume all values from the stream
			for value := range supplier.Stream() {
				// Increment the counter of asynchronous calls
				tt.RegisterCalls(1, "consumed-values", "goroutines", "consumer-one")
				// Log the consumer call
				tt.StackLogf("consumer-one(%v)", value)
			}

			// Assert locally synchronous flow
			tt.AssertCalls(len(mediumSupply), "consumer-one")

			// Set the flag to assert asynchronous consumption
			tt.Flag(true, "asynchronous-time")

			tt.StackLog("consumer1 cancelled")
			// Close the context
			consumer1Exit.Now()
		}, consumer1)

		// Create a new context
		consumer2, consumer2Exit := contexts.WithConditionalCancel(mainCtx)

		// Before second consumer
		tt.AsyncBefore(func() {

			// add bigger delay to ensure it consumes the second stream
			time.Sleep(durations.Ms(90))

			// Register goroutine
			tt.RegisterCalls(1, "goroutines")

			// Consume all values from the stream
			for value := range supplier.Stream() {
				// Increment the counter of asynchronous calls
				tt.RegisterCalls(1, "consumed-values", "goroutines", "consumer-two")
				// Log the consumer call
				tt.StackLogf("consumer-two(%v)", value)
			}

			// Assert locally synchronous flow
			tt.AssertCalls(len(bigSupply), "consumer-two")

			// Set the flag to assert asynchronous consumption
			tt.Flag(true, "asynchronous-time")

			tt.StackLog("consumer2 cancelled")
			// Close the context
			consumer2Exit.Now()
		}, consumer2)

		// Exit after both inner contexts are done
		exit.When(cancelers.All())
		// Flag asynchronous time as false to assert asynchronous flow
		tt.Flag(false, "asynchronous-time")
		// Log end of main context
		tt.StackLog("End of Main ContextRelease\n++++++++ AsyncLite flow:")
		// Assert no calls were added before the context reaches its end
		tt.AssertCalls(0, "goroutines", "consumed-values", "consumer-one", "consumer-two", "supplier")

	}, mainCtx)

	// Assert asynchronous-time flags true
	tt.AssertFlag(true, "asynchronous-time")

	// Assert all callers were fully fetched
	tt.AssertCalls(tt.CallsTo("consumed-values")+4, "goroutines")
	tt.AssertCalls(2, "supplier")
	tt.AssertCalls(len(bigSupply)+len(mediumSupply), "consumed-values")
	tt.AssertCalls(len(mediumSupply), "consumer-one")
	tt.AssertCalls(len(bigSupply), "consumer-two")

	// Print log
	tt.PrintLogStack()
}

func TestCancellableSupplier_AsyncBehaviour(t *testing.T) {
	// Create a supplier
	supplier := CancellableSupplier[string](context.Background(), 10)
	// Call a helper
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)
	// Create a context to run the test
	ctx, closeRunnable := context.WithCancel(context.Background())

	// Before blocking
	tt.RunBlocking(func() {

		// Supply asynchronously
		tt.StackLog("Supplying asynchronously")
		err := supplier.Supply(func() []string {
			// Simulate a long-running supplier
			time.Sleep(durations.Ms(200))
			// Increment the counter of asynchronous calls
			tt.RegisterCalls(1, "goroutines", "supplier")
			// Log the supplier call
			tt.StackLog("supplier()")
			// Return the values
			return bigSupply
		})

		// Log after supplier call
		tt.StackLog("After supplier call")

		tt.AssertCalls(0, "goroutines", "supplier")

		// Shouldn't fail
		if err != nil {
			t.Error("no error was expected while supplying asynchronously, got", err)
		}

		// Consume asynchronously
		go func() {
			time.Sleep(durations.Ms(200))
			// Register call
			tt.RegisterCalls(1, "goroutines")
			// Consume all values from the stream
			for value := range supplier.Stream() {
				// Log the consumer call
				tt.StackLogf("consumer(%v)", value)
				// Increment the counter of consumed values
				tt.RegisterCalls(1, "consumed-values")
			}
			// Set the flag to assert asynchronous consumption
			tt.Flag(true, "asynchronous-time")
			// Close the context
			closeRunnable()
		}()

		// Log after consumer call
		tt.StackLog("After consumer call")

		// Assert none of the asynchronous components registered their calls
		tt.AssertCalls(0, "goroutines", "consumed-values", "supplier")

		// Set the flag as false to assert asynchronous consumption
		tt.Flag(false, "asynchronous-time")

		// Log at the end of the context
		tt.StackLog("End of Main ContextRelease\n++++++++++ Async flow:")
	}, ctx)

	// Assert asynchronous consumption
	tt.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")

	// Assert 2 goroutines calls inside main context
	tt.AssertCalls(2, "goroutines")

	// Assert 10 values consumed
	tt.AssertCalls(len(bigSupply), "consumed-values")

	// Assert 1 supplier call
	tt.AssertCalls(1, "supplier")

	tt.PrintLogStack()
}

// TestCancellableChannelSupplier_Close tests the behavior of the CancellableSupplier when it is closed and then supplied.
// It asserts that the supplier is closed and that the stream is not consumed.
func TestCancellableChannelSupplier_Close(t *testing.T) {
	// Create a supplier with a buffer of 5
	supplier := CancellableSupplier[string](context.Background(), 5)
	// Call a helper
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)

	// Close the supplier
	err := supplier.Close()

	// Assert no error
	if err != nil {
		t.Errorf("expected no error, got %s\n", err)
	}

	// Create a context to run the test
	ctx, closeRunnable := context.WithCancel(context.Background())
	tt.RunBlocking(func() {

		// Supply asynchronously after the supplier is closed
		tt.StackLog("Supplying asynchronously after supplier is closed")
		err = supplier.Supply(func() []string {
			// Simulate a long-running supplier
			time.Sleep(durations.Ms(300))
			// Increment the counter of asynchronous calls
			tt.RegisterCalls(1, "goroutines", "supplier")
			// Log the supplier call
			tt.StackLog("supplier()")
			// Return the values
			return smallSupply
		})

		// Assert error on supplier call
		if err == nil {
			t.Error("'closed supplier' error was expected, got nil")
		}

		// Consume asynchronously
		go func() {
			// Consume all values from the stream
			for value := range supplier.Stream() {
				// Increment the counter of asynchronous calls
				tt.RegisterCalls(1, "consumed-values", "goroutines")
				// Log the consumer call
				tt.StackLogf("consumer(%v)", value)
			}
			// Set the flag to assert asynchronous consumption
			tt.Flag(true, "asynchronous-time")
			// Close the context
			closeRunnable()
		}()

		// Log after consumer call
		tt.StackLog("After consumer call")

		// Assert no calls before stream is consumed
		tt.AssertCalls(0, "goroutines", "supplier", "consumed-values")

		// Set the flag to assert asynchronous consumption
		tt.Flag(false, "asynchronous-time")

		// Log at the end of the context
		tt.StackLog("End of Main ContextRelease\n++++++++++ AsyncLite flow:")
	}, ctx)

	// Assert asynchronous consumption
	tt.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")

	// Assert no goroutines ran, no values consumed and no supplier call
	tt.AssertCalls(0, "goroutines", "supplier", "consumed-values")

	tt.PrintLogStack()
}

// TestCancellableChannelSupplier_CancelOnTheFly tests the cancellableChannelSupplier with a buffer of 5 by supplying a long-running supplier,
// consuming synchronously and cancelling the supplier on the fly.
func TestCancellableChannelSupplier_CancelOnTheFly(t *testing.T) {
	// Create a cancelable context
	supplierCtx, cancel := context.WithCancel(context.Background())
	// Create a supplier with a buffer of 5
	supplier := CancellableSupplier[string](supplierCtx, 5)
	// Call a helper
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)
	// Create a context to run the test
	mainCtx, exitMain := context.WithCancel(context.Background())

	// Before test
	tt.RunBlocking(func() {

		tt.StackLog("Supplying asynchronously")

		err := supplier.Supply(func() []string {
			// Simulate a long-running supplier
			time.Sleep(durations.Ms(300))
			// Increment the counter of asynchronous calls
			tt.RegisterCalls(1, "goroutines", "supplier")
			// Log the supplier call
			tt.StackLog("supplier()")
			// Return the values
			return bigSupply
		})

		// Assert no error
		if err != nil {
			t.Errorf("expected no error, got %s\n", err)
		}

		// Assert asynchronous behaviour
		tt.AssertCalls(0, "goroutines", "supplier")

		// consume synchronously
		for value := range supplier.Stream() {
			// Increment the counter of asynchronous calls
			tt.RegisterCalls(1, "consumed-values", "goroutines")

			// Cancel the supplier on the fly
			if tt.CallsTo("consumed-values") >= 5 {
				// Cancel the supplier
				cancel()
			}

			// Log the consumer call
			tt.StackLogf("consumer(%v)", value)
		}

		tt.StackLog("After consumer call")

		// Exit the main context
		exitMain()

	}, mainCtx)

	// Assert on call to supplier
	tt.AssertCallsTo("supplier", 1, "expected 1 supplier call, got %d\n", tt.CallsTo("supplier"))

	consumes := tt.CallsTo("consumed-values")
	goroutines := tt.CallsTo("goroutines")
	if consumes < 5 {
		t.Errorf("expected at least 5 values consumed, got %d\n", consumes)
	}
	if consumes >= len(bigSupply) {
		t.Errorf("expected less than %d values consumed, got %d\n", len(bigSupply), consumes)
	}
	if goroutines < 5 {
		t.Errorf("expected at least 5 goroutines, got %d\n", goroutines)
	}
	if goroutines >= len(bigSupply) {
		t.Errorf("expected less than %d goroutines, got %d\n", len(bigSupply), goroutines)
	}

	tt.PrintLogStack()
}
