// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package casters

import (
	"testing"
)

func TestTyped(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []int
		ok       bool
	}{
		{"All correct types", []interface{}{1, 2, 3}, []int{1, 2, 3}, true},
		{"Mixed types with correct type", []interface{}{1, "string", 3}, []int{1, 3}, false},
		{"All incorrect types", []interface{}{"string", 3.14}, nil, false},
		{"Empty input", []interface{}{}, []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Typed[int](tt.input...)
			if !Equals(result, tt.expected) || ok != tt.ok {
				t.Errorf("Typed(%v) = (%v, %v); want (%v, %v)", tt.input, result, ok, tt.expected, tt.ok)
			}
		})
	}
}

func TestAssertedTyped(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []int
		panic    bool
	}{
		{"All correct types", []interface{}{1, 2, 3}, []int{1, 2, 3}, false},
		{"Mixed types with correct type", []interface{}{1, "string", 3}, []int{1, 3}, true}, // Expect panic here
		{"All incorrect types", []interface{}{"string", 3.14}, nil, true},                   // Expect panic here
		{"Empty input", []interface{}{}, []int{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("AssertedTyped(%v) did not panic as expected", tt.input)
					}
				}()
			}

			result := AssertedTyped[int](tt.input...)
			if !Equals(result, tt.expected) {
				t.Errorf("AssertedTyped(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func Equals[T comparable](a, b []T) bool {
	wa, wb := a, b
	// Check if the slices have different lengths
	if len(wa) != len(wb) {
		return false
	}

	// Compare each element of the slices
	for i := range wa {
		if wa[i] != wb[i] {
			return false
		}
	}

	// Slices are equal
	return true
}
