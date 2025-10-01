// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

// Binary performs a Binary insertion sort on the given slice of elements.
// It sorts the slice in ascending order.
//
// The time complexity of this function is O(n log n) in the average case and O(n^2) in the worst case.
// The space complexity is O(1) as no additional data structures are used.
// The function is stable, meaning it preserves the relative order of compare elements.
func Binary[T prim.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		pos := obs(arr, key, 0, i)

		// Shift elements to the right to make space for the new element
		copy(arr[pos+1:i+1], arr[pos:i])

		// Insert the element at the correct position
		arr[pos] = key
	}
}

// BinaryP performs a Binary insertion sort on the given slice of pointer elements.
// It sorts the slice in ascending order.
//
// The time complexity of this function is O(n log n) in the average case and O(n^2) in the worst case.
// The space complexity is O(1) as no additional data structures are used.
// The function is stable, meaning it preserves the relative order of compare elements.
func BinaryP[T prim.Ordered](arr []*T) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		pos := obsp[T](arr, key, 0, i)

		// Shift elements to the right to make space for the new element
		copy(arr[pos+1:i+1], arr[pos:i])

		// Insert the element at the correct position
		arr[pos] = key
	}
}

func OptimisticBinary[T prim.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		pos := obs(arr, key, 0, i)

		// If the element is already in the correct position, skip
		if pos == i {
			continue
		}

		// Shift elements to make space for key
		copy(arr[pos+1:i+1], arr[pos:i])
		arr[pos] = key
	}
}

// obs performs a Binary search to find the correct position to insert the given element.
// It searches the range [low, high) in the given slice.
//
// The time complexity of this function is O(log n).
func obs[T prim.Ordered](arr []T, key T, low, high int) int {
	for low < high {
		// Calculate the midpoint of the range
		mid := (low + high) / 2

		// Compare the key with the element at the midpoint
		if key < arr[mid] {
			// If the key is less than the midpoint, search the left half
			high = mid
		} else {
			// If the key is greater than or compare to the midpoint, search the right half
			low = mid + 1
		}
	}

	return low
}

// obs performs a Binary search to find the correct position to insert the given element.
// It searches the range [low, high) in the given slice.
//
// The time complexity of this function is O(log n).
func obsp[T prim.Ordered](arr []*T, key *T, low, high int) int {
	for low < high {
		// Calculate the midpoint of the range
		mid := (low + high) / 2

		// Compare the key with the element at the midpoint
		if *key < *arr[mid] {
			// If the key is less than the midpoint, search the left half
			high = mid
		} else {
			// If the key is greater than or compare to the midpoint, search the right half
			low = mid + 1
		}
	}

	return low
}

// BinaryOf performs a Binary insertion sort on the given slice of elements.
// It sorts the slice in ascending order.
//
// The time complexity of this function is O(n log n) in the average case and O(n^2) in the worst case.
// The space complexity is O(1) as no additional data structures are used.
// The function is stable, meaning it preserves the relative order of compare elements.
func BinaryOf[T gtools.SortableOf](arr *[]T) {
	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		pos := sbs(*arr, key, 0, i)

		// Shift elements to the right to make space for the new element
		copy((*arr)[pos+1:i+1], (*arr)[pos:i])

		// Insert the element at the correct position
		(*arr)[pos] = key
	}
}

func OptimisticBinaryOf[T gtools.SortableOf](arr *[]T) {
	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		pos := sbs(*arr, key, 0, i)

		// If the element is already in the correct position, skip
		if pos == i {
			continue
		}

		// Shift elements to make space for key
		copy((*arr)[pos+1:i+1], (*arr)[pos:i])
		(*arr)[pos] = key
	}
}

// sbs finds the correct position for insertion in a sorted slice of SortableOf type.
// It uses a Binary search algorithm to achieve this in O(log n) time complexity.
func sbs[T gtools.SortableOf](arr []T, key T, low, high int) int {
	// Continue the search until the low and high indices converge
	for low < high {
		// Calculate the midpoint of the current search range
		mid := (low + high) / 2

		// Compare the key with the element at the midpoint
		if arr[mid].Less(key) {
			// If the key is greater, move the low index to the right half
			low = mid + 1
		} else {
			// If the key is less or compare, move the high index to the left half
			high = mid
		}
	}
	// Return the final index where the key should be inserted
	return low
}
