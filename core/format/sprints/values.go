// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprints

import (
	"strconv"
	"strings"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/core/format"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
)

const addrSentence = "0x%d"

func Array[A ~[]T, T any](tab indent.Indentor, size int, typx T, data A) string {
	sb := strings.Builder{}
	sb.WriteString(tab.Sprintf(arraySentence, size, typx))
	for _, value := range data {
		sb.WriteString(tab.Inc().Sprintlnf(valueSentence, value))
	}
	sb.WriteString(tab.Sprintlnf(closedBracket))
	return sb.String()
}

const arraySentence = "[%d]%v{"

func Bool[B ~bool](tab indent.Indentor, v B) string {
	return tab.String() + strconv.FormatBool(bool(v))
}

func Byte[B ~byte](tab indent.Indentor, s B) string {
	return tab.Sprintf(byteSentence, s)
}

const byteSentence = "%d"

const closedBracket = "}"

func Complex[C nums.Complex](tab indent.Indentor, data C) string {
	return tab.Sprintf(complexSentence, data)
}

const complexSentence = "%g"
const digitSentence = "%d"

func Float[T nums.Float](tab indent.Indentor, data T) string {
	return tab.Sprintf(floatSentence, data)
}

const floatSentence = "%g"

func Int[I ~int](tab indent.Indentor, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Int8[I ~int8](tab indent.Indentor, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Int16[I ~int16](tab indent.Indentor, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Int32[I ~int32](tab indent.Indentor, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Int64[I ~int64](tab indent.Indentor, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Invalid(tab indent.Indentor) string {
	return tab.Sprint("<invalid>")
}

func Map[M ~map[K]V, K comparable, V any](tab indent.Indentor, m M, keyValue ...string) string {
	sb := strings.Builder{}
	if len(keyValue) > 0 {
		k, v := keyValue[0], keyValue[1]
		sb.WriteString(tab.Sprintf(mapSentence, k, v))
	} else {
		var k K
		var v V
		sb.WriteString(tab.Sprintf(unknownMapSentence, k, v))
	}
	for key, value := range m {
		sb.WriteString(tab.Inc().Sprintlnf(mapFieldSentence, key, value))
	}
	sb.WriteString(tab.Sprintlnf(closedBracket))
	return sb.String()
}

const unknownMapSentence = "map[%T]%T{"
const mapSentence = "map[%s]%s{}"
const mapFieldSentence = "%v: %v,"

func Nil(tab indent.Indentor) string {
	return tab.Sprint("<nil>")
}

const objectSentence = "%s{"

func Ptr(tab indent.Indentor, s string) string {
	return tab.Sprint("*" + FieldVal(s))
}

func TypedSlice[S ~[]T, T any](tab indent.Indentor, s S, typx ...string) string {
	var zero T
	sb := strings.Builder{}
	if len(typx) > 0 {
		sb.WriteString(tab.Sprintf(sliceSentence, typx[0]))
	} else {
		sb.WriteString(tab.Sprintf(unknownSliceSentence, zero))
	}

	for _, value := range s {
		if sp, ok := any(value).(format.Sprintable); ok {
			sb.WriteString(sp.Sprint(tab.Inc()))
		} else {
			sb.WriteString(tab.Inc().Sprintlnf(valueSentence, value))
		}
	}
	sb.WriteString(tab.Sprintlnf(closedBracket))
	return sb.String()
}

const unknownSliceSentence = "[]%T{"
const sliceSentence = "[]%s{"

func String(s string) string {
	if s == "" {
		return "<empty>"
	}
	return s
}

const stringSentence = "%q"

func Struct(tab indent.Indentor, n string, fields ...Tuple[string, string]) string {
	sb := strings.Builder{}
	sb.WriteString(tab.Sprintf(objectSentence, n))
	for _, field := range fields {
		sb.WriteString(tab.Inc().Sprintlnf(structFieldSentence, field.Key, field.Value))
	}
	sb.WriteString(tab.Sprintlnf(closedBracket))
	return sb.String()
}

const structFieldSentence = "%v %v,"

func Uint[I ~uint](tab indent.Indentor, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uint8[I ~uint8](tab indent.Indentor, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uint16[I ~uint16](tab indent.Indentor, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uint32[I ~uint32](tab indent.Indentor, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uint64[I ~uint64](tab indent.Indentor, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uintptr(tab indent.Indentor, s uintptr) string {
	return tab.Sprintf(addrSentence, s)
}

func UnsafePointer(tab indent.Indentor, s unsafe.Pointer) string {
	return tab.String() + "unsafe.Pointer(" + fmx.Sprintf(addrSentence, uintptr(s)) + ")"
}

const valueSentence = "%v"

func Value(tab indent.Indentor, value any) string {
	switch v := value.(type) {
	case nil:
		return "<nil>"
	case bool:
		return tab.String() + strconv.FormatBool(v)
	case int:
		return tab.String() + strconv.FormatInt(int64(v), 10)
	case int8:
		return tab.String() + strconv.FormatInt(int64(v), 10)
	case int16:
		return tab.String() + strconv.FormatInt(int64(v), 10)
	case int32:
		return tab.String() + strconv.FormatInt(int64(v), 10)
	case int64:
		return tab.String() + strconv.FormatInt(v, 10)

	case uint:
		return tab.String() + strconv.FormatUint(uint64(v), 10)
	case uint8:
		return tab.String() + strconv.FormatUint(uint64(v), 10)
	case uint16:
		return tab.String() + strconv.FormatUint(uint64(v), 10)
	case uint32:
		return tab.String() + strconv.FormatUint(uint64(v), 10)
	case uint64:
		return tab.String() + strconv.FormatUint(v, 10)
	case float32:
		return tab.String() + strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return tab.String() + strconv.FormatFloat(v, 'f', -1, 64)

	case complex64:
		r, i := real(v), imag(v)
		return tab.Sprintf("(%g + %gi)", r, i)
	case complex128:
		r, i := real(v), imag(v)
		return tab.Sprintf("(%g + %gi)", r, i)

	case string:
		return v
	default:
		return Errorf(tab, "not a value: %v", value)
	}
}
