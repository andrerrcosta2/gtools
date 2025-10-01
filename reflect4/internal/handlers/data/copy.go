// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
)

// DeepCopy deep copies the given value.
// The function returns a deep clone of the value and an error.
// The error is nil if the clone was successful.
func DeepCopy[O internal.Option](value reflect.Value, o ...O) (cp reflect.Value, err error) {
	// Call the deepCopy function with the pointer map to avoid cyclic references
	cp, err = deepCopy(value, NewCopyStrategy(o...))
	if err != nil {
		return values.Empty, fmx.Errorf("failed to deep clone value: %s%v: '%v'",
			value.String(), value.Interface(), err)
	}
	return cp, nil
}

// ShallowCopy TODO: unimplemented
func ShallowCopy(value reflect.Value, _ *CopyStrategy) (reflect.Value, error) {
	panic("shallow clone not yet implemented")
	return value, nil
}

// DeepCopyArray performs a deep clone of the given array based on the given
// CopyStrategy.
// It panics if the value isn't an array
func DeepCopyArray[O internal.Option](v reflect.Value, o ...O) (reflect.Value, error) {
	return deepCopyArray(v, NewCopyStrategy(o...))
}

// DeepCopyChannel performs a deep clone of a given channel based on the given
// CopyStrategy.
// It panics if the value isn't a channel
func DeepCopyChannel[O internal.Option](v reflect.Value, o ...O) (reflect.Value, error) {
	return NewCopyStrategy(o...).CopyChan(v, nil)
}

// DeepCopyInterface performs a deep clone of a given interface based on the given
// CopyStrategy.
// It panics if the value isn't an interface
func DeepCopyInterface[O internal.Option](v reflect.Value, o ...O) (reflect.Value, error) {
	return deepCopyInterface(v, NewCopyStrategy(o...))
}

// DeepCopyMap performs a deep clone of a given map based on the given
// CopyStrategy.
// It panics if the value isn't a map
func DeepCopyMap[O internal.Option](v reflect.Value, o ...O) (reflect.Value, error) {
	return deepCopyMap(v, NewCopyStrategy(o...))
}

// DeepCopyPointer performs a deep clone of the given pointer based on the given
// CopyStrategy.
// It panics if the value isn't a pointer
func DeepCopyPointer[O internal.Option](v reflect.Value, o ...O) (reflect.Value, error) {
	s := NewCopyStrategy(o...)
	return s.CopyPtr(v, s)
}

// DeepCopySlice performs a deep clone of the given slice based on the given
// CopyStrategy.
// It panics if the value isn't a slice
func DeepCopySlice[O internal.Option](v reflect.Value, o ...O) (reflect.Value, error) {
	return deepCopySlice(v, NewCopyStrategy(o...))
}

// DeepCopyStruct performs a deep clone of the given struct based on the given
// CopyStrategy
// It panics if the value isn't a struct
func DeepCopyStruct[O internal.Option](v reflect.Value, o ...O) (reflect.Value, error) {
	return deepCopyStruct(v, NewCopyStrategy(o...))
}

// ShallowCopyFunc does the clone we can do in golang.
//
//	functions are shallow-copied by reference.
//
// If the function is a closure, the captured state is shared between the original and clone.
// There's no way to isolate or deep-clone the closure environment in Go.
// Therefore, you should avoid to clone structures containing closure functions
// if their shared states can lead to inconsistency
func ShallowCopyFunc(v reflect.Value) (reflect.Value, error) {
	if v.Kind() != reflect.Func {
		return values.Empty, reflect4.ErrNotFunc
	}
	return shallowCopyFunc(v)
}

func ShallowCopyPointer(v reflect.Value) (reflect.Value, error) {
	if v.Kind() != reflect.Pointer {
		return values.Empty, reflect4.ErrNotPointer
	}
	return identityCopy(v, nil)
}
