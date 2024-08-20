// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/sortables"
)

// DepthFirst performs a depth-first search on the graph starting from the given node.
// It returns a slice of visited nodes in the order they were visited.
func DepthFirst[T any, K constraints.Ordered](g iterables.Deliverer[T], nodeKey K) []T {
	// Initialize the result slice to store the visited nodes.
	var result []T

	// Create a set to track visited nodes and avoid revisiting them.
	visited := map[K]struct{}{}

	// Start the depth-first search from the given start node.
	dfs[T, K](nodeKey, visited, g, result)

	// Return the slice of visited nodes in the order they were visited.
	return result
}

// dfs performs a recursive depth-first search on the graph starting from the given node.
// It marks the node as visited, adds it to the result, and recursively visits all its neighbors.
func dfs[T any, K constraints.Ordered](currentKey K, visited map[K]struct{}, g iterables.Deliverer[T], result []T) {
	// Check if the node has already been visited to avoid infinite loops.
	if visited[currentKey] == struct{}{} {
		// If the node is already visited, return immediately to prevent revisiting.
		return
	}

	// Mark the node as visited to keep track of visited nodes.
	visited[currentKey] = struct{}{}

	// Convert the key back to the node type (T) to add it to the result.
	currentNode := any(currentKey).(T)
	// Add the current node to the result slice.
	result = append(result, currentNode)

	// Recursively visit all neighbors of the current node.
	for _, neighbor := range g.Deliver(currentNode) {
		// Convert the neighbor to the appropriate key type (K) for the recursive call.
		neighborKey := xnk(neighbor).(K)
		// Recursively call dfs on the neighbor node.
		dfs[T, K](neighborKey, visited, g, result)
	}
}

// xnk is a helper function to convert a node to its key type.
// It takes a node of any type and returns its corresponding key type.
func xnk(n any) any {
	// Use type switching to determine the type of the node
	switch n := n.(type) {
	// If the node is of type gtools.SortableOf, return its unique key
	case gtools.SortableOf:
		return sortables.Unique(n)
	// If the node is of any other type, return the node itself
	default:
		return n
	}
}
