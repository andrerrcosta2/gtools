// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprints

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"strconv"
	"strings"
	"unsafe"
)

func Addr[T any](value *T) string {
	return fmx.Sprintf("%p", unsafe.Pointer(value))
}

func BClosedobj(tab indent.Tab, name string, fields ...string) string {
	if len(fields) == 0 {
		return tab.Sprint(fmx.SBold(name) + " {}")
	}
	sb := strings.Builder{}
	sb.Grow(20 + len(name) + len(fields)*10)

	sb.WriteString(tab.Sprint(fmx.SBold(name)))
	sb.WriteString("{")
	for _, field := range fields {
		sb.WriteString("\n" + tab.Inc().Sprint(field) + ",")
	}
	sb.WriteString("\n" + tab.Sprint("}"))

	return sb.String()
}

func NestedBClosedobj(tab indent.Tab, name string, fields ...string) string {
	if len(fields) == 0 {
		return fmx.SBold(name) + " {}"
	}
	sb := strings.Builder{}
	sb.WriteString(fmx.SBold(name))
	sb.WriteString("{")
	for _, field := range fields {
		sb.WriteString("\n" + field + ",")
	}
	sb.WriteString("\n" + tab.Sprint("}"))
	return sb.String()
}

func BClosedObjError(tab indent.Tab, name string, fields ...string) string {
	return fmx.SRed(BClosedobj(tab, name, fields...))
}

func CyclicRef(tab indent.Tab, typ string, addr string) string {
	return tab.Sprintf("%s{ <cyclic-ref|%s> }", typ, addr)
}

func ClosedArray(tab indent.Tab, typ string, size int, fields ...string) string {
	if len(fields) == 0 {
		return tab.Sprint("[" + strconv.Itoa(size) + "]" + typ + "[<empty>]")
	}

	sb := strings.Builder{}
	sb.Grow(len(typ) + len(fields)*10)

	sb.WriteString("[" + strconv.Itoa(size) + "]" + typ + "[")

	for _, field := range fields {
		sb.WriteString("\n" + tab.Indentf(1, "%s,", field)) // tab.Inc().Sprint(field) + ",")
	}

	sb.WriteString("\n" + tab.Sprint("]"))

	return sb.String()
}

func ClosedAnonymous(tab indent.Tab, name string, fields ...string) string {
	if len(fields) == 0 {
		return tab.Sprint(name + " {}")
	}
	sb := strings.Builder{}
	sb.Grow(20 + len(name) + len(fields)*10)

	sb.WriteString(tab.Sprint(name))
	sb.WriteString("{")
	for _, field := range fields {
		sb.WriteString("\n" + tab.Inc().Sprint(field) + ",")
	}
	sb.WriteString("\n" + tab.Sprint("}"))

	return sb.String()
}

func ClosedMap(tab indent.Tab, key string, value string, fields ...string) string {
	if len(fields) == 0 {
		return tab.Sprint("map[" + key + "]" + value + " {}")
	}

	typ := "map[" + key + "]" + value
	sb := strings.Builder{}
	sb.Grow(len(typ) + len(fields)*10)

	sb.WriteString(tab.Sprint(typ + "{"))

	for _, field := range fields {
		sb.WriteString("\n" + tab.Indentf(1, "%s,", field))
	}

	sb.WriteString("\n" + tab.Sprint("}"))

	return sb.String()
}

func Closedobj(tab indent.Tab, name string, fields ...string) string {
	if len(fields) == 0 {
		return tab.Sprint(name + " {}")
	}
	sb := strings.Builder{}
	sb.Grow(20 + len(name) + len(fields)*10)

	sb.WriteString(tab.Sprint(name))
	sb.WriteString("{")
	for _, field := range fields {
		sb.WriteString("\n" + tab.Inc().Sprint(field) + ",")
	}
	sb.WriteString("\n" + tab.Sprint("}"))

	return sb.String()
}

func ClosedSlice(tab indent.Tab, typ string, fields ...string) string {
	if len(fields) == 0 {
		return tab.Sprint("[]" + typ + "[<empty>]")
	}
	sb := strings.Builder{}
	sb.Grow(len(typ) + len(fields)*10)
	sb.WriteString("[]" + typ + "[")
	for _, field := range fields {
		sb.WriteString("\n" + tab.Indentf(1, "%s,", field))
	}
	sb.WriteString("\n" + tab.Sprint("]"))
	return sb.String()
}

func Digit[T nums.Integer](data T) string {
	return fmx.Sprintf("%d", data)
}

func EmptyIterable(tab indent.Tab, typ string) string {
	return tab.Sprint(typ + "[<empty>]")
}

func Errorf(tab indent.Tab, format string, args ...interface{}) string {
	return tab.String() + fmx.SRedf(format, args...)
}

func Error(tab indent.Tab, data string) string { return tab.String() + fmx.SRed(data) }

func Field(tab indent.Tab, key, value string) string {
	return tab.Sprint(key + ": " + FieldVal(value))
}

func Fieldf(tab indent.Tab, key string, format string, args ...any) string {
	value := fmx.Sprintf(format, args...)
	return tab.Sprint(key) + ": " + FieldVal(value)
}

func Interface(tab indent.Tab, typ, elem string) string {
	return tab.Sprintf("%s{\n%s\n%s}", typ, elem, tab.String())
}

func Lclose(name string) string {
	return "\n" + name + "}"
}

func Line(data string) string { return fmx.Sprintf("\n%s", data) }

func Lnf(format string, args ...any) string {
	return "\n" + fmx.Sprintf(format, args...)
}

func Ltab(data string) string { return "\n\t" + data }

func Ltabf(format string, args ...any) string {
	return "\n\t" + fmx.Sprintf(format, args...)
}

func LtError(e string) string {
	return Ltab(e)
}

func LtErrorf(format string, args ...interface{}) string {
	return Ltabf(format, args...)
}

func Lfield(tab indent.Tab, key, value string) string {
	return "\n" + tab.Sprint(key) + ": " + FieldVal(value)
}

func Ltfield(tab indent.Tab, key, value string) string {
	return "\n\t" + tab.Sprint(key) + ": " + FieldVal(value)
}

func Ltfieldf(tab indent.Tab, key string, format string, args ...any) string {
	value := fmx.Sprintf(format, args...)
	field := "\n\t" + tab.Sprint(key) + ": " + FieldVal(value)
	return field
}

func FieldVal(value string) string {
	return strings.TrimLeftFunc(value, func(r rune) bool {
		return r == '\t' || r == '\n' || r == ' '
	})
}

func FieldValf(format string, args ...any) string {
	return FieldVal(fmx.Sprintf(format, args...))
}

func MapEntry(tab indent.Tab, key, value string) string {
	return tab.String() + key + ": " + value + ","
}

func NilPointer(tab indent.Tab) string {
	return tab.Sprint("<nil>")
}

func NilType(tab indent.Tab, typ string) string {
	return tab.Sprint(typ + "<nil>")
}

func Openobj(tab indent.Tab, name string, fields ...string) string {
	sb := strings.Builder{}
	sb.WriteString(tab.Sprint(fmx.SBold(name)))
	sb.WriteString("{")
	sb.WriteString(strings.Join(fields, ""))
	return sb.String()
}

func OptField(tab indent.Tab, field, value string) string {
	if field == "" {
		return tab.Sprint(value)
	}
	return tab.Sprint(field + ": " + FieldVal(value))
}

func PointerElem(tab indent.Tab, elem string) string {
	return tab.String() + "*" + elem
}

func Quoted(s string) string {
	if s == "" {
		return "<empty>"
	}
	return fmx.Sprintf("%q", s)
}

func Structf(tab indent.Tab, name string, fields ...Tuple[string, string]) string {
	sb := strings.Builder{}
	sb.WriteString(tab.Sprint(fmx.SBold(name) + " {\n"))
	for _, field := range fields {
		sb.WriteString(tab.Sprint(fmx.SBoldf("\t%s: ", field.Key) + field.Value))
	}
	sb.WriteString("\n}")
	return sb.String()
}

func Tabf(format string, args ...any) string {
	return fmx.Sprintf("\t%s", fmx.Sprintf(format, args...))
}

func Tfield(tab indent.Tab, key, value string) string {
	return "\t" + tab.Sprint(key) + ": " + FieldVal(value)
}

func Tfieldf(tab indent.Tab, key string, format string, args ...any) string {
	value := fmx.Sprintf(format, args...)
	return "\t" + tab.Sprint(key) + ": " + FieldVal(value)
}

func TypedString(s string) string {
	if s == "" {
		return "<string>(empty)"
	}
	return fmx.Sprintf("<string>%q", s)
}

func Typed(typ string, value any) string {
	return fmx.Sprintf("<%s>%v", typ, Valuef(value))
}

func TypedBool(value bool) string {
	return fmx.Sprintf("<bool>%t", value)
}

func TypedByte(value byte) string {
	return fmx.Sprintf("<byte>%d", value)
}

func TypedComplex64(value complex128) string {
	return fmx.Sprintf("<complex64>%v", value)
}

func TypedComplex128(value complex128) string {
	return fmx.Sprintf("<complex128>%v", value)
}

func TypeValue(tab indent.Tab, typ, value string) string {
	return tab.Sprint(typ + " " + value)
}

func TypedDigit[T nums.Integer](typ string, data T) string {
	return fmx.Sprintf("<%s>%d", typ, data)
}

func TypedFloat32(value float32) string {
	return fmx.Sprintf("<float32>%f", value)
}

func TypedFloat64(value float32) string {
	return fmx.Sprintf("<float64>%f", value)
}

func TypedRoundFloat[T nums.Float](typ string, data T) string {
	return fmx.Sprintf("<%s>%g", typ, data)
}

// StripANSI removes ANSI escape codes (like \x1b[1m) from a string
func StripANSI(input string) string {
	var result []rune
	inEscape := false

	for _, r := range input {
		// If we're inside an escape sequence, ignore characters until the end of the sequence
		if r == '\x1b' {
			inEscape = true
		}

		if inEscape {
			if r == 'm' {
				inEscape = false // end of escape sequence
			}
			continue // skip characters inside escape sequences
		}

		result = append(result, r)
	}

	return string(result)
}

func trimmedFields(fields ...any) []string {
	strs := toStringSlice(fields...)
	for i, s := range strs {
		strs[i] = strings.TrimSpace(s) // Trim both leading and trailing spaces globally
	}
	return strs
}

func toStringSlice(fields ...any) []string {
	out := make([]string, len(fields))
	for i, f := range fields {
		out[i] = fmx.Sprintf("%v", f)
	}
	return out
}

func Uintptrf(value uintptr) string {
	return fmx.Sprintf("0x%x", value)
}

func UnsafeAddrf(value unsafe.Pointer) string {
	return fmx.Sprintf("<unsafe.Pointer>0x%x\"", uintptr(value))
}
