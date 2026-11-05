// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"encoding/json"
	"sync"

	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/keyer/keyers"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/datastr/internal/hash"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
)

// ComparableOf Creates a new str.Mutable[T] of comparable backingMap
//
// This backingMap doesn't hold any order of elements
func ComparableOf[T constraints.Comparable[T]](values ...T) set.Mutable[T] {
	// Initialize a new sortableOfSet with an empty slice of backingMap and an empty idx map.
	keyer := keyers.FingerPrint[T]()
	s := &cmpOfSet[T]{
		// The idx map is initialized with an empty map of string to struct{}.
		table: make(hash.FlatStore[string, T]),
		cmp: comparator.Build[T, string]().
			Key(keyer.Key).
			Equals(func(a, b T) bool { return a.Equal(b) }).
			Build(),
	}

	for _, value := range values {
		s.table.TryStore(keyer.Key(value), value)
	}

	// Return the populated sortableOfSet instance.
	return s
}

type cmpOfSet[T constraints.Comparable[T]] struct {
	table hash.FlatStore[string, T]
	cmp   comparator.EqKeyer[string, T]
}

func (s *cmpOfSet[T]) Add(t T) {
	s.table.TryStore(s.hash(t), t)
}

func (s *cmpOfSet[T]) Clear() {
	s.table.Clear()
}

func (s *cmpOfSet[T]) Equals(o set.Readable[T]) bool {
	if s == o {
		return true
	}
	return eqHashable(s.table, o, s.cmp)
}

func (s *cmpOfSet[T]) Has(t T) bool {
	return s.table.HasHash(s.hash(t))
}

func (s *cmpOfSet[T]) hash(t T) string {
	return s.cmp.Key(t)
}

func (s *cmpOfSet[T]) IsEmpty() bool {
	return s.table.IsEmpty()
}

func (s *cmpOfSet[T]) Len() int {
	return s.table.Len()
}

func (s *cmpOfSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Elements())
}

func (s *cmpOfSet[T]) Range(callback func(T) bool) {
	for _, value := range s.table {
		if !callback(value) {
			return
		}
	}
}

func (s *cmpOfSet[T]) Remove(t T) {
	s.table.TryRemove(s.hash(t))
}

func (s *cmpOfSet[T]) UnmarshalJSON(data []byte) error {
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

func (s *cmpOfSet[T]) Elements() []T {
	return s.table.Values()
}

type S = tests.SortableStub

var _ set.Mutable[S] = (*cmpOfSet[S])(nil)
var _ data.JSONSerializable = (*cmpOfSet[S])(nil)

// ConcComparableOf Creates a new thread-safe str.Mutable[T] of comparable backingMap
//
// This backingMap doesn't hold any order of insertion
func ConcComparableOf[T constraints.Comparable[T]](values ...T) set.Mutable[T] {
	rasher := keyers.FingerPrint[T]()
	s := &concCmpOf[T]{
		// The idx map is initialized with an empty map of string to struct{}.
		table: make(hash.FlatStore[string, T]),
		cmp: comparator.Build[T, string]().
			Key(rasher.Key).
			Equals(func(a, b T) bool { return a.Equal(b) }).
			Build(),
	}

	for _, value := range values {
		s.table.TryStore(rasher.Key(value), value)
	}

	// Return the populated sortableOfSet instance.
	return s
}

type concCmpOf[T constraints.Comparable[T]] struct {
	mtx   sync.RWMutex
	table hash.FlatStore[string, T]
	cmp   comparator.EqKeyer[string, T]
}

func (s *concCmpOf[T]) Add(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.table.TryStore(s.hash(t), t)
}

func (s *concCmpOf[T]) Clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.table.Clear()
}

func (s *concCmpOf[T]) Equals(o set.Readable[T]) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s == o {
		return true
	}
	return eqHashable(s.table, o, s.cmp)
}

func (s *concCmpOf[T]) Has(t T) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.table.HasHash(s.hash(t))
}

func (s *concCmpOf[T]) hash(t T) string {
	return s.cmp.Key(t)
}

func (s *concCmpOf[T]) IsEmpty() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.table.IsEmpty()
}

func (s *concCmpOf[T]) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.table.Len()
}

func (s *concCmpOf[T]) MarshalJSON() ([]byte, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return json.Marshal(s.Elements())
}

func (s *concCmpOf[T]) Range(callback func(T) bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, value := range s.table {
		if !callback(value) {
			return
		}
	}
}

func (s *concCmpOf[T]) Remove(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.table.TryRemove(s.hash(t))
}

func (s *concCmpOf[T]) UnmarshalJSON(data []byte) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	var slice []T
	err := json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}
	s.Clear()
	for _, v := range slice {
		s.table.TryStore(s.hash(v), v)
	}
	return nil
}

func (s *concCmpOf[T]) Elements() []T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	values := make([]T, 0, s.table.Len())
	for _, value := range s.table {
		values = append(values, value)
	}
	return values
}

var _ set.Mutable[S] = (*concCmpOf[S])(nil)
var _ data.JSONSerializable = (*concCmpOf[S])(nil)
