// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package proto

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"reflect"
)

type Prototype struct {
	Methods []Method    // For interface
	Fields  []Prototype // For struct types

	PkgName   string // Package name
	TypeName  string // Name of the type
	FieldName string // Name of the field
	Tag       string // struct field tag

	Kind       reflect.Kind            // Kind of type (Struct, Slice, Array, etc.)
	Key        *Prototype              // For map
	Value      *Prototype              // For slice/array/map, or pointer element type
	EmptyValue functions.Supplier[any] // Empty constructor

	Offset uint32 // Offset of the field in the struct

	CanSet    bool // Whether this field can be set
	IsPointer bool // Whether this type is a pointer
	Vararg    bool // Whether this type is a vararg
	Embedded  bool // Whether this field is embedded
}

type Method struct {
	Params  []Prototype // Parameters (Types of input parameters)
	Returns []Prototype // Return types

	Name     string // Method name
	Receiver string // Receiver type (e.g., "MyType" or "*MyType")

	Func interface{} // Pointer to the function for lazy reflection
}
