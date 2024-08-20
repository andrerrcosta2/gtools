// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package typers

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

func TestXrtSymbol(t *testing.T) {
	tests := []struct {
		input            string
		expectedSymbol   string
		expectedVariable string
	}{
		{"***lib32.ThreeA", "***", "lib32.ThreeA"},
		{"&&ThreeA", "&&", "ThreeA"},
		{"*int", "*", "int"},
		{"**string", "**", "string"},
		{"***MyType", "***", "MyType"},
		{"&&SomeOtherType", "&&", "SomeOtherType"},
		{"&pkg.TypeName", "&", "pkg.TypeName"},
		{"**pkg.Subpkg.TypeName", "**", "pkg.Subpkg.TypeName"},
		{"lib32.ThreeA", "", "lib32.ThreeA"},
		{"int", "", "int"},
	}

	for _, tt := range tests {
		symbol, variable := XrtSymbol(tt.input)
		if symbol != tt.expectedSymbol {
			t.Errorf("XrtSymbol(%q) symbol = %v, want %v", tt.input, symbol, tt.expectedSymbol)
		}
		if variable != tt.expectedVariable {
			t.Errorf("XrtSymbol(%q) variable = %v, want %v", tt.input, variable, tt.expectedVariable)
		}
	}
}
