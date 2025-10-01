// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
	"github.com/andrerrcosta2/gtools/datastr/maps"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/numbers/prog"
	"testing"
)

// TestAddingNodesAndEdges Test the basics of adding nodes and edges
// and expecting them to be present
func TestAddingNodesAndEdges(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edges.SortableSingleTypedWeighted[N, W]
	type G = str.WOrderedGraphOf[N, W, E]
	// Graph
	g := WeightedSortableDigraphOf[N, W]()

	var expEdges []E
	weights, err := prog.Sqrt(1_073_741_824, 2, 30)
	if err != nil {
		tt.Fatal(err)
	}

	random.Struct[N](30).Each(g.AddNode).
		Operation(func(i int, n *iterables.Slice[N]) {
			// LeftChild child
			if 2*i+1 < n.Len() {
				err := g.AddEdge(n.At(i), n.At(2*i+1), weights[i])
				if err != nil {
					tt.Fatal(err)
				}
				expEdges = append(expEdges, SingleTypedWeightedEdge(n.At(i), n.At(2*i+1), weights[i], true))
			}
			// RightChild child
			if 2*i+2 < n.Len() {
				err := g.AddEdge(n.At(i), n.At(2*i+2), weights[i])
				if err != nil {
					tt.Fatal(err)
				}
				expEdges = append(expEdges, SingleTypedWeightedEdge(n.At(i), n.At(2*i+1), weights[i], true))
			}
		})

	shouldContainAllEdgesBy[N, E, N, N](tt, g, expEdges, sortables.Equality[E])

	tt.PrintLogStack()
}

// TestNodeExistence Test the existence of nodes after being inserted
func TestNodeExistence(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edges.SortableSingleTypedWeighted[N, W]
	type G = str.WOrderedGraphOf[N, W, E]
	// Graph
	g := WeightedSortableDigraphOf[N, W]()

	a := tests.NewSortableNode("A")
	b := tests.NewSortableNode("C")
	g.AddNode(a)
	g.AddNode(b)

	exp := []N{a, b}

	shouldContainAllNodesBy[N](tt, g, exp, sortables.Equality[N])

	tt.PrintLogStack()
}

// TestEdgeExistenceAndWeight Test the existence of edges after being inserted
// and their weights
func TestEdgeExistenceAndWeight(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edges.SortableSingleTypedWeighted[N, W]
	type G = str.WOrderedGraphOf[N, W, E]
	// Graph
	g := WeightedSortableDigraphOf[N, W]()

	var expectedEdges = make([]E, 2)

	a := tests.NewSortableNode("A")
	b := tests.NewSortableNode("C")
	c := tests.NewSortableNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddEdge(a, b, 10)
	expectedEdges[0] = SingleTypedWeightedEdge(a, b, 10, true)
	g.AddEdge(b, c, 20)
	expectedEdges[1] = SingleTypedWeightedEdge(b, c, 20, true)

	shouldContainAllEdgesBy[N, E, N, N](tt, g, expectedEdges, sortables.Equality[E])

	tt.PrintLogStack()
}

// TestNeighborsRetrieval Test the retrieval of neighbors
func TestNeighborsRetrieval(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edges.SortableSingleTypedWeighted[N, W]
	type G = str.WOrderedGraphOf[N, W, E]
	// Graph
	g := WeightedSortableDigraphOf[N, W]()

	expNeighbors := maps.SortableOf[N, []N]()

	a := tests.NewSortableNode("A")
	b := tests.NewSortableNode("C")
	c := tests.NewSortableNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddEdge(a, b, 10)
	expNeighbors.Put(a, []N{b})
	g.AddEdge(b, c, 20)
	expNeighbors.Put(b, []N{c})

	shouldContainAllNeighborsBy[N](tt, g, expNeighbors, sortables.Equality[N])

	tt.PrintLogStack()
}

// TestDisconnectedGraph Test a disconnected graph
func TestDisconnectedGraph(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edges.SortableSingleTypedWeighted[N, W]
	type G = str.WOrderedGraphOf[N, W, E]
	// Graph
	g := WeightedSortableDigraphOf[N, W]()

	a := tests.NewSortableNode("A")
	b := tests.NewSortableNode("C")
	c := tests.NewSortableNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddEdge(a, b, 10)

	if !g.HasEdge(a, b) {
		tt.Errorf("Graph should have an edge from A to C.")
	}

	if g.HasEdge(b, c) {
		tt.Errorf("Graph should not have an edge from C to C.")
	}

	tt.PrintLogStack()
}
