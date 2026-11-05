// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package heuristic

import (
	"container/heap"

	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/graph"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/core/search/internal/queues"
)

type record[N comparable, W nums.Ordered] struct {
	parent N
	g      W
}

func AStarBy[E edge.Weighted[N, W], N comparable, W nums.Ordered](
	start N,
	goal N,
	neighbors func(N) []E,
	heuristic func(N, N) W,
) ([]N, W, bool) {

	open := make(queues.Priority[N, W], 0)
	heap.Init(&open)
	heap.Push(&open, queues.NewPqItem(start, 0))

	records := make(map[N]record[N, W])
	records[start] = record[N, W]{}

	for open.Len() > 0 {
		current := heap.Pop(&open).(*queues.PqItem[N, W]).Node()

		if current == goal {
			return reconstructPath(records, current), records[current].g, true
		}

		for _, edge := range neighbors(current) {
			realStep := records[current].g + edge.Weight()

			// if the real step is less than the old weight, that means
			// we found a better prefix path to the same node
			if r, ok := records[edge.To()]; !ok || realStep < r.g {
				records[edge.To()] = record[N, W]{current, realStep}
				f := realStep + heuristic(edge.To(), goal)
				heap.Push(&open, queues.NewPqItem(edge.To(), f))
			}
		}
	}

	return nil, 0, false
}

func reconstructPath[N comparable, W nums.Ordered](records map[N]record[N, W], from N) []N {
	path := []N{from}

	for {
		prev, ok := records[from]
		if !ok {
			break
		}
		from = prev.parent
		path = append([]N{from}, path...)
	}

	return path
}

func AStar[G graph.ReadableEdged[E, N], E edge.Weighted[N, W], N comparable, W nums.Ordered](
	g G,
	start N,
	goal N,
	heuristic func(N, N) W,
) ([]N, W, bool) {
	return AStarBy(start, goal, g.EdgesFrom, heuristic)
}
