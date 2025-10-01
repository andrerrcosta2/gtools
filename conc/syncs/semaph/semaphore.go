// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package semaph

import "sync"

// SingleRoutine creates a new SingleThreadedSemaphore.
func SingleRoutine() *SingleThreadedSemaphore {
	return &SingleThreadedSemaphore{}
}

// SingleThreadedSemaphore is a semaphore that works on a single thread of execution.
// It's a simple but effective way of limiting the number of ops that can be
// performed concurrently.
//
// The semaphore is implemented using a mutex and a boolean flag. The Acq method
// blocks until the flag is reset to false, the Rls method resets the flag.
// Its capacity is always 1, and the RemainingCapacity method
// returns 0 or 1 depending on whether the semaphore is currently acquired.
type SingleThreadedSemaphore struct {
	mu     sync.Mutex
	locked bool
}

// Acq acquires the semaphore, blocking if the semaphore is at capacity.
// This method will block until the semaphore has available slots.
func (s *SingleThreadedSemaphore) Acq() {
	s.mu.Lock()
	s.locked = true
}

// Rls releases a semaphore, allowing another op to proceed.
// It blocks until a slot is available in the semaphore's buffer.
func (s *SingleThreadedSemaphore) Rls() {
	s.mu.Unlock()
	s.locked = false
}

// Cap returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of ops that can be performed
// concurrently.
func (s *SingleThreadedSemaphore) Cap() int {
	return 1
}

// Rem returns the remaining capacity of the semaphore.
// The remaining capacity is the maximum number of ops that can still be performed
// concurrently.
func (s *SingleThreadedSemaphore) Rem() int {
	if s.locked {
		return 0
	}
	return 1
}

// Channel returns a new semaphore that can be used to limit the number of concurrent ops.
// The semaphore is initialized with a buffer of size maxConcurrent, allowing up to maxConcurrent ops to proceed concurrently.
func Channel(maxConcurrent int) *ChannelSemaphore {
	// Create a new semaphore with a buffered channel of size maxConcurrent
	return &ChannelSemaphore{
		// The channel is used to track the number of available slots
		ch: make(chan struct{}, maxConcurrent),
	}
}

type ChannelSemaphore struct {
	mu sync.RWMutex
	ch chan struct{}
}

// Acq acquires the semaphore, blocking if the semaphore is at capacity.
// This method will block until the semaphore has available slots.
func (s *ChannelSemaphore) Acq() {
	// Send a signal to the channel to acquire the semaphore
	s.ch <- struct{}{}
}

// Rls releases a semaphore, allowing another op to proceed.
// It blocks until a slot is available in the semaphore's buffer.
func (s *ChannelSemaphore) Rls() {
	// Receiver from the channel to release a slot
	<-s.ch
}

// Cap returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of ops that can be performed
// concurrently.
func (s *ChannelSemaphore) Cap() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// The capacity is the maximum size of the channel
	return cap(s.ch)
}

// Rem returns the number of available slots in the semaphore.
func (s *ChannelSemaphore) Rem() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cap(s.ch) - len(s.ch)
}

// Counter returns a new counting semaphore with the given initial count.
// The returned semaphore is ready to use and has a condition variable associated with it.
func Counter(initial int) *CounterSemaphore {
	// Create a new counting semaphore with the given initial count
	s := &CounterSemaphore{
		count: initial,
	}
	// Initialize the condition variable with the semaphore's mutex
	s.cond = sync.NewCond(&s.mu)
	return s
}

type CounterSemaphore struct {
	count int
	mu    sync.Mutex
	cond  *sync.Cond
}

func (s *CounterSemaphore) Acq() {
	s.mu.Lock()
	for s.count <= 0 {
		s.cond.Wait()
	}
	s.count--
	s.mu.Unlock()
}

func (s *CounterSemaphore) Rls() {
	s.mu.Lock()
	s.count++
	s.cond.Signal()
	s.mu.Unlock()
}

// Cap returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of ops that can be performed
// concurrently.
func (s *CounterSemaphore) Cap() int {
	// The capacity is the initial count
	return s.count
}

// Rem returns the remaining capacity of the semaphore.
// The remaining capacity is the maximum number of ops that can still be performed
// concurrently.
func (s *CounterSemaphore) Rem() int {
	return s.count
}

func Unbounded() *UnboundedSemaphore {
	return &UnboundedSemaphore{}
}

// UnboundedSemaphore represents a semaphore with no capacity limit.
// It can be useful on interface-bounded ops.
type UnboundedSemaphore struct{}

// Acq simulates acquiring a semaphore. Since it's unbounded, it does nothing.
func (s *UnboundedSemaphore) Acq() {
	// No op, as an unbounded semaphore never blocks.
}

// Rls simulates releasing a semaphore. Since it's unbounded, it does nothing.
func (s *UnboundedSemaphore) Rls() {
	// No op, as an unbounded semaphore doesn't track releases.
}

// Cap always returns -1 for an unbounded semaphore, indicating no limit.
func (s *UnboundedSemaphore) Cap() int {
	return -1
}

// Rem always returns -1, as capacity is unlimited.
func (s *UnboundedSemaphore) Rem() int {
	return -1
}

// ReadWrite returns a new read-write semaphore.
// The returned semaphore is ready to use and has condition variables associated with it for read and write ops.
func ReadWrite(maxWriters, maxReaders int) *ReadWriteSemaphore {
	// Create a new read-write semaphore
	rw := &ReadWriteSemaphore{
		maxReaders: maxReaders,
		maxWriters: maxWriters,
	}
	// Initialize the condition variable for read ops with the semaphore's mutex
	rw.read = sync.NewCond(&rw.mtx)
	// Initialize the condition variable for write ops with the semaphore's mutex
	rw.write = sync.NewCond(&rw.mtx)
	return rw
}

type ReadWriteSemaphore struct {
	mtx            sync.Mutex
	read           *sync.Cond
	write          *sync.Cond
	readers        int
	maxReaders     int
	maxWriters     int
	currentWriters int
	writer         bool
}

func (rw *ReadWriteSemaphore) StartR() {
	rw.mtx.Lock()
	for rw.writer || rw.readers >= rw.maxReaders {
		rw.read.Wait() // SetWaitingPoint if a writer is active or readers have hit max capacity
	}
	rw.readers++
	rw.mtx.Unlock()
}

func (rw *ReadWriteSemaphore) EndR() {
	rw.mtx.Lock()
	rw.readers--
	if rw.readers == 0 {
		rw.write.Signal() // Signal writers if no readers are left
	}
	rw.mtx.Unlock()
}

func (rw *ReadWriteSemaphore) StartW() {
	rw.mtx.Lock()
	for rw.readers > 0 || rw.writer || rw.currentWriters >= rw.maxWriters {
		rw.write.Wait() // SetWaitingPoint if readers are active or writers have hit max capacity
	}
	rw.writer = true
	rw.currentWriters++
	rw.mtx.Unlock()
}

func (rw *ReadWriteSemaphore) EndW() {
	rw.mtx.Lock()
	rw.writer = false
	rw.currentWriters--
	rw.read.Signal()  // Signal readers
	rw.write.Signal() // Signal other writers if any
	rw.mtx.Unlock()
}

// Capacity returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of ops that can be performed
// concurrently.
func (rw *ReadWriteSemaphore) Capacity() (writers int, readers int) {
	rw.mtx.Lock()
	defer rw.mtx.Unlock()
	return rw.maxWriters, rw.maxReaders
}

// RemainingCapacity returns the remaining capacity of the semaphore.
// The remaining capacity is the maximum number of ops that can still be performed
// concurrently.
func (rw *ReadWriteSemaphore) RemainingCapacity() (writers int, readers int) {
	rw.mtx.Lock()
	defer rw.mtx.Unlock()
	writers = rw.maxWriters - rw.currentWriters
	readers = rw.maxReaders - rw.readers
	return
}
