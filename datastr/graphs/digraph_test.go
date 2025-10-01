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

// TestDirectedGraph_AddNode tests adding nodes to the graph.
func TestDirectedGraph_AddNode(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	type N = *tests.SortableNode

	// Create graph
	g := DigraphOf[N]()

	nodes := random.Struct[N](10).Each(g.AddNode)

	shouldContainAllNodesBy[N](tt, g, nodes.Values(), sortables.Equality[N])

	tt.PrintLogStack()
}

// TestDirectedGraph_AddEdge tests adding directed edges to the graph.
func TestDirectedGraph_AddEdge(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Types
	type N = *tests.SortableNode
	type E = edges.SortableSingleTyped[*tests.SortableNode]

	// Create graph
	g := DigraphOf[N]()

	var expEdges = iterables.OfSlice[E]()

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					return
				}
				expEdges.Append(SortableEdge(it, that, true))
			}
		})

	shouldContainAllEdgesBy[N, E, N, N](tt, g, expEdges.Values(), sortables.Equality[E])

	tt.PrintLogStack()
}

// TestDirectedGraph_Neighbors tests retrieving neighbors for a node.
func TestDirectedGraph_Neighbors(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Types
	type N = *tests.SortableNode

	g := DigraphOf[N]()

	var neighbors = maps.SortableOf[N, []N]()

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				g.AddEdge(it, that)
				if nodes, ok := neighbors.Get(it); !ok {
					neighbors.Put(it, []N{that})
				} else {
					neighbors.Put(it, append(nodes, that))
				}
			}
		})

	shouldContainAllNeighborsBy[N](tt, g, neighbors, sortables.Equality[N])

	tt.PrintLogStack()
}

// TestDirectedGraph_HasNode tests the presence of nodes in the graph.
func TestDirectedGraph_HasNode(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Types
	type N = *tests.SortableNode

	g := DigraphOf[N]()
	exp := random.Struct[N](15).Each(g.AddNode).Values()
	// Assert the presence of nodes
	shouldContainAllNodesBy[N](tt, g, exp, sortables.Equality[N])

	tt.PrintLogStack()
}

// TestDirectedGraph_HasEdge tests the presence of edges in the graph.
func TestDirectedGraph_HasEdge(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Types
	type N = *tests.SortableNode
	type E = edges.SortableSingleTyped[*tests.SortableNode]

	g := DigraphOf[N]()

	var expEdges iterables.Slice[E]

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				g.AddEdge(it, that)
				expEdges.Append(SortableEdge(it, that, true))
			}
		})

	shouldContainAllEdgesBy[N, E, N, N](tt, g, expEdges.Values(), sortables.Equality[E])

	tt.PrintLogStack()
}

// TestDirectedGraph_Nodes tests retrieving all nodes from the graph.
func TestDirectedGraph_Nodes(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Types
	type N = *tests.SortableNode

	g := DigraphOf[N]()

	exp := random.Struct[N](100).Each(g.AddNode)

	tt.StackLogf("Nodes: %v\n", g.Nodes())
	tt.StackLogf("Expected nodes: %v\n", exp.Values())

	shouldContainAllNodesBy[N](tt, g, exp.Values(), sortables.Equality[N])
	shouldContainNNodes[N](tt, g, exp.Len())

	tt.PrintLogStack()
}

// TestDirectedGraph_Edges tests retrieving all directed edges from the graph.
func TestDirectedGraph_Edges(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Types
	type N = *tests.SortableNode
	type E = edges.SortableSingleTyped[*tests.SortableNode]

	g := DigraphOf[N]()

	var expEdges = iterables.OfSlice[E]()

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					tt.Errorf("Error adding edge: %v", err)
				}
				expEdges.Append(SortableEdge(it, that, true))
			}
		})

	shouldHaveNEdges[N, E, N, N](tt, g, expEdges.Len())
	shouldContainAllEdgesBy[N, E, N, N](tt, g, expEdges.Values(), sortables.Equality[E])

	tt.PrintLogStack()
}

func TestDirectedGraph_AddEdgeNodesNotExist(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Types
	type N = *tests.SortableNode

	g := DigraphOf[N]()

	nodeA := tests.NewSortableNode("A")
	nodeB := tests.NewSortableNode("C")

	// Add only one node
	g.AddNode(nodeA)

	// Attempt to add an edge with a non-existent node
	g.AddEdge(nodeA, nodeB)

	// Assert that the edge was not added
	if g.HasEdge(nodeA, nodeB) {
		tt.Errorf("Edge from A to C should not be present")
	}

	tt.PrintLogStack()
}
