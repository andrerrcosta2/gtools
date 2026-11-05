// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

import "github.com/andrerrcosta2/gtools/core/domain/constraints/prim"

type Clearable interface {
	Clear()
}

type Ordered[T prim.Ordered] interface {
	Get(i int) (T, bool)
	Exclude(i int) bool
}

type Pushable[T any] interface {
	Push(T)
}

type Peekable[T any] interface {
	Peek() (T, bool)
}

type Sized interface {
	Empty() bool
	Len() int
}
