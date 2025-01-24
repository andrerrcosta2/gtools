// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/core/gtools/constraints/prim"
	"github.com/andrerrcosta2/gtools/datastr/maps"
	"github.com/andrerrcosta2/gtools/datastr/sets"
)

// WeightedSortableDigraphOf returns a new instance of str.WOrderedGraphOf of gtools.SortableOf node type.
// This function initializes the adjacency list as a map of maps.
func WeightedSortableDigraphOf[G gtools.SortableOf, W prim.Ordered](nods ...G) str.WOrderedGraphOf[G, W, edges.SortableSingleTypedWeighted[G, W]] {
	// Initialize the adjacency list with a map of maps.
	graph := &weightedSortableDigraphOf[G, W]{
		adj: maps.SortableOf[G, *maps.SortableOfMap[G, W]](),
	}

	// Addf nodes to the graph
	for _, nod := range nods {
		graph.adj.Put(nod, maps.SortableOf[G, W]())
	}

	// Return the graph
	return graph
}

var _ str.Graph[gtools.SortableOf] = (*weightedSortableDigraphOf[gtools.SortableOf, int])(nil)
var _ str.SingleWeightedEdgesGraph[gtools.SortableOf, int, edges.SortableSingleTypedWeighted[gtools.SortableOf, int]] = (*weightedSortableDigraphOf[gtools.SortableOf, int])(nil)
var _ str.WOrderedGraphOf[gtools.SortableOf, int, edges.SortableSingleTypedWeighted[gtools.SortableOf, int]] = (*weightedSortableDigraphOf[gtools.SortableOf, int])(nil)

type weightedSortableDigraphOf[G gtools.SortableOf, W prim.Ordered] struct {
	adj *maps.SortableOfMap[G, *maps.SortableOfMap[G, W]]
}

// AddNode adds a node to the graph.
func (g *weightedSortableDigraphOf[G, W]) AddNode(node G) {
	if !g.adj.Contains(node) {
		g.adj.Put(node, maps.SortableOf[G, W]())
	}
}

// AddEdge adds a directed, weighted edge from 'from' to 'to' with a given weight.
func (g *weightedSortableDigraphOf[G, W]) AddEdge(from, to G, weight W) error {
	return AddWeightedEdgeOfIfNodesExist(g.adj, from, to, weight)
}

// Neighbors returns the neighbors of a node and their associated weights.
func (g *weightedSortableDigraphOf[G, W]) Neighbors(node G) []G {
	// Attempt to retrieve the neighbors of the node from the adjacency list
	if neighbors, ok := g.adj.Get(node); ok {
		// If the node exists, return its neighbors
		return neighbors.Keys()
	}
	// If the node does not exist, return an empty slice
	return nil
}

// HasNode checks if a node exists in the graph.
func (g *weightedSortableDigraphOf[G, W]) HasNode(node G) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if a weighted edge exists from 'from' to 'to'.
func (g *weightedSortableDigraphOf[G, W]) HasEdge(from, to G) bool {
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
func (g *weightedSortableDigraphOf[G, W]) Weight(from, to G) (weight W, exists bool) {
	if neighbors, ok := g.adj.Get(from); ok {
		weight, exists = neighbors.Get(to)
	}
	return
}

// Nodes returns all nodes in the graph.
func (g *weightedSortableDigraphOf[G, W]) Nodes() []G {
	return g.adj.Keys()
}

func (g *weightedSortableDigraphOf[G, W]) Edges() []edges.SortableSingleTypedWeighted[G, W] {
	e := sets.SortableOf[edges.SortableSingleTypedWeighted[G, W]]()
	fit := g.adj.Iterator()

	for from, tos, ok := fit.Next(); ok; from, tos, ok = fit.Next() {
		tit := tos.Iterator()
		for to, weight, ok := tit.Next(); ok; to, weight, ok = tit.Next() {
			e.Add(SingleTypedWeightedEdge(from, to, weight, true))
		}
	}
	return e.Values()
}

func (g *weightedSortableDigraphOf[G, W]) String() string {
	return g.adj.String()
}

var _ str.EdgedGraph[gtools.SortableOf, edges.SortableSingleTypedWeighted[gtools.SortableOf, int], gtools.SortableOf, gtools.SortableOf] = (*weightedSortableDigraphOf[gtools.SortableOf, int])(nil)
