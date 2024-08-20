// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tuple

import "github.com/andrerrcosta2/gtools/pkg/generics"

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

func DerefPair[A any, B any](b generics.BiTypedInterface[A, B]) (A, B) {
	var aZero A
	var bZero B
	if p, ok := b.(*Pair[A, B]); ok {
		return p.first, p.second
	}
	return aZero, bZero
}

func NewTriple[A any, B any, C any](a A, b B, c C) *Triple[A, B, C] {
	return &Triple[A, B, C]{
		first:  a,
		second: b,
		third:  c,
	}
}

var _ generics.BiTypedInterface[string, string] = (*Pair[string, string])(nil)

type Triple[A any, B any, C any] struct {
	first  A
	second B
	third  C
}

func (t *Triple[A, B, C]) First() A {
	return t.first
}

func (t *Triple[A, B, C]) Second() B {
	return t.second
}

func (t *Triple[A, B, C]) Third() C {
	return t.third
}

func DerefTriple[A any, B any, C any](t generics.TriTypedInterface[A, B, C]) (A, B, C) {
	var aZero A
	var bZero B
	var cZero C
	if p, ok := t.(*Triple[A, B, C]); ok {
		return p.first, p.second, p.third
	}
	return aZero, bZero, cZero
}

var _ generics.TriTypedInterface[string, string, string] = (*Triple[string, string, string])(nil)

type Quad[A any, B any, C any, D any] struct {
	first  A
	second B
	third  C
	fourth D
}

type Quint[A any, B any, C any, D any, E any] struct {
	first  A
	second B
	third  C
	fourth D
	fifth  E
}

type Sextet[A any, B any, C any, D any, E any, F any] struct {
	first  A
	second B
	third  C
	fourth D
	fifth  E
	sixth  F
}
