// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparators

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type Typed[O any] interface {
	Compare(a, b O) int
	Equals(a, b O) bool
}

type KeyTyped[K any, O prim.Ordered] interface {
	Typed[K]
	Hash(key K) O
}

// Functional represents equals using 0, less than using -1 and greater than using 1
type Functional[A any] func(a, b A) int

func (f Functional[A]) Equals(a, b A) bool {
	return f(a, b) == 0
}

func (f Functional[A]) Compare(a, b A) int {
	return f(a, b)
}
