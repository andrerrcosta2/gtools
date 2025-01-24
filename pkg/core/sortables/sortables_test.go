// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/core/internal/testseed"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"testing"
	"unsafe"
)

func TestComparatorSortableOf_Compare(t *testing.T) {
	comp := ComparatorOf[*testseed.SortableNode]()
	a := testseed.NewSortableNode("A")
	b := testseed.NewSortableNode("B")
	c := testseed.NewSortableNode("A")

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
	comp := ComparatorOf[*testseed.SortableNode]()
	a := testseed.NewSortableNode("A")
	b := testseed.NewSortableNode("A")
	c := testseed.NewSortableNode("B")

	if !comp.Equals(a, b) {
		t.Errorf("Equals(A, B) = false, want true")
	}
	if comp.Equals(a, c) {
		t.Errorf("Equals(A, C) = true, want false")
	}
}

func TestUnique_Primitives_Values(t *testing.T) {
	// Define some test values
	intVal := 42
	floatVal := 3.14
	stringVal := "hello"
	boolVal := true

	// Define expected results
	expectedUniqueInt := "<int>42"
	expectedUniqueFloat := "<float64>3.14"
	expectedUniqueString := "<string>hello"
	expectedUniqueBool := "<bool>true"

	// Test for int
	if unique := Unique[int](intVal); unique != expectedUniqueInt {
		t.Errorf("Unique(intVal) = %v, want %v", unique, expectedUniqueInt)
	}

	// Test for float
	if unique := Unique[float64](floatVal); unique != expectedUniqueFloat {
		t.Errorf("Unique(floatVal) = %v, want %v", unique, expectedUniqueFloat)
	}

	// Test for string
	if unique := Unique[string](stringVal); unique != expectedUniqueString {
		t.Errorf("Unique(stringVal) = %v, want %v", unique, expectedUniqueString)
	}

	// Test for bool
	if unique := Unique[bool](boolVal); unique != expectedUniqueBool {
		t.Errorf("Unique(boolVal) = %v, want %v", unique, expectedUniqueBool)
	}
}

func TestUnique_Primitives_Pointers(t *testing.T) {
	// Define some test values
	intVal := 42
	intPtr := &intVal
	floatVal := 3.14
	floatPtr := &floatVal
	stringVal := "hello"
	stringPtr := &stringVal
	boolVal := true
	boolPtr := &boolVal
	uintptrVal := uintptr(42)
	uintptrPtr := &uintptrVal

	// Define expected results
	expectedUniqueIntPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(intPtr)))
	expectedUniqueFloatPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(floatPtr)))
	expectedUniqueStringPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(stringPtr)))
	expectedUniqueBoolPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(boolPtr)))
	expectedUniqueUintptrPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(uintptrPtr)))

	// Test for int
	t.Run("int", func(t *testing.T) {
		if unique := Unique[int](intPtr); unique != expectedUniqueIntPtr {
			t.Errorf("Unique(intPtr) = %v, want %v", unique, expectedUniqueIntPtr)
		}
	})

	// Test for float
	t.Run("float", func(t *testing.T) {
		if unique := Unique[float64](floatPtr); unique != expectedUniqueFloatPtr {
			t.Errorf("Unique(floatPtr) = %v, want %v", unique, expectedUniqueFloatPtr)
		}
	})

	// Test for string
	t.Run("string", func(t *testing.T) {
		if unique := Unique[string](stringPtr); unique != expectedUniqueStringPtr {
			t.Errorf("Unique(stringPtr) = %v, want %v", unique, expectedUniqueStringPtr)
		}
	})

	// Test for bool
	t.Run("bool", func(t *testing.T) {
		if unique := Unique[bool](boolPtr); unique != expectedUniqueBoolPtr {
			t.Errorf("Unique(boolPtr) = %v, want %v", unique, expectedUniqueBoolPtr)
		}
	})

	// Test for uintptr
	t.Run("uintptr", func(t *testing.T) {
		if unique := Unique[uintptr](uintptrPtr); unique != expectedUniqueUintptrPtr {
			t.Errorf("Unique(uintptrPtr) = %v, want %v", unique, expectedUniqueUintptrPtr)
		}
	})
}

func TestUnique_Structs(t *testing.T) {
	// Test for TestStruct
	structVal := TestSortable{Name: "Alice", Age: 30}
	expectedUniqueStruct := "<sortables.TestSortable>{Alice 30}"

	t.Run("TestStruct", func(t *testing.T) {
		if unique := Unique[TestSortable](structVal); unique != expectedUniqueStruct {
			t.Errorf("Unique(structVal) = %v, want %v", unique, expectedUniqueStruct)
		}
	})

	// Test for pointer to TestStruct
	t.Run("pointer to TestStruct", func(t *testing.T) {
		structPtr := &TestSortable{Name: "Bob", Age: 40}
		expectedUniqueStructPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(structPtr)))
		if unique := Unique[TestSortable](structPtr); unique != expectedUniqueStructPtr {
			t.Errorf("Unique(structPtr) = %v, want %v", unique, expectedUniqueStructPtr)
		}
	})
}

func TestUnique_EdgeCases_ComplexTypes(t *testing.T) {
	// Define some test values
	var ch = make(chan int)
	var chNil chan int
	var chPtr = &ch
	var chNilPtr *chan int

	var mapVal = make(map[int]int)
	var mapNil map[int]int
	var mapPtr = &mapVal
	var mapNilPtr *map[int]int

	var sliceVal = make([]int, 0)
	var sliceNil []int
	var slicePtr = &sliceVal
	var sliceNilPtr *[]int

	var expectedUniqueCh = fmt.Sprintf("<%T>%v", ch, ch)
	var expectedUniqueChNil = fmt.Sprintf("<%T>%v", chNil, chNil)
	var expectedUniqueChPtr = "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(chPtr)))
	var expectedUniqueChNilPtr = "0x0"

	var expectedUniqueMap = fmt.Sprintf("<%T>%v", mapVal, mapVal)
	var expectedUniqueMapNil = fmt.Sprintf("<%T>%v", mapNil, mapNil)
	var expectedUniqueMapPtr = "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(mapPtr)))
	var expectedUniqueMapNilPtr = "0x0"

	var expectedUniqueSlice = fmt.Sprintf("<%T>%v", sliceVal, sliceVal)
	var expectedUniqueSliceNil = fmt.Sprintf("<%T>%v", sliceNil, sliceNil)
	var expectedUniqueSlicePtr = "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(slicePtr)))
	var expectedUniqueSliceNilPtr = "0x0"

	t.Run("channel", func(t *testing.T) {
		unique := Unique[chan int](ch)
		if unique != expectedUniqueCh {
			t.Errorf("Unique(ch) = %v, want %v", unique, expectedUniqueCh)
		}
	})

	t.Run("nil channel", func(t *testing.T) {
		unique := Unique[chan int](chNil)
		if unique != expectedUniqueChNil {
			t.Errorf("Unique(chNil) = %v, want %v", unique, expectedUniqueChNil)
		}
	})

	t.Run("pointer to channel", func(t *testing.T) {
		unique := Unique[chan int](chPtr)
		if unique != expectedUniqueChPtr {
			t.Errorf("Unique(chPtr) = %v, want %v", unique, expectedUniqueChPtr)
		}
	})

	t.Run("nil pointer to channel", func(t *testing.T) {
		unique := Unique[chan int](chNilPtr)
		if unique != expectedUniqueChNilPtr {
			t.Errorf("Unique(chNilPtr) = %v, want %v", unique, expectedUniqueChNilPtr)
		}
	})

	t.Run("map", func(t *testing.T) {
		unique := Unique[map[int]int](mapVal)
		if unique != expectedUniqueMap {
			t.Errorf("Unique(mapVal) = %v, want %v", unique, expectedUniqueMap)
		}
	})

	t.Run("pointer to map", func(t *testing.T) {
		unique := Unique[map[int]int](mapPtr)
		if unique != expectedUniqueMapPtr {
			t.Errorf("Unique(mapPtr) = %v, want %v", unique, expectedUniqueMapPtr)
		}
	})

	t.Run("nil map", func(t *testing.T) {
		unique := Unique[map[int]int](mapNil)
		if unique != expectedUniqueMapNil {
			t.Errorf("Unique(mapNil) = %v, want %v", unique, expectedUniqueMapNil)
		}
	})

	t.Run("nil pointer to map", func(t *testing.T) {
		unique := Unique[map[int]int](mapNilPtr)
		if unique != expectedUniqueMapNilPtr {
			t.Errorf("Unique(mapNilPtr) = %v, want %v", unique, expectedUniqueMapNilPtr)
		}
	})

	t.Run("slice", func(t *testing.T) {
		unique := Unique[[]int](sliceVal)
		if unique != expectedUniqueSlice {
			t.Errorf("Unique(sliceVal) = %v, want %v", unique, expectedUniqueSlice)
		}
	})

	t.Run("pointer to slice", func(t *testing.T) {
		unique := Unique[[]int](slicePtr)
		if unique != expectedUniqueSlicePtr {
			t.Errorf("Unique(slicePtr) = %v, want %v", unique, expectedUniqueSlicePtr)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		unique := Unique[[]int](sliceNil)
		if unique != expectedUniqueSliceNil {
			t.Errorf("Unique(sliceNil) = %v, want %v", unique, expectedUniqueSliceNil)
		}
	})

	t.Run("nil pointer to slice", func(t *testing.T) {
		unique := Unique[[]int](sliceNilPtr)
		if unique != expectedUniqueSliceNilPtr {
			t.Errorf("Unique(sliceNilPtr) = %v, want %v", unique, expectedUniqueSliceNilPtr)
		}
	})
}

// TestUnique_EdgeCases_Interfaces tests the Unique function with interface edge cases.
// The expected behaviour is the method to be able to correctly identify if the underlying
// implementation is a pointer or not.
func TestUnique_EdgeCases_Interfaces(t *testing.T) {
	// Test for a nil interface
	t.Run("nil interface", func(t *testing.T) {
		var i testInterface
		unique := Unique[testInterface](i)
		t.Logf("Unique(i) = %v", unique)
	})

	// Test for a non-nil value of implementation 1
	t.Run("value implementation 1", func(t *testing.T) {
		var impl1Val testInterface = TestInterfaceImpl1{"impl value", 1}
		unique := Unique[testInterface](impl1Val)
		t.Logf("Unique(impl1Val) = %v", unique)
	})

	// Test for a nil value of implementation 1
	t.Run("nil value implementation 1", func(t *testing.T) {
		var impl1NilValue TestInterfaceImpl1
		unique := Unique[testInterface](impl1NilValue)
		t.Logf("Unique(impl1Val) = %v", unique)
	})

	// Test for a pointer to implementation 1
	t.Run("pointer implementation 1", func(t *testing.T) {
		var implPtr testInterface = &TestInterfaceImpl1{"impl pointer", 1}
		unique := Unique[testInterface](implPtr)
		t.Logf("Unique(impl1NilPtr) = %v", unique)
	})

	// Test for a nil pointer to implementation 1
	t.Run("nil pointer implementation 1", func(t *testing.T) {
		var impl1NilPtr *TestInterfaceImpl1
		unique := Unique[testInterface](impl1NilPtr)
		t.Logf("Unique(impl1NilPtr) = %v", unique)
	})

	// Test for a pointer to nil implementation 1
	t.Run("nil pointer implementation 1", func(t *testing.T) {
		var impl TestInterfaceImpl1
		var impl1NilPtr = &impl
		unique := Unique[testInterface](impl1NilPtr)
		t.Logf("Unique(impl1NilPtr) = %v", unique)
	})
}

func TestEquality(t *testing.T) {
	type S = testseed.SortableValue
	seed, dup := random.Struct[S](20).Duplicate()

	t.Run("should be equal", func(t *testing.T) {
		seed.EachN(func(i int, v S) {
			if !Equality(v, dup.At(i)) {
				t.Errorf("EqualsOf(%v, %v) = false, want true", v, dup.At(i))
			}
		})
	})

	t.Run("should not be equal", func(t *testing.T) {
		seed.EachN(func(i int, v S) {
			if i+1 < seed.Len() {
				if Equality(v, seed.At(i+1)) {
					t.Errorf("EqualsOf(%v, %v) = true, want false", v, seed.At(i+1))
				}
			}
		})
	})
}

func TestUnsafeEquality(t *testing.T) {
	t.Run("equal pointers", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		pointer := &data
		pointer2 := &data
		if !TryEquality[TestInterfaceImpl1](pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	t.Run("unequal pointers", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		pointer := &data
		pointer2 := &TestInterfaceImpl1{"b", 2}
		if TryEquality[TestInterfaceImpl1](pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = true, want false")
		}
	})

	t.Run("equal comparable values", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		value := data
		value2 := data
		if !TryEquality[TestInterfaceImpl1](value, value2) {
			t.Errorf("Equality(value, value2) = false, want true")
		}
	})

	t.Run("unequal comparable values", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		value := data
		value2 := TestInterfaceImpl1{"b", 2}
		if TryEquality[TestInterfaceImpl1](value, value2) {
			t.Errorf("Equality(value, value2) = true, want false")
		}
	})

	// The pointers here have a valid address to the value.
	// this case it recognizes them as the same. which is different
	// from uninitialized pointers, whose address is 0x0.
	t.Run("equal pointers to nil", func(t *testing.T) {
		var data TestInterfaceImpl1
		pointer := &data
		pointer2 := &data
		if !TryEquality[TestInterfaceImpl1](pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	t.Run("unequal nil pointers to nil", func(t *testing.T) {
		var data1 TestInterfaceImpl1
		var data2 TestInterfaceImpl1
		var pointer = &data1
		var pointer2 = &data2
		if TryEquality[TestInterfaceImpl1](pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	// Here there is not much to do. uninitialized pointers have an address of 0x0
	// and any equality check for them will return true.
	// Introducing complexity to differentiate between the stack addresses of nil
	// pointers can lead to confusion and bugs. Since nil pointers indicate the
	// absence of a value, treating them as equal makes logical sense in most contexts.
	t.Run("uninitialized pointers", func(t *testing.T) {
		var data *TestInterfaceImpl1
		var data2 *TestInterfaceImpl1
		if !TryEquality[TestInterfaceImpl1](data, data2) {
			t.Errorf("Equality(data, data2) = false, want true")
		}
	})

	t.Run("uninitialized pointers", func(t *testing.T) {
		var data *TestInterfaceImpl1
		if !TryEquality[TestInterfaceImpl1](data, data) {
			t.Errorf("Equality(data, data2) = false, want true")
		}
	})

	// Zero Values (Non-Pointer Structs): Go automatically initializes each field of
	// the struct to its zero value. This means the struct is not nil; it’s just empty,
	// containing default values (e.g., empty strings for string fields, zeroes for int
	// fields, etc.). Importantly, it’s fully instantiated in memory, so Go can call
	// its methods even though it holds no actual data.
	t.Run("equal nil values", func(t *testing.T) {
		var value TestInterfaceImpl1
		var value2 TestInterfaceImpl1
		if !TryEquality[TestInterfaceImpl1](value, value2) {
			t.Errorf("Equality(value, value2) = false, want true")
		}
	})

	t.Run("equal implementation pointers as interface 1", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		pointer := &data
		pointer2 := &data
		if !TryEquality[testInterface](pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	// When we declare pointer and pointer2 as testInterface, Go sees them as interface
	// values pointing to the same instance of TestInterfaceImpl1, so TryEquality[testInterface]
	// correctly identifies that they refer to the same data.
	t.Run("equal implementation pointers as interface", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		var pointer testInterface = &data
		var pointer2 testInterface = &data

		// Interface values carry both type and value: An interface value in Go includes both
		// the concrete type and the value it holds. Passing pointer and pointer2 as
		// testInterface, Go performs a comparison on the underlying data for interface{}
		// types. So TryEquality[testInterface] succeeds because both pointer and pointer2
		// point to the same TestInterfaceImpl1 instance.
		if !TryEquality[testInterface](pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}

		// In TryEquality[TestInterfaceImpl1](pointer, pointer2),
		// Go expects both pointer and pointer2 to be directly of type TestInterfaceImpl1.
		// However, they’re actually stored as interface{} types wrapping *TestInterfaceImpl1.
		// The line if _, ok := x.(T); ok fails because x and y are interfaces pointing to
		// TestInterfaceImpl1, not TestInterfaceImpl1 directly.
		// The method has a workaround for that trying to make a dereference conversion.
		if !TryEquality[TestInterfaceImpl1](pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	t.Run("equal comparableOf implementation", func(t *testing.T) {
		a := testseed.NewComparableValue("a", 1)
		b := testseed.NewComparableValue("a", 1)
		if !TryEquality[gtools.ComparableOf](a, b) {
			t.Errorf("Equality(%v, %v) = false, want true", a, b)
		}
	})

	t.Run("unequal comparableOf implementation", func(t *testing.T) {
		a := testseed.NewComparableValue("a", 1)
		b := testseed.NewComparableValue("b", 2)
		if TryEquality[gtools.ComparableOf](a, b) {
			t.Errorf("Equality(%v, %v) = true, want false", a, b)
		}
	})
}
