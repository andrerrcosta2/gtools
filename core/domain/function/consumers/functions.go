// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package consumers

import (
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// Before is a wrapper function that takes a function.Consumer and a timeout duration as parameters.
// It calls the provided consumer function and waits for the timeout duration.
func Before[T any](fn function.Consumer[T], data T, timeout time.Duration) {
	// Call the provided consumer function
	fn(data)

	// SetWaitingPoint for the timeout duration
	time.Sleep(timeout)
}

// After is a wrapper function that takes a function.Consumer and a timeout duration as parameters.
// It calls the provided consumer function and waits for the timeout duration.
func After[T any](fn function.Consumer[T], data T, timeout time.Duration) {
	// SetWaitingPoint for the timeout duration
	time.Sleep(timeout)

	// Call the provided consumer function
	fn(data)
}

// ContextRelease is a wrapper function that takes a syncs.Blocker and a function.Consumer as parameters.
// It calls the provided consumer function and waits for the context to be canceled.
func ContextRelease[T any](blocker interface{ Done() <-chan struct{} }, fn function.Consumer[T], data T) {
	// Call the provided consumer function
	fn(data)

	// SetWaitingPoint for the context to be canceled
	<-blocker.Done()
}

// AcqRlsAndSignal acquires am acquire-releaser before running a ReleaseAndSignal runnable
func AcqRlsAndSignal[T any](completer interface{ Done() }, rls interface {
	Acq()
	Rls()
}, fn function.Consumer[T], data T) {
	rls.Acq()
	defer func() {
		rls.Rls()
		completer.Done()
	}()
	fn(data)
}

// ReleaseAndSignal runs a runnable and release a release, then complete a completer
func ReleaseAndSignal[T any](completer interface{ Done() }, rls interface {
	Acq()
	Rls()
}, fn function.Consumer[T], data T) {
	defer func() {
		rls.Rls()
		completer.Done()
	}()
	fn(data)
}

// Semaphored is a wrapper function that takes a domain.Semaphore and a function.Consumer as parameters.
// It acquires the semaphore before calling the provided consumer function, releases the semaphore after the consumer function returns,
// and does not wait for the consumer function to finish.
//
// The function takes a domain.Semaphore, and a function.Runnable as parameters.
// The returned function.Runnable will acquire the semaphore, call the provided consumer function,
// release the semaphore, then return without waiting.
func Semaphored[T any](rls interface {
	Acq()
	Rls()
}, fn function.Consumer[T], data T) {
	rls.Acq()
	defer rls.Rls()

	fn(data)
}
