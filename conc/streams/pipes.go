// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import "github.com/andrerrcosta2/gtools/core/domain/contracts/stream"

type Sink[T any] interface {
	apply(T) bool
}

type Consumer[T any] interface {
	accept(T) bool
}

type RangeFn[T any] func(e T) bool

func Range[T any](s stream.Readable[T], fn RangeFn[T]) {
	it := s.Iterator()
	it(fn)
}

func Parallel[T any](src stream.Readable[T]) stream.Readable[T] {
	return parallelize(src)
}
