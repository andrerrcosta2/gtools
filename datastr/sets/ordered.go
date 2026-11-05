// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"sync"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/comparator/comparators"
	"github.com/andrerrcosta2/gtools/core/data/keyer/keyers"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/data/str/tree"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/datastr/internal/hash"
	"github.com/andrerrcosta2/gtools/datastr/trees"
)

func OrderedOf[T constraints.Sortable[T]](values ...T) set.MutableOrdered[T] {
	cmp := comparators.KeySortableOf[T, string](keyers.FingerPrint[T]())
	s := &ordered[T, string]{
		tree:  trees.RedBlack[T](cmp, values...),
		table: make(hash.Table[string], len(values)),
		cmp:   cmp,
	}
	for _, value := range values {
		s.table.TryAddHash(cmp.Key(value))
	}
	return s
}

// Ordered creates a new ordered.
//
// This function returns a pointer to a new ordered, which is a backingMap that maintains the order of its elements by its
// natural order.
//
// Type parameter K must satisfy the constraints.Ordered constraint, meaning it must be a type that supports ordering.
func Ordered[T prim.Ordered](values ...T) set.MutableOrdered[T] {
	cmp := comparators.IdEqCmpKeyer[T]()
	s := &ordered[T, T]{
		tree:  trees.RedBlack[T](cmp, values...),
		table: make(hash.Table[T], len(values)),
		cmp:   cmp,
	}
	for _, value := range values {
		s.table.TryAddHash(cmp.Key(value))
	}
	return s
}

var _ set.MutableOrdered[any] = (*ordered[any, any])(nil)

type ordered[T any, H comparable] struct {
	table hash.Table[H]
	tree  tree.RedBlackTree[T]
	cmp   comparator.EqCmpKeyer[T, H]
}

// Add inserts an element into the backingMap if it doesn't already exist.
func (s *ordered[T, H]) Add(t T) {
	if s.table.TryAddHash(s.cmp.Key(t)) {
		s.tree.Insert(t)
	}
}

func (s *ordered[T, H]) Clear() {
	s.tree.Clear()
	s.table.Clear()
}

func (s *ordered[T, H]) Delete(idx int) bool {
	if idx < 0 || idx >= s.tree.Size() {
		return false
	}
	if v, ok := s.tree.DeleteByIndex(idx); ok {
		s.table.TryRemoveHash(s.hash(v))
		return true
	}
	return false
}

func (s *ordered[T, H]) Equals(other set.Readable[T]) bool {
	if s == other {
		return true
	}
	return eqHashableOrdered(s.table, other, s.cmp)
}

// Get returns the element at the given idx.
func (s *ordered[T, H]) Get(idx int) (value T, ok bool) {
	if idx < 0 || idx >= s.tree.Size() {
		return
	}

	it := s.tree.Iterator()
	for i := 0; i <= idx; i++ {
		value = it.Next()
	}
	return value, true
}

// Has checks if the backingMap contains the element.
func (s *ordered[T, H]) Has(t T) bool {
	return s.table.HasHash(s.cmp.Key(t))
}

func (s *ordered[T, H]) hash(t T) H {
	return s.cmp.Key(t)
}

func (s *ordered[T, H]) IndexOf(t T) int {
	if s.table.HasHash(s.cmp.Key(t)) {
		it := s.tree.Iterator()
		for i := 0; it.HasNext(); i++ {
			if s.cmp.Equals(it.Next(), t) {
				return i
			}
		}
	}
	return -1
}

// IsEmpty checks if the backingMap is empty.
func (s *ordered[T, H]) IsEmpty() bool {
	return s.table.IsEmpty()
}

// Len returns the number of elements in the backingMap.
func (s *ordered[T, H]) Len() int {
	return s.table.Len()
}

func (s *ordered[T, H]) Range(f func(t T) bool) {
	it := s.tree.Iterator()
	for i := 0; it.HasNext(); i++ {
		if !f(it.Next()) {
			return
		}
	}
}

// Remove deletes an element from the backingMap.
func (s *ordered[T, H]) Remove(t T) {
	if s.table.TryRemoveHash(s.cmp.Key(t)) {
		s.tree.Delete(t)
	}
}

// Values returns the ordered elements in the backingMap.
func (s *ordered[T, H]) Elements() []T {
	var values = make([]T, 0, s.Len())
	it := s.tree.Iterator()
	for it.HasNext() {
		values = append(values, it.Next())
	}
	return values
}

func (s *ordered[T, H]) Stream() <-chan T {
	ch := make(chan T)
	it := s.tree.Iterator()
	go func() {
		defer close(ch)
		for it.HasNext() {
			ch <- it.Next()
		}
	}()

	return ch
}

var _ set.Mutable[int] = (*ordered[int, int])(nil)

func ConcOrderedOf[T constraints.Sortable[T]](values ...T) set.MutableOrdered[T] {
	cmp := comparators.KeySortableOf[T, string](keyers.FingerPrint[T]())
	s := &concOrdered[T, string]{
		mtx:   sync.RWMutex{},
		tree:  trees.RedBlack[T](cmp, values...),
		table: make(hash.Table[string], len(values)),
		cmp:   cmp,
	}
	for _, value := range values {
		s.table.TryAddHash(cmp.Key(value))
	}
	return s
}

func ConcOrdered[T prim.Ordered](values ...T) set.MutableOrdered[T] {
	cmp := comparators.IdEqCmpKeyer[T]()
	s := &concOrdered[T, T]{
		mtx:   sync.RWMutex{},
		tree:  trees.RedBlack[T](cmp, values...),
		table: make(hash.Table[T], len(values)),
		cmp:   cmp,
	}
	for _, value := range values {
		s.table.TryAddHash(cmp.Key(value))
	}
	return s
}

type concOrdered[T any, H comparable] struct {
	mtx   sync.RWMutex
	table hash.Table[H]
	tree  tree.RedBlackTree[T]
	cmp   comparator.EqCmpKeyer[T, H]
}

func (s *concOrdered[T, H]) Add(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.table.TryAddHash(s.cmp.Key(t)) {
		s.tree.Insert(t)
	}
}

func (s *concOrdered[T, H]) Clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.tree.Clear()
	s.table.Clear()
}

func (s *concOrdered[T, H]) Delete(idx int) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if idx < 0 || idx >= s.tree.Size() {
		return false
	}
	if v, ok := s.tree.DeleteByIndex(idx); ok {
		s.table.TryRemoveHash(s.hash(v))
		return true
	}
	return false
}

func (s *concOrdered[T, H]) Equals(other set.Readable[T]) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s == other {
		return true
	}
	return eqHashableOrdered(s.table, other, s.cmp)
}

// Get returns the element at the given idx.
func (s *concOrdered[T, H]) Get(idx int) (value T, ok bool) {
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

// Has checks if the backingMap contains the element.
func (s *concOrdered[T, H]) Has(t T) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.tree.Contains(t)
}

func (s *concOrdered[T, H]) hash(v T) H {
	return s.cmp.Key(v)
}

func (s *concOrdered[T, H]) IndexOf(t T) int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s.table.HasHash(s.cmp.Key(t)) {
		it := s.tree.Iterator()
		for i := 0; it.HasNext(); i++ {
			if s.cmp.Equals(it.Next(), t) {
				return i
			}
		}
	}
	return -1
}

// IsEmpty checks if the backingMap is empty.
func (s *concOrdered[T, H]) IsEmpty() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.table.IsEmpty()
}

// Len returns the number of elements in the backingMap.
func (s *concOrdered[T, H]) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.table.Len()
}

func (s *concOrdered[T, H]) Range(callback func(value T) bool) {
	it := s.tree.Iterator()
	for i := 0; it.HasNext(); i++ {
		if !callback(it.Next()) {
			return
		}
	}
}

// Remove deletes an element from the backingMap.
func (s *concOrdered[T, H]) Remove(t T) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.table.TryRemoveHash(s.cmp.Key(t)) {
		s.tree.Delete(t)
	}
}

// Values returns the ordered elements in the backingMap.
func (s *concOrdered[T, H]) Elements() []T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	var values = make([]T, 0, s.Len())
	it := s.tree.Iterator()
	for it.HasNext() {
		values = append(values, it.Next())
	}
	return values
}

func (s *concOrdered[T, H]) Stream() <-chan T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	ch := make(chan T)
	it := s.tree.Iterator()
	go func() {
		defer close(ch)
		for it.HasNext() {
			ch <- it.Next()
		}
	}()

	return ch
}
