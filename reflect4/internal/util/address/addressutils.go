// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package address

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
)

// Of returns an uintptr representing the stable identity of a reflect.Value.
// It only supports pointer-like types - Ptr, slices, maps, channel, functions, UnsafePointer.
// Returns an error if the kind doesn't naturally support a stable pointer.
func Of(v reflect.Value) (uintptr, error) {
	switch v.Kind() {

	// These types support the pointers() method which does give you the address of the underlying data for explicit
	// ptrs, and golang's underlying ptrs.
	// However, it isn't always the same as the actual pointer variable you’d get in Go code,
	// but it's usually a valid unique handle for that instance during that reflection session.
	case reflect.Ptr, reflect.UnsafePointer, reflect.Map, reflect.Chan, reflect.Func, reflect.Slice:
		return v.Pointer(), nil
	default:
		return 0, fmx.Errorf("address.Of called by a non pointer kind: %v", v.Kind())
	}
}

// StableOf returns the memory address on which the pointer-like kind points to.
// Stable kinds are channels, maps, functions, ptrs, and unsafe ptrs.
// Despite slices and arrays be underlying ptrs in golang, their pointer method call
// is not reliable since they retrieve the same address of their first element.
func StableOf(v reflect.Value) (uintptr, error) {
	switch v.Kind() {
	case reflect.Chan, reflect.Map, reflect.Func, reflect.Ptr, reflect.UnsafePointer:
		if v.IsNil() {
			return 0, fmx.Errorf("address.StableOf: called by a nil '%v'", v.Kind())
		}
		return v.Pointer(), nil
	default:
		return 0, fmx.Errorf("address.StableOf: called by a non stable pointer kind '%v'", v.Kind())
	}
}

// Unsafe returns the memory address, as uintptr, of the underlying data
// in the given reflect.Value.
func Unsafe(v reflect.Value) (uintptr, error) {
	if !v.CanAddr() {
		return 0, fmx.Errorf("address.Unsafe: called by an unaddressable value: %v", v.Kind())
	}
	return v.UnsafeAddr(), nil
}

// UnsafeReadOnly returns the memory address (uintptr) of the underlying data
// // in the given reflect.Value.
//
// ⚠️ This function does not check for v.IsValid(), and the address returned
// for non-addressable values is from a clone — it should not be used for mutation.
// It also does not safely handle unexported fields and should not be used to
// create Go ptrs, as that breaks the Go memory model.
func UnsafeReadOnly(v reflect.Value) uintptr {
	if !v.CanAddr() {
		v = values.ForceOfUnaddr(v)
	}
	return v.UnsafeAddr()
}
