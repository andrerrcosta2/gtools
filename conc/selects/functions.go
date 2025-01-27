// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package selects

import (
	"github.com/andrerrcosta2/gtools/conc/channels"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/io"
)

// Closeable is a helper function to be used with goroutines.
// It takes an io.Closeable, a channels.Signal, and a functions.Supplier[bool] function.
// It checks if the signal is sent, if so it call the closeable and returns false.
// Otherwise, it runs the supplier function
func Closeable(closeable io.Closeable, done channels.Signal, fn functions.Supplier[bool]) bool {
	select {
	case <-done:
		// if it is signaled, close the closeable
		io.CloseOrLog(closeable, "CloseableSelect")
		return false
	default:
		// run the supplier function
		return fn()
	}
}

// Void is a helper function to be used with goroutines.
// It takes a channels.Signal, and a functions.Supplier[bool] function.
// It checks if the signal is sent, if so it returns.
// Otherwise, it runs the supplier function
func Void(done channels.Signal, fn functions.Runnable) {
	// if the signal is sent, return
	select {
	case <-done:
		return
	default:
		// run the supplier function
		fn()
		return
	}
}

// Bool is a helper function to be used with goroutines.
// It takes a channels.Signal, and a functions.Supplier[bool] function.
// It checks if the signal is sent, if so it returns false.
// Otherwise, it runs the supplier function
func Bool(signal channels.Signal, fn functions.Supplier[bool]) bool {
	// if the signal is sent, return false
	select {
	case <-signal:
		return false
	default:
		// run the supplier function
		return fn()
	}
}
