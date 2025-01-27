// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtools

import (
	"github.com/andrerrcosta2/gtools/core/domain/data"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
)

type AsyncSupplier[T any] interface {
	Streamable[T]
	// Supply receives values of type T from the given channel and sends them
	// to the internal channel returned by Stream.
	//
	// It's thread-safe and can be used concurrently.
	Supply(supplier functions.Supplier[[]T]) error
	// Get returns all the values that have been sent to the supplier.
	//
	// It's thread-safe and can be used concurrently.
	Get() []T
}

type AsyncCloseableSupplier[T any] interface {
	data.Closeable
	AsyncSupplier[T]
}
