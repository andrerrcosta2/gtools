// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"sync"

	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// Concat creates an Observable that subscribes to the given source Observables
// sequentially, one at a time, in the order they are provided.
//
// Example:
//
//	a := obs.SlotSubject[int]()
//	b := obs.SlotSubject[int]()
//
//	c := obs.Concat(a, b)
//
//	c.Sub(&obs.ObserverImpl[int]{
//		Next: func(v int) {
//			fmt.Println(v)
//		},
//		Complete: func() {
//			fmt.Println("done")
//		},
//	})
//
//	a.Next(1)
//	a.Next(2)
//	a.Complete()
//
//	b.Next(3)
//	b.Next(4)
//	b.Complete()
//
//	Output:
//	1
//	2
//	3
//	4
//	done
//
// Diagram:
//
//	a: --1--2--|
//	b: ---------3--4--|
//
//	Concat(a, b):
//	--1--2----3--4--|
//
// Semantics:
//
//   - Only one source Observable is subscribed at any given time.
//   - The next source is subscribed only after the previous source signals Complete.
//   - Values (Next) are forwarded immediately to downstream observers.
//   - Errors (Error) are terminal: on the first error, Concat forwards the error
//     downstream and stops without subscribing to remaining sources.
//   - Completion (Complete) is forwarded only after all source Observables have
//     completed successfully.
//
// Subscription behavior:
//
//   - The Subscription returned by Concat controls the entire concatenation process.
//   - Calling Unsub() cancels the currently active source subscription (if any)
//     and prevents any remaining sources from being subscribed.
//   - Sources that have already completed are not affected by Unsub().
//   - Sources that have not yet been subscribed are never subscribed after Unsub().
//
// Important:
//
//   - Concat does NOT subscribe to all sources eagerly.
//   - Concat does NOT keep multiple active subscriptions.
//   - At most, one upstream Subscription exists at any point in time.
//
// This operator preserves ordering and guarantees that all values from a source
// are emitted before moving to the next source.
//
// Concat is synchronous with respect to the active source and does not introduce
// goroutines by itself.
func Concat[T any](sources ...obs.Observable[T]) obs.Observable[T] {
	return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
		var (
			index int
			sub   obs.Subscription
		)

		cs := subscr.Func(func() {
			if sub != nil {
				sub.Unsub()
			}
		})

		var subscribeNext func() error
		subscribeNext = func() error {
			if cs.Closed() {
				return nil
			}
			if index >= len(sources) {
				o.Complete()
				cs.Close()
				return nil
			}

			src := sources[index]
			index++

			s, err := src.Sub(&obs.Observer[T]{
				Next: o.Next,
				Error: func(err error) {
					o.Error(err)
					cs.Close()
				},
				Complete: func() {
					_ = subscribeNext()
				},
			})

			if err != nil {
				return err
			}

			sub = s
			return nil
		}

		if err := subscribeNext(); err != nil {
			return nil, err
		}

		return cs, nil
	})
}

// Merge creates an Observable that subscribes to all provided sources
// eagerly and forwards values from all active sources to the downstream
// ObserverImpl as they arrive.
//
// Unlike Concat, Merge does not preserve ordering between sources.
// Values are emitted immediately, based solely on when each source produces them.
//
// Completion semantics:
//   - The merged Observable completes only after *all* sources have completed.
//   - Completion order of sources does not matter.
//
// Error semantics:
//   - If any source emits an error, the merged Observable immediately emits that error.
//   - All currently active source subscriptions are unsubscribed (fail-fast behavior).
//
// Subscription semantics:
//   - Sources are subscribed eagerly, in order, when the merged Observable is subscribed.
//   - If a source emits an error synchronously during subscription, remaining sources
//     may never be subscribed.
//   - Unsubscribing from the merged Observable unsubscribes from all active sources.
//
// Concurrency note:
//   - Merge itself does not introduce concurrency.
//   - If sources emit concurrently (e.g. via goroutines), their values may interleave.
//   - With synchronous sources (Of, FromFunc), emissions occur deterministically
//     during subscription.
//
// Example:
//
//	src1: --1-----3----|
//	src2: ----2-----4--|
//
//	Merge(src1, src2):
//	--1--2--3--4-------|
//
// Note:
//
//	Merge is suitable for combining independent streams where ordering
//	is not significant and values should be observed as soon as they are produced.
func Merge[T any](sources ...obs.Observable[T]) obs.Observable[T] {
	return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

		if len(sources) == 0 {
			o.Complete()
			return subscr.Noop{}, nil
		}

		var (
			remaining  = len(sources)
			subs       = make([]obs.Subscription, 0, len(sources))
			terminated bool
			mu         sync.Mutex
		)

		terminateWithError := func(err error) {
			mu.Lock()
			if terminated {
				mu.Unlock()
				return
			}
			terminated = true
			mu.Unlock()

			o.Error(err)
			for _, s := range subs {
				s.Unsub()
			}
		}

		for _, src := range sources {
			s, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					mu.Lock()
					if terminated {
						mu.Unlock()
						return
					}
					mu.Unlock()
					o.Next(v)
				},

				Error: terminateWithError,

				Complete: func() {
					mu.Lock()
					if terminated {
						mu.Unlock()
						return
					}
					remaining--
					if remaining == 0 {
						terminated = true
						mu.Unlock()
						o.Complete()
						return
					}
					mu.Unlock()
				},
			})

			if err != nil {
				return nil, err
			}

			mu.Lock()
			if terminated {
				mu.Unlock()
				s.Unsub()
				break
			}
			subs = append(subs, s)
			mu.Unlock()
		}

		return subscr.Func(func() {
			mu.Lock()
			if terminated {
				mu.Unlock()
				return
			}
			terminated = true
			mu.Unlock()

			for _, s := range subs {
				s.Unsub()
			}
		}), nil
	})
}
