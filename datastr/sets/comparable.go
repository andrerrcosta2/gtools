// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"encoding/json"
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"maps"
	"sync"
)

// Comparable Creates a new str.Set[T] of comparable set
//
// This set doesn't hold any order of elements
func Comparable[T comparable](values ...T) str.Set[T] {
	// Create a new cmpSet instance with an empty map.
	set := &cmpSet[T]{
		set: make(map[T]struct{}),
	}

	// Add each value to the set.
	for _, v := range values {
		set.Add(v)
	}

	// Return the populated cmpSet instance.
	return set
}

type cmpSet[T comparable] struct {
	set map[T]struct{}
}

func (s *cmpSet[T]) Add(t T) {
	s.set[t] = struct{}{}
}

func (s *cmpSet[T]) Clear() {
	s.set = make(map[T]struct{})
}

func (s *cmpSet[T]) Equals(other str.Set[T]) bool {
	if s == other {
		return true
	}
	switch set := other.(type) {
	case *cmpSet[T]:
		return maps.Equal(s.set, set.set)
	default:
		if s.Len() != set.Len() {
			return false
		}
		o := make(map[T]struct{})
		for _, v := range set.Values() {
			o[v] = struct{}{}
		}
		return maps.Equal(s.set, o)
	}
}

func (s *cmpSet[T]) Has(t T) bool {
	_, exists := s.set[t]
	return exists
}

func (s *cmpSet[T]) IsEmpty() bool {
	return len(s.set) == 0
}

func (s *cmpSet[T]) Len() int {
	return len(s.set)
}

func (s *cmpSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

func (s *cmpSet[T]) Remove(t T) {
	delete(s.set, t)
}

func (s *cmpSet[T]) UnmarshalJSON(data []byte) error {
	var slice []T
	err := json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}
	s.Clear()
	for _, v := range slice {
		s.Add(v)
	}
	return nil
}

func (s *cmpSet[T]) Values() []T {
	values := make([]T, 0, len(s.set))
	for key := range s.set {
		values = append(values, key)
	}
	return values
}

var _ str.Set[any] = (*cmpSet[any])(nil)
var _ data.JSONSerializable = (*cmpSet[any])(nil)

// ConcComparable Creates a new thread-safe str.Set[T] of comparable set
//
// This set doesn't hold any order of insertion
func ConcComparable[T comparable](values ...T) str.Set[T] {
	set := &concCmp[T]{
		set: make(map[T]struct{}),
	}
	for _, v := range values {
		set.Add(v)
	}
	return set
}

type concCmp[T comparable] struct {
	mtx sync.RWMutex
	set map[T]struct{}
}

func (s *concCmp[T]) Add(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.set[t] = struct{}{}
}

func (s *concCmp[T]) Clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.set = make(map[T]struct{})
}

func (s *concCmp[T]) Equals(other str.Set[T]) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	switch set := other.(type) {
	case *concCmp[T]:
		return maps.Equal(s.set, set.set)
	default:
		if s.Len() != set.Len() {
			return false
		}
		o := make(map[T]struct{})
		for _, v := range set.Values() {
			o[v] = struct{}{}
		}
		return maps.Equal(s.set, o)
	}
}

func (s *concCmp[T]) Has(t T) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	_, exists := s.set[t]
	return exists
}

func (s *concCmp[T]) IsEmpty() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.set) == 0
}

func (s *concCmp[T]) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.set)
}

func (s *concCmp[T]) MarshalJSON() ([]byte, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return json.Marshal(s.Values())
}

func (s *concCmp[T]) Remove(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	delete(s.set, t)
}

func (s *concCmp[T]) UnmarshalJSON(data []byte) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	var slice []T
	err := json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}
	s.Clear()
	for _, v := range slice {
		s.set[v] = struct{}{}
	}
	return nil
}

func (s *concCmp[T]) Values() []T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	values := make([]T, 0, len(s.set))
	for key := range s.set {
		values = append(values, key)
	}
	return values
}

var _ str.Set[any] = (*concCmp[any])(nil)
var _ data.JSONSerializable = (*concCmp[any])(nil)
