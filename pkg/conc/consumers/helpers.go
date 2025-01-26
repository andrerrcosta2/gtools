// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package consumers

import (
	"errors"
	"github.com/andrerrcosta2/gtools/conc/channels"
	"github.com/andrerrcosta2/gtools/conc/selects"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/functions/runnables"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/io"
	"log"
	"sync"
)

var ClosedStreamable = errors.New("streamable is closed...\n")
var ClosedStreamableOnTheFly = errors.New("streamable was found unexpectedly closed on the fly while consume...\n")
var ClosedConsumer = errors.New("consumer is closed...\n")
var ClosedConsumerOnTheFly = errors.New("consumer was found unexpectedly closed on the fly while consume...\n")
var Closed = errors.New("consumer is already closed...\n")

// CloseableSelect is a helper function to be used with goroutines.
// It receives a domain.Streamable, a domain.CloseableConsumer, a WaitGroup, a Semaphore, a channel as signal, and a functions.BiConsumer[int, T] function.
// It checks if the signal is sent, if so it call the closeable and returns false.
// Otherwise, it tries to consume the value from the channel using the BiConsumer function
func CloseableSelect[T any](closeable io.Closeable, stream gtools.Stream[T], signal channels.Signal, wait *sync.WaitGroup,
	semaphore gtools.Semaphore, fn functions.BiConsumer[int, T], counter int) bool {
	// Return false on signal sent
	return selects.Bool(signal, func() bool {
		// Otherwise consume the value from the channel using the BiConsumer function
		return ConsumeCloseableSynchronously(closeable, stream, wait, semaphore, fn, counter)
	})
}

// ConsumeCloseableSynchronously is a helper function that is used to consume values from a channel using a BiConsumer function.
// It takes a WaitGroup, a Semaphore, a BiConsumer function, the index of the value to be consumed, and the value itself.
// It checks if the consumer is closed, if so it waits for all operations to complete before returning false.
// Otherwise, it tries to consume the value from the channel using the BiConsumer function and returns true.
// It releases the Semaphore and decrements the WaitGroup.
func ConsumeCloseableSynchronously[T any](closeable io.Closeable, stream gtools.Stream[T], wait *sync.WaitGroup,
	semaphore gtools.Semaphore, fn functions.BiConsumer[int, T], counter int) bool {
	// If the consumer is closed, wait for all operations to complete before returning false
	if closeable.IsClosed() {
		log.Printf("\nconsumer is closed, exiting...\n")
		wait.Wait()
		return false
	}
	// Try consuming the value from the channel using the BiConsumer function
	return ConsumeSynchronously[T](stream, wait, semaphore, counter, fn)
}

// ConsumeSynchronously is a helper function that is used to consume values from a channel using a BiConsumer function.
// It takes a WaitGroup, a Semaphore, a BiConsumer function, the index of the value to be consumed, and the value itself.
// It acquires the Semaphore, consumes the value using the BiConsumer function, releases the Semaphore, and decrements the WaitGroup.
// It returns true if a value was consumed, false otherwise.
func ConsumeSynchronously[T any](stream gtools.Stream[T], wait *sync.WaitGroup, semaphore gtools.Semaphore, counter int, consume functions.BiConsumer[int, T]) bool {
	value, ok := <-stream
	if !ok {
		// If the channel is closed, wait for all operations to complete before returning
		wait.Wait()
		return false
	}

	// Increment the WaitGroup to indicate that a new operation is starting
	wait.Add(1)

	// Consume the value using the provided function
	go runnables.SemaphoredSync(wait, semaphore, func() { consume(counter, value) })

	// Signalize that the value was consumed
	return true
}
