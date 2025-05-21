// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprints

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/core/format"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"strconv"
	"strings"
	"unsafe"
)

const addrSentence = "0x%d"

func Array[A ~[]T, T any](tab indent.Tab, size int, typx T, data A) string {
	sb := strings.Builder{}
	sb.WriteString(tab.Sprintf(arraySentence, size, typx))
	for _, value := range data {
		sb.WriteString(tab.Inc().Sprintlnf(valueSentence, value))
	}
	sb.WriteString(tab.Sprintlnf(closedBracket))
	return sb.String()
}

const arraySentence = "[%d]%v{"

func Bool[B ~bool](tab indent.Tab, data B) string {
	return tab.Sprintf(boolSentence, data)
}

const boolSentence = "%t"

func Byte[B ~byte](tab indent.Tab, s B) string {
	return tab.Sprintf(byteSentence, s)
}

const byteSentence = "%d"

const closedBracket = "}"

func Complex[C nums.Complex](tab indent.Tab, data C) string {
	return tab.Sprintf(complexSentence, data)
}

const complexSentence = "%g"
const digitSentence = "%d"

func Float[T nums.Float](tab indent.Tab, data T) string {
	return tab.Sprintf(floatSentence, data)
}

const floatSentence = "%g"

func Int[I ~int](tab indent.Tab, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Int8[I ~int8](tab indent.Tab, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Int16[I ~int16](tab indent.Tab, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Int32[I ~int32](tab indent.Tab, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Int64[I ~int64](tab indent.Tab, s I) string {
	return tab.Sprintf(strconv.Itoa(int(s)))
}

func Invalid(tab indent.Tab) string {
	return tab.Sprint("<invalid>")
}

func Map[M ~map[K]V, K comparable, V any](tab indent.Tab, m M, keyValue ...string) string {
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

func Nil(tab indent.Tab) string {
	return tab.Sprint("<nil>")
}

const objectSentence = "%s{"

func Ptr(tab indent.Tab, s string) string {
	return tab.Sprint("*" + FieldVal(s))
}

func TypedSlice[S ~[]T, T any](tab indent.Tab, s S, typx ...string) string {
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

func Struct(tab indent.Tab, n string, fields ...Tuple[string, string]) string {
	sb := strings.Builder{}
	sb.WriteString(tab.Sprintf(objectSentence, n))
	for _, field := range fields {
		sb.WriteString(tab.Inc().Sprintlnf(structFieldSentence, field.Key, field.Value))
	}
	sb.WriteString(tab.Sprintlnf(closedBracket))
	return sb.String()
}

const structFieldSentence = "%v %v,"

func Uint[I ~uint](tab indent.Tab, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uint8[I ~uint8](tab indent.Tab, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uint16[I ~uint16](tab indent.Tab, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uint32[I ~uint32](tab indent.Tab, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uint64[I ~uint64](tab indent.Tab, s I) string {
	return tab.Sprint(strconv.Itoa(int(s)))
}

func Uintptr(tab indent.Tab, s uintptr) string {
	return tab.Sprintf(addrSentence, s)
}

func UnsafePointer(tab indent.Tab, s unsafe.Pointer) string {
	return tab.Sprintf(addrSentence, uintptr(s))
}

const valueSentence = "%v,"
