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

// TestDirectedGraph_AddNode tests adding node to the graph.
func TestDirectedGraph_AddNode(t *testing.T) {
	// Helper

	type N = *tests.SortableNode

	// Create graph
	g := DigraphOf[N]()

	nodes := random.Struct[N](10).Each(g.AddNode)

	shouldContainAllNodesBy[N](t, g, nodes.Values(), sortables.Equality[N])

}

// TestDirectedGraph_AddEdge tests adding directed edge to the graph.
func TestDirectedGraph_AddEdge(t *testing.T) {
	// Helper

	// Types
	type N = *tests.SortableNode
	type E = edge.Comparable[*tests.SortableNode]

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
				expEdges.Append(edges.Comparable(it, that, true))
			}
		})

	shouldContainAllEdgesBy[E, N](t, g, expEdges.Values(), sortables.Equality[E])

}

// TestDirectedGraph_Neighbors tests retrieving neighbors for a node.
func TestDirectedGraph_Neighbors(t *testing.T) {
	// Helper

	// Types
	type N = *tests.SortableNode

	g := DigraphOf[N]()

	var neighbors = dicts.ComparableOf[N, []N]()

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				_ = g.AddEdge(it, that)
				if nodes, ok := neighbors.Get(it); !ok {
					neighbors.Put(it, []N{that})
				} else {
					neighbors.Put(it, append(nodes, that))
				}
			}
		})

	shouldContainAllNeighborsBy[N](t, g, neighbors, sortables.Equality[N])

}

// TestDirectedGraph_HasNode tests the presence of node in the graph.
func TestDirectedGraph_HasNode(t *testing.T) {
	// Helper

	// Types
	type N = *tests.SortableNode

	g := DigraphOf[N]()
	exp := random.Struct[N](15).Each(g.AddNode).Values()
	// Assert the presence of node
	shouldContainAllNodesBy[N](t, g, exp, sortables.Equality[N])

}

// TestDirectedGraph_HasEdge tests the presence of edge in the graph.
func TestDirectedGraph_HasEdge(t *testing.T) {
	// Helper

	// Types
	type N = *tests.SortableNode
	type E = edge.Comparable[*tests.SortableNode]

	g := DigraphOf[N]()

	var expEdges iterables.Slice[E]

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				g.AddEdge(it, that)
				expEdges.Append(edges.Comparable(it, that, true))
			}
		})

	shouldContainAllEdgesBy[E, N](t, g, expEdges.Values(), sortables.Equality[E])

}

// TestDirectedGraph_Nodes tests retrieving all node from the graph.
func TestDirectedGraph_Nodes(t *testing.T) {
	// Helper

	// Types
	type N = *tests.SortableNode

	g := DigraphOf[N]()

	exp := random.Struct[N](100).Each(g.AddNode)

	t.Logf("Nodes: %v\n", g.Nodes())
	t.Logf("Expected node: %v\n", exp.Values())

	shouldContainAllNodesBy[N](t, g, exp.Values(), sortables.Equality[N])
	shouldContainNNodes[N](t, g, exp.Len())

}

// TestDirectedGraph_Edges tests retrieving all directed edge from the graph.
func TestDirectedGraph_Edges(t *testing.T) {
	// Helper

	// Types
	type N = *tests.SortableNode
	type E = edge.Comparable[*tests.SortableNode]

	g := DigraphOf[N]()

	var expEdges = iterables.OfSlice[E]()

	random.Struct[N](30).
		Operation(func(i int, n *iterables.Slice[N]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				err := g.AddEdge(it, that)
				if err != nil {
					t.Errorf("Error adding edge: %v", err)
				}
				expEdges.Append(edges.Comparable(it, that, true))
			}
		})

	shouldHaveNEdges[E, N](t, g, expEdges.Len())
	shouldContainAllEdgesBy[E, N](t, g, expEdges.Values(), sortables.Equality[E])

}

func TestDirectedGraph_AddEdgeNodesNotExist(t *testing.T) {
	// Helper

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
		t.Errorf("Edge from A to C should not be present")
	}

}
