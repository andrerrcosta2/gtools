// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observables

import (
	"sync"
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/gflux/core/lifecycle/gates"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

func Async[T any](fn func(o *observers.Observer[T])) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			go func(subscribe func(*observers.Observer[T]), o *observers.Observer[T]) {
				subscribe(created)
			}(fn, created)

			return s, nil
		},
	}
}

// AsyncChan is a channel-backed cold-start, single-consumer observable
func AsyncChan[T any](ch <-chan T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			go func() {
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							created.Complete()
							return
						}
						created.Next(v)
					}
				}
			}()

			return s, nil
		},
	}
}

func AsyncDelayedOf[T any](delay IndexedDelayFn[T], v ...T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			go func(o *observers.Observer[T], getDelay IndexedDelayFn[T], v ...T) {
				for i, item := range v {
					time.Sleep(getDelay(i, item))
					created.Next(item)
				}
				created.Complete()
			}(o, delay, v...)

			return s, nil
		},
	}
}

func AsyncIntervalOf[T any](delay time.Duration, v ...T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			go func(o *observers.Observer[T], delay time.Duration, v ...T) {
				time.Sleep(delay)
				for _, item := range v {
					created.Next(item)
					time.Sleep(delay)
				}
				created.Complete()
			}(o, delay, v...)

			return s, nil
		},
	}
}

// AsyncEmpty creates an asynchronous cold Observable that completes immediately.
func AsyncEmpty[T any]() *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			normal := observers.NormalizeComplete(o)
			s := subscr.NewSlotSub(normal, gates.Atomic())

			go func(o *observers.Observer[T]) {
				if s.Closed() {
					return
				}
				normal.Complete()
			}(o)

			return s, nil
		},
	}
}

func AsyncError[T any](err error) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			normal := observers.NormalizeError(o)
			s := subscr.NewSlotSub(normal, gates.Atomic())

			go func(o *observers.Observer[T], err error) {
				if s.Closed() {
					return
				}
				normal.Error(err)
			}(o, err)

			return s, nil
		},
	}
}

func AsyncJust[T any](v T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			go func() {
				if s.Closed() {
					return
				}
				created.Next(v)
				created.Complete()
			}()

			return s, nil
		},
	}
}

func AsyncMulticast[T any](ch <-chan T) *ColdObs[T] {
	var (
		mu   sync.Mutex
		subs = map[*observers.Observer[T]]subscr.Subscription{}
	)

	go func() {
		for v := range ch {
			mu.Lock()
			for o := range subs {
				o.Next(v)
			}
			mu.Unlock()
		}
		mu.Lock()
		for o := range subs {
			o.Complete()
		}
		mu.Unlock()
	}()

	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			mu.Lock()
			subs[created] = s
			mu.Unlock()

			return s, nil
		},
	}
}

func AsyncOf[T any](v ...T) *ColdObs[T] {
	return &ColdObs[T]{
		sub: func(o *observers.Observer[T]) (subscr.Subscription, error) {
			gate := gates.Atomic()
			created := observers.Create(o, gate)
			s := subscr.NewSlotSub(created, gate)

			go func(o *observers.Observer[T], v ...T) {
				for _, item := range v {
					created.Next(item)
				}
				created.Complete()
			}(o, v...)

			return s, nil
		},
	}
}

func AsyncRange[T nums.Ordered](from, to, delta T) *ColdObs[T] {
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

			go func(o *observers.Observer[T], from, to, delta T) {
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
			}(o, from, to, delta)
			return s, nil
		},
	}
}
