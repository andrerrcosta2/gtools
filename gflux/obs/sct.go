package obs

type Subscribable[T any] interface {
	Sub(observer *Observer[T]) (*Sub[T], error)
}
