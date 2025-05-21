// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package polyn

import (
	"testing"
	"testing/quick"
)

// Quick Check for N function
func TestN(t *testing.T) {
	check := func(n int, coefficients ...float64) bool {
		// Calculate the expected result
		var expected float64
		for i, cfc := range coefficients {
			expected += cfc * float64(Pow(n, i))
		}

		// Check if the result from N matches the expected value
		actual := N(n, coefficients...)
		return actual == expected
	}

	if err := quick.Check(check, nil); err != nil {
		t.Errorf("TestN failed: %v", err)
	}
}

// Quick Check for Lin function
func TestLin(t *testing.T) {
	check := func(a, b, n float64) bool {
		// Calculate the expected result
		expected := a + b*n

		// Check if the result from Lin matches the expected value
		actual := Lin(a, b, int(n))
		return actual == expected
	}

	if err := quick.Check(check, nil); err != nil {
		t.Errorf("TestLin failed: %v", err)
	}
}

// Quick Check for Quad function
func TestQuad(t *testing.T) {
	check := func(a, b, c, n float64) bool {
		// Calculate the expected result
		expected := a + b*n + c*(n*n)

		// Check if the result from Quad matches the expected value
		actual := Quad(a, b, c, int(n))
		return actual == expected
	}

	if err := quick.Check(check, nil); err != nil {
		t.Errorf("TestQuad failed: %v", err)
	}
}

// Quick Check for Cub function
func TestCub(t *testing.T) {
	check := func(a, b, c, d, n float64) bool {
		// Calculate the expected result
		expected := a + b*n + c*(n*n) + d*(n*n*n)

		// Check if the result from Cub matches the expected value
		actual := Cub(a, b, c, d, int(n))
		return actual == expected
	}

	if err := quick.Check(check, nil); err != nil {
		t.Errorf("TestCub failed: %v", err)
	}
}

// Quick Check for Pow function
func TestPow(t *testing.T) {
	check := func(base, exp int) bool {
		// Calculate the expected result
		expected := Pow(base, exp)

		// Check if the result from Pow matches the expected value
		actual := 1
		for i := 0; i < exp; i++ {
			actual *= base
		}
		return actual == expected
	}

	if err := quick.Check(check, nil); err != nil {
		t.Errorf("TestPow failed: %v", err)
	}
}
