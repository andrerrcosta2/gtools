// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/search"
	"github.com/andrerrcosta2/gtools/pkg/sortables"
	"github.com/andrerrcosta2/gtools/pkg/sorts"
	"maps"
)

// SortableOf returns a new instance of SortableOfSet.
// It creates a new set with an empty slice of items and an empty index map.
func SortableOf[T gtools.SortableOf](values ...T) *SortableOfSet[T] {
	// Sort the values in ascending order.
	sorts.QuickOf(&values)
	// Initialize a new SortableOfSet with an empty slice of items and an empty index map.
	set := &SortableOfSet[T]{
		// The items slice is initialized with the values passed as arguments.
		items: make([]T, 0),
		// The index map is initialized with an empty map of string to struct{}.
		index: make(map[string]struct{}),
		// The comparator is initialized with the ComparatorOf function.
		comparator: sortables.ComparatorOf[T](),
	}

	// Add each value to the set.
	for i, value := range values {
		set.items[i] = value
		set.index[set.comparator.Hash(value)] = struct{}{}
	}

	// Return the populated SortableOfSet instance.
	return set
}

type SortableOfSet[T gtools.SortableOf] struct {
	items      []T
	index      map[string]struct{}
	comparator comparables.KeyComparator[T, string]
}

func (s *SortableOfSet[T]) Has(t T) bool {
	unique := s.comparator.Hash(t)
	_, exists := s.index[unique]
	return exists
}

func (s *SortableOfSet[T]) Add(t T) {
	unique := s.comparator.Hash(t)
	if _, exists := s.index[unique]; !exists {
		// Binary search for insertion point
		pos := search.BinaryOf(s.items, t)
		// Insert item at the found position
		s.items = append(s.items[:pos], append([]T{t}, s.items[pos:]...)...)
		s.index[unique] = struct{}{}
	}
}

func (s *SortableOfSet[T]) Remove(t T) {
	if s.Has(t) {
		// Binary search for the position of the item
		pos := search.BinaryOf(s.items, t)
		// Remove the item
		s.items = append(s.items[:pos], s.items[pos+1:]...)
		unique := s.comparator.Hash(t)
		delete(s.index, unique)
	}
}

func (s *SortableOfSet[T]) Len() int {
	return len(s.items)
}

func (s *SortableOfSet[T]) Values() []T {
	return s.items
}

func (s *SortableOfSet[T]) Get(i int) (T, bool) {
	if !arrays.OutOfBounds(s.items, i) {
		return s.items[i], true
	}
	var zeroValue T
	return zeroValue, false
}

func (s *SortableOfSet[T]) Exclude(i int) bool {
	if !arrays.OutOfBounds(s.items, i) {
		s.items = append(s.items[:i], s.items[i+1:]...)
		return true
	}
	return false
}

// Loop returns a channel of the items in the set.
//
// Example of use:
//
//	for item := range s.Loop() {
//	  fmt.Println(item)
//	}
func (s *SortableOfSet[T]) Loop() <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for _, item := range s.items {
			ch <- item
		}
	}()

	return ch
}

func (s *SortableOfSet[T]) Clear() {
	s.items = make([]T, 0)
	s.index = make(map[string]struct{})
}

func (s *SortableOfSet[T]) Equals(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	values := other.Values()
	if len(values) == 0 {
		return false
	}

	if set, ok := other.(*SortableOfSet[T]); ok {
		return arrays.SortedEqualsBy[T](s.items, set.items, s.comparator.Equals) &&
			maps.Equal(s.index, set.index)
	}

	if _, ok := any(values[0]).(gtools.SortableOf); ok {
		sortable := SortableOf(values...)
		return arrays.SortedEqualsBy[T](s.items, sortable.items, s.comparator.Equals) &&
			maps.Equal(s.index, sortable.index)
	}

	return false
}

func (s *SortableOfSet[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}
