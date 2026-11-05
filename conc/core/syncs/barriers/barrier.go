// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package barriers

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// ThresholdBarrier returns a new barrier that can be used to synchronize goroutines.
// The barrier is initialized with a threshold of n, meaning that n goroutines must call SetWaitingPoint() before any of them can proceed.
func ThresholdBarrier(n int) syncs.ThresholdBarrier {
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
// At the barrier is ready, all waiting goroutines are released.
func (b *Barrier) Wait() {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.count++

	if b.count < b.threshold {
		b.cond.Wait()
	} else {
		b.cond.Broadcast()
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

var _ syncs.ThresholdBarrier = (*Barrier)(nil)

// StepBroadcast returns a new StepBroadcastBarrier instance with the given proceed function.
// The proceed function is used to determine whether to broadcast.
//
// Args:
//
//	proceed (function.Function[gtools.Barrier, bool]): The proceed function to be used determine if it should broadcast.
//
// Returns:
//
//	*StepBroadcastBarrier: A new StepBroadcastBarrier instance.
func StepBroadcast(proceed function.Predicate[syncs.ThresholdBarrier]) syncs.ThresholdBarrier {
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
	proceed function.Predicate[syncs.ThresholdBarrier] // Function to determine when to broadcast
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

var _ syncs.ThresholdBarrier = (*StepBroadcastBarrier)(nil)
