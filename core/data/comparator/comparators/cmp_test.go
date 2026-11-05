// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparators

import (
	"testing"
)

// TestOrdered_Compare tests the Compare method of the ordered type.
func TestOrdered_Compare(t *testing.T) {
	var ordered ordered[int]

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a < b", 1, 2, -1},
		{"a > b", 3, 2, 1},
		{"a == b", 4, 4, 0},
		{"negative numbers: a < b", -5, -3, -1},
		{"negative numbers: a > b", -1, -2, 1},
		{"zero comparison: a < b", 0, 1, -1},
		{"zero comparison: a > b", 1, 0, 1},
		{"zero comparison: a == b", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ordered.Compare(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Compare(%v, %v): expected %d, got %d", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}

// TestOrdered_Equals tests the Equals method of the ordered type.
func TestOrdered_Equals(t *testing.T) {
	var ordered ordered[int]

	tests := []struct {
		name     string
		a, b     int
		expected bool
	}{
		{"a == b", 1, 1, true},
		{"a != b", 1, 2, false},
		{"negative numbers: a == b", -5, -5, true},
		{"negative numbers: a != b", -5, -3, false},
		{"zero comparison: a == b", 0, 0, true},
		{"zero comparison: a != b", 0, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ordered.Equals(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Equals(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}

// TestOrdered_Generic tests the ordered type with different generic type.
func TestOrdered_Generic(t *testing.T) {
	// Test with strings
	var orderedString ordered[string]

	tests := []struct {
		name     string
		a, b     string
		expected int
	}{
		{"a < b", "apple", "banana", -1},
		{"a > b", "cherry", "banana", 1},
		{"a == b", "apple", "apple", 0},
		{"case sensitivity: a < b", "Apple", "apple", -1},
		{"empty strings: a == b", "", "", 0},
		{"empty strings: a < b", "", "a", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := orderedString.Compare(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Compare(%v, %v): expected %d, got %d", tt.a, tt.b, tt.expected, result)
			}
		})
	}

	// Test with floats
	var orderedFloat ordered[float64]

	floatTests := []struct {
		name     string
		a, b     float64
		expected int
	}{
		{"a < b", 1.5, 2.5, -1},
		{"a > b", 3.5, 2.5, 1},
		{"a == b", 4.0, 4.0, 0},
		{"negative numbers: a < b", -5.5, -3.5, -1},
		{"negative numbers: a > b", -1.5, -2.5, 1},
		{"zero comparison: a == b", 0.0, 0.0, 0},
		{"zero comparison: a < b", 0.0, 1.0, -1},
	}

	for _, tt := range floatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := orderedFloat.Compare(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Compare(%v, %v): expected %d, got %d", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}

// TestKeyOrdered_Hash tests the Hash method of the keyOrdered type.
func TestKeyOrdered_Hash(t *testing.T) {
	var keyOrderedInt idEqCmpKeyer[int]
	var keyOrderedString idEqCmpKeyer[string]
	var keyOrderedFloat idEqCmpKeyer[float64]

	tests := []struct {
		name     string
		input    any
		expected any
	}{
		{"integer hash", 42, 42},
		{"negative integer hash", -10, -10},
		{"zero integer hash", 0, 0},
		{"string hash", "hello", "hello"},
		{"empty string hash", "", ""},
		{"float hash", 3.14, 3.14},
		{"negative float hash", -2.71, -2.71},
		{"zero float hash", 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case int:
				result := keyOrderedInt.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case string:
				result := keyOrderedString.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case float64:
				result := keyOrderedFloat.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			default:
				t.Fatalf("Unsupported type: %T", tt.input)
			}
		})
	}
}

// TestKeyOrdered_Inheritance tests that keyOrdered inherits methods from ordered.
func TestKeyOrdered_Inheritance(t *testing.T) {
	var keyOrderedInt idEqCmpKeyer[int]

	// Test Compare method inherited from ordered
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a < b", 1, 2, -1},
		{"a > b", 3, 2, 1},
		{"a == b", 4, 4, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := keyOrderedInt.Compare(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Compare(%v, %v): expected %d, got %d", tt.a, tt.b, tt.expected, result)
			}
		})
	}

	// Test Equals method inherited from ordered
	equalityTests := []struct {
		name     string
		a, b     int
		expected bool
	}{
		{"a == b", 1, 1, true},
		{"a != b", 1, 2, false},
	}

	for _, tt := range equalityTests {
		t.Run(tt.name, func(t *testing.T) {
			result := keyOrderedInt.Equals(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Equals(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}

// TestStringOrdered_Hash tests the Hash method of the stringOrdered type.
func TestStringOrdered_Hash(t *testing.T) {
	var stringOrderedInt stringOrdered[int]
	var stringOrderedInt8 stringOrdered[int8]
	var stringOrderedInt16 stringOrdered[int16]
	var stringOrderedInt32 stringOrdered[int32]
	var stringOrderedInt64 stringOrdered[int64]
	var stringOrderedUint stringOrdered[uint]
	var stringOrderedUint8 stringOrdered[uint8]
	var stringOrderedUint16 stringOrdered[uint16]
	var stringOrderedUint32 stringOrdered[uint32]
	var stringOrderedUint64 stringOrdered[uint64]
	var stringOrderedFloat32 stringOrdered[float32]
	var stringOrderedFloat64 stringOrdered[float64]
	var stringOrderedString stringOrdered[string]

	tests := []struct {
		name     string
		input    any
		expected string
	}{
		// Integer type
		{"int hash", 42, "42"},
		{"negative int hash", -10, "-10"},
		{"zero int hash", 0, "0"},
		{"int8 hash", int8(127), "127"},
		{"int16 hash", int16(-32768), "-32768"},
		{"int32 hash", int32(2147483647), "2147483647"},
		{"int64 hash", int64(-9223372036854775808), "-9223372036854775808"},

		// Unsigned integer type
		{"uint hash", uint(42), "42"},
		{"uint8 hash", uint8(255), "255"},
		{"uint16 hash", uint16(65535), "65535"},
		{"uint32 hash", uint32(4294967295), "4294967295"},
		{"uint64 hash", uint64(18446744073709551615), "18446744073709551615"},

		// Floating-point type
		{"float32 hash", float32(3.14), "3.14"},
		{"negative float32 hash", float32(-2.71), "-2.71"},
		{"zero float32 hash", float32(0.0), "0"},
		{"float64 hash", 3.141592653589793, "3.141592653589793"},
		{"negative float64 hash", -2.718281828459045, "-2.718281828459045"},
		{"zero float64 hash", 0.0, "0"},

		// String type
		{"string hash", "hello", "hello"},
		{"empty string hash", "", ""},

		// Unsupported type (should panic)
		{"unsupported type", struct{}{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.name == "unsupported type" {
						// Expected panic for unsupported type
						return
					}
					t.Errorf("unexpected panic: %v", r)
				}
			}()

			switch v := tt.input.(type) {
			case int:
				result := stringOrderedInt.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case int8:
				result := stringOrderedInt8.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case int16:
				result := stringOrderedInt16.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case int32:
				result := stringOrderedInt32.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case int64:
				result := stringOrderedInt64.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case uint:
				result := stringOrderedUint.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case uint8:
				result := stringOrderedUint8.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case uint16:
				result := stringOrderedUint16.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case uint32:
				result := stringOrderedUint32.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case uint64:
				result := stringOrderedUint64.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case float32:
				result := stringOrderedFloat32.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case float64:
				result := stringOrderedFloat64.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			case string:
				result := stringOrderedString.Key(v)
				if result != tt.expected {
					t.Errorf("Hash(%v): expected %v, got %v", tt.input, tt.expected, result)
				}
			default:
				t.Logf("unsupported type: %T", tt.input)
			}
		})
	}
}

// TestStringOrdered_Inheritance tests that stringOrdered inherits methods from ordered.
func TestStringOrdered_Inheritance(t *testing.T) {
	var stringOrderedInt stringOrdered[int]

	// Test Compare method inherited from ordered
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a < b", 1, 2, -1},
		{"a > b", 3, 2, 1},
		{"a == b", 4, 4, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringOrderedInt.Compare(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Compare(%v, %v): expected %d, got %d", tt.a, tt.b, tt.expected, result)
			}
		})
	}

	// Test Equals method inherited from ordered
	equalityTests := []struct {
		name     string
		a, b     int
		expected bool
	}{
		{"a == b", 1, 1, true},
		{"a != b", 1, 2, false},
	}

	for _, tt := range equalityTests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringOrderedInt.Equals(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Equals(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}
