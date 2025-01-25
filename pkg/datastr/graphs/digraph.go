// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/datastr/maps"
	"github.com/andrerrcosta2/gtools/datastr/sets"
)

// DigraphOf creates a new DirectedGraphOf instance.
//
// It initializes a DirectedGraphOf with the given nodes.
//
// Returns:
//
//	*DirectedGraphOf[N]: A new DirectedGraphOf instance with the given nodes.
func DigraphOf[G gtools.SortableOf](nods ...G) str.GraphOf[G, edges.SortableSingleTyped[G]] {
	// Create a new DirectedGraphOf instance with an empty adjacency list.
	graph := &directedGraphOf[G]{
		// Initialize the adjacency list as a new, empty map.
		adj: maps.SortableOf[G, []G](),
	}

	// Addf nodes to the graph
	for _, nod := range nods {
		graph.adj.Put(nod, []G{})
	}

	// Return the graph
	return graph
}

var _ str.Graph[gtools.SortableOf] = (*directedGraphOf[gtools.SortableOf])(nil)
var _ str.SingleEdgedGraph[gtools.SortableOf, edges.SortableSingleTyped[gtools.SortableOf]] = (*directedGraphOf[gtools.SortableOf])(nil)
var _ str.GraphOf[gtools.SortableOf, edges.SortableSingleTyped[gtools.SortableOf]] = (*directedGraphOf[gtools.SortableOf])(nil)

// DirectedGraphOf is a basic implementation of a directed graph using an adjacency list.
type directedGraphOf[G gtools.SortableOf] struct {
	adj *maps.SortableOfMap[G, []G]
}

// AddNode adds a node to the graph.
func (g *directedGraphOf[G]) AddNode(node G) {
	AddNodeOfIfNotExists[G](g.adj, node)
}

// AddEdge adds a directed edge from 'from' to 'to'.
func (g *directedGraphOf[G]) AddEdge(from, to G) error {
	return AddDirectEdgeOfIfNodesExist(g.adj, from, to)
}

// Neighbors returns the outgoing neighbors of a node.
func (g *directedGraphOf[G]) Neighbors(node G) []G {
	// Attempt to retrieve the neighbors of the node from the adjacency list
	if neighbors, ok := g.adj.Get(node); ok {
		// If the node exists, return its neighbors
		return neighbors
	}
	// If the node doesn't exist, return an empty slice
	return nil
}

// HasNode checks if a node exists in the graph.
func (g *directedGraphOf[G]) HasNode(node G) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if a directed edge exists from 'from' to 'to'.
func (g *directedGraphOf[G]) HasEdge(from, to G) bool {
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
func (g *directedGraphOf[G]) Nodes() []G {
	// Get the keys from the adjacency list, which represent the nodes in the graph.
	return g.adj.Keys()
}

// Edges returns all directed edges in the graph.
func (g *directedGraphOf[G]) Edges() []edges.SortableSingleTyped[G] {
	// Create a set to store unique edges
	edges := sets.SortableOf[edges.SortableSingleTyped[G]]()

	// Get an iterator over the adjacency list
	iterator := g.adj.Iterator()

	// Iterate over all nodes and their neighbors
	for key, value, ok := iterator.Next(); ok; key, value, ok = iterator.Next() {
		// Iterate over all neighbors of the current node
		for _, neighbor := range value {
			edges.Add(SortableEdge[G](key, neighbor, true))
		}
	}
	// Return the slice of unique edges
	return edges.Values()
}

// String returns a string representation of the graph.
func (g *directedGraphOf[G]) String() string {
	return g.adj.String()
}

// Digraph creates a new DirectedGraph instance.
//
// It initializes a DirectedGraph with the given nodes.
//
// Returns:
//
//	*DirectedGraph[N]: A new DirectedGraph instance with the given nodes.
func Digraph[G prim.Ordered](nods ...G) str.OrderedGraph[G, edges.SortableSingleTyped[G]] {
	// Create a new DirectedGraph instance with the given nodes
	graph := &directedGraph[G]{
		// Initialize the adjacency list as a new, empty map.
		adj: make(map[G][]G, len(nods)),
	}

	// Addf nodes to the graph
	for _, nod := range nods {
		graph.adj[nod] = []G{}
	}

	// Return the graph
	return graph
}

var _ str.Graph[int] = (*directedGraph[int])(nil)
var _ str.SingleEdgedGraph[int, edges.SortableSingleTyped[int]] = (*directedGraph[int])(nil)
var _ str.OrderedGraph[int, edges.SortableSingleTyped[int]] = (*directedGraph[int])(nil)

// DirectedGraph is a basic implementation of a directed graph using an adjacency list.
type directedGraph[G prim.Ordered] struct {
	adj map[G][]G
}

// AddNode adds a node to the graph.
func (g *directedGraph[G]) AddNode(node G) {
	AddNodeIfNotExists[G](&g.adj, node)
}

// AddEdge adds a directed edge from 'from' to 'to'.
func (g *directedGraph[G]) AddEdge(from, to G) error {
	return AddDirectEdgeIfNodesExist(&g.adj, from, to)
}

// Neighbors returns the outgoing neighbors of a node.
func (g *directedGraph[G]) Neighbors(node G) []G {
	// Attempt to retrieve the neighbors of the node from the adjacency list
	if neighbors, ok := g.adj[node]; ok {
		// If the node exists, return its neighbors
		return neighbors
	}
	// If the node doesn't exist, return an empty slice
	return nil
}

// HasNode checks if a node exists in the graph.
func (g *directedGraph[G]) HasNode(node G) bool {
	// Check if the node exists in the adjacency list
	_, ok := g.adj[node]
	return ok
}

// HasEdge checks if a directed edge exists from 'from' to 'to'.
func (g *directedGraph[G]) HasEdge(from, to G) bool {
	// Get the neighbors of 'from' node
	if neighbors, ok := g.adj[from]; ok {
		// Iterate over the neighbors
		for _, neighbor := range neighbors {
			// Check if 'neighbor' is equal to 'to'
			if neighbor == to {
				// If 'to' is a neighbor of 'from', return true
				return true
			}
		}
	}
	// If 'from' does not exist or 'to' is not a neighbor of 'from', return false
	return false
}

// Nodes return all nodes in the graph.
func (g *directedGraph[G]) Nodes() []G {
	// Get the keys from the adjacency list, which represent the nodes in the graph.
	out := make([]G, 0, len(g.adj))
	for key := range g.adj {
		out = append(out, key)
	}
	// Return the slice of nodes
	return out
}

// Edges returns all directed edges in the graph.
func (g *directedGraph[G]) Edges() []edges.SortableSingleTyped[G] {
	// Create a set to store unique edges
	edges := sets.SortableOf[edges.SortableSingleTyped[G]]()

	// Iterate over all nodes and their neighbors
	for key, value := range g.adj {
		// Iterate over all neighbors of the current node
		for _, neighbor := range value {
			edges.Add(SortableEdge[G](key, neighbor, true))
		}
	}
	// Return the slice of unique edges
	return edges.Values()
}

// String returns a string representation of the graph.
func (g *directedGraph[G]) String() string {
	return fmt.Sprintf("%v", g.adj)
}
