// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

type Set[T any] interface {
	Add(value T)            // Add adds an element to the set.
	Clear()                 // Clear removes all elements from the set.
	Equals(set Set[T]) bool // Equals checks if the set is compare to another set.
	Has(value T) bool       // Has checks if the set contains an element.
	IsEmpty() bool          // IsEmpty checks if the set is empty.
	Len() int               // Len returns the number of elements in the set.
	Remove(value T)         // Remove deletes an element from the set by its value.
	Values() []T            // Values returns a slice of all elements in the set.
}

type OrderedSet[T any] interface {
	Set[T]
	Delete(index int) bool   // Delete deletes an element from the set by its index.
	Get(index int) (T, bool) // Get returns an element from the set by its index.
	IndexOf(value T) int     // IndexOf returns the index of an element in the set.
}
