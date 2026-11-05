// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/comparator/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/graph"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
	"github.com/andrerrcosta2/gtools/datastr/sets"
)

// TestIsCyclicOf tests the isCyclicOf function for various graph structures
func TestIsCyclicOf_UndirectedGraph(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type E = edge.Comparable[N]

	tcs := []struct {
		name     string
		setup    func() graph.Iterable[N]
		expected bool
	}{
		{
			name: "No cycle with two node and one edge",
			setup: func() graph.Iterable[N] {
				g := undirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C")) // Edge A-C
				return g
			},
			expected: false,
		},
		{
			name: "Cycle with three node in a triangle",
			setup: func() graph.Iterable[N] {
				g := undirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C")) // Edge A-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("C")) // Edge C-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("A")) // Edge C-A
				return g
			},
			expected: true,
		},
		{
			name: "Cycle with four node in a square",
			setup: func() graph.Iterable[N] {
				g := undirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C")) // Edge A-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("C")) // Edge C-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D")) // Edge C-D
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("A")) // Edge D-A
				return g
			},
			expected: true,
		},
		{
			name: "Cycle with four node and a diagonal",
			setup: func() graph.Iterable[N] {
				g := undirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C")) // Edge A-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("C")) // Edge C-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D")) // Edge C-D
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("A")) // Edge D-A
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C")) // Diagonal A-C
				return g
			},
			expected: true,
		},
		{
			name: "Self-loop detection",
			setup: func() graph.Iterable[N] {
				g := undirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("A")) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph",
			setup: func() graph.Iterable[N] {
				g := undirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C")) // Edge A-C
				// Next C is disconnected
				return g
			},
			expected: false,
		},
		{
			name: "Multiple disconnected components with cycles",
			setup: func() graph.Iterable[N] {
				g := undirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C")) // Edge A-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("C")) // Edge C-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("A")) // Edge C-A (Cycle)
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("D")) // Self-loop on D
				return g
			},
			expected: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			graph := tc.setup()
			visited := sets.HashSet(comparators.Entities[N]())
			recStack := sets.HashSet(comparators.Entities[N]())
			if got := isCyclic[N](graph, visited, recStack); got != tc.expected {
				t.Errorf("isCyclicOf() = %v, want %v\nGraph: %v\n", got, tc.expected, graph)
			}
		})
	}

}

// TestIsCyclicOf_DirectedGraph tests the isCyclicOf function for various directed graph structures
func TestIsCyclicOf_DirectedGraph(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type E = edge.Typed[N]

	tcs := []struct {
		name     string
		setup    func() graph.Iterable[N]
		expected bool
	}{
		{
			name: "No Cycle",
			setup: func() graph.Iterable[N] {
				g := DigraphOf[N]()

				a := tests.NewSortableNode("A")
				b := tests.NewSortableNode("B")
				c := tests.NewSortableNode("C")

				g.AddNode(a)
				g.AddNode(b)
				g.AddNode(c)

				g.AddEdge(a, b) // A -> B
				g.AddEdge(b, c) // B -> C

				return g
			},
			expected: false,
		},
		{
			name: "Executor Cycle",
			setup: func() graph.Iterable[N] {
				g := DigraphOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("A")) // Creates a cycle: A -> C -> C -> A
				return g
			},
			expected: true,
		},
		{
			name: "Self Loop",
			setup: func() graph.Iterable[N] {
				g := DigraphOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("A")) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Multi Component with Cycle",
			setup: func() graph.Iterable[N] {
				g := DigraphOf[N]()
				// First component (no cycle)
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C"))

				// Second component (with a cycle)
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddNode(tests.NewSortableNode("E"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("E"))
				g.AddEdge(tests.NewSortableNode("E"), tests.NewSortableNode("C")) // Creates a cycle: C -> D -> E -> C
				return g
			},
			expected: true,
		},
		{
			name: "Multi Disconnected Acyclic",
			setup: func() graph.Iterable[N] {
				g := DigraphOf[N]()

				// First component
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C"))

				// Second component
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"))
				return g
			},
			expected: false,
		},
		{
			name: "Large Cyclic Graph",
			setup: func() graph.Iterable[N] {
				g := DigraphOf[N]()
				// Adding node
				tests.NewSortableNodes("A", "C", "C", "D", "E", "F").Each(g.AddNode)

				// Adding edge (with a cycle)
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("E"))
				g.AddEdge(tests.NewSortableNode("E"), tests.NewSortableNode("F"))
				g.AddEdge(tests.NewSortableNode("F"), tests.NewSortableNode("C")) // Creates a cycle: C -> D -> E -> F -> C
				return g
			},
			expected: true,
		},
		{
			name: "Single Next No Edges",
			setup: func() graph.Iterable[N] {
				g := DigraphOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				return g
			},
			expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			graph := tc.setup()
			visited := sets.HashSet(comparators.Entities[N]())
			recStack := sets.HashSet(comparators.Entities[N]())
			result := isCyclic[N](graph, visited, recStack)
			if result != tc.expected {
				t.Errorf("isCyclicOf() = %v; want %v\nGraph:\n%v\n", result, tc.expected, graph)
			}
		})
	}
}

// TestIsCyclicOf_WeightedGraph tests the isCyclicOf function for various weighted graph structures
func TestIsCyclicOf_WeightedGraph(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edge.ComparableWeighted[N, W]

	tcs := []struct {
		name     string
		setup    func() graph.MutableWeighted[E, N, W]
		expected bool
	}{
		{
			name: "No cycle with four node in a square",
			setup: func() graph.MutableWeighted[E, N, W] {
				g := WeightedDigraphOf[N, W]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C"), 1) // Edge A-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("C"), 1) // Edge C-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"), 1) // Edge C-D
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("A"), 1) // Edge D-A
				return g
			},
			expected: true,
		},
		{
			name: "No cycle with three node in a line",
			setup: func() graph.MutableWeighted[E, N, W] {
				g := WeightedDigraphOf[N, W]()
				a := tests.NewSortableNode("A")
				b := tests.NewSortableNode("B")
				c := tests.NewSortableNode("C")
				g.AddNode(a)
				g.AddNode(b)
				g.AddNode(c)
				g.AddEdge(a, b, 1) // A -> B (weight 1)
				g.AddEdge(b, c, 1) // B -> C (weight 1)

				return g
			},
			expected: false,
		},

		{
			name: "Cycle with weighted edge",
			setup: func() graph.MutableWeighted[E, N, W] {
				g := WeightedDigraphOf[N, W]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C"), 1)
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("C"), 1)
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("A"), 1) // Creates a cycle: A -> C -> C -> A
				return g
			},
			expected: true,
		},
		{
			name: "Single node with a self-loop",
			setup: func() graph.MutableWeighted[E, N, W] {
				g := WeightedDigraphOf[N, W]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("A"), 1) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph with one component containing a cycle",
			setup: func() graph.MutableWeighted[E, N, W] {
				g := WeightedDigraphOf[N, W]()

				// First component (no cycle)
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C"), 1)

				// Second component (with a cycle)
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddNode(tests.NewSortableNode("E"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"), 1)
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("E"), 1)
				g.AddEdge(tests.NewSortableNode("E"), tests.NewSortableNode("C"), 1) // Creates a cycle: C -> D -> E -> C
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph with all acyclic components",
			setup: func() graph.MutableWeighted[E, N, W] {
				g := WeightedDigraphOf[N, W]()

				// First component
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C"), 1)

				// Second component
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"), 1)
				return g
			},
			expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			graph := tc.setup()
			visited := sets.HashSet(comparators.Entities[N]())
			recstack := sets.HashSet(comparators.Entities[N]())
			if result := isCyclic[N](graph, visited, recstack); result != tc.expected {
				t.Errorf("isCyclicOf() = %v; want %v\nGraph:\n%v\n", result, tc.expected, graph)
			}
		})
	}

}
