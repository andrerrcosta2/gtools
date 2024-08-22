// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"github.com/andrerrcosta2/gtools/pkg/testdata/testsortables"
	"testing"
)

// TestDirectedGraph_AddNode tests adding nodes to the graph.
func TestDirectedGraph_AddNode(t *testing.T) {
	graph := DigraphOf[testsortables.TestNode]()

	exp := testsortables.RandomTestNodes(10, "node").Each(graph.AddNode)
	graphN, expN := graph.Nodes(), exp.Values()
	arrays.ContainsAllBy(&graphN, &expN, functions.ImplementedEquality[testsortables.TestNode])
}

// TestDirectedGraph_AddEdge tests adding directed edges to the graph.
func TestDirectedGraph_AddEdge(t *testing.T) {
	graph := DigraphOf[testsortables.TestNode]()

	var expEdges iterables.Slice[*SingleTypedEdge[testsortables.TestNode]]

	testsortables.RandomTestNodes(30, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			graph.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				graph.AddEdge(it, that)
				expEdges.Append(NewEdge(it, that))
			}
		})

	expEdges.Each(func(e *SingleTypedEdge[testsortables.TestNode]) {
		if !graph.HasEdge(e.From(), e.To()) {
			t.Errorf("Expected edge from %s to %s to exist in the graph", e.From(), e.To())
		}
	})
}

// TestDirectedGraph_Neighbors tests retrieving neighbors for a node.
func TestDirectedGraph_Neighbors(t *testing.T) {
	graph := DigraphOf[testsortables.TestNode]()

	var neighbors iterables.SliceMap[testsortables.TestNode, testsortables.TestNode] = map[testsortables.TestNode][]testsortables.TestNode{}
	testsortables.RandomTestNodes(10, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			graph.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				graph.AddEdge(it, that)
				neighbors.PutOrAppend(it, that)
			}
		})

	for _, node := range graph.Nodes() {
		graphN := graph.Neighbors(node)
		expN := neighbors.At(node)
		if len(graphN) != len(expN) {
			t.Errorf("Expected %d neighbors for node %s, got %d", len(neighbors.At(node)), node, len(graph.Neighbors(node)))
		}
		arrays.ContainsAllBy(&expN, &graphN, functions.ImplementedEquality[testsortables.TestNode])
	}
}

// TestDirectedGraph_HasNode tests the presence of nodes in the graph.
func TestDirectedGraph_HasNode(t *testing.T) {
	graph := DigraphOf[testsortables.TestNode]()
	exp := testsortables.RandomTestNodes(10, "node").Each(graph.AddNode)

	exp.Each(func(n testsortables.TestNode) {
		if !graph.HasNode(n) {
			t.Errorf("Expected node %s to be present in the graph", n)
		}
	})
}

// TestDirectedGraph_HasEdge tests the presence of edges in the graph.
func TestDirectedGraph_HasEdge(t *testing.T) {
	graph := DigraphOf[testsortables.TestNode]()

	var expEdges iterables.Slice[*SingleTypedEdge[testsortables.TestNode]]

	testsortables.RandomTestNodes(30, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			graph.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				graph.AddEdge(it, that)
				expEdges.Append(NewEdge(it, that))
			}
		})

	expEdges.Each(func(e *SingleTypedEdge[testsortables.TestNode]) {
		if !graph.HasEdge(e.From(), e.To()) {
			t.Errorf("Expected edge from %s to %s to exist in the graph", e.From(), e.To())
		}
	})
}

// TestDirectedGraph_Nodes tests retrieving all nodes from the graph.
func TestDirectedGraph_Nodes(t *testing.T) {
	graph := DigraphOf[testsortables.TestNode]()
	exp := testsortables.RandomTestNodes(100, "node").Each(graph.AddNode)

	fmt.Printf("Nodes: %v\n", graph.Nodes())
	fmt.Printf("Expected nodes: %v\n", exp.Values())

	if len(graph.Nodes()) != exp.Len() {
		t.Fatalf("Expected %d nodes, got %d", len(graph.Nodes()), exp.Len())
	}

	graphN, expN := graph.Nodes(), exp.Values()

	arrays.ContainsAllBy(&expN, &graphN, functions.ImplementedEquality[testsortables.TestNode])
}

// TestDirectedGraph_Edges tests retrieving all directed edges from the graph.
func TestDirectedGraph_Edges(t *testing.T) {
	graph := DigraphOf[testsortables.TestNode]()

	var expEdges iterables.Slice[*SingleTypedEdge[testsortables.TestNode]]

	testsortables.RandomTestNodes(30, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			graph.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				graph.AddEdge(it, that)
				expEdges.Append(NewEdge(it, that))
			}
		})

	edges := iterables.OfSlice(graph.Edges()...)

	if edges.Len() != expEdges.Len() {
		t.Fatalf("Expected %d edge(s), got %d\nGraph:\n%v", expEdges.Len(), edges.Len(), graph)
	}

	edgesN, expN := edges.Values(), expEdges.Values()

	arrays.ContainsAllBy(&edgesN, &expN, func(e1, e2 *SingleTypedEdge[testsortables.TestNode]) bool {
		return e1.From().Equal(e2.From()) && e1.To().Equal(e2.To())
	})
}

func TestDirectedGraph_AddEdgeNodesNotExist(t *testing.T) {
	g := DigraphOf[testsortables.TestNode]()

	nodeA := testsortables.TestNode("A")
	nodeB := testsortables.TestNode("B")

	// Add only one node
	g.AddNode(nodeA)

	// Attempt to add an edge with a non-existent node
	g.AddEdge(nodeA, nodeB)

	// Assert that the edge was not added
	if g.HasEdge(nodeA, nodeB) {
		t.Errorf("Edge from A to B should not be present")
	}
}
