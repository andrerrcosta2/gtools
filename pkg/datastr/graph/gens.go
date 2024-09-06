// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/nums/graph/graphs"
)

// grok_system_design_interview.pdf
type Type int

const (
	Directed Type = iota
	Undirected
	Weighted
)

func New[G any, W any](graph Type, nodes ...G) (Graph[G], error) {
	switch graph {
	case Undirected:
		return tryCreateUndirectedGraph[G](nodes)
	case Directed:
		return tryCreateDirectedGraph[G](nodes)
	case Weighted:
		return tryCreateWeightedGraph[G, W](nodes)
	default:
		return nil, fmt.Errorf("graph.New: invalid graph type: %d", graph)
	}
}

func tryCreateUndirectedGraph[N any](nodes []N) (Graph[N], error) {
	switch n := any(nodes).(type) {
	case []gtools.SortableOf:
		return any(graphs.UndirectOf[gtools.SortableOf](n...)).(Graph[N]), nil
	case []int:
		return any(graphs.Undirect[int](n...)).(Graph[N]), nil
	case []int8:
		return any(graphs.Undirect[int8](n...)).(Graph[N]), nil
	case []int16:
		return any(graphs.Undirect[int16](n...)).(Graph[N]), nil
	case []int32:
		return any(graphs.Undirect[int32](n...)).(Graph[N]), nil
	case []int64:
		return any(graphs.Undirect[int64](n...)).(Graph[N]), nil
	case []float32:
		return any(graphs.Undirect[float32](n...)).(Graph[N]), nil
	case []float64:
		return any(graphs.Undirect[float64](n...)).(Graph[N]), nil
	case []uint:
		return any(graphs.Undirect[uint](n...)).(Graph[N]), nil
	case []uint8:
		return any(graphs.Undirect[uint8](n...)).(Graph[N]), nil
	case []uint16:
		return any(graphs.Undirect[uint16](n...)).(Graph[N]), nil
	case []uint32:
		return any(graphs.Undirect[uint32](n...)).(Graph[N]), nil
	case []uint64:
		return any(graphs.Undirect[uint64](n...)).(Graph[N]), nil
	case []string:
		return any(graphs.Undirect[string](n...)).(Graph[N]), nil
	default:
		return nil, fmt.Errorf("the provided nodes do not implement comparable: %T\n", n)
	}
}

func tryCreateDirectedGraph[N any](nodes []N) (Graph[N], error) {
	switch n := any(nodes).(type) {
	case []gtools.SortableOf:
		return any(graphs.DigraphOf[gtools.SortableOf](n...)).(Graph[N]), nil
	case []int:
		return any(graphs.Digraph[int](n...)).(Graph[N]), nil
	case []int8:
		return any(graphs.Digraph[int8](n...)).(Graph[N]), nil
	case []int16:
		return any(graphs.Digraph[int16](n...)).(Graph[N]), nil
	case []int32:
		return any(graphs.Digraph[int32](n...)).(Graph[N]), nil
	case []int64:
		return any(graphs.Digraph[int64](n...)).(Graph[N]), nil
	case []float32:
		return any(graphs.Digraph[float32](n...)).(Graph[N]), nil
	case []float64:
		return any(graphs.Digraph[float64](n...)).(Graph[N]), nil
	case []uint:
		return any(graphs.Digraph[uint](n...)).(Graph[N]), nil
	case []uint8:
		return any(graphs.Digraph[uint8](n...)).(Graph[N]), nil
	case []uint16:
		return any(graphs.Digraph[uint16](n...)).(Graph[N]), nil
	case []uint32:
		return any(graphs.Digraph[uint32](n...)).(Graph[N]), nil
	case []uint64:
		return any(graphs.Digraph[uint64](n...)).(Graph[N]), nil
	case []string:
		return any(graphs.Digraph[string](n...)).(Graph[N]), nil
	default:
		return nil, fmt.Errorf("the provided nodes are not orderable: %T\n", n)
	}
}

func tryCreateWeightedGraph[N any, W any](nodes []N) (Graph[N], error) {
	switch n := any(nodes).(type) {
	case []gtools.SortableOf:
		return any(graphs.WeightedOrderedOf[gtools.SortableOf, constraints.Ordered](n...)).(Graph[N]), nil
	case []int:
		return any(graphs.Digraph[int](n...)).(Graph[N]), nil
	case []int8:
		return any(graphs.Digraph[int8](n...)).(Graph[N]), nil
	case []int16:
		return any(graphs.Digraph[int16](n...)).(Graph[N]), nil
	case []int32:
		return any(graphs.Digraph[int32](n...)).(Graph[N]), nil
	case []int64:
		return any(graphs.Digraph[int64](n...)).(Graph[N]), nil
	case []float32:
		return any(graphs.Digraph[float32](n...)).(Graph[N]), nil
	case []float64:
		return any(graphs.Digraph[float64](n...)).(Graph[N]), nil
	case []uint:
		return any(graphs.Digraph[uint](n...)).(Graph[N]), nil
	case []uint8:
		return any(graphs.Digraph[uint8](n...)).(Graph[N]), nil
	case []uint16:
		return any(graphs.Digraph[uint16](n...)).(Graph[N]), nil
	case []uint32:
		return any(graphs.Digraph[uint32](n...)).(Graph[N]), nil
	case []uint64:
		return any(graphs.Digraph[uint64](n...)).(Graph[N]), nil
	case []string:
		return any(graphs.Digraph[string](n...)).(Graph[N]), nil
	default:
		return nil, fmt.Errorf("the provided nodes do not implement comparable: %T\n", n)
	}
}
