// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
)

// Graph represents the basic interface for a graph
type Graph[G any] interface {
	Iterable[G]
	AddNode(id G)
	HasNode(id G) bool
	HasEdge(from, to G) bool
	Nodes() []G
}

// Iterable represents the basic interface for an iterable graph
type Iterable[G any] interface {
	Neighbors(id G) []G
}

// SingleEdgedGraph represents the basic interface for a graph with a single edge
type SingleEdgedGraph[G any] interface {
	AddEdge(from, to G)
	Edges() []*SingleTypedEdge[G]
}

// SingleWeightedEdgesGraph represents the basic interface for a graph with a single weighted edge
type SingleWeightedEdgesGraph[G any, W any] interface {
	AddEdge(from, to G, weight W)
	Edges() []*SingleTypedWeightedEdge[G, W]
}

// SingleWeightedGraph represents the basic interface for a weighted graph
type SingleWeightedGraph[G any, W any] interface {
	Weight(from, to G) (W, bool)
}

// OrderedGraph represents the basic interface for a graph of a constraints.Ordered type
type OrderedGraph[G constraints.Ordered] interface {
	Graph[G]
	SingleEdgedGraph[G]
}

// GraphOf represents the basic interface for a graph of a gtools.SortableOf type
type GraphOf[G gtools.SortableOf] interface {
	Graph[G]
	SingleEdgedGraph[G]
}

type WGraph[G any, W any] interface {
	Graph[G]
	SingleWeightedEdgesGraph[G, W]
	SingleWeightedGraph[G, W]
}

type WOrderedGraph[G constraints.Ordered, W constraints.Ordered] interface {
	Graph[G]
	SingleWeightedEdgesGraph[G, W]
	SingleWeightedGraph[G, W]
}

type WOrderedGraphOf[G gtools.SortableOf, W constraints.Ordered] interface {
	Graph[G]
	SingleWeightedEdgesGraph[G, W]
	SingleWeightedGraph[G, W]
}

type WComparableGraphOf[G gtools.SortableOf, W gtools.ComparableOf] interface {
	Graph[G]
	SingleWeightedEdgesGraph[G, W]
	SingleWeightedGraph[G, W]
}

type WAggregableGraphOf[G gtools.SortableOf, W gtools.AggregableOf] interface {
	Graph[G]
	SingleWeightedEdgesGraph[G, W]
	SingleWeightedGraph[G, W]
}
