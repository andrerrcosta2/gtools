// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterables

type Iterable[T any] interface {
	Loop() <-chan T
}

type Iterator[T any] interface {
	Next() (T, bool)
}

type MapIterator[T any, U any] interface {
	Next() (T, U, bool)
}
