// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparables

import "github.com/andrerrcosta2/gtools/pkg/constraints"

type Comparator[O any] interface {
	Compare(a, b O) int
	Equals(a, b O) bool
}

type KeyComparator[K any, O constraints.Ordered] interface {
	Hash(key K) O
	Comparator[K]
}

// FunctionalComparator represents equals using 0, less than using -1 and greater than using 1
type FunctionalComparator[A any] func(a, b A) int

func (f FunctionalComparator[A]) Equals(a, b A) bool {
	return f(a, b) == 0
}

func (f FunctionalComparator[A]) Compare(a, b A) int {
	return f(a, b)
}
