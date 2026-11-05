// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package edges

import (
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
)

// Comparable returns a new SingleTyped instance with the given from and to node.
//
// The SingleTyped type represents a directed edge in a graph, where each edge has a
// unique from and to node.
func Comparable[N any](from N, to N, direct bool) edge.Comparable[N] {
	// Create a new SingleTyped instance with the given from and to node.
	return &cmpEdge[N]{
		from:   from, // The node that the edge originates from.
		to:     to,   // The node that the edge points to.
		direct: direct,
	}
}

// SingleTyped represents a directed edge in a graph, where each edge has a unique from and to identifier.
type cmpEdge[N any] struct {
	hash   string
	from   N
	to     N
	direct bool
}

func (e *cmpEdge[N]) Direct() bool {
	return e.direct
}

func (e *cmpEdge[N]) From() N {
	return e.from
}

func (e *cmpEdge[N]) To() N {
	return e.to
}

func (e *cmpEdge[N]) Equal(other edge.Comparable[N]) bool {
	if e.direct != other.Direct() {
		return false
	}
	if !e.direct {
		return undirectedEdgesEquality(e, other)
	}
	return directedEdgesEquality(e, other)
}

// Unique returns a string representation of the edge for use as a map key.
// The string representation is in the format "from->to", where "from" and "to" are the unique string representations of the edge's from and to node.
func (e *cmpEdge[N]) Unique() string {
	return hashEdge[N, N](e, e.direct)
}

func (e *cmpEdge[N]) String() string {
	return e.Unique()
}

func BuildComparableFrom[N any](from N, neighbors set.Readable[N], direct bool) []edge.Comparable[N] {
	out := make([]edge.Comparable[N], neighbors.Len())
	neighbors.Range(func(to N) bool {
		out = append(out, Comparable(from, to, direct))
		return true
	})
	return out
}

func ComparableWeighted[N any, W any](from, to N, weight W, direct bool) edge.ComparableWeighted[N, W] {
	return &typedWeightedEdge[N, W]{
		from:   from,
		to:     to,
		weight: weight,
		direct: direct,
	}
}

type typedWeightedEdge[N any, W any] struct {
	hash   string
	from   N
	to     N
	weight W
	direct bool
}

func (e *typedWeightedEdge[N, W]) Direct() bool {
	return e.direct
}

func (e *typedWeightedEdge[N, W]) From() N {
	return e.from
}

func (e *typedWeightedEdge[N, W]) To() N {
	return e.to
}

func (e *typedWeightedEdge[N, W]) Weight() W {
	return e.weight
}

func (e *typedWeightedEdge[N, W]) Equal(other edge.ComparableWeighted[N, W]) bool {
	if e.direct != other.Direct() {
		return false
	}
	if !e.direct {
		return undirectedEdgesEquality(e, other)
	}
	return directedEdgesEquality(e, other)
}

func (e *typedWeightedEdge[N, W]) Unique() string {
	return hashEdge[N, N](e, e.direct)
}

func (e *typedWeightedEdge[N, W]) String() string {
	return e.Unique()
}

var _ edge.Weighted[string, int] = (*typedWeightedEdge[string, int])(nil)

func BuildComparableWeightedFrom[N any, W any](from N, neighbors dict.Readable[N, W], direct bool) []edge.ComparableWeighted[N, W] {
	out := make([]edge.ComparableWeighted[N, W], neighbors.Len())
	it := neighbors.Iterator()
	for neighbor, weight, hasNext := it.Next(); hasNext; neighbor, weight, hasNext = it.Next() {
		out = append(out, ComparableWeighted(from, neighbor, weight, direct))
	}
	return out
}
