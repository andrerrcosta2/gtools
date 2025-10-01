// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package floats

import (
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

// TestAbs32 test multiple float32 values expecting:
//   - its positive values
func TestAbs32(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		expected float32
	}{
		{
			name:     "Positive number",
			input:    42.5,
			expected: 42.5,
		},
		{
			name:     "Negative number",
			input:    -42.5,
			expected: 42.5,
		},
		{
			name:     "Zero",
			input:    0.0,
			expected: 0.0,
		},
		{
			name:     "Negative zero",
			input:    float32(math.Copysign(0.0, -1)),
			expected: 0.0,
		},
		{
			name:     "Positive infinity",
			input:    float32(math.Inf(1)),
			expected: float32(math.Inf(1)),
		},
		{
			name:     "Negative infinity",
			input:    float32(math.Inf(-1)),
			expected: float32(math.Inf(1)),
		},
		{
			name:     "NaN",
			input:    float32(math.NaN()),
			expected: float32(math.NaN()),
		},
		{
			name:     "Smallest positive non-zero",
			input:    math.SmallestNonzeroFloat32,
			expected: math.SmallestNonzeroFloat32,
		},
		{
			name:     "Smallest negative non-zero",
			input:    -math.SmallestNonzeroFloat32,
			expected: math.SmallestNonzeroFloat32,
		},
		{
			name:     "Largest positive float32",
			input:    math.MaxFloat32,
			expected: math.MaxFloat32,
		},
		{
			name:     "Largest negative float32",
			input:    -math.MaxFloat32,
			expected: math.MaxFloat32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Abs32(tt.input)
			if math.IsNaN(float64(tt.expected)) {
				if !math.IsNaN(float64(got)) {
					t.Errorf("Abs32(%v) = %v; want NaN", tt.input, got)
				}
			} else if got != tt.expected {
				t.Errorf("Abs32(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// TestAbs32 test multiple float64 values expecting:
//   - its positive values
func TestAbs64(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "Positive number",
			input:    42.5,
			expected: 42.5,
		},
		{
			name:     "Negative number",
			input:    -42.5,
			expected: 42.5,
		},
		{
			name:     "Zero",
			input:    0.0,
			expected: 0.0,
		},
		{
			name:     "Negative zero",
			input:    math.Copysign(0.0, -1),
			expected: 0.0,
		},
		{
			name:     "Positive infinity",
			input:    math.Inf(1),
			expected: math.Inf(1),
		},
		{
			name:     "Negative infinity",
			input:    math.Inf(-1),
			expected: math.Inf(1),
		},
		{
			name:     "NaN",
			input:    math.NaN(),
			expected: math.NaN(),
		},
		{
			name:     "Smallest positive non-zero",
			input:    math.SmallestNonzeroFloat64,
			expected: math.SmallestNonzeroFloat64,
		},
		{
			name:     "Smallest negative non-zero",
			input:    -math.SmallestNonzeroFloat64,
			expected: math.SmallestNonzeroFloat64,
		},
		{
			name:     "Largest positive float64",
			input:    math.MaxFloat64,
			expected: math.MaxFloat64,
		},
		{
			name:     "Largest negative float64",
			input:    -math.MaxFloat64,
			expected: math.MaxFloat64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Abs64(tt.input)
			if math.IsNaN(tt.expected) {
				if !math.IsNaN(got) {
					t.Errorf("Abs64(%v) = %v; want NaN", tt.input, got)
				}
			} else if got != tt.expected {
				t.Errorf("Abs64(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// TestFromBytes tests the FromBytes method capability of converting
// correctly byte slices into its float representation
func TestFromBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected any
		err      error
	}{
		// float32 tests (4 bytes)
		{
			name:     "float32 positive number",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(42.5)),
			expected: float32(42.5),
			err:      nil,
		},
		{
			name:     "float32 negative number",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(-42.5)),
			expected: float32(-42.5),
			err:      nil,
		},
		{
			name:     "float32 zero",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(0.0)),
			expected: float32(0.0),
			err:      nil,
		},
		{
			name:     "float32 negative zero",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(float32(math.Copysign(0.0, -1)))),
			expected: float32(math.Copysign(0.0, -1)),
			err:      nil,
		},
		{
			name:     "float32 positive infinity",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(float32(math.Inf(1)))),
			expected: float32(math.Inf(1)),
			err:      nil,
		},
		{
			name:     "float32 negative infinity",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(float32(math.Inf(-1)))),
			expected: float32(math.Inf(-1)),
			err:      nil,
		},
		{
			name:     "float32 NaN",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(float32(math.NaN()))),
			expected: float32(math.NaN()),
			err:      nil,
		},
		{
			name:     "float32 smallest positive non-zero",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(math.SmallestNonzeroFloat32)),
			expected: float32(math.SmallestNonzeroFloat32), // Explicitly float32
			err:      nil,
		},
		{
			name:     "float32 largest positive",
			input:    binary.BigEndian.AppendUint32(nil, math.Float32bits(math.MaxFloat32)),
			expected: float32(math.MaxFloat32),
			err:      nil,
		},

		// float64 tests (8 bytes)
		{
			name:     "float64 positive number",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(42.5)),
			expected: 42.5, // float64
			err:      nil,
		},
		{
			name:     "float64 negative number",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(-42.5)),
			expected: -42.5, // float64
			err:      nil,
		},
		{
			name:     "float64 zero",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(0.0)),
			expected: 0.0, // float64
			err:      nil,
		},
		{
			name:     "float64 negative zero",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(math.Copysign(0.0, -1))),
			expected: math.Copysign(0.0, -1), // float64
			err:      nil,
		},
		{
			name:     "float64 positive infinity",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(math.Inf(1))),
			expected: math.Inf(1), // float64
			err:      nil,
		},
		{
			name:     "float64 negative infinity",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(math.Inf(-1))),
			expected: math.Inf(-1), // float64
			err:      nil,
		},
		{
			name:     "float64 NaN",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(math.NaN())),
			expected: math.NaN(), // float64
			err:      nil,
		},
		{
			name:     "float64 smallest positive non-zero",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(math.SmallestNonzeroFloat64)),
			expected: math.SmallestNonzeroFloat64, // float64
			err:      nil,
		},
		{
			name:     "float64 largest positive",
			input:    binary.BigEndian.AppendUint64(nil, math.Float64bits(math.MaxFloat64)),
			expected: math.MaxFloat64, // float64
			err:      nil,
		},

		// Error cases
		{
			name:     "Empty byte slice",
			input:    []byte{},
			expected: nil,
			err:      fmt.Errorf("invalid byte slice length 0"),
		},
		{
			name:     "Invalid length 3 bytes",
			input:    []byte{0x00, 0x01, 0x02},
			expected: nil,
			err:      fmt.Errorf("invalid byte slice length 3"),
		},
		{
			name:     "Invalid length 5 bytes",
			input:    []byte{0x00, 0x01, 0x02, 0x03, 0x04},
			expected: nil,
			err:      fmt.Errorf("invalid byte slice length 5"),
		},
		{
			name:     "Invalid length 9 bytes",
			input:    []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			expected: nil,
			err:      fmt.Errorf("invalid byte slice length 9"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromBytes(tt.input)

			// Check error
			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("FromBytes(%v) error = %v; want %v", tt.input, err, tt.err)
				}
				if got != nil {
					t.Errorf("FromBytes(%v) got = %v; want nil", tt.input, got)
				}
				return
			}
			if err != nil {
				t.Errorf("FromBytes(%v) unexpected error: %v", tt.input, err)
				return
			}

			// Check value based on expected type
			switch expected := tt.expected.(type) {
			case float32:
				if gotFloat32, ok := got.(float32); !ok {
					t.Errorf("FromBytes(%v) type = %T; want float32", tt.input, got)
				} else if math.IsNaN(float64(expected)) {
					if !math.IsNaN(float64(gotFloat32)) {
						t.Errorf("FromBytes(%v) = %v; want NaN", tt.input, gotFloat32)
					}
				} else if gotFloat32 != expected {
					t.Errorf("FromBytes(%v) = %v; want %v", tt.input, gotFloat32, expected)
				}
			case float64:
				if gotFloat64, ok := got.(float64); !ok {
					t.Errorf("FromBytes(%v) type = %T; want float64", tt.input, got)
				} else if math.IsNaN(expected) {
					if !math.IsNaN(gotFloat64) {
						t.Errorf("FromBytes(%v) = %v; want NaN", tt.input, gotFloat64)
					}
				} else if gotFloat64 != expected {
					t.Errorf("FromBytes(%v) = %v; want %v", tt.input, gotFloat64, expected)
				}
			default:
				t.Errorf("FromBytes(%v) unexpected expected type: %T", tt.input, expected)
			}
		})
	}
}

func TestMax64Fraction(t *testing.T) {
	const tolerance = 1e-15 // Tolerance for floating-point comparisons

	tests := []struct {
		name     string
		input    float64
		expected float64
		isNaN    bool // True if expected result is NaN
	}{
		{
			name:     "Positive number",
			input:    42.5,
			expected: 42.5 / math.MaxFloat64,
			isNaN:    false,
		},
		{
			name:     "Negative number",
			input:    -42.5,
			expected: -42.5 / math.MaxFloat64,
			isNaN:    false,
		},
		{
			name:     "Zero",
			input:    0.0,
			expected: 0.0,
			isNaN:    false,
		},
		{
			name:     "Negative zero",
			input:    math.Copysign(0.0, -1),
			expected: 0.0,
			isNaN:    false,
		},
		{
			name:     "MaxFloat64",
			input:    math.MaxFloat64,
			expected: 1.0,
			isNaN:    false,
		},
		{
			name:     "Negative MaxFloat64",
			input:    -math.MaxFloat64,
			expected: -1.0,
			isNaN:    false,
		},
		{
			name:     "Smallest positive non-zero",
			input:    math.SmallestNonzeroFloat64,
			expected: math.SmallestNonzeroFloat64 / math.MaxFloat64,
			isNaN:    false,
		},
		{
			name:     "Smallest negative non-zero",
			input:    -math.SmallestNonzeroFloat64,
			expected: -math.SmallestNonzeroFloat64 / math.MaxFloat64,
			isNaN:    false,
		},
		{
			name:     "Positive infinity",
			input:    math.Inf(1),
			expected: math.Inf(1),
			isNaN:    false,
		},
		{
			name:     "Negative infinity",
			input:    math.Inf(-1),
			expected: math.Inf(-1),
			isNaN:    false,
		},
		{
			name:     "NaN",
			input:    math.NaN(),
			expected: 0.0, // Expected to be NaN
			isNaN:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max64Fraction(tt.input)
			if tt.isNaN {
				if !math.IsNaN(got) {
					t.Errorf("Max64Fraction(%v) = %v; want NaN", tt.input, got)
				}
			} else if math.Abs(got-tt.expected) > tolerance {
				t.Errorf("Max64Fraction(%v) = %v; want %v (within tolerance %v)", tt.input, got, tt.expected, tolerance)
			}
		})
	}
}
