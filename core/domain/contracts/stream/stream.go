// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package stream

import (
	"errors"
)

var ErrClosed = errors.New("this stream is already closed")

// A Readable represents a consumable sequence.
// Implementations may or may not preserve encounter order;
// ordered streams must document that guarantee
type Readable[T any] interface {
	Iterator() Iterator[T]
}

type Parallel[T any] interface {
	Iterator() Iterator[T]
	Unwrap() Readable[T]
}

type Buffered[T any] interface {
	Readable[T]
	Len() int
	Cap() int
}

type Closeable[T any] interface {
	Readable[T]
	Close() error
}

type Iterator[T any] func(fn func(e T) bool) bool

func CompletedIter[T any](_ func(e T) bool) bool { return false }
