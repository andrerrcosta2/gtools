// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"fmt"

	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/graph"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/datastr/dicts"
	"github.com/andrerrcosta2/gtools/datastr/graphs/edges"
	"github.com/andrerrcosta2/gtools/datastr/graphs/internal/ngb"
	"github.com/andrerrcosta2/gtools/datastr/sets"
)

// DigraphOf creates a new direct graph.MutableEdged instance
// whose edges implements constraints.Comparable
func DigraphOf[N constraints.Comparable[N]](nodes ...N) graph.MutableEdged[edge.Comparable[N], N] {
	g := &graphOf[N]{
		adj:    dicts.ComparableOf[N, ngb.ComparableOf[N]](),
		direct: true,
	}

	for _, nod := range nodes {
		g.adj.Put(nod, sets.ComparableOf[N]())
	}

	return g
}

// UndirectedOf creates a new undirect graph.MutableEdged instance
// whose edges implements constraints.Comparable
func UndirectedOf[N constraints.Comparable[N]](nodes ...N) graph.MutableEdged[edge.Comparable[N], N] {
	return undirectedOf(nodes...)
}

func undirectedOf[N constraints.Comparable[N]](nodes ...N) *graphOf[N] {
	g := &graphOf[N]{
		adj:    dicts.ComparableOf[N, ngb.ComparableOf[N]](),
		direct: false,
	}

	for _, nod := range nodes {
		g.adj.Put(nod, sets.ComparableOf[N]())
	}

	return g
}

type S = sortableStub

var _ graph.Readable[S] = (*graphOf[S])(nil)
var _ graph.Iterable[S] = (*graphOf[S])(nil)

// DirectedGraphOf is a basic implementation of a directed graph using an adjacency list.
type graphOf[N constraints.Comparable[N]] struct {
	adj    dict.Mutable[N, ngb.ComparableOf[N]]
	direct bool
}

// AddEdge adds a directed edge from 'from' to 'to'.
func (g *graphOf[N]) AddEdge(from, to N) error {
	return addEdgeOfIfNodesExist(g.adj, from, to, true)
}

// AddNode adds a node to the graph.
func (g *graphOf[N]) AddNode(node N) {
	addNodeOfIfAbsent(g.adj, node)
}

// Edges returns all directed edge in the graph.
func (g *graphOf[N]) Edges() []edge.Comparable[N] {
	e := sets.ComparableOf[edge.Comparable[N]]()

	it1 := g.adj.Iterator()

	// Iterate over all node and their neighbors
	for key, neighbors, ok := it1.Next(); ok; key, neighbors, ok = it1.Next() {
		neighbors.Range(func(neighbor N) bool {
			e.Add(edges.Comparable[N](key, neighbor, true))
			return true
		})
	}
	// Return the slice of unique edge
	return e.Elements()
}

func (g *graphOf[N]) EdgesFrom(from N) []edge.Comparable[N] {
	if neighbors, ok := g.adj.Get(from); ok {
		out := make([]edge.Comparable[N], neighbors.Len())
		neighbors.Range(func(neighbor N) bool {
			out = append(out, edges.Comparable(from, neighbor, true))
			return true
		})
	}
	// If 'from' does not exist or 'to' is not a neighbor of 'from', return false
	return []edge.Comparable[N]{}
}

// HasNode checks if a node exists in the graph.
func (g *graphOf[N]) HasNode(node N) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if a directed edge exists from 'from' to 'to'.
func (g *graphOf[N]) HasEdge(from, to N) bool {
	if neighbors, ok := g.adj.Get(from); ok {
		return neighbors.Has(to)
	}
	return false
}

// Neighbors returns the outgoing neighbors of a node.
func (g *graphOf[N]) Neighbors(node N) []N {
	if neighbors, ok := g.adj.Get(node); ok {
		return neighbors.Elements()
	}
	return nil
}

// Nodes return all node in the graph.
func (g *graphOf[N]) Nodes() []N {
	return g.adj.Keys()
}

func (g *graphOf[N]) RemoveEdge(from, to N) bool {
	neighbors, ok := g.adj.Get(from)
	if !ok {
		return false
	}
	neighbors.Remove(to)
	return true
}

func (g *graphOf[N]) Range(f func(N, set.Readable[N]) bool) {
	it := g.adj.Iterator()
	for key, value, ok := it.Next(); ok; key, value, ok = it.Next() {
		if !f(key, value) {
			return
		}
	}
}

func (g *graphOf[N]) RangeNeighbors(node N, f func(N) bool) {
	if neighbors, ok := g.adj.Get(node); ok {
		neighbors.Range(f)
	}
}

func (g *graphOf[N]) RemoveNode(node N) bool {
	g.adj.Delete(node)
	return true
}

// String returns a string representation of the graph.
func (g *graphOf[N]) String() string {
	return g.adj.String()
}

// Digraph creates a new DirectedGraph instance.
//
// It initializes a DirectedGraph with the given node.
//
// Returns:
//
//	*DirectedGraph[N]: A new DirectedGraph instance with the given node.
func Digraph[N prim.Ordered](nods ...N) graph.MutableEdged[edge.Comparable[N], N] {
	// Create a new DirectedGraph instance with the given node
	g := &directedGraph[N]{
		// Initialize the adjacency list as a new, empty map.
		adj: make(map[N]ngb.Comparable[N], len(nods)),
	}

	for _, nod := range nods {
		g.adj[nod] = make(ngb.Comparable[N])
	}

	return g
}

var _ graph.Readable[int] = (*directedGraph[int])(nil)
var _ graph.ReadableEdged[edge.Comparable[int], int] = (*directedGraph[int])(nil)

// DirectedGraph is a basic implementation of a directed graph using an adjacency list.
type directedGraph[N comparable] struct {
	adj map[N]ngb.Comparable[N]
}

// AddNode adds a node to the graph.
func (g *directedGraph[N]) AddNode(node N) {
	addNodeIfAbsent[N](g.adj, node)
}

// AddEdge adds a directed edge from 'from' to 'to'.
func (g *directedGraph[N]) AddEdge(from, to N) error {
	return addEdgeIfNodesExist(g.adj, from, to, true)
}

func (g *directedGraph[N]) EdgesFrom(from N) []edge.Comparable[N] {
	if neighbors, ok := g.adj[from]; ok {
		out := make([]edge.Comparable[N], len(neighbors))
		// Iterate over the neighbors
		for neighbor := range neighbors {
			out = append(out, edges.Comparable(from, neighbor, true))
		}
	}
	// If 'from' does not exist or 'to' is not a neighbor of 'from', return false
	return []edge.Comparable[N]{}
}

// Neighbors returns the outgoing neighbors of a node.
func (g *directedGraph[N]) Neighbors(node N) []N {
	// Attempt to retrieve the neighbors of the node from the adjacency list
	if neighbors, ok := g.adj[node]; ok {
		// If the node exists, return its neighbors
		return neighbors.Elements()
	}
	// If the node doesn't exist, return an empty slice
	return []N{}
}

// HasNode checks if a node exists in the graph.
func (g *directedGraph[N]) HasNode(node N) bool {
	_, ok := g.adj[node]
	return ok
}

// HasEdge checks if a directed edge exists from 'from' to 'to'.
func (g *directedGraph[N]) HasEdge(from, to N) bool {
	if neighbors, ok := g.adj[from]; ok {
		_, exists := neighbors[to]
		return exists
	}
	return false
}

// Nodes return all node in the graph.
func (g *directedGraph[N]) Nodes() []N {
	// Get the keys from the adjacency list, which represent the node in the graph.
	out := make([]N, 0, len(g.adj))
	for key := range g.adj {
		out = append(out, key)
	}
	// Return the slice of node
	return out
}

func (g *directedGraph[N]) Edges() []edge.Comparable[N] {
	e := sets.Comparable[edge.Comparable[N]]()
	for from, neighbors := range g.adj {
		for to := range neighbors {
			e.Add(edges.Comparable[N](from, to, true))
		}
	}
	return e.Elements()
}

func (g *directedGraph[N]) Range(f func(N, set.Readable[N]) bool) {
	for key, value := range g.adj {
		if !f(key, value) {
			return
		}
	}
}

func (g *directedGraph[N]) RangeNeighbors(node N, f func(N) bool) {
	if neighbors, ok := g.adj[node]; ok {
		neighbors.Range(f)
	}
}

func (g *directedGraph[N]) RemoveEdge(from, to N) bool {
	if neighbors, ok := g.adj[from]; ok {
		delete(neighbors, to)
		return true
	}
	return false
}

func (g *directedGraph[N]) RemoveNode(node N) bool {
	delete(g.adj, node)
	return true
}

// String returns a string representation of the graph.
func (g *directedGraph[N]) String() string {
	return fmt.Sprintf("%v", g.adj)
}
