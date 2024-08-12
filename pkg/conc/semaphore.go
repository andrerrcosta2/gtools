// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import "sync"

// NewChannelSemaphore returns a new semaphore that can be used to limit the number of concurrent operations.
// The semaphore is initialized with a buffer of size maxConcurrent, allowing up to maxConcurrent operations to proceed concurrently.
func NewChannelSemaphore(maxConcurrent int) *ChannelSemaphore {
	// Create a new semaphore with a buffered channel of size maxConcurrent
	return &ChannelSemaphore{
		// The channel is used to track the number of available slots
		ch: make(chan struct{}, maxConcurrent),
	}
}

type ChannelSemaphore struct {
	ch chan struct{}
}

// Acq acquires the semaphore, blocking if the semaphore is at capacity.
// This method will block until the semaphore has available slots.
func (s *ChannelSemaphore) Acq() {
	// Send a signal to the channel to acquire the semaphore
	s.ch <- struct{}{}
}

// Rls releases a semaphore, allowing another operation to proceed.
// It blocks until a slot is available in the semaphore's buffer.
func (s *ChannelSemaphore) Rls() {
	// Receive from the channel to release a slot
	<-s.ch
}

// NewCountingSemaphore returns a new counting semaphore with the given initial count.
// The returned semaphore is ready to use and has a condition variable associated with it.
func NewCountingSemaphore(initial int) *CountingSemaphore {
	// Create a new counting semaphore with the given initial count
	s := &CountingSemaphore{
		count: initial,
	}
	// Initialize the condition variable with the semaphore's mutex
	s.cond = sync.NewCond(&s.mu)
	return s
}

type CountingSemaphore struct {
	count int
	mu    sync.Mutex
	cond  *sync.Cond
}

func (s *CountingSemaphore) Acq() {
	s.mu.Lock()
	for s.count <= 0 {
		s.cond.Wait()
	}
	s.count--
	s.mu.Unlock()
}

func (s *CountingSemaphore) Rls() {
	s.mu.Lock()
	s.count++
	s.cond.Signal()
	s.mu.Unlock()
}

// NewReadWriteSemaphore returns a new read-write semaphore.
// The returned semaphore is ready to use and has condition variables associated with it for read and write operations.
func NewReadWriteSemaphore() *ReadWriteSemaphore {
	// Create a new read-write semaphore
	rw := &ReadWriteSemaphore{}
	// Initialize the condition variable for read operations with the semaphore's mutex
	rw.read = sync.NewCond(&rw.mu)
	// Initialize the condition variable for write operations with the semaphore's mutex
	rw.write = sync.NewCond(&rw.mu)
	return rw
}

type ReadWriteSemaphore struct {
	readers int
	writer  bool
	mu      sync.Mutex
	read    *sync.Cond
	write   *sync.Cond
}

func (rw *ReadWriteSemaphore) StartR() {
	rw.mu.Lock()
	for rw.writer {
		rw.read.Wait()
	}
	rw.readers++
	rw.mu.Unlock()
}

func (rw *ReadWriteSemaphore) EndR() {
	rw.mu.Lock()
	rw.readers--
	if rw.readers == 0 {
		rw.write.Signal()
	}
	rw.mu.Unlock()
}

func (rw *ReadWriteSemaphore) StartW() {
	rw.mu.Lock()
	for rw.readers > 0 || rw.writer {
		rw.write.Wait()
	}
	rw.writer = true
	rw.mu.Unlock()
}

func (rw *ReadWriteSemaphore) EndW() {
	rw.mu.Lock()
	rw.writer = false
	rw.read.Signal()
	rw.write.Signal()
	rw.mu.Unlock()
}
