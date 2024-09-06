// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testdata

func NewPair[A any, B any](a A, b B) *Pair[A, B] {
	return &Pair[A, B]{
		first:  a,
		second: b,
	}
}

type Pair[A any, B any] struct {
	first  A
	second B
}

func (p *Pair[A, B]) First() A {
	return p.first
}

func (p *Pair[A, B]) Second() B {
	return p.second
}
