// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/str/arrays"
	"github.com/andrerrcosta2/gtools/core/data/str/matrices"
	"github.com/andrerrcosta2/gtools/core/sortables"
)

// backtrackBranchable inserts a branchable value backtracked from the given branchable.
// It will insert the branchable value and all its parent rows into the given slice until
// a branch that is not a comparable value, or if it was already visited, or if it is the same
// as its own parent.
//
// A branchable value is a value that implements the data.Branchable interface.
// The branch must be a comparable value.
func backtrackBranchable[B data.Branchable[B]](branchable B, action func(B)) {
	// Track visited rows to avoid infinite cycles.
	var visited = make(map[string]bool)
	var next = branchable

	for {
		hash := sortables.Unique[B](next)
		// Check if the branch was already visited
		if visited[hash] {
			break
		}

		// Perform the action (collect the branchable)
		action(next)

		// Mark the branch as visited
		visited[hash] = true

		// Check for backtracking conditions
		nb, ok := next.Branch()
		if !ok || sortables.TryEquality[B](nb, next) {
			break
		}

		// Get the next branch
		// The overhead this kind of type checking adds to the performance
		// is negligible mostly due to golang compiler optimizations
		if next, ok = any(nb).(B); !ok {
			panic("gtools:collectors: recursive collector with different types\n")
		}
	}
}

func addBranchToMatrixIfAbsent[H any, O any, C ~[][]O](heads map[string]int, branches *C, head H) (index int, exists bool) {
	// To dereference here is necessary to rely only on
	// the branch natural comparable properties.
	if ptr, ok := any(head).(*H); ok {
		head = *ptr
	}
	// Get the hash of the head
	hash := sortables.Unique[H](head)
	// Get the index of the head from the map
	index, exists = heads[hash]
	// Add to the matrix if it doesn't exist
	if !exists {
		// Get the next index
		index = len(heads)
		// Store the index in the map
		heads[hash] = index
		// Initialize the rows slice at the new index
		matrices.GrowIfLessThan(branches, index+1)
		// Initialize the rows slice at the new index
		(*branches)[index] = make([]O, 0)
	}
	return
}

func addBranchToSliceIfAbsent[H any, O any, C ~[]O](heads map[string]int, branches *C, head H) (index int, exists bool) {
	// To dereference here is necessary to rely only on
	// the branch natural comparable properties.
	if ptr, ok := any(head).(*H); ok {
		head = *ptr
	}
	// Get the hash of the head
	hash := sortables.Unique[H](head)
	// Get the index of the head from the map
	index, exists = heads[hash]
	// Add to the matrix if it doesn't exist
	if !exists {
		// Get the next index
		index = len(heads)
		// Store the index in the map
		heads[hash] = index
		// Initialize the rows slice at the new index
		arrays.GrowIfLessThan(branches, index+1)
	}
	return
}
