// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prng

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"testing"
)

// shouldNotOverflowRandomIntInterval tests the Int function to check if the
// generated integer is within the range specified by min and max.
func shouldNotOverflowRandomIntInterval(t *testing.T, min, max int) {
	t.Helper()
	result := Int(min, max)
	if result > max || result < min {
		t.Fatalf(
			"Int() = %d; want value in range [%d, %d]",
			result, min, max)
	}
}

// shouldNotOverflowRandomInt8Interval tests the Int8 function to ensure
// that the generated integer is within the range specified by min and max.
func shouldNotOverflowRandomInt8Interval(t *testing.T, min, max int8) {
	t.Helper()
	// Generate a random int8 within the specified range
	result := Int8(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Int8() = %d; want value in range [%d, %d]", result, min, max)
	}
}

// shouldNotOverflowRandomInt16Interval tests the Int16 function to check if the
// generated integer is within the range specified by min and max.
func shouldNotOverflowRandomInt16Interval(t *testing.T, min, max int16) {
	t.Helper()
	// Generate a random int16 within the specified range
	result := Int16(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Int16() = %d; want value in range [%d, %d]", result, min, max)
	}
}

// shouldNotOverflowRandomInt32Interval tests the Int32 function to check if the
// generated integer is within the range specified by min and max.
func shouldNotOverflowRandomInt32Interval(t *testing.T, min, max int32) {
	t.Helper()
	// Generate a random int32 within the specified range
	result := Int32(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Int32() = %d; want value in range [%d, %d]", result, min, max)
	}
}

func shouldNotOverflowRandomInt64Interval(t *testing.T, min, max int64) {
	t.Helper()
	// Generate a random int64 within the specified range
	result := Int64(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Int64() = %d; want value in range [%d, %d]", result, min, max)
	}
}

func shouldNotOverflowRandomUintInterval(t *testing.T, min, max uint) {
	t.Helper()
	// Generate a random uint within the specified range
	result := Uint(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Uint() = %d; want value in range [%d, %d]", result, min, max)
	}
}

func shouldNotOverflowRandomUint8Interval(t *testing.T, min, max uint8) {
	t.Helper()
	// Generate a random uint8 within the specified range
	result := Uint8(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Uint8() = %d; want value in range [%d, %d]", result, min, max)
	}
}

func shouldNotOverflowRandomUint16Interval(t *testing.T, min, max uint16) {
	t.Helper()
	// Generate a random uint16 within the specified range
	result := Uint16(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Uint16() = %d; want value in range [%d, %d]", result, min, max)
	}
}

func shouldNotOverflowRandomUint32Interval(t *testing.T, min, max uint32) {
	t.Helper()
	// Generate a random uint32 within the specified range
	result := Uint32(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Uint32() = %d; want value in range [%d, %d]", result, min, max)
	}
}

func shouldNotOverflowRandomUint64Interval(t *testing.T, min, max uint64) {
	t.Helper()
	// Generate a random uint64 within the specified range
	result := Uint64(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Uint64() = %d; want value in range [%d, %d]", result, min, max)
	}
}

func shouldNotOverflowRandomFloat32Interval(t *testing.T, min, max float32) {
	t.Helper()
	// Generate a random float32 within the specified range
	result := Float32(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Float32() = %f; want value in range [%f, %f]", result, min, max)
	}
}

func shouldNotOverflowRandomFloat64Interval(t *testing.T, min, max float64) {
	t.Helper()
	// Generate a random float64 within the specified range
	result := Float64(min, max)

	// Check if the result is outside the desired range
	if result > max || result < min {
		// Fail the test if the result is not within the expected range
		t.Fatalf("Float64() = %f; want value in range [%f, %f]", result, min, max)
	}
}

// shouldNotOverflowRandomComplex64Interval tests the Complex64 function to
// ensure that the generated complex64 number's real and imaginary components
// are within the specified ranges.
func shouldNotOverflowRandomComplex64Interval(t *testing.T, rm, rM, im, iM float32) {
	t.Helper()
	// Generate a random complex64 within the specified range
	result := Complex64(rm, rM, im, iM)

	// Extract real and imaginary components
	realPart := real(result)
	imagPart := imag(result)

	// Check if the real part is within range
	if realPart < rm || realPart > rM {
		// Fail the test if the real part is outside the expected range
		t.Fatalf("Real part of Complex64() = %f; want value in range [%f, %f]", realPart, rm, rM)
	}

	// Check if the imaginary part is within range
	if imagPart < im || imagPart > iM {
		// Fail the test if the imaginary part is outside the expected range
		t.Fatalf("Imaginary part of Complex64() = %f; want value in range [%f, %f]", imagPart, im, iM)
	}
}

// shouldNotOverflowRandomComplex128Interval tests the Complex128 function
// to ensure that the generated complex number's real and imaginary components
// are within the specified ranges.
func shouldNotOverflowRandomComplex128Interval(t *testing.T, rm, rM, im, iM float64) {
	t.Helper()
	// Generate a random complex128 number within the specified range
	result := Complex128(rm, rM, im, iM)

	// Extract real and imaginary components
	realPart := real(result)
	imagPart := imag(result)

	// Check if the real part is within the specified range
	if realPart < rm || realPart > rM {
		// Fail the test if the real part is outside the expected range
		t.Fatalf("Real part of Complex128() = %f; want value in range [%f, %f]", realPart, rm, rM)
	}

	// Check if the imaginary part is within the specified range
	if imagPart < im || imagPart > iM {
		// Fail the test if the imaginary part is outside the expected range
		t.Fatalf("Imaginary part of Complex128() = %f; want value in range [%f, %f]", imagPart, im, iM)
	}
}

func shouldHaveNBuckets[k prim.Comparable](t *testing.T, groups map[k]int, nbuckets int) {
	t.Helper()
	if len(groups) != nbuckets {
		t.Fatalf("Expected %d buckets, got %d", nbuckets, len(groups))
	}
}

func shouldNotDeviateMoreThan[K prim.Comparable](t *testing.T, groups map[K]int, total int, maxDeviation float64) {
	t.Helper()
	size := len(groups)
	expRatio := 1.0 / float64(size) // Expected ratio of each group (uniform distribution)

	for k, v := range groups {
		actualRatio := float64(v) / float64(total) // Proportion of values in the current group
		deviation := actualRatio/expRatio - 1.0    // Deviation from the expected ratio

		if deviation < -maxDeviation || deviation > maxDeviation {
			t.Fatalf("Group %v has a deviation of %f; the maximum defined is %f", k, deviation, maxDeviation)
		}
	}
}

//func chiSquaredTest[K prim.Comparable](t *testing.T, groups map[K]int, total int, alpha float64) {
//	size := len(groups)
//	expected := float64(total) / float64(size)
//	chiSquared := 0.0
//
//	for _, v := range groups {
//		observed := float64(v)
//		chiSquared += (observed - expected) * (observed - expected) / expected
//	}
//
//	// Critical value for the chi-squared distribution with (size - 1) degrees of freedom.
//	criticalValue := chiSquaredCriticalValue(size-1, alpha)
//
//	if chiSquared > criticalValue {
//		t.Fatalf("Chi-squared test failed: χ² = %f, critical value = %f", chiSquared, criticalValue)
//	}
//}
