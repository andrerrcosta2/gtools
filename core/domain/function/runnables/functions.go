// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package runnables

import (
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// Void is a no-op function that takes no arguments and returns no values.
// It can be used as a placeholder or default value for function.
func Void() {}

// Before calls the provided runnable then waits for the timeout duration.
func Before(fn function.Runnable, timeout time.Duration) {
	defer time.Sleep(timeout)
	fn()
}

// After calls the provided runnable after the timeout duration.
func After(fn function.Runnable, timeout time.Duration) {
	time.Sleep(timeout)
	fn()
}

func AfterSignal(fn function.Runnable, blocker interface{ Done() <-chan struct{} }) {
	<-blocker.Done()
	fn()
}

// ContextRelease is a wrapper function that takes a syncs.Blocker and a function.Runnable as parameters.
// It calls the provided runnable function and waits for the context to be canceled.
func ContextRelease(fn function.Runnable, blocker interface{ Done() <-chan struct{} }) {
	fn()
	<-blocker.Done()
}

// Notifier is a runnable that signals when it is over
func Notifier(signal interface{ Signal() }, fn function.Runnable) {
	defer signal.Signal()
	fn()
}

// AcqRlsAndSignal runs a function and release a syncs.Semaphore and a syncs.Completer after the
// runnable is done.
func AcqRlsAndSignal(completer interface{ Done() }, rls interface {
	Acq()
	Rls()
}, fn function.Runnable) {
	rls.Acq()
	defer func() {
		rls.Rls()
		completer.Done()
	}()
	fn()
}

// ReleaseAndSignal runs a function and release a syncs.Semaphore and a syncs.Completer after the
// runnable is done.
func ReleaseAndSignal(completer interface{ Done() }, rls interface{ Rls() }, fn function.Runnable) {
	defer func() {
		rls.Rls()
		completer.Done()
	}()
	fn()
}

// AcqRelease is a wrapper function that takes a domain.Semaphore and a function.Runnable as parameters.
// It acquires the semaphore before calling the provided runnable function, releases the semaphore after the runnable function returns,
// and does not wait for the runnable function to finish.
//
// The function takes a domain.Semaphore, and a function.Runnable as parameters.
// The returned function.Runnable will acquire the semaphore, call the provided runnable function,
// release the semaphore, then return without waiting.
func AcqRelease(rls interface {
	Acq()
	Rls()
}, fn function.Runnable) {
	rls.Acq()
	defer rls.Rls()
	fn()
}
