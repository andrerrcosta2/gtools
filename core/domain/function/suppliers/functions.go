// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package suppliers

import (
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// Before is a blocking wrapper function that takes a function.Supplier[T] and a timeout duration as parameters.
// It calls the provided supplier function, waits for the timeout duration, and returns the result of the supplier function.
//
// The function takes a function.Supplier[T] and a time.Duration as parameters.
// The returned value is of type T.
func Before[T any](fn function.Supplier[T], timeout time.Duration) T {
	// Call the provided supplier function
	value := fn()

	// SetWaitingPoint for the timeout duration
	time.Sleep(timeout)

	// Return the result of the supplier function
	return value
}

func After[T any](fn function.Supplier[T], timeout time.Duration) T {
	// SetWaitingPoint for the timeout duration
	time.Sleep(timeout)

	// Call the provided supplier function
	return fn()
}

// AcqRlsAndSignal is a wrapper function that takes a domain.Semaphore and a function.Supplier[T]
// and returns a new the result of that function call. The function acquires the semaphore before calling the provided
// supplier function, releases the semaphore after the supplier function returns, and decrements the sync.WaitGroup
//
// The function takes a domain.Semaphore, a sync.WaitGroup, and a function.Supplier[T] as parameters.
// The returned function.Supplier[T] acquires the semaphore, call the provided supplier function,
// release the semaphore, and decrements the sync.WaitGroup.
func AcqRlsAndSignal[T any](ar interface {
	Acq()
	Rls()
}, completer interface{ Done() }, fn function.Supplier[T]) T {
	ar.Acq()

	defer func() {
		ar.Rls()
		completer.Done()
	}()

	return fn()
}

// AcqRelease is a wrapper function that takes a domain.Semaphore and a function.Supplier[T]
// and returns a new the result of that function call. The function acquires the semaphore before calling the provided
// supplier function, releases the semaphore after the supplier function returns, and decrements the sync.WaitGroup
//
// The function takes a domain.Semaphore, a sync.WaitGroup, and a function.Supplier[T] as parameters.
// The returned function.Supplier[T] acquires the semaphore, call the provided supplier function,
// release the semaphore, and decrements the sync.WaitGroup.
func AcqRelease[T any](ar interface {
	Acq()
	Rls()
}, fn function.Supplier[T]) T {
	ar.Acq()
	defer ar.Rls()

	// Call the provided supplier function
	return fn()
}

// Async is a function that takes a function.Supplier[T] and returns a channel that sends the result of the supplier function call.
// The supplier function is called asynchronously using a goroutine and the result is sent to the channel.
// The channel is closed after the supplier function returns.
func Async[T any](supplier function.Supplier[T]) chan T {
	// size minumum to avoid the goroutine to be blocked
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
