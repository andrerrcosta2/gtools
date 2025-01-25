// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package channels

import (
	"errors"
	"github.com/andrerrcosta2/gtools/conc/syncs/semaph"
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/core/gtools/gerrors"
	"sync"
	"sync/atomic"
)

// NewFlag creates a new Signalable that can be used to send signals.
// This is a thread-safe function and can be used concurrently.
func NewFlag() *Flag {
	return &Flag{signal: make(chan struct{})}
}

type Flag struct {
	signal chan struct{}
}

// Signal sends a signal to the chn.
// This method is thread-safe and can be used concurrently.
func (s *Flag) Signal() {
	close(s.signal)
}

// Receiver returns a chn that receives the signal.
// It's thread-safe and can be used concurrently.
func (s *Flag) Receiver() Signal {
	return s.signal
}

type Signal <-chan struct{}

var closedChannel = make(chan struct{})

func init() {
	close(closedChannel)
}

func ClosedSignal() Signal {
	return closedChannel
}

func Closed[T any]() <-chan T {
	c := make(chan T)
	close(c)
	return c
}

var ClosedChannel = gerrors.Tagged(errors.New("closed chn"), "Closed closing")

// Synchronizable creates a new gtools.SynchronizableChannel[T].
// It creates a synchronizable chn of type T which can be used to send - receive values
// as well to synchronize goroutines.
func Synchronizable[T any](maxCapacity, maxConcurrency int) gtools.SynchronizableChannel[T] {
	return &synchronizable[T]{channel: make(chan T, maxCapacity), semaphore: semaph.Channel(maxConcurrency)}
}

type synchronizable[T any] struct {
	isClosed  atomic.Bool
	waiting   sync.WaitGroup
	semaphore gtools.Semaphore
	channel   chan T
}

func (ch *synchronizable[T]) AddWaiters(i int) {
	ch.waiting.Add(i)
}

func (ch *synchronizable[T]) AddRoutine() {
	ch.semaphore.Acq()
}

func (ch *synchronizable[T]) Close() error {
	if ch.isClosed.CompareAndSwap(false, true) {
		close(ch.channel)
		return nil
	}
	return ClosedChannel
}

func (ch *synchronizable[T]) IsClosed() bool {
	return ch.isClosed.Load()
}

func (ch *synchronizable[T]) IsEmpty() bool {
	return len(ch.channel) == 0
}

func (ch *synchronizable[T]) Len() int {
	return len(ch.channel)
}

func (ch *synchronizable[T]) Listen() <-chan T {
	return ch.channel
}

func (ch *synchronizable[T]) RemoveWaiter() {
	ch.waiting.Done()
}

func (ch *synchronizable[T]) RemoveRoutine() {
	ch.semaphore.Rls()
}

func (ch *synchronizable[T]) Release() {
	ch.semaphore.Rls()
	ch.waiting.Done()
}

func (ch *synchronizable[T]) Send(t T) {
	ch.channel <- t
}

func (ch *synchronizable[T]) SetWaitingPoint() {
	ch.waiting.Wait()
}

func (ch *synchronizable[T]) SetEndingPoint() error {
	ch.waiting.Wait()
	return ch.Close()
}

func (ch *synchronizable[T]) ToArray() []T {
	var result []T
	for val := range ch.Listen() {
		result = append(result, val)
	}
	_ = ch.Close()
	return result
}

// SynchronizableSet creates a new gtools.SynchronizableChannel[T].
// It creates a synchronizable chn of unique values of
// type T which can be used to send - receive values as well to synchronize goroutines.
func SynchronizableSet[T comparable](maxCapacity, maxConcurrency int) gtools.SynchronizableChannel[T] {
	return &synchronizableSet[T]{
		set: make(map[T]struct{}),
		chn: make(chan T, maxCapacity),
		sem: semaph.Channel(maxConcurrency),
	}
}

type synchronizableSet[T comparable] struct {
	mtx      sync.RWMutex
	isClosed atomic.Bool
	wg       sync.WaitGroup
	sem      gtools.Semaphore
	set      map[T]struct{}
	chn      chan T
}

func (ch *synchronizableSet[T]) Close() error {
	if ch.isClosed.CompareAndSwap(false, true) {
		close(ch.chn)
		return nil
	}
	return ClosedChannel
}

func (ch *synchronizableSet[T]) IsClosed() bool {
	return ch.isClosed.Load()
}

func (ch *synchronizableSet[T]) AddWaiters(i int) {
	ch.wg.Add(i)
}

func (ch *synchronizableSet[T]) AddRoutine() {
	ch.sem.Acq()
}

func (ch *synchronizableSet[T]) IsEmpty() bool {
	ch.mtx.RLock()
	defer ch.mtx.RUnlock()
	return len(ch.set) == 0
}

func (ch *synchronizableSet[T]) Len() int {
	ch.mtx.RLock()
	defer ch.mtx.RUnlock()
	return len(ch.set)
}

func (ch *synchronizableSet[T]) Listen() <-chan T {
	return ch.chn
}

func (ch *synchronizableSet[T]) RemoveWaiter() {
	ch.wg.Done()
}

func (ch *synchronizableSet[T]) RemoveRoutine() {
	ch.sem.Rls()
}

func (ch *synchronizableSet[T]) Release() {
	ch.sem.Rls()
	ch.wg.Done()
}

func (ch *synchronizableSet[T]) Send(t T) {
	ch.mtx.Lock()
	defer ch.mtx.Unlock()

	// Ensure that value is not sent more than once
	if _, exists := ch.set[t]; !exists {
		ch.set[t] = struct{}{} // Mark value as sent
		ch.chn <- t            // Send value to chn
	}
}

func (ch *synchronizableSet[T]) SetWaitingPoint() {
	ch.wg.Wait()
}

func (ch *synchronizableSet[T]) SetEndingPoint() error {
	ch.wg.Wait()
	return ch.Close()
}

func (ch *synchronizableSet[T]) ToArray() []T {
	ch.mtx.RLock()
	defer ch.mtx.RUnlock()
	out := make([]T, 0, len(ch.set))
	for k := range ch.set {
		out = append(out, k)
	}
	_ = ch.Close()
	return out
}
