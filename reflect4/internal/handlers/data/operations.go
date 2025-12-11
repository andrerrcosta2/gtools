// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
)

func deepCopy(v reflect.Value, s *CopyStrategy) (reflect.Value, error) {
	if v.Kind() == reflect.Invalid {
		return values.Empty, reflect4.ErrInvalidValue
	}

	switch v.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		return v, nil
	case reflect.Array:
		return deepCopyArray(v, s)
	case reflect.Slice:
		return deepCopySlice(v, s)
	case reflect.Chan:
		return s.CopyChan(v, nil)
	case reflect.Map:
		return deepCopyMap(v, s)
	case reflect.Struct:
		return deepCopyStruct(v, s)
	case reflect.Ptr:
		return s.CopyPtr(v, s)
	case reflect.Interface:
		return deepCopyInterface(v, s)
	case reflect.Func:
		return s.CopyFunc(v, nil)
	case reflect.UnsafePointer:
		return v, nil
	default:
		return values.Empty, fmt.Errorf("unsupported kind: %v", v.Kind())
	}
}

// deepCopyArray copies an array deeply based on the CopyStrategy
func deepCopyArray(v reflect.Value, s *CopyStrategy) (reflect.Value, error) {
	// array sizes are part of its type definitions in golang,
	cp := reflect.New(v.Type()).Elem()
	for i := 0; i < v.Len(); i++ {
		elem, err := deepCopy(v.Index(i), s)
		if err != nil {
			return values.Empty, err
		}
		cp.Index(i).Set(elem)
	}
	return cp, nil
}

// deepCopyInterface copies an interface deeply based on the CopyStrategy
func deepCopyInterface(v reflect.Value, s *CopyStrategy) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	elem := v.Elem()
	cp, err := deepCopy(elem, s)
	if err != nil {
		return values.Empty, err
	}
	return reflect.ValueOf(cp.Interface()).Convert(v.Type()), nil
}

func deepCopyMap(v reflect.Value, s *CopyStrategy) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	cache, exists := s.Check(v)
	if exists {
		return cache, nil
	}
	cp := reflect.MakeMapWithSize(v.Type(), v.Len())
	for _, key := range v.MapKeys() {
		valCopy, err := deepCopy(v.MapIndex(key), s)
		if err != nil {
			return values.Empty, err
		}
		keyCopy, err := deepCopy(key, s)
		if err != nil {
			return values.Empty, err
		}
		cp.SetMapIndex(keyCopy, valCopy)
	}
	// cache result to shortcut the clone
	return s.Cache(v, cp), nil
}

// deepCopySlice is a std function for slice copying.
func deepCopySlice(v reflect.Value, s *CopyStrategy) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	cp := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
	for i := 0; i < v.Len(); i++ {
		elem, err := deepCopy(v.Index(i), s)
		if err != nil {
			return reflect.Value{}, err
		}
		cp.Index(i).Set(elem)
	}
	return cp, nil
}

// deepCopyStruct deep copies each exported field of a struct.
func deepCopyStruct(v reflect.Value, s *CopyStrategy) (reflect.Value, error) {
	if !v.CanAddr() {
		v = values.ForceOfUnaddr(v)
	}
	cp := reflect.New(v.Type()).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ptr := unsafe.Pointer(field.UnsafeAddr())
		src := reflect.NewAt(field.Type(), ptr).Elem()
		fieldCopy, err := deepCopy(src, s)
		if err != nil {
			return values.Empty, err
		}

		dest := cp.Field(i)
		reflect.NewAt(dest.Type(), unsafe.Pointer(dest.UnsafeAddr())).
			Elem().Set(fieldCopy)
	}

	return cp, nil
}

// defaultDeepCopyPointer deep copies a pointer element
func defaultDeepCopyPointer(v reflect.Value, s *CopyStrategy) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	// handle pointer cache to avoid cyclic references
	cache, exists := s.Check(v)
	if exists {
		//sc := cache.strings()
		//fmx.Purplef("\t\t[cached:already-exists]: %s\n", sc)
		return cache, nil
	}

	// This is a kind of deadlock-like recursion prevention.
	// We basically create a placeholder as the pointer itself,
	// then we deliver the same object at once
	//
	// 1. Allocate an empty target
	elemType := v.Type().Elem()
	placeholder := reflect.New(elemType)

	// 2. mark the pointer preemptively
	s.Cache(v, placeholder)

	// 3. clone the underlying value into placeholder
	copied, err := deepCopy(v.Elem(), s)
	if err != nil {
		return values.Empty, err
	}
	err = values.UnsafeSet(placeholder.Elem(), copied)
	if err != nil {
		return values.Empty, err
	}

	// 4. Return the pointer to the placeholder
	return placeholder, nil
}

func defaultDeepCopyChan(v reflect.Value, _ *CopyStrategy) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}

	// As channels are not addressable, we can't use reflect's "New" to create a new one.
	// Instead, we create a new one using the "make" function.
	elemType := v.Type().Elem()
	capacity := v.Cap()
	// Always make a bidirectional version first
	bidirChanType := reflect.ChanOf(reflect.BothDir, elemType)

	newChan := reflect.MakeChan(bidirChanType, capacity)

	// Convert back to the same direction as the original
	newChan = newChan.Convert(v.Type())

	return newChan, nil
}

func defaultDeepCopyFunc(v reflect.Value, _ *CopyStrategy) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	// I suppose that we can't actually clone functions in golang
	// if we have closures, the locked value might be modifiable be both.
	return reflect.New(v.Type()).Elem(), nil
}

func identityCopy(v reflect.Value, _ *CopyStrategy) (reflect.Value, error) {
	return v, nil
}

func readAllFields(v reflect.Value, i int) (field reflect.Value, canRead bool) {
	return v.Field(i), true
}

func readAddrStruct(v reflect.Value) (s reflect.Value, canRead bool) {
	return reflect.New(v.Type()).Elem(), s.CanAddr()
}

func readAnyStruct(v reflect.Value) (s reflect.Value, canRead bool) {
	return reflect.New(v.Type()).Elem(), true
}

func readSettableField(v reflect.Value, i int) (field reflect.Value, canRead bool) {
	f := v.Field(i)
	return f, f.CanSet()
}

func rideSettableFields(a, b reflect.Value, fn functions.TriPredicate[int, reflect.Value, reflect.Value]) {
	for i := 0; i < a.NumField(); i++ {
		f1 := a.Field(i)
		if !f1.CanSet() {
			continue
		}
		f2 := b.Field(i)
		if !fn(i, f1, f2) {
			return
		}
	}
}

func rideAllFields(a, b reflect.Value, fn functions.TriPredicate[int, reflect.Value, reflect.Value]) {
	for i := 0; i < a.NumField(); i++ {
		f1 := a.Field(i)
		f2 := b.Field(i)
		if !fn(i, f1, f2) {
			return
		}
	}
}

func shallowCopyFunc(v reflect.Value) (reflect.Value, error) {
	if v.IsNil() {
		return reflect.Zero(v.Type()), nil
	}
	// I suppose that we can't actually clone functions in golang
	// if we have closures the locked value might be modifiable be both.
	return reflect.New(v.Type()).Elem(), nil
}

func skipCopy(v reflect.Value, _ *CopyStrategy) (reflect.Value, error) {
	return reflect.Zero(v.Type()), nil
}
