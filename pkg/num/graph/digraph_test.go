// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/sorts"
	"github.com/andrerrcosta2/gtools/pkg/testdata"
	"testing"
)

// TestDirectedGraph_AddNode tests adding nodes to the graph.
func TestDirectedGraph_AddNode(t *testing.T) {
	graph := DigraphOf[testdata.TestNode]()

	graph.AddNode(testdata.TestNode("A"))
	graph.AddNode(testdata.TestNode("B"))

	if !graph.HasNode(testdata.TestNode("A")) || !graph.HasNode(testdata.TestNode("B")) {
		t.Errorf("Expected nodes A and B to be present in the graph")
	}
}

// TestDirectedGraph_AddEdge tests adding directed edges to the graph.
func TestDirectedGraph_AddEdge(t *testing.T) {
	graph := DigraphOf[testdata.TestNode]()

	graph.AddNode(testdata.TestNode("A"))
	graph.AddNode(testdata.TestNode("B"))
	graph.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))

	if !graph.HasEdge(testdata.TestNode("A"), testdata.TestNode("B")) {
		t.Errorf("Expected edge from A to B to exist in the graph")
	}

	if graph.HasEdge(testdata.TestNode("B"), testdata.TestNode("A")) {
		t.Errorf("Expected no edge from B to A in the graph")
	}
}

// TestDirectedGraph_Neighbors tests retrieving neighbors for a node.
func TestDirectedGraph_Neighbors(t *testing.T) {
	graph := DigraphOf[testdata.TestNode]()

	graph.AddNode(testdata.TestNode("A"))
	graph.AddNode(testdata.TestNode("B"))
	graph.AddNode(testdata.TestNode("C"))
	graph.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))
	graph.AddEdge(testdata.TestNode("A"), testdata.TestNode("C"))

	neighbors := graph.Neighbors(testdata.TestNode("A"))
	expectedNeighbors := []testdata.TestNode{"B", "C"}

	if len(neighbors) != len(expectedNeighbors) {
		t.Fatalf("Expected %d neighbors, got %d", len(expectedNeighbors), len(neighbors))
	}

	for i, neighbor := range neighbors {
		if neighbor != expectedNeighbors[i] {
			t.Errorf("Expected neighbor %s, got %s", expectedNeighbors[i], neighbor)
		}
	}
}

// TestDirectedGraph_HasNode tests the presence of nodes in the graph.
func TestDirectedGraph_HasNode(t *testing.T) {
	graph := DigraphOf[testdata.TestNode]()

	graph.AddNode(testdata.TestNode("A"))

	if !graph.HasNode(testdata.TestNode("A")) {
		t.Errorf("Expected node A to be present in the graph")
	}

	if graph.HasNode(testdata.TestNode("B")) {
		t.Errorf("Expected node B to not be present in the graph")
	}
}

// TestDirectedGraph_HasEdge tests the presence of edges in the graph.
func TestDirectedGraph_HasEdge(t *testing.T) {
	graph := DigraphOf[testdata.TestNode]()

	graph.AddNode(testdata.TestNode("A"))
	graph.AddNode(testdata.TestNode("B"))
	graph.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))

	if !graph.HasEdge(testdata.TestNode("A"), testdata.TestNode("B")) {
		t.Errorf("Expected edge from A to B to exist in the graph")
	}

	if graph.HasEdge(testdata.TestNode("B"), testdata.TestNode("A")) {
		t.Errorf("Expected no edge from B to A in the graph")
	}
}

// TestDirectedGraph_Nodes tests retrieving all nodes from the graph.
func TestDirectedGraph_Nodes(t *testing.T) {
	graph := DigraphOf[testdata.TestNode]()

	graph.AddNode(testdata.TestNode("A"))
	graph.AddNode(testdata.TestNode("B"))

	nodes := graph.Nodes()
	expectedNodes := []testdata.TestNode{"A", "B"}

	sorts.QuickOf(&nodes)
	sorts.QuickOf(&expectedNodes)

	if len(nodes) != len(expectedNodes) {
		t.Fatalf("Expected %d nodes, got %d", len(expectedNodes), len(nodes))
	}

	for i, node := range nodes {
		if node != expectedNodes[i] {
			t.Errorf("Expected node %s, got %s", expectedNodes[i], node)
		}
	}
}

// TestDirectedGraph_Edges tests retrieving all directed edges from the graph.
func TestDirectedGraph_Edges(t *testing.T) {
	graph := DigraphOf[testdata.TestNode]()

	graph.AddNode(testdata.TestNode("A"))
	graph.AddNode(testdata.TestNode("B"))
	graph.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))

	edges := graph.Edges()
	if len(edges) != 1 {
		t.Fatalf("Expected 1 edge, got %d", len(edges))
	}

	expectedEdge := NewEdge(testdata.TestNode("A"), testdata.TestNode("B"))
	if !edges[0].Equal(expectedEdge) {
		t.Errorf("Expected edge %v, got %v", expectedEdge, edges[0])
	}
}
