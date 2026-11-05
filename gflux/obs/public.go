package obs

import (
	"sync"
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/gflux/core/lifecycle/gates"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observables"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

type (
	Observable[T any] interface {
		Sub(*Observer[T]) (Subscription, error)
	}
	Observer[T any] observers.Observer[T]
)

func NewObserver[T any](nxt func(T), err func(error), cpt func()) *Observer[T] {
	return (*Observer[T])(observers.New(nxt, err, cpt))
}

func CompleteObserver[T any](cpt func()) *Observer[T] {
	return &Observer[T]{
		Complete: cpt,
	}
}

func ErrorObserver[T any](err func(error)) *Observer[T] {
	return &Observer[T]{
		Error: err,
	}
}

func NextObserver[T any](nxt func(T)) *Observer[T] {
	return &Observer[T]{
		Next: nxt,
	}
}

// Async creates a cold, asynchronous Observable from a user-provided
// subscription function.
//
// Async allows defining an Observable imperatively by providing a
// function that receives an Observer and emits values by calling
// Next, Error, and Complete directly.
//
// Async is a cold Observable:
//
//	The provided function f is not executed until a subscription
//	is created. Range subscription invokes f independently.
//
// Asynchronous semantics:
//
//	The function f is executed in its own goroutine per subscription.
//	Observer callbacks (Next, Error, Complete) are therefore invoked
//	asynchronously with respect to the caller of Sub.
//
// Subscription semantics:
//
//	Range call to Sub:
//	  - Creates a new Observer instance
//	  - Starts a new goroutine executing f
//	  - Returns a Subscription immediately
//
// There is no sharing between subscriptions unless explicitly
// implemented by the user inside f.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	If Unsub is called:
//	  - The Observer becomes closed
//	  - Further calls to Next, Error, or Complete are ignored
//
// Important:
//
//	Async does NOT automatically stop the goroutine executing f.
//	It is the responsibility of f to respect cancellation by
//	observing the Observer’s closed states or coordinating through
//	external signals (channels, contexts, etc.).
//
// Error semantics:
//
//	Async does not perform panic recovery.
//	If f panics, the panic will escape the goroutine and crash
//	the program unless recovered by the user.
//
// Completion semantics:
//
//	The Observable completes only if f explicitly calls Complete
//	or Error on the Observer.
//
// Diagram:
//
//	Async(f):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	--Next--Next--Complete
//
// Example:
//
//	src := Async(func(o *Observer[int]) {
//		for i := 1; i <= 3; i++ {
//			o.Next(i)
//		}
//		o.Complete()
//	})
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output:
//
//	1
//	2
//	3
//	done
//
// Notes:
//
//	Async is the lowest-level constructor for creating Observables.
//	It provides maximal flexibility but minimal safety.
//
//	It is best used for:
//	  - Integrating callback-based or goroutine-based APIs
//	  - Prototyping custom sources
//	  - Building higher-level constructors
//
//	For channel-based sources, prefer AsyncChan.
//	For synchronous generation, prefer Sync or Defer.
func Async[T any](f func(*Observer[T])) Observable[T] {
	inner := func(o *observers.Observer[T]) {
		f((*Observer[T])(o))
	}
	return observableAdapter[T]{inner: observables.Async[T](inner)}
}

// AsyncChan creates a channel-backed Observable that asynchronously
// emits values received from the provided channel.
//
// AsyncChan is a cold, asynchronous Observable with single-consumer
// semantics. Values are only read from the channel after a subscription
// is created, and each subscription starts its own goroutine that
// competes to read from the same channel.
//
// Important: AsyncChan is NOT multicast.
//
// If multiple observers subscribe to the same AsyncChan, values sent
// on the underlying channel will be received by only one subscriber,
// according to Go channel scheduling rules. No guarantees are made
// about fairness or distribution.
//
// Cold semantics:
//
//	The underlying channel is not read until Sub is called.
//	Values sent to the channel before subscription are not observed.
//	If no subscription exists, sending to an unbuffered channel
//	will block.
//
// Asynchronous semantics:
//
//	Values are emitted from a dedicated goroutine per subscription.
//	Observer callbacks (Next, Complete) are invoked asynchronously.
//
// Completion semantics:
//
//   - At the channel is closed, Complete is signaled downstream.
//   - No Error is produced by channel closure.
//   - After completion, no further values are emitted.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation.
//	If Unsub is called, the observer stops receiving values.
//	The goroutine will exit once the channel is closed or the
//	subscription is closed.
//
// Diagram:
//
//	ch := make(chan int)
//
//	AsyncChan(ch):
//
//	subscribe
//	   |
//	   v
//	(ch goroutine)
//	   |
//	   v
//	--1--2--3--|
//
// Example:
//
//	ch := make(chan int)
//
//	src := AsyncChan(ch)
//
//	go func() {
//		ch <- 1
//		ch <- 2
//		ch <- 3
//		close(ch)
//	}()
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v) },
//		Error: func(err error) {},
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output (order preserved):
//
//	1
//	2
//	3
//	done
//
// Notes:
//
//	AsyncChan is intended as the simplest bridge between Go channels
//	and Observables. It models Go’s native channel behavior directly.
//
//	For fan-out, replay, buffering, or backpressure-aware behavior,
//	use a multicast or subject-based Observable instead.
func AsyncChan[T any](channel <-chan T) Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncChan[T](channel)}
}

// AsyncEmpty creates a cold, asynchronous Observable that emits no
// values and completes asynchronously.
//
// AsyncEmpty is the asynchronous counterpart of Empty.
// It immediately schedules a completion signal on a separate
// goroutine for each subscription.
//
// AsyncEmpty is a cold Observable:
//
//	The completion is not scheduled until Sub is called.
//	Range subscription schedules its own asynchronous completion.
//
// Asynchronous semantics:
//
//	Complete is invoked from a goroutine created per subscription.
//	Observers are never completed synchronously during Sub.
//
// Emission semantics:
//
//	No values are emitted.
//	No Error is emitted.
//
// Completion semantics:
//
//	Complete is always signaled exactly once, unless the
//	subscription is cancelled before the goroutine runs.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	If Unsub is called before completion is delivered,
//	the completion signal is suppressed.
//
// Diagram:
//
//	AsyncEmpty():
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	   --|
//
// Example:
//
//	src := AsyncEmpty[int]()
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println("unexpected:", v) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output:
//
//	done
//
// Notes:
//
//	AsyncEmpty is useful when an asynchronous boundary is required,
//	such as:
//	  - Forcing downstream operators to observe completion
//	    asynchronously
//	  - Testing async behavior without emitting values
//	  - Composing with operators that distinguish sync vs async
//	    completion paths
//
//	For synchronous completion, use Empty instead.
func AsyncEmpty[T any]() Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncEmpty[T]()}
}

// AsyncError creates a cold, asynchronous Observable that emits no
// values and terminates with an error asynchronously.
//
// AsyncError is the asynchronous counterpart of Error.
// It schedules an Error signal on a separate goroutine for each
// subscription.
//
// AsyncError is a cold Observable:
//
//	The error is not scheduled until Sub is called.
//	Range subscription schedules its own asynchronous error signal.
//
// Asynchronous semantics:
//
//	Error is invoked from a goroutine created per subscription.
//	Observers are never errored synchronously during Sub.
//
// Emission semantics:
//
//	No values are emitted.
//	The provided error is emitted exactly once.
//
// Error semantics:
//
//	The error passed to AsyncError is delivered verbatim to the
//	observer’s Error callback.
//	No wrapping, transformation, or recovery is performed.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	If Unsub is called before the error is delivered,
//	the error signal is suppressed.
//
// Diagram:
//
//	AsyncError(err):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	   --X
//
// Example:
//
//	err := errors.New("error.")
//
//	src := AsyncError[int](err)
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println("unexpected:", v) },
//		Error: func(e error) { fmt.Println("error:", e) },
//	})
//
// Output:
//
//	error: error.
//
// Notes:
//
//	AsyncError is useful when an asynchronous failure boundary is
//	required, such as:
//	  - Simulating async failures in tests
//	  - Forcing downstream operators to observe errors
//	    asynchronously
//	  - Maintaining consistency with async sources
//
//	For synchronous error emission, use Error instead.
func AsyncError[T any](err error) Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncError[T](err)}
}

// AsyncDelayedOf creates a cold, asynchronous Observable that emits the
// provided values sequentially, delaying each emission according to
// a user-defined function.
//
// Range value v[i] is emitted after sleeping for the duration returned
// by delay(i, v[i]). Values are emitted in order, and completion is
// signaled immediately after the final value is emitted.
//
// AsyncDelayedOf is a cold Observable:
//
//	The emission process does not start until Sub is called.
//	Range subscription schedules its own goroutine and timing sequence.
//
// Asynchronous semantics:
//
//	Values and completion are emitted from a dedicated goroutine per
//	subscription.
//	Observer callbacks (Next, Complete) are invoked asynchronously.
//
// Emission semantics:
//
//	Values are emitted in the same order as provided.
//	Range value is delayed independently using the delay function.
//	Delays are applied sequentially.
//
// Timing semantics:
//
//	Delays are attached to values, not to the stream itself.
//	The delay for value i starts only after value i-1 has been emitted.
//	This is not a timer-based scheduler and does not compensate for
//	execution drift.
//
// Completion semantics:
//
//	AsyncDelayedOf does NOT delay completion.
//	Complete is emitted immediately after the last value is emitted.
//
//	If no values are provided, the Observable completes immediately.
//
//	This behavior differs from AsyncIntervalOf, which delays completion
//	according to the interval even when no values are present.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	Cancellation does not interrupt sleeps or stop the emission loop.
//	After Unsub, emitted values are ignored, but the goroutine continues
//	running until all delays have elapsed.
//
// Diagram:
//
//	AsyncDelayedOf(delay, 1, 2, 3):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	--(d0)--1--(d1)--2--(d2)--3--|
//
// Example:
//
//	src := AsyncDelayedOf(
//		func(i int, v int) time.Duration {
//			return time.Duration(i+1) * time.Second
//		},
//		1, 2, 3,
//	)
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v, time.AtNow()) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output (timing approximate):
//
//	1  +1s
//	2  +2s
//	3  +3s
//	done
//
// Notes:
//
//	AsyncDelayedOf models value-relative timing.
//	It is best suited for sequential, per-item delays.
//
//	For stream-relative timing, periodic emission, or delayed
//	completion semantics, use AsyncIntervalOf instead.
func AsyncDelayedOf[T any](
	delay func(index int, value T) time.Duration,
	v ...T,
) Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncDelayedOf[T](delay, v...)}
}

// AsyncIntervalOf creates a cold, asynchronous Observable that emits the
// provided values at fixed time intervals.
//
// AsyncIntervalOf applies a uniform delay to the stream itself.
// Emissions and completion are scheduled according to a fixed interval,
// independent of individual values.
//
// AsyncIntervalOf is a cold Observable:
//
//	The interval schedule does not start until Sub is called.
//	Range subscription creates its own goroutine and timing sequence.
//
// Asynchronous semantics:
//
//	Values and completion are emitted from a dedicated goroutine per
//	subscription.
//	Observer callbacks (Next, Complete) are invoked asynchronously.
//
// Timing semantics:
//
//	An initial delay is applied before the first emission.
//	After each emission, the interval delay is applied again.
//
//	The interval applies even when no values are provided.
//
// Emission semantics:
//
//	Values are emitted in the same order as provided.
//	Exactly one value is emitted per interval tick.
//
// Completion semantics:
//
//	Completion is delayed according to the interval.
//
//	If values are provided:
//	  - Completion occurs after the final interval delay.
//
//	If no values are provided:
//	  - Completion occurs after a single interval delay.
//
// This behavior differs from AsyncDelayedOf, which applies delays
// only to individual values and completes immediately when empty.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	Cancellation does not interrupt the sleep or stop the scheduling loop.
//	After Unsub, emitted values are ignored, but the goroutine continues
//	running until completion.
//
// Diagram:
//
//	AsyncIntervalOf(delay, 1, 2, 3):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	--(d)--1--(d)--2--(d)--3--(d)--|
//
//	AsyncIntervalOf(delay):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	--(d)--|
//
// Example:
//
//	src := AsyncIntervalOf(1*time.Second, 1, 2, 3)
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v, time.AtNow()) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output (timing approximate):
//
//	1  +1s
//	2  +2s
//	3  +3s
//	done +4s
//
// Notes:
//
//	AsyncIntervalOf models stream-relative timing.
//	The interval governs when emissions and completion occur,
//	not the values themselves.
//
//	For per-value delays without delayed completion,
//	use AsyncDelayedOf instead.
func AsyncIntervalOf[T any](delay time.Duration, v ...T) Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncIntervalOf[T](delay, v...)}
}

// AsyncJust creates a cold, asynchronous Observable that emits a single
// value and then completes.
//
// AsyncJust introduces an asynchronous boundary for a single value.
// The value and completion are scheduled on a separate goroutine
// for each subscription.
//
// AsyncJust is a cold Observable:
//
//	The value is not emitted until Sub is called.
//	Range subscription schedules its own emission.
//
// Asynchronous semantics:
//
//	The value and completion are emitted from a dedicated goroutine
//	per subscription.
//	Observer callbacks (Next, Complete) are invoked asynchronously.
//
// Emission semantics:
//
//	The provided value is emitted exactly once.
//	Complete is signaled immediately after the value is emitted.
//	No Error is emitted.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	If Unsub is called before the emission occurs, the value and
//	completion are suppressed.
//
// Diagram:
//
//	AsyncJust(1):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	   --1--|
//
// Example:
//
//	src := AsyncJust(42)
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output:
//
//	42
//	done
//
// Notes:
//
//	AsyncJust is useful when an asynchronous boundary is required
//	for a single, known value.
//
//	For synchronous emission, use Just instead.
func AsyncJust[T any](v T) Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncJust[T](v)}
}

// AsyncMulticast creates a channel-backed Observable that asynchronously
// emits values received from the provided channel to all active subscribers.
//
// AsyncMulticast is a cold, asynchronous Observable with multicast
// semantics. Values are only read from the channel after the first
// subscription is created, and a single shared goroutine is responsible
// for consuming the channel and broadcasting values to all subscribers.
//
// Important: AsyncMulticast IS multicast.
//
// All active subscribers receive every value read from the underlying
// channel, in the same order. Late subscribers only receive values
// emitted after they subscribe; no replay or buffering is performed.
//
// Cold semantics:
//
//	The underlying channel is not read until at least one subscription
//	is created.
//	Values sent to the channel before the first subscription are not observed.
//	If no subscription exists, sending to an unbuffered channel will block.
//
// Asynchronous semantics:
//
//	Values are emitted from a single shared goroutine.
//	Observer callbacks (Next, Complete) are invoked asynchronously
//	with respect to the subscriber.
//
// Multicast semantics:
//
//	A single goroutine reads from the channel and fans out each value
//	to all currently subscribed observers.
//	Subscription order does not affect value ordering.
//	Observers added after a value is emitted will not receive it.
//
// Completion semantics:
//
//   - At the channel is closed, Complete is signaled to all
//     active subscribers.
//   - No Error is produced by channel closure.
//   - After completion, no further values are emitted.
//
// Cancellation semantics:
//
//	Range subscription can be independently cancelled.
//	If Unsub is called, the observer stops receiving values.
//	Cancelling a subscription does not affect the channel reader
//	or other subscribers.
//
// Diagram:
//
//	ch := make(chan int)
//
//	AsyncMulticast(ch):
//
//	          subscribe A
//	              |
//	          subscribe B
//	              |
//	              v
//	        (shared goroutine)
//	              |
//	              v
//	        --1--2--3--|
//	          |  |  |
//	          v  v  v
//	          A  A  A
//	          B  B  B
//
// Example:
//
//	ch := make(chan int)
//
//	src := AsyncMulticast(ch)
//
//	go func() {
//		ch <- 1
//		ch <- 2
//		ch <- 3
//		close(ch)
//	}()
//
//	src.Sub(observer.NewObserver(
//		func(v int) { fmt.Println("A:", v) },
//		func(err error) {},
//		func() { fmt.Println("A done") },
//	))
//
//	src.Sub(observer.NewObserver(
//		func(v int) { fmt.Println("B:", v) },
//		func(err error) {},
//		func() { fmt.Println("B done") },
//	))
//
// Output (order preserved per subscriber):
//
//	A: 1
//	B: 1
//	A: 2
//	B: 2
//	A: 3
//	B: 3
//	A done
//	B done
//
// Notes:
//
//	AsyncMulticast is the appropriate bridge when a Go channel represents
//	a hot, shared data source that must be observed by multiple consume.
//
//	It does not provide replay, buffering, or backpressure control.
//	For replay or stateful fan-out, use an explicit replay or subject-based
//	Observable instead.
func AsyncMulticast[T any](channel <-chan T) Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncMulticast[T](channel)}
}

// AsyncOf creates a cold, asynchronous Observable that emits the
// provided values sequentially and then completes.
//
// AsyncOf introduces an asynchronous boundary for a finite sequence
// of known values. All values and the completion signal are emitted
// from a separate goroutine per subscription.
//
// AsyncOf is a cold Observable:
//
//	The sequence is not emitted until Sub is called.
//	Range subscription schedules its own emission.
//
// Asynchronous semantics:
//
//	Values and completion are emitted from a dedicated goroutine
//	per subscription.
//	Observer callbacks (Next, Complete) are invoked asynchronously.
//
// Emission semantics:
//
//	Values are emitted in the same order as provided.
//	Range value is emitted exactly once.
//	Complete is signaled after the final value is emitted.
//	No Error is emitted.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	Cancellation does not stop the emission goroutine.
//	After Unsub, emitted values are ignored, but the goroutine
//	continues until completion.
//
// Diagram:
//
//	AsyncOf(1, 2, 3):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	   --1--2--3--|
//
// Example:
//
//	src := AsyncOf(1, 2, 3)
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output:
//
//	1
//	2
//	3
//	done
//
// Notes:
//
//	AsyncOf is useful when an asynchronous boundary is required
//	for a fixed, in-memory sequence.
//
//	For synchronous emission, use Of instead.
//	For delayed or interval-based emission, use AsyncDelayedOf
//	or AsyncIntervalOf instead.
func AsyncOf[T any](v ...T) Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncOf(v...)}
}

// AsyncRange creates a cold, asynchronous Observable that emits a sequence
// of numeric values starting from `from` and ending at or before `to`,
// incrementing (or decrementing) by `delta`.
//
// AsyncRange is asynchronous:
//
//   - Range subscription starts a dedicated goroutine
//   - Values and completion are emitted asynchronously
//   - Subscribe returns immediately, before any values are emitted
//
// Parameters:
//
//   - from: the starting value of the sequence
//   - to: the inclusive upper (or lower) bound of the sequence
//   - delta: the step between consecutive values
//
// Panics:
//
//   - delta == 0 → panic("delta must be non-zero")
//   - delta < 0 && from < to → panic("delta must be non-negative when from < to")
//   - delta > 0 && from > to → panic("delta must be negative when from > to")
//
// Timing semantics:
//
//   - Values are emitted asynchronously in order, one after another
//   - Completion is emitted after the last value
//   - No delay is applied between values unless you wrap it with a delay operator
//
// Emission semantics:
//
//   - Values are emitted in ascending order if delta > 0
//   - Values are emitted in descending order if delta < 0
//   - Range value is emitted exactly once
//   - No errors are emitted
//
// Completion semantics:
//
//   - Completion occurs after the last value is emitted
//   - If no values satisfy the range constraints, completion occurs immediately
//
// Cancellation semantics:
//
//   - The returned Subscription supports cancellation via Unsub
//   - After Unsub, further emitted values are ignored
//   - The goroutine may continue running until the sequence completes
//
// Diagram:
//
//	AsyncRange(1, 5, 1):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	--1--2--3--4--5--|
//
//	AsyncRange(5, 1, -2):
//
//	subscribe
//	   |
//	   v
//	(goroutine)
//	   |
//	   v
//	--5--3--1--|
//
// Example:
//
//	src := AsyncRange(1, 5, 1)
//
//	src.Sub(&obs.Observer{
//	    Next: func(v int) { fmt.Println(v, time.AtNow()) },
//	    Complete: func() { fmt.Println("done") },
//	})
//
// Output (timing approximate):
//
//	1  <async>
//	2  <async>
//	3  <async>
//	4  <async>
//	5  <async>
//	done
//
// Notes:
//
//   - AsyncRange introduces a goroutine per subscription to emit values
//     asynchronously, unlike the synchronous Range which blocks the
//     subscribing goroutine.
//   - For interval-based or delayed emissions, consider using
//     AsyncDelayedOf or AsyncIntervalOf wrapped around a numeric range.
func AsyncRange[T nums.Ordered](from, to, delta T) Observable[T] {
	return observableAdapter[T]{inner: observables.AsyncRange(from, to, delta)}
}

// Create creates a cold Observable by invoking a user-provided
// subscription function.
//
// Create is a low-level constructor that gives full control over
// emission and lifecycle to the user.
//
// The provided function fn is invoked once per subscription and
// receives an Observer through which values may be emitted.
// fn must return a Subscription representing the underlying work.
//
// Create is a cold Observable:
//
//	The function fn is not invoked until Sub is called.
//	Range subscription invokes fn independently.
//
// Execution semantics:
//
//	Create does not introduce an asynchronous boundary.
//	The function fn is invoked synchronously during Sub.
//
//	Any asynchronous behavior must be explicitly introduced
//	by fn (for example, by spawning goroutines).
//
// Observer semantics:
//
//	The Observer passed to fn is wrapped with a lifecycle gate:
//
//	  - Next calls after cancellation are ignored
//	  - Error closes the subscription and is delivered at most once
//	  - Complete closes the subscription and is delivered at most once
//
//	If Next, Error, or Complete are nil, they are replaced with no-op
//	implementations.
//
// Subscription semantics:
//
//	The Subscription returned by fn is bound to the Observable’s
//	infra lifecycle gate.
//
//	If fn returns a non-nil error:
//	  - The lifecycle gate is immediately closed
//	  - The error is returned from Sub
//	  - No values are emitted
//
//	Create does not recover from panics inside fn.
//
// Cancellation semantics:
//
//	If Unsub is called on the returned Subscription:
//	  - The lifecycle gate is closed
//	  - Further emissions via the Observer are ignored
//
//	It is the responsibility of fn to stop any ongoing work
//	when cancellation occurs.
//
// Diagram:
//
//	Create(fn):
//
//	subscribe
//	   |
//	   v
//	fn(o)
//	   |
//	   +--> Next / Error / Complete
//
// Example:
//
//	src := Create(func(o *Observer[int]) (Subscription, error) {
//		go func() {
//			for i := 0; i < 3; i++ {
//				o.Next(i)
//			}
//			o.Complete()
//		}()
//		return subscr.Empty(), nil
//	})
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output:
//
//	0
//	1
//	2
//	done
//
// Notes:
//
//	Create provides maximal flexibility and minimal guarantees.
//	It is intended for:
//
//	  - Bridging external or callback-based APIs
//	  - Integrating systems with custom lifecycles
//	  - Implementing higher-level Observable constructors
//
//	Most users should prefer higher-level constructors such as
//	Async, AsyncOf, or AsyncChan.
func Create[T any](fn func(*Observer[T]) (Subscription, error)) Observable[T] {
	inner := observables.Func[T](func(o *observers.Observer[T]) (subscr.Subscription, error) {
		gate := gates.Atomic()

		created := observers.Create(o, gate)

		sub, err := fn((*Observer[T])(created))
		if err != nil {
			gate.Close()
			return nil, err
		}

		return subscr.Bind(sub, gate), nil
	})

	return observableAdapter[T]{inner: inner}
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
func Defer[T any](fn func() Observable[T]) Observable[T] {
	return Create[T](func(o *Observer[T]) (Subscription, error) {
		var sub Subscription

		inner := fn()
		s, err := inner.Sub(o)
		if err != nil {
			return nil, err
		}

		sub = s
		return sub, nil
	})
}

// DelayedOf creates a cold, synchronous Observable that emits the
// provided values sequentially, blocking between emissions according
// to a user-defined delay function.
//
// Range value v[i] is emitted after blocking for the duration returned
// by delay(i, v[i]). Values are emitted in order, and completion is
// signaled after the final value is emitted.
//
// DelayedOf is a cold Observable:
//
//	The emission process does not start until Sub is called.
//	Range subscription executes independently.
//
// Synchronous semantics:
//
//	Emission occurs synchronously during Sub.
//	The calling goroutine is blocked while delays are applied.
//	Observer callbacks (Next, Complete) are invoked inline.
//
// Emission semantics:
//
//	Values are emitted in the same order as provided.
//	Range value is emitted exactly once.
//	No Error is emitted.
//
// Timing semantics:
//
//	Delays are attached to values, not to the stream itself.
//	The delay for value i starts only after value i-1 has been emitted.
//
// Completion semantics:
//
//	Completion is emitted immediately after the last value is emitted.
//	If no values are provided, the Observable completes immediately.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	However, because emission is synchronous and blocking,
//	cancellation cannot interrupt the delay or emission loop.
//	After Unsub, further emissions are ignored.
//
// Diagram:
//
//	DelayedOf(delay, 1, 2, 3):
//
//	subscribe
//	   |
//	   v
//	--(d0)--1--(d1)--2--(d2)--3--|
//
// Example:
//
//	src := DelayedOf(
//		func(i int, v int) time.Duration {
//			return time.Duration(i+1) * time.Second
//		},
//		1, 2, 3,
//	)
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v, time.AtNow()) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output (timing approximate):
//
//	1  +1s
//	2  +2s
//	3  +3s
//	done
//
// Notes:
//
//	DelayedOf blocks the calling goroutine.
//	It is best suited for controlled, deterministic execution,
//	such as tests or examples.
//
//	For non-blocking behavior, use AsyncDelayedOf instead.
func DelayedOf[T any](
	delay func(index int, value T) time.Duration,
	v ...T,
) Observable[T] {
	return observableAdapter[T]{inner: observables.DelayedOf[T](delay, v...)}
}

// Empty creates a cold Observable that completes immediately upon subscription
// without emitting any values.
//
// The completion is synchronous and happens during Subscribe.
// No goroutines are started and no asynchronous boundary is introduced.
//
// Notes:
//   - Complete() is called immediately when subscribing
//   - Subscribe returns after completion
//   - Unsubscription has no effect
//   - No values are ever emitted
//
// Empty is typically used as a terminal source or as a neutral element
// in composition.
func Empty[T any]() Observable[T] {
	return observableAdapter[T]{inner: observables.Empty[T]()}
}

// Error creates a cold Observable that terminates immediately with the
// provided error upon subscription.
//
// The error is delivered synchronously during Subscribe.
// No values are emitted and no asynchronous boundary is introduced.
//
// Notes:
//   - Error(err) is called immediately when subscribing
//   - Subscribe returns after the error is delivered
//   - The observable is terminal
//   - Unsubscription has no effect
func Error[T any](err error) Observable[T] {
	return observableAdapter[T]{inner: observables.Error[T](err)}
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
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output:
//
//	1
//	2
//	done
func FromFunc[T any](fn func(*Observer[T])) Observable[T] {
	internal := observables.Func[T](func(o *observers.Observer[T]) (subscr.Subscription, error) {
		gate := gates.Atomic()
		created := observers.Create(o, gate)
		sub := subscr.NewSlotSub(created, gate)
		fn((*Observer[T])(created))
		return sub, nil
	})
	return observableAdapter[T]{inner: internal}
}

// IntervalOf creates a cold, synchronous Observable that emits the
// provided values at fixed time intervals.
//
// IntervalOf applies a uniform delay to the stream itself.
// Emissions and completion are spaced according to a fixed interval,
// independent of individual values.
//
// IntervalOf is a cold Observable:
//
//	The interval schedule does not start until Sub is called.
//	Range subscription executes independently on the calling goroutine.
//
// Synchronous semantics:
//
//	Values and completion are emitted synchronously during Subscribe.
//	The calling goroutine is blocked while the interval sleeps are applied.
//	Observer callbacks (Next, Complete) are invoked inline.
//
// Timing semantics:
//
//	An initial delay is applied before the first emission.
//	After each emission, the interval delay is applied again.
//
//	The interval applies even when no values are provided.
//
// Emission semantics:
//
//	Values are emitted in the same order as provided.
//	Exactly one value is emitted per interval tick.
//
// Completion semantics:
//
//	Completion is delayed according to the interval.
//
//	If values are provided:
//	  - Completion occurs after the final interval delay.
//
//	If no values are provided:
//	  - Completion occurs after a single interval delay.
//
// This behavior differs from DelayedOf, which applies delays
// only to individual values and completes immediately when empty.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation via Unsub.
//	Since IntervalOf is synchronous, cancellation cannot interrupt
//	sleep or emission. After Unsub, further emissions are ignored,
//	but the calling goroutine remains blocked until completion.
//
// Diagram:
//
//	IntervalOf(delay, 1, 2, 3):
//
//	subscribe
//	   |
//	   v
//	--(d)--1--(d)--2--(d)--3--(d)--|
//
//	IntervalOf(delay):
//
//	subscribe
//	   |
//	   v
//	--(d)--|
//
// Example:
//
//	src := IntervalOf(1*time.Second, 1, 2, 3)
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v, time.AtNow()) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output (timing approximate):
//
//	1  +1s
//	2  +2s
//	3  +3s
//	done +4s
//
// Notes:
//
//	IntervalOf models stream-relative timing in a blocking manner.
//	The interval governs when emissions and completion occur,
//	not the values themselves.
//
//	For non-blocking, goroutine-based emission, use AsyncIntervalOf
//	instead.
func IntervalOf[T any](delay time.Duration, v ...T) Observable[T] {
	return observableAdapter[T]{inner: observables.IntervalOf[T](delay, v...)}
}

// Just creates a cold, synchronous Observable that immediately emits
// a single value and then completes.
//
// Just is a blocking, synchronous source:
//
//   - The value is emitted during Subscribe
//   - Complete is called immediately after the value
//   - Subscribe returns only after the value and completion are delivered
//
// Emission semantics:
//
//   - Exactly one value is emitted
//   - Observer callbacks (Next, Complete) are invoked inline
//
// Completion semantics:
//
//   - Completion occurs immediately after the value
//   - No error is ever emitted
//
// Cancellation semantics:
//
//   - The returned Subscription is a no-op
//   - Since emission is synchronous, cancellation cannot interrupt the emission
//
// Diagram:
//
//	Just(42):
//
//	subscribe
//	   |
//	   v
//	--42--|
//
// Example:
//
//	src := Just(42)
//
//	src.Sub(&obs.Observer{
//		Next: func(v int) { fmt.Println(v) },
//		Complete: func() { fmt.Println("done") },
//	})
//
// Output:
//
//	42
//	done
//
// Notes:
//
//   - Just is the synchronous counterpart to AsyncJust.
//   - For an asynchronous emission boundary, use AsyncJust instead.
func Just[T any](v T) Observable[T] {
	return observableAdapter[T]{inner: observables.Just[T](v)}
}

// Never creates a cold Observable that never emits any values,
// never signals completion, and never signals an error.
//
// Never is a blocking-free, infinite idle source:
//
//   - Subscribe returns immediately
//   - No Next, Complete, or Error calls are ever made
//   - The Observable remains idle indefinitely
//
// Subscription semantics:
//
//   - Range subscription is independent
//   - The returned Subscription is a no-op
//   - Unsub has no effect on emissions (because there are none)
//
// Timing semantics:
//
//   - No delays are introduced
//   - No asynchronous boundary is created
//
// Diagram:
//
//	Never():
//
//	subscribe
//	   |
//	   v
//	(no commands ever)
//
// Example:
//
//	src := Never[int]()
//
//	src.Sub(&obs.Observer{
//	    Next: func(v int) { fmt.Println(v) },      // never called
//	    Error: func(err error) { fmt.Println(err) }, // never called
//	    Complete: func() { fmt.Println("done") }, // never called
//	})
//
// Output:
//
//	(nothing)
//
// Notes:
//
//   - Never is typically used for testing, placeholders, or
//     as a building block in combination with operators.
//   - Unlike Empty, it never completes, so any operators that
//     wait for completion will never terminate.
func Never[T any]() Observable[T] {
	return Create[T](func(o *Observer[T]) (Subscription, error) {
		return subscr.Noop{}, nil
	})
}

// Of creates a cold Observable that synchronously emits the provided values
// in order and then completes.
//
// Of is a synchronous, finite Observable. All values are emitted immediately
// upon subscription, on the caller’s goroutine, unless the subscription is
// cancelled.
//
// Cancellation semantics:
//
//	The returned Subscription supports cancellation. If Unsub is called
//	while values are being emitted, emission stops immediately and the
//	remaining values are discarded.
//
// This is important because downstream pipes (such as Delay, Debounce,
// Timeout, etc.) may introduce asynchronous behavior and rely on proper
// cancellation propagation.
//
// Diagram:
//
//	Of(1, 2, 3):
//
//	subscribe
//	   |
//	   v
//	--1--2--3--|
//
// Example:
//
//	src := Of(1, 2, 3)
//
//	src.Sub(observer.NewObserver(
//		func(v int) { fmt.Println(v) },
//		func(err error) {},
//		func() { fmt.Println("done") },
//	))
//
// Output:
//
//	1
//	2
//	3
//	done
func Of[T any](v ...T) Observable[T] {
	return observableAdapter[T]{inner: observables.Of(v...)}
}

func Serialize[T any](o *Observer[T]) *Observer[T] {
	// src.Sub(obs.Serialize(o))
	var (
		mu         sync.Mutex
		terminated bool
	)

	return &Observer[T]{
		Next: func(v T) {
			mu.Lock()
			defer mu.Unlock()

			if terminated {
				return
			}
			if o.Next != nil {
				o.Next(v)
			}
		},

		Error: func(err error) {
			mu.Lock()
			defer mu.Unlock()

			if terminated {
				return
			}
			terminated = true

			if o.Error != nil {
				o.Error(err)
			}
		},

		Complete: func() {
			mu.Lock()
			defer mu.Unlock()

			if terminated {
				return
			}
			terminated = true

			if o.Complete != nil {
				o.Complete()
			}
		},
	}
}

// Range creates a cold, synchronous Observable that emits a sequence
// of numeric values starting from `from` and ending at or before `to`,
// incrementing (or decrementing) by `delta`.
//
// Range is synchronous:
//
//   - All values are emitted during Subscribe
//   - Observer callbacks (Next, Complete) are invoked inline
//   - Subscribe returns only after the last value and completion
//
// Parameters:
//
//   - from: the starting value of the sequence
//   - to: the inclusive upper (or lower) bound of the sequence
//   - delta: the step between consecutive values
//
// Panics:
//
//   - delta == 0 → panic("delta must be non-zero")
//   - delta < 0 && from < to → panic("delta must be non-negative when from < to")
//   - delta > 0 && from > to → panic("delta must be negative when from > to")
//
// Emission semantics:
//
//   - Values are emitted in ascending order if delta > 0
//   - Values are emitted in descending order if delta < 0
//   - Emission stops when the next value would pass the `to` bound
//   - No errors are emitted
//
// Completion semantics:
//
//   - Completion is signaled immediately after the last value
//   - If no values satisfy the range constraints, completion is immediate
//
// Cancellation semantics:
//
//   - The returned Subscription supports Unsub, but since emission is
//     synchronous, cancellation cannot interrupt the emission loop
//
// Diagram:
//
//	Range(1, 5, 1):
//
//	subscribe
//	   |
//	   v
//	--1--2--3--4--5--|
//
//	Range(5, 1, -2):
//
//	subscribe
//	   |
//	   v
//	--5--3--1--|
//
// Example:
//
//	src := Range(1, 5, 1)
//
//	src.Sub(&obs.Observer{
//	    Next: func(v int) { fmt.Println(v) },
//	    Complete: func() { fmt.Println("done") },
//	})
//
// Output:
//
//	1
//	2
//	3
//	4
//	5
//	done
//
// Notes:
//
//   - Range is synchronous and blocking. Pipes or operators may alter
//     this behavior if they introduce async boundaries.
//   - Use AsyncRange (or a similar async wrapper) if you want non-blocking
//     emissions or interval-based sequencing.
func Range[T nums.Ordered](from, to, delta T) Observable[T] {
	return observableAdapter[T]{inner: observables.Range(from, to, delta)}
}
