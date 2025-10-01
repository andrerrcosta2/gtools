// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtools

import (
	"github.com/andrerrcosta2/gtools/core/domain/data"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
)

type Semaphore interface {
	// Acq acquires the semaphore, blocking if the semaphore is at capacity.
	// This method blocks until the semaphore has available slots.
	Acq()
	// Rls releases a semaphore, allowing another op to proceed.
	// It blocks until a slot is available in the semaphore's buffer.
	Rls()
	// Cap returns the maximum number of slots in the semaphore.
	//
	// The capacity is the maximum number of ops that can be performed
	// concurrently.
	Cap() int

	// Rem returns the remaining capacity of the semaphore.
	//
	// The remaining capacity is the maximum number of ops that can still be performed
	// concurrently.
	Rem() int
}

type RWSemaphore interface {
	// StartR starts a read op
	StartR()
	// EndR ends a read op
	EndR()
	// StartW starts a write op
	StartW()
	// EndW ends a write op
	EndW()
	// Capacity returns the maximum number of slots in the semaphore.
	//
	// The capacity is the maximum number of ops that can be performed
	// concurrently.
	Capacity() (writers int, readers int)
	// RemainingCapacity returns the remaining capacity of the semaphore.
	//
	// The remaining capacity is the maximum number of ops that can still be performed
	// concurrently.
	RemainingCapacity() (writers int, readers int)
}

type Barrier interface {
	// Wait causes the current goroutine to wait until the barrier is ready.
	Wait()
	// Count returns the number of goroutines waiting at the barrier.
	//
	// It returns -1 if the barrier is not ready.
	Count() int
	// Threshold returns the number of goroutines that must call Wait() before any of them can proceed.
	//
	// It returns -1 if the barrier is not ready.
	Threshold() int
}

// Streamable is any type which implements the Stream method.
type Streamable[T any] interface {
	// Stream returns a channel that can be used to receive values of type T.
	//
	// It's thread-safe and can be used concurrently.
	Stream() Stream[T]
}

// CloseableStreamable is a Streamable that is also Closeable.
// Once it is closed, it cannot be reopened or supply new values.
type CloseableStreamable[S any] interface {
	data.Closeable
	Streamable[S]
}

type Supplier[T any] interface {
	Streamable[T]
	// Supply takes a value of type T and sends it to the channel.
	//
	// It's thread-safe and can be used concurrently.
	//
	// If the channel is closed, the method returns an error.
	Supply(...T) error
	// Get returns all the values that have been sent to the supplier.
	//
	// It's thread-safe and can be used concurrently.
	Get() []T
	// Clear removes all the values that have been sent to the supplier.
	//
	// It's thread-safe and can be used concurrently.
	Clear()
}

type CloseableSupplier[T any] interface {
	data.Closeable
	Supplier[T]
}

type Consumer[T any] interface {
	// Consume consumes values from the channel using the provided function.
	//
	// The provided function takes two parameters: the index of the value and the value itself.
	//
	// It's thread-safe and can be used concurrently.
	Consume(functions.BiConsumer[int, T]) error
}

type CloseableConsumer[T any] interface {
	data.Closeable
	Consumer[T]
}

type SemaphoredChannel[T any] interface {
	data.Closeable
	// AddRoutine decrease the number of concurrent slots.
	AddRoutine()
	// IsEmpty returns true if the channel is empty, false otherwise.
	IsEmpty() bool
	// Len returns the number of elements in the channel.
	Len() int
	// Listen returns a channel that can be used to receive values of type T.
	Listen() <-chan T
	// RemoveRoutine increase the number of concurrent slots.
	RemoveRoutine()
	// Send sends a value of type T to the channel.
	Send(value T)
	// ToArray consumes and returns all the values that have been sent to the channel.
	ToArray() []T
}

type SynchronizableChannel[T any] interface {
	SemaphoredChannel[T]
	// AddWaiters increases the number of waiting routines.
	AddWaiters(i int)
	// Release releases a slot in the channel and decrease the number of waiting routines.
	Release()
	// RemoveWaiter decreases the number of waiting routines.
	RemoveWaiter()
	// SetWaitingPoint sets a point to block the goroutines until there are no more workloads.
	SetWaitingPoint()
	// SetEndingPoint sets a point to block the goroutines until there are no more workloads.
	// Then it closes the channel
	SetEndingPoint() error
}

type ChanSignal[T any] interface {
	// Close closes the channel.
	Close()
	// Send sends a value of type T to the channel.
	Send(value T)
	// ToArray consumes and returns all the values that have been sent to the channel.
	ToArray() []T
}
