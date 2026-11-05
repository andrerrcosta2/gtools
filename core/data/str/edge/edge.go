// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package edge

import "github.com/andrerrcosta2/gtools/core/domain/constraints"

// Typed represents a directed connection (an "arrow") between two nodes.
//
// In graph theory, an edge is not a container of nodes, but a relation
// connecting a source node to a destination node.
//
//	A ──→ B
type Typed[N any] interface {
	From() N
	To() N
	// Direct reports whether the edge is direct or bidirectional
	Direct() bool
}

type Comparable[N any] interface {
	Typed[N]
	constraints.Comparable[Comparable[N]]
}

type Hyper[N any] interface {
	From() N
	To() []N
}

type ComparableHyper[N any] interface {
	Hyper[N]
	constraints.Comparable[ComparableHyper[N]]
}

type Weighted[N, W any] interface {
	Typed[N]
	Weight() W
}

type ComparableWeighted[N, W any] interface {
	Weighted[N, W]
	constraints.Comparable[ComparableWeighted[N, W]]
}

type HyperWeighted[N, W any] interface {
	Hyper[N]
	Weight() W
}

type ComparableHyperWeighted[N, W any] interface {
	HyperWeighted[N, W]
	constraints.Comparable[ComparableHyperWeighted[N, W]]
}
