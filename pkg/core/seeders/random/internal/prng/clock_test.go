// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prng

import (
	"math"
	"testing"
)

// TODO: Finish distribution test, float64 and complex,
// maybe test chisquared, maybe test autocorrelation, periodogram
// TODO: For sure test parallelization
const (
	Range         = 1000
	Distributions = 1_000_000
)

func TestRandomIntRange(t *testing.T) {
	t.Run("Positive Range - (10 to 100)", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomIntInterval(t, 10, 100)
		}
	})

	t.Run("Negative Range - (-100 to -10)", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomIntInterval(t, -100, -10)
		}
	})

	t.Run("Zero Range - (0 to 0)", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomIntInterval(t, 0, 0)
		}
	})

	t.Run("Mixed Range - (-100 to 100)", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomIntInterval(t, -100, 100)
		}
	})
}

func TestRandomIntDistribution(t *testing.T) {
	t.Run("Positive Distribution (10 to 51)", func(t *testing.T) {
		m, M := 10, 51
		count := make(map[int]int)

		// Before the random number generator 1000 times
		for i := 0; i < Distributions; i++ {
			result := Int(m, M)
			if result < m || result > M {
				t.Fatalf("Int() = %d; want value in range [%d, %d]", result, m, M)
			}
			count[result]++
		}

		shouldHaveNBuckets(t, count, M-m+1)
		shouldNotDeviateMoreThan(t, count, Distributions, 0.05)
	})

	t.Run("Negative Distribution (-51 to -10)", func(t *testing.T) {
		m, M := -51, -10
		count := make(map[int]int)

		// Before the random number generator enough times
		for i := 0; i < Distributions; i++ {
			result := Int(m, M)
			if result < m || result > M {
				t.Fatalf("Int() = %d; want value in range [%d, %d]", result, m, M)
			}
			count[result]++
		}

		shouldHaveNBuckets(t, count, M-m+1)
		shouldNotDeviateMoreThan(t, count, Distributions, 0.05)
	})

	t.Run("High Edge Distribution (MaxInt - 149 to MaxInt)", func(t *testing.T) {
		m := math.MaxInt - 149
		count := make(map[int]int)

		// Before the random number generator enough times
		for i := 0; i < Distributions; i++ {
			result := Int(m)

			// Debug: Check for invalid values
			if result < m || result > math.MaxInt {
				t.Fatalf("Generated value out of bounds: %d (expected range: %d to %d)", result, m, math.MaxInt)
			}

			count[result]++
		}

		shouldHaveNBuckets(t, count, 150)
		shouldNotDeviateMoreThan(t, count, Distributions, 0.05)
	})

	t.Run("Low Edge Distribution (MinInt to MinInt + 149)", func(t *testing.T) {
		m, M := math.MinInt, math.MinInt+149
		count := make(map[int]int)

		// Before the random number generator enough times
		for i := 0; i < Distributions; i++ {
			result := Int(m, M)

			// Debug: Check for invalid values
			if result < math.MinInt || result > M {
				t.Fatalf("Generated value out of bounds: %d (expected range: %d to %d)", result, m, M)
			}

			count[result]++
		}

		shouldHaveNBuckets(t, count, 150)
		shouldNotDeviateMoreThan(t, count, Distributions, 0.05)
	})

	t.Run("Zero Distribution (0 to 0)", func(t *testing.T) {
		m, M := 0, 0
		count := make(map[int]int)

		// Before the random number generator enough times
		for i := 0; i < Distributions; i++ {
			result := Int(m, M)
			if result > M || result < m {
				t.Fatalf("Int() = %d; want value in range [%d, %d]", result, m, M)
			}
			count[result]++
		}

		if len(count) != 1 {
			t.Errorf("Not fully distributed. Expected 1, got: %d\n", len(count))
		}

		if count[0] != Distributions {
			t.Errorf("Number %d should appear in all %d trials, but appeared %d times", 0, Distributions, count[0])
		}
	})

	t.Run("Mixed Distribution (-51 to 51)", func(t *testing.T) {
		m, M := -51, 51
		count := make(map[int]int)

		// Before the random number generator enough times
		for i := 0; i < Distributions; i++ {
			result := Int(m, M)
			if result > M || result < m {
				t.Fatalf("Int() = %d; want value in range [%d, %d]", result, m, M)
			}
			count[result]++
		}

		shouldHaveNBuckets(t, count, M-m+1)
		shouldNotDeviateMoreThan(t, count, Distributions, 0.05)
	})
}

func TestRandomInt8Range(t *testing.T) {
	// full range
	t.Run("Full range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt8Interval(t, math.MinInt8, math.MaxInt8)
		}
	})

	// half range
	t.Run("Half range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt8Interval(t, math.MinInt8/2, math.MaxInt8/2)
		}
	})

	// zero range
	t.Run("Zero range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt8Interval(t, 0, 0)
		}
	})
}

func TestRandomInt8Distribution(t *testing.T) {
	// Positive range (0 to 127)
	t.Run("Positive range", func(t *testing.T) {
		var m, M int8 = 0, 127
		count := make(map[int8]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int8(m, M)
			if result > 127 || result < 0 {
				t.Fatalf("Int8() = %d; want value in range [0, 127]", result)
			}
			count[result]++
		}

		shouldHaveNBuckets(t, count, int(M)-int(m)+1)
		shouldNotDeviateMoreThan(t, count, Distributions, 0.05)
	})

	// Negative range (-128 to 0)
	t.Run("Negative range", func(t *testing.T) {
		count := make(map[int8]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int8(-128, 0)
			if result > 0 || result < -128 {
				t.Fatalf("Int8() = %d; want value in range [-128, 0]", result)
			}
			count[result]++
		}

		if len(count) != 129 {
			t.Errorf("Int8() = %d; want value in range [-128, 0]", len(count))
		}

		// Ensure each number appears at least a few times
		for i := -128; i <= 0; i++ {
			if count[int8(i)] == 0 {
				t.Fatalf("Number %d did not appear in 100000 trials", i)
			}
		}
	})

	// Full range
	t.Run("Full range", func(t *testing.T) {
		fullRange := make(map[int8]int)

		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int8()
			if result > 127 || result < -128 {
				t.Fatalf("Int8() = %d; want value in range [-128, 127]", result)
			}
			fullRange[result]++
		}

		if len(fullRange) != 256 {
			t.Errorf("\nRANDOM INT8 SIZE = %d; should be 256\n\n", len(fullRange))
		}

		// Ensure each number appears at least a few times
		for i := -128; i <= 127; i++ {
			if fullRange[int8(i)] == 0 {
				t.Errorf("Number %d did not appear in 100000 trials", i)
			}
		}
	})

	// Zero range
	t.Run("Zero range", func(t *testing.T) {
		count := make(map[int8]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int8(0, 0)
			if result != 0 {
				t.Fatalf("Int8() = %d; want value in range [0, 0]", result)
			}
			count[result]++
		}

		if len(count) != 1 {
			t.Errorf("Int8() = %d; want value in range [0, 0]", len(count))
		}

		// Ensure each number appears at least a few times
		if count[int8(0)] != Distributions {
			t.Errorf("Number %d should appear in all %d trials, but appeared %d times", 0, Distributions, count[int8(0)])
		}
	})
}

func TestRandomInt16Range(t *testing.T) {
	// Positive range
	t.Run("Positive range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt16Interval(t, 0, 1000)
		}
	})

	// Negative range
	t.Run("Negative range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt16Interval(t, -1000, 0)
		}
	})

	// Zero range
	t.Run("Zero range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt16Interval(t, 0, 0)
		}
	})

	// Mixed range
	t.Run("Mixed range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt16Interval(t, -1000, 1000)
		}
	})

	// High edge range
	t.Run("High edge range", func(t *testing.T) {
		m := int16(math.MaxInt16 - 149)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt16Interval(t, m, math.MaxInt16)
		}
	})

	// Low edge range
	t.Run("Low edge range", func(t *testing.T) {
		m, M := int16(math.MinInt16), int16(math.MinInt16+149)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt16Interval(t, m, M)
		}
	})
}

func TestRandomInt16Distribution(t *testing.T) {
	// Positive range
	t.Run("Positive range (0 - 250)", func(t *testing.T) {
		count := make(map[int16]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int16(0, 250)
			if result > 250 || result < 0 {
				t.Fatalf("Int16() = %d; want value in range [0, 250]", result)
			}
			count[result]++
		}

		if len(count) != 251 {
			t.Errorf("Int16() = %d; want value in range [0, 250]", len(count))
		}

		// Ensure each number appears at least a few times
		for i := 0; i <= 250; i++ {
			if count[int16(i)] == 0 {
				t.Fatalf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Negative range
	t.Run("Negative range (-250 - 0)", func(t *testing.T) {
		count := make(map[int16]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int16(-250, 0)
			if result > 0 || result < -250 {
				t.Fatalf("Int16() = %d; want value in range [-250, 0]", result)
			}
			count[result]++
		}

		if len(count) != 251 {
			t.Errorf("Int16() = %d; want value in range [-250, 0]", len(count))
		}

		// Ensure each number appears at least a few times
		for i := -250; i <= 0; i++ {
			if count[int16(i)] == 0 {
				t.Fatalf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Zero range
	t.Run("Zero range (0 - 0)", func(t *testing.T) {
		count := make(map[int16]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int16(0, 0)
			if result != 0 {
				t.Fatalf("Int16() = %d; want value in range [0, 0]", result)
			}
			count[result]++
		}

		if len(count) != 1 {
			t.Errorf("Int16() = %d; want value in range [0, 0]", len(count))
		}

		// Ensure each number appears at least a few times
		if count[0] != Distributions {
			t.Errorf("Number %d should appear in all %d trials, but appeared %d times", 0, Distributions, count[0])
		}
	})

	// High edge range
	t.Run("High edge range", func(t *testing.T) {
		count := make(map[int16]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int16(math.MaxInt16 - 250)
			if result < math.MaxInt16-250 {
				t.Fatalf("Int16() = %d; want value in range [%d, %d]", result, math.MaxInt16-250, math.MaxInt16)
			}
			count[result]++
		}

		if len(count) != 251 {
			t.Errorf("Int16() = %d; want value in range [0, 250]", len(count))
		}

		// Ensure each number appears at least a few times
		for i := math.MaxInt16 - 250; i <= math.MaxInt16; i++ {
			if count[int16(i)] == 0 {
				t.Fatalf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	t.Run("Low edge range", func(t *testing.T) {
		count := make(map[int16]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int16(math.MinInt16, math.MinInt16+250)
			if result > math.MinInt16+250 || result < math.MinInt16 {
				t.Fatalf("Int16() = %d; want value in range [%d, %d]", result, math.MinInt16, math.MinInt16+250)
			}
			count[result]++
		}

		if len(count) != 251 {
			t.Errorf("Int16() = %d; want value in range [0, 250]", len(count))
		}

		// Ensure each number appears at least a few times
		for i := math.MinInt16; i <= math.MinInt16+250; i++ {
			if count[int16(i)] == 0 {
				t.Fatalf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})
}

func TestRandomInt32Range(t *testing.T) {
	// Positive range
	t.Run("Positive range (0 - 2147483647)", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt32Interval(t, 0, 2147483647)
		}
	})

	// Negative range
	t.Run("Negative range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt32Interval(t, -2147483647, 0)
		}
	})

	t.Run("Zero range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt32Interval(t, 0, 0)
		}
	})

	t.Run("Mixed range", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomInt32Interval(t, -600, 600)
		}
	})
}

func TestRandomInt32Distribution(t *testing.T) {
	// Positive range
	t.Run("Positive range (250 - 411)", func(t *testing.T) {
		count := make(map[int32]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int32(250, 411)
			if result > 411 || result < 250 {
				t.Fatalf("Int32() = %d; want value in range [250, 411]", result)
			}
			count[result]++
		}

		if len(count) != 162 {
			t.Errorf("Int32() = %d; want value in range [250, 411]", len(count))
		}

		// Ensure each number appears at least a few times
		for i := 250; i <= 411; i++ {
			if count[int32(i)] == 0 {
				t.Fatalf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Negative range
	t.Run("Negative range (-411 to -250)", func(t *testing.T) {
		count := make(map[int32]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int32(-411, -250)
			if result > -250 || result < -411 {
				t.Fatalf("Int32() = %d; want value in range [-411, -250]", result)
			}
			count[result]++
		}

		if len(count) != 162 {
			t.Errorf("Int32() = %d; want value in range [-411, -250]", len(count))
		}

		// Ensure each number appears at least a few times
		for i := -411; i <= -250; i++ {
			if count[int32(i)] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Fixed range
	t.Run("Fixed range", func(t *testing.T) {
		count := make(map[int32]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int32(200, 200)
			if result != 200 {
				t.Fatalf("Int32() = %d; want value in range [200, 200]", result)
			}
			count[result]++
		}

		if len(count) != 1 {
			t.Errorf("Int32() = %d; want value in range [200, 200]", len(count))
		}

		// Ensure each number appears at least a few times
		if count[200] == 0 {
			t.Errorf("Number %d should appear in all %d trials, but appeared %d times", 200, Distributions, count[200])
		}
	})

	// Zero range
	t.Run("Zero range", func(t *testing.T) {
		count := make(map[int32]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int32(0, 0)
			if result != 0 {
				t.Fatalf("Int32() = %d; want value in range [0, 0]", result)
			}
			count[result]++
		}

		if len(count) != 1 {
			t.Errorf("Int32() = %d; want value in range [0, 0]", len(count))
		}

		// Ensure each number appears at least a few times
		if count[0] == 0 {
			t.Errorf("Number %d should appear in all %d trials, but appeared %d times", 0, Distributions, count[0])
		}
	})

	// High edge range
	t.Run("High edge range", func(t *testing.T) {
		count := make(map[int32]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int32(math.MaxInt32 - 250)
			if result < math.MaxInt32-250 {
				t.Fatalf("Int32() = %d; want value in range [%d, %d]", result, math.MaxInt32-250, math.MaxInt32)
			}
			count[result]++
		}

		if len(count) != 251 {
			t.Errorf("Int32() = %d; want value in range [%d, %d]", len(count), math.MaxInt32-250, math.MaxInt32)
		}

		// Ensure each number appears at least a few times
		for i := math.MaxInt32 - 250; i <= math.MaxInt32; i++ {
			if count[int32(i)] == 0 {
				t.Fatalf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Low edge range
	t.Run("Low edge range", func(t *testing.T) {
		count := make(map[int32]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Int32(math.MinInt32, math.MinInt32+250)
			if result > math.MinInt32+250 {
				t.Fatalf("Int32() = %d; want value in range [%d, %d]", result, math.MinInt32, math.MinInt32+250)
			}
			count[result]++
		}

		if len(count) != 251 {
			t.Errorf("Int32() = %d; want value in range [%d, %d]", len(count), math.MinInt32, math.MinInt32+250)
		}

		// Ensure each number appears at least a few times
		for i := math.MinInt32 + 250; i >= math.MinInt32; i-- {
			if count[int32(i)] == 0 {
				t.Fatalf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})
}

func TestRandomUintRange(t *testing.T) {
	// Small range
	t.Run("Positive range", func(t *testing.T) {
		m, M := uint(10), uint(51)
		shouldNotOverflowRandomUintInterval(t, m, M)
	})

	// Big range
	t.Run("Negative range", func(t *testing.T) {
		m, M := uint(10), uint(51555)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUintInterval(t, m, M)
		}
	})

	// Fixed Range
	t.Run("Fixed range", func(t *testing.T) {
		m, M := uint(10), uint(10)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUintInterval(t, m, M)
		}
	})

	// Zero range
	t.Run("Zero range", func(t *testing.T) {
		m, M := uint(0), uint(0)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUintInterval(t, m, M)
		}
	})

	// High edge range
	t.Run("High edge range (MaxUint - 149 to MaxUint)", func(t *testing.T) {
		m := uint(math.MaxUint - 149)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUintInterval(t, m, math.MaxUint)
		}
	})

	// Low edge range
	t.Run("Low edge range (0 to 149)", func(t *testing.T) {
		M := uint(149)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUintInterval(t, 0, M)
		}
	})
}

func TestRandomUintDistribution(t *testing.T) {
	// Small range
	t.Run("Small range (464660 to 464710)", func(t *testing.T) {
		m, M := uint(464660), uint(464710)
		count := make(map[uint]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Uint(m, M)
			if result > M || result < m {
				t.Fatalf("Uint() = %d; want value in range [%d, %d]", result, m, M)
			}
			count[result]++
		}

		if len(count) != 51 {
			t.Errorf("Uint() = %d; want value in range [%d, %d]", len(count), m, M)
		}

		// Ensure each number appears at least a few times
		for i := m; i <= M; i++ {
			if count[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Big range
	t.Run("Big range (464660 to 465245)", func(t *testing.T) {
		m, M := uint(464660), uint(465245)
		count := make(map[uint]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Uint(m, M)
			if result > M || result < m {
				t.Fatalf("Uint() = %d; want value in range [%d, %d]", result, m, M)
			}
			count[result]++
		}

		if uint(len(count)) != M-m+1 {
			t.Errorf("\nDistribution size = %d; Should be %d\n\n", len(count), M-m+1)
			// Ensure each number appears at least a few times
			for i := m; i <= M; i++ {
				if count[i] == 0 {
					t.Errorf("Number %d did not appear in %d trials", i, Distributions)
				}
			}
		}
	})

	// Zero range
	t.Run("Zero range (0 to 0)", func(t *testing.T) {
		m, M := uint(0), uint(0)
		count := make(map[uint]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Uint(m, M)
			if result > M || result < m {
				t.Errorf("Uint() = %d; want value in range [%d, %d]", result, m, M)
			}
			count[result]++
		}

		if len(count) != 1 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(count), m, M)
		}

		// Ensure each number appears at least a few times
		if count[0] != Distributions {
			t.Errorf("Number %d should appear in all %d trials, but appeared %d times", 0, Distributions, count[0])
		}
	})

	// High edge range
	t.Run("High edge range (math.MaxUint - 149 to math.MaxUint)", func(t *testing.T) {
		m := uint(math.MaxUint - 149)
		count := make(map[uint]int)
		// Before the random number generator a few times
		for i := 0; i < Distributions; i++ {
			result := Uint(m)
			if result < m {
				t.Fatalf("Uint() = %d; want value in range [%d, %d]", result, m, uint64(math.MaxUint))
			}
			count[result]++
		}

		if len(count) != 150 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(count), m, uint64(math.MaxUint))
		}

		// Ensure each number appears at least a few times
		for i := m; i <= m+149; i++ {
			if count[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
			if i == m+149 {
				// Prevent overflow
				break
			}
		}
	})
}

func TestRandomUint8Range(t *testing.T) {
	// full range
	t.Run("Full range (0 to 255)", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint8Interval(t, 0, 255)
		}
	})

	// half range
	t.Run("Half range (0 to 127)", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint8Interval(t, 0, 127)
		}
	})

	// zero range
	t.Run("Zero range (0 to 0)", func(t *testing.T) {
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint8Interval(t, 0, 0)
		}
	})
}

func TestRandomUint8Distribution(t *testing.T) {
	// Full range
	t.Run("Full range (0 to 255)", func(t *testing.T) {
		// Full range
		fullRange := make(map[uint8]int)

		// Before the random number generator 1000 times
		for i := 0; i < Distributions; i++ {
			result := Uint8()
			if result > 255 {
				t.Fatalf("Uint8() = %d; want value in range [0, 255]", result)
			}
			fullRange[result]++
		}

		if len(fullRange) != 256 {
			t.Errorf("Uint8() = %d; want value in range [0, 255]", len(fullRange))
		}
	})

	// Half range
	t.Run("Half range (0 to 127)", func(t *testing.T) {
		// Half range
		halfRange := make(map[uint8]int)

		// Before the random number generator 1000 times
		for i := 0; i < Distributions; i++ {
			result := Uint8(10, 137)
			if result > 137 || result < 10 {
				t.Fatalf("Uint8() = %d; want value in range [0, 127]", result)
			}
			halfRange[result]++
		}

		if len(halfRange) != 128 {
			t.Errorf("Uint8() = %d; want value in range [0, 127]", len(halfRange))
		}
	})

	// Zero range
	t.Run("Zero range (0 to 0)", func(t *testing.T) {
		// Zero range
		zeroRange := make(map[uint8]int)

		// Before the random number generator 1000 times
		for i := 0; i < Distributions; i++ {
			result := Uint8(0, 0)
			if result != 0 {
				t.Fatalf("Uint8() = %d; want value in range [0, 0]", result)
			}
			zeroRange[result]++
		}

		if len(zeroRange) != 1 {
			t.Errorf("Uint8() = %d; want value in range [0, 0]", len(zeroRange))
		}

		if zeroRange[0] != Distributions {
			t.Errorf("Uint8() = %d; want value in range [0, 0]", zeroRange[0])
		}
	})
}

func TestRandomUint16Range(t *testing.T) {
	// Small Range
	t.Run("Small range (3560 to 3600)", func(t *testing.T) {
		m, M := uint16(3560), uint16(3600)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint16Interval(t, m, M)
		}
	})

	// Big Range
	t.Run("Big range (3560 to 3600)", func(t *testing.T) {
		m, M := uint16(3560), uint16(36000)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint16Interval(t, m, M)
		}
	})

	// Fixed Range
	t.Run("Fixed range (3560 to 3560)", func(t *testing.T) {
		m, M := uint16(3560), uint16(3560)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint16Interval(t, m, M)
		}
	})

	// Zero Range
	t.Run("Zero range (0 to 0)", func(t *testing.T) {
		m, M := uint16(0), uint16(0)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint16Interval(t, m, M)
		}
	})
}

func TestRandomUint16Distribution(t *testing.T) {
	// Small Range
	t.Run("Small range (3560 to 3600)", func(t *testing.T) {
		m, M := uint16(3560), uint16(3600)
		counter := make(map[uint16]int)
		for i := 0; i < Distributions; i++ {
			result := Uint16(m, M)
			if result < m || result > M {
				t.Fatalf("Uint16() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 41 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		// Ensure each number appears at least a few times
		for i := m; i <= M; i++ {
			if counter[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Big Range
	t.Run("Big range (3560 to 4000)", func(t *testing.T) {
		m, M := uint16(3560), uint16(4000)
		counter := make(map[uint16]int)
		for i := 0; i < Distributions; i++ {
			result := Uint16(m, M)
			if result < m || result > M {
				t.Fatalf("Uint16() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 441 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		// Ensure each number appears at least a few times
		for i := m; i <= M; i++ {
			if counter[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Fixed Range
	t.Run("Fixed range (3560 to 3560)", func(t *testing.T) {
		m, M := uint16(3560), uint16(3560)
		counter := make(map[uint16]int)
		for i := 0; i < Distributions; i++ {
			result := Uint16(m, M)
			if result < m || result > M {
				t.Fatalf("Uint16() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 1 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		if counter[m] != Distributions {
			t.Errorf("Number %d should appear in all %d trials, but appeared %d times", m, Distributions, counter[m])
		}
	})

	// High Edge Range
	t.Run("High edge range (math.MaxUint16 - 149 to math.MaxUint16)", func(t *testing.T) {
		m := uint16(math.MaxUint16 - 149)
		counter := make(map[uint16]int)
		for i := 0; i < Distributions; i++ {
			result := Uint16(m)
			if result < m {
				t.Fatalf("Uint16() = %d; want value in range [%d, %d]", result, m, uint64(math.MaxUint16))
			}
			counter[result]++
		}

		if len(counter) != 150 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, uint64(math.MaxUint16))
		}

		// Ensure each number appears at least a few times
		for i := math.MaxUint16 - 149; i <= math.MaxUint16; i++ {
			if counter[uint16(i)] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})
}

func TestRandomUint32Range(t *testing.T) {
	// Small Range
	t.Run("Small range (53560 to 53600)", func(t *testing.T) {
		m, M := uint32(53560), uint32(53600)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint32Interval(t, m, M)
		}
	})

	// Big Range
	t.Run("Big range (3560 to 4000)", func(t *testing.T) {
		m, M := uint32(53560), uint32(54000)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint32Interval(t, m, M)
		}
	})

	// Fixed Range
	t.Run("Fixed range (53560 to 53560)", func(t *testing.T) {
		m, M := uint32(53560), uint32(53560)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint32Interval(t, m, M)
		}
	})

	// Zero Range
	t.Run("Zero range (0 to 0)", func(t *testing.T) {
		m, M := uint32(0), uint32(0)
		for i := 0; i < Range; i++ {
			shouldNotOverflowRandomUint32Interval(t, m, M)
		}
	})
}

func TestRandomUint32Distribution(t *testing.T) {
	// Small Range
	t.Run("Small range (53560 to 53600)", func(t *testing.T) {
		m, M := uint32(53560), uint32(53600)
		counter := make(map[uint32]int)
		for i := 0; i < Distributions; i++ {
			result := Uint32(m, M)
			if result < m || result > M {
				t.Fatalf("Uint32() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 41 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		// Ensure each number appears at least a few times
		for i := m; i <= M; i++ {
			if counter[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Big Range
	t.Run("Big range (53560 to 54000)", func(t *testing.T) {
		m, M := uint32(53560), uint32(54000)
		counter := make(map[uint32]int)
		for i := 0; i < Distributions; i++ {
			result := Uint32(m, M)
			if result < m || result > M {
				t.Fatalf("Uint32() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 441 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		// Ensure each number appears at least a few times
		for i := m; i <= M; i++ {
			if counter[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Fixed Range
	t.Run("Fixed range (53560 to 53560)", func(t *testing.T) {
		m, M := uint32(53560), uint32(53560)
		counter := make(map[uint32]int)
		for i := 0; i < Distributions; i++ {
			result := Uint32(m, M)
			if result < m || result > M {
				t.Fatalf("Uint32() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 1 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		// Ensure each number appears at least a few times
		if counter[m] != Distributions {
			t.Errorf("Number %d should appear in all %d trials, but appeared %d times", m, Distributions, counter[m])
		}
	})

	// High Edge Range
	t.Run("High edge range (MaxUint32 - 360 to MaxUint32)", func(t *testing.T) {
		m := uint32(math.MaxUint32 - 360)
		counter := make(map[uint32]int)
		for i := 0; i < Distributions; i++ {
			res := Uint32(m)
			if res < m {
				t.Fatalf("Uint32() = %d; want value in range [%d, %d]", res, m, math.MaxUint32)
			}
			counter[res]++
		}

		if len(counter) != 361 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, math.MaxUint32)
		}

		// Ensure each number appears at least a few times
		for i := math.MaxUint32 - 360; i <= math.MaxUint32; i++ {
			if counter[uint32(i)] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})
}

func TestRandomUint64Range(t *testing.T) {
	// Small Range
	t.Run("Small range (9593560 to 9593600)", func(t *testing.T) {
		m, M := uint64(9593560), uint64(9593600)
		for i := 0; i < Distributions; i++ {
			shouldNotOverflowRandomUint64Interval(t, m, M)
		}
	})

	// Big Range
	t.Run("Big range (9573560 to 9574000)", func(t *testing.T) {
		m, M := uint64(9573560), uint64(9574000)
		for i := 0; i < Distributions; i++ {
			shouldNotOverflowRandomUint64Interval(t, m, M)
		}
	})

	// Fixed Range
	t.Run("Fixed range (9573560 to 9573560)", func(t *testing.T) {
		m, M := uint64(9573560), uint64(9573560)
		for i := 0; i < Distributions; i++ {
			shouldNotOverflowRandomUint64Interval(t, m, M)
		}
	})

	// Zero Range
	t.Run("Zero range (0 to 0)", func(t *testing.T) {
		m, M := uint64(0), uint64(0)
		for i := 0; i < Distributions; i++ {
			shouldNotOverflowRandomUint64Interval(t, m, M)
		}
	})
}

func TestRandomUint64Distribution(t *testing.T) {
	// Small Range
	t.Run("Small range (9593560 to 9593600)", func(t *testing.T) {
		m, M := uint64(9593560), uint64(9593600)
		counter := make(map[uint64]int)
		for i := 0; i < Distributions; i++ {
			result := Uint64(m, M)
			if result < m || result > M {
				t.Fatalf("Uint64() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 41 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		// Ensure each number appears at least a few times
		for i := m; i <= M; i++ {
			if counter[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Big Range
	t.Run("Big range (9573560 to 9574000)", func(t *testing.T) {
		m, M := uint64(9573560), uint64(9574000)
		counter := make(map[uint64]int)
		for i := 0; i < Distributions; i++ {
			result := Uint64(m, M)
			if result < m || result > M {
				t.Fatalf("Uint64() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 441 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		// Ensure each number appears at least a few times
		for i := m; i <= M; i++ {
			if counter[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
		}
	})

	// Fixed Range
	t.Run("Fixed range (9573560 to 9573560)", func(t *testing.T) {
		m, M := uint64(9573560), uint64(9573560)
		counter := make(map[uint64]int)
		for i := 0; i < Distributions; i++ {
			result := Uint64(m, M)
			if result < m || result > M {
				t.Fatalf("Uint64() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 1 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}
	})

	// Zero Range
	t.Run("Zero range (0 to 0)", func(t *testing.T) {
		m, M := uint64(0), uint64(0)
		counter := make(map[uint64]int)
		for i := 0; i < Distributions; i++ {
			result := Uint64(m, M)
			if result < m || result > M {
				t.Fatalf("Uint64() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 1 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}
	})

	// High Edge Range
	t.Run("High edge range (math.MaxUint64 - 260 to math.MaxUint64)", func(t *testing.T) {
		m, M := uint64(math.MaxUint64-260), uint64(math.MaxUint64)
		counter := make(map[uint64]int)
		for i := 0; i < Distributions; i++ {
			result := Uint64(m, M)
			if result < m || result > M {
				t.Fatalf("Uint64() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[result]++
		}

		if len(counter) != 261 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}

		// Ensure each number appears at least a few times
		for i := m; i <= uint64(math.MaxUint64); i++ {
			if counter[i] == 0 {
				t.Errorf("Number %d did not appear in %d trials", i, Distributions)
			}
			if i == uint64(math.MaxUint64) {
				break
			}
		}
	})
}

func TestRandomBoolDistribution(t *testing.T) {
	counter := make(map[bool]int)
	for i := 0; i < Distributions; i++ {
		result := Bool()
		counter[result]++
	}

	// Ensure both true and false were generated
	if len(counter) != 2 {
		t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), 2, 2)
	}

	shouldNotDeviateMoreThan(t, counter, Distributions, 0.05)
}

func TestRandomBytesRange(t *testing.T) {
	// Small Range
	t.Run("Small range (0 to 10)", func(t *testing.T) {
		m, M := 0, 10
		for i := 0; i < Range; i++ {
			result := Bytes(m, M)
			if len(result) < m || len(result) > M {
				t.Fatalf("RandomByte() = %d; want value in range [%d, %d]", result, m, M)
			}
		}
	})

	// Big Range
	t.Run("Big range (0 to 100)", func(t *testing.T) {
		m, M := 50, 200
		for i := 0; i < Range; i++ {
			result := Bytes(m, M)
			if len(result) < m || len(result) > M {
				t.Fatalf("RandomByte() = %d; want value in range [%d, %d]", result, m, M)
			}
		}
	})

	// Fixed Range
	t.Run("Fixed range (10 to 10)", func(t *testing.T) {
		m, M := 10, 10
		for i := 0; i < Range; i++ {
			result := Bytes(m, M)
			if len(result) < m || len(result) > M {
				t.Fatalf("RandomByte() = %d; want value in range [%d, %d]", result, m, M)
			}
		}
	})
}

func TestRandomBytesDistribution(t *testing.T) {
	// Small Range
	t.Run("Small range (0 to 10)", func(t *testing.T) {
		m, M := 0, 10
		counter := make(map[int]int)
		for i := 0; i < Distributions; i++ {
			result := Bytes(m, M)
			if len(result) < m || len(result) > M {
				t.Fatalf("RandomByte() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[len(result)]++
		}

		if len(counter) != M-m+1 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}
	})

	// Big Range
	t.Run("Big range (0 to 100)", func(t *testing.T) {
		m, M := 50, 200
		counter := make(map[int]int)
		for i := 0; i < Distributions; i++ {
			result := Bytes(m, M)
			if len(result) < m || len(result) > M {
				t.Fatalf("RandomByte() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[len(result)]++
		}

		if len(counter) != M-m+1 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}
	})

	// Fixed Range
	t.Run("Fixed range (10 to 10)", func(t *testing.T) {
		m, M := 10, 10
		counter := make(map[int]int)
		for i := 0; i < Distributions; i++ {
			result := Bytes(m, M)
			if len(result) < m || len(result) > M {
				t.Fatalf("RandomByte() = %d; want value in range [%d, %d]", result, m, M)
			}
			counter[len(result)]++
		}

		if len(counter) != M-m+1 {
			t.Errorf("Distribution size = %d; want value in range [%d, %d]", len(counter), m, M)
		}
	})
}

func TestRandomFloat32Range(t *testing.T) {
	// Positive Range
	t.Run("Positive range (5550.3 to 80900.8)", func(t *testing.T) {
		var m, M float32 = 5550.3, 80900.8
		for i := 0; i < Range; i++ {
			result := Float32(m, M)
			if result < m || result > M {
				t.Fatalf("Float32() = %f; want value in range [%f, %f]", result, m, M)
			}
		}
	})

	// Negative Range
	t.Run("Negative range (-5550.3 to -80900.8)", func(t *testing.T) {
		var m, M float32 = -80900.8, -5550.3
		for i := 0; i < Range; i++ {
			result := Float32(m, M)
			if result < m || result > M {
				t.Fatalf("Float32() = %f; want value in range [%f, %f]", result, m, M)
			}
		}
	})

	// Fixed Range
	t.Run("Fixed range (10.8 to 10.8)", func(t *testing.T) {
		var m, M float32 = 10.8, 10.8
		for i := 0; i < Range; i++ {
			result := Float32(m, M)
			if result != m {
				t.Fatalf("Float32() = %f; want value in range [%f, %f]", result, m, M)
			}
		}
	})
}

func TestRandomFloat32Distribution(t *testing.T) {
	t.Run("Small range and distribution [0, 1]", func(t *testing.T) {
		var m, M float32 = 0.0, 2.0
		bucketSize := 50
		buckets := make(map[uint64]uint64, bucketSize)
		factor := (M - m) / float32(bucketSize)

		for i := 0; i < 1_000_000; i++ {
			result := Float32(m, M)
			if result < m || result > M {
				t.Fatalf("[%d]Float64 = %f; want value in range [%f, %f)", i, result, m, M)
			}

			buckets[uint64(result/factor)]++
		}

		//var _, higher uint64
		//var _, lower uint64 = 0, math.MaxUint64
		//for _, v := range buckets {
		//	if v > higher {
		//		higher = v
		//hid = k
		//}
		//if v < lower {
		//	lower = v
		//lid = k
		//	}
		//}
		//fmt.Printf("Higher[%d]: %d\n", hid, higher)
		//fmt.Printf("Lower[%d]: %d\n", lid, lower)
		//fmt.Printf("Clocks: %v\n", buckets)
	})

	// This test doesn't work very well because
	// floats of 32 bytes lose is precision after 23 bits => Approximately 7 decimal digits of precision
	// to make these calculations properly. But the distribution still working
	t.Run("High edge range [MaxFloat32 - 20000000000000000000000000000000.0, MaxFloat32]", func(t *testing.T) {
		var m, M float32 = math.MaxFloat32 - 20000000000000000000000000000000.0, math.MaxFloat32
		bucketSize := 50
		buckets := make(map[float32]uint64, bucketSize)
		factor := (M - m) / float32(bucketSize)

		for i := 0; i < Distributions; i++ {
			result := Float32(m)
			if result < m || result > M {
				t.Fatalf("Float64() = %f; want value in range [%f, %f)", result, m, M)
			}
			// Determine the bucket
			buckets[result/factor]++
		}

		//var lower, higher uint64
		//var lid, hid float32 = 0, math.MaxFloat32
		//for _, v := range buckets {
		//	if v > higher {
		//		higher = v
		//hid = k
		//}
		//if v < lower {
		//	lower = v
		//lid = k
		//}
		//}
		//fmt.Printf("Higher[%f]: %d\n", hid, higher)
		//fmt.Printf("Lower[%f]: %d\n", lid, lower)
		//fmt.Printf("Clocks: %v\n", buckets)
	})
}
