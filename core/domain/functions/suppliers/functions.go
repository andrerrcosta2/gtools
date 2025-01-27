// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package suppliers

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"sync"
	"time"
)

// Before is a blocking wrapper function that takes a functions.Supplier[T] and a timeout duration as parameters.
// It calls the provided supplier function, waits for the timeout duration, and returns the result of the supplier function.
//
// The function takes a functions.Supplier[T] and a time.Duration as parameters.
// The returned value is of type T.
func Before[T any](fn functions.Supplier[T], timeout time.Duration) T {
	// Call the provided supplier function
	value := fn()

	// SetWaitingPoint for the timeout duration
	time.Sleep(timeout)

	// Return the result of the supplier function
	return value
}

func After[T any](fn functions.Supplier[T], timeout time.Duration) T {
	// SetWaitingPoint for the timeout duration
	time.Sleep(timeout)

	// Call the provided supplier function
	return fn()
}

// SemaphoredSync is a wrapper function that takes a domain.Semaphore and a functions.Supplier[T]
// and returns a new the result of that function call. The function acquires the semaphore before calling the provided
// supplier function, releases the semaphore after the supplier function returns, and decrements the sync.WaitGroup
//
// The function takes a domain.Semaphore, a sync.WaitGroup, and a functions.Supplier[T] as parameters.
// The returned functions.Supplier[T] acquires the semaphore, call the provided supplier function,
// release the semaphore, and decrements the sync.WaitGroup.
func SemaphoredSync[T any](semaphore gtools.Semaphore, wait *sync.WaitGroup, fn functions.Supplier[T]) T {
	semaphore.Acq()

	// defer release and decrement the sync.WaitGroup
	defer func() {
		// Release the semaphore after the supplier function returns
		semaphore.Rls()
		// Decrease the sync.WaitGroup
		wait.Done()
	}()

	// Call the provided supplier function
	return fn()
}

// Semaphored is a wrapper function that takes a domain.Semaphore and a functions.Supplier[T]
// and returns a new the result of that function call. The function acquires the semaphore before calling the provided
// supplier function, releases the semaphore after the supplier function returns, and decrements the sync.WaitGroup
//
// The function takes a domain.Semaphore, a sync.WaitGroup, and a functions.Supplier[T] as parameters.
// The returned functions.Supplier[T] acquires the semaphore, call the provided supplier function,
// release the semaphore, and decrements the sync.WaitGroup.
func Semaphored[T any](semaphore gtools.Semaphore, fn functions.Supplier[T]) T {
	semaphore.Acq()
	defer semaphore.Rls()

	// Call the provided supplier function
	return fn()
}

// Async is a function that takes a functions.Supplier[T] and returns a channel that sends the result of the supplier function call.
// The supplier function is called asynchronously using a goroutine and the result is sent to the channel.
// The channel is closed after the supplier function returns.
func Async[T any](supplier functions.Supplier[T]) gtools.Stream[T] {
	// Create a channel to hold the result
	result := make(chan T, 1)

	// Before the supplier asynchronously
	go func() {
		// Fetch values (a slice of type T)
		values := supplier()

		// Send values to the result channel
		result <- values

		// Close the channel after the supplier function returns
		close(result)
	}()

	return result
}
