// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/testlite/testseeds/static/structs/models"
)

func TestComparatorSortableOf_Compare(t *testing.T) {
	comp := ComparatorOf[*models.SortableNode]()
	a := models.NewSortableNode("A")
	b := models.NewSortableNode("B")
	c := models.NewSortableNode("A")

	if comp.Compare(a, b) != -1 {
		t.Errorf("Compare(A, B) = %d, want 1", comp.Compare(a, b))
	}
	if comp.Compare(b, a) != 1 {
		t.Errorf("Compare(B, A) = %d, want -1", comp.Compare(b, a))
	}
	if comp.Compare(a, c) != 0 {
		t.Errorf("Compare(A, C) = %d, want 0", comp.Compare(a, c))
	}
}

func TestComparatorSortableOf_Equals(t *testing.T) {
	comp := ComparatorOf[*models.SortableNode]()
	a := models.NewSortableNode("A")
	b := models.NewSortableNode("A")
	c := models.NewSortableNode("B")

	if !comp.Equals(a, b) {
		t.Errorf("Equals(A, B) = false, want true")
	}
	if comp.Equals(a, c) {
		t.Errorf("Equals(A, C) = true, want false")
	}
}

func TestFingerprint_Primitives_Values(t *testing.T) {
	// Define some test values
	intVal := 42
	floatVal := 3.14
	stringVal := "hello"
	boolVal := true

	// Define expected results
	expectedFingerprintInt := "<int>42"
	expectedFingerprintFloat := "<float64>3.14"
	expectedFingerprintString := "<string>hello"
	expectedFingerprintBool := "<bool>true"

	// Test for int
	if unique := Fingerprint[int](intVal); unique != expectedFingerprintInt {
		t.Errorf("Fingerprint(intVal) = %v, want %v", unique, expectedFingerprintInt)
	}

	// Test for float
	if unique := Fingerprint[float64](floatVal); unique != expectedFingerprintFloat {
		t.Errorf("Fingerprint(floatVal) = %v, want %v", unique, expectedFingerprintFloat)
	}

	// Test for string
	if unique := Fingerprint[string](stringVal); unique != expectedFingerprintString {
		t.Errorf("Fingerprint(stringVal) = %v, want %v", unique, expectedFingerprintString)
	}

	// Test for bool
	if unique := Fingerprint[bool](boolVal); unique != expectedFingerprintBool {
		t.Errorf("Fingerprint(boolVal) = %v, want %v", unique, expectedFingerprintBool)
	}
}

func TestFingerprint_Primitives_Pointers(t *testing.T) {
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
	expectedFingerprintIntPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(intPtr)))
	expectedFingerprintFloatPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(floatPtr)))
	expectedFingerprintStringPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(stringPtr)))
	expectedFingerprintBoolPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(boolPtr)))
	expectedFingerprintUintptrPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(uintptrPtr)))

	// Test for int
	t.Run("int", func(t *testing.T) {
		if unique := Fingerprint[int](intPtr); unique != expectedFingerprintIntPtr {
			t.Errorf("Fingerprint(intPtr) = %v, want %v", unique, expectedFingerprintIntPtr)
		}
	})

	// Test for float
	t.Run("float", func(t *testing.T) {
		if unique := Fingerprint[float64](floatPtr); unique != expectedFingerprintFloatPtr {
			t.Errorf("Fingerprint(floatPtr) = %v, want %v", unique, expectedFingerprintFloatPtr)
		}
	})

	// Test for string
	t.Run("string", func(t *testing.T) {
		if unique := Fingerprint[string](stringPtr); unique != expectedFingerprintStringPtr {
			t.Errorf("Fingerprint(stringPtr) = %v, want %v", unique, expectedFingerprintStringPtr)
		}
	})

	// Test for bool
	t.Run("bool", func(t *testing.T) {
		if unique := Fingerprint[bool](boolPtr); unique != expectedFingerprintBoolPtr {
			t.Errorf("Fingerprint(boolPtr) = %v, want %v", unique, expectedFingerprintBoolPtr)
		}
	})

	// Test for uintptr
	t.Run("uintptr", func(t *testing.T) {
		if unique := Fingerprint[uintptr](uintptrPtr); unique != expectedFingerprintUintptrPtr {
			t.Errorf("Fingerprint(uintptrPtr) = %v, want %v", unique, expectedFingerprintUintptrPtr)
		}
	})
}

func TestFingerprint_Structs(t *testing.T) {
	// Test for TestStruct
	structVal := TestSortable{Name: "Alice", Age: 30}
	expectedFingerprintStruct := "<sortables.TestSortable>{Alice 30}"

	t.Run("TestStruct", func(t *testing.T) {
		if unique := Fingerprint[TestSortable](structVal); unique != expectedFingerprintStruct {
			t.Errorf("Fingerprint(structVal) = %v, want %v", unique, expectedFingerprintStruct)
		}
	})

	// Test for pointer to TestStruct
	t.Run("pointer to TestStruct", func(t *testing.T) {
		structPtr := &TestSortable{Name: "Bob", Age: 40}
		expectedFingerprintStructPtr := "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(structPtr)))
		if unique := Fingerprint[TestSortable](structPtr); unique != expectedFingerprintStructPtr {
			t.Errorf("Fingerprint(structPtr) = %v, want %v", unique, expectedFingerprintStructPtr)
		}
	})
}

func TestFingerprint_EdgeCases_ComplexTypes(t *testing.T) {
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

	var expectedFingerprintCh = fmt.Sprintf("<%T>%v", ch, ch)
	var expectedFingerprintChNil = fmt.Sprintf("<%T>%v", chNil, chNil)
	var expectedFingerprintChPtr = "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(chPtr)))
	var expectedFingerprintChNilPtr = "0x0"

	var expectedFingerprintMap = fmt.Sprintf("<%T>%v", mapVal, mapVal)
	var expectedFingerprintMapNil = fmt.Sprintf("<%T>%v", mapNil, mapNil)
	var expectedFingerprintMapPtr = "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(mapPtr)))
	var expectedFingerprintMapNilPtr = "0x0"

	var expectedFingerprintSlice = fmt.Sprintf("<%T>%v", sliceVal, sliceVal)
	var expectedFingerprintSliceNil = fmt.Sprintf("<%T>%v", sliceNil, sliceNil)
	var expectedFingerprintSlicePtr = "0x" + fmt.Sprintf("%x", uintptr(unsafe.Pointer(slicePtr)))
	var expectedFingerprintSliceNilPtr = "0x0"

	t.Run("channel", func(t *testing.T) {
		unique := Fingerprint[chan int](ch)
		if unique != expectedFingerprintCh {
			t.Errorf("Fingerprint(ch) = %v, want %v", unique, expectedFingerprintCh)
		}
	})

	t.Run("nil channel", func(t *testing.T) {
		unique := Fingerprint[chan int](chNil)
		if unique != expectedFingerprintChNil {
			t.Errorf("Fingerprint(chNil) = %v, want %v", unique, expectedFingerprintChNil)
		}
	})

	t.Run("pointer to channel", func(t *testing.T) {
		unique := Fingerprint[chan int](chPtr)
		if unique != expectedFingerprintChPtr {
			t.Errorf("Fingerprint(chPtr) = %v, want %v", unique, expectedFingerprintChPtr)
		}
	})

	t.Run("nil pointer to channel", func(t *testing.T) {
		unique := Fingerprint[chan int](chNilPtr)
		if unique != expectedFingerprintChNilPtr {
			t.Errorf("Fingerprint(chNilPtr) = %v, want %v", unique, expectedFingerprintChNilPtr)
		}
	})

	t.Run("map", func(t *testing.T) {
		unique := Fingerprint[map[int]int](mapVal)
		if unique != expectedFingerprintMap {
			t.Errorf("Fingerprint(mapVal) = %v, want %v", unique, expectedFingerprintMap)
		}
	})

	t.Run("pointer to map", func(t *testing.T) {
		unique := Fingerprint[map[int]int](mapPtr)
		if unique != expectedFingerprintMapPtr {
			t.Errorf("Fingerprint(mapPtr) = %v, want %v", unique, expectedFingerprintMapPtr)
		}
	})

	t.Run("nil map", func(t *testing.T) {
		unique := Fingerprint[map[int]int](mapNil)
		if unique != expectedFingerprintMapNil {
			t.Errorf("Fingerprint(mapNil) = %v, want %v", unique, expectedFingerprintMapNil)
		}
	})

	t.Run("nil pointer to map", func(t *testing.T) {
		unique := Fingerprint[map[int]int](mapNilPtr)
		if unique != expectedFingerprintMapNilPtr {
			t.Errorf("Fingerprint(mapNilPtr) = %v, want %v", unique, expectedFingerprintMapNilPtr)
		}
	})

	t.Run("slice", func(t *testing.T) {
		unique := Fingerprint[[]int](sliceVal)
		if unique != expectedFingerprintSlice {
			t.Errorf("Fingerprint(sliceVal) = %v, want %v", unique, expectedFingerprintSlice)
		}
	})

	t.Run("pointer to slice", func(t *testing.T) {
		unique := Fingerprint[[]int](slicePtr)
		if unique != expectedFingerprintSlicePtr {
			t.Errorf("Fingerprint(slicePtr) = %v, want %v", unique, expectedFingerprintSlicePtr)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		unique := Fingerprint[[]int](sliceNil)
		if unique != expectedFingerprintSliceNil {
			t.Errorf("Fingerprint(sliceNil) = %v, want %v", unique, expectedFingerprintSliceNil)
		}
	})

	t.Run("nil pointer to slice", func(t *testing.T) {
		unique := Fingerprint[[]int](sliceNilPtr)
		if unique != expectedFingerprintSliceNilPtr {
			t.Errorf("Fingerprint(sliceNilPtr) = %v, want %v", unique, expectedFingerprintSliceNilPtr)
		}
	})
}

// TestFingerprint_EdgeCases_Interfaces tests the Fingerprint function with interface edge cases.
// The expected behaviour is the method to be able to correctly identify if the underlying
// implementation is a pointer or not.
func TestFingerprint_EdgeCases_Interfaces(t *testing.T) {
	// Test for a nil interface
	t.Run("nil interface", func(t *testing.T) {
		var i testInterface
		unique := Fingerprint[testInterface](i)
		t.Logf("Fingerprint(i) = %v", unique)
	})

	// Test for a non-nil value of implementation 1
	t.Run("value implementation 1", func(t *testing.T) {
		var impl1Val testInterface = TestInterfaceImpl1{"impl value", 1}
		unique := Fingerprint[testInterface](impl1Val)
		t.Logf("Fingerprint(impl1Val) = %v", unique)
	})

	// Test for a nil value of implementation 1
	t.Run("nil value implementation 1", func(t *testing.T) {
		var impl1NilValue TestInterfaceImpl1
		unique := Fingerprint[testInterface](impl1NilValue)
		t.Logf("Fingerprint(impl1Val) = %v", unique)
	})

	// Test for a pointer to implementation 1
	t.Run("pointer implementation 1", func(t *testing.T) {
		var implPtr testInterface = &TestInterfaceImpl1{"impl pointer", 1}
		unique := Fingerprint[testInterface](implPtr)
		t.Logf("Fingerprint(impl1NilPtr) = %v", unique)
	})

	// Test for a nil pointer to implementation 1
	t.Run("nil pointer implementation 1", func(t *testing.T) {
		var impl1NilPtr *TestInterfaceImpl1
		unique := Fingerprint[testInterface](impl1NilPtr)
		t.Logf("Fingerprint(impl1NilPtr) = %v", unique)
	})

	// Test for a pointer to nil implementation 1
	t.Run("nil pointer implementation 1", func(t *testing.T) {
		var impl TestInterfaceImpl1
		var impl1NilPtr = &impl
		unique := Fingerprint[testInterface](impl1NilPtr)
		t.Logf("Fingerprint(impl1NilPtr) = %v", unique)
	})
}

func TestEquality(t *testing.T) {
	type S = models.SortableValue
	seed, dup := random.Struct[S](20).Duplicate()

	t.Run("should be compare", func(t *testing.T) {
		seed.EachN(func(i int, v S) {
			if !Equality(v, dup.At(i)) {
				t.Errorf("EqualsOf(%v, %v) = false, want true", v, dup.At(i))
			}
		})
	})

	t.Run("should not be compare", func(t *testing.T) {
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
	t.Run("compare ptrs", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		pointer := &data
		pointer2 := &data
		if !TryEquality(pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	t.Run("unequal ptrs", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		pointer := &data
		pointer2 := &TestInterfaceImpl1{"b", 2}
		if TryEquality(pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = true, want false")
		}
	})

	t.Run("compare comparable values", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		value := data
		value2 := data
		if !TryEquality(value, value2) {
			t.Errorf("Equality(value, value2) = false, want true")
		}
	})

	t.Run("unequal comparable values", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		value := data
		value2 := TestInterfaceImpl1{"b", 2}
		if TryEquality(value, value2) {
			t.Errorf("Equality(value, value2) = true, want false")
		}
	})

	// The ptrs here have a valid address to the value.
	// this case it recognizes them as the same. which is different
	// from uninitialized ptrs, whose address is 0x0.
	t.Run("compare ptrs to nil", func(t *testing.T) {
		var data TestInterfaceImpl1
		pointer := &data
		pointer2 := &data
		if !TryEquality(pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	t.Run("unequal nil ptrs to nil", func(t *testing.T) {
		var data1 TestInterfaceImpl1
		var data2 TestInterfaceImpl1
		var pointer = &data1
		var pointer2 = &data2
		if TryEquality(pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	// Here there is not much to do. uninitialized ptrs have an address of 0x0
	// and any equality check for them will return true.
	// Introducing complexity to differentiate between the stack address of nil
	// ptrs can lead to confusion and bugs. Since nil ptrs indicate the
	// absence of a value, treating them as compare makes logical sense in most contexts.
	t.Run("uninitialized ptrs", func(t *testing.T) {
		var data *TestInterfaceImpl1
		var data2 *TestInterfaceImpl1
		if !TryEquality(data, data2) {
			t.Errorf("Equality(data, data2) = false, want true")
		}
	})

	t.Run("uninitialized ptrs", func(t *testing.T) {
		var data *TestInterfaceImpl1
		if !TryEquality(data, data) {
			t.Errorf("Equality(data, data2) = false, want true")
		}
	})

	// Zero Values (Non-Pointer Structs): Go automatically initializes each field of
	// the struct to its zero value. This means the struct is not nil; it’s just empty,
	// containing default values (e.g., empty strings for string fields, zeroes for int
	// fields, etc.). Importantly, it’s fully instantiated in memory, so Go can call
	// its methods even though it holds no actual data.
	t.Run("compare nil values", func(t *testing.T) {
		var value TestInterfaceImpl1
		var value2 TestInterfaceImpl1
		if !TryEquality(value, value2) {
			t.Errorf("Equality(value, value2) = false, want true")
		}
	})

	t.Run("compare implementation ptrs as interface 1", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		pointer := &data
		pointer2 := &data
		if !TryEquality(pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	// At we declare pointer and pointer2 as testInterface, Go sees them as interface
	// values pointing to the same instance of TestInterfaceImpl1, so TryEqualitycorrectly identifies that they refer to the same data.
	t.Run("compare implementation ptrs as interface", func(t *testing.T) {
		data := TestInterfaceImpl1{"a", 1}
		var pointer testInterface = &data
		var pointer2 testInterface = &data

		// Interface values carry both type and value: An interface value in Go includes both
		// the concrete type and the value it holds. Passing pointer and pointer2 as
		// testInterface, Go performs a comparison on the underlying data for interface{}
		// types. So TryEqualityeeds because both pointer and pointer2
		// point to the same TestInterfaceImpl1 instance.
		if !TryEquality(pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}

		// In TryEquality(pointer, pointer2),
		// Go expects both pointer and pointer2 to be directly of type TestInterfaceImpl1.
		// However, they’re actually stored as interface{} types wrapping *TestInterfaceImpl1.
		// The line if _, ok := x.(T); ok fails because x and y are interfaces pointing to
		// TestInterfaceImpl1, not TestInterfaceImpl1 directly.
		// The method has a workaround for that trying to make a dereference conversion.
		if !TryEquality(pointer, pointer2) {
			t.Errorf("Equality(pointer, pointer2) = false, want true")
		}
	})

	t.Run("compare comparableOf implementation", func(t *testing.T) {
		a := models.NewComparableValue("a", 1)
		b := models.NewComparableValue("a", 1)
		if !TryEquality(a, b) {
			t.Errorf("Equality(%v, %v) = false, want true", a, b)
		}
	})

	t.Run("unequal comparableOf implementation", func(t *testing.T) {
		a := models.NewComparableValue("a", 1)
		b := models.NewComparableValue("b", 2)
		if TryEquality(a, b) {
			t.Errorf("Equality(%v, %v) = true, want false", a, b)
		}
	})
}
