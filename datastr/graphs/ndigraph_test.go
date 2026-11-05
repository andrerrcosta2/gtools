// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/datastr/dicts"
	"github.com/andrerrcosta2/gtools/datastr/graphs/edges"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
)

// TestAddNode tests adding random node to the undirected graph then check if they are in the graph
func TestAddNode(t *testing.T) {
	// Types
	type N = *tests.SortableNode

	g := UndirectedOf[N]()

	exp := random.Struct[N](10).Each(g.AddNode)
	shouldContainAllNodesBy[N](t, g, exp.Values(), sortables.Equality[N])

}

// TestAddEdge tests adding random edge to the undirected graph then check if they are in the graph
func TestAddEdge(t *testing.T) {

	// Types
	type N = *tests.SortableNode
	type E = edge.Comparable[N]

	g := UndirectedOf[N]()

	var expEdges iterables.Slice[E]

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					t.Errorf("Error adding edge: %v", err)
				}
				expEdges.Append(edges.Comparable(it, that, false))
			}
		})

	shouldHaveNEdges(t, g, expEdges.Len())
	shouldContainAllEdgesBy(t, g, expEdges, sortables.Equality[E])
}

// TestNeighbors tests retrieving all neighbors from the graph and checking if they are
// correctly stored by its keys
func TestNeighbors(t *testing.T) {
	// Types
	type N = *tests.SortableNode

	g := UndirectedOf[N]()

	var neighbors = dicts.ComparableOf[N, []N]()

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {

				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					t.Errorf("Error adding edge: %v", err)
				}

				// Undirected graphs have neighbors in both directions
				if arr, ok := neighbors.Get(it); !ok {
					neighbors.Put(it, []N{that})
				} else {
					neighbors.Put(it, append(arr, that))
				}

				if arr, ok := neighbors.Get(that); !ok {
					neighbors.Put(that, []N{it})
				} else {
					neighbors.Put(that, append(arr, it))
				}
			}
		})

	shouldContainAllNeighborsBy[N](t, g, neighbors, sortables.Equality[N])

}

// TestEdges tests retrieving all edge from the graph
func TestEdges(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type E = edge.Comparable[N]

	g := UndirectedOf[N]()

	var expEdges iterables.Slice[E]

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					t.Errorf("Error adding edge: %v", err)
				}
				expEdges.Append(edges.Comparable(it, that, false))
			}
		})

	shouldContainAllEdgesBy[E, N](t, g, expEdges.Values(), sortables.Equality[E])

}

func TestAddEdgeNodesNotExist(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type E = edge.Comparable[N]

	g := UndirectedOf[N]()

	nodeA := tests.NewSortableNode("A")
	nodeB := tests.NewSortableNode("C")

	//Add only one node
	g.AddNode(nodeA)

	// Attempt to add an edge with a non-existent node
	err := g.AddEdge(nodeA, nodeB)
	if err == nil {
		t.Errorf("Should not be able to add non-existent edge: '%v'", err)
	}

	// Assert that the edge was not added
	shouldNotContainEdge[E, N](t, g, edges.Comparable(nodeA, nodeB, false))
}
