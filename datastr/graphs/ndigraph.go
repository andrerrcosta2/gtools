// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/datastr/maps"
	"github.com/andrerrcosta2/gtools/datastr/sets"
)

// UndirectedOf creates a new undirectGraphOf instance.
// It returns a pointer to a undirectGraphOf struct with an initialized adjacency list.
//
// This function is used to initialize a new undirected graph with a list of nodes.
func UndirectedOf[G gtools.SortableOf](nods ...G) str.GraphOf[G, edges.SortableSingleTyped[G]] {
	// Create a new undirectGraphOf instance with the given nodes in the graph
	// The adjacency list is represented as a map of type N to a slice of type N.
	// The map is initialized using the maps.SortableOf function from the gtools package.
	// The comparator is initialized using the sortables.ComparatorOf function from the gtools package.
	graph := &undirectGraphOf[G]{
		// Initialize the adjacency list (adj) as an empty map.
		adj: maps.SortableOf[G, []G](),

		// Initialize the comparator for the graph.
		comparator: sortables.ComparatorOf[G](),
	}

	// Addf nodes to the graph
	for _, nod := range nods {
		graph.adj.Put(nod, []G{})
	}

	// Return the graph
	return graph
}

var _ str.Graph[gtools.SortableOf] = (*undirectGraphOf[gtools.SortableOf])(nil)
var _ str.SingleEdgedGraph[gtools.SortableOf, edges.SortableSingleTyped[gtools.SortableOf]] = (*undirectGraphOf[gtools.SortableOf])(nil)
var _ str.GraphOf[gtools.SortableOf, edges.SortableSingleTyped[gtools.SortableOf]] = (*undirectGraphOf[gtools.SortableOf])(nil)

// undirectGraphOf is a basic implementation of an undirected graph using an adjacency list.
type undirectGraphOf[G gtools.SortableOf] struct {
	adj        str.Map[G, []G]
	comparator comparators.KeyTyped[G, string]
}

// AddNode adds a node to the graph.
// It checks if the node already exists in the graph before adding it.
// If the node doesn't exist, it adds the node to the adjacency list with an empty slice of neighbors.
func (g *undirectGraphOf[G]) AddNode(node G) {
	AddNodeOfIfNotExists[G](g.adj, node)
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
func (g *undirectGraphOf[G]) AddEdge(from, to G) error {
	return AddUndirectEdgeOfIfNodesExist[G](g.adj, from, to)
}

// Neighbors returns the neighbors of a node in the undirected graph.
// It returns a slice of nodes that are directly connected to the given node.
func (g *undirectGraphOf[G]) Neighbors(node G) []G {
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
func (g *undirectGraphOf[G]) HasNode(node G) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if an edge exists between two nodes.
// It returns true if there is an edge from 'from' to 'to', false otherwise.
func (g *undirectGraphOf[G]) HasEdge(from, to G) bool {
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
func (g *undirectGraphOf[G]) Nodes() []G {
	// Get the keys from the adjacency list, which represent the nodes in the graph.
	return g.adj.Keys()
}

// Edges returns all edges in the graph.
// Edges returns all edges in the graph.
// This function iterates over all nodes in the graph and their neighbors,
// and returns a slice of edges, where each edge is represented as a
// *SingleTyped[N].
func (g *undirectGraphOf[G]) Edges() []edges.SortableSingleTyped[G] {
	// Create a set to store unique edges
	edges := sets.SortableOf[edges.SortableSingleTyped[G]]()

	// Get an iterator over the adjacency list
	iterator := g.adj.Iterator()

	// Iterate over all nodes and their neighbors
	for key, value, ok := iterator.Next(); ok; key, value, ok = iterator.Next() {
		// Iterate over all neighbors of the current node
		for _, neighbor := range value {
			// Check if the edge should be added in the "from -> to" direction
			if g.comparator.Compare(key, neighbor) < 0 {
				// Addf the edge to the set
				edges.Add(SortableEdge(key, neighbor, false))
			} else {
				// Addf the edge in the "to -> from" direction
				edges.Add(SortableEdge(neighbor, key, false))
			}
		}
	}
	// Return the slice of unique edges
	return edges.Values()
}

// String returns a string representation of the graph.
func (g *undirectGraphOf[G]) String() string {
	return g.adj.String()
}

// Undirected creates a new undirected graph.
// It returns a pointer to the new graph.
//
// The function creates an empty graph and returns a pointer to it.
func Undirected[G prim.Ordered](nods ...G) str.OrderedGraph[G, edges.SortableSingleTyped[G]] {
	// Create an empty graph
	graph := &undirectGraph[G]{
		adj: make(map[G][]G),
	}

	// Addf nodes to the graph
	for _, nod := range nods {
		graph.AddNode(nod)
	}

	// Return the graph
	return graph
}

var _ str.Graph[int] = (*undirectGraph[int])(nil)
var _ str.SingleEdgedGraph[int, edges.SortableSingleTyped[int]] = (*undirectGraph[int])(nil)
var _ str.OrderedGraph[int, edges.SortableSingleTyped[int]] = (*undirectGraph[int])(nil)

// undirectGraph is a graph with undirected edges.
type undirectGraph[G prim.Ordered] struct {
	adj map[G][]G
}

// AddNode adds a node to the graph.
// It checks if the node already exists in the graph before adding it.
// If the node doesn't exist, it adds the node to the adjacency list with an empty slice of neighbors.
func (g *undirectGraph[G]) AddNode(node G) {
	AddNodeIfNotExists(&g.adj, node)
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
func (g *undirectGraph[G]) AddEdge(from, to G) error {
	return AddUndirectEdgeIfNodesExist(&g.adj, from, to)
}

// Neighbors returns the neighbors of a node in the undirected graph.
// It returns a slice of nodes that are directly connected to the given node.
func (g *undirectGraph[G]) Neighbors(node G) []G {
	// Attempt to retrieve the neighbors of the node from the adjacency list
	if neighbors, ok := g.adj[node]; ok {
		// If the node exists, return its neighbors
		return neighbors
	}
	// If the node does not exist, return an empty slice
	return nil
}

// HasNode checks if a node exists in the graph.
// It returns true if the node is found, false otherwise.
func (g *undirectGraph[G]) HasNode(node G) bool {
	// Check if the node exists in the adjacency list
	_, ok := g.adj[node]
	return ok
}

// HasEdge checks if an edge exists between two nodes.
// It returns true if there is an edge from 'from' to 'to', false otherwise.
func (g *undirectGraph[G]) HasEdge(from, to G) bool {
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

// Nodes returns a list of all nodes in the graph.
func (g *undirectGraph[G]) Nodes() []G {
	// Create a slice to store the nodes
	out := make([]G, 0, len(g.adj))
	// Get the keys from the adjacency list, which represent the nodes in the graph.
	for key := range g.adj {
		out = append(out, key)
	}
	// Return the slice of nodes
	return out
}

// Edges returns all edges in the graph.
// Edges returns all edges in the graph.
// This function iterates over all nodes in the graph and their neighbors,
// and returns a slice of edges, where each edge is represented as a
// *SingleTyped[N].
func (g *undirectGraph[G]) Edges() []edges.SortableSingleTyped[G] {
	// Create a set to store unique edges
	edges := sets.SortableOf[edges.SortableSingleTyped[G]]()

	// Iterate over all nodes and their neighbors
	for key, value := range g.adj {
		// Iterate over all neighbors of the current node
		for _, neighbor := range value {
			// Check if the edge should be added in the "from -> to" direction
			if key < neighbor {
				// Addf the edge to the set
				edges.Add(SortableEdge(key, neighbor, false))
			} else {
				// Addf the edge in the "to -> from" direction
				edges.Add(SortableEdge(neighbor, key, false))
			}
		}
	}
	// Return the slice of unique edges
	return edges.Values()
}

// String returns a string representation of the graph.
func (g *undirectGraph[G]) String() string {
	return fmt.Sprintf("%v", g.adj)
}
