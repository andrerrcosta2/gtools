// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package channel

type Signal = <-chan struct{}

type Semaphored[T any] interface {
	AddRoutine()
	IsEmpty() bool
	Len() int
	Listen() <-chan T
	RemoveRoutine()
	Send(value T)
	ToArray() []T
	IsClosed() bool
	Close() error
}

type Synchronizable[T any] interface {
	Semaphored[T]
	AddWaiters(i int)
	Release()
	RemoveWaiter()
	SetWaitingPoint()
	SetEndingPoint() error
}
