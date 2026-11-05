// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package list

import "github.com/andrerrcosta2/gtools/core/data/str/node"

type Readable[T any] interface {
	Contains(value T) bool // Returns true if the list contains the specified element.
	Size() int             // Returns the number of elements in the list.
	ToSlice() []T          // Returns a slice containing all elements in order.
}

type Immutable[T any] interface {
	Readable[T]
	Get(index int) (T, bool) // Returns the element at the specified position.
	IndexOf(value T) int     // Returns the index of the first occurrence, or -1 if not found.
	IsEmpty() bool           // Returns true if the list contains no elements.
}

// Mutable interface
type Mutable[T any] interface {
	Immutable[T]
	Add(value T) int                // Appends the specified element to the end of the list and retrieve its index.
	HardClear()                     // Explicitly clears the list, breaking all links
	Insert(index int, value T) bool // Inserts an element at the specified index.
	Remove(value T) bool            // Removes all occurrences of the specified element.
	RemoveAt(index int) (T, bool)   // Removes the element at the specified position.
	Set(index int, value T) bool    // Replaces the element at the specified position.
	SoftClear()                     // Soft reset, relies on GC to clean up
}

type GraphNode[N node.Graph[X], X any] interface {
	All() []N
	Iter(f func(N) bool)
}
