// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package semaph

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums/ints"
)

// Serial creates a semaphore that only allows a serial processing (one thread per processing).
func Serial() syncs.Semaphore {
	return &singleThreaded{}
}

// singleThreaded is a semaphore that works on a single thread of execution.
// It's a simple but effective way of limiting the number of ops that can be
// performed concurrently.
//
// The semaphore is implemented using a mutex and a boolean flag. The Acq method
// blocks until the flag is reset to false, the Rls method resets the flag.
// Its capacity is always 1, and the Cap method
// returns 0 or 1 depending on whether the semaphore is currently acquired.
type singleThreaded struct {
	mu     sync.Mutex
	locked bool
}

// Acq acquires the semaphore, blocking if the semaphore is at capacity.
// This method will block until the semaphore has available slots.
func (s *singleThreaded) Acq() {
	s.mu.Lock()
	s.locked = true
}

// Rls releases a semaphore, allowing another op to proceed.
// It blocks until a slot is available in the semaphore's buffer.
func (s *singleThreaded) Rls() {
	s.mu.Unlock()
	s.locked = false
}

// Cap returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of ops that can be performed
// concurrently.
func (s *singleThreaded) Cap() int {
	return 1
}

// Len returns the remaining capacity of the semaphore.
// The remaining capacity is the maximum number of ops that can still be performed
// concurrently.
func (s *singleThreaded) Len() int {
	if s.locked {
		return 0
	}
	return 1
}

// Limiter returns a new semaphore that can be used to limit the number of concurrent ops.
// The semaphore is initialized with a buffer of size maxConcurrent, allowing up to maxConcurrent ops to proceed concurrently.
func Limiter(maxConcurrent int) syncs.Semaphore {
	// Create a new semaphore with a buffered channel of size maxConcurrent
	return &limiter{
		// The channel is used to track the number of available slots
		ch: make(chan struct{}, maxConcurrent),
	}
}

type limiter struct {
	mu sync.RWMutex
	ch chan struct{}
}

// Acq acquires the semaphore, blocking if the semaphore is at capacity.
// This method will block until the semaphore has available slots.
func (s *limiter) Acq() {
	// Send a signal to the channel to acquire the semaphore
	s.ch <- struct{}{}
}

// Rls releases a semaphore, allowing another op to proceed.
// It blocks until a slot is available in the semaphore's buffer.
func (s *limiter) Rls() {
	// Receiver from the channel to release a slot
	<-s.ch
}

// Cap returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of ops that can be performed
// concurrently.
func (s *limiter) Cap() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// The capacity is the maximum size of the channel
	return cap(s.ch)
}

// Len returns the number of available slots in the semaphore.
func (s *limiter) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cap(s.ch) - len(s.ch)
}

// Counter returns a new counting semaphore with the given initial count.
// The returned semaphore is ready to use and has a condition variable associated with it.
func Counter(initial int) syncs.Semaphore {
	// Create a new counting semaphore with the given initial count
	s := &counter{
		count: initial,
	}
	// Initialize the condition variable with the semaphore's mutex
	s.cond = sync.NewCond(&s.mu)
	return s
}

type counter struct {
	count int
	mu    sync.Mutex
	cond  *sync.Cond
}

func (s *counter) Acq() {
	s.mu.Lock()
	for s.count <= 0 {
		s.cond.Wait()
	}
	s.count--
	s.mu.Unlock()
}

func (s *counter) Rls() {
	s.mu.Lock()
	s.count++
	s.cond.Signal()
	s.mu.Unlock()
}

// Cap returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of ops that can be performed
// concurrently.
func (s *counter) Cap() int {
	// The capacity is the initial count
	return s.count
}

// Len returns the remaining capacity of the semaphore.
// The remaining capacity is the maximum number of ops that can still be performed
// concurrently.
func (s *counter) Len() int {
	return s.count
}

func Unbounded() syncs.Semaphore {
	return &unbounded{}
}

// unbounded represents a semaphore with no capacity limit.
// It can be useful on interface-bounded ops.
type unbounded struct{}

// Acq simulates acquiring a semaphore. Since it's unbounded, it does nothing.
func (s *unbounded) Acq() {
	// No op, as an unbounded semaphore never blocks.
}

// Rls simulates releasing a semaphore. Since it's unbounded, it does nothing.
func (s *unbounded) Rls() {
	// No op, as an unbounded semaphore doesn't track releases.
}

// Cap always returns ints.Max for an unbounded semaphore, indicating no limit.
func (s *unbounded) Cap() int {
	return ints.Max
}

// Len always returns ints.Max, as capacity is unlimited.
func (s *unbounded) Len() int {
	return ints.Max
}

// ReadWrite returns a new read-write semaphore.
// The returned semaphore is ready to use and has condition variables associated with it for read and write ops.
func ReadWrite(maxWriters, maxReaders int) syncs.RWSemaphore {
	// Create a new read-write semaphore
	rw := &rwsmp{
		maxReaders: maxReaders,
		maxWriters: maxWriters,
	}
	// Initialize the condition variable for read ops with the semaphore's mutex
	rw.read = sync.NewCond(&rw.mtx)
	// Initialize the condition variable for write ops with the semaphore's mutex
	rw.write = sync.NewCond(&rw.mtx)
	return rw
}

type rwsmp struct {
	mtx            sync.Mutex
	read           *sync.Cond
	write          *sync.Cond
	readers        int
	maxReaders     int
	maxWriters     int
	currentWriters int
	writer         bool
}

func (rw *rwsmp) StartR() {
	rw.mtx.Lock()
	for rw.writer || rw.readers >= rw.maxReaders {
		rw.read.Wait() // SetWaitingPoint if a writer is active or readers have hit max capacity
	}
	rw.readers++
	rw.mtx.Unlock()
}

func (rw *rwsmp) EndR() {
	rw.mtx.Lock()
	rw.readers--
	if rw.readers == 0 {
		rw.write.Signal() // Wait writers if no readers are left
	}
	rw.mtx.Unlock()
}

func (rw *rwsmp) StartW() {
	rw.mtx.Lock()
	for rw.readers > 0 || rw.writer || rw.currentWriters >= rw.maxWriters {
		rw.write.Wait() // SetWaitingPoint if readers are active or writers have hit max capacity
	}
	rw.writer = true
	rw.currentWriters++
	rw.mtx.Unlock()
}

func (rw *rwsmp) EndW() {
	rw.mtx.Lock()
	rw.writer = false
	rw.currentWriters--
	rw.read.Signal()  // Wait readers
	rw.write.Signal() // Wait other writers if any
	rw.mtx.Unlock()
}

// Len returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of ops that can be performed
// concurrently.
func (rw *rwsmp) Len() (writers int, readers int) {
	rw.mtx.Lock()
	defer rw.mtx.Unlock()
	return rw.maxWriters, rw.maxReaders
}

// Cap returns the remaining capacity of the semaphore.
// The remaining capacity is the maximum number of ops that can still be performed
// concurrently.
func (rw *rwsmp) Cap() (writers int, readers int) {
	rw.mtx.Lock()
	defer rw.mtx.Unlock()
	writers = rw.maxWriters - rw.currentWriters
	readers = rw.maxReaders - rw.readers
	return
}
