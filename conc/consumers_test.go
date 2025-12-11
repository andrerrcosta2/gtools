// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"context"
	"errors"
	"github.com/andrerrcosta2/gtools/conc/consumers"
	"github.com/andrerrcosta2/gtools/conc/streams"
	"github.com/andrerrcosta2/gtools/conc/syncs/semaph"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"testing"
)

// TestSemaphoredConsumer_FullConsume tests the SemaphoredConsumer.Consume() method with a HotCloseable supplier.
// It asserts that the consumer consumes all values from the supplier synchronously and concurrent.
// This test also asserts its automatic close after all values have been consumed.
func TestSemaphoredConsumer_FullConsume(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	// Create a stream with 10 values
	stream := streams.HotCloseable(4, "a", "b", "c", "d", "e", "A", "C", "T", "D", "E")

	// Consume all values from the stream
	err := SemaphoredConsumer[string](stream, semaph.Channel(3)).
		Consume(func(i int, s string) {
			// Count the number of values consumed
			tt.RegisterCalls(1, "consumes")
			// Log the value consumed
			tt.StackLogf("consumer(%d, %v)", i, s)
		})

	// Assert no error
	if err != nil {
		t.Errorf("no error was expected while consuming stream, got %s\n", err)
	}

	// Assert 10 values consumed
	tt.AssertCallsTo("consumes", 10, "expected 10 values to be consumed, got %d\n", tt.CallsTo("consumes"))

	tt.PrintLogStack()
}

// TestSemaphoredConsumer_ClosedStreamEarly tests the SemaphoredConsumer.Consume() method with a Streamable supplier that has been closed early.
// It asserts that the consumer consumes no values from the supplier synchronously and concurrent.
// This test also asserts no error is returned if the supplier is closed before the Consume() method is called.
func TestSemaphoredConsumer_ClosedStreamEarly(t *testing.T) {
	// A closed Streamable shouldn't stream any values
	stream := streams.HotCloseable(3, bigSupply...)

	// Close early
	err := stream.Close()

	// No error should be returned
	if err != nil {
		t.Errorf("no error was expected while closing stream, got %s\n", err)
	}

	// Block the main goroutine until all values have been consumed
	_ = SemaphoredConsumer[string](stream, semaph.Channel(3)).
		Consume(func(i int, s string) {
			// fail the test on consume
			t.Errorf("expected no values to be consumed, got %d\n", i)
		})

}

// TestSemaphoredConsumer_ClosedStreamOnTheFly tests the behavior of a SemaphoredConsumer when the Streamable supplier is closed on the fly.
// It asserts that not all values are consumed.
func TestSemaphoredConsumer_ClosedStreamOnTheFly(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	stream := streams.HotCloseable(3, bigSupply...)

	// close the stream after 3 values
	tt.RegisterCallback(func() {
		err := stream.Close()
		if err != nil {
			t.Errorf("no error was expected while closing stream, got %s\n", err)
		}
	}, 3, "consumes")

	// Append flag to assert synchronous flow
	tt.Flag(false, "synchronous")

	err := SemaphoredConsumer[string](stream, semaph.Channel(3)).
		Consume(func(i int, s string) {
			// Assert synchronous flow
			tt.Flag(true, "synchronous")

			// Count the number of values consumed
			tt.RegisterCalls(1, "consumes")

			// Log the value consumed
			tt.StackLogf("consumer(%d, %v)", i, s)
		})

	tt.AssertFlag(true, "synchronous")

	// Expect no error
	if err != nil {
		t.Errorf("no error was expected while consuming stream, got %s", err)
	}

	// Expect less than all values
	consumes := tt.CallsTo("consumes")
	tt.Condition(consumes < len(bigSupply), "expected to consume less than %d values, got %d", len(bigSupply), consumes)

	tt.PrintLogStack()
}

// TestSemaphoredConsumer_ClosedConsumer tests the behavior of a SemaphoredConsumer that is closed before consuming any values.
// It asserts that no values are consumed.
func TestSemaphoredConsumer_ClosedConsumer(t *testing.T) {

	// Create a supplier of 30 values
	stream := streams.HotCloseable(4, bigSupply...)

	// Create a consumer with a semaphore of 3
	consumer := SemaphoredConsumer[string](stream, semaph.Channel(3))

	// Close the consumer
	err := consumer.Close()

	// Assert no error
	if err != nil {
		t.Errorf("no error was expected while closing consumer, got %s", err)
	}

	// Consume the values
	_ = consumer.Consume(func(i int, s string) {
		// Fail the on consume
		t.Errorf("expected no values to be consumed, got %d\n", i)
	})
}

// TestSemaphoredConsumer_ClosedConsumerOnTheFly tests the behavior of the
// SemaphoredConsumer when it is closed on the fly while consuming values from
// a supplier.
func TestSemaphoredConsumer_ClosedConsumerOnTheFly(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	// Create a supplier
	stream := streams.HotCloseable(5, bigSupply...)

	// Create a semaphored consumer
	consumer := SemaphoredConsumer[string](stream, semaph.Channel(4))

	// Flag to assert synchronous flow
	tt.Flag(false, "synchronous")

	tt.RegisterCallback(func() {
		// Close the consumer after 5 values have been consumed
		err := consumer.Close()
		if err != nil {
			t.Errorf("no error was expected while closing consumer, got %s", err)
		}
	}, 5, "consumes")

	err := consumer.Consume(func(i int, s string) {
		// Count the number of values consumed
		tt.StackLogf("consumer(%d, %v)", i, s)

		// Register consume
		tt.RegisterCalls(1, "consumes")

		// Append flag to assert synchronous flow
		tt.Flag(true, "synchronous")
	})

	// Assert synchronous flow
	tt.AssertFlag(true, "synchronous")

	// Assert no error
	if err != nil {
		t.Errorf("no error was expected while consuming, got %s", err)
	}

	// Expect less than all values
	consumes := tt.CallsTo("consumes")
	tt.Condition(consumes < len(bigSupply), "expected to consume less than %d values, got %d", len(bigSupply), consumes)

	tt.PrintLogStack()
}

// TestCancellableConsumer_FullConsume tests the CancellableConsumer.Consume() method against a HotCloseable supplier.
// It asserts that the consumer consumes all values from the supplier synchronously and concurrently.
// This test also asserts its automatic close after all values have been consumed.
func TestCancellableConsumer_FullConsume(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	stream := streams.HotCloseable(5, bigSupply...)

	// Consume all values from the stream
	err := CancellableConsumer[string](stream, context.Background(), semaph.Channel(3)).
		Consume(func(i int, s string) {
			// Count the number of values consumed
			tt.RegisterCalls(1, "consumes")
			// Log the value consumed
			tt.StackLogf("consumer(%d, %v)", i, s)
		})

	// Assert no error
	if err != nil {
		t.Errorf("no error was expected, got %s", err)
	}

	// Assert all values consumed
	tt.AssertCalls(len(bigSupply), "consumes")

	tt.PrintLogStack()
}

// TestCancellableConsumer_EarlyClosedStream tests the CancellableConsumer.Consume() method against a stream that is closed before it is consumed.
// It asserts that the consumer consumes no values from the supplier and that the supplier is closed.
func TestCancellableConsumer_EarlyClosedStream(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	stream := streams.HotCloseable(7, bigSupply...)
	// Close the stream before consuming
	if err := stream.Close(); err != nil {
		t.Errorf("no error was expected while closing stream, got %s", err)
	}

	// Consume all values from the stream
	err := CancellableConsumer[string](stream, context.Background(), semaph.Channel(3)).
		Consume(func(i int, s string) {
			// Count the number of values consumed
			tt.Errorf("unexpected value consumed, got %d", i)
		})

	// Assert closed streamable error
	if !errors.Is(err, consumers.ClosedStreamable) {
		t.Errorf("expect '%s' error, got '%s'", consumers.ClosedStreamable, err)
	}

	tt.PrintLogStack()
}

// TestCancellableConsumer_OnTheFlyClosedStream tests the CancellableConsumer.Consume() method against a stream that is closed on the fly.
// It asserts that the consumer consumes fewer values than the stream size and that the supplier is closed.
func TestCancellableConsumer_OnTheFlyClosedStream(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	stream := streams.HotCloseable(7, bigSupply...)

	// flag synchronous flow
	tt.Flag(false, "synchronous")

	// close stream on the fly
	tt.RegisterCallback(func() {
		err := stream.Close()
		if err != nil {
			t.Errorf("no error was expected while closing stream, got %s", err)
		}
	}, 5, "consumes")

	// Consume all values from the stream
	err := CancellableConsumer[string](stream, context.Background(), semaph.Channel(3)).
		Consume(func(i int, s string) {
			// Log the value consumed
			tt.StackLogf("consumer(%d, %v)", i, s)

			// Register consume
			tt.RegisterCalls(1, "consumes")

			// Append flag to assert synchronous flow
			tt.Flag(true, "synchronous")
		})

	// Assert synchronous flow
	tt.AssertFlag(true, "synchronous")

	// Assert no error
	if err != nil {
		t.Errorf("no error was expected, got %s", err)
	}

	// Assert fewer values consumed
	tt.Condition(tt.CallsTo("consumes") < len(bigSupply), "expected to consume less than %d values, got %d", len(bigSupply), tt.CallsTo("consumes"))

	tt.PrintLogStack()
}

// TestCancellableConsumer_ClosedConsumer tests that calling Consume on a closed consumer
// doesn't consume any values and returns an error.
func TestCancellableConsumer_ClosedConsumer(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	// Create a stream
	values := []string{"a", "b", "c", "d", "e", "A", "C", "T", "D", "E", "f", "g", "h", "i", "j", "F", "G", "H", "I",
		"J", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	stream := streams.HotCloseable(7, values...)

	// Create a cancellable consumer
	consumer := CancellableConsumer[string](stream, context.Background(), semaph.Channel(3))

	// Close the consumer
	if err := consumer.Close(); err != nil {
		t.Errorf("no error was expected while closing consumer, got %s", err)
	}

	// Try to consume all values from the stream
	err := consumer.Consume(func(i int, s string) {
		// Log the value consumed
		tt.StackLogf("consumer(%d, %v)", i, s)
		// Count the number of values consumed
		tt.RegisterCalls(1, "consumes")
	})

	// Assert an error
	if err == nil {
		t.Errorf("a 'consumer is closed...' error was expected, got nil")
	}

	// Assert no values consumed
	if tt.CallsTo("consumes") != 0 {
		tt.Errorf("expected no values to be consumed, got %d", tt.CallsTo("consumes"))
	}

	tt.PrintLogStack()
}

func TestCancellableConsumer_OnTheFlyClosedConsumer(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	stream := streams.HotCloseable(7, bigSupply...)

	// Create a cancellable consumer
	consumer := CancellableConsumer[string](stream, context.Background(), semaph.Channel(3))

	// Close the consumer on the fly
	tt.RegisterCallback(func() {
		err := consumer.Close()
		if err != nil {
			t.Errorf("no error was expected while closing consumer, got %s", err)
		}
	}, 4, "consumes")

	// Flag synchronous flow
	tt.Flag(false, "synchronous")

	// Try to consume all values from the stream
	err := consumer.Consume(func(i int, s string) {
		// Log the value consumed
		tt.StackLogf("consumer(%d, %v)", i, s)

		// Flag synchronous flow
		tt.Flag(true, "synchronous")

		// Register consume
		tt.RegisterCalls(1, "consumes")
	})

	// Assert synchronous flow
	tt.AssertFlag(true, "synchronous")

	// Assert no error
	if err != nil {
		t.Errorf("no error was expected while consuming stream, got %s", err)
	}

	// Assert fewer values consumed
	tt.Condition(tt.CallsTo("consumes") < len(bigSupply), "expected to consume less than %d values, got %d", len(bigSupply), tt.CallsTo("consumes"))

	tt.PrintLogStack()
}

// TestCancellableConsumer_OnTheFlyCancelledConsumer tests the CancellableConsumer with a stream that is closed on the fly.
// It creates a cancellable context and a cancellable consumer. It then consumes all values from the stream and asserts that no values are consumed if the consumer is closed on the fly.
func TestCancellableConsumer_OnTheFlyCancelledConsumer(t *testing.T) {
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	stream := streams.HotCloseable(7, mediumSupply...)

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Close the consumer on the fly
	tt.RegisterCallback(functions.Runnable(cancel), 4, "consumes")

	// Flag synchronous flow
	tt.Flag(false, "synchronous")

	// Try to consume all values from the stream
	err := CancellableConsumer[string](stream, ctx, semaph.Channel(3)).
		Consume(func(i int, s string) {
			// Log the value consumed
			tt.StackLogf("consumer(%d, %v)", i, s)

			// Flag synchronous flow
			tt.Flag(true, "synchronous")

			// Register consume
			tt.RegisterCalls(1, "consumes")
		})

	// Assert synchronous flow
	tt.AssertFlag(true, "synchronous")

	// Assert no error
	if err != nil {
		t.Errorf("no error was expected while consuming stream, got %s", err)
	}

	// Assert fewer values consumed
	tt.Condition(len(mediumSupply) > tt.CallsTo("consumes"), "expected fewer values than the stream size '%d' to be consumed, got %d", len(mediumSupply), tt.CallsTo("consumes"))

	tt.PrintLogStack()
}
