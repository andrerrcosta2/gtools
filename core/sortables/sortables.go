// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"unsafe"
)

// ComparatorOf returns a new Comparator instance for the given type K.
// This comparator is used to compare and hash values of type K that implement the core.SortableOf interface.
func ComparatorOf[K gtools.SortableOf]() *Comparator[K] {
	// Return a new instance of Typed with the given type K.
	return &Comparator[K]{}
}

type Comparator[K gtools.SortableOf] struct {
}

func (s *Comparator[K]) Compare(a, b K) int {
	if a.Equal(b) {
		return 0
	}
	if a.Less(b) {
		return 1
	}
	return -1
}

func (s *Comparator[K]) Hash(sortable K) string {
	return Unique[K](sortable)
}

func (s *Comparator[K]) Equals(a, b K) bool {
	return a.Equal(b)
}

var _ comparators.KeyTyped[gtools.SortableOf, string] = (*Comparator[gtools.SortableOf])(nil)
var _ comparators.Typed[gtools.SortableOf] = (*Comparator[gtools.SortableOf])(nil)

// Equality returns true if the two values are implemented equal, false otherwise.
// It uses the Equal method of the ComparableOf interface to compare values.
func Equality[T gtools.ComparableOf](x T, y T) bool {
	// Check if the values are equal
	// Use the Equal method of the ComparableOf interface to compare values
	// If the values are equal, return true
	// Otherwise, return false
	return x.Equal(y)
}

var _ functions.BiPredicate[gtools.SortableOf, gtools.SortableOf] = Equality[gtools.SortableOf]

// MultipleEquality returns true if all the values are equal, false otherwise.
// It uses the Equal method of the ComparableOf interface to compare values.
func MultipleEquality[T gtools.ComparableOf](values ...T) bool {
	// Check if the values are equal
	// Iterate over the values and compare each pair of adjacent values
	// If any pair of values is not equal, return false
	for i := 0; i < len(values)-1; i++ {
		if !values[i].Equal(values[i+1]) {
			return false
		}
	}
	// If all values are equal, return true
	return true
}

var _ functions.VarPredicate[gtools.SortableOf] = MultipleEquality[gtools.SortableOf]

// TryEquality checks if two values are equal without deep reflection.
// The method assumes that x and y are of the same type, otherwise it returns false.
// It follows the order:
//
// 1. ImplementationOf ComparableOf
// 2. Is a natural comparable type
// 3. Uses pointer comparison for reference types
func TryEquality[T any](x, y any) bool {
	// If x or y implements ComparableOf
	if comparableX, ok := x.(gtools.ComparableOf); ok {
		return comparableX.Equal(y)
	}
	if comparableY, ok := y.(gtools.ComparableOf); ok {
		return comparableY.Equal(x)
	}

	// If x and y are from the primitive Comparable types
	if equal, err := prim.Compare(x, y); err == nil {
		return equal
	}

	// it makes some sense though
	if _, ok := x.(T); ok {
		if _, ok = y.(T); ok {
			return x == y
		}
	}

	// Pointer comparison
	xPtr, okX := x.(*T)
	yPtr, okY := y.(*T)

	// this is handful when a variable of an implementation
	// is initialized as interface.
	if okX && okY {
		// Keeping the equality check straightforward makes it easier to understand
		// and maintain. When dealing with pointers, especially nil pointers, their
		// behavior is well-defined: two nil pointers are always considered equal.
		if xPtr == nil && yPtr == nil {
			return true
		}

		return xPtr == yPtr // Compare the dereferenced pointers
	}

	// Lost scope
	return false
}

// TryLess checks if x is less than y without deep reflection.
// The method assumes that x and y are of the same type, otherwise it returns false.
// It follows the order:
//
// 1. ImplementationOf SortableOf
// 2. Is a natural sortable type
//
// If you need a different scope don't use this method
func TryLess(x, y any) bool {
	// If x or y implements ComparableOf
	if comparableX, ok := x.(gtools.SortableOf); ok {
		return comparableX.Less(y)
	}
	if comparableY, ok := y.(gtools.SortableOf); ok {
		return comparableY.Less(x)
	}

	// If x and y are from the primitive Comparable types
	if less, err := prim.TryLess(x, y); err == nil {
		return less
	}

	// Lost scope
	return false
}

// Unique returns a unique string identifier for the given object.
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
func Unique[T any](sortable any) string {
	switch s := sortable.(type) {
	case gtools.UniqueOf:
		// Return the unique identifier as a string
		if s == nil {
			return "<nil>"
		}
		return s.Unique()
	case error:
		if s == nil {
			return "<nil>"
		}
		return s.Error()
	case *T:
		return fmt.Sprintf("0x%x", uintptr(unsafe.Pointer(s)))
	case int:
		return fmt.Sprintf("<int>%d", s)
	case int8:
		return fmt.Sprintf("<int8>%d", s)
	case int16:
		return fmt.Sprintf("<int16>%d", s)
	case int32:
		return fmt.Sprintf("<int32>%d", s)
	case int64:
		return fmt.Sprintf("<int64>%d", s)
	case uint:
		return fmt.Sprintf("<uint>%d", s)
	case uint8:
		return fmt.Sprintf("<uint8>%d", s)
	case uint16:
		return fmt.Sprintf("<uint16>%d", s)
	case uint32:
		return fmt.Sprintf("<uint32>%d", s)
	case uint64:
		return fmt.Sprintf("<uint64>%d", s)
	case uintptr:
		return fmt.Sprintf("<uintptr>%d", s)
	case float32:
		return fmt.Sprintf("<float32>%g", s)
	case float64:
		return fmt.Sprintf("<float64>%g", s)
	case string:
		return fmt.Sprintf("<string>%s", s)
	case bool:
		return fmt.Sprintf("<bool>%t", s)
	case complex64:
		return fmt.Sprintf("<complex64>%v", s)
	case complex128:
		return fmt.Sprintf("<complex128>%v", s)
	default:
		// I guess it still can be a pointer if an interface is passed
		// as generic type, so it should be reflected to return the memory address
		// Here it is returning its value
		return fmt.Sprintf("<%T>%v", sortable, sortable)
	}
}
