// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import "github.com/andrerrcosta2/gtools/pkg/reflect4/internal/handlers/data"

//import (
//	"github.com/andrerrcosta2/gtools/core/util/reflectiv/internal/handlers/data"
//	"github.com/andrerrcosta2/gtools/core/util/reflectiv/internal/proto"
//)
//
//// PrototypeOf returns the Prototype representation of a given interf of type `T`.
//// It extracts metadata about the type and its nested types if `deep` is true.
//// The function supports both deep and shallow extraction.
//func PrototypeOf[T any](deep bool) (*Prototype, error) {
//	// Declare a variable of type T
//	var v T
//
//	// Perform deep extraction if the flag is set
//	if deep {
//		// Use the proto.Deep function to extract a detailed Type
//		return proto.Deep(v)
//	}
//
//	// Perform shallow extraction otherwise
//	// Use the proto.Shallow function to extract a basic Type
//	return proto.Shallow(v), nil
//}
//

// CopyOf creates a copy of a given interf `v`.
// If `deep` is true, it performs a deep copy of the interf.
// Otherwise, it performs a shallow copy.
func CopyOf[T any](v T, deep bool) (T, error) {
	if deep {
		// Perform a deep copy of the interf
		return data.DeepCopy(v)
	}

	// Perform a shallow copy of the interf
	return data.ShallowCopy(v)
}
