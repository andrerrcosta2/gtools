// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	//"github.com/andrerrcosta2/gtools/pkg/sorts"
)

// MinSpanningTree returns the minimum spanning tree a weighted graph.
// It uses Kruskal's algorithm to find the minimum spanning tree.
func MinSpanningTree[N any, W any](g WGraph[N, W]) WGraph[N, W] {

}

func SortEdgesByWeight[N any, W any](g WGraph[N, W], comparator comparables.Comparator[W]) WGraph[N, W] {
	//sorter := sorts.NewQuicksort(comparator).Sort
	//edges := g.Edges()
	//sorter.Sort(&edges)
}
