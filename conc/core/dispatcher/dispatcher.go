package dispatcher

import "errors"

var (
	ErrClosed = errors.New("dispatcher is closed")
)

type Handler[T any] func(T) error
type Listener[T any] func(T, error)

// Executor is the minimal contract expected by orchestration layers.
type Executor[T any] interface {
	Dispatch(item T) error
}

type Listenable[T any] interface {
	Executor[T]
	Listen(listener Listener[T]) // Sets a listener to react to dispatch ops
}

type Closeable[T any] interface {
	Executor[T]
	Close()
	IsClosed() bool
}

type Buffered[T any] interface {
	Executor[T]
	QueueSize() int
}

type Startable[T any] interface {
	Executor[T]
	Start() // Starts the dispatcher
}

type Strategy[T any] interface {
	Feedback(meta *Record[T])             // Called after each dispatch attempt with the result.
	BeforeDispatch(meta *Record[T]) error // Called before dispatching. Can block or influence dispatch.
	OnStart()
	OnStop()
}
