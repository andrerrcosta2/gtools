// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"sync"

	"github.com/andrerrcosta2/gtools/gflux/obs"
)

// Buffer collects values emitted by the source Observable into fixed-size slices
// and emits each slice downstream when it reaches the configured size.
//
// Buffer preserves the order of elements and does not overlap buffers.
// Emission occurs only when the buffer is full, except on completion.
//
// At the source Observable completes, any remaining buffered values
// (if the buffer is not empty) are emitted as a final slice before
// the completion signal is forwarded.
//
// Errors from the source are propagated immediately and discard
// any buffered values.
//
// Buffer is a stateful transformation and is safe to use with
// both hot and cold Observables.
//
// Diagram (size = 3):
//
//	source:  --1--2--3--4--5--|
//	buffer:  --------[1 2 3]----[4 5]|
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	buffered := Buffer(src)
//
//	buffered.Sub(obs.NewObserver(
//		func(v []int) {
//			fmt.Println(v)
//		},
//		func(err error) {},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
//	src.Next(1)
//	src.Next(2)
//	src.Next(3)
//	src.Next(4)
//	src.Next(5)
//	src.Complete()
//
// Output:
//
//	[1 2 3]
//	[4 5]
//	done
//
// Panics if size <= 0.
func Buffer[T any](size int) obs.TransformPipe[T, []T] {
	if size <= 0 {
		panic("buffer size must be > 0")
	}
	return func(src obs.Observable[T]) obs.Observable[[]T] {
		return obs.Create[[]T](func(o *obs.Observer[[]T]) (obs.Subscription, error) {
			var (
				mu  sync.Mutex
				buf = make([]T, 0, size)
			)

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					mu.Lock()
					defer mu.Unlock()

					buf = append(buf, v)
					if len(buf) == size {
						out := buf
						buf = make([]T, 0, size)
						o.Next(out)
					}
				},

				Error: func(err error) {
					mu.Lock()
					defer mu.Unlock()
					o.Error(err)
				},

				Complete: func() {
					mu.Lock()
					defer mu.Unlock()

					if len(buf) > 0 {
						o.Next(buf)
					}
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

// Pairwise emits the previous and current values from the source Observable
// as a pair for each emission after the first.
//
// The first value from the source is stored but not emitted.
// Range subsequent value produces a pair containing:
//
//	(previous, current)
//
// Pairwise preserves ordering and uses constant memory.
//
// Completion:
//   - Completion from the source is forwarded normally.
//   - No additional pairs are emitted on completion.
//
// Error handling:
//   - Errors from the source are propagated immediately.
//
// Diagram:
//
//	source:    --1--2--3--4--|
//	pairwise:    --(1,2)--(2,3)--(3,4)--|
//
// Example:
//
//	src := obs.Of(1, 2, 3, 4)
//
//	paired := Pairwise[int]()(src)
//
//	paired.Sub(obs.NewObserver(
//		func(p [2]int) {
//			fmt.Println(p[0], p[1])
//		},
//		nil,
//		func() { fmt.Println("done") },
//	))
//
// Output:
//
//	1 2
//	2 3
//	3 4
//	done
func Pairwise[T any]() obs.TransformPipe[T, [2]T] {
	return func(src obs.Observable[T]) obs.Observable[[2]T] {
		return obs.Create[[2]T](func(o *obs.Observer[[2]T]) (obs.Subscription, error) {
			var (
				prev T
				set  bool
				mu   sync.Mutex
			)

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					mu.Lock()
					defer mu.Unlock()

					if set {
						o.Next([2]T{prev, v})
					}
					prev = v
					set = true
				},

				Error: func(err error) {
					o.Error(err)
				},

				Complete: func() {
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

// Window splits the source Observable into consecutive inner Observables
// ("windows"), each emitting up to a fixed number of values.
//
// Range window is itself an Observable[T]. A new window is emitted
// downstream when the previous one reaches the configured size.
//
// Windows do not overlap and preserve the original order of values.
//
// At the source Observable completes:
//   - The currently open window (if any) is completed.
//   - The outer Observable then completes.
//
// Error handling:
//   - If the source emits an error, the active window (if any) receives
//     the same error.
//   - The error is then propagated downstream and the outer Observable
//     terminates.
//
// Diagram (size = 3):
//
//	source:  --1--2--3--4--5--6--7--|
//
//	window$: ----w1------w2------w3-|
//	w1:      --1--2--3-|
//	w2:              --4--5--6-|
//	w3:                      --7-|
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	windowed := Window(src)
//
//	windowed.Sub(obs.NewObserver(
//		func(win obs.Observable[int]) {
//			win.Sub(obs.NewObserver(
//				func(v int) {
//					fmt.Println("value:", v)
//				},
//				func(err error) {},
//				func() {
//					fmt.Println("window done")
//				},
//			))
//		},
//		func(err error) {
//			fmt.Println("error:", err)
//		},
//		func() {
//			fmt.Println("all done")
//		},
//	))
//
//	src.Next(1)
//	src.Next(2)
//	src.Next(3)
//
//	src.Next(4)
//	src.Next(5)
//	src.Next(6)
//
//	src.Next(7)
//	src.Complete()
//
// Output:
//
//	value: 1
//	value: 2
//	value: 3
//	window done
//	value: 4
//	value: 5
//	value: 6
//	window done
//	value: 7
//	window done
//	all done
//
// Panics if size <= 0.
func Window[T any](size int) obs.TransformPipe[T, obs.Observable[T]] {
	if size <= 0 {
		panic("Window size must be > 0")
	}

	return func(src obs.Observable[T]) obs.Observable[obs.Observable[T]] {
		return obs.Create[obs.Observable[T]](func(o *obs.Observer[obs.Observable[T]]) (obs.Subscription, error) {

			var (
				mu    sync.Mutex
				count int
				win   obs.Subject[T]
			)

			openWindow := func() {
				win = obs.ReplaySubject[T]()
				count = 0
				o.Next(win)
			}

			closeWindow := func() {
				if win != nil {
					win.Complete()
					win = nil
				}
			}

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					mu.Lock()
					defer mu.Unlock()

					if win == nil {
						openWindow()
					}

					win.Next(v)
					count++

					if count == size {
						closeWindow()
					}
				},

				Error: func(err error) {
					mu.Lock()
					defer mu.Unlock()

					if win != nil {
						win.Error(err)
						win = nil
					}
					o.Error(err)
				},

				Complete: func() {
					mu.Lock()
					defer mu.Unlock()

					if win != nil {
						win.Complete()
						win = nil
					}
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
