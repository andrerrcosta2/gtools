package obs

type Obs[T any] interface {
	Emit(value T)          // Emit a value
	OnError(err error)     // Emit an error
	OnComplete()           // Signal completion
	RemoveSub(sub *Sub[T]) // Remove a subscription
	Subscribable[T]        // Implement Subscribable
}
