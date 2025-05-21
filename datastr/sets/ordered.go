// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/datastr/trees"
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
		tree: trees.RedBlack[T](comparators.Ordered[T]{}, values...),
		set:  make(map[T]struct{}, len(values)),
	}
	for _, value := range values {
		set.set[value] = struct{}{}
	}
	return set
}

type ordered[T prim.Ordered] struct {
	set  map[T]struct{}
	tree str.RedBlackTree[T]
}

// Add inserts an element into the set if it doesn't already exist.
func (s *ordered[T]) Add(t T) {
	s.tree.Insert(t)
	s.set[t] = struct{}{}
}

func (s *ordered[T]) Clear() {
	s.tree.Clear()
	s.set = make(map[T]struct{})
}

func (s *ordered[T]) Delete(idx int) bool {
	if idx < 0 || idx >= s.tree.Size() {
		return false
	}
	it := s.tree.Iterator()
	var del T
	for i := 0; i <= idx; i++ {
		del = it.Next()
	}
	// I might have to implement this on the red black tree
	// this is performing 2 lookups
	s.tree.Delete(del)
	delete(s.set, del)
	return true
}

func (s *ordered[T]) Equals(other str.Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	switch set := other.(type) {
	case *ordered[T]:
		if len(s.set) != len(set.set) {
			return false
		}
		for k := range s.set {
			if _, ok := set.set[k]; !ok {
				return false
			}
		}
		return true
	default:
		if s.Len() != set.Len() {
			return false
		}
		for _, v := range set.Values() {
			if !s.Has(v) {
				return false
			}
		}
		return true
	}
}

// Get returns the element at the given idx.
func (s *ordered[T]) Get(idx int) (value T, ok bool) {
	if idx < 0 || idx >= s.tree.Size() {
		return
	}

	it := s.tree.Iterator()
	for i := 0; i <= idx; i++ {
		value = it.Next()
	}
	return value, true
}

// Has checks if the set contains the element.
func (s *ordered[T]) Has(t T) bool {
	return s.tree.Contains(t)
}

func (s *ordered[T]) IndexOf(t T) int {
	it := s.tree.Iterator()
	for i := 0; it.HasNext(); i++ {
		if it.Next() == t {
			return i
		}
	}
	return -1
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
		s.tree.Delete(t)
		delete(s.set, t)
	}
}

// Values returns the ordered elements in the set.
func (s *ordered[T]) Values() []T {
	var values = make([]T, 0, len(s.set))
	for k := range s.set {
		values = append(values, k)
	}
	return values
}

func (s *ordered[T]) Stream() gtools.Stream[T] {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for item := range s.set {
			ch <- item
		}
	}()

	return ch
}

var _ str.Set[int] = (*ordered[int])(nil)

func ConcOrdered[T prim.Ordered](values ...T) str.OrderedSet[T] {
	set := &concOrdered[T]{
		tree: trees.RedBlack[T](comparators.Ordered[T]{}, values...),
		set:  make(map[T]struct{}, len(values)),
	}
	for _, value := range values {
		set.set[value] = struct{}{}
	}
	return set
}

type concOrdered[T prim.Ordered] struct {
	mtx  sync.RWMutex
	set  map[T]struct{}
	tree str.RedBlackTree[T]
}

func (s *concOrdered[T]) Add(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.set[t] = struct{}{}
	s.tree.Insert(t)
}

func (s *concOrdered[T]) Clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.tree.Clear()
	s.set = make(map[T]struct{})
}

func (s *concOrdered[T]) Delete(idx int) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if idx < 0 || idx >= s.tree.Size() {
		return false
	}
	it := s.tree.Iterator()
	var del T
	for i := 0; i <= idx; i++ {
		del = it.Next()
	}
	// I might have to implement this on the red black tree
	// this is performing 2 lookups
	s.tree.Delete(del)
	delete(s.set, del)
	return true
}

func (s *concOrdered[T]) Equals(other str.Set[T]) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s.Len() != other.Len() {
		return false
	}
	switch set := other.(type) {
	case *ordered[T]:
		if len(s.set) != len(set.set) {
			return false
		}
		for k := range s.set {
			if _, ok := set.set[k]; !ok {
				return false
			}
		}
		return true
	default:
		if s.Len() != set.Len() {
			return false
		}
		for _, v := range set.Values() {
			if !s.Has(v) {
				return false
			}
		}
		return true
	}
}

// Get returns the element at the given idx.
func (s *concOrdered[T]) Get(idx int) (value T, ok bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if idx < 0 || idx >= s.tree.Size() {
		return
	}

	it := s.tree.Iterator()
	for i := 0; i <= idx; i++ {
		value = it.Next()
	}
	return value, true
}

// Has checks if the set contains the element.
func (s *concOrdered[T]) Has(t T) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.tree.Contains(t)
}

func (s *concOrdered[T]) IndexOf(t T) int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	it := s.tree.Iterator()
	for i := 0; it.HasNext(); i++ {
		if it.Next() == t {
			return i
		}
	}
	return -1
}

// IsEmpty checks if the set is empty.
func (s *concOrdered[T]) IsEmpty() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.set) == 0
}

// Len returns the number of elements in the set.
func (s *concOrdered[T]) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.set)
}

// Remove deletes an element from the set.
func (s *concOrdered[T]) Remove(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.set[t]; ok {
		s.tree.Delete(t)
		delete(s.set, t)
	}
}

// Values returns the ordered elements in the set.
func (s *concOrdered[T]) Values() []T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	var values = make([]T, 0, len(s.set))
	for k := range s.set {
		values = append(values, k)
	}
	return values
}

func (s *concOrdered[T]) Stream() gtools.Stream[T] {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	ch := make(chan T)
	go func() {
		defer close(ch)
		for item := range s.set {
			ch <- item
		}
	}()
	return ch
}
