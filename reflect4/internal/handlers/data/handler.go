// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"fmt"
	"reflect"
)

// DeepCopy deep copies the given value.
// The function returns a deep copy of the value and an error.
// The error is nil if the copy was successful.
func DeepCopy[T any](value T) (cp T, err error) {
	// Create a pointer map to avoid cyclic references
	ptr := make(map[uintptr]interface{})
	var res any
	var ok bool
	// Call the deepCopy function with the pointer map and the value
	res, err = deepCopy(value, ptr)
	if err != nil {
		return
	}
	// Cast the deep copied value to the original type
	if cp, ok = res.(T); !ok {
		return cp, fmt.Errorf("failed to cast deep copied value to original type")
	}
	return
}

func ShallowCopy[T any](value T) (T, error) {
	return value, nil
}

func deepCopy(data any, ptr map[uintptr]interface{}) (any, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Invalid {
		return data, nil
	}

	switch val.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		return data, nil
	case reflect.Array, reflect.Slice:
		return deepCopyIterable(data, ptr, val.Kind())
	case reflect.Map:
		return copyMap(data, ptr)
	case reflect.Struct:
		return deepCopyStruct(data, ptr)
	case reflect.Ptr:
		return copyPointer(data, ptr)
	default:
		return nil, fmt.Errorf("unsupported kind: %v", val.Kind())
	}
}

func DeepCopyArray[T any](data []T) ([]T, error) {
	cp, err := deepCopyIterable(data, make(map[uintptr]any), reflect.Array)
	return cp.([]T), err
}

func DeepCopySlice[T any](data []T) ([]T, error) {
	cp, err := deepCopyIterable(data, make(map[uintptr]any), reflect.Slice)
	return cp.([]T), err
}

// deepCopyIterable is a helper function for array and slice copying.
func deepCopyIterable(data any, ptr map[uintptr]any, kind reflect.Kind) (any, error) {
	val := reflect.ValueOf(data)
	if val.Kind() != kind {
		return nil, fmt.Errorf("must pass a value of kind %v; got %v", kind, val.Kind())
	}

	// Create a new array or slice to hold the copied elements
	cp := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
	for i := 0; i < val.Len(); i++ {
		elemCopy, err := deepCopy(val.Index(i).Interface(), ptr)
		if err != nil {
			return nil, err
		}
		cp.Index(i).Set(reflect.ValueOf(elemCopy))
	}
	return cp.Interface(), nil
}

func CopyMap[K comparable, V any](data map[K]V) (map[K]V, error) {
	cp, err := copyMap(data, make(map[uintptr]any))
	return cp.(map[K]V), err
}

// copyMap deep copies a map, including keys and values.
func copyMap(x interface{}, ptr map[uintptr]interface{}) (interface{}, error) {
	val := reflect.ValueOf(x)
	cp := reflect.MakeMapWithSize(val.Type(), val.Len())

	for _, key := range val.MapKeys() {
		valCopy, err := deepCopy(val.MapIndex(key).Interface(), ptr)
		if err != nil {
			return nil, err
		}
		keyCopy, err := deepCopy(key.Interface(), ptr)
		if err != nil {
			return nil, err
		}
		cp.SetMapIndex(reflect.ValueOf(keyCopy), reflect.ValueOf(valCopy))
	}
	return cp.Interface(), nil
}

func CopyPointer[T any](data *T) (any, error) {
	ptr := make(map[uintptr]any)
	val, err := copyPointer(data, ptr)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// copyPointer handles pointers by cloning the referenced value and avoiding cyclic references.
func copyPointer(data any, ptr map[uintptr]any) (any, error) {
	val := reflect.ValueOf(data)
	ptrAddr := val.Pointer()

	// If the pointer was already cloned, return the reference to avoid cyclic copying
	if existing, ok := ptr[ptrAddr]; ok {
		return existing, nil
	}

	cp := reflect.New(val.Elem().Type())
	ptr[ptrAddr] = cp.Interface()
	elemCopy, err := deepCopy(val.Elem().Interface(), ptr)
	if err != nil {
		return nil, err
	}
	cp.Elem().Set(reflect.ValueOf(elemCopy))
	return cp.Interface(), nil
}

func DeepCopyStruct[T any](data T) (T, error) {
	cp, err := deepCopyStruct(data, make(map[uintptr]any))
	return cp.(T), err
}

// deepCopyStruct deep copies each exported field of a struct.
func deepCopyStruct(data interface{}, ptr map[uintptr]interface{}) (interface{}, error) {
	val := reflect.ValueOf(data)
	cp := reflect.New(val.Type()).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.CanInterface() { // Check for exported fields
			fieldCopy, err := deepCopy(field.Interface(), ptr)
			if err != nil {
				return nil, fmt.Errorf("error copying struct field: %v", err)
			}
			cp.Field(i).Set(reflect.ValueOf(fieldCopy))
		}
	}
	return cp.Interface(), nil
}
