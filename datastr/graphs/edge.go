// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/sortables"
)

func undirectedEdgesEquality[A any, B any](a, b edges.Edge[A, B]) bool {
	fromA, toA := a.From(), a.To()
	fromB, toB := b.From(), b.To()

	// First check the "normal" order of comparison
	if sortables.TryEquality[A](fromA, fromB) && sortables.TryEquality[A](toA, toB) {
		return true
	}

	// Check if the edge direction is reversed (for undirected graphs)
	toBA, ok := any(toB).(A)
	if !ok {
		return false
	}
	if fromBA, ok := any(fromB).(B); ok {
		return sortables.TryEquality[A](fromA, toBA) && sortables.TryEquality[A](toA, fromBA)
	}

	return false
}

func directedEdgesEquality[A any, B any](a, b edges.Edge[A, B]) bool {
	fromA, toA := a.From(), a.To()
	fromB, toB := b.From(), b.To()
	return sortables.TryEquality[A](fromA, fromB) && sortables.TryEquality[A](toA, toB)
}

// hashEdge creates a unique string representation of a directed or undirected edge in a graph.
//
// The function takes three parameters: 'a', which represents the edge to be hashed;
// 'direct', which is a boolean indicating whether the edge is directed or undirected;
// and 'A' and 'B', which are the types of the nodes that the edge connects.
//
// The function first gets the unique string representation of the 'from' and 'to' nodes
// of the edge.  If the edge is directed, the function returns a string in the format
// "from->to".  If the edge is undirected, the function returns a string in the format
// "from<->to", where the nodes are sorted alphabetically.
func hashEdge[A any, B any](a edges.Edge[A, B], direct bool) string {
	// Get the unique string representation of the 'from' node
	fromUnique := sortables.Unique[A](a.From())

	// Get the unique string representation of the 'to' node
	toUnique := sortables.Unique[B](a.To())

	// Represent the directed edge as "from->to"
	if direct {
		return fromUnique + "->" + toUnique
	}

	// Represent the undirected edge as "from<->to"
	first := min(fromUnique, toUnique)
	second := max(fromUnique, toUnique)

	return first + "<->" + second
}
