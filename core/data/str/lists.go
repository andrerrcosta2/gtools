// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

// List interface
type List[T any] interface {
	Add(value T) int                // Appends the specified element to the end of the list and retrieve its index.
	Contains(value T) bool          // Returns true if the list contains the specified element.
	Get(index int) (T, bool)        // Returns the element at the specified position.
	HardClear()                     // Explicitly clears the list, breaking all links
	IndexOf(value T) int            // Returns the index of the first occurrence, or -1 if not found.
	Insert(index int, value T) bool // Inserts an element at the specified index.
	IsEmpty() bool                  // Returns true if the list contains no elements.
	Remove(value T) bool            // Removes all occurrences of the specified element.
	RemoveAt(index int) (T, bool)   // Removes the element at the specified position.
	Set(index int, value T) bool    // Replaces the element at the specified position.
	Size() int                      // Returns the number of elements in the list.
	SoftClear()                     // Soft reset, relies on GC to clean up
	ToSlice() []T                   // Returns a slice containing all elements in order.
}
