// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ptrs

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	type MySlice []int
	tests := []struct {
		name     string
		input    any
		call     func() any
		expected any
	}{
		{
			name:  "unnamed slice",
			input: []int{1, 2, 3},
			call: func() any {
				return Slice[[]int](1, 2, 3)
			},
			expected: &[]int{1, 2, 3},
		},
		{
			name:  "named slice",
			input: MySlice{4, 5, 6},
			call: func() any {
				return Slice[MySlice](4, 5, 6)
			},
			expected: &MySlice{4, 5, 6},
		},
		{
			name:  "preexisting slice spread into unnamed slice",
			input: []int{7, 8, 9},
			call: func() any {
				s := []int{7, 8, 9}
				return Slice[[]int](s...)
			},
			expected: &[]int{7, 8, 9},
		},
		{
			name:  "preexisting slice spread into named slice",
			input: []int{10, 11, 12},
			call: func() any {
				s := []int{10, 11, 12}
				return Slice[MySlice](s...)
			},
			expected: &MySlice{10, 11, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.call()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("%s:\n  got:      %v\n  expected: %v", tt.name, got, tt.expected)
			}
		})
	}
}
