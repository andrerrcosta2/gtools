// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/search"
	"maps"
	"sort"
)

// Ordered creates a new OrderedSet.
//
// This function returns a pointer to a new OrderedSet, which is a set that maintains the order of its elements.
//
// Type parameter T must satisfy the constraints.Ordered constraint, meaning it must be a type that supports ordering.
func Ordered[T constraints.Ordered](values ...T) *OrderedSet[T] {
	// Sort the values in ascending order.
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	// Create a new OrderedSet with an empty map and slice.
	set := &OrderedSet[T]{
		// The set is used to keep track of the elements in the set for efficient lookups.
		index: make(map[T]struct{}),
		// The keys slice is used to maintain the order of the elements.
		items: make([]T, len(values)),
	}

	// Populate the index map
	for i, value := range values {
		set.index[value] = struct{}{}
		set.items[i] = value
	}

	// Return the populated OrderedSet instance.
	return set
}

type OrderedSet[T constraints.Ordered] struct {
	index map[T]struct{}
	items []T
}

// Has checks if the set contains the element.
func (o *OrderedSet[T]) Has(t T) bool {
	_, exists := o.index[t]
	return exists
}

// Add inserts an element into the set if it doesn't already exist.
func (o *OrderedSet[T]) Add(t T) {
	if !o.Has(t) {
		// Binary search for insertion point
		pos := search.Binary(o.items, t)
		o.items = append(o.items[:pos], append([]T{t}, o.items[pos:]...)...)
		o.index[t] = struct{}{}
	}
}

// Remove deletes an element from the set.
func (o *OrderedSet[T]) Remove(t T) {
	if o.Has(t) {
		// Binary search for the position of the item
		pos := search.Binary(o.items, t)
		// Remove the item
		o.items = append(o.items[:pos], o.items[pos+1:]...)
		delete(o.index, t)
	}
}

// Len returns the number of elements in the set.
func (o *OrderedSet[T]) Len() int {
	return len(o.items)
}

// Values returns the ordered elements in the set.
func (o *OrderedSet[T]) Values() []T {
	return o.items
}

// Get returns the element at the given index.
func (o *OrderedSet[T]) Get(i int) (T, bool) {
	if !arrays.OutOfBounds(o.items, i) {
		return o.items[i], true
	}
	var zeroValue T
	return zeroValue, false
}

// Exclude removes an element at the given index.
func (o *OrderedSet[T]) Exclude(i int) bool {
	if !arrays.OutOfBounds(o.items, i) {
		o.items = append(o.items[:i], o.items[i+1:]...)
		return true
	}
	return false
}

func (o *OrderedSet[T]) Loop() <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for _, item := range o.items {
			ch <- item
		}
	}()

	return ch
}

func (o *OrderedSet[T]) Clear() {
	o.items = make([]T, 0)
	o.index = make(map[T]struct{})
}

func (o *OrderedSet[T]) Equals(other Set[T]) bool {
	if o.Len() != other.Len() {
		return false
	}
	switch set := other.(type) {
	case *OrderedSet[T]:
		return arrays.Equals[T](o.items, set.items) && maps.Equal(o.index, set.index)
	default:
		return arrays.Equals[T](o.items, set.Values())
	}
}
