// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

import "github.com/andrerrcosta2/gtools/core/data"

type Space[S any] interface {
	Expand(S) data.Iterator[S]
}

type WeightedSpace[S any] interface {
	Space[S]
	Cost(from, to S) float64
}

type HeuristicSpace[S any] interface {
	WeightedSpace[S]
	Heuristic(from, goal S) float64
}
