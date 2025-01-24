// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package random

import (
	"github.com/andrerrcosta2/gtools/core/gtools/constraints/prim/nums"
	"math"
	"time"
)

func xrng[T nums.Integer](minMax ...T) (min, max T) {
	switch (interface{})(*new(T)).(type) {
	case int8:
		min = any(math.MinInt8).(T)
		max = T(math.MaxInt8)
	case int16:
		min, max = any(math.MinInt16).(T), any(math.MaxInt16).(T)
	case int32:
		min, max = any(math.MinInt32).(T), any(math.MaxInt32).(T)
	case int64, int:
		min, max = any(math.MinInt64).(T), any(math.MaxInt64).(T)
	}

	switch len(minMax) {
	case 1:
		min = minMax[0]
	case 2:
		min, max = minMax[0], minMax[1]
		if min > max {
			min, max = max, min
		}
	}
	return
}

func xurng[T nums.Natural](minMax ...T) (min, max uint64) {
	var zero T
	switch any(zero).(type) {
	case uint8:
		min, max = 0, uint64(math.MaxUint8)
	case uint16:
		min, max = 0, uint64(math.MaxUint16)
	case uint32:
		min, max = 0, uint64(math.MaxUint32)
	case uint64, uintptr, uint:
		min, max = 0, uint64(math.MaxUint64)
	}

	switch len(minMax) {
	case 1:
		min = uint64(minMax[0])
	case 2:
		min, max = uint64(minMax[0]), uint64(minMax[1])
		if min > max {
			min, max = max, min
		}
	}
	return
}

func xrngf32(minMax ...float32) (min, max float32) {
	switch len(minMax) {
	case 0:
		// No m or M provided, use default range
		min, max = -math.MaxFloat32, math.MaxFloat32
	case 1:
		// Only m provided, set M to a large value
		min, max = minMax[0], math.MaxFloat32
	default:
		// Both m and M provided
		min, max = minMax[0], minMax[1]
		// Ensure m is less than M
		if min > max {
			min, max = max, min
		}
	}
	return
}

func xrngf64(minMax ...float64) (min, max float64) {
	switch len(minMax) {
	case 0:
		// No m or M provided, use default range
		// float64 is a 64-bit IEEE 754 floating-point number.
		// Its range is symmetric, meaning the negative and positive sides
		// of the range have the same magnitude
		min, max = -math.MaxFloat64, math.MaxFloat64
	case 1:
		// Only m provided, set M to a large value
		min, max = minMax[0], math.MaxFloat64
	default:
		// Both m and M provided
		min, max = minMax[0], minMax[1]
		// Ensure m is less than M
		if min > max {
			min, max = max, min
		}
	}
	return
}

func xtimergn(fromTo ...time.Time) (from, to time.Time) {
	switch len(fromTo) {
	case 0:
		// No "from" or "to" provided, use default range
		from = time.Unix(0, 0)
		to = time.Date(9999, 12, 31, 23, 59, 59, 999999999, time.UTC)
	case 1:
		// Only from provided, set "to" to a large value
		from = fromTo[0]
		to = time.Date(9999, 12, 31, 23, 59, 59, 999999999, time.UTC)
	default:
		// Both from and to provided
		from, to = fromTo[0], fromTo[1]
		// Ensure "from" is before "to"
		if from.After(to) {
			from, to = to, from // Swap if "from" is after "to"
		}
	}
	return
}
