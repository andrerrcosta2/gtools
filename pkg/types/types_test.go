// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import "testing"

func TestIsBuiltinType(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Test with builtin types
		{"string", true},
		{"*string", true},
		{"&string", true},
		{"int", true},
		{"*int", true},
		{"&int", true},
		{"float64", true},
		{"*float64", true},
		{"&float64", true},
		{"map", true},
		{"*map", true},
		{"&map", true},

		// Test with non-builtin types
		{"customType", false},
		{"*customType", false},
		{"&customType", false},
		{"some/other/type", false},
		{"*some/other/type", false},
		{"&some/other/type", false},

		// Test with malformed types
		{"", false},
		{"*", false},
		{"&*", false},
		{"*int*", false},
		{"&int&", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := IsBuiltinType(tt.input)
			if result != tt.expected {
				t.Errorf("IsBuiltinType(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
