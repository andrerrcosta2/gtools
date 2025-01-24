// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

//import (
//	"fmt"
//	"github.com/andrerrcosta2/gtools/core/gtools"
//)
//
//type Type int
//
//const (
//	Directed Type = iota
//	Undirected
//	Weighted
//)
//
//func New[G any, W any](graph Type, nodes ...G) (Graph[G], error) {
//	switch graph {
//	case Undirected:
//		return tryCreateUndirectedGraph[G](nodes)
//	case Directed:
//		return tryCreateDirectedGraph[G](nodes)
//	case Weighted:
//		return tryCreateWeightedGraph[G, W](nodes)
//	default:
//		return nil, fmt.Errorf("graph.New: invalid graph type: %d", graph)
//	}
//}
//
//func tryCreateUndirectedGraph[N any](nodes []N) (Graph[N], error) {
//	switch n := any(nodes).(type) {
//	case []gtools.SortableOf:
//		return any(UndirectedOf[gtools.SortableOf](n...)).(Graph[N]), nil
//	case []int:
//		return any(Undirected[int](n...)).(Graph[N]), nil
//	case []int8:
//		return any(Undirected[int8](n...)).(Graph[N]), nil
//	case []int16:
//		return any(Undirected[int16](n...)).(Graph[N]), nil
//	case []int32:
//		return any(Undirected[int32](n...)).(Graph[N]), nil
//	case []int64:
//		return any(Undirected[int64](n...)).(Graph[N]), nil
//	case []float32:
//		return any(Undirected[float32](n...)).(Graph[N]), nil
//	case []float64:
//		return any(Undirected[float64](n...)).(Graph[N]), nil
//	case []uint:
//		return any(Undirected[uint](n...)).(Graph[N]), nil
//	case []uint8:
//		return any(Undirected[uint8](n...)).(Graph[N]), nil
//	case []uint16:
//		return any(Undirected[uint16](n...)).(Graph[N]), nil
//	case []uint32:
//		return any(Undirected[uint32](n...)).(Graph[N]), nil
//	case []uint64:
//		return any(Undirected[uint64](n...)).(Graph[N]), nil
//	case []string:
//		return any(Undirected[string](n...)).(Graph[N]), nil
//	default:
//		return nil, fmt.Errorf("the provided nodes do not implement comparable: %T\n", n)
//	}
//}
//
//func tryCreateDirectedGraph[N any](nodes []N) (Graph[N], error) {
//	switch n := any(nodes).(type) {
//	case []gtools.SortableOf:
//		return any(DigraphOf[gtools.SortableOf](n...)).(Graph[N]), nil
//	case []int:
//		return any(Digraph[int](n...)).(Graph[N]), nil
//	case []int8:
//		return any(Digraph[int8](n...)).(Graph[N]), nil
//	case []int16:
//		return any(Digraph[int16](n...)).(Graph[N]), nil
//	case []int32:
//		return any(Digraph[int32](n...)).(Graph[N]), nil
//	case []int64:
//		return any(Digraph[int64](n...)).(Graph[N]), nil
//	case []float32:
//		return any(Digraph[float32](n...)).(Graph[N]), nil
//	case []float64:
//		return any(Digraph[float64](n...)).(Graph[N]), nil
//	case []uint:
//		return any(Digraph[uint](n...)).(Graph[N]), nil
//	case []uint8:
//		return any(Digraph[uint8](n...)).(Graph[N]), nil
//	case []uint16:
//		return any(Digraph[uint16](n...)).(Graph[N]), nil
//	case []uint32:
//		return any(Digraph[uint32](n...)).(Graph[N]), nil
//	case []uint64:
//		return any(Digraph[uint64](n...)).(Graph[N]), nil
//	case []string:
//		return any(Digraph[string](n...)).(Graph[N]), nil
//	default:
//		return nil, fmt.Errorf("the provided nodes are not orderable: %T\n", n)
//	}
//}
//
//func tryCreateWeightedGraph[N any, W any](nodes []N) (Graph[N], error) {
//	switch n := any(nodes).(type) {
//	case []gtools.SortableOf:
//		return any(WeightedSortableDigraphOf[gtools.SortableOf, W](n...)).(Graph[N]), nil
//	case []int:
//		return any(Digraph[int](n...)).(Graph[N]), nil
//	case []int8:
//		return any(Digraph[int8](n...)).(Graph[N]), nil
//	case []int16:
//		return any(Digraph[int16](n...)).(Graph[N]), nil
//	case []int32:
//		return any(Digraph[int32](n...)).(Graph[N]), nil
//	case []int64:
//		return any(Digraph[int64](n...)).(Graph[N]), nil
//	case []float32:
//		return any(Digraph[float32](n...)).(Graph[N]), nil
//	case []float64:
//		return any(Digraph[float64](n...)).(Graph[N]), nil
//	case []uint:
//		return any(Digraph[uint](n...)).(Graph[N]), nil
//	case []uint8:
//		return any(Digraph[uint8](n...)).(Graph[N]), nil
//	case []uint16:
//		return any(Digraph[uint16](n...)).(Graph[N]), nil
//	case []uint32:
//		return any(Digraph[uint32](n...)).(Graph[N]), nil
//	case []uint64:
//		return any(Digraph[uint64](n...)).(Graph[N]), nil
//	case []string:
//		return any(Digraph[string](n...)).(Graph[N]), nil
//	default:
//		return nil, fmt.Errorf("the provided nodes do not implement comparable: %T\n", n)
//	}
//}
