// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal/pointers"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
	"unsafe"
)

// DeepCopy deep copies the given value.
// The function returns a deep copy of the value and an error.
// The error is nil if the copy was successful.
func DeepCopy(value reflect.Value) (cp reflect.Value, err error) {
	fmx.Greenf("\n[DeepCopy]: %s\n", value.String())
	// Call the deepCopy function with the pointer map to avoid cyclic references
	cp, err = deepCopy(value, tracker.Reference())
	if err != nil {
		return values.Empty, fmx.Errorf("error while deep copying %s%v: '%v'", value.String(), value.Interface(), err)
	}
	return cp, nil
}

// ShallowCopy TODO: unimplemented
func ShallowCopy(value reflect.Value) (reflect.Value, error) {
	return value, nil
}

func deepCopy(v reflect.Value, t *tracker.RefTracker) (reflect.Value, error) {
	if v.Kind() == reflect.Invalid {
		return values.Empty, reflect4.ErrInvalidValue
	}

	switch v.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		return v, nil
	case reflect.Array:
		return deepCopyArray(v, t)
	case reflect.Slice:
		return deepCopySlice(v, t)
	case reflect.Chan:
		return deepCopyChan(v)
	case reflect.Map:
		return deepCopyMap(v, t)
	case reflect.Struct:
		return deepCopyStruct(v, t)
	case reflect.Ptr:
		return deepCopyPointer(v, t)
	case reflect.Interface:
		return deepCopyInterface(v, t)
	case reflect.Func:
		return shallowCopyFunc(v)
	default:
		return values.Empty, fmt.Errorf("unsupported kind: %v", v.Kind())
	}
}

// DeepCopyArray copies an array deeply
func DeepCopyArray(v reflect.Value) (reflect.Value, error) {
	if v.Kind() != reflect.Array {
		return values.Empty, fmt.Errorf("deep array copy called with a non array value: '%v'", v.Kind())
	}
	return deepCopyArray(v, tracker.Reference())
}

// deepCopyArray is a helper function for array copying.
func deepCopyArray(v reflect.Value, t *tracker.RefTracker) (reflect.Value, error) {
	// As arrays sizes are part of its type definitions in golang,
	// using reflect's "New" is enough to define its size.
	cp := reflect.New(v.Type()).Elem()
	for i := 0; i < v.Len(); i++ {
		elem, err := deepCopy(v.Index(i), t)
		if err != nil {
			return values.Empty, err
		}
		cp.Index(i).Set(elem)
	}
	return cp, nil
}

func deepCopyChan(v reflect.Value) (reflect.Value, error) {
	// As channels are not addressable, we can't use reflect's "New" to create a new one.
	// Instead, we create a new one using the "make" function.
	typ := v.Type()
	capacity := v.Cap() // only works for buffered channels
	newChan := reflect.MakeChan(typ, capacity)
	return newChan, nil
}

func ShallowCopyFunc(v reflect.Value) (reflect.Value, error) {
	if v.Kind() != reflect.Func {
		return values.Empty, reflect4.ErrNotFunc
	}
	return shallowCopyFunc(v)
}

func shallowCopyFunc(v reflect.Value) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	return reflect.Zero(v.Type()), nil
}

func DeepCopyInterface(v reflect.Value) (reflect.Value, error) {
	if v.Kind() != reflect.Interface {
		return values.Empty, reflect4.ErrNotInterface
	}
	return deepCopyInterface(v, tracker.Reference())
}

func deepCopyInterface(v reflect.Value, t *tracker.RefTracker) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	elem := v.Elem()
	cp, err := deepCopy(elem, t)
	if err != nil {
		return values.Empty, err
	}
	return reflect.ValueOf(cp.Interface()).Convert(v.Type()), nil
}

func DeepCopyMap(v reflect.Value) (reflect.Value, error) {
	if v.Kind() != reflect.Map {
		return values.Empty, fmt.Errorf("deep map copy called with a non map value: '%v'", v.Kind())
	}
	return deepCopyMap(v, tracker.Reference())
}

// deepCopyMap deep copies a map, including keys and values.
func deepCopyMap(v reflect.Value, t *tracker.RefTracker) (reflect.Value, error) {
	// maps are underlying pointers
	cache, exists, err := t.Get(v)
	if err != nil {
		return values.Empty, err
	}
	if exists {
		return cache, nil
	}
	cp := reflect.MakeMapWithSize(v.Type(), v.Len())
	for _, key := range v.MapKeys() {
		valCopy, err := deepCopy(v.MapIndex(key), t)
		if err != nil {
			return values.Empty, err
		}
		keyCopy, err := deepCopy(key, t)
		if err != nil {
			return values.Empty, err
		}
		cp.SetMapIndex(keyCopy, valCopy)
	}
	// cache result to shortcut the copy
	return t.Mark(v, cp)
}

func DeepCopyPointer(v reflect.Value) (reflect.Value, error) {
	if v.Kind() != reflect.Pointer {
		return values.Empty, fmt.Errorf("deep copy pointer called with a non pointer kind: '%v'", v.Kind())
	}
	return deepCopyPointer(v, tracker.Reference())
}

// deepCopyPointer handles pointers by cloning the referenced value and avoiding cyclic references.
func deepCopyPointer(v reflect.Value, t *tracker.RefTracker) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	// handle pointer cache to avoid cyclic references
	cache, exists, err := t.Get(v)
	str := v.String()
	fmx.Sprintf("v: %s", str)
	if err != nil {
		return values.Empty, err
	}
	if exists {
		sc := cache.String()
		fmx.Sprintf("cache: %s", sc)
		return cache, nil
	}
	cp, err := deepCopy(v.Elem(), t)
	if err != nil {
		return values.Empty, err
	}
	// cache result to shortcut the copy
	return t.Mark(v, pointers.NewOf(cp))
}

// DeepCopySlice copies a slice deeply
func DeepCopySlice(v reflect.Value) (reflect.Value, error) {
	if v.Kind() != reflect.Slice {
		return values.Empty, fmt.Errorf("deep slice copy called with a non slice value: '%v'", v.Kind())
	}
	return deepCopySlice(v, tracker.Reference())
}

// deepCopySlice is a helper function for slice copying.
func deepCopySlice(v reflect.Value, t *tracker.RefTracker) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	// Slices are natural pointers
	cache, exists, err := t.Get(v)
	if err != nil {
		return values.Empty, err
	}
	if exists {
		return cache, nil
	}
	cp := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
	for i := 0; i < v.Len(); i++ {
		elem, err := deepCopy(v.Index(i), t)
		if err != nil {
			return reflect.Value{}, err
		}
		cp.Index(i).Set(elem)
	}
	return t.Mark(v, cp)
}

func DeepCopyStruct(v reflect.Value, t *tracker.RefTracker) (reflect.Value, error) {
	if v.Kind() != reflect.Struct {
		return values.Empty, fmt.Errorf("deep copy struct called on non struct kind: '%v'", v.Kind())
	}
	return deepCopyStruct(v, t)
}

// deepCopyStruct deep copies each exported field of a struct.
func deepCopyStruct(v reflect.Value, t *tracker.RefTracker) (reflect.Value, error) {
	fmx.Greenf("\t[copy-Struct]: %s\n", v.String())
	cp := reflect.New(v.Type()).Elem()

	if !v.CanAddr() {
		v = values.OfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ptr := unsafe.Pointer(field.UnsafeAddr())

		src := reflect.NewAt(field.Type(), ptr).Elem()
		str := field.String()
		fmx.Sprintf(str)
		fieldCopy, err := deepCopy(src, t)
		if err != nil {
			return values.Empty, err
		}

		if field.Type().Kind() != fieldCopy.Kind() {
			panic(fmx.Sprintf("error deep copying '%s'. fieldCopy '%s' mismatches source "+
				"field '%s'", cp.String(), fieldCopy.String(), field.String()))
		}

		dest := cp.Field(i)

		if fieldCopy.Kind() != dest.Kind() {
			panic(fmx.Sprintf("error deep copying '%s'. fieldCopy '%s' mismatches destination "+
				"field '%s'", cp.String(), fieldCopy.String(), dest.String()))
		}

		reflect.NewAt(dest.Type(), unsafe.Pointer(dest.UnsafeAddr())).
			Elem().Set(fieldCopy)
	}

	return cp, nil
}
