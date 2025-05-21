// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/search"
	"github.com/andrerrcosta2/gtools/gflux/core/pipes/arrays"
	"maps"
	"sync"
)

// Ordered creates a new ordered.
//
// This function returns a pointer to a new ordered, which is a set that maintains the order of its elements by its
// natural order.
//
// Type parameter K must satisfy the constraints.Ordered constraint, meaning it must be a type that supports ordering.
func Ordered[T prim.Ordered](values ...T) str.OrderedSet[T] {
	set := &ordered[T]{
		idx: make(map[T]struct{}),
		set: make([]T, 0, len(values)), // Preallocate with capacity
	}
	OrderedFilter(values, func(t T) {
		set.set = append(set.set, t)
		set.idx[t] = struct{}{}
	})
	return set
}

type ordered[T prim.Ordered] struct {
	idx map[T]struct{}
	set []T
}

// Add inserts an element into the set if it doesn't already exist.
func (s *ordered[T]) Add(t T) {
	if !s.Has(t) {
		// Binary search for insSet point
		pos, _ := search.Binary(s.set, t)
		s.set = append(s.set[:pos], append([]T{t}, s.set[pos:]...)...)
		s.idx[t] = struct{}{}
	}
}

func (s *ordered[T]) Clear() {
	s.set = make([]T, 0)
	s.idx = make(map[T]struct{})
}

func (s *ordered[T]) Delete(i int) bool {
	if i < 0 || i >= len(s.set) {
		return false
	}
	s.set = arrays.RemoveByIndex(s.set, i)
	delete(s.idx, s.set[i])
	return true
}

func (s *ordered[T]) Equals(other str.Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	switch set := other.(type) {
	case *ordered[T]:
		return arrays.Equals[T](s.set, set.set) && maps.Equal(s.idx, set.idx)
	default:
		setValues := set.Values()
		return arrays.Equals[T](s.set, setValues)
	}
}

// Get returns the element at the given idx.
func (s *ordered[T]) Get(i int) (T, bool) {
	if i >= 0 && len(s.set) < i {
		return s.set[i], true
	}
	var zeroValue T
	return zeroValue, false
}

// Has checks if the set contains the element.
func (s *ordered[T]) Has(t T) bool {
	_, exists := s.idx[t]
	return exists
}

func (s *ordered[T]) IndexOf(t T) int {
	i, _ := search.Binary(s.set, t)
	return i
}

// IsEmpty checks if the set is empty.
func (s *ordered[T]) IsEmpty() bool {
	return len(s.set) == 0
}

// Len returns the number of elements in the set.
func (s *ordered[T]) Len() int {
	return len(s.set)
}

// Remove deletes an element from the set.
func (s *ordered[T]) Remove(t T) {
	if s.Has(t) {
		// Binary search for the position of the item
		pos, _ := search.Binary(s.set, t)
		// Remove the item
		s.set = append(s.set[:pos], s.set[pos+1:]...)
		delete(s.idx, t)
	}
}

// Values returns the ordered elements in the set.
func (s *ordered[T]) Values() []T {
	return s.set
}

func (s *ordered[T]) Loop() <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for _, item := range s.set {
			ch <- item
		}
	}()

	return ch
}

var _ str.Set[int] = (*ordered[int])(nil)

type concOrder[T prim.Ordered] struct {
	mtx   sync.RWMutex
	items []T
	idx   map[T]struct{}
}

func (s *concOrder[T]) Add(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if !s.Has(t) {
		// Binary search for insSet point
		pos, _ := search.Binary(s.items, t)
		s.items = append(s.items[:pos], append([]T{t}, s.items[pos:]...)...)
		s.idx[t] = struct{}{}
	}
}

func (s *concOrder[T]) Has(t T) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	_, exists := s.idx[t]
	return exists
}
