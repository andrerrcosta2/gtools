// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprints

import (
	"strconv"
	"strings"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
)

// Slicef formats the elements of a slice into a single string by applying
// the function `f` to each element. Elements are joined with ", ".
//
// Example:
//
//	nums := []int{1, 2, 3}
//	res := Slicef(nums, func(i int, t indent.Tab) string {
//	    return strconv.Itoa(i)
//	}, indent.None)
//
//	// Output: "1, 2, 3"
func Slicef[I any](t indent.Indentor, arr []I, f func(t indent.Indentor, i I) string) string {
	if len(arr) == 0 {
		return ""
	}

	sb := strings.Builder{}

	for i := 0; i < len(arr)-1; i++ {
		sb.WriteString(f(t, arr[i]))
		sb.WriteString(", ")
	}

	sb.WriteString(f(t, arr[len(arr)-1]))
	return sb.String()
}

// AnySlicef safely formats a slice of any type (`any`) into a string using the
// function `f`. It handles nils and unexpected types gracefully.
//
// If arr is nil → returns "<nil>"
// If arr is not of type S → returns "<unexpected type: %T>"
//
// Example:
//
//	var data any = []string{"a", "b", "c"}
//	res := AnySlicef[[]string, string](data, func(s string, t indent.Tab) string {
//	    return "'" + s + "'"
//	}, indent.None)
//
//	// Output: "'a', 'b', 'c'"
func AnySlicef[S ~[]T, T any](t indent.Indentor, arr any, f func(t indent.Indentor, i T) string) string {
	if arr == nil {
		return "<nil>"
	}
	a, ok := arr.(S)

	if !ok {
		return t.Sprintf("<unexpected type: %T>", arr)
	}

	return Slicef(t, a, f)
}

// SliceNf generates a comma-separated string by calling the function `f`
// exactly `n` times — once for each index from 0 to n-1.
//
// Example:
//
//	res := SliceNf(3, func(i int, t indent.Tab) string {
//	    return "Item" + strconv.Itoa(i)
//	}, indent.None)
//
//	// Output: "Item0, Item1, Item2"
func SliceNf(t indent.Indentor, n int, f func(t indent.Indentor, i int) string) string {
	sb := strings.Builder{}

	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(f(t, i))
	}

	return sb.String()
}

// Mapf formats a map values to a string using the output of function f
// for each element.
func Mapf[K comparable, V any, M ~map[K]V](t indent.Indentor, m M, f func(t indent.Indentor, k K, v V) string) string {
	sb := strings.Builder{}
	sb.WriteString("{\n")
	k := 0
	for key, value := range m {
		sb.WriteString(f(t, key, value) + ",\n")
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
