// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/testdata"
	"testing"
)

// TestIsCyclicOf tests the isCyclicOf function for various graph structures
func TestIsCyclicOf_UndirectedGraph(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() GraphOf[testdata.TestNode]
		expected bool
	}{
		{
			name: "No cycle with two nodes and one edge",
			setup: func() GraphOf[testdata.TestNode] {
				g := UndirectOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B")) // Edge A-B
				return g
			},
			expected: false,
		},
		{
			name: "Cycle with three nodes in a triangle",
			setup: func() GraphOf[testdata.TestNode] {
				g := UndirectOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B")) // Edge A-B
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C")) // Edge B-C
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("A")) // Edge C-A
				return g
			},
			expected: true,
		},
		{
			name: "Cycle with four nodes in a square",
			setup: func() GraphOf[testdata.TestNode] {
				g := UndirectOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddNode(testdata.TestNode("D"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B")) // Edge A-B
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C")) // Edge B-C
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("D")) // Edge C-D
				g.AddEdge(testdata.TestNode("D"), testdata.TestNode("A")) // Edge D-A
				return g
			},
			expected: true,
		},
		{
			name: "Cycle with four nodes and a diagonal",
			setup: func() GraphOf[testdata.TestNode] {
				g := UndirectOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddNode(testdata.TestNode("D"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B")) // Edge A-B
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C")) // Edge B-C
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("D")) // Edge C-D
				g.AddEdge(testdata.TestNode("D"), testdata.TestNode("A")) // Edge D-A
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("C")) // Diagonal A-C
				return g
			},
			expected: true,
		},
		{
			name: "Self-loop detection",
			setup: func() GraphOf[testdata.TestNode] {
				g := UndirectOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("A")) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph",
			setup: func() GraphOf[testdata.TestNode] {
				g := UndirectOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B")) // Edge A-B
				// Node C is disconnected
				return g
			},
			expected: false,
		},
		{
			name: "Multiple disconnected components with cycles",
			setup: func() GraphOf[testdata.TestNode] {
				g := UndirectOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddNode(testdata.TestNode("D"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B")) // Edge A-B
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C")) // Edge B-C
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("A")) // Edge C-A (Cycle)
				g.AddEdge(testdata.TestNode("D"), testdata.TestNode("D")) // Self-loop on D
				return g
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			if got := isCyclicOf[testdata.TestNode](graph); got != tt.expected {
				t.Errorf("isCyclicOf() = %v, want %v\nGraph: %v\n", got, tt.expected, graph)
			}
		})
	}
}

func TestIsCyclicOf_DirectedGraph(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() GraphOf[testdata.TestNode]
		expected bool
	}{
		{
			name: "No Cycle",
			setup: func() GraphOf[testdata.TestNode] {
				g := DigraphOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C"))
				return g
			},
			expected: false,
		},
		{
			name: "Simple Cycle",
			setup: func() GraphOf[testdata.TestNode] {
				g := DigraphOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C"))
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("A")) // Creates a cycle: A -> B -> C -> A
				return g
			},
			expected: true,
		},
		{
			name: "Self Loop",
			setup: func() GraphOf[testdata.TestNode] {
				g := DigraphOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("A")) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Multi Component with Cycle",
			setup: func() GraphOf[testdata.TestNode] {
				g := DigraphOf[testdata.TestNode]()
				// First component (no cycle)
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))

				// Second component (with a cycle)
				g.AddNode(testdata.TestNode("C"))
				g.AddNode(testdata.TestNode("D"))
				g.AddNode(testdata.TestNode("E"))
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("D"))
				g.AddEdge(testdata.TestNode("D"), testdata.TestNode("E"))
				g.AddEdge(testdata.TestNode("E"), testdata.TestNode("C")) // Creates a cycle: C -> D -> E -> C
				return g
			},
			expected: true,
		},
		{
			name: "Multi Disconnected Acyclic",
			setup: func() GraphOf[testdata.TestNode] {
				g := DigraphOf[testdata.TestNode]()

				// First component
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))

				// Second component
				g.AddNode(testdata.TestNode("C"))
				g.AddNode(testdata.TestNode("D"))
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("D"))
				return g
			},
			expected: false,
		},
		{
			name: "Large Cyclic Graph",
			setup: func() GraphOf[testdata.TestNode] {
				g := DigraphOf[testdata.TestNode]()
				// Adding nodes
				nodes := []testdata.TestNode{"A", "B", "C", "D", "E", "F"}
				for _, node := range nodes {
					g.AddNode(node)
				}

				// Adding edges (with a cycle)
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"))
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C"))
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("D"))
				g.AddEdge(testdata.TestNode("D"), testdata.TestNode("E"))
				g.AddEdge(testdata.TestNode("E"), testdata.TestNode("F"))
				g.AddEdge(testdata.TestNode("F"), testdata.TestNode("C")) // Creates a cycle: C -> D -> E -> F -> C
				return g
			},
			expected: true,
		},
		{
			name: "Single Node No Edges",
			setup: func() GraphOf[testdata.TestNode] {
				g := DigraphOf[testdata.TestNode]()
				g.AddNode(testdata.TestNode("A"))
				return g
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			result := isCyclicOf[testdata.TestNode](graph)
			if result != tt.expected {
				t.Errorf("isCyclicOf() = %v; want %v\nGraph:\n%v\n", result, tt.expected, graph)
			}
		})
	}
}

func TestIsCyclicOf_WeightedGraph(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() Graph[testdata.TestNode]
		expected bool
	}{
		{
			name: "No cycle with four nodes in a square",
			setup: func() Graph[testdata.TestNode] {
				g := WeightedOrderedOf[testdata.TestNode, int]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddNode(testdata.TestNode("D"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"), 1) // Edge A-B
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C"), 1) // Edge B-C
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("D"), 1) // Edge C-D
				g.AddEdge(testdata.TestNode("D"), testdata.TestNode("A"), 1) // Edge D-A
				return g
			},
			expected: true,
		},
		{
			name: "No cycle with three nodes in a line",
			setup: func() Graph[testdata.TestNode] {
				g := WeightedOrderedOf[testdata.TestNode, int]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"), 1)
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C"), 1)
				return g
			},
			expected: false,
		},
		{
			name: "Cycle with weighted edges",
			setup: func() Graph[testdata.TestNode] {
				g := WeightedOrderedOf[testdata.TestNode, int]()
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddNode(testdata.TestNode("C"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"), 1)
				g.AddEdge(testdata.TestNode("B"), testdata.TestNode("C"), 1)
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("A"), 1) // Creates a cycle: A -> B -> C -> A
				return g
			},
			expected: true,
		},
		{
			name: "Single node with a self-loop",
			setup: func() Graph[testdata.TestNode] {
				g := WeightedOrderedOf[testdata.TestNode, int]()
				g.AddNode(testdata.TestNode("A"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("A"), 1) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph with one component containing a cycle",
			setup: func() Graph[testdata.TestNode] {
				g := WeightedOrderedOf[testdata.TestNode, int]()

				// First component (no cycle)
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"), 1)

				// Second component (with a cycle)
				g.AddNode(testdata.TestNode("C"))
				g.AddNode(testdata.TestNode("D"))
				g.AddNode(testdata.TestNode("E"))
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("D"), 1)
				g.AddEdge(testdata.TestNode("D"), testdata.TestNode("E"), 1)
				g.AddEdge(testdata.TestNode("E"), testdata.TestNode("C"), 1) // Creates a cycle: C -> D -> E -> C
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph with all acyclic components",
			setup: func() Graph[testdata.TestNode] {
				g := WeightedOrderedOf[testdata.TestNode, int]()

				// First component
				g.AddNode(testdata.TestNode("A"))
				g.AddNode(testdata.TestNode("B"))
				g.AddEdge(testdata.TestNode("A"), testdata.TestNode("B"), 1)

				// Second component
				g.AddNode(testdata.TestNode("C"))
				g.AddNode(testdata.TestNode("D"))
				g.AddEdge(testdata.TestNode("C"), testdata.TestNode("D"), 1)
				return g
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			if result := isCyclicOf(graph); result != tt.expected {
				t.Errorf("isCyclicOf() = %v; want %v\nGraph:\n%v\n", result, tt.expected, graph)
			}
		})
	}
}
