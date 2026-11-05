// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"encoding/json"
	"slices"
	"sync"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/list"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/datastr/lists"
)

// Insertion is a backingMap that preserves insertion order of elements
// This map uses a linked list to keep the elements orders and avoid overhead during removal.
// for that reason its performance for traversals is less efficient than slice based sets
func Insertion[T any, H prim.Hashable](cmp comparator.EqCmpKeyer[T, H], values ...T) set.MutableOrdered[T] {
	return &insertion[T, H]{
		list: lists.HashLinked[T, H](cmp, Filter[T, H](cmp, values...)...),
		cmp:  cmp,
	}
}

type insertion[T any, H prim.Hashable] struct {
	list list.Mutable[T]
	cmp  comparator.EqCmpKeyer[T, H]
}

func (s *insertion[T, H]) Add(t T) {
	if !s.list.Contains(t) {
		s.list.Add(t)
	}
}

func (s *insertion[T, H]) Clear() {
	s.list.SoftClear()
}

func (s *insertion[T, H]) contains(t T) bool {
	return s.list.Contains(t)
}

func (s *insertion[T, H]) Delete(i int) bool {
	_, ok := s.list.RemoveAt(i)
	return ok
}

func (s *insertion[T, H]) Equals(other set.Readable[T]) bool {
	if other == nil {
		return false
	}
	if s == other {
		return true
	}
	if s.Len() != other.Len() {
		return false
	}
	return slices.EqualFunc(s.Elements(), other.Elements(), s.cmp.Equals)
}

func (s *insertion[T, H]) Get(i int) (T, bool) {
	return s.list.Get(i)
}

func (s *insertion[T, H]) Has(t T) bool {
	return s.list.Contains(t)
}

func (s *insertion[T, H]) IndexOf(t T) int {
	return s.list.IndexOf(t)
}

func (s *insertion[T, H]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *insertion[T, H]) Len() int {
	return s.list.Size()
}

func (s *insertion[T, H]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.list.ToSlice())
}

func (s *insertion[T, H]) Range(callback func(t T) bool) {
	for _, value := range s.Elements() {
		if !callback(value) {
			return
		}
	}
}

func (s *insertion[T, H]) Remove(t T) {
	s.list.Remove(t)
}

func (s *insertion[T, H]) String() string {
	return s.sprint(indent.Zero())
}

func (s *insertion[T, H]) Sprint(tab indent.Indentor) string {
	return s.sprint(tab)
}

func (s *insertion[T, H]) sprint(tab indent.Indentor) string {
	return sprints.TypedSlice(tab, s.list.ToSlice(), "sets.Insertion")
}

func (s *insertion[T, H]) UnmarshalJSON(data []byte) error {
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

func (s *insertion[T, H]) Elements() []T {
	return s.list.ToSlice()
}

var _ set.Mutable[int] = (*insertion[int, int])(nil)
var _ set.MutableOrdered[int] = (*insertion[int, int])(nil)

// ConcInsertion is a thread safe backingMap that preserves the insertion order of elements
func ConcInsertion[T any, H prim.Hashable](cmp comparator.EqCmpKeyer[T, H], values ...T) set.MutableOrdered[T] {
	return &insertion[T, H]{
		list: lists.ConcHashLinked[T, H](cmp, Filter[T, H](cmp, values...)...),
		cmp:  cmp,
	}
}

type concIns[T any, H prim.Hashable] struct {
	mtx  sync.RWMutex
	list list.Mutable[T]
	cmp  comparator.EqCmpKeyer[T, H]
}

func (s *concIns[T, H]) Add(t T) {
	s.list.Add(t)
}

func (s *concIns[T, H]) Clear() {
	s.list.SoftClear()
}

func (s *concIns[T, H]) Delete(i int) bool {
	_, ok := s.list.RemoveAt(i)
	return ok
}

func (s *concIns[T, H]) Equals(other set.Readable[T]) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if other == s {
		return true
	}
	if s.Len() != other.Len() {
		return false
	}
	return slices.EqualFunc(s.Elements(), other.Elements(), s.cmp.Equals)
}

func (s *concIns[T, H]) Get(i int) (T, bool) {
	return s.list.Get(i)
}

func (s *concIns[T, H]) Has(t T) bool {
	return s.list.Contains(t)
}

func (s *concIns[T, H]) IndexOf(t T) int {
	return s.list.IndexOf(t)
}

func (s *concIns[T, H]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *concIns[T, H]) Len() int {
	return s.list.Size()
}

func (s *concIns[T, H]) MarshalJSON() ([]byte, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return json.Marshal(s.list.ToSlice())
}

func (s *concIns[T, H]) Range(callback func(t T) bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	for _, value := range s.Elements() {
		if !callback(value) {
			return
		}
	}
}

func (s *concIns[T, H]) Remove(t T) {
	s.list.Remove(t)
}

func (s *concIns[T, H]) String() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.sprint(indent.Zero())
}

func (s *concIns[T, H]) Sprint(tab indent.Indentor) string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.sprint(tab)
}

func (s *concIns[T, H]) sprint(tab indent.Indentor) string {
	return sprints.TypedSlice(tab, s.list.ToSlice(), "sets.ConcInsertion")
}

func (s *concIns[T, H]) UnmarshalJSON(data []byte) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
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

func (s *concIns[T, H]) Elements() []T {
	return s.list.ToSlice()
}

var _ set.Mutable[int] = (*concIns[int, int])(nil)
var _ set.MutableOrdered[int] = (*concIns[int, int])(nil)
