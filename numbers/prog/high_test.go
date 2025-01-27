// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prog

import "testing"

func TestHigh_Int(t *testing.T) {
	coefficients := []int{1, 2, 3} // Represents polynomial 1 + 2x + 3x^2
	length := 5
	expected := []int{1, 6, 17, 34, 57}

	prog := High[int](coefficients, length)

	if len(prog) != length {
		t.Errorf("Expected length %d, but got %d", length, len(prog))
	}

	for i := range expected {
		if prog[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], prog[i])
		}
	}
}

func TestHigh_Float64(t *testing.T) {
	coefficients := []float64{1.0, 2.0, 3.0} // Represents polynomial 1 + 2x + 3x^2
	length := 4
	expected := []float64{1.0, 6.0, 17.0, 34.0}

	prog := High[float64](coefficients, length)

	if len(prog) != length {
		t.Errorf("Expected length %d, but got %d", length, len(prog))
	}

	for i := range expected {
		if prog[i] != expected[i] {
			t.Errorf("At index %d, expected %f, but got %f", i, expected[i], prog[i])
		}
	}
}

func TestHigh_EmptyCoefficients(t *testing.T) {
	coefficients := []int{} // Polynomial with no coefficients
	length := 3
	expected := []int{0, 0, 0}

	prog := High[int](coefficients, length)

	if len(prog) != length {
		t.Errorf("Expected length %d, but got %d", length, len(prog))
	}

	for i := range expected {
		if prog[i] != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], prog[i])
		}
	}
}
