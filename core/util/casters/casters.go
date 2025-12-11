// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package casters

import (
	"fmt"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/bins"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
)

type emptyInterface struct {
	typx unsafe.Pointer
	word unsafe.Pointer
}

// Assert casts the provided values to the type G and returns a slice of type G.
// Panics if any value is not of type G.
func Assert[G any](values ...any) (result []G) {
	result = make([]G, 0, len(values))
	for _, value := range values {
		if castValue, ok := Type[G](value); ok {
			result = append(result, castValue)
		} else {
			panic(fmt.Sprintf("Assert will panic if the provided value are not of the expected type: '%T'",
				value))
		}
	}
	return result
}

// Bytes casts the provided value to a byte slice and returns it.
// It returns an error if the value is not of a type that can be cast to a byte slice.
func Bytes(value any) ([]byte, error) {
	// Use the ToBytes function from the bins package to cast the value to a byte slice.
	// This function will return an error if the value is not of a type that can be cast to a byte slice.
	return bins.ToBytes(value)
}

// Map is just a functional style of casting
// it receives a value of F and casts to T
func Map[F any, T any](value F) T {
	return any(value).(T)
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

// Type casts the provided value to the type T and returns it.
func Type[T any](value any) (cast T, ok bool) {
	cast, ok = value.(T)
	return
}

// Types casts the provided values to the type G and returns a slice of type G.
// The second return value allMatches is true if all values were successfully cast
// to type G, and false otherwise.
func Types[G any](values ...any) (result []G, allMatches bool) {
	result = make([]G, 0, len(values))
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

// UnsafeValueOf casts the provided value to the type T - for reference-like variables use UnsafeReferenceOf -
// using unsafe ops and return it.
// It works only with interface{} types since it forcibly strips the type descriptor.
//
// Important considerations:
//   - The output pointer will discard the type descriptor of the interface{}.
//   - Lifetime problem: If you take a pointer to a value that’s allocated on the stack,
//     that pointer might become invalid when the stack frame ends (function returns).
//   - The GC might reclaim stack-allocated values if they're no longer referenced
//   - Because `unsafe.Pointer` isn't tracked by the GC, you must ensure the original "value"
//     is still considered live by the compiler.
//
// @Important: Keep the "value" parameter alive in the caller scope (or explicitly call runtime.KeepAlive)
// to prevent premature garbage collection of the underlying memory and possible corruption.
func UnsafeValueOf[T any](value any) T {
	// We cast the interface{} to its internal representation (emptyInterface)
	// then extract the data pointer to the underlying value.
	ei := (*emptyInterface)(unsafe.Pointer(&value))
	return *(*T)(ei.word)
}

// UnsafeReferenceOf casts the provided reference-like value - ptrs and golang natural ptrs - to the type T
// using unsafe ops and returns it.
// It works only with interface{} types since it forcibly strips the type descriptor.
//
// Important considerations:
//   - The output pointer will discard the type descriptor of the interface{}.
//   - Lifetime problem: If you take a pointer to a value that’s allocated on the stack,
//     that pointer might become invalid when the stack frame ends (function returns).
//   - The GC might reclaim stack-allocated values if they're no longer referenced
//   - Because `unsafe.Pointer` isn't tracked by the GC, you must ensure the original "value"
//     is still considered live by the compiler.
//
// @Important: Keep the "value" parameter alive in the caller scope (or explicitly call runtime.KeepAlive)
// to prevent premature garbage collection of the underlying memory and possible corruption.
func UnsafeReferenceOf[T any](v any) T {
	// Extract interface layout
	ei := (*emptyInterface)(unsafe.Pointer(&v))
	// Reinterpret the address of the word (pointer) as T
	return *(*T)(unsafe.Pointer(&ei.word))
}
