// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/domain/gerrors"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/pointers"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/arrays"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/chans"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/funcs"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/interfaces"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/maps"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/primitives"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/slices"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/structs"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/unsafes"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
	"sort"
	"strings"
)

func Of(tab indent.Tab, value reflect.Value) (string, error) {
	return sprintOf(tab, value, tracker.Sprint())
}

func sprintOf(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if !value.IsValid() {
		return "invalid type: <invalid>", gerrors.NewError(gerrors.Input_err, reflect4.ErrInvalidValue)
	}

	if value.Kind() <= reflect.Complex128 || value.Kind() == reflect.String {
		return sprintPrimitive(tab, value), nil
	}

	switch value.Kind() {
	case reflect.Array:
		return sprintArray(tab, value, t)
	case reflect.Chan:
		return sprintChan(tab, value), nil
	case reflect.Func:
		return sprintFunc(tab, value), nil
	case reflect.Interface:
		return sprintInterface(tab, value, t)
	case reflect.Map:
		return sprintMap(tab, value, t)
	case reflect.Ptr:
		return sprintPointer(tab, value, t)
	case reflect.Slice:
		return sprintSlice(tab, value, t)
	case reflect.Struct:
		return sprintStruct(tab, value, t)
	case reflect.UnsafePointer:
		return sprintUnsafePointer(tab, value, t)
	default:
		return sprints.Errorf(tab, "invalid type: %s", value.Kind().String()), nil
	}
}

// sprintArray extracts the sprint from an array
func sprintArray(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	// Arrays aren't underlying pointers. So only its values can present cyclic references.
	name, _ := arrays.Name(value.Type())

	sb := strings.Builder{}
	if value.Len() == 0 {
		return tab.Sprint(name + "[<empty>]"), nil
	}
	sb.WriteString(tab.Sprint(name) + "[")

	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		v, err := sprintOf(tab.Inc(), elem, t)
		if err != nil {
			return "", err
		}
		sb.WriteString("\n" + v + ",")
	}

	sb.WriteString("\n" + tab.Sprint("]"))
	return sb.String(), nil
}

// sprintChan extracts the sprint from a channel
func sprintChan(tab indent.Tab, value reflect.Value) string {
	name, _ := chans.Name(value.Type())
	if value.IsNil() {
		return sprints.NilType(tab, name)
	}
	return tab.Sprint(name)
}

// sprintField extracts the sprint from a field
func sprintField(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	// CanInterface() means:
	// - Calling .Interface() is safe (won't panic).
	// - The value is exported.
	// - If value is a pointer, calling .Interface() gives you that pointer wrapped inside interface{}.
	// - If value is a non-pointer value, .Interface() gives you a copy of that value inside interface{}.
	if value.CanInterface() || value.CanAddr() {
		return sprintOf(tab, value, t)
	}

	// If neither CanInterface() nor CanAddr() is true, the value:
	// - Is unexported (hence unsafe to use .Interface()).
	// - Is not directly addressable.
	// - We must use unsafe.Pointer and reflect.NewAt() to access it.
	return sprintOf(tab.Inc(), values.OfUnaddr(value).Elem(), t)
}

func sprintInterface(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	// interfaces aren't underlying pointers
	name, _ := interfaces.Name(value.Type())
	// so we forward the tracker to the recursive function
	val, err := sprintOf(tab.Inc(), value.Elem(), t)
	if err != nil {
		return "", err
	}
	return tab.Sprintf("%s{\n%s\n%s}", name, val, tab.String()), nil
}

// sprintFunc extracts the sprint from a function
func sprintFunc(tab indent.Tab, value reflect.Value) string {
	name, _ := funcs.Name(value.Type())
	return tab.Sprint(name)
}

// sprintMap extracts the sprint from a map
func sprintMap(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, value, t); hasCache {
		return cache, err
	}

	addr, err := pointers.AddrOf(value)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return.
	// Because it isn't a cycle just because it'll be printed again later.
	name, _ := maps.Name(value.Type())
	err = t.Mark(value, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}

	sb := strings.Builder{}
	if value.IsNil() {
		res := sprints.NilType(tab, name)
		err = t.Mark(value, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}
	// Empty maps can hold nil elements, we need to extract the kind using its
	// type declaration
	if value.Len() == 0 {
		res := sprints.Closedobj(tab, name)
		err = t.Mark(value, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}

	// Sort the map keys alphabetically
	var keys []reflect.Value
	for _, key := range value.MapKeys() {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].String() < keys[j].String()
	})

	// Start building the map representation
	sb.WriteString(tab.Sprint(name) + "{")
	for _, key := range keys {
		// Get the value for each sorted key
		val := value.MapIndex(key)
		k, err := sprintOf(tab.Inc(), key, t)
		if err != nil {
			return "", err
		}
		sb.WriteString("\n" + k)
		sb.WriteString(": ")
		v, err := sprintOf(tab.Inc(), val, t)
		if err != nil {
			return "", err
		}
		sb.WriteString(sprints.FieldVal(v) + ",")
	}

	sb.WriteString("\n" + tab.Sprint("}"))
	res := sb.String()
	err = t.Mark(value, res)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	return res, nil
}

func sprintPointer(tab indent.Tab, v reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, v, t); hasCache {
		return cache, err
	}

	addr := v.Pointer()
	name, _ := pointers.Name(v.Type())
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
	err := t.Mark(v, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	if v.IsNil() {
		res := sprints.NilType(tab, name)
		err = t.Mark(v, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}
	val, err := sprintOf(tab, v.Elem(), t)
	return tab.Sprintf("*%s", strings.TrimSpace(val)), err
}

// sprintPrimitive extracts a sprint from a primitive
func sprintPrimitive(tab indent.Tab, s reflect.Value) string {
	name, _ := primitives.Name(s.Type())
	switch s.Kind() {
	case reflect.Bool:
		return tab.Sprintf("<%s>%t", name, s.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return tab.Sprintf("<%s>%d", name, s.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return tab.Sprintf("<%s>%d", name, s.Uint())
	case reflect.Float32, reflect.Float64:
		return tab.Sprintf("<%s>%g", name, s.Float())
	case reflect.Complex64, reflect.Complex128:
		return tab.Sprintf("<%s>%g", name, s.Complex())
	case reflect.String:
		return tab.Sprintf("<%s>%q", name, s.String())
	default:
		return sprints.Errorf(tab, "invalid primitive type: %s", s.Kind().String())
	}
}

// sprintSlice extracts the sprint from a slice
func sprintSlice(tab indent.Tab, v reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, v, t); hasCache {
		return cache, err
	}

	name, _ := slices.Name(v.Type())
	addr := v.Pointer()
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
	err := t.Mark(v, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	sb := strings.Builder{}
	if v.IsNil() {
		res := sprints.NilType(tab, name)
		err = t.Mark(v, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}
	if v.Len() == 0 {
		res := sprints.Closedobj(tab, name)
		err = t.Mark(v, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}
	sb.WriteString(tab.Sprint(name + "["))

	for i := 0; i < v.Len(); i++ {
		elem, err := sprintOf(tab.Inc(), v.Index(i), t)
		if err != nil {
			return "", err
		}
		sb.WriteString("\n" + elem + ",")
	}

	sb.WriteString("\n" + tab.Sprint("]"))
	res := sb.String()
	err = t.Mark(v, res)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	return res, nil
}

// sprintStruct extracts the sprint from a struct
func sprintStruct(tab indent.Tab, v reflect.Value, t *tracker.SprintTracker) (string, error) {
	// structs aren't underlying pointers.
	name, _ := structs.Name(v.Type())
	if v.NumField() == 0 {
		return sprints.BClosedobj(tab, name), nil
	}
	sb := strings.Builder{}
	sb.WriteString(tab.Sprint(fmx.SBold(name) + "{"))

	if !v.CanAddr() {
		v = values.OfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		fv, err := sprintField(tab.Inc(), v.Field(i), t)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		sb.WriteString(sprints.Ltfield(tab, v.Type().Field(i).Name, fv) + ",")
	}

	sb.WriteString("\n" + tab.Sprint("}"))
	return sb.String(), nil
}

func sprintUnsafePointer(tab indent.Tab, v reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, v, t); hasCache {
		return cache, err
	}

	name, _ := unsafes.Name(v.Type())
	addr := v.Pointer()
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
	err := t.Mark(v, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	res := tab.Sprintf("%s<%s>", name, sprints.Uintptrf(addr))
	err = t.Mark(v, res)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	return res, nil
}

func handleCache(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (output string, hasCache bool, err error) {
	cache, exists, err := t.Get(value)
	if exists {
		if err != nil {
			return sprints.Errorf(tab, "%s", err), exists, err
		}
		return tab.Sprint(cache), exists, nil
	}
	return
}
