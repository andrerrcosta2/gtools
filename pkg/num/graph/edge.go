// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/sortables"
)

type Edge[G any, W any] interface {
	From() G
	To() W
}

type WeightedEdge[G any, R any, W any] interface {
	From() G
	To() R
	Weight() W
}

// NewEdge returns a new SingleTypedEdge instance with the given from and to nodes.
//
// The SingleTypedEdge type represents a directed edge in a graph, where each edge has a
// unique from and to node.
func NewEdge[G any](from G, to G) *SingleTypedEdge[G] {
	// Create a new SingleTypedEdge instance with the given from and to nodes.
	return &SingleTypedEdge[G]{
		from: from, // The node that the edge originates from.
		to:   to,   // The node that the edge points to.
	}
}

// SingleTypedEdge represents a directed edge in a graph, where each edge has a unique from and to identifier.
type SingleTypedEdge[G any] struct {
	from   G
	to     G
	direct bool
}

func (e *SingleTypedEdge[G]) From() G {
	return e.from
}

func (e *SingleTypedEdge[G]) To() G {
	return e.to
}

func (e *SingleTypedEdge[G]) Less(o interface{}) bool {
	return e.Unique() < o.(*SingleTypedEdge[G]).Unique()
}

func (e *SingleTypedEdge[G]) Equal(o any) bool {
	other, ok := o.(*SingleTypedEdge[G])
	if !ok {
		return false
	}

	if e.direct != other.direct {
		return false
	}

	// For undirected edges, check if the endpoints are equal in either order
	if !e.direct && !other.direct {
		return (sortables.Unique(e.From()) == sortables.Unique(other.From()) &&
			sortables.Unique(e.To()) == sortables.Unique(other.To())) ||
			(sortables.Unique(e.From()) == sortables.Unique(other.From()) &&
				sortables.Unique(e.To()) == sortables.Unique(other.To())) ||
			(sortables.Unique(e.From()) == sortables.Unique(other.To()) &&
				sortables.Unique(e.To()) == sortables.Unique(other.From()))
	}

	// For directed edges, check if they are exactly equal
	return sortables.Unique(e.From()) == sortables.Unique(other.From()) &&
		sortables.Unique(e.To()) == sortables.Unique(other.To())
}

// Unique returns a string representation of the edge for use as a map key.
// The string representation is in the format "from->to", where "from" and "to" are the unique string representations of the edge's from and to nodes.
func (e *SingleTypedEdge[G]) Unique() string {
	// Get the unique string representation of the from node
	fromUnique := sortables.Unique(e.From())

	// Get the unique string representation of the to node
	toUnique := sortables.Unique(e.To())

	// Represent a directed edge as "from->to"
	if e.direct {
		return fromUnique + " -> " + toUnique
	}

	// Represent a undirected edge as "from<->to"
	return fromUnique + " <-> " + toUnique
}

func (e *SingleTypedEdge[G]) String() string {
	return e.Unique()
}

var _ Edge[string, string] = (*SingleTypedEdge[string])(nil)
var _ gtools.SortableOf = (*SingleTypedEdge[string])(nil)

func NewWeightedEdge[G any, W any](from G, to G, weight W) *SingleTypedWeightedEdge[G, W] {
	return &SingleTypedWeightedEdge[G, W]{
		from:   from,
		to:     to,
		weight: weight,
	}
}

type SingleTypedWeightedEdge[G any, W any] struct {
	from   G
	to     G
	weight W
}

func (e *SingleTypedWeightedEdge[G, W]) From() G {
	return e.from
}

func (e *SingleTypedWeightedEdge[G, W]) To() G {
	return e.to
}

func (e *SingleTypedWeightedEdge[G, W]) Weight() W {
	return e.weight
}

func (e *SingleTypedWeightedEdge[G, W]) Less(o interface{}) bool {
	return e.Unique() < o.(*SingleTypedWeightedEdge[G, W]).Unique()
}

func (e *SingleTypedWeightedEdge[G, W]) Equal(o interface{}) bool {
	other, ok := o.(*SingleTypedWeightedEdge[G, W])
	if !ok {
		return false
	}

	return sortables.Unique(e.From()) == sortables.Unique(other.From()) &&
		sortables.Unique(e.To()) == sortables.Unique(other.To())
}

func (e *SingleTypedWeightedEdge[G, W]) Unique() string {
	// Get the unique string representation of the from node
	fromUnique := sortables.Unique(e.From())

	// Get the unique string representation of the to node
	toUnique := sortables.Unique(e.To())

	// Combine the from and to unique strings with an arrow in between
	return fromUnique + "->" + toUnique
}

func (e *SingleTypedWeightedEdge[G, W]) String() string {
	return e.Unique()
}

var _ WeightedEdge[string, string, int] = (*SingleTypedWeightedEdge[string, int])(nil)
var _ gtools.SortableOf = (*SingleTypedWeightedEdge[string, int])(nil)
