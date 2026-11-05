// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package producer

import (
	"errors"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

var (
	ErrClosed = errors.New("producer is closed")
)

type Sync[T any] interface {
	// Produce takes a value of type T and sends it to the channel.
	Produce(...T) error
	// Stream returns a stream containing the produced data
	Stream() stream.Readable[T]
}

type Async[T any] interface {
	// Produce accepts a supplier function to be called asynchronously
	// to provide data
	Produce(supplier function.Supplier[[]T]) error
	// Stream returns a stream containing the produced data
	Stream() stream.Readable[T]
}

type Batched[S ~[]E, E any] interface {
	// Produce accepts a batched supplier function to be called asynchronously
	// to provide data in batches
	Produce(supplier function.BatchSupplier[S, E]) error
	// Stream returns a stream containing the produced data
	Stream() stream.Readable[E]
}

type Closeable[T any] interface {
	IsClosed() bool
	Close() error
}

type SyncCloseable[T any] interface {
	Sync[T]
	Closeable[T]
}

type AsyncCloseable[T any] interface {
	Async[T]
	Closeable[T]
}

type BatchedCloseable[S ~[]E, E any] interface {
	Batched[S, E]
	Closeable[S]
}
