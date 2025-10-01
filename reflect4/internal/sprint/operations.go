// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
	"strings"
	"unsafe"
)

// defaultArray extracts the sprint from an array
func defaultArray(tab indent.Tab, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	sb := strings.Builder{}
	if v.Len() == 0 {
		return sprints.EmptyIterable(tab, name)
	}
	sb.WriteString(tab.Sprint(name) + "[")

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		vv := sprintOf(tab.Inc(), elem, s)
		sb.WriteString("\n" + vv + ",")
	}

	sb.WriteString("\n" + tab.Sprint("]"))
	return sb.String()
}

// defaultChan extracts the sprint from a channel
func defaultChan(tab indent.Tab, v reflect.Value) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	return tab.Sprint(name)
}

// defaultField extracts the sprint from a field
func defaultField(tab indent.Tab, value reflect.Value, s *Strategy) string {
	// CanInterface() means:
	// - Calling .interfaces() is safe (won't panic).
	// - The value is exported.
	// - If value is a pointer, calling .interfaces() gives you that pointer wrapped inside interface{}.
	// - If value is a non-pointer value, .interfaces() gives you a clone of that value inside interface{}.
	if value.CanInterface() || value.CanAddr() {
		return sprintOf(tab, value, s)
	}

	// If neither CanInterface() nor CanAddr() is true, the value:
	// - Is unexported (hence unsafe to use .interfaces()).
	// - Is not directly addressable.
	// - Is not directly addressable.
	// - We must use unsafe.pointers and reflect.NewAt() to access it.
	return sprintOf(tab.Inc(), values.UnsafeOfUnaddr(value), s)
}

// defaultFunc extracts the sprint from a function
func defaultFunc(tab indent.Tab, v reflect.Value) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	return tab.Sprint(name)
}

func defaultInterface(tab indent.Tab, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	vv := sprintOf(tab.Inc(), v.Elem(), s)
	return sprints.Interface(tab, name, vv)
}

// defaultMap extracts the sprint from a map
func defaultMap(tab indent.Tab, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	spr, ok := s.Check(v)
	if ok {
		return spr
	}
	addr := v.UnsafePointer()
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return.
	// Because it isn't a cycle just because it'll be printed again later.
	s.Mark(v, sprints.CyclicRef(tab, name, sprints.Uintptrf(uintptr(addr))))

	sb := strings.Builder{}
	// Empty maps can hold nil elements, we need to extract the kind using its
	// type declaration
	if v.Len() == 0 {
		return sprints.Closedobj(tab, name)
	}

	iter := v.MapRange()
	sb.WriteString(tab.Sprint(name) + "{")
	for iter.Next() {
		// Get the v for each sorted key
		val := sprintOf(indent.Zero(), iter.Value(), s)
		key := sprintOf(indent.Zero(), iter.Key(), s)
		sb.WriteString("\n" + sprints.MapEntry(tab.Inc(), key, val))
	}
	sb.WriteString("\n" + tab.Sprint("}"))
	return sb.String()
}

func defaultPointer(tab indent.Tab, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	spr, ok := s.Check(v)
	if ok {
		return spr
	}
	addr := v.Pointer()
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
	s.Mark(v, sprints.CyclicRef(tab, name, sprints.Uintptrf(addr)))
	val := sprintOf(tab, v.Elem(), s)
	return sprints.PointerElem(tab, val)
}

// sprintPrimitive extracts a sprint from a primitive
func sprintPrimitive(tab indent.Tab, v reflect.Value) string {
	name := types.ValidValueName(v.Type())
	switch v.Kind() {
	case reflect.Bool:
		return tab.Sprintf("<%v>%t", name, v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return tab.Sprintf("<%v>%d", name, v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return tab.Sprintf("<%v>%d", name, v.Uint())
	case reflect.Float32, reflect.Float64:
		return tab.Sprintf("<%v>%g", name, v.Float())
	case reflect.Complex64, reflect.Complex128:
		return tab.Sprintf("<%v>%g", name, v.Complex())
	case reflect.String:
		return tab.Sprintf("<%v>%q", name, v.String())
	default:
		return sprints.Errorf(tab, "invalid primitive type: %v", v.Kind().String())
	}
}

// defaultSlice extracts the sprint from a slice
func defaultSlice(tab indent.Tab, v reflect.Value, s *Strategy) string {
	name := types.ValidValueName(v.Type())
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	spr, ok := s.Check(v)
	if ok {
		return spr
	}
	s.Mark(v, sprints.CyclicRef(tab, name, sprints.Uintptrf(uintptr(v.UnsafePointer()))))
	sb := strings.Builder{}
	if v.Len() == 0 {
		return sprints.Closedobj(tab, name)
	}
	sb.WriteString(tab.Sprint(name + "["))

	for i := 0; i < v.Len(); i++ {
		elem := sprintOf(tab.Inc(), v.Index(i), s)
		sb.WriteString("\n" + elem + ",")
	}

	sb.WriteString("\n" + tab.Sprint("]"))
	return sb.String()
}

// defaultStruct extracts the sprint from a struct
func defaultStruct(tab indent.Tab, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.NumField() == 0 {
		return sprints.BClosedobj(tab, name)
	}
	sb := strings.Builder{}
	sb.WriteString(tab.Sprint(fmx.SBold(name) + "{"))

	if !v.CanAddr() {
		v = values.UnsafeOfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		val := defaultField(tab.Inc(), v.Field(i), s)
		fieldName := typ.Field(i).Name
		sb.WriteString(sprints.Ltfield(tab, fieldName, val) + ",")
	}

	sb.WriteString("\n" + tab.Sprint("}"))
	return sb.String()
}

func defaultUnsafe(tab indent.Tab, v reflect.Value) string {
	if v.IsNil() {
		return sprints.NilType(tab, "unsafe.pointers")
	}
	ptrVal := v.Interface().(unsafe.Pointer)
	return sprints.UnsafePointer(tab, ptrVal)
}
