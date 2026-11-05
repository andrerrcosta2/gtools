// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package edges

import (
	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/sortables"
)

func undirectedEdgesEquality[A any](a, b edge.Typed[A]) bool {
	fromA, toA := a.From(), a.To()
	fromB, toB := b.From(), b.To()

	// First check the "normal" order of comparison
	if sortables.TryEquality(fromA, fromB) && sortables.TryEquality(toA, toB) {
		return true
	}

	// Check if the edge direction is reversed (for undirected graphs)
	toBA, ok := any(toB).(A)
	if !ok {
		return false
	}
	if fromBA, ok := any(fromB).(A); ok {
		return sortables.TryEquality(fromA, toBA) && sortables.TryEquality(toA, fromBA)
	}

	return false
}

func directedEdgesEquality[A any](a, b edge.Typed[A]) bool {
	fromA, toA := a.From(), a.To()
	fromB, toB := b.From(), b.To()
	return sortables.TryEquality(fromA, fromB) && sortables.TryEquality(toA, toB)
}

// hashEdge creates a unique string representation of a directed or undirected edge in a graph.
//
// The function takes three parameters: 'a', which represents the edge to be hashed;
// 'direct', which is a boolean indicating whether the edge is directed or undirected;
// and 'A' and 'B', which are the types of the node that the edge connects.
//
// The function first gets the unique string representation of the 'from' and 'to' node
// of the edge.  If the edge is directed, the function returns a string in the format
// "from->to".  If the edge is undirected, the function returns a string in the format
// "from<->to", where the node are sorted alphabetically.
func hashEdge[A any, B any](a edge.Typed[A], direct bool) string {
	// Get the unique string representation of the 'from' node
	fromUnique := sortables.Fingerprint[A](a.From())

	// Get the unique string representation of the 'to' node
	toUnique := sortables.Fingerprint[B](a.To())

	// Represent the directed edge as "from->to"
	if direct {
		return fromUnique + "->" + toUnique
	}

	// Represent the undirected edge as "from<->to"
	first := min(fromUnique, toUnique)
	second := max(fromUnique, toUnique)

	return first + "<->" + second
}
