// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package subsets

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/sortables/sorts"
)

// Xor returns all the subsets of a set of size n with k elements
// (i.e. all possible combinations of k elements in a set of size n)
func Xor(min, max, k int) [][]int {
	var subsets [][]int
	n := max - min

	// Loop over all possible combinations of k elements in a set of size n
	for subset := (1 << k) - 1; subset < (1 << n); {
		// Convert the current subset (in binary) to a list of elements, offset by min
		var currentSubset []int
		for i := 0; i < n; i++ {
			if subset&(1<<i) != 0 {
				currentSubset = append(currentSubset, min+i)
			}
		}
		subsets = append(subsets, currentSubset)

		// Calculate the next combination using bitwise operations
		c := subset & -subset
		r := subset + c
		subset = (((r ^ subset) >> 2) / c) | r
	}

	return subsets
}

func Or(min, max, k int) [][]int {
	combinations := Xor(min, max, k)
	var allPermutations [][]int
	for _, combination := range combinations {
		permutations := Permute(combination)
		allPermutations = append(allPermutations, permutations...)
	}
	return allPermutations
}

func Permute(arr []int) [][]int {
	var result [][]int
	var permute func([]int, int)
	permute = func(arr []int, i int) {
		if i == len(arr) {
			// Make a copy of the array and append it to result
			perm := make([]int, len(arr))
			copy(perm, arr)
			result = append(result, perm)
			return
		}
		for j := i; j < len(arr); j++ {
			arr[i], arr[j] = arr[j], arr[i]
			permute(arr, i+1)
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	permute(arr, 0)
	return result
}

func XorFilter(min, max, k int, filter functions.BiPredicate2[[]int, []int]) [][]int {
	var subsets [][]int
	n := max - min

	for subset := (1 << k) - 1; subset < (1 << n); {
		var currentSubset []int
		var excludedSubset []int

		// Create the subsets
		for i := 0; i < n; i++ {
			if subset&(1<<i) != 0 {
				currentSubset = append(currentSubset, min+i)
			} else {
				excludedSubset = append(excludedSubset, min+i)
			}
		}

		// Apply the filter function
		include, stop := filter(excludedSubset, currentSubset)
		if include {
			subsets = append(subsets, currentSubset)
		}
		if stop {
			break
		}

		// Calculate the next combination
		c := subset & -subset
		r := subset + c
		subset = (((r ^ subset) >> 2) / c) | r
	}

	return subsets
}

func XorFilterOf[T prim.Ordered](arr []T, min, max, k int, filter func(excl, incl []int) (bool, bool)) [][]int {
	var subsets [][]int
	n := max - min

	for subset := (1 << k) - 1; subset < (1 << n); {
		var currentSubset []int
		var excludedSubset []int

		// Create the subsets
		for i := 0; i < n; i++ {
			if subset&(1<<i) != 0 {
				currentSubset = append(currentSubset, min+i)
			} else {
				excludedSubset = append(excludedSubset, min+i)
			}
		}

		// Apply the filter function
		include, stop := filter(excludedSubset, currentSubset)
		if include {
			subsets = append(subsets, currentSubset)
		}
		if stop {
			break
		}

		// Calculate the next combination
		c := subset & -subset
		r := subset + c
		subset = (((r ^ subset) >> 2) / c) | r
	}

	return subsets
}

// Of returns the elements in arr that aren't in indexes and the
// elements in arr that are in indexes in different slices.
func Of[T any](arr []T, indexes ...int) ([]T, []T) {
	if len(indexes) == 0 {
		return arr, nil
	}

	// Sorter the indexes to ensure they are in ascending order
	sorts.Quick[int](indexes)

	// Allocate slices with the correct size
	excl := make([]T, 0, len(arr)-len(indexes))
	incl := make([]T, len(indexes))

	// Track the current index in the original array
	currIdx := 0

	for i, idx := range indexes {
		// Append the elements before the current index to the excluded slice
		excl = append(excl, arr[currIdx:idx]...)
		// Addf the current index element to the included slice
		incl[i] = arr[idx]
		// Move the current index forward
		currIdx = idx + 1
	}

	// Append any remaining elements to the excluded slice
	excl = append(excl, arr[currIdx:]...)

	return excl, incl
}

//func OperatorAggregator[T prim.Ordered](arr []T, filter func(arr []T, left, right T) bool) []*tuple.Pair[T, T] {
//	var filtered []*tuple.Pair[T, T]
//	op := make([]T, len(arr)-2)
//	sen := sentinels.Double(0, 1)
//
//	for sen.Right.Less(len(arr)) {
//		op, _ = Of(arr, sen.Left.Val(), sen.Right.Val())
//		if filter(op, arr[sen.Left.Val()], arr[sen.Right.Val()]) {
//			filtered = append(filtered, tuple.NewPair(arr[sen.Left.Val()], arr[sen.Right.Val()]))
//			sen.Left.Next()
//			sen.Right.Next()
//		} else {
//			sen.Right.Next()
//		}
//	}
//	return filtered
//}
//
//func ExclusiveOperatorAggregator[T prim.Ordered](arr []T, filter func(arr []T, left, right T) bool) []*tuple.Pair[T, T] {
//	var filtered []*tuple.Pair[T, T]
//	op := make([]T, len(arr)-2)
//	sen := sentinels.Double(0, 1)
//
//	for sen.Right.Less(len(arr)) {
//		op, _ = Of(arr, sen.Left.Val(), sen.Right.Val())
//		if filter(op, arr[sen.Left.Val()], arr[sen.Right.Val()]) {
//			filtered = append(filtered, tuple.NewPair(arr[sen.Left.Val()], arr[sen.Right.Val()]))
//			sen.Left.Next()
//			sen.Right.Next()
//		} else {
//			sen.Right.Next()
//		}
//	}
//	return filtered
//}
//
//func Aggregator[T any](arr []T, filter func(left, right T) bool) []*tuple.Pair[T, T] {
//	var filtered []*tuple.Pair[T, T]
//	sen := sentinels.Double(0, 1)
//	for sen.Right.Less(len(arr)) {
//		if filter(arr[sen.Left.Val()], arr[sen.Right.Val()]) {
//			filtered = append(filtered, tuple.NewPair(arr[sen.Left.Val()], arr[sen.Right.Val()]))
//			sen.Left.Next()
//			sen.Right.Next()
//		} else {
//			sen.Right.Next()
//		}
//	}
//	return filtered
//}
