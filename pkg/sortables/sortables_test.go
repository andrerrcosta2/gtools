// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/sorts"
	"github.com/andrerrcosta2/gtools/pkg/testdata"
	"testing"
	"unsafe"
)

func TestComparatorSortableOf_Compare(t *testing.T) {
	comp := ComparatorOf[testdata.TestNode]()
	a := testdata.TestNode("A")
	b := testdata.TestNode("B")
	c := testdata.TestNode("A")

	if comp.Compare(a, b) != 1 {
		t.Errorf("Compare(A, B) = %d, want 1", comp.Compare(a, b))
	}
	if comp.Compare(b, a) != -1 {
		t.Errorf("Compare(B, A) = %d, want -1", comp.Compare(b, a))
	}
	if comp.Compare(a, c) != 0 {
		t.Errorf("Compare(A, C) = %d, want 0", comp.Compare(a, c))
	}
}

func TestComparatorSortableOf_Equals(t *testing.T) {
	comp := ComparatorOf[testdata.TestNode]()
	a := testdata.TestNode("A")
	b := testdata.TestNode("A")
	c := testdata.TestNode("B")

	if !comp.Equals(a, b) {
		t.Errorf("Equals(A, B) = false, want true")
	}
	if comp.Equals(a, c) {
		t.Errorf("Equals(A, C) = true, want false")
	}
}

func TestUnique(t *testing.T) {
	// Define some test values
	intVal := 42
	floatVal := 3.14
	stringVal := "hello"
	boolVal := true

	// Define expected results
	expectedUniqueInt := "42"
	expectedUniqueFloat := "3.14"
	expectedUniqueString := "hello"
	expectedUniqueBool := "true"

	// Test for int
	if unique := Unique(intVal); unique != expectedUniqueInt {
		t.Errorf("Unique(intVal) = %v, want %v", unique, expectedUniqueInt)
	}

	// Test for float
	if unique := Unique(floatVal); unique != expectedUniqueFloat {
		t.Errorf("Unique(floatVal) = %v, want %v", unique, expectedUniqueFloat)
	}

	// Test for string
	if unique := Unique(stringVal); unique != expectedUniqueString {
		t.Errorf("Unique(stringVal) = %v, want %v", unique, expectedUniqueString)
	}

	// Test for bool
	if unique := Unique(boolVal); unique != expectedUniqueBool {
		t.Errorf("Unique(boolVal) = %v, want %v", unique, expectedUniqueBool)
	}

	// Test for TestStruct
	structVal := testdata.TestStruct{Name: "Alice", Age: 30}
	expectedUniqueStruct := "{Alice 30}"

	if unique := Unique(structVal); unique != expectedUniqueStruct {
		t.Errorf("Unique(structVal) = %v, want %v", unique, expectedUniqueStruct)
	}

	// Test for pointer to TestStruct
	structPtr := &testdata.TestStruct{Name: "Bob", Age: 40}
	expectedUniqueStructPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(structPtr)))

	if unique := Unique(structPtr); unique != expectedUniqueStructPtr {
		t.Errorf("Unique(structPtr) = %v, want %v", unique, expectedUniqueStructPtr)
	}
}

func TestSort(t *testing.T) {
	sortAlgorithm := sorts.NewQuicksort[testdata.TestNode](testdata.NewTestNodeComparator())

	data := []testdata.TestNode{testdata.TestNode("C"), testdata.TestNode("A"), testdata.TestNode("B")}

	expected := []testdata.TestNode{testdata.TestNode("A"), testdata.TestNode("B"), testdata.TestNode("C")}

	Sort(&data, sortAlgorithm)

	for i, v := range expected {
		if data[i] != v {
			t.Errorf("Sort() = %v, want %v", data, expected)
			return
		}
	}
}

func TestEqualsOf(t *testing.T) {
	a := testdata.TestNode("A")
	b := testdata.TestNode("A")
	c := testdata.TestNode("B")

	if !EqualsOf(a, b) {
		t.Errorf("EqualsOf(A, B) = false, want true")
	}
	if EqualsOf(a, c) {
		t.Errorf("EqualsOf(A, C) = true, want false")
	}
}
