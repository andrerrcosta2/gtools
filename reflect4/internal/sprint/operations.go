// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"reflect"
	"sort"
	"strings"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
)

// defaultArray extracts the sprint from an array
func defaultArray(tab indent.Indentor, v reflect.Value, s *Strategy) string {
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
func defaultChan(tab indent.Indentor, v reflect.Value) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	return tab.Sprint(name)
}

func defaultFields(tab indent.Indentor, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	sb := strings.Builder{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		name := typ.Field(i).Name

		if field.IsZero() && values.CanNil(field) {
			sb.WriteString(sprints.Lfield(tab, name, sprints.NilType(indent.Zero(),
				types.ValidValueName(field.Type()))) + ",")
			continue
		}

		//if !field.CanInterface() {
		//	// since only fields can be unexported
		//	// here is the only method we need to use unsafe operations
		//	if !field.IsZero() {
		//		field = values.UnsafeForceOfUnaddr(field)
		//	}
		//}
		sb.WriteString(sprints.Lfield(tab, name, sprintOf(tab, field, s)) + ",")
	}
	return sb.String()
}

func exportedFields(tab indent.Indentor, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	var sb strings.Builder

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		name := typ.Field(i).Name

		if field.IsZero() && values.CanNil(field) {
			sb.WriteString(sprints.Lfield(tab, name, sprints.NilType(indent.Zero(), name)) + ",")
			continue
		}

		if !typ.Field(i).IsExported() {
			continue
		}
		val := sprintOf(tab, v.Field(i), s)
		sb.WriteString(sprints.Lfield(tab, name, val) + ",")
	}

	return sb.String()
}

// defaultFunc extracts the sprint from a function
func defaultFunc(tab indent.Indentor, v reflect.Value) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	return tab.Sprint(name)
}

func defaultInterface(tab indent.Indentor, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.IsNil() {
		return sprints.NilType(tab, name)
	}
	vv := sprintOf(tab.Inc(), v.Elem(), s)
	return sprints.Interface(tab, name, vv)
}

// defaultMap extracts the sprint from a map
func defaultMap(tab indent.Indentor, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	name := types.ValidValueName(typ)

	if v.IsNil() || v.IsZero() {
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
	keys := make([]reflect.Value, 0, v.Len())

	// order standard
	for iter.Next() {
		keys = append(keys, iter.Key())
	}

	sort.Slice(keys, func(i, j int) bool {
		ki := sprintOf(tab.Inc(), keys[i], s)
		kj := sprintOf(tab.Inc(), keys[j], s)
		return ki < kj
	})

	for _, k := range keys {
		next := tab.Inc()
		key := next.Trim(sprintOf(next, k, s))
		val := next.Trim(sprintOf(next, v.MapIndex(k), s))

		sb.WriteString("\n" + sprints.MapEntry(next, key, val))
	}
	sb.WriteString("\n" + tab.Sprint("}"))
	return sb.String()
}

func defaultPointer(tab indent.Indentor, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	if v.IsNil() {
		return sprints.NilType(tab, "*"+types.ValidValueName(typ.Elem()))
	}
	spr, ok := s.Check(v)
	if ok {
		return spr
	}
	addr := v.Pointer()
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
	s.Mark(v, sprints.CyclicRef(tab, "*"+types.ValidValueName(typ.Elem()), sprints.Uintptrf(addr)))
	val := tab.Trim(sprintOf(tab, v.Elem(), s))
	return sprints.PointerElem(tab, val)
}

// sprintPrimitive extracts a sprint from a primitive
func sprintPrimitive(tab indent.Indentor, v reflect.Value) string {
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
func defaultSlice(tab indent.Indentor, v reflect.Value, s *Strategy) string {
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
func defaultStruct(tab indent.Indentor, v reflect.Value, s *Strategy) string {
	typ := v.Type()
	name := types.ValidValueName(typ)
	if v.NumField() == 0 {
		return sprints.BClosedobj(tab, name)
	}
	sb := strings.Builder{}
	sb.WriteString(tab.Sprint(fmx.SBold(name) + "{"))

	//if !v.CanAddr() {
	//	v = values.ForceOfUnaddr(v)
	//}
	sb.WriteString(s.Fields(tab.Inc(), v, s))
	sb.WriteString("\n" + tab.Sprint("}"))
	return sb.String()
}

func defaultUnsafe(tab indent.Indentor, v reflect.Value) string {
	if v.IsNil() {
		return sprints.NilType(tab, "unsafe.pointer")
	}
	ptrVal := v.Interface().(unsafe.Pointer)
	return sprints.UnsafePointer(tab, ptrVal)
}

func skipDeep(tab indent.Indentor, v reflect.Value, s *Strategy) string {
	name := types.ValidValueName(v.Type())
	return sprints.Errorf(tab, "<%v>(skipped)", name)
}

func skipShallow(tab indent.Indentor, v reflect.Value) string {
	name := types.ValidValueName(v.Type())
	return sprints.Errorf(tab, "<%v>(skipped)", name)
}
