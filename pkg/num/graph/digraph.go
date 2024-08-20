// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/datastr/maps"
	"github.com/andrerrcosta2/gtools/pkg/datastr/sets"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
)

// DigraphOf creates a new DirectedGraphOf instance.
//
// It initializes a DirectedGraphOf with an empty adjacency list.
//
// Returns:
//
//	*DirectedGraphOf[G]: A new DirectedGraphOf instance with an empty adjacency list.
func DigraphOf[G gtools.SortableOf]() *DirectedGraphOf[G] {
	// Create a new DirectedGraphOf instance with an empty adjacency list.
	return &DirectedGraphOf[G]{
		// Initialize the adjacency list as a new, empty map.
		adj: maps.SortableOf[G, []G](),
	}
}

var _ Graph[gtools.SortableOf] = (*DirectedGraphOf[gtools.SortableOf])(nil)
var _ SingleEdgedGraph[gtools.SortableOf] = (*DirectedGraphOf[gtools.SortableOf])(nil)
var _ GraphOf[gtools.SortableOf] = (*DirectedGraphOf[gtools.SortableOf])(nil)

// DirectedGraphOf is a basic implementation of a directed graph using an adjacency list.
type DirectedGraphOf[G gtools.SortableOf] struct {
	adj *maps.SortableOfMap[G, []G]
}

// AddNode adds a node to the graph.
func (g *DirectedGraphOf[G]) AddNode(node G) {
	addNodeOfIfNotExists(g.adj, node)
}

// AddEdge adds a directed edge from 'from' to 'to'.
func (g *DirectedGraphOf[G]) AddEdge(from, to G) {
	addDirectEdgeOfIfNodesExist(g.adj, from, to)
}

// Neighbors returns the outgoing neighbors of a node.
func (g *DirectedGraphOf[G]) Neighbors(node G) []G {
	// Attempt to retrieve the neighbors of the node from the adjacency list
	if neighbors, ok := g.adj.Get(node); ok {
		// If the node exists, return its neighbors
		return neighbors
	}
	// If the node doesn't exist, return an empty slice
	return nil
}

// HasNode checks if a node exists in the graph.
func (g *DirectedGraphOf[G]) HasNode(node G) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if a directed edge exists from 'from' to 'to'.
func (g *DirectedGraphOf[G]) HasEdge(from, to G) bool {
	// Get the neighbors of 'from' node
	if neighbors, ok := g.adj.Get(from); ok {
		// Iterate over the neighbors
		for _, neighbor := range neighbors {
			// Check if 'neighbor' is equal to 'to'
			if neighbor.Equal(to) {
				// If 'to' is a neighbor of 'from', return true
				return true
			}
		}
	}
	// If 'from' does not exist or 'to' is not a neighbor of 'from', return false
	return false
}

// Nodes return all nodes in the graph.
func (g *DirectedGraphOf[G]) Nodes() []G {
	// Get the keys from the adjacency list, which represent the nodes in the graph.
	return g.adj.Keys()
}

// Edges returns all directed edges in the graph.
func (g *DirectedGraphOf[G]) Edges() []*SingleTypedEdge[G] {
	// Create a set to store unique edges
	edges := sets.SortableOf[*SingleTypedEdge[G]]()

	// Get an iterator over the adjacency list
	iterator := g.adj.Iterator()

	// Iterate over all nodes and their neighbors
	for key, value, ok := iterator.Next(); ok; key, value, ok = iterator.Next() {
		// Iterate over all neighbors of the current node
		for _, neighbor := range value {
			edges.Add(NewEdge(key, neighbor))
		}
	}
	// Return the slice of unique edges
	return edges.Values()
}

// String returns a string representation of the graph.
func (g *DirectedGraphOf[G]) String() string {
	return g.adj.String()
}
