// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/testdata"
	"testing"
)

func TestAddNode(t *testing.T) {
	g := UndirectOf[testdata.TestNode]()

	nodeA := testdata.TestNode("A")
	g.AddNode(nodeA)

	// Assert that the node was added
	if !g.HasNode(nodeA) {
		t.Errorf("Node A should be present")
	}
}

func TestAddEdge(t *testing.T) {
	g := UndirectOf[testdata.TestNode]()

	nodeA := testdata.TestNode("A")
	nodeB := testdata.TestNode("B")

	g.AddNode(nodeA)
	g.AddNode(nodeB)
	g.AddEdge(nodeA, nodeB)

	if !g.HasEdge(nodeA, nodeB) {
		t.Errorf("Edge from A to B should be present")
	}

	if !g.HasEdge(nodeB, nodeA) {
		t.Errorf("Edge from B to A should be present")
	}
}

func TestNeighbors(t *testing.T) {
	g := UndirectOf[testdata.TestNode]()

	nodeA := testdata.TestNode("A")
	nodeB := testdata.TestNode("B")
	nodeC := testdata.TestNode("C")

	g.AddNode(nodeA)
	g.AddNode(nodeB)
	g.AddNode(nodeC)
	g.AddEdge(nodeA, nodeB)
	g.AddEdge(nodeA, nodeC)

	// Assert neighbors of nodeA
	neighbors := g.Neighbors(nodeA)

	if len(neighbors) != 2 {
		t.Errorf("Node A should have 2 neighbors")
	}

	if !g.HasEdge(nodeA, nodeB) {
		t.Errorf("Node A should have an edge to B")
	}

	if !g.HasEdge(nodeA, nodeC) {
		t.Errorf("Node A should have an edge to C")
	}

	// Assert neighbors of nodeB
	neighbors = g.Neighbors(nodeB)

	if len(neighbors) != 1 {
		t.Errorf("Node B should have 1 neighbors")
	}

	if !g.HasEdge(nodeB, nodeA) {
		t.Errorf("Node B should have an edge to A")
	}
}

func TestEdges(t *testing.T) {
	g := UndirectOf[testdata.TestNode]()

	nodeA := testdata.TestNode("A")
	nodeB := testdata.TestNode("B")
	nodeC := testdata.TestNode("C")

	g.AddNode(nodeA)
	g.AddNode(nodeB)
	g.AddNode(nodeC)
	g.AddEdge(nodeA, nodeB)
	g.AddEdge(nodeA, nodeC)

	edges := g.Edges()
	expectedEdges := []*SingleTypedEdge[testdata.TestNode]{
		NewEdge(nodeA, nodeB),
		NewEdge(nodeA, nodeC),
	}

	if !arrays.ContainsAllBy(edges, expectedEdges, func(edg, exp *SingleTypedEdge[testdata.TestNode]) bool {
		return edg.Equal(exp)
	}) {
		t.Errorf("should contain all expected edges: \nexpected: %v, \ngot: %v\n", expectedEdges, edges)
	}
}

func TestNodes(t *testing.T) {
	g := UndirectOf[testdata.TestNode]()

	nodeA := testdata.TestNode("A")
	nodeB := testdata.TestNode("B")
	nodeC := testdata.TestNode("C")

	g.AddNode(nodeA)
	g.AddNode(nodeB)
	g.AddNode(nodeC)

	nodes := g.Nodes()

	if len(nodes) != 3 {
		t.Errorf("Graph should have 3 nodes, got %d: %v", len(nodes), nodes)
	}

	if !arrays.ContainsAllBy(nodes, []testdata.TestNode{nodeA, nodeB, nodeC}, func(n1, n2 testdata.TestNode) bool {
		return n1 == n2
	}) {
		t.Errorf("Nodes should contain all expected nodes")
	}
}

func TestAddEdgeNodesNotExist(t *testing.T) {
	g := UndirectOf[testdata.TestNode]()

	nodeA := testdata.TestNode("A")
	nodeB := testdata.TestNode("B")

	// Add only one node
	g.AddNode(nodeA)

	// Attempt to add an edge with a non-existent node
	g.AddEdge(nodeA, nodeB)

	// Assert that the edge was not added
	if g.HasEdge(nodeA, nodeB) {
		t.Errorf("Edge from A to B should not be present")
	}
}
