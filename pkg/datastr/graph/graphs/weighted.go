// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/datastr/maps"
	"github.com/andrerrcosta2/gtools/pkg/datastr/sets"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/nums/graph"
)

// WeightedOrderedOf returns a new instance of WeightedOrderedGraphOf.
// This function initializes the adjacency list as a map of maps.
func WeightedOrderedOf[G gtools.SortableOf, W constraints.Ordered](nods ...G) *WeightedOrderedGraphOf[G, W] {
	// Initialize the adjacency list with a map of maps.
	graph := &WeightedOrderedGraphOf[G, W]{
		adj: maps.SortableOf[G, *maps.SortableOfMap[G, W]](),
	}

	// Add nodes to the graph
	for _, nod := range nods {
		graph.adj.Put(nod, maps.SortableOf[G, W]())
	}

	// Return the graph
	return graph
}

var _ graph.Graph[gtools.SortableOf] = (*WeightedOrderedGraphOf[gtools.SortableOf, int])(nil)
var _ graph.SingleWeightedEdgesGraph[gtools.SortableOf, int] = (*WeightedOrderedGraphOf[gtools.SortableOf, int])(nil)
var _ graph.WOrderedGraphOf[gtools.SortableOf, int] = (*WeightedOrderedGraphOf[gtools.SortableOf, int])(nil)

type WeightedOrderedGraphOf[G gtools.SortableOf, W constraints.Ordered] struct {
	adj *maps.SortableOfMap[G, *maps.SortableOfMap[G, W]]
}

// AddNode adds a node to the graph.
func (g *WeightedOrderedGraphOf[G, W]) AddNode(node G) {
	if !g.adj.Contains(node) {
		g.adj.Put(node, maps.SortableOf[G, W]())
	}
}

// AddEdge adds a directed, weighted edge from 'from' to 'to' with a given weight.
func (g *WeightedOrderedGraphOf[G, W]) AddEdge(from, to G, weight W) {
	graph.addWeightedEdgeOfIfNodesExist(g.adj, from, to, weight)
}

// Neighbors returns the neighbors of a node and their associated weights.
func (g *WeightedOrderedGraphOf[G, W]) Neighbors(node G) []G {
	// Attempt to retrieve the neighbors of the node from the adjacency list
	if neighbors, ok := g.adj.Get(node); ok {
		// If the node exists, return its neighbors
		return neighbors.Keys()
	}
	// If the node does not exist, return an empty slice
	return nil
}

// HasNode checks if a node exists in the graph.
func (g *WeightedOrderedGraphOf[G, W]) HasNode(node G) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if a weighted edge exists from 'from' to 'to'.
func (g *WeightedOrderedGraphOf[G, W]) HasEdge(from, to G) bool {
	// Get the neighbors of 'from' node
	if neighbors, ok := g.adj.Get(from); ok {
		it := neighbors.Iterator()
		for key, _, ok := it.Next(); ok; key, _, ok = it.Next() {
			// Check if 'key' is equal to 'to'
			if key.Equal(to) {
				// If 'to' is a neighbor of 'from', return true
				return true
			}
		}
	}
	// If 'from' does not exist or 'to' is not a neighbor of 'from', return false
	return false
}

// Weight returns the weight of the edge from 'from' to 'to'.
func (g *WeightedOrderedGraphOf[G, W]) Weight(from, to G) (W, bool) {
	if neighbors, ok := g.adj.Get(from); ok {
		if weight, ok := neighbors.Get(to); ok {
			return weight, true
		}
	}
	var zero W
	return zero, false
}

// Nodes returns all nodes in the graph.
func (g *WeightedOrderedGraphOf[G, W]) Nodes() []G {
	return g.adj.Keys()
}

func (g *WeightedOrderedGraphOf[G, W]) Edges() []*graph.SingleTypedWeightedEdge[G, W] {
	edges := sets.SortableOf[*graph.SingleTypedWeightedEdge[G, W]]()
	fit := g.adj.Iterator()

	for from, tos, ok := fit.Next(); ok; from, tos, ok = fit.Next() {
		tit := tos.Iterator()
		for to, weight, ok := tit.Next(); ok; to, weight, ok = tit.Next() {
			edges.Add(graph.NewWeightedEdge(from, to, weight))
		}
	}
	return edges.Values()
}

func (g *WeightedOrderedGraphOf[G, W]) String() string {
	return g.adj.String()
}
