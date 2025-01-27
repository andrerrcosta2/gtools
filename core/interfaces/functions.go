// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interfaces

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

// ImplementationOf checks if the given data implements the given interface T.
// The meaning of this method is to avoid primitives which can always implement
// interfaces with no constraints, while working with generics.
//
// It returns true if the data implements the interface T,
//
// Parameters:
//   - data: A value of any type.
//
// Returns:
//   - bool: A boolean indicating if the data implements the interface T.
//
// Example:
//
//	var myInterface interface{}
//	if interfaces.ImplementationOf[io.Closer](&myInterface) {
//	    fmt.Println("myValue implements io.Closer")
//	} else {
//	    fmt.Println("myValue does not implement io.Closer")
//	}
func ImplementationOf[T any](data any) bool {
	// If the data is a primitive, it shouldn't be handled
	// by the interfaces package because it could mislead
	// logics
	if prim.IsPrimitive(data) {
		return false
	}

	// If the data implements the interface, return true
	if _, ok := data.(T); ok {
		return true
	}

	return false
}
