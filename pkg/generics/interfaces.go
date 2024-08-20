// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package generics

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
)

// TypedInterface primary single generic typed interface
type TypedInterface[A any] interface{}

// BiTypedInterface primary doubled generic typed interface
type BiTypedInterface[A any, B any] interface{}

// TriTypedInterface primary tripled generic typed interface
type TriTypedInterface[A any, B any, C any] interface{}

type TypedOrdered[A constraints.Ordered] interface{}

type BiTypedOrdered[A constraints.Ordered, B constraints.Ordered] interface{}

type TriTypedOrdered[A constraints.Ordered, B constraints.Ordered, C constraints.Ordered] interface{}
