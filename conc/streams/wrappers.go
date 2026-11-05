// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/async"
	"github.com/andrerrcosta2/gtools/conc/core/channel/channels"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/semaph"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
)

func parallelize[T any](src stream.Readable[T], maxConcurrency ...int) stream.Readable[T] {
	var unw stream.Readable[T]
	if wrapped, ok := src.(stream.Parallel[T]); ok {
		unw = wrapped.Unwrap()
	}

	if len(maxConcurrency) == 0 {
		return unboundParallel[T]{src: unw}
	}
	return semParallel[T]{src: unw, sem: semaph.Limiter(maxConcurrency[0])}
}

type unboundParallel[T any] struct {
	src stream.Readable[T]
}

func (s unboundParallel[T]) Iterator() stream.Iterator[T] {
	return func(fn func(T) bool) bool {
		var wg sync.WaitGroup
		flag := channels.Flag()

		s.src.Iterator()(func(e T) bool {
			select {
			case <-flag.Done():
				return false
			default:
			}

			async.CountedConsume(&wg, func(v T) {
				ok := fn(v)
				if !ok {
					flag.Signal()
				}
			}, e)

			return true
		})

		wg.Wait()
		return false
	}
}

func (s unboundParallel[T]) Unwrap() stream.Readable[T] {
	return s.src
}

type semParallel[T any] struct {
	src stream.Readable[T]
	sem syncs.Semaphore
}

func (s semParallel[T]) Iterator() stream.Iterator[T] {
	return func(fn func(T) bool) bool {
		var wg sync.WaitGroup
		flag := channels.Flag()

		s.src.Iterator()(func(e T) bool {
			select {
			case <-flag.Done():
				return false
			default:
			}

			async.SemCountedConsume(&wg, s.sem, func(v T) {
				ok := fn(v)
				if !ok {
					flag.Signal()
				}
			}, e)

			return true
		})

		wg.Wait()
		return false
	}
}

func (s semParallel[T]) Unwrap() stream.Readable[T] {
	return s.src
}
