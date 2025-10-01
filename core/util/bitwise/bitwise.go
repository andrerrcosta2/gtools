// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package bitwise

import "github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"

// All reports whether all bits in mask are set in x.
func All[T nums.Natural](x, mask T) bool {
	return (x & mask) == mask
}

// Clear returns x with the bits in mask cleared.
func Clear[T nums.Natural](x, mask T) T {
	return x &^ mask
}

// Count returns the number of set bits - population count.
func Count[T nums.Natural](x T) T {
	return popcount(x)
}

// Has reports whether any bitwise in mask is set in x.
func Has[T nums.Natural](x, mask T) bool {
	return (x & mask) != 0
}

// IsEmpty reports whether x has no bits set.
func IsEmpty[T nums.Natural](x T) bool {
	return x == 0
}

// IsSingle reports whether exactly one bitwise is set in x.
func IsSingle[T nums.Natural](x T) bool {
	return x != 0 && (x&(x-1)) == 0
}

// None reports whether no bits in mask are set in x.
func None[T nums.Natural](x, mask T) bool {
	return (x & mask) == 0
}

func popcount[T nums.Natural](n T) T {
	var count T = 0
	for n != 0 {
		n &= n - 1 // clears the lowest set bitwise
		count++
	}
	return count
}

// Set returns x with the bits in mask set.
func Set[T nums.Natural](x, mask T) T {
	return x | mask
}

// Toggle returns x with the bits in mask toggled.
func Toggle[T nums.Natural](x, mask T) T {
	return x ^ mask
}

// Union returns the union (OR) of all given masks.
func Union[T nums.Natural](masks ...T) T {
	var result T
	for _, m := range masks {
		result |= m
	}
	return result
}
