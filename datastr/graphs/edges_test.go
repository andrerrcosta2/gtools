// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
	"github.com/google/uuid"
	"testing"
)

func TestSingleTypedEdge_Construction(t *testing.T) {
	from := tests.NewSortableNode("A")
	to := tests.NewSortableNode("B")
	edge := PersistentSortableEdge(uuid.NewString(), from, to, false)

	if edge.From() != from {
		t.Errorf("From() = %v, want %v", edge.From(), from)
	}
	if edge.To() != to {
		t.Errorf("To() = %v, want %v", edge.To(), to)
	}
}

func TestSingleTypedEdge_Unique(t *testing.T) {
	from := tests.NewSortableNode("A")
	to := tests.NewSortableNode("B")

	edge := SortableEdge(from, to, false)
	expectedUnique := "A<->B"

	if edge.Unique() != expectedUnique {
		t.Errorf("Unique() = %v, want %v", edge.Unique(), expectedUnique)
	}
}

func TestSingleTypedEdge_Equal(t *testing.T) {
	t.Run("PersistentSortableEdge", func(t *testing.T) {
		hash1 := uuid.NewString()
		hash2 := uuid.NewString()

		from1 := tests.NewSortableNode("A")
		to1 := tests.NewSortableNode("B")
		edge1 := PersistentSortableEdge(hash1, from1, to1, false)

		from2 := tests.NewSortableNode("B")
		to2 := tests.NewSortableNode("A")
		edge2 := PersistentSortableEdge(hash1, from2, to2, false)

		if !edge1.Equal(edge2) {
			t.Errorf("Equal() = false, want true")
		}

		// Test with different hashes
		edge2 = PersistentSortableEdge(hash2, from2, to2, false)
		if edge1.Equal(edge2) {
			t.Errorf("Equal() = true, want false")
		}
	})

	t.Run("SortableEdge", func(t *testing.T) {
		from1 := tests.NewSortableNode("A")
		to1 := tests.NewSortableNode("B")
		edge1 := SortableEdge(from1, to1, false)

		from2 := tests.NewSortableNode("A")
		to2 := tests.NewSortableNode("B")
		edge2 := SortableEdge(from2, to2, false)

		if !edge1.Equal(edge2) {
			t.Errorf("Equal() = true, want false")
		}
	})

}

func TestSingleTypedEdge_Less(t *testing.T) {
	from1 := tests.NewSortableNode("A")
	to1 := tests.NewSortableNode("B")
	edge1 := PersistentSortableEdge(uuid.NewString(), from1, to1, false)

	from2 := tests.NewSortableNode("A")
	to2 := tests.NewSortableNode("C")
	edge2 := PersistentSortableEdge(uuid.NewString(), from2, to2, false)

	if edge1.Less(edge2) != (edge1.Unique() < edge2.Unique()) {
		t.Errorf("Less() = false, want true")
	}
}

func TestSingleTypedWeightedEdge_Construction(t *testing.T) {
	from := tests.NewSortableNode("A")
	to := tests.NewSortableNode("B")
	weight := 10
	edge := PersistentSingleTypedWeightedEdge(uuid.NewString(), from, to, weight, true)

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

func TestSingleTypedWeightedEdge_Unique_String(t *testing.T) {
	from := tests.NewSortableNode("A")
	to := tests.NewSortableNode("B")
	weight := 10
	hash := uuid.NewString()
	edge := PersistentSingleTypedWeightedEdge(hash, from, to, weight, true)
	expString := "A->B"

	if fmt.Sprintf("%s", edge) != expString {
		t.Errorf("Unique() = %v, want %v\n", fmt.Sprintf("%v", edge), expString)
	}

	if edge.Unique() != hash {
		t.Errorf("Unique() = %v, want %v\n", edge.Unique(), hash)
	}
}

func TestSingleTypedWeightedEdge_Equal(t *testing.T) {
	hash1 := uuid.NewString()
	hash2 := uuid.NewString()

	from1 := tests.NewSortableNode("A")
	to1 := tests.NewSortableNode("B")
	weight1 := 10
	edge1 := PersistentSingleTypedWeightedEdge(hash1, from1, to1, weight1, true)

	from2 := tests.NewSortableNode("A")
	to2 := tests.NewSortableNode("B")
	weight2 := 10
	edge2 := PersistentSingleTypedWeightedEdge(hash1, from2, to2, weight2, true)

	if !edge1.Equal(edge2) {
		t.Errorf("Equal() = false, want true")
	}

	// Test with different hashes
	edge2 = PersistentSingleTypedWeightedEdge(hash2, from2, to2, weight2, true)
	if edge1.Equal(edge2) {
		t.Errorf("Equal() = true, want false")
	}
}

func TestSingleTypedWeightedEdge_Less(t *testing.T) {
	from1 := tests.NewSortableNode("A")
	to1 := tests.NewSortableNode("B")
	weight1 := random.Int(1, 100).At(0)
	edge1 := PersistentSingleTypedWeightedEdge(uuid.NewString(), from1, to1, weight1, false)

	from2 := tests.NewSortableNode("A")
	to2 := tests.NewSortableNode("C")
	weight2 := random.Int(1, 100).At(0)
	edge2 := PersistentSingleTypedWeightedEdge(uuid.NewString(), from2, to2, weight2, false)

	if edge1.Less(edge2) != (weight1 < weight2) {
		t.Errorf("Less() = false, want true")
	}
}
