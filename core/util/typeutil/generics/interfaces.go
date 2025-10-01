// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package generics

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
)

// Typed primary single generic-typed interface
type Typed[A any] interface{}

// BiTyped primary doubled generic-typed interface
type BiTyped[A, B any] interface{}

// TriTyped primary tripled generic-typed interface
type TriTyped[A, B, C any] interface{}

type QuadTyped[A, B, C, D any] interface{}

// TypedOrdered primary single generic-ordered interface
type TypedOrdered[A prim.Ordered] interface{}

// BiTypedOrdered primary doubled generic-ordered interface
type BiTypedOrdered[A, B prim.Ordered] interface{}

// TriTypedOrdered primary tripled generic-ordered interface
type TriTypedOrdered[A, B, C prim.Ordered] interface{}

type QuadTypedOrdered[A, B, C, D prim.Ordered] interface{}

// TypedNumeric primary single generic-numeric interface
type TypedNumeric[N nums.Any] interface{}

// BiTypedNumeric primary doubled generic-numeric interface
type BiTypedNumeric[N, M nums.Any] interface{}

type TriTypedNumeric[N, M, L nums.Any] interface{}

type QuadTypedNumeric[N, M, L, K nums.Any] interface{}

type Tuple2[A, B any] interface {
	GetFirst() A
	GetSecond() B
	SetFirst(first A)
	SetSecond(second B)
}

type Tuple3[A, B, C any] interface {
	GetFirst() A
	GetSecond() B
	GetThird() C
	SetFirst(first A)
	SetSecond(second B)
	SetThird(third C)
}

type Tuple4[A, B, C, D any] interface {
	GetFirst() A
	GetSecond() B
	GetThird() C
	GetFourth() D
	SetFirst(first A)
	SetSecond(second B)
	SetThird(third C)
	SetFourth(fourth D)
}
