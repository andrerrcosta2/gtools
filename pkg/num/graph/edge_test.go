// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/testdata/testsortables"
	"testing"
)

func TestSingleTypedEdge_Construction(t *testing.T) {
	from := testsortables.TestNode("A")
	to := testsortables.TestNode("B")
	edge := NewEdge(from, to)

	if edge.From() != from {
		t.Errorf("From() = %v, want %v", edge.From(), from)
	}
	if edge.To() != to {
		t.Errorf("To() = %v, want %v", edge.To(), to)
	}
}

func TestSingleTypedEdge_Unique(t *testing.T) {
	from := testsortables.TestNode("A")
	to := testsortables.TestNode("B")

	edge := NewEdge(from, to)
	expectedUnique := "A <-> B"

	if edge.Unique() != expectedUnique {
		t.Errorf("Unique() = %v, want %v", edge.Unique(), expectedUnique)
	}
}

func TestSingleTypedEdge_Equal(t *testing.T) {
	from1 := testsortables.TestNode("A")
	to1 := testsortables.TestNode("B")
	edge1 := NewEdge(from1, to1)

	from2 := testsortables.TestNode("A")
	to2 := testsortables.TestNode("B")
	edge2 := NewEdge(from2, to2)

	if !edge1.Equal(edge2) {
		t.Errorf("Equal() = false, want true")
	}
}

func TestSingleTypedEdge_Less(t *testing.T) {
	from1 := testsortables.TestNode("A")
	to1 := testsortables.TestNode("B")
	edge1 := NewEdge(from1, to1)

	from2 := testsortables.TestNode("A")
	to2 := testsortables.TestNode("C")
	edge2 := NewEdge(from2, to2)

	if !edge1.Less(edge2) {
		t.Errorf("Less() = false, want true")
	}
}

func TestSingleTypedWeightedEdge_Construction(t *testing.T) {
	from := testsortables.TestNode("A")
	to := testsortables.TestNode("B")
	weight := 10
	edge := NewWeightedEdge(from, to, weight)

	if edge.From() != from {
		t.Errorf("From() = %v, want %v", edge.From(), from)
	}
	if edge.To() != to {
		t.Errorf("To() = %v, want %v", edge.To(), to)
	}
	if edge.Weight() != weight {
		t.Errorf("Weight() = %v, want %v", edge.Weight(), weight)
	}
}

func TestSingleTypedWeightedEdge_Unique(t *testing.T) {
	from := testsortables.TestNode("A")
	to := testsortables.TestNode("B")
	weight := 10
	edge := NewWeightedEdge(from, to, weight)
	expectedUnique := "A->B"

	if edge.Unique() != expectedUnique {
		t.Errorf("Unique() = %v, want %v", edge.Unique(), expectedUnique)
	}
}

func TestSingleTypedWeightedEdge_Equal(t *testing.T) {
	from1 := testsortables.TestNode("A")
	to1 := testsortables.TestNode("B")
	weight1 := 10
	edge1 := NewWeightedEdge(from1, to1, weight1)

	from2 := testsortables.TestNode("A")
	to2 := testsortables.TestNode("B")
	weight2 := 10
	edge2 := NewWeightedEdge(from2, to2, weight2)

	if !edge1.Equal(edge2) {
		t.Errorf("Equal() = false, want true")
	}
}

func TestSingleTypedWeightedEdge_Less(t *testing.T) {
	from1 := testsortables.TestNode("A")
	to1 := testsortables.TestNode("B")
	weight1 := 10
	edge1 := NewWeightedEdge(from1, to1, weight1)

	from2 := testsortables.TestNode("A")
	to2 := testsortables.TestNode("C")
	weight2 := 15
	edge2 := NewWeightedEdge(from2, to2, weight2)

	if !edge1.Less(edge2) {
		t.Errorf("Less() = false, want true")
	}
}
