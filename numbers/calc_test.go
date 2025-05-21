// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package numbers

import "testing"

func TestBinomialCoefficient(t *testing.T) {

	// Test 1
	if BinCoef(5, 2).Int64() != 10 {
		t.Errorf("Expected 10, got %d", BinCoef(5, 2).Int64())
	}

	// Test 2
	if BinCoef(5, 3).Int64() != 10 {
		t.Errorf("Expected 10, got %d", BinCoef(5, 3).Int64())
	}

	// Test 3
	if BinCoef(5, 4).Int64() != 5 {
		t.Errorf("Expected 5, got %d", BinCoef(5, 4).Int64())
	}
}
