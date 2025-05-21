// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package generics

import (
	"reflect"
	"testing"
)

// Test cases for Zero function
func TestZero(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
	}{
		{
			name:     "int",
			expected: 0,
		},
		{
			name:     "float64",
			expected: 0.0,
		},
		{
			name:     "string",
			expected: "",
		},
		{
			name:     "bool",
			expected: false,
		},
		{
			name:     "*int",
			expected: (*int)(nil),
		},
		{
			name:     "[]int",
			expected: ([]int)(nil),
		},
		{
			name:     "map[string]int",
			expected: (map[string]int)(nil),
		},
		{
			name:     "chan int",
			expected: (chan int)(nil),
		},
		{
			name:     "struct{}",
			expected: struct{}{},
		},
		{
			name: "custom struct",
			expected: struct {
				Field1 int
				Field2 string
			}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "int":
				if got := Zero[int](); got != tt.expected {
					t.Errorf("Zero[int]() = %v, expected %v", got, tt.expected)
				}
			case "float64":
				if got := Zero[float64](); got != tt.expected {
					t.Errorf("Zero[float64]() = %v, expected %v", got, tt.expected)
				}
			case "string":
				if got := Zero[string](); got != tt.expected {
					t.Errorf("Zero[string]() = %v, expected %v", got, tt.expected)
				}
			case "bool":
				if got := Zero[bool](); got != tt.expected {
					t.Errorf("Zero[bool]() = %v, expected %v", got, tt.expected)
				}
			case "*int":
				if got := Zero[*int](); !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Zero[*int]() = %v, expected %v", got, tt.expected)
				}
			case "[]int":
				if got := Zero[[]int](); !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Zero[[]int]() = %v, expected %v", got, tt.expected)
				}
			case "map[string]int":
				if got := Zero[map[string]int](); !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Zero[map[string]int]() = %v, expected %v", got, tt.expected)
				}
			case "chan int":
				if got := Zero[chan int](); !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Zero[chan int]() = %v, expected %v", got, tt.expected)
				}
			case "struct{}":
				if got := Zero[struct{}](); !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Zero[struct{}] = %v, expected %v", got, tt.expected)
				}
			case "custom struct":
				if got := Zero[struct {
					Field1 int
					Field2 string
				}](); !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Zero[custom struct] = %v, expected %v", got, tt.expected)
				}
			}
		})
	}
}
