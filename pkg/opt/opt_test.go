// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package opt

import (
	"testing"
)

func TestOfNullable(t *testing.T) {

	settedPrimitive := 10
	if !OfNullable(&settedPrimitive).IsPresent() {
		t.Errorf("The value offers its own memory address as a pointer")
	}

	var unsettedPrimitive int // It's a zero value
	if !OfNullable(&unsettedPrimitive).IsPresent() {
		t.Errorf("A primitive variable holds a zero-value")
	}

	var unsettedPrimitivePointer *int // It's a nil pointer, since the address isn't set
	if OfNullable(unsettedPrimitivePointer).IsPresent() {
		t.Errorf("A pointer to a primitive address has no address as value")
	}

	// The optional was designed to reject non-nil pointers that point to nil pointers.
	//
	//	val := reflect.ValueOf(value).Elem()
	//	if val.Kind() == reflect.Ptr && val.IsNil() { return None[T]() }
	//
	// `unsettedPrimitivePointer` is a nil pointer to an int.
	// When passing the address of `unsettedPrimitivePointer`, it's of type **int.
	// Although the address itself isn't nil, the underlying value of the pointer it points to is nil.
	// This means the `OfNullable` function correctly identifies that the underlying value is nil and returns `None`.
	// Therefore, `IsPresent()` should return false.
	if OfNullable(&unsettedPrimitivePointer).IsPresent() {
		t.Errorf("The underlying value which the pointer points to is nil")
	}

	if !OfNullable(&struct{ value int }{10}).IsPresent() {
		t.Errorf("The value offers its own memory address as a pointer")
	}

	settedStruct := struct{ value int }{10} // It's a value
	if !OfNullable(&settedStruct).IsPresent() {
		t.Errorf("The value offers its own memory address as a pointer")
	}

	settedStructAddress := &struct{ value int }{10} // Is a pointer to an anonymous struct with a field value initialized to 10.

	if !OfNullable(settedStructAddress).IsPresent() {
		t.Errorf("The pointer holds a non nil value")
	}
	// This is just redundancy
	if !OfNullable(&settedStructAddress).IsPresent() {
		t.Errorf("The pointer holds a non nil value")
	}

	//
	var unsettedStruct struct{ value int } // It's a zero value
	if !OfNullable(&unsettedStruct).IsPresent() {
		t.Error("This is a value with its own address as zero")
	}
	// Later i finish that
	//
	//var unsettedStructPointer *struct {
	//	value int
	//}
	//unsettedStructPointerVar := OfNullable(unsettedStructPointer).IsPresent()
	//unsettedStrunctPointerAsAddress := OfNullable(&unsettedStructPointer).IsPresent()
	//// It must panic because unsettedStructPointer is nil, and it is trying to dereference a nil pointer
	//assert.Panics(t, func() { OfNullable(*unsettedStructPointer).IsPresent() })
	//
	//rawArray := OfNullable([]int{1, 2, 3}).IsPresent()
	//rawEmptyArray := OfNullable([]int{}).IsPresent()
	//
	//settedArray := []int{1, 2, 3}
	//settedArrayVar := OfNullable(settedArray).IsPresent()
	//
	//var unsettedArray []int
	//unsettedArrayVar := OfNullable(unsettedArray).IsPresent()
	//unsettedArrayAddress := OfNullable(&unsettedArray).IsPresent()
	//
	//var unsettedArrayPointer *[]int
	//unsettedArrayPointerVar := OfNullable(unsettedArrayPointer).IsPresent()
	//unsettedArrayPointerAsAddress := OfNullable(&unsettedArrayPointer).IsPresent()
	//assert.Panics(t, func() { OfNullable(*unsettedArrayPointer).IsPresent() })
	//
	//var unsettedArrayOfPointers []*int
	//unsettedArrayOfPointersVar := OfNullable(unsettedArrayOfPointers).IsPresent()
	//unsettedArrayOfPointersAsAddress := OfNullable(&unsettedArrayOfPointers).IsPresent()
	//
	//var unsettedArrayPointerOfPointers *[]*int
	//unsettedArrayPointerOfPointersVar := OfNullable(unsettedArrayPointerOfPointers).IsPresent()
	//unsettedArrayPointerOfPointersAsAddress := OfNullable(&unsettedArrayPointerOfPointers).IsPresent()
	//assert.Panics(t, func() { OfNullable(*unsettedArrayPointerOfPointers).IsPresent() })
	//
	//rawMap := OfNullable(map[string]int{"key": 1}).IsPresent()
	//rawEmptyMap := OfNullable(map[string]int{}).IsPresent()
	//
	//settedMap := map[string]int{"key": 1}
	//settedMapVar := OfNullable(settedMap).IsPresent()
	//settedMapVarAsAddress := OfNullable(&settedMap).IsPresent()
	//
	//var settedMapPointer *map[string]int
	//settedMapPointerVar := OfNullable(settedMapPointer).IsPresent()
	//settedMapPointerAsAddress := OfNullable(&settedMapPointer).IsPresent()
	//assert.Panics(t, func() { OfNullable(*settedMapPointer).IsPresent() })
	//
	//var settedMapOfPointers map[string]*int
	//settedMapOfPointersVar := OfNullable(settedMapOfPointers).IsPresent()
	//settedMapOfPointersAsAddress := OfNullable(&settedMapOfPointers).IsPresent()
	//
	//var settedMapPointerOfPointers *map[string]*int
	//settedMapPointerOfPointersVar := OfNullable(settedMapPointerOfPointers).IsPresent()
	//settedMapPointerOfPointersAsAddress := OfNullable(&settedMapPointerOfPointers).IsPresent()
	//assert.Panics(t, func() { OfNullable(*settedMapPointerOfPointers).IsPresent() })
	//
	//var unsettedMap map[string]int
	//unsettedMapVar := OfNullable(unsettedMap).IsPresent()
	//unsettedMapAddress := OfNullable(&unsettedMap).IsPresent()
	//
	//var unsettedMapPointer *map[string]int
	//unsettedMapPointerVar := OfNullable(unsettedMapPointer).IsPresent()
	//unsettedMapPointerAsAddress := OfNullable(&unsettedMapPointer).IsPresent()
	//assert.Panics(t, func() { OfNullable(*unsettedMapPointer).IsPresent() })
}

func TestOrAssert_ReturnExistingValue(t *testing.T) {
	existingValue := 100
	opt := Of(existingValue)

	newValue := 200
	result := opt.OrAssert(&newValue)

	if !result.isSet {
		t.Errorf("Expected isSet to be true, but got false")
	}

	if *result.value != existingValue {
		t.Errorf("Expected value %v, but got %v", existingValue, *result.value)
	}
}

// TestOrAssert_PanicOnNilPointer tests that OrAssert panics when a nil pointer is provided.
func TestOrAssert_PanicOnNilPointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when passing a nil pointer, but got none")
		}
	}()

	var nilValue *int
	opt := None[int]()
	opt.OrAssert(nilValue)
}

// TestOrAssert_PanicOnNilInnerPointer tests that OrAssert panics when a pointer to a nil value is provided.
func TestOrAssert_PanicOnNilInnerPointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when passing a pointer to a nil value, but got none")
		}
	}()

	var nilPointer *int = nil
	opt := None[*int]()
	opt.OrAssert(&nilPointer)
}

// TestOrAssert_ReturnNewValue tests that OrAssert returns a new Option when the original Option is not set.
func TestOrAssert_ReturnNewValue(t *testing.T) {
	newValue := 300
	opt := None[int]()

	result := opt.OrAssert(&newValue)

	if !result.isSet {
		t.Errorf("Expected isSet to be true, but got false")
	}

	if *result.value != newValue {
		t.Errorf("Expected value %v, but got %v", newValue, *result.value)
	}
}

// TestOrAssert_HandleComplexType tests OrAssert with a more complex type.
func TestOrAssert_HandleComplexType(t *testing.T) {
	type ComplexStruct struct {
		Name  string
		Value int
	}

	cs := ComplexStruct{Name: "Test", Value: 42}
	opt := None[ComplexStruct]()

	result := opt.OrAssert(&cs)

	if !result.isSet {
		t.Errorf("Expected isSet to be true, but got false")
	}

	if result.value.Name != "Test" || result.value.Value != 42 {
		t.Errorf("Expected ComplexStruct {Name: 'Test', Value: 42}, but got {Name: %v, Value: %v}", result.value.Name, result.value.Value)
	}
}

func TestReflectHardNullable(t *testing.T) {
	// Test case: nil slice
	var nilSlice []string
	if !isDeepReflectedNullable(&nilSlice) {
		t.Errorf("Expected true for nil slice, got false")
	}

	// Test case: non-nil slice
	nonNilSlice := []string{"test"}
	if isDeepReflectedNullable(&nonNilSlice) {
		t.Errorf("Expected false for non-nil slice, got true")
	}

	// Test case: nil map
	var nilMap map[string]string
	if !isDeepReflectedNullable(&nilMap) {
		t.Errorf("Expected true for nil map, got false")
	}

	// Test case: non-nil map
	nonNilMap := map[string]string{"key": "value"}
	if isDeepReflectedNullable(&nonNilMap) {
		t.Errorf("Expected false for non-nil map, got true")
	}

	// Test case: nil pointer
	var nilPtr *int
	if !isDeepReflectedNullable(nilPtr) {
		t.Errorf("Expected true for nil pointer, got false")
	}

	// Test case: non-nil pointer
	value := 10
	nonNilPtr := &value
	if isDeepReflectedNullable(nonNilPtr) {
		t.Errorf("Expected false for non-nil pointer, got true")
	}

	// Test case: struct (non-nil, always)
	type TestStruct struct {
		Field1 string
		Field2 int
	}
	structVal := TestStruct{Field1: "test", Field2: 42}
	if isDeepReflectedNullable(&structVal) {
		t.Errorf("Expected false for non-nil struct, got true")
	}

	// Test case: nil interface
	var nilInterface interface{}
	if isDeepReflectedNullable(&nilInterface) {
		t.Errorf("Expected false for nil interface, got true")
	}

	// Test case: non-nil interface
	var nonNilInterface interface{} = "some value"
	if isDeepReflectedNullable(&nonNilInterface) {
		t.Errorf("Expected false for non-nil interface, got true")
	}

	// Test case: empty but non-nil slice
	emptySlice := []string{}
	if isDeepReflectedNullable(&emptySlice) {
		t.Errorf("Expected false for non-nil empty slice, got true")
	}

	// Test case: empty but non-nil map
	emptyMap := map[string]string{}
	if isDeepReflectedNullable(&emptyMap) {
		t.Errorf("Expected false for non-nil empty map, got true")
	}
}
