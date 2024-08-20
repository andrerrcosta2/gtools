// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/datastr/maps"
	"github.com/andrerrcosta2/gtools/pkg/datastr/sets"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/sortables"
)

// UndirectOf creates a new UndirectGraphOf instance.
// It returns a pointer to a UndirectGraphOf struct with an initialized adjacency list.
//
// This function is used to initialize a new undirected graph with an empty adjacency list.
func UndirectOf[G gtools.SortableOf]() *UndirectGraphOf[G] {
	// Create a new UndirectGraphOf instance with an empty adjacency list.
	// The adjacency list is represented as a map of type G to a slice of type G.
	// The map is initialized using the maps.SortableOf function from the gtools package.
	// The comparator is initialized using the sortables.ComparatorOf function from the gtools package.
	return &UndirectGraphOf[G]{
		// Initialize the adjacency list (adj) as an empty map.
		adj: maps.SortableOf[G, []G](),

		// Initialize the comparator for the graph.
		comparator: sortables.ComparatorOf[G](),
	}
}

var _ Graph[gtools.SortableOf] = (*UndirectGraphOf[gtools.SortableOf])(nil)
var _ SingleEdgedGraph[gtools.SortableOf] = (*UndirectGraphOf[gtools.SortableOf])(nil)
var _ GraphOf[gtools.SortableOf] = (*UndirectGraphOf[gtools.SortableOf])(nil)

// UndirectGraphOf is a basic implementation of an undirected graph using an adjacency list.
type UndirectGraphOf[G gtools.SortableOf] struct {
	adj        *maps.SortableOfMap[G, []G]
	comparator comparables.KeyComparator[G, string]
}

// AddNode adds a node to the graph.
// It checks if the node already exists in the graph before adding it.
// If the node doesn't exist, it adds the node to the adjacency list with an empty slice of neighbors.
func (g *UndirectGraphOf[G]) AddNode(node G) {
	addNodeOfIfNotExists(g.adj, node)
}

// AddEdge adds an undirected edge between two nodes.
//
// The function takes two parameters: 'from' and 'to', which represent the nodes
// between which the edge is to be added.
//
// The function first checks if both 'from' and 'to' nodes exist in the graph.
// If either of the nodes is not present, the function returns without performing
// any further operations.
//
// If both nodes exist, the function retrieves the neighbors of 'from' and 'to'
// from the adjacency list.
//
// The function then appends 'to' to the list of neighbors of 'from' and 'from'
// to the list of neighbors of 'to'.
//
// Finally, the function updates the adjacency list with the new neighbor lists.
func (g *UndirectGraphOf[G]) AddEdge(from, to G) {
	addUndirectEdgeOfIfNodesExist(g.adj, from, to)
}

// Neighbors returns the neighbors of a node in the undirected graph.
// It returns a slice of nodes that are directly connected to the given node.
func (g *UndirectGraphOf[G]) Neighbors(node G) []G {
	// Attempt to retrieve the neighbors of the node from the adjacency list
	if neighbors, ok := g.adj.Get(node); ok {
		// If the node exists, return its neighbors
		return neighbors
	}
	// If the node does not exist, return an empty slice
	return nil
}

// HasNode checks if a node exists in the graph.
// It returns true if the node is found, false otherwise.
func (g *UndirectGraphOf[G]) HasNode(node G) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if an edge exists between two nodes.
// It returns true if there is an edge from 'from' to 'to', false otherwise.
func (g *UndirectGraphOf[G]) HasEdge(from, to G) bool {
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

// Nodes returns a list of all nodes in the graph.
func (g *UndirectGraphOf[G]) Nodes() []G {
	// Get the keys from the adjacency list, which represent the nodes in the graph.
	return g.adj.Keys()
}

// Edges returns all edges in the graph.
// Edges returns all edges in the graph.
// This function iterates over all nodes in the graph and their neighbors,
// and returns a slice of edges, where each edge is represented as a
// *SingleTypedEdge[G].
func (g *UndirectGraphOf[G]) Edges() []*SingleTypedEdge[G] {
	// Create a set to store unique edges
	edges := sets.SortableOf[*SingleTypedEdge[G]]()

	// Get an iterator over the adjacency list
	iterator := g.adj.Iterator()

	// Iterate over all nodes and their neighbors
	for key, value, ok := iterator.Next(); ok; key, value, ok = iterator.Next() {
		// Iterate over all neighbors of the current node
		for _, neighbor := range value {
			// Check if the edge should be added in the "from -> to" direction
			if g.comparator.Compare(key, neighbor) < 0 {
				// Add the edge to the set
				edges.Add(NewEdge(key, neighbor))
			} else {
				// Add the edge in the "to -> from" direction
				edges.Add(NewEdge(neighbor, key))
			}
		}
	}
	// Return the slice of unique edges
	return edges.Values()
}

// String returns a string representation of the graph.
func (g *UndirectGraphOf[G]) String() string {
	return g.adj.String()
}
