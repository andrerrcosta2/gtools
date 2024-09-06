// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package generics

import (
	"github.com/andrerrcosta2/gtools/core/constraints/nums"
	"github.com/andrerrcosta2/gtools/core/constraints/prim"
)

// TypedInterface primary single generic-typed interface
type TypedInterface[A any] interface{}

// BiTypedInterface primary doubled generic-typed interface
type BiTypedInterface[A any, B any] interface{}

// TriTypedInterface primary tripled generic-typed interface
type TriTypedInterface[A any, B any, C any] interface{}

// TypedOrdered primary single generic-ordered interface
type TypedOrdered[A prim.Ordered] interface{}

// BiTypedOrdered primary doubled generic-ordered interface
type BiTypedOrdered[A prim.Ordered, B prim.Ordered] interface{}

// TriTypedOrdered primary tripled generic-ordered interface
type TriTypedOrdered[A prim.Ordered, B prim.Ordered, C prim.Ordered] interface{}

// TypedNumeric primary single generic-numeric interface
type TypedNumeric[N nums.Any] interface{}

// BiTypedNumeric primary doubled generic-numeric interface
type BiTypedNumeric[N nums.Any, M nums.Any] interface{}
