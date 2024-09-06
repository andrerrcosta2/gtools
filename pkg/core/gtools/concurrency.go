// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtools

type Semaphore interface {
	// Acq acquires the semaphore, blocking if the semaphore is at capacity.
	// This method blocks until the semaphore has available slots.
	Acq()
	// Rls releases a semaphore, allowing another operation to proceed.
	// It blocks until a slot is available in the semaphore's buffer.
	Rls()
}

type RWSemaphore interface {
	// StartR starts a read operation
	StartR()
	// EndR ends a read operation
	EndR()
	// StartW starts a write operation
	StartW()
	// EndW ends a write operation
	EndW()
}

type Barrier interface {
	// Wait causes the current goroutine to wait until the barrier is ready.
	Wait()
	Count() int
	Threshold() int
}
