// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
)

// Supply runs a supplier asynchronously and returns a channel to receive its values.
// It takes a functions.Supplier[T] function, runs it asynchronously and returns a channel
// of type T that can be used to receive the values returned by the supplier.
// The channel is closed after the supplier finishes.
func Supply[T any](supplier functions.Supplier[T]) <-chan T {
	// Create a channel to hold the result
	result := make(chan T, 1)

	// Before the supplier asynchronously
	go func() {
		// Fetch values (a slice of type T)
		values := supplier()
		// Send values to the result channel
		result <- values
		// Close the result channel
		close(result)
	}()

	return result
}
