// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/printer"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"strings"
)

// Message returns a formatted standard diff message
func Message(msg, diff string) string {
	if diff == "" {
		return clean(msg)
	}
	return clean(printer.Sprintf("%s\n%s", msg, diff))
}

func Append(tab indent.Tab, whole, msg string) string {
	out := tab.Smarkf("%s\n%s", whole, tab.Indent(1, msg))
	return out
}

func BothNils(tab indent.Tab, kind, typxA, typxB string) string {
	return tab.Sprintf(BothNilsFmt, kind, recExp(tab, printer.Sprintf("%s<nil>", typxA), printer.Sprintf("%s<nil>", typxB)))
}

const BothNilsFmt = "Both %s are nil:%s"

func BothInvalid(tab indent.Tab) string {
	return tab.Sprintf(BothInvalidFmt)
}

const BothInvalidFmt = "Both values are invalid"

func Chain(tab indent.Tab, msg ...string) string {
	var cb strings.Builder
	for i, m := range msg {
		if m == "" {
			continue
		}
		ind := tab.Indent(i, m)
		cb.WriteString(tab.Smarkf("%s", ind))

		if i < len(msg)-1 {
			cb.WriteString("\n")
		}
	}
	return cb.String()
}

func ChainMessage(tab indent.Tab, diff string, msg ...string) string {
	if len(msg) == 0 {
		return tab.Sprint(diff)
	}
	var cb strings.Builder
	for i, m := range msg {
		ct := tab.Plus(i)
		if i > 0 {
			cb.WriteString("\n")
		}
		cb.WriteString(ct.Smarkf(m))
	}
	return Message(cb.String(), tab.Indent(len(msg), diff))
}

func Diff(tab indent.Tab, a, b string) string {
	t := tab.String()
	diff := t + "(Received...):\n" +
		tab.Indent(1, a) + "\n" +
		t + "(...)\n\n" +
		t + "(Expected...):\n" +
		tab.Indent(1, b) + "\n" +
		t + "(...)"
	return diff
}

func ArrayLenMismatch(tab indent.Tab, a, b int) string {
	return tab.Smarkf(ArrayLenMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const ArrayLenMismatchFmt = "array lengths mismatch:%s"

// ArrayElem returns a formatted standard diff message for an array with different
// element values at a specific index
func ArrayElem(tab indent.Tab, idx int) string {
	return tab.Sprintf(ArrayElemFmt, idx)
}

const ArrayElemFmt = "array elements mismatch at index '%d':"

func ArrayTypesMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(ArrayTypesMismatchFmt, recExp(tab, a, b))
}

const ArrayTypesMismatchFmt = "array types mismatch:%s"

func ChanTypesMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(ChanTypesMismatchFmt, recExp(tab, a, b))
}

const ChanTypesMismatchFmt = "chan types mismatch:%s"

func ChanDirMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(ChanDirMismatchFmt, recExp(tab, a, b))
}

const ChanDirMismatchFmt = "chan direction mismatch:%s"

func ChanBufSizeMismatch(tab indent.Tab, a, b int) string {
	return tab.Sprintf(ChanBufSizeMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const ChanBufSizeMismatchFmt = "chan buffer sizes mismatch:%s"

func Empty(tab indent.Tab) string {
	return tab.Sprint(EmptyFmt)
}

const EmptyFmt = ""

func EqualsCyclicRef(tab indent.Tab, typx string, addr uintptr) string {
	return tab.Sprint(EqualsCyclicRefFmt, typx, sprints.Uintptrf(addr))
}

const EqualsCyclicRefFmt = "%s<cyclic-reference|%s>"

func FuncSignMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(FuncSignMismatchFmt, recExp(tab, a, b))
}

const FuncSignMismatchFmt = "func mismatch:%s"

func FuncTypesMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(FuncTypesMismatchFmt, recExp(tab, a, b))
}

const FuncTypesMismatchFmt = "func types mismatch:%s"

func InterfaceTypesMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(InterfaceTypesMismatchFmt, recExp(tab, a, b))
}

const InterfaceTypesMismatchFmt = "interface types mismatch:%s"

func InterfaceImpl(tab indent.Tab) string {
	return tab.Sprint(InterfaceImplFmt)
}

const InterfaceImplFmt = "implementations are different:"

func InvalidExpected(tab indent.Tab, a string) string {
	return tab.Sprintf(InvalidExpectedFmt, recExp(indent.Zero(), a, sprints.Invalid(indent.Zero())))
}

const InvalidExpectedFmt = "Expected value is invalid:%s"

func InvalidReceived(tab indent.Tab, b string) string {
	return tab.Sprintf(InvalidReceivedFmt, recExp(indent.Zero(), sprints.Invalid(indent.Zero()), b))
}

const InvalidReceivedFmt = "Received value is invalid:%s"

func MapTypesMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(MapTypesMismatchFmt, recExp(tab, a, b))
}

const MapTypesMismatchFmt = "map types mismatch:%s"

func MapLenMismatch(tab indent.Tab, a, b int) string {
	return tab.Sprintf(MapLenMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const MapLenMismatchFmt = "map lengths mismatch:%s"

func MapKeys(tab indent.Tab) string {
	return tab.Sprint(MapKeysFmt)
}

const MapKeysFmt = "map keys mismatch:"

func MapKeysDiff(tab indent.Tab, missingKeysA, extraKeysA, missingKeysB, extraKeysB []string) string {
	t := tab.String()
	result := t + "(Received...):\n"

	if len(missingKeysA) > 0 {
		result += tab.Indent(1, "missing keys", strings.Join(missingKeysA, ", "), "\n")
	}

	if len(extraKeysA) > 0 {
		result += tab.Indent(1, "extra keys:", strings.Join(extraKeysA, ", "), "\n")
	}

	result += t + "(...)\n\n" +
		t + "(Expected...):\n"

	if len(missingKeysB) > 0 {
		result += tab.Indent(1, "missing keys:", strings.Join(missingKeysB, ", "), "\n")
	}

	if len(extraKeysB) > 0 {
		result += tab.Indent(1, "extra keys: ", strings.Join(extraKeysB, ", "), "\n")
	}

	result += t + "(...)"
	return result
}

func MapValue(tab indent.Tab, key string) string {
	return tab.Sprintf(MapValueFmt, key)
}

const MapValueFmt = "map values are different for key '%s':"

func MissingField(tab indent.Tab, value string) string {
	return tab.Sprintf(MissingFieldFmt, value)
}

const MissingFieldFmt = "Field '%s' is missing"

func NilChan(tab indent.Tab, id string) string {
	return tab.Sprintf(NilChanFmt, id)
}

const NilChanFmt = "The %s channel is nil"

func NilFunc(tab indent.Tab, id string) string {
	return tab.Sprintf(NilFuncFmt, id)
}

const NilFuncFmt = "The %s function is nil"

func NilInterface(tab indent.Tab, id string) string {
	return tab.Sprintf(NilInterfaceFmt, id)
}

const NilInterfaceFmt = "The %s interface is nil"

func NilMap(tab indent.Tab, id string) string {
	return tab.Sprintf(NilMapFmt, id)
}

const NilMapFmt = "The %s map is nil"

func NilPointer(tab indent.Tab, id string) string {
	return tab.Sprintf(NilPointerFmt, id)
}

const NilPointerFmt = "The %s pointer is nil"

func NilSlice(tab indent.Tab, id string) string {
	return tab.Sprintf(NilSliceFmt, id)
}

const NilSliceFmt = "The %s slice is nil"

func NilReceived(tab indent.Tab, typx, expected string) string {
	return tab.Sprintf(NilReceivedFmt, typx, recExp(tab, printer.Sprintf("%s<nil>", typx), expected))
}

const NilReceivedFmt = "The '%s' received is nil:%s"

func NilExpected(tab indent.Tab, typx, received string) string {
	return tab.Sprintf(NilExpectedFmt, typx, recExp(tab, received, printer.Sprintf("%s<nil>", typx)))
}

const NilExpectedFmt = "The '%s' expected is nil:%s"

func Pointers(tab indent.Tab) string {
	return tab.Sprint(PointerValueFmt)
}

const PointerValueFmt = "pointer values are different:"

func Space(tab indent.Tab, amount int) string {
	return tab.Sprint(strings.Repeat("␣", amount))
}

func SliceCapMismatch(tab indent.Tab, a, b int) string {
	return tab.Sprintf(SliceCapMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const SliceCapMismatchFmt = "slice capacities mismatch:%s"

func SliceLenMismatch(tab indent.Tab, a, b int) string {
	return tab.Sprintf(SliceLenMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const SliceLenMismatchFmt = "slice lengths mismatch:%s"

func SliceTypesMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(SliceTypesMismatchFmt, recExp(tab, a, b))
}

const SliceTypesMismatchFmt = "slice types mismatch:%s"

func SliceValues(tab indent.Tab, idx int) string {
	return tab.Sprintf(SliceValuesFmt, idx)
}

const SliceValuesFmt = "slice values differ at index %d:"

func Strings(tab indent.Tab, idx int) string {
	return tab.Sprintf(StringsFmt, idx)
}

const StringsFmt = "strings differ at index %d:"

func StructTypesMismatch(tab indent.Tab, a, b string) string {
	return tab.Sprintf(StructTypesMismatchFmt, recExp(tab, a, b))
}

const StructTypesMismatchFmt = "struct types mismatch:%s"

func StructFields(tab indent.Tab, field string) string {
	return tab.Sprintf(StructFieldsFmt, field)
}

const StructFieldsFmt = "struct field '%s' mismatch:"

// TypesMismatch returns a types mismatch Message
func TypesMismatch(tab indent.Tab, a, b string) string {
	a = "<" + a + ">"
	b = "<" + b + ">"
	return tab.Sprintf(TypesMismatchFmt, recExp(tab, a, b))
}

const TypesMismatchFmt = "types mismatch:%s"

func UnsafePointers(tab indent.Tab) string {
	return tab.Sprint(UnsafePointersFmt)
}

const UnsafePointersFmt = "unsafe pointers values are different:"

func UnsafePointersAddr(tab indent.Tab, a, b string) string {
	return tab.Sprintf(UnsafePointersAddrFmt, recExp(tab, a, b))
}

const UnsafePointersAddrFmt = "unsafe pointers addresses mismatch:%s"

func Values(tab indent.Tab, a, b string) string {
	return tab.Sprintf(ValuesFmt, recExp(tab, a, b))
}

const ValuesFmt = "values mismatch:%s"

func clean(s string) string {
	return strings.TrimRight(s, "\t\n ")
}

func recExp(tab indent.Tab, rec, exp string) string {
	received := tab.Inc().Sprintf("→ Received: %s", rec)
	expected := tab.Inc().Sprintf("→ Expected: %s", exp)

	// Combine the results
	return printer.Sprintf("\n%s\n%s", received, expected)
}
