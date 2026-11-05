// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tests

import (
	"github.com/andrerrcosta2/gtools/core/data/str/tuple"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// EqualSlicesFunc compares two slices using a comparison function and returns the mismatches
func EqualSlicesFunc[S ~[]E, E any](
	slice1 S,
	slice2 S,
	compare function.Equals[E],
) (equal bool, mismatches []*tuple.Pair[E, E]) {
	// Check if both slices are nil or both empty
	if len(slice1) == 0 && len(slice2) == 0 {
		return true, nil
	}

	// If lengths differ, consider all elements as mismatches (up to the shorter length)
	maxLen := len(slice1)
	if len(slice2) > maxLen {
		maxLen = len(slice2)
	}

	mismatches = make([]*tuple.Pair[E, E], 0)

	for i := 0; i < maxLen; i++ {
		var a, b E
		var aExists, bExists bool

		if i < len(slice1) {
			a = slice1[i]
			aExists = true
		}
		if i < len(slice2) {
			b = slice2[i]
			bExists = true
		}

		// If only one exists at this position, it's a mismatch
		if !aExists || !bExists {
			var zeroA, zeroB E
			if !aExists && bExists {
				// Only second has value
				mismatches = append(mismatches, tuple.NewPair(zeroA, b))
			} else if aExists && !bExists {
				// Only first has value
				mismatches = append(mismatches, tuple.NewPair(a, zeroB))
			}
			continue
		}

		if !compare(a, b) {
			mismatches = append(mismatches, tuple.NewPair(a, b))
		}
	}

	return len(mismatches) == 0, mismatches
}
