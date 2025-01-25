// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streamables

import (
	"github.com/andrerrcosta2/gtools/conc/channels"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/funcs/suppliers"
	"github.com/andrerrcosta2/gtools/core/io"
)

// OpenCloseableSync is a helper function that opens a domain.CloseableStreamable[T] synchronously
// and fills it with the given values. It uses a goroutine to send all values
// to the channel and then closes the channel. If a signal is sent to the
// signal channel, it stops sending values and closes the channel.
func OpenCloseableSync[T any](values []T, channel chan T, closeSignal channels.Signal) {
	go func() {
		defer close(channel)
		for _, value := range values {
			select {
			case channel <- value:
				// Value successfully sent to the channel
			case <-closeSignal:
				// Stop sending if an early close signal is received
				return
			}
		}
	}()
}

// OpenCloseableAsync is a helper function that opens a domain.CloseableStreamable[T] asynchronously
// and fills it with the given values. It uses a goroutine to send all values
// to the channel and then closes the channel. If a signal is sent to the
// signal channel, it stops sending values and closes the channel.
func OpenCloseableAsync[T any](channel chan T, signal channels.Signal, supplier functions.Supplier[[]T]) {
	// Use the supplier to generate values asynchronously
	asyncOutput := suppliers.Async(supplier)

	// Start a goroutine that sends the values to the channel
	go func() {
		defer close(channel)
		// Iterate over the results
		for values := range asyncOutput {
			// Iterate over the values
			for _, value := range values {
				select {
				case channel <- value:
					// Value successfully sent to the channel
				case <-signal:
					// Stop sending if an early close signal is received
					return
				}
			}
		}
	}()
}

// StreamCloseable is a helper function that sends a value to a channel if the supplier isn't closed.
// It takes a io.Closeable, a channel and a value.
// It checks if the supplier is closed, and if so, it returns false.
// If the supplier isn't closed, it sends the value to the channel and returns true.
func StreamCloseable[T any](closeable io.Closeable, ch chan T, value T) bool {
	if closeable.IsClosed() {
		return false
	}
	ch <- value
	return true
}
