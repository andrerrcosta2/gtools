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

func Value(tab indent.Tab, value reflect.Value) (string, error) {
	return vl(tab, value, tracker.Sprint())
}

func vl(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if !value.IsValid() {
		return "invalid type: <invalid>", gerrors.NewError(gerrors.Input_err, reflect4.ErrInvalidValue)
	}

	if value.Kind() <= reflect.Complex128 || value.Kind() == reflect.String {
		return pr(tab, value), nil
	}

	switch value.Kind() {
	case reflect.Array:
		return ar(tab, value, t)
	case reflect.Chan:
		return ch(tab, value), nil
	case reflect.Func:
		return fn(tab, value), nil
	case reflect.Interface:
		return in(tab, value, t)
	case reflect.Map:
		return mp(tab, value, t)
	case reflect.Ptr:
		return pt(tab, value, t)
	case reflect.Slice:
		return sl(tab, value, t)
	case reflect.Struct:
		return st(tab, value, t)
	case reflect.UnsafePointer:
		return up(tab, value, t)
	default:
		return sprints.Errorf(tab, "invalid type: %s", value.Kind().String()), nil
	}
}

// ar extracts the sprint from an array
func ar(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, value, t); hasCache {
		return cache, err
	}

	name, _ := arrays.Name(value.Type())
	addr, err := pointers.Of(value)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	err = t.Mark(value, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}

	sb := strings.Builder{}
	if value.Len() == 0 {
		res := tab.Sprint(name + "[<empty>]")
		err = t.Mark(value, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}
	sb.WriteString(tab.Sprint(name) + "[")

	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		v, err := vl(tab.Inc(), elem, t)
		if err != nil {
			return "", err
		}
		sb.WriteString("\n" + v + ",")
	}

	sb.WriteString("\n" + tab.Sprint("]"))
	res := sb.String()
	err = t.Mark(value, res)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	return res, nil
}

// ch extracts the sprint from a channel
func ch(tab indent.Tab, value reflect.Value) string {
	name, _ := chans.Name(value.Type())
	if value.IsNil() {
		return sprints.NilType(tab, name)
	}
	return name
}

// fd extracts the sprint from a field
func fd(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	// CanInterface() means:
	// - Calling .Interface() is safe (won't panic).
	// - The value is exported.
	// - If value is a pointer, calling .Interface() gives you that pointer wrapped inside interface{}.
	// - If value is a non-pointer value, .Interface() gives you a copy of that value inside interface{}.
	if value.CanInterface() || value.CanAddr() {
		return vl(tab, value, t)
	}

	// If neither CanInterface() nor CanAddr() is true, the value:
	// - Is unexported (hence unsafe to use .Interface()).
	// - Is not directly addressable.
	// - We must use unsafe.Pointer and reflect.NewAt() to access it.
	return vl(tab.Inc(), values.OfUnaddr(value).Elem(), t)
}

func in(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, value, t); hasCache {
		return cache, err
	}

	name, _ := interfaces.Name(value.Type())
	addr, err := pointers.Of(value)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	err = t.Mark(value, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	val, err := vl(tab.Inc(), value.Elem(), t)
	if err != nil {
		return "", err
	}
	res := tab.Sprintf("%s(%s)", name, val)
	err = t.Mark(value, res)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	return res, nil
}

// fn extracts the sprint from a function
func fn(tab indent.Tab, value reflect.Value) string {
	name, _ := funcs.Name(value.Type())
	return tab.Sprint(name)
}

// mp extracts the sprint from a map
func mp(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, value, t); hasCache {
		return cache, err
	}

	name, _ := maps.Name(value.Type())
	addr, err := pointers.Of(value)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
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
	var keys []string
	for _, key := range value.MapKeys() {
		keys = append(keys, key.String())
	}
	sort.Strings(keys)

	// Start building the map representation
	sb.WriteString(tab.Sprint(name) + "{")
	for _, key := range keys {
		// Get the value for each sorted key
		val := value.MapIndex(reflect.ValueOf(key))

		// Print the key-value pair
		k, err := vl(tab.Inc(), reflect.ValueOf(key), t)
		if err != nil {
			return "", err
		}
		sb.WriteString("\n" + k)
		sb.WriteString(": ")
		v, err := vl(tab.Inc(), val, t)
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

func pt(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, value, t); hasCache {
		return cache, err
	}

	addr, err := pointers.Of(value)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	name, _ := pointers.Name(value.Type())
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
	err = t.Mark(value, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	if value.IsNil() {
		res := sprints.NilType(tab, name)
		err = t.Mark(value, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}
	val, err := vl(tab, value.Elem(), t)
	return "*" + strings.TrimSpace(val), err
}

// pr extracts a sprint from a primitive
func pr(tab indent.Tab, s reflect.Value) string {
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

// sl extracts the sprint from a slice
func sl(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, value, t); hasCache {
		return cache, err
	}

	name, _ := slices.Name(value.Type())
	addr, err := pointers.Of(value)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
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
	if value.Len() == 0 {
		res := sprints.Closedobj(tab, name)
		err = t.Mark(value, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}
	sb.WriteString(tab.Sprint(name + "["))

	for i := 0; i < value.Len(); i++ {
		elem, err := vl(tab.Inc(), value.Index(i), t)
		if err != nil {
			return "", err
		}
		sb.WriteString("\n" + elem + ",")
	}

	sb.WriteString("\n" + tab.Sprint("]"))
	res := sb.String()
	err = t.Mark(value, res)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	return res, nil
}

// st extracts the sprint from a struct
func st(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, value, t); hasCache {
		return cache, err
	}

	name, _ := structs.Name(value.Type())
	addr, err := pointers.Of(value)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
	err = t.Mark(value, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	if value.NumField() == 0 {
		res := sprints.BClosedobj(tab, name)
		err = t.Mark(value, res)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		return res, nil
	}
	sb := strings.Builder{}
	sb.WriteString(tab.Sprint(fmx.SBold(name) + "{"))

	if !value.CanAddr() {
		value = values.OfUnaddr(value)
	}

	for i := 0; i < value.NumField(); i++ {
		v, err := fd(tab.Inc(), value.Field(i), t)
		if err != nil {
			return sprints.Errorf(tab, "%s", err), err
		}
		sb.WriteString(sprints.Ltfield(tab, value.Type().Field(i).Name, v) + ",")
	}

	sb.WriteString("\n" + tab.Sprint("}"))
	res := sb.String()
	err = t.Mark(value, res)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	return res, nil
}

func up(tab indent.Tab, value reflect.Value, t *tracker.SprintTracker) (string, error) {
	if cache, hasCache, err := handleCache(tab, value, t); hasCache {
		return cache, err
	}

	name, _ := unsafes.Name(value.Type())
	addr, err := pointers.Of(value)
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	// The correct approach is to mark it as cycle early and replace it on the same level
	// before the return. Because it isn't a cycle just because it will be printed
	// again later.
	err = t.Mark(value, sprints.CyclicRef(indent.Zero(), name, sprints.Uintptrf(addr)))
	if err != nil {
		return sprints.Errorf(tab, "%s", err), err
	}
	res := tab.Sprintf("%s<%s>", name, sprints.Uintptrf(addr))
	err = t.Mark(value, res)
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
