// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package pointers

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
	"unsafe"
)

// Name returns the kind string representation
func Name(value reflect.Type) (string, error) {
	if value.Kind() != reflect.Ptr {
		return "", reflect4.ErrNotPointer
	}
	return value.String(), nil
}

// NewOf creates a new pointer with a given reflect.Value as element.
func NewOf(x reflect.Value) reflect.Value {
	ptrCopy := reflect.New(x.Type())
	ptrCopy.Elem().Set(x)
	return ptrCopy
}

// AddrOf returns an uintptr representing the stable identity of a reflect.Value.
// It only supports pointer-like types Ptr, Slice, Map, Chan, Func, UnsafePointer).
// Returns an error if the kind doesn't naturally support a stable pointer.
func AddrOf(v reflect.Value) (uintptr, error) {
	switch v.Kind() {

	// These types support the Pointer() method which does give you the address of the underlying data for explicit
	// pointers, and golang's underlying pointers (slices, maps, etc).
	// However, it isn't always the same as the actual pointer variable you’d get in Go code,
	// but it's usually a valid unique handle for that instance during that reflection session.
	case reflect.Ptr, reflect.UnsafePointer, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.Pointer(), nil
	default:
		// Handle invalid values gracefully
		return 0, fmx.Errorf("pointers.AddrOf called by a non pointer kind: %v", v.Kind())
	}
}

// ForcedAddrOf tries to extract an uintptr identifier from any reflect.Value,
// using the most stable method available.
//
// WARNING: for non-pointer and non-addressable values, the result may be unstable
// and shouldn't be used for cycle detection or strict identity comparison.
//
//	Supported Cases:
//	- Invalid Values: it returns 0 and reflect4.ErrInvalidValue
//	- Pointer-Like Types (Support .Pointer() directly): returns its uintptr address.
//	- Addressable Values: if the value is addressable (value.CanAddr() == true)
//	  the function takes its address and returns the pointer:
//	For not addressable and not pointer-like types, it'll retrieve unstable addresses.
//	That means:
//	- Every time you call Forced on the same logical value using this fallback, you'll get a different pointer
func ForcedAddrOf(v reflect.Value) (uintptr, error) {
	if !v.IsValid() {
		// Handle invalid values gracefully
		return 0, reflect4.ErrInvalidValue
	}

	switch v.Kind() {

	// These types support the Pointer() method which does give you the address of the underlying data for explicit
	// pointers, and golang's underlying pointers as slices, maps, etc.
	// However, it isn't always the same as the actual pointer variable you’d get in Go code,
	// but it's usually a valid unique handle for that instance during that reflection session.
	case reflect.Ptr, reflect.UnsafePointer, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.Pointer(), nil

	default:
		// This case handles any value that's addressable, (has a real memory address).
		// That includes:
		// - Struct fields if exported
		// - Local variables
		// - Elements of addressable slices/arrays
		// - Interface-wrapped values where .Elem() is addressable
		// That means the output of this method is unique but only as long as the original object exists.
		// The returned pointer will:
		// - Be consistent within the same call to Forced.
		// - Still be stable during the current session, assuming the object isn't garbage collected or moved.
		//
		// It won't be stable:
		// - It could retrieve different addresses if a re-wrap of the same object if performed multiple times.
		// - Not guaranteed to match the address of the original variable in Go code due to escape
		//	 analysis/compiler optimizations.
		if v.CanAddr() {
			// Use unsafe.Addr for performance and clarity
			return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Pointer(), nil
		}
		// Fallback: create a new interface and take its address
		// Note: This is unstable and should be used cautiously
		interfacePtr := reflect.ValueOf(v.Interface()).Pointer()
		return interfacePtr, nil
	}
}
