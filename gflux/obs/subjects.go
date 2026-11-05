// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package obs

import (
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observables"
)

// Subject is a hot observable
// It owns its observers
// It is both producer and multicaster
type Subject[T any] interface {
	Observable[T] // The Subject is Observable
	Next(T)
	Error(error)
	Complete()
	UnsubAll()
}

// AsyncSubject returns a Subject that emits only the last value upon completion.
//
// An AsyncSubject captures the most recent value it receives, but does not emit
// any values until the subject is completed. Only the final value (if any) is
// emitted, followed by the terminal signal.
//
// If the subject has already terminated:
//   - the last value is emitted (if present and the subject completed successfully)
//   - then the terminal signal is delivered (Error or Complete)
//
// After termination, further Next calls are ignored.
//
// Error handling:
//   - Error terminates the subject immediately and is broadcast to all
//     current observers.
//   - Any stored value is discarded.
//   - Late subscribers receive only the error.
//
// Completion:
//   - Completion emits the last stored value (if any) to all current observers,
//     then broadcasts completion.
//   - Late subscribers receive the last value (if any), then completion.
//
// AsyncSubject multicasts to all active subscribers and uses slot-based
// subscriptions, allowing safe reentrancy during emission.
//
// Diagram:
//
//	subscribe A        subscribe B
//	    |                   |
//	    v                   v
//	    |
//	    |  Next(1)
//	    |
//	    |  Next(2)
//	    |
//	    |  Next(3)
//	    |
//	    |  Complete()
//	    v
//	    |
//	A:        3--|
//	B:        3--|
//
// Example:
//
//	subj := obs.AsyncSubject[int]()
//
//	subj.Next(1)
//	subj.Next(2)
//	subj.Next(3)
//
//	subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println("A:", v) },
//		nil,
//		func() { fmt.Println("A: done") },
//	))
//
//	subj.Complete()
//
//	subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println("B:", v) },
//		nil,
//		func() { fmt.Println("B: done") },
//	))
//
// Output:
//
//	A: 3
//	A: done
//	B: 3
//	B: done
func AsyncSubject[T any]() Subject[T] {
	return subjectAdapter[T]{observables.AsyncSubject[T]()}
}

// BehaviorSubject returns a Subject that holds and replays the latest value.
//
// A BehaviorSubject always has a current value. Upon subscription, the observer
// immediately receives the most recent value (starting with the provided
// initial value), followed by any subsequent emissions.
//
// If the subject has already terminated:
//   - the last value is emitted (if present)
//   - then the terminal signal is delivered (Error or Complete)
//
// After termination, further Next calls are ignored.
//
// Error handling:
//   - Error terminates the subject and is broadcast to all current observers.
//   - Late subscribers receive the last value, then the error.
//
// Completion:
//   - Completion terminates the subject and is broadcast to all current observers.
//   - Late subscribers receive the last value, then completion.
//
// BehaviorSubject multicasts to all active subscribers and uses slot-based
// subscriptions, allowing safe reentrancy during emission.
//
// Diagram:
//
//	subscribe A        subscribe B
//	    |                   |
//	    v                   v
//	(value=0)
//	    |
//	    |  Next(1)
//	    v
//	    1
//	    |
//	    |  Next(2)
//	    v
//	    2
//	    |
//	    |  Complete()
//	    v
//	    |
//
//	A:  0--1--2--|
//	B:        2--|
//
// Example:
//
//	subj := obs.BehaviorSubject(0)
//
//	subj.Next(1)
//	subj.Next(2)
//
//	subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println("A:", v) },
//		nil,
//		func() { fmt.Println("A: done") },
//	))
//
//	subj.Complete()
//
//	subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println("B:", v) },
//		nil,
//		func() { fmt.Println("B: done") },
//	))
//
// Output:
//
//	A: 2
//	A: done
//	B: 2
//	B: done
func BehaviorSubject[T any](initial T) Subject[T] {
	return subjectAdapter[T]{observables.BehaviorSubject[T](initial)}
}

// ConcSlotSubject returns a concurrent-safe Subject.
//
// A concurrent Subject serializes all interactions (Sub, Next, Error,
// Complete, UnsubAll) using mutual exclusion, making it safe to use from
// multiple goroutines.
//
// Emission semantics:
//   - Values are multicast to all active subscribers.
//   - Calls to Next, Error, and Complete are mutually exclusive.
//   - Relative ordering between concurrent calls from different goroutines
//     is not guaranteed.
//
// Termination:
//   - Error or Complete terminates the subject.
//   - After termination, further Next calls are ignored.
//   - Late subscribers immediately receive the terminal signal.
//
// Concurrency guarantees:
//   - No data races or panics due to concurrent access.
//   - Observer callbacks are invoked while holding the infra lock.
//   - Observers MUST NOT synchronously call Sub, Unsub, Next, Error, or
//     Complete from within callbacks, or a deadlock may occur.
//
// Error handling:
//   - Error is broadcast to all current observers.
//   - Late subscribers receive the terminal error.
//
// Completion:
//   - Completion is broadcast to all current observers.
//   - Late subscribers complete immediately.
//
// This Subject uses slot-based subscriptions internally and favors
// correctness and safety over maximal throughput.
//
// Diagram (two goroutines):
//
//	go1: Next(1) --------->
//	go2: Error(e) -------->
//
//	Actual order depends on lock acquisition:
//
//	Possible A:
//	  Next(1) → Error(e)
//
//	Possible B:
//	  Error(e) → Next(1) (ignored)
//
// Example:
//
//	subj := obs.ConcSlotSubject[int]()
//
//	go subj.Next(1)
//	go subj.Error(errors.New("fail"))
//
//	// Outcome is deterministic in safety, but not in ordering.
func ConcSlotSubject[T any]() Subject[T] {
	return subjectAdapter[T]{observables.ConcSubject[T]()}
}

// DirectSubject returns a Subject with immediate unsubscribe semantics.
//
// subscr.Subscriptions are removed eagerly and directly from the subject’s infra
// structure. This avoids per-emission liveness checks and favors fast emission
// at the cost of higher mutation overhead and stricter lifecycle discipline.
//
// Observers MUST NOT unsubscribe synchronously from Next, Error, or Complete.
// Doing so may cause a runtime panic due to map mutation during iteration.
//
// This Subject is not safe for concurrent mutation unless otherwise stated.
func DirectSubject[T any]() Subject[T] {
	return subjectAdapter[T]{observables.DirectSubject[T]()}
}

// ManySubject returns a Subject that allows a bounded number of subscribers.
//
// A ManySubject multicasts values to all active subscribers, but enforces
// an upper limit on the number of concurrent subscriptions. At most `max`
// observers may be subscribed at any given time.
//
// If the subscription limit is reached, additional Sub calls fail with
// an error and the observer is not registered.
//
// At a subscriber unsubscribes, its slot is freed, allowing new
// subscriptions up to the configured capacity.
//
// Termination semantics:
//   - Error or Complete terminates the subject and is broadcast to all
//     active subscribers.
//   - Late subscribers immediately receive the terminal signal.
//   - After termination, further Next calls are ignored.
//   - Len limits are ignored after termination.
//
// Error handling:
//   - Error terminates the subject and is broadcast to all current observers.
//   - Late subscribers receive the terminal error immediately.
//
// Completion:
//   - Completion terminates the subject and is broadcast to all current observers.
//   - Late subscribers receive completion immediately.
//
// ManySubject uses slot-based subscriptions, allowing safe reentrancy
// during emission.
//
// Concurrency:
//   - ManySubject is NOT safe for concurrent Sub, Next, Error, or Complete calls.
//   - For concurrent usage, wrap it with a serialized or concurrent subject.
//
// Panics:
//   - Panics if max <= 0.
//
// Diagram (max = 2):
//
//	subscribe A        subscribe B        subscribe C (rejected)
//	    |                   |                   |
//	    v                   v                   x
//	    |                   |
//	    |   Next(1)          |
//	    v                   v
//	    1                   1
//	    |                   |
//	    |   Unsub A          |
//	    v                   |
//	                        |
//	    subscribe C          |
//	        |               |
//	        v               |
//	        |   Next(2)     |
//	        v               v
//	        2               2
//	        |               |
//	        |   Complete()  |
//	        v               v
//	        |               |
//
//	A:  1
//	B:  1--2--|
//	C:      2--|
//
// Example:
//
//	subj := obs.ManySubject
//
//	a := obs.ObserverImpl[int]{Next: func(v int) { fmx.Print(i) }, Complete: func() { fmx.Print("done") }}
//	b := obs.ObserverImpl[int]{Next: func(v int) { fmx.Print(i) }, Complete: func() { fmx.Print("done") }}
//
//	subj.Sub(a)
//	subj.Sub(b)
//
//	subj.Next(1)
//
//	subA.Unsub()
//
//	recC, c := obs.ObserverImpl[int]{Next: func(v int) { fmx.Print(i) }, Complete: func() { fmx.Print("done") }}
//	subj.Sub(c)
//
//	subj.Next(2)
//	subj.Complete()
//
// Output (conceptual):
//
//	A: 1
//	B: 1 2 done
//	C: 2 done
func ManySubject[T any](max int) Subject[T] {
	return subjectAdapter[T]{observables.ManySubject[T](max)}
}

// MonoSubject returns a Subject that allows at most one active subscriber.
//
// A MonoSubject behaves like a Subject with strict single-subscriber semantics.
// Only one observer may subscribe at a time. Any attempt to subscribe while an
// active subscription exists results in an error.
//
// The subject does not buffer values. Emissions are delivered only to the
// currently subscribed observer, if any.
//
// Termination semantics:
//   - Error or Complete terminates the subject permanently.
//   - If a subscriber is present at termination, the terminal signal is
//     delivered immediately.
//   - If no subscriber is present, the terminal states is recorded.
//   - Late subscribers receive the terminal signal immediately.
//
// After termination, further Next calls are ignored.
//
// Error handling:
//   - Error terminates the subject and is delivered to the current observer
//     if present.
//   - Late subscribers receive the stored error.
//
// Completion:
//   - Complete terminates the subject and is delivered to the current observer
//     if present.
//   - Late subscribers receive completion.
//
// MonoSubject is synchronous and does not introduce goroutines or buffering.
// It uses direct (non-slot) subscriptions and supports immediate unsubscribe
// semantics.
//
// This Subject is not safe for concurrent use unless externally synchronized.
//
// Diagram:
//
//	subscribe A
//	    |
//	    v
//	    |  Next(1)
//	    v
//	    1
//	    |
//	    |  Complete()
//	    v
//	    |
//
//	A:  1--|
//
// Example:
//
//	subj := obs.MonoSubject[int]()
//
//	subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println(v) },
//		func(err error) { fmt.Println("error:", err) },
//		func() { fmt.Println("done") },
//	))
//
//	subj.Next(1)
//	subj.Complete()
//
//	// Late subscription
//	subj.Sub(obs.NewObserver(
//		nil,
//		nil,
//		func() { fmt.Println("late done") },
//	))
//
// Output:
//
//	1
//	done
//	late done
func MonoSubject[T any]() Subject[T] {
	return subjectAdapter[T]{observables.MonoSubject[T]()}
}

// ReplaySubject returns a Subject that replays previously emitted values
// to new subscribers.
//
// A ReplaySubject buffers values emitted via Next and replays them
// synchronously to each new subscriber, in the original emission order,
// before forwarding any subsequent live emissions.
//
// Buffering behavior:
//   - If called with no arguments, the ReplaySubject buffers all values
//     (unbounded replay).
//   - If called with a single integer n > 0, the ReplaySubject buffers
//     only the last n values.
//   - At the buffer reaches capacity, the oldest values are discarded.
//
// Termination semantics:
//   - On Error, the subject terminates, broadcasts the error to all
//     current subscribers, and stops accepting new values.
//   - On Complete, the subject terminates, broadcasts completion, and
//     stops accepting new values.
//   - Late subscribers receive the buffered values first, then the
//     terminal signal (Error or Complete).
//
// After termination, further Next calls are ignored.
//
// ReplaySubject multicasts to all active subscribers and uses slot-based
// subscriptions, allowing safe reentrancy during emission.
//
// Diagram (capacity = 2):
//
//	source:     Next(1)  Next(2)  Next(3)  Complete
//	               |       |       |        |
//	buffer:        [1]    [1 2]    [2 3]     [2 3]
//
//	subscribe A ------------------------------|
//	A:            2--3--|
//
//	subscribe B --------------------|
//	B:            2--3--|
//
// Example:
//
//	subj := obs.ReplaySubject
//
//	subj.Next(1)
//	subj.Next(2)
//	subj.Next(3)
//
//	subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println("A:", v) },
//		nil,
//		func() { fmt.Println("A: done") },
//	))
//
//	subj.Complete()
//
//	subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println("B:", v) },
//		nil,
//		func() { fmt.Println("B: done") },
//	))
//
// Output:
//
//	A: 2
//	A: 3
//	A: done
//	B: 2
//	B: 3
//	B: done
func ReplaySubject[T any](n ...int) Subject[T] {
	return subjectAdapter[T]{observables.ReplaySubject[T](n...)}
}

// SlotSubject returns a Subject that multicasts values to all active subscribers.
//
// SlotSubject is a hot, synchronous Subject that allows multiple observers
// to subscribe and receive values emitted via Next.
//
// Key characteristics:
//   - Multicasting: Range value emitted via Next is broadcast to all
//     currently subscribed observers.
//   - Hot: The Subject is a live producer; subscriptions only receive
//     values emitted after they subscribe.
//   - Slot-based subscriptions: Range subscriber receives its own slot,
//     allowing safe reentrant calls to Next, Error, or Complete.
//
// Termination semantics:
//   - On Error, the Subject terminates immediately, broadcasts the error
//     to all active subscribers, and stops accepting new values.
//   - On Complete, the Subject terminates immediately, broadcasts
//     completion, and stops accepting new values.
//   - After termination, further Next calls are ignored.
//   - Late subscribers receive the terminal signal immediately upon
//     subscribing (Error or Complete).
//
// Unsubscription:
//   - Individual subscriptions can be unsubscribed using the returned
//     Subscription.
//   - UnsubAll closes all active subscriptions and clears the infra slot list.
//
// Diagram:
//
//	SlotSubject:
//	source:    Next(1)  Next(2)  Complete
//	              |       |       |
//	subscribers:  A       B      A,B receive Complete
//
// Example:
//
//	subj := obs.SlotSubject[int]()
//
//	subj.Sub(obs.NewObserver(
//	    func(v int) { fmt.Println("A:", v) },
//	    nil,
//	    func() { fmt.Println("A: done") },
//	))
//
//	subj.Next(1)
//	subj.Next(2)
//
//	subj.Sub(obs.NewObserver(
//	    func(v int) { fmt.Println("B:", v) },
//	    nil,
//	    func() { fmt.Println("B: done") },
//	))
//
//	subj.Complete()
//
// Output:
//
//	A: 1
//	A: 2
//	B: (receives nothing; joined after 2)
//	A: done
//	B: done
//
// Notes:
//
//   - SlotSubject is suitable when you want a simple hot Subject that
//     multicasts values to all active subscribers without buffering.
//   - For replaying past values to late subscribers, use ReplaySubject
//     instead.
func SlotSubject[T any]() Subject[T] {
	return subjectAdapter[T]{observables.SlotSubject[T]()}
}
