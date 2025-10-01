// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"strings"
)

// PersistentSortableEdge returns a new SingleTyped instance with the given from and to nodes.
//
// The SingleTyped type represents a directed edge in a graph, where each edge has a
// unique from and to node.
func PersistentSortableEdge[G any](hash string, from G, to G, direct bool) edges.UniqueSortableSingleTyped[G] {
	// Create a new SingleTyped instance with the given from and to nodes.
	return &persistentSortableEdge[G]{
		hash:   hash,
		from:   from, // The node that the edge originates from.
		to:     to,   // The node that the edge points to.
		direct: direct,
	}
}

// SingleTyped represents a directed edge in a graph, where each edge has a unique from and to identifier.
type persistentSortableEdge[G any] struct {
	hash   string
	from   G
	to     G
	direct bool
}

func (e *persistentSortableEdge[G]) From() G {
	return e.from
}

func (e *persistentSortableEdge[G]) To() G {
	return e.to
}

// Less returns true if the edge is less than the given edge. This method
// is used just for sorting purposes.
func (e *persistentSortableEdge[G]) Less(o interface{}) bool {
	return strings.Compare(e.Unique(), o.(*persistentSortableEdge[G]).Unique()) < 0
}

func (e *persistentSortableEdge[G]) Equal(o any) bool {
	other, ok := o.(*persistentSortableEdge[G])
	if !ok {
		return false
	}
	return e.hash == other.hash
}

// Unique returns a string representation of the edge for use as a map key.
func (e *persistentSortableEdge[G]) Unique() string {
	return e.hash
}

func (e *persistentSortableEdge[G]) String() string {
	return hashEdge[G, G](e, e.direct)
}

var _ edges.Edge[string, string] = (*persistentSortableEdge[string])(nil)
var _ edges.SortableSingleTyped[string] = (*persistentSortableEdge[string])(nil)
var _ edges.UniqueSortableSingleTyped[int] = (*persistentSortableEdge[int])(nil)

// SortableEdge returns a new SingleTyped instance with the given from and to nodes.
//
// The SingleTyped type represents a directed edge in a graph, where each edge has a
// unique from and to node.
func SortableEdge[G any](from G, to G, direct bool) edges.UniqueSortableSingleTyped[G] {
	// Create a new SingleTyped instance with the given from and to nodes.
	return &sortableEdge[G]{
		from:   from, // The node that the edge originates from.
		to:     to,   // The node that the edge points to.
		direct: direct,
	}
}

// SingleTyped represents a directed edge in a graph, where each edge has a unique from and to identifier.
type sortableEdge[G any] struct {
	hash   string
	from   G
	to     G
	direct bool
}

func (e *sortableEdge[G]) From() G {
	return e.from
}

func (e *sortableEdge[G]) To() G {
	return e.to
}

// Less returns true if the edge is less than the given edge. This method
// is used just for sorting purposes.
func (e *sortableEdge[G]) Less(o interface{}) bool {
	return strings.Compare(e.Unique(), o.(*sortableEdge[G]).Unique()) < 0
}

func (e *sortableEdge[G]) Equal(o any) bool {
	other, ok := o.(*sortableEdge[G])
	if !ok {
		return false
	}

	if e.direct != other.direct {
		return false
	}

	// For undirected edges, check if the endpoints are compare in either order
	if !e.direct {
		return undirectedEdgesEquality[G, G](e, other)
	}

	// For directed edges, check if they are exactly compare
	return directedEdgesEquality[G, G](e, other)
}

// Unique returns a string representation of the edge for use as a map key.
// The string representation is in the format "from->to", where "from" and "to" are the unique string representations of the edge's from and to nodes.
func (e *sortableEdge[G]) Unique() string {
	return hashEdge[G, G](e, e.direct)
}

func (e *sortableEdge[G]) String() string {
	return e.Unique()
}

var _ edges.Edge[string, string] = (*sortableEdge[string])(nil)
var _ edges.SortableSingleTyped[string] = (*sortableEdge[string])(nil)
var _ edges.UniqueSortableSingleTyped[int] = (*sortableEdge[int])(nil)

func PersistentSingleTypedWeightedEdge[G any, W any](hash string, from G, to G, weight W, direct bool) edges.UniqueSortableSingleTypedWeighted[G, W] {
	return &persistentSingleTypedWeightedEdge[G, W]{
		hash:   hash,
		from:   from,
		to:     to,
		weight: weight,
		direct: direct,
	}
}

type persistentSingleTypedWeightedEdge[G any, W any] struct {
	hash   string
	from   G
	to     G
	weight W
	direct bool
}

func (e *persistentSingleTypedWeightedEdge[G, W]) From() G {
	return e.from
}

func (e *persistentSingleTypedWeightedEdge[G, W]) To() G {
	return e.to
}

func (e *persistentSingleTypedWeightedEdge[G, W]) Weight() W {
	return e.weight
}

// Less returns true if the edge is less than the given edge. This method
// is used just for sorting purposes.
func (e *persistentSingleTypedWeightedEdge[G, W]) Less(o interface{}) bool {
	return sortables.TryLess(e.Weight(), o.(*persistentSingleTypedWeightedEdge[G, W]).Weight())
}

func (e *persistentSingleTypedWeightedEdge[G, W]) Equal(o interface{}) bool {
	other, ok := o.(*persistentSingleTypedWeightedEdge[G, W])
	if !ok {
		return false
	}
	return e.hash == other.hash
}

func (e *persistentSingleTypedWeightedEdge[G, W]) Unique() string {
	return e.hash
}

func (e *persistentSingleTypedWeightedEdge[G, W]) String() string {
	return hashEdge[G, G](e, e.direct)
}

var _ edges.Weighted[string, string, int] = (*persistentSingleTypedWeightedEdge[string, int])(nil)
var _ edges.UniqueSortableSingleTypedWeighted[float64, string] = (*persistentSingleTypedWeightedEdge[float64, string])(nil)

func SingleTypedWeightedEdge[G any, W any](from G, to G, weight W, direct bool) edges.UniqueSortableSingleTypedWeighted[G, W] {
	return &singleTypedWeightedEdge[G, W]{
		from:   from,
		to:     to,
		weight: weight,
		direct: direct,
	}
}

type singleTypedWeightedEdge[G any, W any] struct {
	hash   string
	from   G
	to     G
	weight W
	direct bool
}

func (e *singleTypedWeightedEdge[G, W]) From() G {
	return e.from
}

func (e *singleTypedWeightedEdge[G, W]) To() G {
	return e.to
}

func (e *singleTypedWeightedEdge[G, W]) Weight() W {
	return e.weight
}

// Less returns true if the edge is less than the given edge. This method
// is used just for sorting purposes.
func (e *singleTypedWeightedEdge[G, W]) Less(o interface{}) bool {
	return strings.Compare(e.Unique(), o.(*singleTypedWeightedEdge[G, W]).Unique()) < 0
}

func (e *singleTypedWeightedEdge[G, W]) Equal(o interface{}) bool {
	other, ok := o.(*singleTypedWeightedEdge[G, W])
	if !ok {
		return false
	}

	if e.direct != other.direct {
		return false
	}

	// For undirected edges, check if the endpoints are compare in either order
	if !e.direct {
		return undirectedEdgesEquality[G, G](e, other)
	}

	// For directed edges, check if they are exactly compare
	return directedEdgesEquality[G, G](e, other)
}

func (e *singleTypedWeightedEdge[G, W]) Unique() string {
	return hashEdge[G, G](e, e.direct)
}

func (e *singleTypedWeightedEdge[G, W]) String() string {
	return e.Unique()
}

var _ edges.Weighted[string, string, int] = (*singleTypedWeightedEdge[string, int])(nil)
var _ edges.UniqueSortableSingleTypedWeighted[float64, string] = (*singleTypedWeightedEdge[float64, string])(nil)
