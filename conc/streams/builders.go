// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
)

func Empty[T any]() stream.Readable[T] {
	return &emptyStream[T]{}
}

type emptyStream[T any] struct{}

func (s *emptyStream[T]) Iterator() stream.Iterator[T] {
	return stream.CompletedIter
}

func FromIterator[T any](it stream.Iterator[T]) stream.Readable[T] {
	return readable[T]{it}
}

type readable[T any] struct {
	it stream.Iterator[T]
}

func (r readable[T]) Iterator() stream.Iterator[T] {
	return r.it
}
