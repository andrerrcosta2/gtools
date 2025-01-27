// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package datastr

func PairOf[A any, B any](first A, second B) *Pair[A, B] {
	return &Pair[A, B]{
		first:  first,
		second: second,
	}
}

type Pair[A any, B any] struct {
	first  A
	second B
}

func (p Pair[A, B]) First() A {
	return p.first
}

func (p Pair[A, B]) Second() B {
	return p.second
}
