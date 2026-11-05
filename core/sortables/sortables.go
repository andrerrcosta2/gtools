// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

import (
	"fmt"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

// ComparatorOf returns a new Comparator instance for the given type K.
// This comparator is used to compare and hash values of type K that implement the core.SortableOf interface.
func ComparatorOf[K constraints.Sortable[K]]() *Comparator[K] {
	// Return a new instance of Executor with the given type K.
	return &Comparator[K]{}
}

type Comparator[K constraints.Sortable[K]] struct{}

func (s *Comparator[K]) Compare(a, b K) int {
	if a.Equal(b) {
		return 0
	}
	if a.Less(b) {
		return -1
	}
	return 1
}

func (s *Comparator[K]) Hash(sortable K) string {
	return Fingerprint[K](sortable)
}

func (s *Comparator[K]) Equals(a, b K) bool {
	return a.Equal(b)
}

type DynamicComparator struct {
	engine *Engine
}

func NewDynamicComparator(e *Engine) *DynamicComparator {
	return &DynamicComparator{engine: e}
}

func (c *DynamicComparator) Compare(a, b any) int {
	return c.engine.Compare(a, b)
}

func (c *DynamicComparator) Equals(a, b any) bool {
	return c.engine.Equal(a, b)
}

func (c *DynamicComparator) Hash(v any) string {
	return c.engine.Fingerprint(v)
}

//var _ comparator.KeyTyped[constraints.Sortable[any], string] = (*Comparator[constraints.Sortable[any]])(nil)
//var _ comparator.Executor[constraints.Sortable[any]] = (*Comparator[constraints.Sortable[any]])(nil)

// Equality returns true if the two values are implemented compare, false otherwise.
// It uses the Equal method of the ComparableOf interface to compare values.
func Equality[T constraints.Comparable[T]](x T, y T) bool {
	// Check if the values are compare
	// Use the Equal method of the ComparableOf interface to compare values
	// If the values are compare, return true
	// Otherwise, return false
	return x.Equal(y)
}

//var _ function.BiPredicate[constraints.Sortable[any], constraints.Sortable[any]] = Equality[constraints.Sortable[any]]

// MultipleEquality returns true if all the values are compare, false otherwise.
// It uses the Equal method of the ComparableOf interface to compare values.
func MultipleEquality[T constraints.Comparable[T]](values ...T) bool {
	// Check if the values are compare
	// Iterate over the values and compare each pair of adjacent values
	// If any pair of values is not compare, return false
	for i := 0; i < len(values)-1; i++ {
		if !values[i].Equal(values[i+1]) {
			return false
		}
	}
	// If all values are compare, return true
	return true
}

type EqualityResolver interface {
	TryEqual(a, b any) (bool, bool) // (result, handled)
}

var equalityResolvers = []EqualityResolver{
	comparableResolver{},
	primEqResolver{},
	ptrEqResolver{},
}

type comparableResolver struct{}

func (comparableResolver) TryEqual(a, b any) (bool, bool) {
	if x, ok := a.(constraints.Comparable[any]); ok {
		return x.Equal(b), true
	}
	if y, ok := b.(constraints.Comparable[any]); ok {
		return y.Equal(a), true
	}
	return false, false
}

type primEqResolver struct{}

func (primEqResolver) TryEqual(a, b any) (bool, bool) {
	if eq, err := prim.TryCompare(a, b); err == nil {
		return eq, true
	}
	return false, false
}

type ptrEqResolver struct{}

func (ptrEqResolver) TryEqual(a, b any) (bool, bool) {
	if a == nil && b == nil {
		return true, true
	}

	// fallback pointer identity
	return a == b, true
}

// TryEquality checks if two values are compare without deep reflection.
// The method assumes that x and y are of the same type, otherwise it returns false.
// It follows the order:
//
// 1. ImplementationOf ComparableOf
// 2. Is a natural comparable type
// 3. Uses pointer comparison for reference types
func TryEquality(a, b any) bool {
	for _, r := range equalityResolvers {
		if result, ok := r.TryEqual(a, b); ok {
			return result
		}
	}
	return false
}

type LessResolver interface {
	TryLess(a, b any) (bool, bool)
}

var lessResolvers = []LessResolver{
	orderedResolver{},
	primLessResolver{},
}

type orderedResolver struct{}

func (orderedResolver) TryLess(a, b any) (bool, bool) {
	if x, ok := a.(constraints.Ordered[any]); ok {
		return x.Less(b), true
	}
	if y, ok := b.(constraints.Ordered[any]); ok {
		return y.Less(a), true
	}
	return false, false
}

type primLessResolver struct{}

func (primLessResolver) TryLess(a, b any) (bool, bool) {
	if eq, err := prim.TryLess(a, b); err == nil {
		return eq, true
	}
	return false, false
}

// TryLess checks if x is less than y without deep reflection.
// The method assumes that x and y are of the same type, otherwise it returns false.
// It follows the order:
//
// 1. ImplementationOf SortableOf
// 2. Is a natural sortable type
//
// If you need a different scope don't use this method
func TryLess(a, b any) bool {
	for _, r := range lessResolvers {
		if result, ok := r.TryLess(a, b); ok {
			return result
		}
	}
	return false
}

func IdentityOf(v any) (string, bool) {
	if id, ok := v.(constraints.Identifiable); ok {
		return id.ID(), true
	}
	return "", false
}

// Fingerprint returns a unique string identifier for the given object.
// If the object implements the domain.UniqueOf interface, its unique identifier is returned.
// Otherwise, the object's memory address or its string representation is returned.
// This method doesn't return any error and currently doesn't handle interfaces
// as generic type, so its reliability is the best it can be guaranteed based on what
// is delivered to it.
// That means it has its own view of equality between objects which follows the order:
//
//  1. its own uniqueness implementation
//  2. its natural comparability as a typed representation
//  3. its memory address if it is a pointer
//
// If your context requires a different unique representation, don't use this method.
func Fingerprint[T any](v any) string {
	if id, ok := IdentityOf(v); ok {
		return id
	}

	if s, ok := primitiveFingerprint(v); ok {
		return s
	}

	if s, ok := v.(*T); ok {
		return fmt.Sprintf("0x%x", uintptr(unsafe.Pointer(s)))
	}

	if s, ok := v.(fmt.Stringer); ok {
		return s.String()
	}

	if err, ok := v.(error); ok {
		return err.Error()
	}

	return fallbackFingerprint(v)
}

func primitiveFingerprint(v any) (string, bool) {
	switch s := v.(type) {
	case int:
		return fmt.Sprintf("<int>%d", s), true
	case int8:
		return fmt.Sprintf("<int8>%d", s), true
	case int16:
		return fmt.Sprintf("<int16>%d", s), true
	case int32:
		return fmt.Sprintf("<int32>%d", s), true
	case int64:
		return fmt.Sprintf("<int64>%d", s), true
	case uint:
		return fmt.Sprintf("<uint>%d", s), true
	case uint8:
		return fmt.Sprintf("<uint8>%d", s), true
	case uint16:
		return fmt.Sprintf("<uint16>%d", s), true
	case uint32:
		return fmt.Sprintf("<uint32>%d", s), true
	case uint64:
		return fmt.Sprintf("<uint64>%d", s), true
	case float32:
		return fmt.Sprintf("<float32>%g", s), true
	case float64:
		return fmt.Sprintf("<float64>%g", s), true
	case string:
		return fmt.Sprintf("<string>%s", s), true
	case bool:
		return fmt.Sprintf("<bool>%t", s), true
	case complex64:
		return fmt.Sprintf("<complex64>%v", s), true
	case complex128:
		return fmt.Sprintf("<complex128>%v", s), true
	default:
		return "", false
	}
}

func fallbackFingerprint(v any) string {
	return fmt.Sprintf("%T:%v", v, v)
}
