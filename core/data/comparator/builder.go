// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparator

import "github.com/andrerrcosta2/gtools/core/domain/function"

type Builder[T any, H comparable] interface {
	Equals(equals function.Equals[T]) Builder[T, H]
	Compare(compare function.Compare[T]) Builder[T, H]
	Key(key function.Function[T, H]) Builder[T, H]
	Build() EqCmpKeyer[T, H]
}

func Build[T any, H comparable]() Builder[T, H] {
	return builder[T, H]{}
}

type builder[T any, H comparable] struct {
	comparator composite[T, H]
}

func (b builder[T, H]) Equals(equals function.Equals[T]) Builder[T, H] {
	b.comparator.equals = equals
	return b
}

func (b builder[T, H]) Compare(compare function.Compare[T]) Builder[T, H] {
	b.comparator.compare = compare
	return b
}

func (b builder[T, H]) Key(key function.Function[T, H]) Builder[T, H] {
	b.comparator.key = key
	return b
}

func (b builder[T, H]) Build() EqCmpKeyer[T, H] {
	return b.comparator
}

type composite[T any, H comparable] struct {
	equals  function.Equals[T]
	compare function.Compare[T]
	key     function.Function[T, H]
}

func (c composite[T, H]) Equals(a, b T) bool {
	return c.equals(a, b)
}

func (c composite[T, H]) Compare(a, b T) int {
	return c.compare(a, b)
}

func (c composite[T, H]) Key(t T) H {
	return c.key(t)
}
