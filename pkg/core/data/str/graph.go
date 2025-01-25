// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

import (
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

// Graph represents the basic interface for a graph
type Graph[N any] interface {
	IterableGraph[N]
	AddNode(id N)
	HasNode(id N) bool
	Nodes() []N
}

// IterableGraph represents the basic interface for an iterable graph
type IterableGraph[N any] interface {
	Neighbors(id N) []N
}

type EdgedGraph[N any, E edges.Edge[F, T], F any, T any] interface {
	Graph[N]
	HasEdge(from F, to T) bool
	Edges() []E
}

// SingleEdgedGraph represents the basic interface for a graph with a single edge
type SingleEdgedGraph[N any, E edges.SingleTyped[N]] interface {
	EdgedGraph[N, E, N, N]
	AddEdge(from, to N) error
}

// SingleWeightedEdgesGraph represents the basic interface for a graph with a single weighted edge
type SingleWeightedEdgesGraph[N any, W any, E edges.SingleTypedWeighted[N, W]] interface {
	EdgedGraph[N, E, N, N]
	AddEdge(from, to N, weight W) error
}

// SingleWeightedGraph represents the basic interface for a weighted graph
type SingleWeightedGraph[N any, W any] interface {
	Graph[N]
	Weight(from, to N) (W, bool)
}

// OrderedGraph represents the basic interface for a graph of a constraints.Ordered type
type OrderedGraph[N prim.Ordered, E edges.SingleTyped[N]] interface {
	Graph[N]
	SingleEdgedGraph[N, E]
}

// GraphOf represents the basic interface for a graph of a core.SortableOf type
type GraphOf[N gtools.SortableOf, E edges.SortableSingleTyped[N]] interface {
	SingleEdgedGraph[N, E]
}

type WGraph[N any, W any, E edges.SingleTypedWeighted[N, W]] interface {
	Graph[N]
	SingleWeightedEdgesGraph[N, W, E]
	SingleWeightedGraph[N, W]
}

type WOrderedGraph[N prim.Ordered, W prim.Ordered, E edges.SingleTypedWeighted[N, W]] interface {
	Graph[N]
	SingleWeightedEdgesGraph[N, W, E]
	SingleWeightedGraph[N, W]
}

type WOrderedGraphOf[N gtools.SortableOf, W prim.Ordered, E edges.SortableSingleTypedWeighted[N, W]] interface {
	Graph[N]
	SingleWeightedEdgesGraph[N, W, E]
	SingleWeightedGraph[N, W]
}

type WComparableGraphOf[N gtools.SortableOf, W gtools.ComparableOf, E edges.SingleTypedWeighted[N, W]] interface {
	Graph[N]
	SingleWeightedEdgesGraph[N, W, E]
	SingleWeightedGraph[N, W]
}

type WAggregableGraphOf[N gtools.SortableOf, W gtools.AggregableOf, E edges.SingleTypedWeighted[N, W]] interface {
	Graph[N]
	SingleWeightedEdgesGraph[N, W, E]
	SingleWeightedGraph[N, W]
}
