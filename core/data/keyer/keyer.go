// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package keyer

type Comparable[T any, K comparable] interface {
	Key(T) K
}

type Func[T any, K comparable] func(T) K

func (f Func[T, K]) Key(t T) K {
	return f(t)
}
