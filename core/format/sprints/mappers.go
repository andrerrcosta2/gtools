// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprints

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"strconv"
	"strings"
)

// Slicef formats a slice values to a string using the output of function f
// for each element.
func Slicef[I any](arr []I, f func(i I, t indent.Tab) string, t indent.Tab) string {
	if len(arr) == 0 {
		return ""
	}

	sb := strings.Builder{}

	for i := 0; i < len(arr)-1; i++ {
		sb.WriteString(f(arr[i], t))
		sb.WriteString(", ")
	}

	sb.WriteString(f(arr[len(arr)-1], t))
	return sb.String()
}

// AnySlicef formats a slice of any type values to a string using the output of function f
// for each element.
func AnySlicef[S ~[]T, T any](arr any, f func(i T, t indent.Tab) string, t indent.Tab) string {
	if arr == nil {
		return "<nil>"
	}
	a, ok := arr.(S)

	if !ok {
		return t.Sprintf("<unexpected type: %T>", arr)
	}

	return Slicef(a, f, t)
}

func SliceNf(n int, f func(i int, t indent.Tab) string, t indent.Tab) string {
	sb := strings.Builder{}

	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(f(i, t))
	}

	return sb.String()
}

// Mapf formats a map values to a string using the output of function f
// for each element.
func Mapf[K comparable, V any, M ~map[K]V](m M, f func(k K, v V, t indent.Tab) string, t indent.Tab) string {
	sb := strings.Builder{}
	sb.WriteString("{\n")
	k := 0
	for key, value := range m {
		sb.WriteString(f(key, value, t) + ",\n")
		k++
	}
	sb.WriteString(t.Sprintf("}"))
	return sb.String()
}

func Valuef(value any) string {
	switch v := value.(type) {
	case nil:
		return "<nil>"
	case bool:
		return strconv.FormatBool(v)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case complex64:
		return strconv.FormatComplex(complex128(v), 'g', -1, 64)
	case complex128:
		return strconv.FormatComplex(v, 'g', -1, 128)
	case string:
		return String(v)
	default:
		return fmx.Sprintf("%T(%v)", v, v) // Keep fmx.Sprintf only for unknown types
	}
}
