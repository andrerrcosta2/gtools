// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package casters

import (
	"fmt"
)

// Type casts the provided value to the type T and returns it.
// It just feels stupid to have a casters package and do not have
// its most basic use cases.
func Type[T any](value any) (cast T, ok bool) {
	cast, ok = value.(T)
	return
}

// Types casts the provided values to the type G and returns a slice of type G.
// The second return value allMatches is true if all values were successfully cast
// to type G, and false otherwise.
func Types[G any](values ...any) (result []G, allMatches bool) {
	allMatches = true
	for _, value := range values {
		if castValue, ok := Type[G](value); ok {
			result = append(result, castValue)
		} else {
			allMatches = false
		}
	}
	return
}

// AssertedTypes casts the provided values to the type G and returns a slice of type G.
// Panics if any value is not of type G.
func AssertedTypes[G any](values ...any) []G {
	var result []G
	for _, value := range values {
		if castValue, ok := Type[G](value); ok {
			result = append(result, castValue)
		} else {
			panic(fmt.Sprintf("AssertedTypes will panic if the provided value are not of the expected type: %T", value))
		}
	}
	return result
}

// Primitive casts the provided value to the type T and returns it.
// T must be a primitive type.
// If the value is not of type G, an error is returned.
func Primitive[T prim.Any](value any) (T, error) {
	return prim.ToPrimitive[T](value)
}

// PrimitiveOrdered casts the provided value to the type T and returns it.
// T must be a primitive ordered type.
// If the value is not of type G, an error is returned.
func PrimitiveOrdered[T prim.Ordered](value any) (T, error) {
	return prim.ToOrdered[T](value)
}

// PrimitiveComparable casts the provided value to the type T and returns it.
// T must be a primitive comparable type.
// If the value is not of type G, an error is returned.
func PrimitiveComparable[T prim.Comparable](value any) (T, error) {
	return prim.ToComparable[T](value)
}

// NaturalComparable casts the provided values to the type T as value.
//
// Natural comparable types are the types that are comparable by default, such as
// primitives, strings, booleans, and complex numbers. See the documentation for
// the comparable type for more information. https://go.dev/blog/comparable
func NaturalComparable[C comparable](value any) (cmp C, ok bool) {
	// Cast the value to the type C and check if it's ok.
	// If the value is already of type C, this will just return the value.
	if cmp, ok = value.(C); ok {
		// If the value is a pointer to C, dereference the pointer and return the value.
		if ptr, isPointer := any(cmp).(*C); isPointer && ptr != nil {
			cmp = *ptr
		}
	}
	return
}

// Numeric casts the provided value to the type T and returns it.
// T must be a numeric type.
// If the value is not of type T, an error is returned.
func Numeric[T nums.Any](value any) (T, error) {
	return nums.ToNumeric[T](value)
}

// Bytes casts the provided value to a byte slice and returns it.
// It returns an error if the value is not of a type that can be cast to a byte slice.
func Bytes(value any) ([]byte, error) {
	// Use the ToBytes function from the bins package to cast the value to a byte slice.
	// This function will return an error if the value is not of a type that can be cast to a byte slice.
	return bins.ToBytes(value)
}
