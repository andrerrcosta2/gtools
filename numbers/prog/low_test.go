// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prog

import "testing"

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		name     string
		start    float64
		ratio    float64
		length   int
		expected []float64
		err      bool
	}{
		{
			name:     "Simple test with ratio 2",
			start:    16,
			ratio:    2,
			length:   4,
			expected: []float64{16, 8, 4, 2},
			err:      false,
		},
		{
			name:     "Test with ratio 1.5",
			start:    27,
			ratio:    1.5,
			length:   4,
			expected: []float64{27, 18, 12, 8},
			err:      false,
		},
		{
			name:     "Test with ratio 1",
			start:    5,
			ratio:    1,
			length:   3,
			expected: []float64{5, 5, 5},
			err:      false,
		},
		{
			name:     "Test with length 1",
			start:    9,
			ratio:    3,
			length:   1,
			expected: []float64{9},
			err:      false,
		},
		{
			name:     "Zero ratio",
			start:    10,
			ratio:    0,
			length:   3,
			expected: []float64{},
			err:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Sqrt(tt.start, tt.ratio, tt.length)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
			}
			if (err != nil) != tt.err {
				t.Errorf("Expected no error, got %v", err)
			}

			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("At index %d: expected %v, got %v", i, tt.expected[i], v)
				}
			}
		})
	}
}

func TestCubic(t *testing.T) {
	tests := []struct {
		name       string
		a, b, c, d float64
		length     int
		expected   []float64
		err        bool
	}{
		{
			name:     "Simple cubic test",
			a:        1,
			b:        2,
			c:        3,
			d:        4,
			length:   5,
			expected: []float64{1, 10, 49, 142, 313},
			err:      false,
		},
		{
			name:     "Cub with zero coefficients",
			a:        0,
			b:        0,
			c:        0,
			d:        4,
			length:   4,
			expected: []float64{0, 4, 32, 108},
			err:      false,
		},
		{
			name:     "Cub with negative coefficients",
			a:        -1,
			b:        -2,
			c:        -3,
			d:        -4,
			length:   3,
			expected: []float64{-1, -10, -49},
			err:      false,
		},
		{
			name:     "Cub with one element",
			a:        1,
			b:        1,
			c:        1,
			d:        1,
			length:   1,
			expected: []float64{1},
			err:      false,
		},
		{
			name:     "Cub with larger sequence",
			a:        1,
			b:        2,
			c:        3,
			d:        4,
			length:   10,
			expected: []float64{1, 10, 49, 142, 313, 586, 985, 1534, 2257, 3178},
			err:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Cub(tt.a, tt.b, tt.c, tt.d, tt.length)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
			}
			if (err != nil) != tt.err {
				t.Errorf("Expected no error, got %v", err)
			}

			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("At index %d: expected %v, got %v", i, tt.expected[i], v)
				}
			}
		})
	}
}

func TestQuadratic(t *testing.T) {
	tests := []struct {
		name     string
		a, b, c  float64
		length   int
		expected []float64
		err      bool
	}{
		{
			name:     "Simple quadratic test",
			a:        1,
			b:        2,
			c:        1,
			length:   5,
			expected: []float64{1, 4, 9, 16, 25},
			err:      false,
		},
		{
			name:     "Quad with zero coefficients",
			a:        0,
			b:        0,
			c:        1,
			length:   5,
			expected: []float64{0, 1, 4, 9, 16},
			err:      false,
		},
		{
			name:     "Quad with a negative coefficient",
			a:        -1,
			b:        2,
			c:        1,
			length:   5,
			expected: []float64{-1, 2, 7, 14, 23},
			err:      false,
		},
		{
			name:     "Quad with mixed coefficients",
			a:        0.5,
			b:        1.5,
			c:        2,
			length:   5,
			expected: []float64{0.5, 4, 11.5, 23, 38.5},
			err:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Quad(tt.a, tt.b, tt.c, tt.length)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d, but got %d", len(tt.expected), len(result))
			}
			if (err != nil) != tt.err {
				t.Errorf("Expected no error, but got %v", err)
			}

			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("At index %d, expected %v, but got %v", i, tt.expected[i], v)
				}
			}
		})
	}
}

func TestRandom_Int(t *testing.T) {
	length := 10
	minimum := 1
	maximum := 10

	randoms, _ := Rand[int](length, minimum, maximum)

	if len(randoms) != length {
		t.Errorf("Expected length %d, but got %d", length, len(randoms))
	}

	for _, v := range randoms {
		if v < minimum || v > maximum {
			t.Errorf("Value %d is out of the expected range [%d, %d]", v, minimum, maximum)
		}
	}
}

func TestRandom_Float64(t *testing.T) {
	length := 10
	minimum := 0.5
	maximum := 5.5

	randoms, _ := Rand[float64](length, minimum, maximum)

	if len(randoms) != length {
		t.Errorf("Expected length %d, but got %d", length, len(randoms))
	}

	for _, v := range randoms {
		if v < minimum || v > maximum {
			t.Errorf("Value %f is out of the expected range [%f, %f]", v, minimum, maximum)
		}
	}
}

func TestRandom_EmptySlice(t *testing.T) {
	length := 0
	minimum := 1
	maximum := 10

	randoms, _ := Rand[int](length, minimum, maximum)

	if len(randoms) != 0 {
		t.Errorf("Expected length 0, but got %d", len(randoms))
	}
}

func TestFibonacci_Int(t *testing.T) {
	length := 10
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	fib, _ := Fib[int](length)

	if len(fib) != length {
		t.Errorf("Expected length %d, but got %d", length, len(fib))
	}

	for i := range expected {
		if fib[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], fib[i])
		}
	}
}

func TestFibonacci_Float64(t *testing.T) {
	length := 10
	expected := []float64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	fib, _ := Fib[float64](length)

	if len(fib) != length {
		t.Errorf("Expected length %d, but got %d", length, len(fib))
	}

	for i := range expected {
		if fib[i] != expected[i] {
			t.Errorf("At index %d, expected %f, but got %f", i, expected[i], fib[i])
		}
	}
}

func TestFibonacci_ShortLength(t *testing.T) {
	length := 2
	expected := []int{0, 1}

	fib, _ := Fib[int](length)

	if len(fib) != length {
		t.Errorf("Expected length %d, but got %d", length, len(fib))
	}

	for i := range expected {
		if fib[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], fib[i])
		}
	}
}

func TestFibonacci_Empty(t *testing.T) {
	length := 0

	fib, _ := Fib[int](length)

	if len(fib) != length {
		t.Errorf("Expected length %d, but got %d", length, len(fib))
	}
}

func TestGeometric_Int(t *testing.T) {
	start := 2
	ratio := 3
	length := 5
	expected := []int{2, 6, 18, 54, 162}

	gp, _ := Geom[int](start, ratio, length)

	if len(gp) != length {
		t.Errorf("Expected length %d, but got %d", length, len(gp))
	}

	for i := range expected {
		if gp[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], gp[i])
		}
	}
}

func TestGeometric_Float64(t *testing.T) {
	start := 1.5
	ratio := 2.0
	length := 5
	expected := []float64{1.5, 3.0, 6.0, 12.0, 24.0}

	gp, _ := Geom[float64](start, ratio, length)

	if len(gp) != length {
		t.Errorf("Expected length %d, but got %d", length, len(gp))
	}

	for i := range expected {
		if gp[i] != expected[i] {
			t.Errorf("At index %d, expected %f, but got %f", i, expected[i], gp[i])
		}
	}
}

func TestGeometric_ShortLength(t *testing.T) {
	start := 3
	ratio := 4
	length := 3
	expected := []int{3, 12, 48}

	gp, _ := Geom[int](start, ratio, length)

	if len(gp) != length {
		t.Errorf("Expected length %d, but got %d", length, len(gp))
	}

	for i := range expected {
		if gp[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], gp[i])
		}
	}
}

func TestGeometric_Empty(t *testing.T) {
	start := 5
	ratio := 2
	length := 0

	gp, _ := Geom[int](start, ratio, length)

	if len(gp) != length {
		t.Errorf("Expected length %d, but got %d", length, len(gp))
	}
}

func TestArithmetic_Int(t *testing.T) {
	start := 1
	step := 3
	length := 5
	expected := []int{1, 4, 7, 10, 13}

	ap, _ := Arit[int](start, step, length)

	if len(ap) != length {
		t.Errorf("Expected length %d, but got %d", length, len(ap))
	}

	for i := range expected {
		if ap[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], ap[i])
		}
	}
}

func TestArithmetic_Float64(t *testing.T) {
	start := 0.5
	step := 2.5
	length := 4
	expected := []float64{0.5, 3.0, 5.5, 8.0}

	ap, _ := Arit[float64](start, step, length)

	if len(ap) != length {
		t.Errorf("Expected length %d, but got %d", length, len(ap))
	}

	for i := range expected {
		if ap[i] != expected[i] {
			t.Errorf("At index %d, expected %f, but got %f", i, expected[i], ap[i])
		}
	}
}

func TestArithmetic_ShortLength(t *testing.T) {
	start := 10
	step := -2
	length := 3
	expected := []int{10, 8, 6}

	ap, _ := Arit[int](start, step, length)

	if len(ap) != length {
		t.Errorf("Expected length %d, but got %d", length, len(ap))
	}

	for i := range expected {
		if ap[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], ap[i])
		}
	}
}

func TestArithmetic_Empty(t *testing.T) {
	start := 5
	step := 1
	length := 0

	ap, _ := Arit[int](start, step, length)

	if len(ap) != length {
		t.Errorf("Expected length %d, but got %d", length, len(ap))
	}
}

func TestSequence_Int(t *testing.T) {
	length := 5
	expected := []int{0, 1, 2, 3, 4}

	seq := Seq[int](length)

	if len(seq) != length {
		t.Errorf("Expected length %d, but got %d", length, len(seq))
	}

	for i := range expected {
		if seq[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], seq[i])
		}
	}
}

func TestSequence_Float64(t *testing.T) {
	length := 4
	expected := []float64{0, 1, 2, 3}

	seq := Seq[float64](length)

	if len(seq) != length {
		t.Errorf("Expected length %d, but got %d", length, len(seq))
	}

	for i := range expected {
		if seq[i] != expected[i] {
			t.Errorf("At index %d, expected %f, but got %f", i, expected[i], seq[i])
		}
	}
}

func TestSequence_Empty(t *testing.T) {
	length := 0

	seq := Seq[int](length)

	if len(seq) != length {
		t.Errorf("Expected length %d, but got %d", length, len(seq))
	}
}
