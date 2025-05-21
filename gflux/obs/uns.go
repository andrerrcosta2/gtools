package obs

type Unsubscribable[T any] interface {
	Unsub()
	Closed() bool
}
