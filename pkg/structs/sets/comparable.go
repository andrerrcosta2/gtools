// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

type ComparableSet[T comparable] struct {
	set map[T]struct{}
}

func Comparable[T comparable]() *ComparableSet[T] {
	return &ComparableSet[T]{
		set: make(map[T]struct{}),
	}
}

func (c ComparableSet[T]) Has(t T) bool {
	_, exists := c.set[t]
	return exists
}

func (c ComparableSet[T]) Add(t T) {
	c.set[t] = struct{}{}
}

func (c ComparableSet[T]) Remove(t T) {
	delete(c.set, t)
}

func (c ComparableSet[T]) Len() int {
	return len(c.set)
}

func (c ComparableSet[T]) Values() []T {
	values := make([]T, 0, len(c.set))
	for key := range c.set {
		values = append(values, key)
	}
	return values
}
