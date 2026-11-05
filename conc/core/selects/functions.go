// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package selects

import (
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/io"
)

// Closeable is a std function to be used with goroutines.
// It takes an io.Closeable, a channels.Wait, and a function.Supplier[bool] function.
// It checks if the signal is sent, if so it call the closeable and returns false.
// Otherwise, it runs the supplier function
func Closeable(closeable io.Closeable, blocker syncs.Blocker, fn function.Supplier[bool]) bool {
	select {
	case <-blocker.Done():
		// if it is signaled, close the closeable
		io.CloseOrLog(closeable, "CloseableSelect")
		return false
	default:
		// run the supplier function
		return fn()
	}
}

// Void is a std function to be used with goroutines.
// It takes a syncs.Blocker, and a function.Runnable.
// It checks if the signal is sent, if so it returns.
// Otherwise, it runs the runnable function
func Void(blocker syncs.Blocker, fn function.Runnable) {
	// if the signal is sent, return
	select {
	case <-blocker.Done():
		return
	default:
		// run the supplier function
		fn()
		return
	}
}

// Bool is a std function to be used with goroutines.
// It takes a channels.Wait, and a function.Supplier[bool] function.
// It checks if the signal is sent, if so it returns false.
// Otherwise, it runs the supplier function
func Bool(blocker syncs.Blocker, fn function.Supplier[bool]) bool {
	// if the signal is sent, return false
	select {
	case <-blocker.Done():
		return false
	default:
		// run the supplier function
		return fn()
	}
}
