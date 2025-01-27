// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterables

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/comparables"
)

// DetectCycle detects if a list has a cycle.
//
// If the list has a cycle, it returns the node where the cycle starts.
// If the list does not have a cycle, it returns nil.
// If the list is empty or has only one element, it also returns nil.
//
// The algorithm used is the Floyd's Cycle Detection Algorithm.
//
// This function is a generic function, which means it works with any type T.
//
// The time complexity of this function is O(n) where n is the length of the list.
// The space complexity is O(1) as no additional data structures are used.
func DetectCycle[T any](head data.Iterator[T], comparator comparables.Comparator[T]) T {
	var zero T
	// If the list is empty or has only one element, return nil
	if head == nil || !head.HasNext() {
		return zero // No cycle if the list is empty or has only one element
	}

	tortoise := head
	hare := head

	// Detect if a cycle exists
	for hare.HasNext() {
		tortoise = head // Move tortoise forward by 1
		tortoiseValue := tortoise.Next()

		hareValue := hare.Next() // 1/2
		if !hare.HasNext() {
			return zero
		}
		hareValue = hare.Next()

		// If there's a cycle, find the starting point of the cycle
		if comparator.Equals(tortoiseValue, hareValue) {
			tortoise = head // Move tortoise to the start
			for !comparator.Equals(tortoise.Next(), hareValue) {
				tortoiseValue = tortoise.Next()
				hareValue = hare.Next()
			}
			return tortoiseValue
		}
	}

	return zero
}
