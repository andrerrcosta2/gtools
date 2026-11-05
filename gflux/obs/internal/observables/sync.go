// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observables

import (
	"time"

	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/gflux/core/lifecycle/gates"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

type (
	Observable[T any] interface {
		Sub(*observers.Observer[T]) (subscr.Subscription, error)
	}

	// ColdObs represents an Observable whose execution starts
	// independently for each subscription.
	//	- New execution per subscription
	//	- No shared states
	//	- No implicit replay
	//	- Deterministic only if the source is deterministic
	ColdObs[T any] struct {
		sub func(o *observers.Observer[T]) (subscr.Subscription, error)
	}

	// ReplayObs represents an Observable whose emissions are
	// recorded and deterministically replayed for each subscriber.
	// Conceptually, ReplayObs is cold in execution, but hot in memory:
	// 	- First execution produces values
	//	- Values are stored
	//	- Every subscriber sees the same sequence
	//	- No concurrency hazards
	//	- Deterministic tests
	//	- Replay defines time travel
	ReplayObs[T any] struct {
		src     Observable[T]
		values  []T
		err     error
		done    bool
		started bool
	}
)

func (o *ColdObs[T]) Sub(obs *observers.Observer[T]) (subscr.Subscription, error) {
	return o.sub(obs)
}

func (r *ReplayObs[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	// First subscriber executes the source
	if !r.started {
		r.started = true

		_, err := r.src.Sub(&observers.Observer[T]{
			Next: func(v T) {
				r.values = append(r.values, v)
			},
			Error: func(err error) {
				r.err = err
				r.done = true
			},
			Complete: func() {
				r.done = true
			},
		})

		if err != nil {
			r.err = err
			r.done = true
		}
	}

	gate := gates.Atomic()
	created := observers.Create(o, gate)
	s := subscr.NewSlotSub(created, gate)

	for _, v := range r.values {
		created.Next(v)
	}

	if r.err != nil {
		created.Error(r.err)
		return s, nil
	}

	if r.done {
		created.Complete()
	}

	return s, nil
}

func Chan[T any](ch <-chan T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)
			for {
				select {
				case v, ok := <-ch:
					if !ok {
						created.Complete()
						return s, nil
					}
					created.Next(v)
				}
			}
		},
	}
}

// Defer creates a cold Observable that defers the creation of the actual
// source Observable until subscription time.
//
// The function fn is invoked separately for each subscription, allowing
// the resulting Observable to capture fresh states or perform side effects
// per subscriber.
//
// Defer does not emit any values by itself; all emissions come from the
// Observable returned by fn.
//
// Cancellation semantics:
//
//	Cancellation is delegated entirely to the Observable returned by fn.
//	Unsubscribing from the resulting subscription unsubscribes from the
//	inner Observable.
//
// Diagram:
//
//	Defer(fn):
//
//	subscribe
//	   |
//	   v
//	fn() -> Observable
//	   |
//	   v
//	(inner emissions)
//
// Example:
//
//	attempts := 0
//
//	src := Defer(func() Observable[int] {
//		attempts++
//		return Of(attempts)
//	})
//
//	src.Sub(...)
//	src.Sub(...)
//
// Output:
//
//	1
//	2
func Defer[T any](fn func() Observable[T]) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			return fn().Sub(o)
		},
	}
}

type IndexedDelayFn[T any] func(index int, value T) time.Duration

func DelayedOf[T any](fn IndexedDelayFn[T], v ...T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)
			for i, item := range v {
				time.Sleep(fn(i, item))
				created.Next(item)
			}
			created.Complete()
			return s, nil
		},
	}
}

func Error[T any](err error) Observable[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)
			created.Error(err)
			return s, nil
		},
	}
}

// Empty creates a cold Observable that completes immediately.
func Empty[T any]() *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			if o.Complete != nil {
				o.Complete()
			}
			return subscr.Noop{}, nil
		},
	}
}

// FromFunc creates a cold Observable from a function that is given direct
// access to an Observer.
//
// The provided function fn is executed synchronously at subscription time
// and is responsible for calling Observer methods (Next, Error).
// Completion is automatically signaled after fn returns.
//
// FromFunc is useful for creating custom synchronous Observables without
// implementing the Observable interface directly.
//
// Cancellation semantics:
//
//	A cancellable Subscription is returned. If Unsub is called, downstream
//	pipes and observers will observe cancellation, but fn itself is not
//	interrupted once execution has begun.
//
// Diagram:
//
//	FromFunc(fn):
//
//	subscribe
//	   |
//	   v
//	fn(o)
//	   |
//	   v
//	o.Complete()
//
// Example:
//
//	src := FromFunc(func(o *Observer[int]) {
//		o.Next(1)
//		o.Next(2)
//	})
//
//	src.Sub(observer.NewObserver(
//		func(v int) { fmt.Println(v) },
//		nil,
//		func() { fmt.Println("done") },
//	))
//
// Output:
//
//	1
//	2
//	done
func FromFunc[T any](fn func(*observers.Observer[T])) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)
			fn(created)
			return s, nil
		},
	}
}

// FromIter creates a cold Observable that synchronously pulls values
// from the provided Iterator until it is exhausted.
//
// FromIter is a synchronous, linear Observable.
// Values are pulled and emitted one-by-one on the caller’s goroutine
// when a subscription occurs.
//
// Iterator semantics:
//
//   - HasNext() is checked before each emission
//   - Next() is called only when HasNext() returns true
//
// Cancellation semantics:
//
//   - If the subscription is canceled, iteration stops immediately
//     and no further calls to HasNext or Next are made.
//
// FromIter does not introduce buffering, concurrency, or replay.
// Range subscription consumes the iterator independently.
func FromIter[T any](it data.Iterator[T]) Observable[T] {
	return Func[T](func(o *observers.Observer[T]) (subscr.Subscription, error) {
		sub := subscr.NewSlotSub(o, gates.Atomic())
		o = observers.Normalize(o)

		for {
			if sub.Closed() {
				return sub, nil
			}

			if !it.HasNext() {
				o.Complete()
				return sub, nil
			}

			v := it.Next()
			o.Next(v)
		}
	})
}

type Func[T any] func(*observers.Observer[T]) (subscr.Subscription, error)

func (f Func[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	return f(observers.Create(o, gates.Atomic()))
}

func IntervalOf[T any](delay time.Duration, v ...T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)
			time.Sleep(delay)
			for _, item := range v {
				created.Next(item)
				time.Sleep(delay)
			}
			created.Complete()
			return s, nil
		},
	}
}

func Just[T any](v T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)
			created.Next(v)
			created.Complete()
			// this is sync, but pipes may alter its behavior
			return s, nil
		},
	}
}

func Of[T any](v ...T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			for _, item := range v {
				created.Next(item)
			}
			created.Complete()

			// this is sync but pipes may alter its behavior
			return s, nil
		},
	}
}

func Range[T nums.Ordered](from, to, delta T) *ColdObs[T] {
	if delta == 0 {
		panic("delta must be non-zero")
	}
	// decreasing ranges
	if delta < 0 && from < to {
		panic("delta must be non-negative when from < to")
	}
	if delta > 0 && from > to {
		panic("delta must be negative when from > to")
	}
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			current := from
			for {
				created.Next(current)
				next := current + delta
				if (delta > 0 && next > to) || (delta < 0 && next < to) {
					break
				}
				if next == current {
					break
				}
				current = next
			}

			created.Complete()
			// this is sync, but pipes may alter its behavior
			return s, nil
		},
	}
}

func Replay[T any](src Observable[T]) *ReplayObs[T] {
	return &ReplayObs[T]{src: src}
}
