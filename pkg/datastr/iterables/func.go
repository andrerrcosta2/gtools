// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterables

import "github.com/andrerrcosta2/gtools/pkg/comparables"

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
func DetectCycle[T any](head Iterable[T], comparator comparables.Comparator[T]) Iterable[T] {
	if head == nil || !head.HasNext() {
		return nil // No cycle if the list is empty or has only one element
	}

	tortoise := head
	hare := head

	// Detect if a cycle exists
	for hare.HasNext() {
		tortoise = tortoise.Next() // 1/2
		if !hare.Next().HasNext() {
			return nil // No cycle if hare can't move forward safely
		}

		// if there's a cycle, return the node where the cycle starts
		if comparator.Equals(tortoise.Get(), hare.Get()) {
			tortoise = head                                      // Move tortoise to the start
			for !comparator.Equals(tortoise.Get(), hare.Get()) { // Move both pointers by 1 step
				tortoise = tortoise.Next()
				hare = hare.Next()
			}
			return tortoise
		}
	}

	return nil // No cycle found
}
