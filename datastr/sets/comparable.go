// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"encoding/json"
	"sync"

	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/datastr/internal/hash"
)

// Comparable Creates a new str.Mutable[T] of comparable backingMap
//
// This backingMap doesn't hold any order of elements
func Comparable[T comparable](values ...T) set.Mutable[T] {
	// Create a new cmpSet instance with an empty map.
	set := &cmpSet[T]{
		table: make(map[T]struct{}),
	}

	// Add each value to the backingMap.
	for _, v := range values {
		set.Add(v)
	}

	// Return the populated cmpSet instance.
	return set
}

type cmpSet[T comparable] struct {
	table hash.Table[T]
}

func (s *cmpSet[T]) Add(t T) {
	s.table.TryAddHash(t)
}

func (s *cmpSet[T]) Clear() {
	s.table.Clear()
}

func (s *cmpSet[T]) Equals(other set.Readable[T]) bool {
	if s == other {
		return true
	}
	return eqComparable(s.table, other)
}

func (s *cmpSet[T]) Has(t T) bool {
	_, exists := s.table[t]
	return exists
}

func (s *cmpSet[T]) IsEmpty() bool {
	return len(s.table) == 0
}

func (s *cmpSet[T]) Len() int {
	return len(s.table)
}

func (s *cmpSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Elements())
}

func (s *cmpSet[T]) Range(f func(T) bool) {
	for k := range s.table {
		if !f(k) {
			return
		}
	}
}

func (s *cmpSet[T]) Remove(t T) {
	delete(s.table, t)
}

func (s *cmpSet[T]) UnmarshalJSON(data []byte) error {
	var slice []T
	err := json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}
	s.Clear()
	for _, v := range slice {
		s.table.TryAddHash(v)
	}
	return nil
}

func (s *cmpSet[T]) Elements() []T {
	values := make([]T, 0, len(s.table))
	for key := range s.table {
		values = append(values, key)
	}
	return values
}

func (s *cmpSet[T]) backingMap() map[T]struct{} {
	return s.table
}

var _ set.Mutable[any] = (*cmpSet[any])(nil)
var _ data.JSONSerializable = (*cmpSet[any])(nil)

// ConcComparable Creates a new thread-safe str.Mutable[T] of comparable backingMap
//
// This backingMap doesn't hold any order of insertion
func ConcComparable[T comparable](values ...T) set.Mutable[T] {
	s := &concCmp[T]{
		table: make(map[T]struct{}),
	}
	for _, v := range values {
		s.table.TryAddHash(v)
	}
	return s
}

type concCmp[T comparable] struct {
	mtx   sync.RWMutex
	table hash.Table[T]
}

func (s *concCmp[T]) Add(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.table.TryAddHash(t)
}

func (s *concCmp[T]) Clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.table.Clear()
}

func (s *concCmp[T]) Equals(other set.Readable[T]) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return eqComparable(s.table, other)
}

func (s *concCmp[T]) Has(t T) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.table.HasHash(t)
}

func (s *concCmp[T]) IsEmpty() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.table.IsEmpty()
}

func (s *concCmp[T]) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.table.Len()
}

func (s *concCmp[T]) MarshalJSON() ([]byte, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return json.Marshal(s.Elements())
}

func (s *concCmp[T]) Range(f func(T) bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	for k := range s.table {
		if !f(k) {
			return
		}
	}
}

func (s *concCmp[T]) Remove(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.table.TryRemoveHash(t)
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
		s.table.TryAddHash(v)
	}
	return nil
}

func (s *concCmp[T]) Elements() []T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	values := make([]T, 0, s.table.Len())
	for key := range s.table {
		values = append(values, key)
	}
	return values
}

var _ set.Mutable[any] = (*concCmp[any])(nil)
var _ data.JSONSerializable = (*concCmp[any])(nil)

func ImmutableComparable[T comparable](values ...T) set.Immutable[T] {
	m := make(map[T]struct{}, len(values))

	for _, v := range values {
		m[v] = struct{}{}
	}

	return &immCmpSet[T]{
		table: m,
	}
}

type immCmpSet[T comparable] struct {
	table hash.Table[T]
}

func (s *immCmpSet[T]) Copy() set.Mutable[T] {
	return Comparable(s.Elements()...)
}

func (s *immCmpSet[T]) Equals(other set.Readable[T]) bool {
	return eqComparable(s.table, other)
}

func (s *immCmpSet[T]) Has(value T) bool {
	_, exists := s.table[value]
	return exists
}

func (s *immCmpSet[T]) IsEmpty() bool {
	return len(s.table) == 0
}

func (s *immCmpSet[T]) Len() int {
	return len(s.table)
}

func (s *immCmpSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Elements())
}

func (s *immCmpSet[T]) Range(f func(T) bool) {
	for k := range s.table {
		if !f(k) {
			return
		}
	}
}

func (s *immCmpSet[T]) Elements() []T {
	out := make([]T, 0, len(s.table))
	for v := range s.table {
		out = append(out, v)
	}
	return out
}

func (s *immCmpSet[T]) backingMap() map[T]struct{} {
	return s.table
}

var _ set.Readable[any] = (*immCmpSet[any])(nil)
var _ set.Immutable[any] = (*immCmpSet[any])(nil)
