// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observables

import (
	"context"
	"errors"
	"sync"

	"github.com/andrerrcosta2/gtools/gflux/core/lifecycle/gates"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

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
	return &asyncSubj[T]{
		subj: &slotSubj[T]{slots: []*subscr.Slot[T]{}},
	}
}

type asyncSubj[T any] struct {
	subj        *slotSubj[T]
	last        T
	hasValue    bool
	terminated  bool
	terminalErr error
}

func (s *asyncSubj[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	if s.terminated {
		if s.terminalErr != nil {
			observers.Error(o, s.terminalErr)
		} else {
			if s.hasValue {
				observers.Next(o, s.last)
			}
			o.Complete()
		}
		return &subscr.Noop{}, nil
	}
	return s.subj.Sub(o)
}

func (s *asyncSubj[T]) Next(v T) {
	if s.terminated {
		return
	}
	s.last = v
	s.hasValue = true
}

func (s *asyncSubj[T]) Error(err error) {
	if s.terminated {
		return
	}
	s.terminated = true
	s.terminalErr = err
	s.subj.Error(err)
}

func (s *asyncSubj[T]) Complete() {
	if s.terminated {
		return
	}
	s.terminated = true

	if s.hasValue {
		s.subj.Next(s.last)
	}
	s.subj.Complete()
}

func (s *asyncSubj[T]) UnsubAll() {
	s.subj.UnsubAll()
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
	s := &behaviorSubj[T]{
		subj:     &slotSubj[T]{slots: []*subscr.Slot[T]{}},
		value:    initial,
		hasValue: true,
	}
	return s
}

type behaviorSubj[T any] struct {
	subj        *slotSubj[T]
	value       T
	hasValue    bool
	terminalErr error
	terminated  bool
}

func (s *behaviorSubj[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	if s.terminated {
		if s.hasValue {
			o.Next(s.value)
		}
		if s.terminalErr != nil {
			o.Error(s.terminalErr)
		} else {
			o.Complete()
		}
		return &subscr.Noop{}, nil
	}

	sub, err := s.subj.Sub(o)
	if err != nil {
		return nil, err
	}

	if s.hasValue {
		o.Next(s.value)
	}

	return sub, nil
}

func (s *behaviorSubj[T]) Next(v T) {
	if s.terminated {
		return
	}
	s.value = v
	s.hasValue = true
	s.subj.Next(v)
}

func (s *behaviorSubj[T]) Error(err error) {
	if s.terminated {
		return
	}
	s.terminated = true
	s.terminalErr = err
	s.subj.Error(err)
}

func (s *behaviorSubj[T]) Complete() {
	if s.terminated {
		return
	}
	s.terminated = true
	s.subj.Complete()
}

func (s *behaviorSubj[T]) UnsubAll() {
	s.subj.UnsubAll()
}

// ConcSubject returns a concurrent-safe Subject.
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
//	subj := obs.ConcSubject[int]()
//
//	go subj.Next(1)
//	go subj.Error(errors.New("fail"))
//
//	// Outcome is deterministic in safety, but not in ordering.
func ConcSubject[T any]() Subject[T] {
	return &concSubj[T]{
		slotSubj: slotSubj[T]{slots: []*subscr.Slot[T]{}},
	}
}

type concSubj[T any] struct {
	slotSubj[T]
	mu  sync.Mutex
	ctx context.Context
}

func (s *concSubj[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	if s.terminated {
		if s.terminalErr != nil {
			observers.Error(o, s.terminalErr)
		} else {
			observers.Complete(o)
		}
		return &subscr.Noop{}, nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	return s.sub(o)
}

func (s *concSubj[T]) Next(t T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next(t)
}

func (s *concSubj[T]) Error(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.error(err)
}

func (s *concSubj[T]) Complete() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.complete()
}

func (s *concSubj[T]) UnsubAll() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.unsubAll()
}

// DirectSubject returns a Subject that directly multicasts to all subscribers.
//
// DirectSubject is a low-level Subject that maintains a direct map of
// subscriptions to observers. Range Next call is immediately delivered
// to all active observers. Terminal signals (Error or Complete) are
// broadcast to all subscribers.
//
// Key behaviors:
//
//   - Multicast: each emission goes to all active subscribers.
//   - Cold semantics: not applicable; subscribers start receiving values
//     immediately after subscription.
//   - Terminal signals: Error or Complete terminates the subject; late
//     subscribers receive the terminal signal immediately.
//   - Slotless: DirectSubject does not use slot buffers; reentrant emissions
//     should be handled with care.
//
// Termination:
//
//   - After Error or Complete, further Next calls are ignored.
//   - Late subscribers receive the terminal signal (error or complete).
//
// Subscription management:
//
//   - Range subscription can be independently cancelled via Close.
//   - UnsubAll closes all active subscriptions and clears the infra map.
//
// Diagram:
//
//	subscribe A        subscribe B
//	    |                   |
//	    v                   v
//	    |  Next(1) -------->|
//	    |                   |
//	    |  Next(2) -------->|
//	    v                   v
//	A:  1--2
//	B:  1--2
//
// Example:
//
//	subj := obs.DirectSubject[int]()
//
//	subA := subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println("A:", v) },
//		nil,
//		func() { fmt.Println("A: done") },
//	))
//
//	subB := subj.Sub(obs.NewObserver(
//		func(v int) { fmt.Println("B:", v) },
//		nil,
//		func() { fmt.Println("B: done") },
//	))
//
//	subj.Next(1)
//	subj.Next(2)
//	subj.Complete()
//
//	subA.Close()
//	subB.Close()
//
// Output:
//
//	A: 1
//	B: 1
//	A: 2
//	B: 2
//	A: done
//	B: done
func DirectSubject[T any]() Subject[T] {
	return &diSubj[T]{
		obs: make(map[*subscr.DiSubscr[T]]*observers.Observer[T]),
	}
}

type diSubj[T any] struct {
	obs         map[*subscr.DiSubscr[T]]*observers.Observer[T]
	terminalErr error
	terminated  bool
}

func (s *diSubj[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	if s.terminated {
		if s.terminalErr != nil {
			observers.Error(o, s.terminalErr)
		} else {
			observers.Complete(o)
		}
		return &subscr.Noop{}, nil
	}

	sub := subscr.NewDiSub[T](s)
	s.obs[sub] = o
	return sub, nil
}

func (s *diSubj[T]) Next(t T) {
	if s.terminated {
		return
	}
	for _, o := range s.obs {
		observers.Next(o, t)
	}
}

func (s *diSubj[T]) Error(err error) {
	if s.terminated {
		return
	}
	s.terminated = true
	s.terminalErr = err
	for _, o := range s.obs {
		observers.Error(o, err)
	}
	s.UnsubAll()
}

func (s *diSubj[T]) Complete() {
	if s.terminated {
		return
	}
	s.terminated = true
	for _, o := range s.obs {
		observers.Complete(o)
	}
	s.UnsubAll()
}

func (s *diSubj[T]) UnsubAll() {
	for sub := range s.obs {
		sub.Close()
	}
	s.obs = make(map[*subscr.DiSubscr[T]]*observers.Observer[T])
}

func (s *diSubj[T]) RemoveSub(subscription subscr.Subscription) {
	delete(s.obs, subscription.(*subscr.DiSubscr[T]))
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
//	a := &obs.Observer[int]{Next: func(v int){ fmx.Print(i) }, Complete: func() { fmx.Print("done") }}
//	b := &obs.Observer[int]{Next: func(v int){ fmx.Print(i) }, Complete: func() { fmx.Print("done") }}
//
//	subj.Sub(a)
//	subj.Sub(b)
//
//	subj.Next(1)
//
//	subA.Unsub()
//
//	c := &obs.Observer[int]{Next: func(v int){ fmx.Print(i) }, Complete: func() { fmx.Print("done") }}
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
	if max <= 0 {
		panic("max must be > 0")
	}
	return &manySubj[T]{
		slotSubj: slotSubj[T]{
			slots: []*subscr.Slot[T]{},
		},
		capacity: max,
	}
}

var MaxSubsError = errors.New("maximum subscribers reached")

type manySubj[T any] struct {
	slotSubj[T]
	capacity int
}

func (s *manySubj[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	if s.terminated {
		if s.terminalErr != nil {
			observers.Error(o, s.terminalErr)
		} else {
			observers.Complete(o)
		}
		return &subscr.Noop{}, nil
	}

	// prevent delayed unsubscribed slots
	s.slots = subscr.CompactAliveSlots(s.slots)
	if len(s.slots) >= s.capacity {
		return nil, MaxSubsError
	}

	sub := subscr.NewSlotSub(o, gates.Atomic())
	s.slots = append(s.slots, sub.Slot())
	return sub, nil
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
	return &monoSubj[T]{}
}

type monoSubj[T any] struct {
	observer     *observers.Observer[T]
	subscription *subscr.DiSubscr[T]
	terminated   bool
	terminalErr  error
}

func (s *monoSubj[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	if s.terminated {
		if s.terminalErr != nil {
			observers.Error(o, s.terminalErr)
		} else {
			observers.Complete(o)
		}
		return &subscr.Noop{}, nil
	}

	if s.observer != nil {
		return nil, errors.New("mono subject allows only one subscriber")
	}

	s.subscription = subscr.NewDiSub[T](s)
	s.observer = o
	return s.subscription, nil
}

func (s *monoSubj[T]) Next(t T) {
	if s.terminated || s.observer == nil {
		return
	}
	observers.Next(s.observer, t)
}

func (s *monoSubj[T]) Error(err error) {
	if s.terminated {
		return
	}
	s.terminated = true
	s.terminalErr = err
	if s.observer != nil {
		observers.Error(s.observer, err)
	}
}

func (s *monoSubj[T]) Complete() {
	if s.terminated {
		return
	}
	s.terminated = true
	if s.observer != nil {
		observers.Complete(s.observer)
	}
}

func (s *monoSubj[T]) UnsubAll() {
	s.subscription.Close()
	s.observer = nil
	s.subscription = nil
}

func (s *monoSubj[T]) RemoveSub(_ subscr.Subscription) {
	s.subscription = nil
	s.observer = nil
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
	if len(n) == 0 {
		return &replaySubj[T]{
			subj:   &slotSubj[T]{slots: []*subscr.Slot[T]{}},
			buffer: []T{},
		}
	}
	return &replaySubj[T]{
		subj:     &slotSubj[T]{slots: []*subscr.Slot[T]{}},
		buffer:   make([]T, 0, n[0]),
		capacity: n[0],
	}
}

type replaySubj[T any] struct {
	subj        *slotSubj[T]
	buffer      []T
	capacity    int // 0 = unbounded
	terminalErr error
	terminated  bool
}

func (s *replaySubj[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	for _, v := range s.buffer {
		observers.Next(o, v)
	}

	if s.terminated {
		if s.terminalErr != nil {
			observers.Error(o, s.terminalErr)
		} else {
			observers.Complete(o)
		}
		return &subscr.Noop{}, nil
	}

	return s.subj.Sub(o)
}

func (s *replaySubj[T]) Next(v T) {
	if s.terminated {
		return
	}

	if s.capacity > 0 && len(s.buffer) == s.capacity {
		copy(s.buffer, s.buffer[1:])
		s.buffer[len(s.buffer)-1] = v
	} else {
		s.buffer = append(s.buffer, v)
	}

	s.subj.Next(v)
}

func (s *replaySubj[T]) Error(err error) {
	if s.terminated {
		return
	}
	s.terminated = true
	s.terminalErr = err
	s.subj.Error(err)
}

func (s *replaySubj[T]) Complete() {
	if s.terminated {
		return
	}
	s.terminated = true
	s.subj.Complete()
}

func (s *replaySubj[T]) UnsubAll() {
	s.subj.UnsubAll()
}

// SlotSubject returns the default Subject implementation.
//
// This slotSubj uses slot-based subscriptions with lazy cleanup performed at0
// emission boundaries. Unsubscribed observers are removed eventually, not
// immediately, favoring iteration performance and reentrancy safety.
func SlotSubject[T any]() Subject[T] {
	return &slotSubj[T]{
		slots: []*subscr.Slot[T]{},
	}
}

type slotSubj[T any] struct {
	slots       []*subscr.Slot[T]
	terminated  bool
	terminalErr error
}

func (s *slotSubj[T]) Sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	if s.terminated {
		if s.terminalErr != nil {
			observers.Error(o, s.terminalErr)
		} else {
			observers.Complete(o)
		}
		return &subscr.Noop{}, nil
	}
	return s.sub(o)
}

func (s *slotSubj[T]) sub(o *observers.Observer[T]) (subscr.Subscription, error) {
	sub := subscr.NewSlotSub(o, gates.Atomic())
	s.slots = append(s.slots, sub.Slot())
	return sub, nil
}

func (s *slotSubj[T]) Next(t T) {
	s.next(t)
}

func (s *slotSubj[T]) next(t T) {
	if s.terminated {
		return
	}

	dst := s.slots[:0]
	for _, slt := range s.slots {
		if !slt.Closed() {
			observers.Next(slt.Observer(), t)
			dst = append(dst, slt)
		}
	}
	s.slots = dst
}

func (s *slotSubj[T]) Error(err error) {
	s.error(err)
}

func (s *slotSubj[T]) error(err error) {
	if s.terminated {
		return
	}

	s.terminated = true
	s.terminalErr = err

	for _, slt := range s.slots {
		if !slt.Closed() {
			observers.Error(slt.Observer(), err)
			slt.Close()
		}
	}
	s.slots = nil
}

func (s *slotSubj[T]) Complete() {
	s.complete()
}

func (s *slotSubj[T]) complete() {
	if s.terminated {
		return
	}
	s.terminated = true
	for _, slt := range s.slots {
		if !slt.Closed() {
			observers.Complete(slt.Observer())
			slt.Close()
		}
	}

	s.slots = nil
}

func (s *slotSubj[T]) UnsubAll() {
	s.unsubAll()
}

func (s *slotSubj[T]) unsubAll() {
	for _, slt := range s.slots {
		slt.Close()
	}
	s.slots = nil
}
