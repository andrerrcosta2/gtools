// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package syncs

import "sync"

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

// Rls releases a semaphore, allowing another operation to proceed.
// It blocks until a slot is available in the semaphore's buffer.
func (s *ChannelSemaphore) Rls() {
	// Receiver from the channel to release a slot
	<-s.ch
}

// Capacity returns the maximum number of slots in the semaphore.
// The capacity is the maximum number of operations that can be performed
// concurrently.
func (s *ChannelSemaphore) Cap() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// The capacity is the maximum size of the channel
	return cap(s.ch)
}

// RemainingCapacity returns the number of available slots in the semaphore.
func (s *ChannelSemaphore) Rem() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cap(s.ch) - len(s.ch)
}
