// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import "errors"

var (
	ErrInvalidValue = errors.New("invalid type: <invalid>")
	ErrNotArray     = errors.New("target is not an array")
	ErrNotChan      = errors.New("target is not a channel")
	ErrNotFunc      = errors.New("target is not a function")
	ErrNotInterface = errors.New("target is not an interface")
	ErrNotMap       = errors.New("target is not a map")
	ErrNotPointer   = errors.New("target is not a pointer")
	ErrNotPrimitive = errors.New("target is not a primitive")
	ErrNotSlice     = errors.New("target is not a slice")
	ErrNotStruct    = errors.New("target is not a struct")
	ErrNotUnsafePtr = errors.New("target is not an unsafe.pointers")
)
