// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/comparables"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/search"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/core/sortables/sorts"
	"github.com/andrerrcosta2/gtools/pipes/arrays"
	"maps"
)

// SortableOf returns a new instance of sortableOfSet.
// It creates a new set with an empty slice of items and an empty index map.
func SortableOf[T gtools.SortableOf](values ...T) str.Set[T] {
	// Sorter the values in ascending order.
	sorts.QuickOf(values)
	// Initialize a new sortableOfSet with an empty slice of items and an empty index map.
	set := &sortableOfSet[T]{
		// The items slice is initialized with the values passed as arguments.
		items: make([]T, 0),
		// The index map is initialized with an empty map of string to struct{}.
		index: make(map[string]struct{}),
		// The comparator is initialized with the ComparatorOf function.
		comparator: sortables.ComparatorOf[T](),
	}

	// Addf each value to the set.
	for i, value := range values {
		set.items[i] = value
		hash := set.comparator.Hash(value)
		set.index[hash] = struct{}{}
	}

	// Return the populated sortableOfSet instance.
	return set
}

type sortableOfSet[T gtools.SortableOf] struct {
	items      []T
	index      map[string]struct{}
	comparator comparables.KeyComparator[T, string]
}

func (s *sortableOfSet[T]) Has(t T) bool {
	unique := s.comparator.Hash(t)
	_, exists := s.index[unique]
	return exists
}

func (s *sortableOfSet[T]) Add(t T) {
	unique := s.comparator.Hash(t)
	if _, exists := s.index[unique]; !exists {
		// binary search for insertion point
		pos := search.BinaryOf(s.items, t)
		// Insert item at the found position
		s.items = append(s.items[:pos], append([]T{t}, s.items[pos:]...)...)
		s.index[unique] = struct{}{}
	}
}

func (s *sortableOfSet[T]) Remove(t T) {
	if s.Has(t) {
		// binary search for the position of the item
		pos := search.BinaryOf(s.items, t)
		// Remove the item
		s.items = append(s.items[:pos], s.items[pos+1:]...)
		unique := s.comparator.Hash(t)
		delete(s.index, unique)
	}
}

func (s *sortableOfSet[T]) Len() int {
	return len(s.items)
}

func (s *sortableOfSet[T]) Values() []T {
	return s.items
}

func (s *sortableOfSet[T]) Get(i int) (T, bool) {
	if !arrays.OutOfBounds(s.items, i) {
		return s.items[i], true
	}
	var zeroValue T
	return zeroValue, false
}

func (s *sortableOfSet[T]) Exclude(i int) bool {
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
func (s *sortableOfSet[T]) Loop() <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for _, item := range s.items {
			ch <- item
		}
	}()

	return ch
}

func (s *sortableOfSet[T]) Clear() {
	s.items = make([]T, 0)
	s.index = make(map[string]struct{})
}

func (s *sortableOfSet[T]) Equals(other str.Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	values := other.Values()
	if len(values) == 0 {
		return false
	}

	if set, ok := other.(*sortableOfSet[T]); ok {
		return arrays.SortedEqualsBy[T](s.items, set.items, s.comparator.Equals) &&
			maps.Equal(s.index, set.index)
	}

	if _, ok := any(values[0]).(gtools.SortableOf); ok {
		sortable := SortableOf(values...).(*sortableOfSet[T])
		return arrays.SortedEqualsBy[T](s.items, sortable.items, s.comparator.Equals) &&
			maps.Equal(s.index, sortable.index)
	}

	return false
}

func (s *sortableOfSet[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

var _ str.Set[gtools.SortableOf] = (*sortableOfSet[gtools.SortableOf])(nil)
