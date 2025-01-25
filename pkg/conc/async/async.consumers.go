// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

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

// Consumer returns a new asynchronous domain.CloseableConsumer that consumes values
// from the provided supplier and applies the given function to each value.
//
// Parameters:
// - supplier: The supplier to consume values from.
//
// Returns:
//   - A new channelConsumer that consumes values from the provided supplier and
//     applies the given function to each value.
func Consumer[T any](streamable gtools.CloseableStreamable[T], sem gtools.Semaphore) gtools.CloseableConsumer[T] {
	return &channelConsumer[T]{streamable: streamable, semaphore: sem, flag: channels.NewFlag()}
}

type channelConsumer[T any] struct {
	closed     atomic.Bool
	flag       *channels.Flag
	semaphore  gtools.Semaphore
	streamable gtools.CloseableStreamable[T]
	wait       sync.WaitGroup
}

// Consume consumes values from the channel using the provided function.
// The provided function takes two parameters: the index of the value and the value itself.
// This method is intended to be used as a goroutine. Due to the nature of an asynchronous consumer, the
// consumption shouldn't block the main goroutine.
//
// Parameters:
// - fn: The function to be applied to each value consumed.
// - sem: The semaphore to be used to limit the number of concurrent operations.
//
// Returns:
//   - An error if the supplier is flag.
func (c *channelConsumer[T]) Consume(fn functions.BiConsumer[int, T]) error {
	if c.IsClosed() {
		return consumers.ClosedConsumer
	}

	// Consume asynchronously
	go c.consume(fn)

	// Return no errors
	return nil
}

// Consume consumes values from the channel using the provided function.
// The provided function takes two parameters: the index of the value and the value itself.
// This method is intended to be used as a goroutine. Due to the nature of an asynchronous consumer, the
// consumption shouldn't block the main goroutine.
func (c *channelConsumer[T]) consume(fn functions.BiConsumer[int, T]) {
	// Capture the stream once outside the loop
	stream := c.streamable.Stream()
	// Use a counter to provide the indexes of values consumed
	counter := 0
	// Use a waitgroup to wait for all operations to complete
	for {
		// TODO: Check that double closing, it is a possible bug
		if c.IsClosed() {
			log.Printf("rejecting supplier values and closing channel...\n")
			c.wait.Wait()
			return
		}
		// Consume the value using the provided function
		consumers.CloseableSelect[T](c, stream, c.flag.Receiver(), &c.wait, c.semaphore, fn, counter)

		// Increment the counter to provide the indexes of values consumed
		counter++
	}
}

// Close closes the consumer and stops it from consuming values from the supplier.
// It rejects new calls to the Consume method and closes the consumer.
// It's thread-safe and can be used concurrently.
func (c *channelConsumer[T]) Close() error {
	if c.closed.CompareAndSwap(false, true) {
		c.flag.Signal()
		return nil
	}
	return consumers.Closed
}

// IsClosed returns true if the consumer has been flag.
// It's thread-safe and can be used concurrently.
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
func CancellableConsumer[T any](streamable gtools.CloseableStreamable[T], context context.Context, semaphore gtools.Semaphore) gtools.CloseableConsumer[T] {
	// Create a new channelConsumer with the provided supplier and semaphore
	return &cancellableChannelConsumer[T]{streamable: streamable, ctx: context, semaphore: semaphore}
}

type cancellableChannelConsumer[T any] struct {
	ctx        context.Context
	closer     atomic.Bool
	streamable gtools.CloseableStreamable[T]
	semaphore  gtools.Semaphore
	wait       sync.WaitGroup
}

func (c *cancellableChannelConsumer[T]) Consume(fn functions.BiConsumer[int, T]) error {
	if c.IsClosed() {
		return consumers.ClosedConsumer
	}

	// Exit if the context is already canceled
	if c.ctx.Err() != nil {
		return c.ctx.Err()
	}

	// Consume asynchronously
	go c.consume(fn)

	// Return no error
	return nil
}

func (c *cancellableChannelConsumer[T]) consume(fn functions.BiConsumer[int, T]) {
	// Capture the stream once outside the loop
	stream := c.streamable.Stream()
	// Use a counter to provide the indexes of values consumed
	counter := 0
	// Use a waitgroup to wait for all operations to complete
	for {
		if c.ctx.Err() != nil {
			log.Printf("rejecting supplier value and closing channel...\n")
			c.wait.Wait()
			return
		}
		// SetWaitingPoint for a new value to be consumed
		consumers.CloseableSelect[T](c, stream, c.ctx.Done(), &c.wait, c.semaphore, fn, counter)
		// Increment the counter to provide the indexes of values consumed
		counter++
	}
}

// Close closes the consumer and stops it from consuming values from the supplier.
// It rejects new calls to the Consume method and closes the consumer.
// It's thread-safe and can be used concurrently.
func (c *cancellableChannelConsumer[T]) Close() error {
	// Check for error
	if err := c.ctx.Err(); err != nil {
		return err
	}
	// Signal the consumer to close
	c.ctx.Done()
	// Return no errors
	return nil
}

func (c *cancellableChannelConsumer[T]) IsClosed() bool {
	return c.closer.Load()
}

var _ gtools.CloseableConsumer[int] = (*cancellableChannelConsumer[int])(nil)
