// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package casters

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/core/testlite/testseed"
)

// TestTyped tests the capability of the Types function.
//
// Types casts the provided values to the type G and returns a slice of type G.
// The second return value allMatches is true if all values were successfully cast
// to type G, and false otherwise.
func TestTyped(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []int
		ok       bool
	}{
		{"WhenAllCancels correct types", []interface{}{1, 2, 3}, []int{1, 2, 3}, true},
		{"Mixed types with correct type", []interface{}{1, "string", 3}, []int{1, 3}, false},
		{"WhenAllCancels incorrect types", []interface{}{"string", 3.14}, []int{}, false},
		{"IsEmpty input", []interface{}{}, []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Types[int](tt.input...)
			assertlite.EqualSlices(t, result, tt.expected)
			assertlite.True(t, ok == tt.ok)
		})
	}
}

// TestAssertedTyped tests the capability of the Assert function.
//
// Assert casts the provided values to the type G and returns a slice of type G.
// If any value can't be cast to type G, Assert panics.
func TestAssertedTyped(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []int
		panic    bool
	}{
		{"WhenAllCancels correct types", []interface{}{1, 2, 3}, []int{1, 2, 3}, false},
		{"Mixed types with correct type", []interface{}{1, "string", 3}, []int{1, 3}, true}, // Expect panic here
		{"WhenAllCancels incorrect types", []interface{}{"string", 3.14}, nil, true},        // Expect panic here
		{"IsEmpty input", []interface{}{}, []int{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assertlite.Panic(t, func() {
					Assert[int](tt.input...)
				})
			} else {
				result := Assert[int](tt.input...)
				assertlite.EqualSlices(t, result, tt.expected)
			}
		})
	}
}

// Test_UnsafeInlineOf_Primitives tests the UnsafeValueOf cast function.
//
// This test should assert the values returned by "UnsafeValueOf" meet its original
// type before being passed as 'interface{}'
//
// Inline are values that aren't golang underlying references
// such as primitives, structs
func Test_UnsafeValueOf_Primitives(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		var i any = 42
		result := UnsafeValueOf[int](i)
		assertlite.IsTypeOf[int](t, result)
		assertlite.Equals(t, result, 42)
	})

	t.Run("bool", func(t *testing.T) {
		var b any = true
		result := UnsafeValueOf[bool](b)
		assertlite.IsTypeOf[bool](t, result)
		assertlite.Equals(t, result, true)
	})

	t.Run("float64", func(t *testing.T) {
		var f any = 3.14
		result := UnsafeValueOf[float64](f)
		assertlite.IsTypeOf[float64](t, result)
		assertlite.Equals(t, result, 3.14)
	})

	t.Run("string", func(t *testing.T) {
		var s any = "hello"
		result := UnsafeValueOf[string](s)
		assertlite.IsTypeOf[string](t, result)
		assertlite.Equals(t, result, "hello")
	})

	t.Run("slice", func(t *testing.T) {
		var s any = []int{1, 2, 3}
		result := UnsafeValueOf[[]int](s)
		assertlite.IsTypeOf[[]int](t, result)
		assertlite.Equals(t, result, s)
	})
}

func TestUnsafeValueOf_Structs(t *testing.T) {
	t.Run("comparable struct", func(t *testing.T) {
		var s any = testseed.ComparableValue{Name: "John", Age: 31}
		result := UnsafeValueOf[testseed.ComparableValue](s)
		assertlite.IsTypeOf[testseed.ComparableValue](t, result)
		assertlite.Equals(t, result, s)
	})

	t.Run("struct one data", func(t *testing.T) {
		var s any = testseed.StructOneData{Name: "John"}
		result := UnsafeValueOf[testseed.StructOneData](s)
		assertlite.IsTypeOf[testseed.StructOneData](t, result)
		assertlite.Equals(t, result, s)
	})

	t.Run("struct with ptrs", func(t *testing.T) {
		name := "John"
		var s any = testseed.StructWithPointers{Name: &name}
		result := UnsafeValueOf[testseed.StructWithPointers](s)
		assertlite.IsTypeOf[testseed.StructWithPointers](t, result)
		assertlite.Equals(t, result, s)
	})
}

func TestUnsafeReferenceOf(t *testing.T) {
	t.Run("channel", func(t *testing.T) {
		var s any = make(chan int)

		// ❌ Attempt to cast as if it's an inline value
		result := UnsafeValueOf[chan int](s)
		assertlite.IsTypeOf[chan int](t, result)
		assertlite.NotEquals(t, s, result)

		// ✅ Cast as a reference
		result = UnsafeReferenceOf[chan int](s)
		assertlite.IsTypeOf[chan int](t, result)
		assertlite.Equals(t, s, result)
		assertlite.Same(t, s, result)
	})

	t.Run("map", func(t *testing.T) {
		var s any = map[string]int{"one": 1, "two": 2}

		// ❌ Attempt to cast as if it's an inline value
		result := UnsafeValueOf[map[string]int](s)
		assertlite.IsTypeOf[map[string]int](t, result)
		assertlite.Panic(t, func() {
			assertlite.NotEquals(t, s, result)
		}, "expected panic on deep equality reflection due to bad pointer in frame")
		assertlite.NotSame(t, s, result)

		// ✅ Cast as a reference
		result = UnsafeReferenceOf[map[string]int](s)
		assertlite.IsTypeOf[map[string]int](t, result)
		assertlite.Equals(t, s, result)
		assertlite.Same(t, s, result)
	})

	t.Run("pointer", func(t *testing.T) {
		var x = 123
		var s any = &x

		// ❌ Cast as a value
		result := UnsafeValueOf[*int](s)
		assertlite.IsTypeOf[*int](t, result)
		assertlite.Panic(t, func() {
			assertlite.NotEquals(t, s, result)
		})
		assertlite.NotSame(t, s, result)

		// ✅ Cast as a reference
		result = UnsafeReferenceOf[*int](s)
		assertlite.IsTypeOf[*int](t, result)
		assertlite.Equals(t, s, result)
		assertlite.Same(t, s, result)
	})
}
