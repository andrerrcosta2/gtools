// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"maps"
	"sync"
)

// Comparable Creates a new str.Set[T] of comparable values
//
// This set doesn't hold any order of insSet
func Comparable[T comparable](values ...T) str.Set[T] {
	// Create a new cmpSet instance with an empty map.
	set := &cmpSet[T]{
		set: make(map[T]struct{}),
	}

	// Addf each value to the set.
	for _, v := range values {
		set.Add(v)
	}

	// Return the populated cmpSet instance.
	return set
}

type cmpSet[T comparable] struct {
	set map[T]struct{}
}

func (c *cmpSet[T]) Add(t T) {
	c.set[t] = struct{}{}
}

func (c *cmpSet[T]) Clear() {
	c.set = make(map[T]struct{})
}

func (c *cmpSet[T]) Equals(o str.Set[T]) bool {
	if c == o {
		return true
	}
	switch set := o.(type) {
	case *cmpSet[T]:
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

func (c *cmpSet[T]) Has(t T) bool {
	_, exists := c.set[t]
	return exists
}

func (c *cmpSet[T]) IsEmpty() bool {
	return len(c.set) == 0
}

func (c *cmpSet[T]) Len() int {
	return len(c.set)
}

func (c *cmpSet[T]) Remove(t T) {
	delete(c.set, t)
}

func (c *cmpSet[T]) Values() []T {
	values := make([]T, 0, len(c.set))
	for key := range c.set {
		values = append(values, key)
	}
	return values
}

var _ str.Set[any] = (*cmpSet[any])(nil)

// ConcComparable Creates a new thread-safe str.Set[T] of comparable values
//
// This set doesn't hold any order of insSet
func ConcComparable[T comparable](values ...T) str.Set[T] {
	set := &concCmpSet[T]{
		set: make(map[T]struct{}),
	}
	for _, v := range values {
		set.Add(v)
	}
	return set
}

type concCmpSet[T comparable] struct {
	mtx sync.RWMutex
	set map[T]struct{}
}

func (c *concCmpSet[T]) Add(t T) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.set[t] = struct{}{}
}

func (c *concCmpSet[T]) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.set = make(map[T]struct{})
}

func (c *concCmpSet[T]) Equals(o str.Set[T]) bool {
	switch set := o.(type) {
	case *concCmpSet[T]:
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

func (c *concCmpSet[T]) Has(t T) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	_, exists := c.set[t]
	return exists
}

func (c *concCmpSet[T]) IsEmpty() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return len(c.set) == 0
}

func (c *concCmpSet[T]) Len() int {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return len(c.set)
}

func (c *concCmpSet[T]) Remove(t T) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	delete(c.set, t)
}

func (c *concCmpSet[T]) Values() []T {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	values := make([]T, 0, len(c.set))
	for key := range c.set {
		values = append(values, key)
	}
	return values
}

var _ str.Set[any] = (*concCmpSet[any])(nil)
