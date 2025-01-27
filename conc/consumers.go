// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"context"
	"github.com/andrerrcosta2/gtools/conc/channels"
	"github.com/andrerrcosta2/gtools/conc/consumers"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"log"
	"sync"
	"sync/atomic"
)

// SemaphoredConsumer returns a new domain.CloseableConsumer that consumes values from the provided supplier and applies the given function to each value.
// The returned consumer uses the provided semaphore to limit the number of concurrent operations.
//
// It's a blocking function that waits for all operations to complete before returning.
func SemaphoredConsumer[T any](streamable gtools.CloseableStreamable[T], semaphore gtools.Semaphore) gtools.CloseableConsumer[T] {
	// Create a new channelConsumer with the provided supplier and semaphore
	return &channelConsumer[T]{streamable: streamable, semaphore: semaphore, flag: channels.NewFlag()}
}

// channelConsumer returns a new channelConsumer that consumes values from the provided supplier and applies the given function to each value.
type channelConsumer[T any] struct {
	wait       sync.WaitGroup
	streamable gtools.CloseableStreamable[T]
	semaphore  gtools.Semaphore
	flag       *channels.Flag
	closed     atomic.Bool
}

// Consume consumes values synchronously blocking the main thread until all values have been consumed.
// The consumer applies the given function to each value as it is received.
// The consumer uses the provided semaphore to limit the number of concurrent operations.
// The consumer waits for all operations to complete before returning.
func (c *channelConsumer[T]) Consume(fn functions.BiConsumer[int, T]) error {
	if c.IsClosed() {
		return consumers.ClosedConsumer
	}

	if c.streamable.IsClosed() {
		return consumers.ClosedStreamable
	}

	// Block outer goroutine until all values have been consumed
	c.consume(fn)

	// Return no errors
	return nil
}

func (c *channelConsumer[T]) consume(fn functions.BiConsumer[int, T]) {
	// Capture the stream once outside the loop
	stream := c.streamable.Stream()
	// Use a counter to provide the indexes of values consumed
	counter := 0
	// Loop until all values have been consumed
	for {
		if !consumers.CloseableSelect[T](c, stream, c.flag.Receiver(), &c.wait, c.semaphore, fn, counter) {
			return
		}
		// Increment the counter to provide the indexes of values consumed
		counter++
	}
}

func (c *channelConsumer[T]) Close() error {
	if c.closed.CompareAndSwap(false, true) {
		return nil
	}
	return consumers.Closed
}

func (c *channelConsumer[T]) IsClosed() bool {
	return c.closed.Load()
}

var _ gtools.CloseableConsumer[int] = (*channelConsumer[int])(nil)

// CancellableConsumer creates a new channelConsumer that consumes values from the provided supplier and applies the given function to each value.
// The consumer is cancellable via the context.
// The context is used to stop the goroutine that supplies values to the channel.
//
// Parameters:
// - supplier: The supplier to consume values from.
// - context: The context to be used to cancel the supplier.
//
// Returns:
//   - A new channelConsumer that consumes values from the provided supplier and applies the given function to each value.
func CancellableConsumer[T any](streamable gtools.CloseableStreamable[T], ctx context.Context, semaphore gtools.Semaphore) gtools.CloseableConsumer[T] {
	// Create a new channelConsumer with the provided supplier and semaphore
	return &cancellableCloseableConsumer[T]{streamable: streamable, ctx: ctx, semaphore: semaphore}
}

// cancellableCloseableConsumer is a consumer that uses a channel-based supplier. It consumes values from the provided
// supplier and applies the given function to each value.
type cancellableCloseableConsumer[T any] struct {
	closed     atomic.Bool
	streamable gtools.CloseableStreamable[T]
	semaphore  gtools.Semaphore
	ctx        context.Context
	wait       sync.WaitGroup
}

// Consume consumes values from the supplier using the provided function.
// The consumer is cancellable via the context.
// The context is used to stop the goroutine that supplies values to the channel.
//
// Parameters:
// - fn: The function to be applied to each value consumed.
//
// Returns:
//   - An error if the supplier is closed.
func (c *cancellableCloseableConsumer[T]) Consume(fn functions.BiConsumer[int, T]) error {
	if c.IsClosed() {
		return consumers.ClosedConsumer
	}

	// Exit if streamable is closed
	if c.streamable.IsClosed() {
		return consumers.ClosedStreamable
	}

	// Block outer goroutine until all values have been consumed
	c.consume(fn)

	// Return no errors
	return nil
}

// consume consumes values from the channel using the provided function.
// It uses a waitgroup to wait for all operations to complete.
// It's thread-safe and can be used concurrently.
func (c *cancellableCloseableConsumer[T]) consume(fn functions.BiConsumer[int, T]) {
	// Capture the stream once outside the loop
	stream := c.streamable.Stream()
	// Use a counter to provide the indexes of values consumed
	counter := 0
	for {
		if c.ctx.Err() != nil {
			// If the consumer is force-closed, reject the supplier value and close the channel
			log.Printf("rejecting supplier value and closing channel...\n")
			// SetWaitingPoint for all operations to complete
			c.wait.Wait()
			// Return immediately
			return
		}

		// SetWaitingPoint for a new value to be consumed
		// It's important to use the same waitgroup and context as the consumer
		if !consumers.CloseableSelect[T](c, stream, c.ctx.Done(), &c.wait, c.semaphore, fn, counter) {
			return
		}

		// Increment the counter to provide the indexes of values consumed
		counter++
	}
}

// Close closes the consumer from consuming new values.
// It rejects new calls to the Consume method and closes the consumer.
// It's thread-safe and can be used concurrently.
func (c *cancellableCloseableConsumer[T]) Close() error {
	if c.closed.CompareAndSwap(false, true) {
		return nil
	}
	return consumers.Closed
}

// IsClosed returns true if the consumer has been closed.
func (c *cancellableCloseableConsumer[T]) IsClosed() bool { return c.closed.Load() }

var _ gtools.CloseableConsumer[int] = (*cancellableCloseableConsumer[int])(nil)
