// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

// Map takes a slice of elements of type K and a function that maps each element of
// type K to an element of type V. It returns a ToSet of elements of type V.
// The function f is applied to each element of the slice and the results are added
// to the ToSet.
//
// Parameters:
// - arr: a slice of elements of type K
// - f: a function that maps an element of type K to an element of type V
//
// Returns:
// - A ToSet of elements of type V
func Map[T any, S comparable](arr []T, f func(v T) S) str.Set[S] {
	// Create an empty ToSet of type V
	s := Comparable[S]()
	// Iterate over the elements of the slice
	for _, v := range arr {
		// Apply the function f to each element and add the result to the ToSet
		s.Add(f(v))
	}
	// Return the ToSet
	return s
}

// OfStrings creates a ToSet from a slice of strings.
// It uses the Map function to create the ToSet.
func OfStrings(arr ...string) str.Set[string] {
	// Use the Map function to create the ToSet
	// The function takes a slice of strings and a function that maps each string to itself
	return Map(arr, func(v string) string {
		// Return the string itself
		return v
	})
}

func addToInsCmp[T any, H prim.Hashable](set insertionSet[T, H], t T) {
	hash := set.hash(t)
	if !set.contains(t) {
		// register the unique
		set.setHash(hash, set.size())
		// Insert item at the next position
		set.append(t)
	}
}

// Snap Helper function to extract a Snap from any set
func Snap[T gtools.SortableOf](other str.Set[T]) ([]T, map[string]struct{}) {
	if set, ok := other.(*concSortableOf[T]); ok {
		return set.snap()
	}

	// Fallback for non-concSortableOf sets
	values := other.Values()
	sortable := SortableOf(values...).(*sortableOfSet[T])
	return sortable.items, sortable.index
}

func MarshalJSON[T any](set str.Set[T]) ([]byte, error) {
	cast, ok := set.(data.JSONSerializable)
	if !ok {
		panic("set is not JSONSerializable")
	}
	return cast.MarshalJSON()
}

func UnmarshalJSON[T any](set str.Set[T], json []byte) (str.Set[T], error) {
	cast, ok := set.(data.JSONSerializable)
	if !ok {
		panic("set is not JSONSerializable")
	}

	err := cast.UnmarshalJSON(json)
	return set, err
}
