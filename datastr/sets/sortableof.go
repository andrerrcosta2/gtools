// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/search"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/core/sortables/sorts"
	"github.com/andrerrcosta2/gtools/gflux/core/pipes/arrays"
	"maps"
	"sync"
)

// SortableOf returns a new instance of sortableOfSet.
// It creates a new set with an empty slice of set and an empty idx map.
func SortableOf[T gtools.SortableOf](values ...T) str.Set[T] {
	// Sorter the values in ascending order.
	sorts.QuickOf(values)
	// Initialize a new sortableOfSet with an empty slice of set and an empty idx map.
	set := &sortableOfSet[T]{
		// The set slice is initialized with the values passed as arguments.
		items: make([]T, 0),
		// The idx map is initialized with an empty map of string to struct{}.
		index: make(map[string]struct{}),
		// The cmp is initialized with the ComparatorOf function.
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
	comparator comparators.KeyTyped[T, string]
}

func (s *sortableOfSet[T]) Add(t T) {
	unique := s.comparator.Hash(t)
	if _, exists := s.index[unique]; !exists {
		// binary search for insSet point
		pos, _ := search.BinaryOf(s.items, t)
		// Insert item at the found position
		s.items = append(s.items[:pos], append([]T{t}, s.items[pos:]...)...)
		s.index[unique] = struct{}{}
	}
}

func (s *sortableOfSet[T]) Clear() {
	s.items = make([]T, 0)
	s.index = make(map[string]struct{})
}

func (s *sortableOfSet[T]) Delete(i int) bool {
	if !arrays.OutOfBounds(s.items, i) {
		s.items = append(s.items[:i], s.items[i+1:]...)
		return true
	}
	return false
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

func (s *sortableOfSet[T]) Get(i int) (T, bool) {
	if !arrays.OutOfBounds(s.items, i) {
		return s.items[i], true
	}
	var zeroValue T
	return zeroValue, false
}

func (s *sortableOfSet[T]) Has(t T) bool {
	unique := s.comparator.Hash(t)
	_, exists := s.index[unique]
	return exists
}

func (s *sortableOfSet[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *sortableOfSet[T]) Len() int {
	return len(s.items)
}

// Loop returns a channel of the set in the set.
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

func (s *sortableOfSet[T]) Remove(t T) {
	if s.Has(t) {
		// binary search for the position of the item
		pos, _ := search.BinaryOf(s.items, t)
		// Remove the item
		s.items = append(s.items[:pos], s.items[pos+1:]...)
		unique := s.comparator.Hash(t)
		delete(s.index, unique)
	}
}

func (s *sortableOfSet[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

func (s *sortableOfSet[T]) Values() []T {
	return s.items
}

var _ str.Set[gtools.SortableOf] = (*sortableOfSet[gtools.SortableOf])(nil)

func ConcSortableOf[T gtools.SortableOf](values ...T) str.Set[T] {
	return &concSortableOf[T]{
		items:      make([]T, len(values)),
		index:      make(map[string]struct{}),
		comparator: sortables.ComparatorOf[T](),
	}
}

type concSortableOf[T gtools.SortableOf] struct {
	mtx        sync.RWMutex
	items      []T
	index      map[string]struct{}
	comparator comparators.KeyTyped[T, string]
}

func (s *concSortableOf[T]) Add(t T) {
	unique := s.comparator.Hash(t)
	s.mtx.Lock()
	if _, exists := s.index[unique]; !exists {
		// binary search for insSet point
		pos, _ := search.BinaryOf(s.items, t)
		// Insert item at the found position
		s.items = append(s.items[:pos], append([]T{t}, s.items[pos:]...)...)
		s.index[unique] = struct{}{}
	}
	s.mtx.Unlock()
}

func (s *concSortableOf[T]) Clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.items = make([]T, 0)
	s.index = make(map[string]struct{})
}

func (s *concSortableOf[T]) Delete(i int) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if !arrays.OutOfBounds(s.items, i) {
		s.items = append(s.items[:i], s.items[i+1:]...)
		return true
	}
	return false
}

func (s *concSortableOf[T]) Get(i int) (T, bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if !arrays.OutOfBounds(s.items, i) {
		return s.items[i], true
	}
	var zeroValue T
	return zeroValue, false
}

func (s *concSortableOf[T]) Has(t T) bool {
	unique := s.comparator.Hash(t)
	s.mtx.RLock()
	_, exists := s.index[unique]
	s.mtx.RUnlock()
	return exists
}

func (s *concSortableOf[T]) IsEmpty() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.items) == 0
}

// Loop returns a channel of the set in the set.
//
// Example of use:
//
//	for item := range s.Loop() {
//	  fmt.Println(item)
//	}
func (s *concSortableOf[T]) Loop() <-chan T {
	items, _ := s.snap()
	ch := make(chan T)

	go func() {
		defer close(ch)
		for _, item := range items {
			ch <- item
		}
	}()

	return ch
}

func (s *concSortableOf[T]) Remove(t T) {
	unique := s.comparator.Hash(t)
	s.mtx.Lock()
	if _, exists := s.index[unique]; exists {
		pos, _ := search.BinaryOf(s.items, t)
		s.items = append(s.items[:pos], s.items[pos+1:]...)
		delete(s.index, unique)
	}
	s.mtx.Unlock()
}

func (s *concSortableOf[T]) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.items)
}

func (s *concSortableOf[T]) snap() ([]T, map[string]struct{}) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	// Copy the set slice
	items := make([]T, len(s.items))
	copy(items, s.items)

	// Copy the idx map
	index := make(map[string]struct{}, len(s.index))
	for k, v := range s.index {
		index[k] = v
	}

	return items, index
}

func (s *concSortableOf[T]) Values() []T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.items
}

func (s *concSortableOf[T]) Equals(other str.Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}

	sItems, sIndex := s.snap()
	oItems, oIndex := Snap(other)

	// Compare the set slices
	if !arrays.SortedEqualsBy[T](sItems, oItems, s.comparator.Equals) {
		return false
	}

	// Compare the idx maps
	return maps.Equal(sIndex, oIndex)
}
