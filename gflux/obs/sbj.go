package obs

type Subject[T any] interface {
	Obs[T]             // Subject is Observable
	Subscribable[T]    // Subject is Subscribable
	OnError(err error) // Emit an error
	OnComplete()       // Signal completion
	UnsubAll()         // Unsub all observers
}
