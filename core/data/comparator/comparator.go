// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparator

import "github.com/andrerrcosta2/gtools/core/data/keyer"

type Equals[T any] interface {
	Equals(a, b T) bool
}

type Comparer[T any] interface {
	Compare(a, b T) int
}

type EqComparer[T any] interface {
	Equals[T]
	Comparer[T]
}

type EqKeyer[H comparable, T any] interface {
	Equals[T]
	keyer.Comparable[T, H]
}

type CmpKeyer[T any, H comparable] interface {
	Comparer[T]
	keyer.Comparable[T, H]
}

type EqCmpKeyer[T any, H comparable] interface {
	EqComparer[T]
	keyer.Comparable[T, H]
}

// Func represents compare using 0, less than using -1 and greater than using 1
type Func[A any] func(a, b A) int

func (f Func[A]) Equals(a, b A) bool {
	return f(a, b) == 0
}

func (f Func[A]) Compare(a, b A) int {
	return f(a, b)
}
