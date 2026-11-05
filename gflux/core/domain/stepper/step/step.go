// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package step

type Step[T any] interface {
	At() T
	Exec() error
}

// Keyed decorates a step with a key
type Keyed[K comparable] interface {
	Key() K
}

// Stoppable is a step that can be marked as stopped and later queried.
type Stoppable[T any] interface {
	Step[T]
	Stop()
	Stopped() bool
}

// Priority decorates a step with an integer priority.
type Priority[T any] interface {
	Step[T]
	Priority() int
}

// Conflatable is a step that can be identified by a key.
type Conflatable[T any, K comparable] interface {
	Step[T]
	Keyed[K]
}

// Forwarder is a step that can yield a successor step.
type Forwarder[T any] interface {
	Step[T]
	Next() (Forwarder[T], bool)
}

type Releasable[T any] interface {
	Step[T]
	Release()
	Released() bool
}

type Future[T any] interface {
	Step[T]
	Ready() bool
	Wait()
	Try() (error, error, bool)
}

// Queue is a push/peek/pop source of ordered broadcasters.
type Queue[S Step[T], T any] interface {
	Push(S)
	Peek() (S, bool)
	Next() (S, bool)
}
