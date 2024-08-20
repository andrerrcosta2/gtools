// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import "maps"

// Comparable Creates a new ComparableSet from a variable number of values.
// The set is initialized with the given values.
//
// This set doesn't hold any order of insertion
func Comparable[T comparable](values ...T) *ComparableSet[T] {
	// Create a new ComparableSet instance with an empty map.
	set := &ComparableSet[T]{
		set: make(map[T]struct{}),
	}

	// Add each value to the set.
	for _, v := range values {
		set.Add(v)
	}

	// Return the populated ComparableSet instance.
	return set
}

type ComparableSet[T comparable] struct {
	set map[T]struct{}
}

func (c *ComparableSet[T]) Has(t T) bool {
	_, exists := c.set[t]
	return exists
}

func (c *ComparableSet[T]) Add(t T) {
	c.set[t] = struct{}{}
}

func (c *ComparableSet[T]) Remove(t T) {
	delete(c.set, t)
}

func (c *ComparableSet[T]) Len() int {
	return len(c.set)
}

func (c *ComparableSet[T]) Values() []T {
	values := make([]T, 0, len(c.set))
	for key := range c.set {
		values = append(values, key)
	}
	return values
}

func (c *ComparableSet[T]) Clear() {
	c.set = make(map[T]struct{})
}

func (c *ComparableSet[T]) Equals(o Set[T]) bool {
	switch set := o.(type) {
	case *ComparableSet[T]:
		return maps.Equal(c.set, set.set)
	default:
		if c.Len() != set.Len() {
			return false
		}
		s := make(map[T]struct{})
		for _, v := range set.Values() {
			s[v] = struct{}{}
		}
		return maps.Equal(c.set, s)
	}
}
