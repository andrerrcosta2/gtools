// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/objects"
	"github.com/andrerrcosta2/gtools/pkg/sorts"
	"github.com/andrerrcosta2/gtools/pkg/typers"
	"reflect"
)

// ComparatorOf returns a new ComparatorSortableOf instance for the given type K.
// This comparator is used to compare and hash values of type K that implement the gtools.SortableOf interface.
func ComparatorOf[K gtools.SortableOf]() *ComparatorSortableOf[K] {
	// Return a new instance of ComparatorSortableOf with the given type K.
	return &ComparatorSortableOf[K]{}
}

type ComparatorSortableOf[K gtools.SortableOf] struct {
}

func (s *ComparatorSortableOf[K]) Compare(a, b K) int {
	if a.Equal(b) {
		return 0
	}
	if a.Less(b) {
		return 1
	}
	return -1
}

func (s *ComparatorSortableOf[K]) Hash(sortable K) string {
	return Unique(sortable)
}

func (s *ComparatorSortableOf[K]) Equals(a, b K) bool {
	return a.Equal(b)
}

var _ comparables.KeyComparator[gtools.SortableOf, string] = (*ComparatorSortableOf[gtools.SortableOf])(nil)
var _ comparables.Comparator[gtools.SortableOf] = (*ComparatorSortableOf[gtools.SortableOf])(nil)

// Unique returns a unique string identifier for the given sortable object.
// If the object implements the gtools.PersistentSortableOf interface, its unique identifier is returned.
// Otherwise, the object's memory address or its string representation is returned.
func Unique[T any](sortable T) string {
	// Try to cast the sortable object to a gtools.PersistentSortableOf
	switch s := any(sortable).(type) {
	case gtools.PersistentSortableOf:
		// Return the unique identifier as a string
		return string(s.Unique())
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string,
		complex128, complex64, bool, gtools.SortableOf:
		return fmt.Sprintf("%v", sortable)
	case error:
		return s.Error()
	case *gtools.SortableOf, *gtools.PersistentSortableOf, *gtools.ComparableOf:
		return fmt.Sprintf("%p", s)
	default:
		val := reflect.ValueOf(sortable)
		if val.Kind() == reflect.Ptr {
			// Use reflect to get the pointer's address and format it as hexadecimal
			return fmt.Sprintf("%p", val.Interface())
		}
		return fmt.Sprintf("%v", sortable)
	}
}

// Sort sorts a slice of sortables using the provided sort algorithm.
//
// This function takes a slice of sortables and a sort algorithm, and returns the sorted slice.
// The sort algorithm must implement the sorts.Sort interface.
func Sort[T any](sortables *[]T, sort sorts.Sort[T]) {
	// Sort the slice using the provided sort algorithm
	sort.Sort(sortables)

}

// Equals checks if all the given values are equal.
// It supports both gtools.SortableOf and constraints.Ordered types.
// It returns true if all the values are equal, false otherwise.
func Equals[T constraints.Ordered](values ...interface{}) bool {
	// If there are less than 2 values, all values are equal
	if len(values) < 2 {
		return true
	}

	// Separate the values into two slices: ofs (gtools.SortableOf) and ords (constraints.Ordered)
	ofs, ords := typers.Ors[gtools.SortableOf, T](values...)

	// If all values are of type gtools.SortableOf, use EqualsOf
	if len(ords) == 0 && len(ofs) == len(values) {
		return EqualsOf[gtools.SortableOf](ofs...)
	}

	// If all values are of type constraints.Ordered, use objects.Equals
	if len(ofs) == 0 && len(ords) == len(values) {
		return objects.Equals(ords...)
	}

	// If there are values of different types, return false
	return false
}

// EqualsOf checks if all the given values are equal.
// It uses the Equal method of the gtools.SortableOf interface to compare values.
func EqualsOf[G gtools.SortableOf](values ...G) bool {
	// Use the objects.EqualsBy function to compare values using the Equal method
	return objects.EqualsBy(func(a, b G) bool {
		// Compare two values using the Equal method
		return a.Equal(b)
	}, values...)
}
