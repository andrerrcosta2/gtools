// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package runnables

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/io"
	"sync"
	"time"
)

// Void is a no-op function that takes no arguments and returns no values.
// It can be used as a placeholder or default value for functions.
func Void() {}

// Before calls the provided runnable then waits for the timeout duration.
func Before(fn functions.Runnable, timeout time.Duration) {
	// Call the provided runnable function
	fn()

	// SetWaitingPoint for the timeout duration
	time.Sleep(timeout)
}

// After calls the provided runnable after the timeout duration.
func After(fn functions.Runnable, timeout time.Duration) {
	// SetWaitingPoint for the timeout duration
	time.Sleep(timeout)

	// Call the provided runnable function
	fn()
}

// ContextRelease is a wrapper function that takes a context.Context and a functions.Runnable as parameters.
// It calls the provided runnable function and waits for the context to be canceled.
func ContextRelease(fn functions.Runnable, ctx context.Context) {
	// Call the provided runnable function
	fn()

	// SetWaitingPoint for the context to be canceled
	<-ctx.Done()
}

// SemaphoredSync is a wrapper function that takes a sync.WaitGroup, a domain.Semaphore, and a functions.Runnable
// and calls the provided runnable function. The function acquires the semaphore before calling the provided runnable function,
// releases the semaphore after the runnable function returns, and decreases the sync.WaitGroup
//
// The function takes a sync.WaitGroup, a domain.Semaphore, and a functions.Runnable as parameters.
// The returned functions.Runnable will acquire the semaphore, call the provided runnable function,
// release the semaphore, then decreases the sync.WaitGroup
func SemaphoredSync(wait *sync.WaitGroup, semaphore gtools.Semaphore, fn functions.Runnable) {
	semaphore.Acq()

	// defer release and decrement the sync.WaitGroup
	defer func() {
		// Release the semaphore after the runnable function returns
		semaphore.Rls()
		// Decrease the sync.WaitGroup
		wait.Done()
	}()

	// Call the provided runnable function
	fn()
}

// Semaphored is a wrapper function that takes a domain.Semaphore and a functions.Runnable as parameters.
// It acquires the semaphore before calling the provided runnable function, releases the semaphore after the runnable function returns,
// and does not wait for the runnable function to finish.
//
// The function takes a domain.Semaphore, and a functions.Runnable as parameters.
// The returned functions.Runnable will acquire the semaphore, call the provided runnable function,
// release the semaphore, then return without waiting.
func Semaphored(semaphore gtools.Semaphore, fn functions.Runnable) {
	// Acquire the semaphore before calling the provided runnable function
	semaphore.Acq()
	defer semaphore.Rls()

	// Call the provided runnable function
	fn()
}

// CloseableSelect is a helper function that takes a io.Closeable, a channel as signal, and a functions.Runnable as parameters.
// It checks if the signal is sent, if so, it calls the closeable and returns false.
// Otherwise, it runs the supplier function and returns true.
func CloseableSelect(closeable io.Closeable, done <-chan struct{}, fn functions.Runnable) {
	select {
	case <-done:
		// If the signal is sent, close the closeable
		io.CloseOrLog(closeable, "CloseableSelect")
	default:
		// Otherwise, run the function
		fn()
	}
}
