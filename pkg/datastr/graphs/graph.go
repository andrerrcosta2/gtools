// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/core/gtools/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/datastr/maps"
	"github.com/andrerrcosta2/gtools/datastr/sets"
)

func AddNodeOfIfNotExists[G gtools.SortableOf](adj maps.StructMap[G, []G], node G) {
	if !hasNodeOf(adj, node) {
		adj.Put(node, []G{})
	}
}

func AddNodeIfNotExists[G prim.Ordered](adj *map[G][]G, node G) {
	if nodes, err := getNodes(adj, node); nodes.Len() == 0 && err == nil {
		(*adj)[node] = []G{}
	}
}

func AddUndirectEdgeOfIfNodesExist[G gtools.SortableOf](adj maps.StructMap[G, []G], from, to G) error {
	if !hasNodeOf(adj, from) || !hasNodeOf(adj, to) {
		return fmt.Errorf("The nodes %v and/or %v do not exist in the graph: %v\n", from, to, adj.Keys())
	}
	f, _ := adj.Get(from)
	t, _ := adj.Get(to)
	f = append(f, to)
	t = append(t, from)
	adj.Put(from, f)
	adj.Put(to, t)
	return nil
}

func AddUndirectEdgeIfNodesExist[G prim.Ordered](adj *map[G][]G, from, to G) error {
	nodes, err := getNodes(adj, from, to)
	if nodes.Len() != 2 {
		return fmt.Errorf("The nodes %v and/or %v do not exist in the graph: %v\n", from, to, nodes)
	}
	if err != nil {
		return err
	}

	froms, _ := nodes.Get(from)
	tos, _ := nodes.Get(to)
	froms = append(froms, to)
	tos = append(tos, from)
	work := *adj
	work[from] = froms
	work[to] = tos
	return nil
}

// AddDirectEdgeOfIfNodesExist adds a directed edge from 'from' to 'to' to the adjacency list.
// If 'from' or 'to' does not exist in the adjacency list, it will be created.
func AddDirectEdgeOfIfNodesExist[G gtools.SortableOf](adj *maps.SortableOfMap[G, []G], from, to G) error {
	if !hasNodeOf[G](adj, from) || !hasNodeOf[G](adj, to) {
		return fmt.Errorf("The nodes %v and/or %v do not exist in the graph: %v\n", from, to, adj.Keys())
	}
	// For a pattern reason, the nodes must exist before adding of be added as an edge
	if f, ok := adj.Get(from); ok {
		f = append(f, to)
		adj.Put(from, f)
		return nil
	}
	adj.Put(from, []G{to})
	return nil
}

func AddDirectEdgeIfNodesExist[G prim.Ordered](adj *map[G][]G, from, to G) error {
	nodes, err := getNodes(adj, from, to)
	if nodes.Len() != 2 {
		return fmt.Errorf("The nodes %v and/or %v do not exist in the graph: %v\n", from, to, nodes)
	}
	if err != nil {
		return err
	}

	froms, _ := nodes.Get(from)
	froms = append(froms, to)
	(*adj)[from] = froms
	return nil
}

// AddWeightedEdgeOfIfNodesExist adds a weighted edge from 'from' to 'to' with a given weight to the adjacency list.
// If 'from' or 'to' does not exist in the adjacency list, it will be created.
func AddWeightedEdgeOfIfNodesExist[G gtools.SortableOf, W any](adj *maps.SortableOfMap[G, *maps.SortableOfMap[G, W]], from, to G, weight W) error {
	if !adj.Contains(from) || !adj.Contains(to) {
		return fmt.Errorf("The nodes %v and/or %v do not exist in the graph: %v\n", from, to, adj.Keys())
	}
	// For a pattern reason, the nodes must exist before adding of be added as an edge
	mfrom, _ := adj.Get(from)
	mfrom.Put(to, weight)
	return nil
}

func hasNodeOf[G gtools.SortableOf](adj maps.StructMap[G, []G], node G) bool {
	return adj.Contains(node)
}

func getNodes[G prim.Ordered](adj *map[G][]G, nodes ...G) (*maps.EntrySet[G, []G], error) {
	out := maps.NewEntrySet[G, []G]()
	if len(nodes) == 0 {
		return out, nil
	}
	work := *adj
	for _, node := range nodes {
		if n, ok := (work)[node]; ok {
			err := out.Add(maps.NewComparableEntry(node, n))
			if err != nil {
				return nil, err
			}
		}
	}
	return out, nil
}

// IsCyclic This is tricky, I'm not sure about that. the problem is
// you can't type check constraints because type check can't produce a false-positive.
// This isn't a reliable pattern, and i should change it as soon as i find a better solution
// that doesn't demand to define multiple generics. It could be done using 3 generics
// then type checking against correct interfaces. But the real assertion for this method
// is against the type used for comparison, not the type of the graph.
func IsCyclic[G any, K prim.Ordered](graph str.Graph[G]) (bool, error) {
	if g, ok := any(graph).(str.Graph[gtools.SortableOf]); ok {
		return isCyclicOf(g), nil
	}

	if g, ok := any(graph).(str.Graph[K]); ok {
		return isCyclicOrdered(g), nil
	}

	return false, errors.New("the type does not implement the interface graph.Graph\n")
}

func IsAcyclic[G any, K prim.Ordered](graph str.Graph[G]) (bool, error) {
	if g, ok := any(graph).(str.Graph[gtools.SortableOf]); ok {
		return isAcyclicOf(g), nil
	}

	if g, ok := any(graph).(str.Graph[K]); ok {
		return isAcyclicOrdered(g), nil
	}

	return false, errors.New("the type does not implement the interface graph.Graph\n")
}

// IsCyclic checks if the graph is cyclic.
// A graph is considered cyclic if there's a path that starts and ends at the same node.
func isCyclicOrdered[G prim.Ordered](g str.Graph[G]) bool {
	visited := sets.Comparable[G]()
	recStack := sets.Comparable[G]()

	// Here it searches vertically. Nodes are map keys
	for _, node := range g.Nodes() {
		if !visited.Has(node) {
			if searchCycleVertically[G, G](visited, recStack, g, node) {
				return true
			}
		}
	}
	return false
}

// IsCyclicOf checks if the graph is cyclic.
// A graph is considered cyclic if there is a path that starts and ends at the same node.
func isCyclicOf[T gtools.SortableOf](g str.Graph[T]) bool {
	visited := sets.SortableOf[T]()
	recStack := sets.SortableOf[T]()

	// Here it searches vertically. Nodes are map keys
	for _, node := range g.Nodes() {
		if searchCycleVertically[T, string](visited, recStack, g, node) {
			return true
		}
	}
	return false
}

func isAcyclicOrdered[T prim.Ordered](g str.Graph[T]) bool {
	return !isCyclicOrdered(g)
}

// IsAcyclicOf checks if the graph is acyclic.
// A graph is considered acyclic if it does not contain any cycles.
func isAcyclicOf[T gtools.SortableOf](g str.Graph[T]) bool {
	// Check if the graph is cyclic and return the opposite result
	return !isCyclicOf(g)
}

func searchCycleVertically[G any, K prim.Ordered](visited, recStack str.Set[G], g str.IterableGraph[G], node G) bool {
	return searchCycleHorizontally[G, K](visited, recStack, g, node, node)
}

func searchCycleHorizontally[G any, K prim.Ordered](visited, recStack str.Set[G], g str.IterableGraph[G], start, node G) bool {
	// Mark the node as visited and add it to the recursion stack
	visited.Add(node)
	recStack.Add(node)

	neighborhood := g.Neighbors(node)
	// Get neighbors and recursively check for cycles
	for _, neighbor := range neighborhood {
		// Self-loop detection
		if sortables.TryEquality[G](neighbor, node) {
			//fmt.Println("Will print true")
			return true
		}
		// Avoids considering the direction as a cycle when it's bidirectional
		if !visited.Has(neighbor) {
			if searchCycleHorizontally[G, K](visited, recStack, g, node, neighbor) {
				return true
			}
		} else if recStack.Has(neighbor) && !sortables.TryEquality[G](neighbor, start) {
			//fmt.Println("Will print true")
			// If the neighbor is visited and isn't the parent, then there's a cycle
			return true
		}
	}

	// Remove the node from the recursion stack
	recStack.Remove(node)
	return false
}

// IsConnectedOf returns true if the undirected graph is connected.
// A graph is considered connected if there's a path between every pair of nodes.
func IsConnectedOf[T gtools.SortableOf](g str.GraphOf[T, edges.SortableSingleTyped[T]]) bool {
	// Handle the edge case where the graph has no nodes
	if len(g.Nodes()) == 0 {
		return true // An empty graph is considered connected
	}

	// Create a map to keep track of visited nodes
	visited := sets.SortableOf[T]() // Use unique string identifiers for nodes

	// Define a recursive function to visit nodes
	var visit func(T)
	visit = func(node T) {
		visited.Add(node) // Mark the node as visited

		// Visit all unvisited neighbors of the current node
		for _, neighbor := range g.Neighbors(node) {
			if !visited.Has(neighbor) {
				visit(neighbor) // Recursively visit the neighbor
			}
		}
	}

	// Get all nodes in the graph
	nodes := g.Nodes()

	// Start the visitation process from the first node
	visit(nodes[0])

	// The graph is connected if all nodes were visited
	return visited.Len() == len(g.Nodes())
}

// IsDisconnectedOf returns true if the undirected graph is disconnected.
// A graph is considered disconnected if there is no path between every pair of nodes.
func IsDisconnectedOf[T gtools.SortableOf](g str.GraphOf[T, edges.SortableSingleTyped[T]]) bool {
	// Simply return the opposite of IsConnectedOf, as a graph is disconnected if it's not connected
	return !IsConnectedOf[T](g)
}
