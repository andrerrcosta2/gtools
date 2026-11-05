// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package set

type Readable[T any] interface {
	IsEmpty() bool                     // IsEmpty checks if the set is empty.
	Equals(other Readable[T]) bool     // Equals checks if the set is compare to another set.
	Has(value T) bool                  // Has checks if the set contains an element.
	Len() int                          // Len returns the number of elements in the set.
	Range(callback func(value T) bool) // Range iterates over all values from the set
	Elements() []T                     // Values returns a slice of all elements in the set.
}

type ReadableOrdered[T any] interface {
	Get(index int) (T, bool) // Get returns an element from the set by its index.
}

type Writable[T any] interface {
	Add(value T)    // Add adds an element to the set.
	Clear()         // Clear removes all elements from the set.
	Remove(value T) // Remove deletes an element from the set by its value.
}

type Mutable[T any] interface {
	Readable[T]
	Writable[T]
}

// WritableOrdered is a Mutable set with all its elements ordered
type WritableOrdered[T any] interface {
	Delete(index int) bool   // Delete deletes an element from the set by its index.
	Get(index int) (T, bool) // Get returns an element from the set by its index.
	IndexOf(value T) int     // IndexOf returns the index of an element in the set.
}

type MutableOrdered[T any] interface {
	ReadableOrdered[T]
	WritableOrdered[T]
}
