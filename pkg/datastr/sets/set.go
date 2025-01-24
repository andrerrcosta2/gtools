// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import "github.com/andrerrcosta2/gtools/core/data/str"

// Map takes a slice of elements of type T and a function that maps each element of
// type T to an element of type S. It returns a Set of elements of type S.
// The function f is applied to each element of the slice and the results are added
// to the Set.
//
// Parameters:
// - arr: a slice of elements of type T
// - f: a function that maps an element of type T to an element of type S
//
// Returns:
// - A Set of elements of type S
func Map[T any, S comparable](arr []T, f func(v T) S) str.Set[S] {
	// Create an empty Set of type S
	s := Comparable[S]()
	// Iterate over the elements of the slice
	for _, v := range arr {
		// Apply the function f to each element and add the result to the Set
		s.Add(f(v))
	}
	// Return the Set
	return s
}

// NewString creates a Set from a slice of strings.
// It uses the Map function to create the Set.
func NewString(arr ...string) str.Set[string] {
	// Use the Map function to create the Set
	// The function takes a slice of strings and a function that maps each string to itself
	return Map(arr, func(v string) string {
		// Return the string itself
		return v
	})
}
