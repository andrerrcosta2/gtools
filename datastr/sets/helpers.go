// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/domain/function/functions"
	"github.com/andrerrcosta2/gtools/core/sortables/sorts"
)

// Map takes a slice of elements of type K and a function that dicts each element of
// type K to an element of type V. It returns a ToSet of elements of type V.
// The function f is applied to each element of the slice and the results are added
// to the ToSet.
func Map[T any, S comparable](arr []T, f func(v T) S) set.Mutable[S] {
	s := Comparable[S]()
	for _, v := range arr {
		s.Add(f(v))
	}
	return s
}

// OfStrings creates a set from a slice of strings.
// It uses the Map function to create the ToSet.
func OfStrings(arr ...string) set.Mutable[string] {
	return Map(arr, functions.Identity[string])
}

func addToInsCmp[T any, H prim.Hashable](set insertable[T, H], t T) {
	hash := set.hash(t)
	if !set.contains(t) {
		// register hash
		set.setHash(hash, set.size())
		// Insert item at the next position
		set.append(t)
	}
}

// Snap Helper function to extract a Snap from any backingMap
func Snap[T constraints.Sortable[T]](other set.Mutable[T]) []T {
	if set, ok := other.(snapshotter[T]); ok {
		return set.snap()
	}

	// Fallback for non-concSortableOf sets
	return other.Elements()
}

func MarshalJSON[T any](set set.Mutable[T]) ([]byte, error) {
	cast, ok := set.(data.JSONSerializable)
	if !ok {
		panic("backingMap is not JSONSerializable")
	}
	return cast.MarshalJSON()
}

func UnmarshalJSON[T any](set set.Mutable[T], json []byte) (set.Mutable[T], error) {
	cast, ok := set.(data.JSONSerializable)
	if !ok {
		panic("backingMap is not JSONSerializable")
	}

	err := cast.UnmarshalJSON(json)
	return set, err
}

func Filter[T any, H comparable](cmp comparator.EqCmpKeyer[T, H], e ...T) []T {
	values := make(map[H]bool)
	output := make([]T, 0, len(e))

	for _, v := range e {
		h := cmp.Key(v)
		if _, ok := values[h]; !ok {
			values[h] = true
			output = append(output, v)
		}
	}
	return output
}

func CmpFilter[T comparable](e ...T) []T {
	values := make(map[T]bool)
	output := make([]T, 0, len(e))
	for _, v := range e {
		if _, ok := values[v]; !ok {
			values[v] = true
			output = append(output, v)
		}
	}
	return output
}

func OrderedFilter[S ~[]E, E prim.Ordered](e S, fn function.Consumer[E]) {
	sorts.Quick(e)
	var prev E
	first := true
	for _, value := range e {
		// just skip duplicates
		if first || value != prev {
			fn(value)
			prev = value
			first = false
		}
	}
}
