// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/graph"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/datastr/dicts"
	"github.com/andrerrcosta2/gtools/datastr/graphs/edges"
	"github.com/andrerrcosta2/gtools/datastr/graphs/internal/ngb"
	"github.com/andrerrcosta2/gtools/datastr/sets"
)

// WeightedDigraphOf returns a new instance of graph.MutableWeighted with edge.ComparableWeighted edges
func WeightedDigraphOf[N constraints.Comparable[N], W any](nodes ...N) graph.MutableWeighted[edge.ComparableWeighted[N, W], N, W] {

	// Initialize the adjacency list with a map of dicts.
	g := &weightedOf[N, W]{
		adj:    dicts.ComparableOf[N, ngb.WComparableOf[N, W]](),
		direct: true,
	}

	//Add node to the graph
	for _, nod := range nodes {
		g.adj.Put(nod, dicts.ComparableOf[N, W]())
	}

	// Return the graph
	return g
}

type weightedOf[N constraints.Comparable[N], W any] struct {
	adj    dict.Mutable[N, ngb.WComparableOf[N, W]]
	direct bool
}

// AddNode adds a node to the graph.
func (g *weightedOf[N, W]) AddNode(node N) {
	dicts.PutIfAbsent(g.adj, node, ngb.NewWComparableOf[N, W]())
}

// AddEdge adds a directed, weighted edge from 'from' to 'to' with a given weight.
func (g *weightedOf[N, W]) AddEdge(from, to N, weight W) error {
	return addWeightedEdgeOfIfNodesExist(g.adj, from, to, weight, g.direct)
}

func (g *weightedOf[N, W]) Edges() []edge.ComparableWeighted[N, W] {
	e := sets.ComparableOf[edge.ComparableWeighted[N, W]]()
	fit := g.adj.Iterator()

	for from, tos, ok := fit.Next(); ok; from, tos, ok = fit.Next() {
		tit := tos.Iterator()
		for to, weight, ok := tit.Next(); ok; to, weight, ok = tit.Next() {
			e.Add(edges.ComparableWeighted(from, to, weight, g.direct))
		}
	}
	return e.Elements()
}

func (g *weightedOf[N, W]) EdgesFrom(from N) []edge.ComparableWeighted[N, W] {
	neighbors, ok := g.adj.Get(from)
	if !ok {
		return nil
	}
	return edges.BuildComparableWeightedFrom[N, W](from, neighbors, g.direct)
}

// HasNode checks if a node exists in the graph.
func (g *weightedOf[N, W]) HasNode(node N) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if a weighted edge exists from 'from' to 'to'.
func (g *weightedOf[N, W]) HasEdge(from, to N) bool {
	if neighbors, ok := g.adj.Get(from); ok {
		return neighbors.Contains(to)
	}
	return false
}

// Neighbors returns the neighbors of a node and their associated weights.
func (g *weightedOf[N, W]) Neighbors(node N) []N {
	if neighbors, ok := g.adj.Get(node); ok {
		// If the node exists, return its neighbors
		return neighbors.Keys()
	}
	return nil
}

// Nodes returns all node in the graph.
func (g *weightedOf[N, W]) Nodes() []N {
	return g.adj.Keys()
}

func (g *weightedOf[N, W]) Range(callback func(N, set.Readable[N]) bool) {
	it := g.adj.Iterator()
	for key, value, ok := it.Next(); ok; key, value, ok = it.Next() {
		if !callback(key, ngb.KeySet(value)) {
			return
		}
	}
}

func (g *weightedOf[N, W]) RangeNeighbors(node N, callback func(N) bool) {
	neighbors, ok := g.adj.Get(node)
	if !ok {
		return
	}
	it := neighbors.Iterator()
	for key, _, ok := it.Next(); ok; key, _, ok = it.Next() {
		if !callback(key) {
			return
		}
	}
}

func (g *weightedOf[N, W]) RemoveEdge(from, to N) bool {
	if neighbors, ok := g.adj.Get(from); ok {
		neighbors.Delete(to)
		return true
	}
	return false
}

func (g *weightedOf[N, W]) RemoveNode(node N) bool {
	g.adj.Delete(node)
	return true
}

// Weight returns the weight of the edge from 'from' to 'to'.
func (g *weightedOf[N, W]) Weight(from, to N) (weight W, exists bool) {
	if neighbors, ok := g.adj.Get(from); ok {
		weight, exists = neighbors.Get(to)
	}
	return
}

func (g *weightedOf[N, W]) String() string {
	return g.adj.String()
}

var _ graph.ReadableWeighted[edge.ComparableWeighted[sortableStub, any], sortableStub, any] = (*weightedOf[sortableStub, any])(nil)
var _ graph.MutableWeighted[edge.ComparableWeighted[sortableStub, any], sortableStub, any] = (*weightedOf[sortableStub, any])(nil)

// WeightedDigraph returns a new instance of graph.MutableWeighted with edge.ComparableWeighted edges.
func WeightedDigraph[N comparable, W any](nodes ...N) graph.MutableWeighted[edge.ComparableWeighted[N, W], N, W] {

	g := &weighted[N, W]{
		adj:    dicts.Comparable[N, ngb.WComparable[N, W]](),
		direct: true,
	}

	for _, nod := range nodes {
		g.adj.Put(nod, make(ngb.WComparable[N, W]))
	}

	// Return the graph
	return g
}

type weighted[N comparable, W any] struct {
	adj    dict.Mutable[N, ngb.WComparable[N, W]]
	direct bool
}

// AddNode adds a node to the graph.
func (g *weighted[N, W]) AddNode(node N) {
	dicts.PutIfAbsent(g.adj, node, make(ngb.WComparable[N, W]))
}

// AddEdge adds a directed, weighted edge from 'from' to 'to' with a given weight.
func (g *weighted[N, W]) AddEdge(from, to N, weight W) error {
	return addWeightedEdgeIfNodesExist(g.adj, from, to, weight, g.direct)
}

func (g *weighted[N, W]) Edges() []edge.ComparableWeighted[N, W] {
	e := sets.Comparable[edge.ComparableWeighted[N, W]]()
	fit := g.adj.Iterator()

	for from, tos, ok := fit.Next(); ok; from, tos, ok = fit.Next() {
		for to, weight := range tos {
			e.Add(edges.ComparableWeighted(from, to, weight, g.direct))
		}
	}
	return e.Elements()
}

func (g *weighted[N, W]) EdgesFrom(from N) []edge.ComparableWeighted[N, W] {
	neighbors, ok := g.adj.Get(from)
	if !ok {
		return nil
	}
	return edges.BuildComparableWeightedFrom[N, W](from, neighbors, g.direct)
}

// HasNode checks if a node exists in the graph.
func (g *weighted[N, W]) HasNode(node N) bool {
	// Check if the node exists in the adjacency list
	return g.adj.Contains(node)
}

// HasEdge checks if a weighted edge exists from 'from' to 'to'.
func (g *weighted[N, W]) HasEdge(from, to N) bool {
	if neighbors, ok := g.adj.Get(from); ok {
		return neighbors.Has(to)
	}
	return false
}

// Neighbors returns the neighbors of a node and their associated weights.
func (g *weighted[N, W]) Neighbors(node N) []N {
	if neighbors, ok := g.adj.Get(node); ok {
		// If the node exists, return its neighbors
		return neighbors.Elements()
	}
	return nil
}

// Nodes returns all node in the graph.
func (g *weighted[N, W]) Nodes() []N {
	return g.adj.Keys()
}

func (g *weighted[N, W]) Range(callback func(N, set.Readable[N]) bool) {
	it := g.adj.Iterator()
	for key, value, ok := it.Next(); ok; key, value, ok = it.Next() {
		if !callback(key, value) {
			return
		}
	}
}

func (g *weighted[N, W]) RangeNeighbors(node N, callback func(N) bool) {
	neighbors, ok := g.adj.Get(node)
	if !ok {
		return
	}
	neighbors.Range(callback)
}

func (g *weighted[N, W]) RemoveEdge(from, to N) bool {
	if neighbors, ok := g.adj.Get(from); ok {
		delete(neighbors, to)
		return true
	}
	return false
}

func (g *weighted[N, W]) RemoveNode(node N) bool {
	g.adj.Delete(node)
	return true
}

// Weight returns the weight of the edge from 'from' to 'to'.
func (g *weighted[N, W]) Weight(from, to N) (weight W, exists bool) {
	if neighbors, ok := g.adj.Get(from); ok {
		weight, exists = neighbors.Get(to)
	}
	return
}

func (g *weighted[N, W]) String() string {
	return g.adj.String()
}

var _ graph.ReadableWeighted[edge.ComparableWeighted[sortableStub, any], sortableStub, any] = (*weighted[sortableStub, any])(nil)
var _ graph.MutableWeighted[edge.ComparableWeighted[sortableStub, any], sortableStub, any] = (*weighted[sortableStub, any])(nil)
