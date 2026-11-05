// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package consumer

import (
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type Typed[T any] interface {
	// Consume consumes values from the producer using the provided function.
	Consume(function.BiConsumer[int, T]) error
}

type Closeable[T any] interface {
	IsClosed() bool
	Close() error
	Typed[T]
}
