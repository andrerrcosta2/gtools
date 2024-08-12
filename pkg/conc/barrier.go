// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"sync"
)

// ThresholdBarrier returns a new barrier that can be used to synchronize goroutines.
// The barrier is initialized with a threshold of n, meaning that n goroutines must call Wait() before any of them can proceed.
func ThresholdBarrier(n int) *Barrier {
	// Create a new barrier with the given threshold
	b := &Barrier{
		// The threshold is the number of goroutines that must wait before any can proceed
		threshold: n,
		// The count is initialized to 0, indicating that no goroutines have yet reached the barrier
		count: 0,
	}
	// Create a new condition variable for the barrier, using the barrier's mutex
	b.cond = sync.NewCond(&b.mtx)
	return b
}

type Barrier struct {
	count     int
	threshold int
	mtx       sync.Mutex
	cond      *sync.Cond
}

// Wait causes the current goroutine to wait until the barrier is ready.
// When the barrier is ready, all waiting goroutines are released.
func (b *Barrier) Wait() {
	// Acquire the mutex to ensure exclusive access to the barrier's state
	b.mtx.Lock()
	defer b.mtx.Unlock()

	// Increment the count of waiting goroutines
	b.count++

	// If the threshold has not been reached, wait for the barrier to be released
	if b.count < b.threshold {
		// Wait for the condition variable to be signaled
		b.cond.Wait()
	} else {
		// If the threshold has been reached, release all waiting goroutines
		b.cond.Broadcast()
		// Reset the count for the next use
		b.count = 0
	}
}

func (b *Barrier) Count() int {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	return b.count
}

func (b *Barrier) Threshold() int {
	return b.threshold
}

var _ gtools.Barrier = (*Barrier)(nil)

// StepBroadcast returns a new StepBroadcastBarrier instance with the given proceed function.
// The proceed function is used to determine whether to broadcast.
//
// Args:
//
//	proceed (functions.Function[gtools.Barrier, bool]): The proceed function to be used determine if it should broadcast.
//
// Returns:
//
//	*StepBroadcastBarrier: A new StepBroadcastBarrier instance.
func StepBroadcast(proceed functions.Function[gtools.Barrier, bool]) *StepBroadcastBarrier {
	// Create a new StepBroadcastBarrier instance with the given threshold and proceed function.
	// Initialize the count to 0, indicating no goroutines have reached the barrier.
	b := &StepBroadcastBarrier{
		count:   0,       // Initial count of goroutines at the barrier.
		proceed: proceed, // Proceed function to be used when the threshold is reached.
	}
	// Create a new condition variable for the StepBroadcastBarrier instance.
	// This is used to synchronize goroutines waiting at the barrier.
	b.cond = sync.NewCond(&b.mtx)
	return b // Return the newly created StepBroadcastBarrier instance.
}

type StepBroadcastBarrier struct {
	count   int
	mtx     sync.Mutex
	cond    *sync.Cond
	proceed functions.Function[gtools.Barrier, bool] // Function to determine when to broadcast
}

func (b *StepBroadcastBarrier) Wait() {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.count++
	if !b.proceed(b) { // Use the custom function to decide whether to broadcast
		b.cond.Wait()
	} else {
		b.cond.Broadcast() // Release all waiting threads if the condition is met
		b.count = 0        // Reset count for reuse
	}
}

func (b *StepBroadcastBarrier) Count() int {
	return b.count
}

func (b *StepBroadcastBarrier) Threshold() int {
	return 0
}

var _ gtools.Barrier = (*StepBroadcastBarrier)(nil)
