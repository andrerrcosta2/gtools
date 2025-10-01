// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ask

import (
	"bytes"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/testlite/internal/meml"
	"reflect"
	"unsafe"
)

func AreEquals(a, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	if ba, ok := a.([]byte); ok {
		if bb, ok := b.([]byte); ok {
			return bytes.Equal(ba, bb)
		}
		return false
	}
	return reflect.DeepEqual(a, b)
}

func AreEqualMaps[M ~map[K]V, K comparable, V any](a, b M) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		vv, ok := b[k]
		if !ok || !AreEquals(v, vv) {
			return false
		}
	}
	return true
}

// AreSameInstances returns true if both arguments are ptrs to the same instance.
// For values, it'll always return false
func AreSameInstances(a, b any) (bool, [2]uintptr, error) {
	ta := reflect.TypeOf(a)
	if !IsReferenceType(ta) {
		return false, [2]uintptr{}, fmt.Errorf("non-reference type cannot be on two different variables")
	}
	tb := reflect.TypeOf(b)
	if !IsReferenceType(tb) {
		return false, [2]uintptr{}, fmt.Errorf("non-reference type cannot be on two different variables")
	}
	if ta.Kind() != tb.Kind() {
		return false, [2]uintptr{}, fmt.Errorf("different kinds: '(%v) != (%v)'", ta.Kind(), tb.Kind())
	}
	//fmt.Printf("Checking if '%p == %p'\n", a, b)
	ua := uintptr(meml.UnsafePointerOf(a))
	ub := uintptr(meml.UnsafePointerOf(b))
	//fmt.Printf("ua: '%d', ub: '%d'\n", ua, ub)
	return ua == ub, [2]uintptr{ua, ub}, nil
}

// HasAnyNilField reports whether the struct has any nil field
func HasAnyNilField(v reflect.Value) (bool, []string, error) {
	v = unwrapValue(v)
	if v.Kind() != reflect.Struct {
		return false, nil, fmt.Errorf("expected a struct, got %s", v.String())
	}
	nf := nilFields(v)
	return len(nf) > 0, nf, nil
}

func HasAnyNonInterfaceNilField(v reflect.Value) (bool, []string, error) {
	v = unwrapValue(v)
	if v.Kind() != reflect.Struct {
		return false, nil, fmt.Errorf("expected a struct, got %s", v.Kind())
	}
	nf := nonInterfaceNilFields(v)
	return len(nf) > 0, nf, nil
}

// HasAnyNonNilField reports whether the struct has any non nil field
func HasAnyNonNilField(v reflect.Value) (bool, []string, error) {
	v = unwrapValue(v)
	if v.Kind() != reflect.Struct {
		return false, nil, fmt.Errorf("expected a struct, got %s", v.String())
	}
	nf := nonNilFields(v)
	return len(nf) > 0, nf, nil
}

func HasExportedNilField(v reflect.Value) (bool, []string, error) {
	v = unwrapValue(v)
	if v.Kind() != reflect.Struct {
		return false, nil, fmt.Errorf("expected a struct, got %s", v.Kind())
	}
	nf := nilExportedFields(v)
	return len(nf) > 0, nf, nil
}

// HasExportedNonInterfaceNilField reports whether a struct has exported non-interface nil fields
func HasExportedNonInterfaceNilField(v reflect.Value) (bool, []string, error) {
	v = unwrapValue(v)
	if v.Kind() != reflect.Struct {
		return false, nil, fmt.Errorf("expected a struct, got %s", v.Kind())
	}
	nf := nilExportedNonInterfaceFields(v)
	return len(nf) > 0, nf, nil
}

func IsNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()

	default:
		return false
	}
}

// IsNonEmptyInterface reports whether a type is not a interface{}
func IsNonEmptyInterface(typ reflect.Type) bool {
	return typ.Kind() == reflect.Interface && typ.NumMethod() > 0
}

// IsReferenceType reports whether a type can be referenced
func IsReferenceType(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Ptr, reflect.Chan, reflect.Map, reflect.Func, reflect.UnsafePointer, reflect.Interface:
		return true
	}
	return false
}

// nilFields returns all nil fields from a struct
func nilFields(v reflect.Value) []string {
	if !v.CanAddr() {
		ptr := reflect.New(v.Type())
		ptr.Elem().Set(v)
		v = ptr.Elem()
	}
	var vt = v.Type()
	var nf = make([]string, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ftype := vt.Field(i)
		if !ftype.IsExported() {
			ptr := unsafe.Pointer(field.UnsafeAddr())
			field = reflect.NewAt(field.Type(), ptr).Elem()
		}
		if IsNil(v.Field(i)) {
			nf = append(nf, ftype.Name)
		}
	}
	return nf
}

// nonNilFields returns all non-nil fields from a struct
func nonNilFields(v reflect.Value) []string {
	if !v.CanAddr() {
		ptr := reflect.New(v.Type())
		ptr.Elem().Set(v)
		v = ptr.Elem()
	}
	var vt = v.Type()
	var nf = make([]string, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ftype := vt.Field(i)
		if !ftype.IsExported() {
			ptr := unsafe.Pointer(field.UnsafeAddr())
			field = reflect.NewAt(field.Type(), ptr).Elem()
		}
		if !IsNil(v.Field(i)) {
			nf = append(nf, ftype.Name)
		}
	}
	return nf
}

// nilExportedFields returns the exported nil fields from a struct
func nilExportedFields(v reflect.Value) []string {
	var vt = v.Type()
	var nf = make([]string, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		ft := vt.Field(i)
		if !ft.IsExported() {
			continue
		}
		if IsNil(v.Field(i)) {
			nf = append(nf, ft.Name)
		}
	}
	return nf
}

// nonInterfaceNilFields returns all nil fields excluding interfaces
// with at least one method from a struct
func nonInterfaceNilFields(v reflect.Value) []string {
	if !v.CanAddr() {
		ptr := reflect.New(v.Type())
		ptr.Elem().Set(v)
		v = ptr.Elem()
	}
	var vt = v.Type()
	var nf = make([]string, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ftype := vt.Field(i)
		typ := field.Type()

		if !ftype.IsExported() {
			ptr := unsafe.Pointer(field.UnsafeAddr())
			field = reflect.NewAt(field.Type(), ptr).Elem()
		}

		// Skip all interfaces that aren't interface{}
		if IsNonEmptyInterface(typ) {
			continue
		}

		if IsNil(field) {
			nf = append(nf, ftype.Name)
		}
	}
	return nf
}

// nilExportedNonInterfaceFields returns the exported nil fields excluding interfaces
// with at least one method from a struct
func nilExportedNonInterfaceFields(v reflect.Value) []string {
	var vt = v.Type()
	var nf = make([]string, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ftype := vt.Field(i)
		typ := field.Type()

		// Skip nil interfaces that aren't interface{}
		if IsNonEmptyInterface(typ) || !ftype.IsExported() {
			continue
		}

		if IsNil(field) {
			nf = append(nf, ftype.Name)
		}
	}
	return nf
}

func unwrapValue(t reflect.Value) reflect.Value {
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		t = t.Elem()
	}
	return t
}
