// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"github.com/andrerrcosta2/gtools/pkg/testdata/testsortables"
	"testing"
)

// TestAddNode tests adding a node to the undirected graph
func TestAddNode(t *testing.T) {
	g := UndirectOf[testsortables.TestNode]()

	exp := testsortables.RandomTestNodes(10, "node").Each(g.AddNode)
	nodesN, expN := g.Nodes(), exp.Values()

	arrays.ContainsAllBy(&nodesN, &expN, functions.ImplementedEquality[testsortables.TestNode])
}

func TestAddEdge(t *testing.T) {
	g := UndirectOf[testsortables.TestNode]()

	var expEdges iterables.Slice[*SingleTypedEdge[testsortables.TestNode]]

	testsortables.RandomTestNodes(30, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				g.AddEdge(it, that)
				expEdges.Append(NewEdge(it, that))
			}
		})

	expEdges.Each(func(e *SingleTypedEdge[testsortables.TestNode]) {
		if !g.HasEdge(e.From(), e.To()) {
			t.Errorf("Expected edge from %s to %s to exist in the graph", e.From(), e.To())
		}
	})
}

func TestNeighbors(t *testing.T) {
	g := UndirectOf[testsortables.TestNode]()

	var neighbors iterables.SliceMap[testsortables.TestNode, testsortables.TestNode] = map[testsortables.TestNode][]testsortables.TestNode{}
	testsortables.RandomTestNodes(10, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				g.AddEdge(it, that)
				// Undirected graphs have neighbors in both directions
				neighbors.PutOrAppend(it, that)
				neighbors.PutOrAppend(that, it)
			}
		})

	for _, node := range g.Nodes() {
		graphN := g.Neighbors(node)
		expN := neighbors.At(node)

		// Check if the number of neighbors is as expected
		if len(graphN) != len(expN) {
			t.Errorf("Expected %d neighbors for node %s, got %d", len(expN), node, len(graphN))
		}

		// Check if all expected neighbors are present
		if !arrays.ContainsAllBy(&expN, &graphN, functions.ImplementedEquality[testsortables.TestNode]) {
			t.Errorf("Neighbors for node %s do not match expected neighbors", node)
		}

		// Check if the graph's neighbor list is a subset of the expected neighbors (and vice versa)
		if !arrays.ContainsAllBy(&graphN, &expN, functions.ImplementedEquality[testsortables.TestNode]) {
			t.Errorf("Neighbors for node %s do not match expected neighbors", node)
		}

		// Check if the expected neighbors are a subset of the graph's neighbor list (and vice versa)
		if !arrays.ContainsAllBy(&expN, &graphN, functions.ImplementedEquality[testsortables.TestNode]) {
			t.Errorf("Neighbors for node %s do not match expected neighbors", node)
		}
	}
}

func TestEdges(t *testing.T) {
	g := UndirectOf[testsortables.TestNode]()

	var expEdges iterables.Slice[*SingleTypedEdge[testsortables.TestNode]]

	testsortables.RandomTestNodes(30, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			g.AddNode(n.At(i))
			if i%3 == 0 && i+1 < n.Len() && i-1 >= 0 {
				it, that := n.At(i), n.At(i-1)
				g.AddEdge(it, that)
				expEdges.Append(NewEdge(it, that))
			}
		})

	expEdges.Each(func(e *SingleTypedEdge[testsortables.TestNode]) {
		if !g.HasEdge(e.From(), e.To()) {
			t.Errorf("Expected edge from %s to %s to exist in the graph", e.From(), e.To())
		}
	})
}

func TestNodes(t *testing.T) {
	g := UndirectOf[testsortables.TestNode]()

	exp := testsortables.RandomTestNodes(10, "node").Each(g.AddNode)

	exp.Each(func(n testsortables.TestNode) {
		if !g.HasNode(n) {
			t.Errorf("Expected node %s to be present in the graph", n)
		}
	})
}

func TestAddEdgeNodesNotExist(t *testing.T) {
	g := UndirectOf[testsortables.TestNode]()

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
