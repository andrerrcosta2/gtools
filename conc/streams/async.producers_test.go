// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"context"
	"testing"

	"github.com/andrerrcosta2/gtools/conc/core/async"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/contexts"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/contexts/cancelers"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono/durations"
	"github.com/andrerrcosta2/gtools/core/domain/function/runnables"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

// TestBufferedProducer_Async tests the asynchronous behavior of the
// AsyncProducer. It asserts that the supplier is called asynchronously
// when the whole stream is being consumed.
func TestBufferedProducer_AsyncBehaviour(t *testing.T) {
	// Create a producer with a buffer of 5
	producer := AsyncProducer[string](5)

	// Create a context to run the test
	ctx, closeRunnable := context.WithCancel(context.Background())

	counter := testlite.Map(t)
	flags := testlite.Flag(t)

	// Before blocking
	runnables.ContextRelease(func() {
		// Produce asynchronously
		err := producer.Produce(func() []string {
			// Increment the counter of asynchronous calls
			counter.Add(1, "goroutines", "producer")
			// Simulate a long-running producer
			chrono.Sleep(durations.Ms(30))
			// Log the producer call
			// Return the values
			return smallSupply
		})

		// Log after producer call

		// Assert asynchronous behaviour
		counter.AssertCounts(0, "goroutines", "producer")

		// Shouldn't fail
		if err != nil {
			t.Error("no error was expected while supplying asynchronously, got", err)
		}

		// Consume asynchronously
		go func() {
			// Register call
			counter.Add(1, "goroutines")
			// Consume all values from the stream
			Range(producer.Stream(), func(item string) bool {
				// Log the consumer call
				// Increment the counter of consumed values
				counter.Add(1, "consumed-values")
				return true
			})
			// ToSet the flag to assert asynchronous consumption
			flags.Set(true, "asynchronous-time")
			// Close the context
			closeRunnable()
		}()

		// Log after consumer call

		// Assert none of the asynchronous components registered their commands
		counter.AssertCounts(0, "goroutines", "consumed-values", "producer")

		// ToSet the flag as false to assert asynchronous consumption
		flags.Set(false, "asynchronous-time")

		// Log at the end of the context
	}, ctx)

	// Assert asynchronous consumption
	flags.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")

	// Assert 2 goroutines calls inside main context
	counter.AssertCounts(2, "goroutines")

	// Assert all values were consumed
	counter.AssertCounts(len(smallSupply), "consumed-values")

	// Assert 1 producer call
	counter.AssertCounts(1, "producer")

}

// TestBufferedChannelProducer_Closed tests that a supplier with a buffer of 5 can be closed and that it doesn't
// accept new values after it is closed.
func TestBufferedChannelProducer_Closed(t *testing.T) {

	counter := testlite.Map(t)
	flags := testlite.Flag(t)

	// Create a producer with a buffer of 5
	producer := AsyncProducer[string](5)

	// Close the producer
	err := producer.Close()

	// Assert no error
	if err != nil {
		t.Errorf("expected no error, got %s\n", err)
	}

	// Create a context to run the test
	ctx, closeRunnable := context.WithCancel(context.Background())
	runnables.ContextRelease(func() {

		// Produce asynchronously after the producer is closed
		err = producer.Produce(func() []string {
			// Simulate a long-running producer
			chrono.Sleep(durations.Ms(300))
			// Increment the counter of asynchronous calls
			counter.Add(1, "goroutines", "producer")
			// Log the producer call
			// Return the values
			return smallSupply
		})

		// Assert no calls were added before the stream is consumed
		counter.AssertCounts(0, "goroutines", "producer")

		// Should fail supplying after the producer is closed
		if err == nil {
			t.Error("'closed producer' error was expected, got nil")
		}

		// Consume asynchronously
		go func() {
			// Consume all values from the stream
			Range(producer.Stream(), func(item string) bool {
				// Increment the counter of asynchronous calls
				counter.Add(1, "consumed-values", "goroutines")
				// Log the consumer call
				return true
			})
			// ToSet the flag to assert asynchronous consumption
			flags.Set(true, "asynchronous-time")
			// Close the context
			closeRunnable()
		}()

		// Log after consumer call

		// Assert no calls before stream is consumed
		counter.AssertCounts(0, "goroutines", "producer", "consumed-values")

		// ToSet the flag to assert asynchronous consumption
		flags.Set(false, "asynchronous-time")

		// Log at the end of the context
	}, ctx)

	// Assert asynchronous consumption
	flags.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")

	// Assert no goroutines ran, no values consumed and no producer call
	counter.AssertCounts(0, "goroutines", "producer", "consumed-values")

}

// TestBufferedProducer_MultipleAsyncSupplies tests the behavior of a supplier of values when it is supplied asynchronously
// multiple times. It asserts that the supplier and the consumer are run concurrently with different streams.
func TestBufferedProducer_MultipleAsyncSupplies(t *testing.T) {
	counter := testlite.Map(t)
	flags := testlite.Flag(t)

	// Create a producer
	var producer = AsyncProducer[string](5)
	// Create a context to run the test
	mainCtx, exit := contexts.Synchronized(context.Background())

	// Before the test
	runnables.ContextRelease(func() {

		// Before the asynchronous producer
		err := producer.Produce(func() []string {
			// Simulate a long-running producer
			chrono.Sleep(durations.Ms(130))
			// Increment the counter of asynchronous calls
			counter.Add(1, "goroutines", "producer")
			// Log the producer call
			// Return the values
			return bigSupply
		})

		if err != nil {
			t.Errorf("Expected no error while supplying asynchronously, got %s", err)
		}

		// Assert no calls were added before the context reaches its end
		counter.AssertCounts(0, "goroutines", "producer")

		// Log after producer call

		// Create a new context
		consumer1, exit1 := contexts.Synchronized(mainCtx)

		// Consume asynchronously once
		async.ContextReleaseRun(func() {
			defer exit1.Now()
			// Consume all values from the stream
			Range(producer.Stream(), func(item string) bool {
				// consume with some delay
				runnables.Before(func() {
					// Increment the counter of asynchronous calls
					counter.Add(1, "consumed-values", "goroutines", "consumer-one")
					// Log the consumer call
				}, durations.Ms(30).ToStd())
				return true
			})
			// Assert locally blocking consumption
			counter.AssertCounts(len(bigSupply), "consumer-one")
			// ToSet the flag to assert asynchronous consumption
			flags.Set(true, "asynchronous-time")

		}, consumer1)

		// Create a new context
		consumer2, exit2 := contexts.Synchronized(mainCtx)

		// Consume asynchronously twice
		async.ContextReleaseRun(func() {
			defer exit2.Now()
			// Consume all values from the stream
			Range(producer.Stream(), func(item string) bool {
				runnables.Before(func() {
					// Increment the counter of asynchronous calls
					counter.Add(1, "consumed-values", "goroutines", "consumer-two")
					// Log the consumer call
				}, durations.Ms(30).ToStd())
				return true
			})
			// Assert locally blocking consumption
			counter.AssertCounts(len(bigSupply), "consumer-two")
			// ToSet the flag to assert asynchronous consumption
			flags.Set(true, "asynchronous-time")

		}, consumer2)

		// exit main context when all nested contexts are done
		exit.When(cancelers.All())

		// Assert no calls were added before the context reaches its end
		counter.AssertCounts(0, "goroutines", "producer", "consumed-values", "consumer-one", "consumer-two")

		// Flag asynchronous-time as false to assert asynchronous consumption
		flags.Set(false, "asynchronous-time")

		// Log end of main context
	}, mainCtx)

	// Assert asynchronous consumption
	flags.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")
	counter.AssertCounts(len(bigSupply)*2+2, "goroutines")
	counter.AssertCounts(len(bigSupply), "consumer-one", "consumer-two")
	counter.AssertCounts(2, "producer")
	counter.AssertCounts(len(bigSupply)*2, "consumed-values")

	// Print the log stack

}

// TestBufferedProducer_UpdatedValues tests the supplier subsequent calls to a same function which
// delivers different values over time. The consumption should be different on each stream.
func TestBufferedProducer_UpdatedValues(t *testing.T) {
	counter := testlite.Map(t)
	flags := testlite.Flag(t)

	// Create a producer
	producer := AsyncProducer[string](15)
	// Create a main context
	mainCtx, exit := contexts.WithConditionalCancel(context.Background())

	// Create an asynchronous function to supply values differently
	asyncFunc := func() []string {
		// Simulate a long-running producer
		chrono.Sleep(durations.Ms(130))
		// add a call to the producer after delay
		counter.Add(1, "producer", "goroutines")
		// Return the values
		if counter.Get("producer")%2 != 0 {
			return mediumSupply
		}
		return bigSupply
	}

	// Before main context
	runnables.ContextRelease(func() {

		// Log start of main context

		// Produce asynchronously
		err := producer.Produce(asyncFunc)
		if err != nil {
			t.Errorf("no error was expected while supplying asynchronously, got %s", err)
		}

		// Assert no calls were added before the context reaches its end
		counter.AssertCounts(0, "goroutines", "producer")

		// Create a new context
		consumer1, consumer1Exit := contexts.WithConditionalCancel(mainCtx)

		// Before first consumer
		async.ContextReleaseRun(func() {

			// add small delay to ensure fully asynchronous flow
			chrono.Sleep(durations.Ms(60))

			// Register goroutine
			counter.Add(1, "goroutines")

			// Consume all values from the stream
			Range(producer.Stream(), func(item string) bool {
				// Increment the counter of asynchronous calls
				counter.Add(1, "consumed-values", "goroutines", "consumer-one")
				// Log the consumer call
				return true
			})

			// Assert locally synchronous flow
			counter.AssertCounts(len(mediumSupply), "consumer-one")

			// ToSet the flag to assert asynchronous consumption
			flags.Set(true, "asynchronous-time")

			// Close the context
			consumer1Exit.Now()
		}, consumer1)

		// Create a new context
		consumer2, consumer2Exit := contexts.WithConditionalCancel(mainCtx)

		// Before second consumer
		async.ContextReleaseRun(func() {

			// add bigger delay to ensure it consumes the second stream
			chrono.Sleep(durations.Ms(90))

			// Register goroutine
			counter.Add(1, "goroutines")

			// Consume all values from the stream
			Range(producer.Stream(), func(item string) bool {
				// Increment the counter of asynchronous calls
				counter.Add(1, "consumed-values", "goroutines", "consumer-two")
				// Log the consumer call
				return true
			})

			// Assert locally synchronous flow
			counter.AssertCounts(len(bigSupply), "consumer-two")

			// ToSet the flag to assert asynchronous consumption
			flags.Set(true, "asynchronous-time")

			// Close the context
			consumer2Exit.Now()
		}, consumer2)

		// Exit after both inner contexts are done
		exit.When(cancelers.All())
		// Flag asynchronous time as false to assert asynchronous flow
		flags.Set(false, "asynchronous-time")
		// Log end of main context
		// Assert no calls were added before the context reaches its end
		counter.AssertCounts(0, "goroutines", "consumed-values", "consumer-one", "consumer-two", "producer")

	}, mainCtx)

	// Assert asynchronous-time flags true
	flags.AssertFlagTo("asynchronous-time", true, "asynchronous-time was not set to true")

	// Assert all callers were fully fetched
	counter.AssertCounts(counter.Get("consumed-values")+4, "goroutines")
	counter.AssertCounts(2, "producer")
	counter.AssertCounts(len(bigSupply)+len(mediumSupply), "consumed-values")
	counter.AssertCounts(len(mediumSupply), "consumer-one")
	counter.AssertCounts(len(bigSupply), "consumer-two")

	// Print log

}

func TestCancellableProducer_AsyncBehaviour(t *testing.T) {
	counter := testlite.Map(t)
	flags := testlite.Flag(t)

	// Create a producer
	producer := AsyncCancellableProducer[string](context.Background(), 10)
	// Create a context to run the test
	ctx, closeRunnable := context.WithCancel(context.Background())

	// Before blocking
	runnables.ContextRelease(func() {

		// Produce asynchronously
		err := producer.Produce(func() []string {
			// Simulate a long-running producer
			chrono.Sleep(durations.Ms(200))
			// Increment the counter of asynchronous calls
			counter.Add(1, "goroutines", "producer")
			// Log the producer call
			// Return the values
			return bigSupply
		})

		counter.AssertCounts(0, "goroutines", "producer")

		// Shouldn't fail
		if err != nil {
			t.Error("no error was expected while supplying asynchronously, got", err)
		}

		// Consume asynchronously
		go func() {
			chrono.Sleep(durations.Ms(200))
			// Register call
			counter.Add(1, "goroutines")
			// Consume all values from the stream
			Range(producer.Stream(), func(item string) bool {
				// Log the consumer call
				// Increment the counter of consumed values
				counter.Add(1, "consumed-values")
				return true
			})
			// ToSet the flag to assert asynchronous consumption
			flags.Set(true, "asynchronous-time")
			// Close the context
			closeRunnable()
		}()

		// Log after consumer call

		// Assert none of the asynchronous components registered their calls
		counter.AssertCounts(0, "goroutines", "consumed-values", "producer")

		// ToSet the flag as false to assert asynchronous consumption
		flags.Set(false, "asynchronous-time")

		// Log at the end of the context
	}, ctx)

	// Assert asynchronous consumption
	flags.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")

	// Assert 2 goroutines calls inside main context
	counter.AssertCounts(2, "goroutines")

	// Assert 10 values consumed
	counter.AssertCounts(len(bigSupply), "consumed-values")

	// Assert 1 producer call
	counter.AssertCounts(1, "producer")

}

// TestCancellableChannelProducer_Close tests the behavior of the AsyncCancellableProducer when it is closed and then supplied.
// It asserts that the supplier is closed and that the stream is not consumed.
func TestCancellableChannelProducer_Cancel(t *testing.T) {
	counter := testlite.Map(t)
	flags := testlite.Flag(t)

	// Create a producer with a buffer of 5
	background, cancel := context.WithCancel(context.Background())
	producer := AsyncCancellableProducer[string](background, 5)

	// Close the producer
	cancel()

	// Create a context to run the test
	ctx, closeRunnable := context.WithCancel(context.Background())
	runnables.ContextRelease(func() {

		// Produce asynchronously after the producer is closed
		err := producer.Produce(func() []string {
			// Simulate a long-running producer
			chrono.Sleep(durations.Ms(300))
			// Increment the counter of asynchronous calls
			counter.Add(1, "goroutines", "producer")
			// Log the producer call
			// Return the values
			return smallSupply
		})

		// Assert error on producer call
		assertlite.IsErrorOf(t, context.Canceled, err)

		// Consume asynchronously
		go func() {
			// Consume all values from the stream
			Range(producer.Stream(), func(item string) bool {
				// Increment the counter of asynchronous calls
				counter.Add(1, "consumed-values", "goroutines")
				// Log the consumer call
				return true
			})
			// ToSet the flag to assert asynchronous consumption
			flags.Set(true, "asynchronous-time")
			// Close the context
			closeRunnable()
		}()

		// Log after consumer call

		// Assert no calls before stream is consumed
		counter.AssertCounts(0, "goroutines", "producer", "consumed-values")

		// ToSet the flag to assert asynchronous consumption
		flags.Set(false, "asynchronous-time")

		// Log at the end of the context
	}, ctx)

	// Assert asynchronous consumption
	flags.AssertFlagTo("asynchronous-time", true, "Expected asynchronous consumption, got synchronous")

	// Assert no goroutines ran, no values consumed and no producer call
	counter.AssertCounts(0, "goroutines", "producer", "consumed-values")

}

// TestCancellableChannelProducer_CancelOnTheFly tests the cancellableProd with a buffer of 5 by supplying a long-running supplier,
// consuming synchronously and cancelling the supplier on the fly.
func TestCancellableChannelProducer_CancelOnTheFly(t *testing.T) {
	counter := testlite.Map(t)

	// Create a cancelable context
	supplierCtx, cancel := context.WithCancel(context.Background())
	// Create a producer with a buffer of 5
	producer := AsyncCancellableProducer[string](supplierCtx, 5)
	// Create a context to run the test
	mainCtx, exitMain := context.WithCancel(context.Background())

	// Before test
	runnables.ContextRelease(func() {

		err := producer.Produce(func() []string {
			// Simulate a long-running producer
			chrono.Sleep(durations.Ms(300))
			// Increment the counter of asynchronous calls
			counter.Add(1, "goroutines", "producer")
			// Log the producer call
			// Return the values
			return bigSupply
		})

		// Assert no error
		if err != nil {
			t.Errorf("expected no error, got %s\n", err)
		}

		// Assert asynchronous behaviour
		counter.AssertCounts(0, "goroutines", "producer")

		// consume synchronously
		Range(producer.Stream(), func(item string) bool {

			// Increment the counter of asynchronous calls
			counter.Add(1, "consumed-values", "goroutines")

			// Cancel the producer on the fly
			if counter.Get("consumed-values") >= 5 {
				// Cancel the producer
				cancel()
			}

			// Log the consumer call
			return true
		})

		// Exit the main context
		exitMain()

	}, mainCtx)

	// Assert on call to producer
	counter.AssertCountTo("producer", 1, "expected 1 producer call, got %d\n", counter.Get("producer"))

	consumes := counter.Get("consumed-values")
	goroutines := counter.Get("goroutines")
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

}
