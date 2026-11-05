// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
)

type IterationAlgorithm interface {
}

// Readable represents the basic interface for a graph
type Readable[N any] interface {
	HasNode(n N) bool
	Nodes() []N
}

type Immutable[N any] interface {
	Readable[N]
}

type Writable[N any] interface {
	AddNode(N)
	RemoveNode(N) bool
}

type ReadableEdged[E edge.Typed[N], N any] interface {
	Readable[N]
	HasEdge(from, to N) bool
	Edges() []E
	EdgesFrom(from N) []E
}

type WritableEdged[E edge.Typed[N], N any] interface {
	Writable[N]
	AddEdge(N, N) error
	RemoveEdge(N, N) bool
}

type Mutable[N any] interface {
	Readable[N]
	Writable[N]
}

// Iterable represents the basic interface for an iterable graph
type Iterable[N any] interface {
	Range(func(N, set.Readable[N]) bool)
	RangeNeighbors(N, func(N) bool)
	Neighbors(N) []N
	Nodes() []N
}

type MutableEdged[E edge.Typed[N], N any] interface {
	ReadableEdged[E, N]
	WritableEdged[E, N]
	Iterable[N]
}

type WritableWeighted[N any, W any] interface {
	Writable[N]
	AddEdge(N, N, W) error
	RemoveEdge(N, N) bool
}

type ReadableWeighted[E edge.Weighted[N, W], N any, W any] interface {
	ReadableEdged[E, N]
}

type MutableWeighted[E edge.Weighted[N, W], N any, W any] interface {
	ReadableWeighted[E, N, W]
	WritableWeighted[N, W]
	Iterable[N]
}
