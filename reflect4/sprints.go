// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/core/format/sprints/builder"
	"reflect"
	"strings"
	"unsafe"
)

func Any(tab indent.Tab, value any) string {
	return an(tab, value)
}

func an(tab indent.Tab, value any) string {
	switch v := value.(type) {
	case string:
		return pr(tab, v, "string")
	case bool:
		return pr(tab, v, "bool")
	case uint8:
		return pr(tab, v, "uint8")
	case uint16:
		return pr(tab, v, "uint16")
	case uint32:
		return pr(tab, v, "uint32")
	case uint64:
		return pr(tab, v, "uint64")
	case int8:
		return pr(tab, v, "int8")
	case int16:
		return pr(tab, v, "int16")
	case int32:
		return pr(tab, v, "int32")
	case int64:
		return pr(tab, v, "int64")
	case float32:
		return pr(tab, v, "float32")
	case float64:
		return pr(tab, v, "float64")
	default:
		return ob(tab, reflect.ValueOf(v))
	}
}

func pr(tab indent.Tab, s any, typx string) string {
	return tab.Sprintf("<%s>%s", typx, sprints.Valuef(s))
}

func ob(tab indent.Tab, value reflect.Value) string {
	if value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		if !value.IsValid() || value.IsNil() {
			return sprints.Nil(tab)
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Bool, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return pr(tab, value.Interface(), value.Kind().String())
	case reflect.Struct:
		return st(tab, value)
	case reflect.Slice:
		return sl(tab, value)
	case reflect.Map:
		return mp(tab, value)
	case reflect.Chan:
		return ch(tab, value)
	case reflect.Func:
		return fn(tab, value)
	case reflect.Ptr, reflect.Interface:
		return an(tab, value.Elem().Interface())
	default:
		return sprints.Errorf(tab, "invalid type: %s", value.Kind().String())
	}
}

func st(tab indent.Tab, value reflect.Value) string {
	stc := builder.Struct(tab, value.Type().Name())
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		addr := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr()))
		stc.Field(value.Type().Field(i).Name, ob(tab.Inc(), addr))
	}
	return stc.String()
}

func sl(tab indent.Tab, value reflect.Value) string {
	// Check if the slice is nil
	if value.IsNil() {
		return sprints.NilType(tab, value.Type().String())
	}

	sb := strings.Builder{}
	sb.WriteString(tab.Sprint(fmx.SBold(value.Type().String()) + "["))

	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		sb.WriteString(ob(tab.Inc(), elem))

		if i < value.Len()-1 {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

func mp(tab indent.Tab, value reflect.Value) string {
	if value.IsNil() {
		return sprints.NilType(tab, value.Type().String())
	}

	sb := strings.Builder{}
	sb.WriteString(sprints.TypeValue(tab, value.Type().String(), "{"))

	for i, key := range value.MapKeys() {
		val := value.MapIndex(key)

		sb.WriteString(ob(tab.Inc(), key))
		sb.WriteString(": ")
		sb.WriteString(ob(tab.Inc(), val))

		if i < len(value.MapKeys())-1 {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("}")
	return sb.String()
}

func ch(tab indent.Tab, value reflect.Value) string {
	if value.IsNil() {
		return sprints.NilType(tab, value.Type().String())
	}

	dir := value.Type().ChanDir()
	var dirStr string
	switch dir {
	case reflect.RecvDir:
		dirStr = "<-chan"
	case reflect.SendDir:
		dirStr = "chan<-"
	case reflect.BothDir:
		dirStr = "chan"
	default:
		dirStr = "<?>chan"
	}

	return sprints.TypeValue(tab, dirStr, value.Type().Elem().String())
}

func fn(tab indent.Tab, value reflect.Value) string {
	strings.TrimPrefix(value.Type().String(), "func")
	return sprints.TypeValue(tab, "func", value.Type().String())
}
