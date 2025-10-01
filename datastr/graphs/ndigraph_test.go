// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
	"github.com/andrerrcosta2/gtools/datastr/maps"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"testing"
)

// TestAddNode tests adding random nodes to the undirected graph then check if they are in the graph
func TestAddNode(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode

	g := UndirectedOf[N]()

	exp := random.Struct[N](10).Each(g.AddNode)
	shouldContainAllNodesBy[N](tt, g, exp.Values(), sortables.Equality[N])

	tt.PrintLogStack()
}

// TestAddEdge tests adding random edges to the undirected graph then check if they are in the graph
func TestAddEdge(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Types
	type N = *tests.SortableNode
	type E = edges.SortableSingleTyped[*tests.SortableNode]

	g := UndirectedOf[N]()

	var expEdges iterables.Slice[E]

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					tt.Errorf("Error adding edge: %v", err)
				}
				expEdges.Append(SortableEdge(it, that, false))
			}
		})

	shouldHaveNEdges[N, E, N, N](tt, g, expEdges.Len())
	shouldContainAllEdgesBy[N, E, N, N](tt, g, expEdges.Values(), sortables.Equality[E])

	tt.PrintLogStack()
}

// TestNeighbors tests retrieving all neighbors from the graph and checking if they are
// correctly stored by its keys
func TestNeighbors(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode

	g := UndirectedOf[N]()

	var neighbors = maps.SortableOf[N, []N]()

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {

				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					tt.Errorf("Error adding edge: %v", err)
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

	shouldContainAllNeighborsBy[N](tt, g, neighbors, sortables.Equality[N])

	tt.PrintLogStack()
}

// TestEdges tests retrieving all edges from the graph
func TestEdges(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode
	type E = edges.SortableSingleTyped[*tests.SortableNode]

	g := UndirectedOf[N]()

	var expEdges iterables.Slice[E]

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					tt.Errorf("Error adding edge: %v", err)
				}
				expEdges.Append(SortableEdge(it, that, false))
			}
		})

	shouldContainAllEdgesBy[N, E, N, N](tt, g, expEdges.Values(), sortables.Equality[E])

	tt.PrintLogStack()
}

func TestAddEdgeNodesNotExist(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode
	type E = edges.SortableSingleTyped[*tests.SortableNode]

	g := UndirectedOf[N]()

	nodeA := tests.NewSortableNode("A")
	nodeB := tests.NewSortableNode("C")

	// Addf only one node
	g.AddNode(nodeA)

	// Attempt to add an edge with a non-existent node
	err := g.AddEdge(nodeA, nodeB)
	if err == nil {
		tt.Errorf("Should not be able to add non-existent edge: %v", err)
	}

	// Assert that the edge was not added
	shouldNotContainEdge[N, E, N, N](tt, g, SortableEdge(nodeA, nodeB, false))

	tt.PrintLogStack()
}
