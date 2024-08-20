// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/datastr/maps"
	"github.com/andrerrcosta2/gtools/pkg/datastr/sets"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/sortables"
)

func addNodeOfIfNotExists[G gtools.SortableOf](adj *maps.SortableOfMap[G, []G], node G) {
	if !hasNodeOf(adj, node) {
		adj.Put(node, []G{})
	}
}

func hasNodeOf[G gtools.SortableOf](adj *maps.SortableOfMap[G, []G], node G) bool {
	return adj.Contains(node)
}

func addUndirectEdgeOfIfNodesExist[G gtools.SortableOf](adj *maps.SortableOfMap[G, []G], from, to G) {
	if !hasNodeOf(adj, from) || !hasNodeOf(adj, to) {
		return
	}
	f, _ := adj.Get(from)
	t, _ := adj.Get(to)
	f = append(f, to)
	t = append(t, from)
	adj.Put(from, f)
	adj.Put(to, t)
}

func addDirectEdgeOfIfNodesExist[G gtools.SortableOf](adj *maps.SortableOfMap[G, []G], from, to G) {
	// For a pattern reason, the nodes must exist before adding of be added as an edge
	if neighbors, ok := adj.Get(from); ok {
		// Check if 'to' exists in the graph
		if _, sure := adj.Get(to); !sure {
			return
		}
		// If 'from' exists, append 'to' to the list of neighbors of 'from'.
		neighbors = append(neighbors, to)
		// Update the adjacency list
		adj.Put(from, neighbors)
	}
}

// addWeightedEdgeOfIfNodesExist adds a weighted edge from 'from' to 'to' with a given weight to the adjacency list.
// If 'from' or 'to' does not exist in the adjacency list, it will be created.
func addWeightedEdgeOfIfNodesExist[G gtools.SortableOf, W any](adj *maps.SortableOfMap[G, *maps.SortableOfMap[G, W]], from, to G, weight W) {
	// For a pattern reason, the nodes must exist before adding of be added as an edge
	if mfrom, ok := adj.Get(from); ok {
		// Check if 'to' exists in the graph
		if _, sure := adj.Get(to); !sure {
			return
		}
		// If 'from' and to exists, append 'to' to the list of neighbors of 'from'.
		mfrom.Put(to, weight)
		m := mfrom
		fmt.Printf("mfrom: %v\n", m)
	}
}

// IsCyclic This is tricky, I'm not sure about that. the problem is
// you can't type check constraints because type check can't produce a false-positive.
// This isn't a reliable pattern, and i should change it as soon as i find a better solution
// that doesn't demand to define multiple generics. It could be done using 3 generics
// then type checking against correct interfaces. But the real assertion for this method
// is against the type used for comparison, not the type of the graph.
func IsCyclic[G any, K constraints.Ordered](graph Graph[G]) (bool, error) {
	if g, ok := any(graph).(Graph[gtools.SortableOf]); ok {
		return isCyclicOf(g), nil
	}

	if g, ok := any(graph).(Graph[K]); ok {
		return isCyclicOrdered(g), nil
	}

	return false, errors.New("the type does not implement the interface graph.Graph\n")
}

func IsAcyclic[G any, K constraints.Ordered](graph Graph[G]) (bool, error) {
	if g, ok := any(graph).(Graph[gtools.SortableOf]); ok {
		return isAcyclicOf(g), nil
	}

	if g, ok := any(graph).(Graph[K]); ok {
		return isAcyclicOrdered(g), nil
	}

	return false, errors.New("the type does not implement the interface graph.Graph\n")
}

// IsCyclic checks if the graph is cyclic.
// A graph is considered cyclic if there's a path that starts and ends at the same node.
func isCyclicOrdered[G constraints.Ordered](g Graph[G]) bool {
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
func isCyclicOf[T gtools.SortableOf](g Graph[T]) bool {
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

func isAcyclicOrdered[T constraints.Ordered](g Graph[T]) bool {
	return !isCyclicOrdered(g)
}

// IsAcyclicOf checks if the graph is acyclic.
// A graph is considered acyclic if it does not contain any cycles.
func isAcyclicOf[T gtools.SortableOf](g Graph[T]) bool {
	// Check if the graph is cyclic and return the opposite result
	return !isCyclicOf(g)
}

func searchCycleVertically[G any, K constraints.Ordered](visited, recStack sets.Set[G], g Iterable[G], node G) bool {
	return searchCycleHorizontally[G, K](visited, recStack, g, node, node)
}

func searchCycleHorizontally[G any, K constraints.Ordered](visited, recStack sets.Set[G], g Iterable[G], start, node G) bool {
	// Mark the node as visited and add it to the recursion stack
	visited.Add(node)
	recStack.Add(node)

	neighborhood := g.Neighbors(node)
	// Get neighbors and recursively check for cycles
	for _, neighbor := range neighborhood {
		// Self-loop detection
		if sortables.Equals[K](neighbor, node) {
			//fmt.Println("Will print true")
			return true
		}
		// Avoids considering the direction as a cycle when it's bidirectional
		if !visited.Has(neighbor) {
			if searchCycleHorizontally[G, K](visited, recStack, g, node, neighbor) {
				return true
			}
		} else if recStack.Has(neighbor) && !sortables.Equals[K](neighbor, start) {
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
func IsConnectedOf[T gtools.SortableOf](g *UndirectGraphOf[T]) bool {
	// Handle the edge case where the graph has no nodes
	if len(g.Nodes()) == 0 {
		return true // An empty graph is considered connected
	}

	// Create a map to keep track of visited nodes
	visited := sets.Comparable[string]() // Use unique string identifiers for nodes

	// Define a recursive function to visit nodes
	var visit func(T)
	visit = func(node T) {
		// Get the unique key for the current node
		nodeKey := sortables.Unique(node)
		visited.Add(nodeKey) // Mark the node as visited

		// Visit all unvisited neighbors of the current node
		for _, neighbor := range g.Neighbors(node) {
			neighborKey := sortables.Unique(neighbor)
			if !visited.Has(neighborKey) {
				visit(neighbor) // Recursively visit the neighbor
			}
		}
	}

	// Get all nodes in the graph
	nodes := g.Nodes()

	// Start the visitation process from the first node
	visit(nodes[0])

	// The graph is connected if all nodes were visited
	return visited.Len() == g.adj.Len()
}

// IsDisconnectedOf returns true if the undirected graph is disconnected.
// A graph is considered disconnected if there is no path between every pair of nodes.
func IsDisconnectedOf[T gtools.SortableOf](g *UndirectGraphOf[T]) bool {
	// Simply return the opposite of IsConnectedOf, as a graph is disconnected if it's not connected
	return !IsConnectedOf(g)
}
