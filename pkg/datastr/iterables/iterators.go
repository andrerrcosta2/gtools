// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterables

type Loopable[T any] interface {
	Loop() <-chan T
}

type Iterable[T any] interface {
	Next() Iterable[T]
	HasNext() bool
	Get() T
}

type BiTypedIterable[T any, U any] interface {
	Next() BiTypedIterable[T, U]
	HasNext() bool
	Get() (T, U)
}

type MapIterator[T any, U any] interface {
	Next() (T, U, bool)
}

type Iterator[T any] interface {
	Next() (T, bool)
}
