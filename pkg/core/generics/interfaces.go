// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package generics

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
)

// Typed primary single generic-typed interface
type Typed[A any] interface{}

// BiTyped primary doubled generic-typed interface
type BiTyped[A any, B any] interface{}

// TriTyped primary tripled generic-typed interface
type TriTyped[A any, B any, C any] interface{}

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
