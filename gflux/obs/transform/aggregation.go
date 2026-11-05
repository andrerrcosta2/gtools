// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"github.com/andrerrcosta2/gtools/gflux/obs"
)

// Reduce applies an accumulator function over the source Observable
// and emits only the final accumulated value when the source completes.
//
// Reduce behaves like Scan, but instead of emitting intermediate
// accumulated values, it emits exactly one value on completion.
//
// The accumulator is initialized with the provided seed value.
//
// Completion:
//   - At the source completes, Reduce emits the final accumulated value
//     and then completes.
//   - If the source emits no values, the seed is emitted.
//
// Error handling:
//   - Errors from the source are propagated immediately.
//   - No value is emitted on error.
//
// Reduce is synchronous, deterministic, and does not introduce
// goroutines or buffering.
//
// Diagram (seed = 0, fn = sum):
//
//	source: --1--2--3--|
//	reduce: ---------6|
//
// Example:
//
//	src := obs.Of(1, 2, 3, 4)
//
//	reduced := Reduce(0, func(acc, v int) int {
//		return acc + v
//	})(src)
//
//	reduced.Sub(obs.NewObserver(
//		func(v int) { fmt.Println(v) },
//		nil,
//		func() { fmt.Println("done") },
//	))
//
// Output:
//
//	10
//	done
func Reduce[T, A any](seed A, fn func(A, T) A) obs.TransformPipe[T, A] {
	return func(src obs.Observable[T]) obs.Observable[A] {
		return obs.Create[A](func(o *obs.Observer[A]) (obs.Subscription, error) {

			acc := seed

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					acc = fn(acc, v)
				},

				Error: func(err error) {
					o.Error(err)
				},

				Complete: func() {
					o.Next(acc)
					o.Complete()
				},
			})

			if err != nil {
				return nil, err
			}

			return sub, nil
		})
	}
}

// Scan applies an accumulator function over the source Observable
// and emits each intermediate accumulated value.
//
// Scan behaves like Reduce, but instead of emitting only the final
// accumulated result, it emits the accumulator after every source value.
//
// The accumulator is initialized with the provided seed value.
//
// Completion:
//   - Completion from the source is forwarded immediately.
//   - No additional values are emitted on completion.
//
// Error handling:
//   - Errors from the source are propagated immediately.
//
// Scan is synchronous, deterministic, and does not introduce
// goroutines or buffering.
//
// Diagram (seed = 0, fn = sum):
//
//	source: --1--2--3--4--|
//	scan:   --1--3--6--10-|
//
// Example:
//
//	src := obs.Of(1, 2, 3, 4)
//
//	scanned := Scan(0, func(acc, v int) int {
//		return acc + v
//	})(src)
//
//	scanned.Sub(obs.NewObserver(
//		func(v int) { fmt.Println(v) },
//		nil,
//		func() { fmt.Println("done") },
//	))
//
// Output:
//
//	1
//	3
//	6
//	10
//	done
func Scan[T, A any](seed A, fn func(A, T) A) obs.TransformPipe[T, A] {
	return func(src obs.Observable[T]) obs.Observable[A] {
		return obs.Create[A](func(o *obs.Observer[A]) (obs.Subscription, error) {
			acc := seed
			terminated := false

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if terminated {
						return
					}
					acc = fn(acc, v)
					o.Next(acc)
				},

				Error: func(err error) {
					if terminated {
						return
					}
					terminated = true
					o.Error(err)
				},

				Complete: func() {
					if terminated {
						return
					}
					terminated = true
					o.Complete()
				},
			})

			if err != nil {
				return nil, err
			}

			return sub, nil
		})
	}
}
