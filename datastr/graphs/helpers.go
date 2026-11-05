// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"fmt"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/graph"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/datastr/graphs/internal/ngb"
	"github.com/andrerrcosta2/gtools/datastr/sets"
)

// AddNodeOfIfAbsent Add a new node to the
func AddNodeOfIfAbsent[N constraints.Comparable[N]](g graph.Mutable[N], node N) {
	if sg, ok := g.(structGraph[N]); ok {
		addNodeOfIfAbsent(sg.adj(), node)
	}
}

func addNodeOfIfAbsent[N constraints.Comparable[N]](adj ngb.StructDict[N], node N) {
	if !adj.Contains(node) {
		adj.Put(node, sets.ComparableOf[N]())
	}
}

func AddNodeIfAbsent[N comparable](g graph.Mutable[N], node N) {
	if cg, ok := g.(cmpGraph[N]); ok {
		addNodeIfAbsent(cg.adj(), node)
	}
}

func addNodeIfAbsent[N comparable](adj ngb.CmpDict[N], node N) {
	if _, ok := adj[node]; !ok {
		adj[node] = make(ngb.Comparable[N])
	}
}

func AddEdgeOfIfNodesExist[N constraints.Comparable[N]](g graph.Mutable[N], from, to N, direct bool) error {
	if sg, ok := g.(structGraph[N]); ok {
		return addEdgeOfIfNodesExist(sg.adj(), from, to, direct)
	}
	return fmt.Errorf("expected a graph of constraints.Comparable[N] nodes but got '%v'", g)
}

func addEdgeOfIfNodesExist[N constraints.Comparable[N]](adj ngb.StructDict[N], from, to N, direct bool) error {
	froms, ok := adj.Get(from)
	if !ok {
		return fmt.Errorf("The node '%v' does not exist in the graph: '%v'\n", from, adj)
	}
	tos, ok := adj.Get(to)
	if !ok {
		return fmt.Errorf("The node '%v' does not exist in the graph: '%v'\n", to, adj)
	}

	froms.Add(to)
	if !direct {
		tos.Add(from)
	}
	return nil
}

func AddEdgeIfNodesExist[N comparable](g graph.Mutable[N], from, to N, direct bool) error {
	if sg, ok := g.(cmpGraph[N]); ok {
		return addEdgeIfNodesExist(sg.adj(), from, to, direct)
	}
	return fmt.Errorf("expected a graph of comparable nodes but got '%v'", g)
}

func addEdgeIfNodesExist[N comparable](adj ngb.CmpDict[N], from, to N, direct bool) error {
	froms, ok := adj[from]
	if !ok {
		return fmt.Errorf("The node '%v' does not exist in the graph: '%v'\n", from, adj)
	}
	tos, ok := adj[to]
	if !ok {
		return fmt.Errorf("The node '%v' does not exist in the graph: '%v'\n", to, adj)
	}
	froms.Add(to)
	if !direct {
		tos.Add(from)
	}
	return nil
}

func AddWeightedEdgeOfIfNodesExist[E edge.Weighted[N, W], N constraints.Comparable[N], W any](
	g graph.MutableWeighted[E, N, W],
	from, to N, weight W, direct bool) error {
	if sg, ok := g.(wStructGraph[N, W]); ok {
		return addWeightedEdgeOfIfNodesExist(sg.adj(), from, to, weight, direct)
	}
	return fmt.Errorf("expected a graph of weighted constraints.Comparable[N] nodes but got '%v'", g)
}

// addWeightedEdgeOfIfNodesExist adds a weighted edge from 'from' to 'to' with a given weight to the adjacency list.
// If 'from' or 'to' does not exist in the adjacency list, it will be created.
func addWeightedEdgeOfIfNodesExist[N constraints.Comparable[N], W any](
	adj ngb.WStructDict[N, W],
	from, to N, weight W, direct bool) error {
	froms, ok := adj.Get(from)
	if !ok {
		return fmt.Errorf("The node '%v' does not exist in the graph: '%v'\n", from, adj)
	}
	tos, ok := adj.Get(to)
	if !ok {
		return fmt.Errorf("The node '%v' does not exist in the graph: '%v'\n", to, adj)
	}
	froms.Put(to, weight)
	if !direct {
		tos.Put(from, weight)
	}
	return nil
}

func AddWeightedEdgeIfNodesExist[E edge.Weighted[N, W], N comparable, W any](
	g graph.MutableWeighted[E, N, W],
	from, to N, weight W, direct bool) error {
	if sg, ok := g.(wCmpGraph[N, W]); ok {
		return addWeightedEdgeIfNodesExist(sg.adj(), from, to, weight, direct)
	}
	return fmt.Errorf("expected a graph of weighted constraints.Comparable[N] nodes but got '%v'", g)
}

// addWeightedEdgeOfIfNodesExist adds a weighted edge from 'from' to 'to' with a given weight to the adjacency list.
// If 'from' or 'to' does not exist in the adjacency list, it will be created.
func addWeightedEdgeIfNodesExist[N comparable, W any](
	adj ngb.WCmpDict[N, W],
	from, to N, weight W, direct bool) error {
	froms, ok := adj.Get(from)
	if !ok {
		return fmt.Errorf("The node '%v' does not exist in the graph: '%v'\n", from, adj)
	}
	tos, ok := adj.Get(to)
	if !ok {
		return fmt.Errorf("The node '%v' does not exist in the graph: '%v'\n", to, adj)
	}
	froms[to] = weight
	if !direct {
		tos[from] = weight
	}
	return nil
}

// IsCyclic reports whether a graph has a cycle
func IsCyclic[N any, H comparable](g graph.Iterable[N], cmp comparator.EqCmpKeyer[N, H]) bool {
	visited := sets.HashSet(cmp)
	recStack := sets.HashSet(cmp)
	return isCyclic(g, visited, recStack)
}

// IsCyclicOf checks if the graph is cyclic.
// A graph is considered cyclic if there is a path that starts and ends at the same node.
func isCyclic[N any](g graph.Iterable[N], visited, recStack set.Mutable[N]) bool {
	hasCycle := false
	g.Range(func(node N, _ set.Readable[N]) bool {
		if !visited.Has(node) {
			visited.Add(node)
			hasCycle = searchCycleHorizontally[N](g, visited, recStack, node)
		}
		return hasCycle
	})
	return hasCycle
}

func searchCycleHorizontally[N any](
	g graph.Iterable[N],
	visited, recursion set.Mutable[N],
	node N) bool {

	if recursion.Has(node) {
		return true
	}
	recursion.Add(node)
	defer recursion.Remove(node)

	found := false

	g.RangeNeighbors(node, func(neighbor N) bool {
		if !visited.Has(neighbor) {
			visited.Add(neighbor)

			if searchCycleHorizontally(g, visited, recursion, neighbor) {
				found = true
				return false // interrompe a iteração
			}
		}
		return true
	})

	return found
}

// IsConnected reports whether the undirected graph is connected.
// A graph is considered connected if there's a path between every pair of node.
func IsConnected[N any, H comparable](g graph.Iterable[N], cmp comparator.EqCmpKeyer[N, H]) bool {
	nodes := g.Nodes()
	// Handle the edge case where the graph has no node
	if len(nodes) == 0 {
		return true // An empty graph is considered connected
	}

	visited := sets.HashSet(cmp)

	var visit func(N)
	visit = func(node N) {
		visited.Add(node)

		for _, neighbor := range g.Neighbors(node) {
			if !visited.Has(neighbor) {
				visit(neighbor)
			}
		}
	}

	visit(nodes[0])

	// The graph is connected if all node were visited
	return visited.Len() == len(g.Nodes())
}
