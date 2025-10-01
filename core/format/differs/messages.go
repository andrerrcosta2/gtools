// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/printer"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"strings"
)

// Message sprints as:
//
//	'<msg>
//	 <diff>'
func Message(msg, diff string) string {
	if diff == "" {
		return clean(msg)
	}
	return clean(printer.Sprintf("%s\n%s", msg, diff))
}

// Append sprints as:
//
//	'<whole>
//	 <msg>'
func Append(tab indent.Indentor, whole, msg string) string {
	return tab.Sprintf("%s\n%s", whole, msg)
}

func BothInvalid(tab indent.Indentor) string {
	return tab.Sprintf(BothInvalidFmt)
}

const BothInvalidFmt = "Both values are invalid"

func Chain(msg ...string) string {
	var cb strings.Builder
	for i, m := range msg {
		if m == "" {
			continue
		}
		cb.WriteString(m)
		if i < len(msg)-1 {
			cb.WriteString("\n")
		}
	}
	return cb.String()
}

func ChainMessage(tab indent.Indentor, diff string, msg ...string) string {
	if len(msg) == 0 {
		return tab.Sprint(diff)
	}
	var cb strings.Builder
	for i, m := range msg {
		ct := tab.Plus(i)
		if i > 0 {
			cb.WriteString("\n")
		}
		cb.WriteString(ct.Sprint(m))
	}
	return Message(cb.String(), tab.Indent(len(msg), diff))
}

func Diff(tab indent.Indentor, a, b string) string {
	t := tab.String()
	diff := t + "(Received...):\n" +
		tab.Indent(1, a) + "\n" +
		t + "(...)\n\n" +
		t + "(Expected...):\n" +
		tab.Indent(1, b) + "\n" +
		t + "(...)"
	return diff
}

// ArrayLenMismatch sprint as:
//
//	'array lengths mismatch
//	 → Received: %s
//	 → Expected: %s'
func ArrayLenMismatch(tab indent.Indentor, a, b int) string {
	return tab.Sprintf(ArrayLenMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const ArrayLenMismatchFmt = "array lengths mismatch:%s"

// ArrayElem returns a formatted standard diff message for an array with different
// element values at a specific index
func ArrayElem(tab indent.Indentor, idx int) string {
	return tab.Sprintf(ArrayElemFmt, idx)
}

const ArrayElemFmt = "array elements mismatch at index '%d':"

// ArrayTypesMismatch sprint as:
//
//	'array types mismatch
//	 → Received: %s
//	 → Expected: %s'
func ArrayTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(ArrayTypesMismatchFmt, recExp(tab, a, b))
}

const ArrayTypesMismatchFmt = "array types mismatch:%s"

// ChanAddressMismatch sprint as:
//
//	'chan addresses mismatch:
//	 → Received: %s
//	 → Expected: %s'
func ChanAddressMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(ChanAddressMismatchFmt, recExp(tab, a, b))
}

const ChanAddressMismatchFmt = "chan addresses mismatch:%s>"

// ChanBufSizeMismatch sprint as:
//
//	'chan buffer sizes mismatch:
//	 → Received: %s
//	 → Expected: %s'
func ChanBufSizeMismatch(tab indent.Indentor, a, b int) string {
	return tab.Sprintf(ChanBufSizeMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const ChanBufSizeMismatchFmt = "chan buffer sizes mismatch:%s"

// ChanDirMismatch sprint as:
//
//	'chan direction mismatch:
//	 → Received: %s
//	 → Expected: %s'
func ChanDirMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(ChanDirMismatchFmt, recExp(tab, a, b))
}

const ChanDirMismatchFmt = "chan direction mismatch:%s"

// ChanElemTypesMismatch sprint as:
//
//	'chan element types mismatch:
//	 → Received: %s
//	 → Expected: %s'
func ChanElemTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(ChanElemTypesMismatchFmt, recExp(tab, a, b))
}

const ChanElemTypesMismatchFmt = "chan element types mismatch:%s"

// ChanTypesMismatch sprint as:
//
//	'chan types mismatch:
//	 → Received: %s
//	 → Expected: %s'
func ChanTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(ChanTypesMismatchFmt, recExp(tab, a, b))
}

const ChanTypesMismatchFmt = "chan types mismatch:%s"

// Empty sprints en indented empty string
func Empty(tab indent.Indentor) string {
	return tab.Sprint(EmptyFmt)
}

const EmptyFmt = ""

// EqualsCyclicRef sprints as:
//
// %s<cyclic-reference|%s>
func EqualsCyclicRef(tab indent.Indentor, typx string, addr uintptr) string {
	return tab.Sprint(EqualsCyclicRefFmt, typx, sprints.Uintptrf(addr))
}

const EqualsCyclicRefFmt = "%s<cyclic-reference|%s>"

// FuncAddressMismatch sprint as:
//
//	'func addresses mismatch:
//	 → Received: %s
//	 → Expected: %s'
func FuncAddressMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(FuncAddressMismatchFmt, recExp(tab, a, b))
}

const FuncAddressMismatchFmt = "func addresses mismatch:%s"

// FuncSignMismatch sprint as:
//
//	'func mismatch:
//	 → Received: %s
//	 → Expected: %s'
func FuncSignMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(FuncSignMismatchFmt, recExp(tab, a, b))
}

const FuncSignMismatchFmt = "func mismatch:%s"

// FuncTypesMismatch sprint as:
//
//	'func types mismatch:
//	 → Received: %s
//	 → Expected: %s'
func FuncTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(FuncTypesMismatchFmt, recExp(tab, a, b))
}

const FuncTypesMismatchFmt = "func types mismatch:%s"

// GenericAddrMismatch sprint as:
//
//	'%s addresses mismatch:
//	 → Received: %s
//	 → Expected: %s'
func GenericAddrMismatch(tab indent.Indentor, typ, a, b string) string {
	return tab.Sprintf(GenericAddrMismatchFmt, typ, recExp(tab, a, b))
}

const GenericAddrMismatchFmt = "%s addresses mismatch:%s"

// InterfaceTypesMismatch sprint as:
//
//	'interface types mismatch:
//	 → Received: %s
//	 → Expected: %s'
func InterfaceTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(InterfaceTypesMismatchFmt, recExp(tab, a, b))
}

const InterfaceTypesMismatchFmt = "interface types mismatch:%s"

// InterfaceImpl sprint as:
//
//	'implementations are different:'
func InterfaceImpl(tab indent.Indentor) string {
	return tab.Sprint(InterfaceImplFmt)
}

const InterfaceImplFmt = "implementations are different:"

// InvalidExpected
//
//	'expected value is invalid:
//	 → Received: %s
//	 → Expected: %s'
func InvalidExpected(tab indent.Indentor, a string) string {
	return tab.Sprintf(InvalidExpectedFmt, recExp(tab, a,
		sprints.Invalid(indent.Zero())))
}

const InvalidExpectedFmt = "expected value is invalid:%s"

// InvalidReceived
//
//	'received value is invalid:
//	 → Received: %s
//	 → Expected: %s'
func InvalidReceived(tab indent.Indentor, b string) string {
	return tab.Sprintf(InvalidReceivedFmt, recExp(tab, sprints.Invalid(indent.Zero()), b))
}

const InvalidReceivedFmt = "received value is invalid:%s"

// MapTypesMismatch
//
//	'map types mismatch:
//	 → Received: %s
//	 → Expected: %s'
func MapTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(MapTypesMismatchFmt, recExp(tab, a, b))
}

const MapTypesMismatchFmt = "map types mismatch:%s"

// MapLenMismatch
//
//	'map lengths mismatch:
//	 → Received: %s
//	 → Expected: %s'
func MapLenMismatch(tab indent.Indentor, a, b int) string {
	return tab.Sprintf(MapLenMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const MapLenMismatchFmt = "map lengths mismatch:%s"

// MapKeys
//
//	'map keys mismatch:'
func MapKeys(tab indent.Indentor) string {
	return tab.Sprint(MapKeysFmt)
}

const MapKeysFmt = "map keys mismatch:"

// MapKeysDiff
//
//	(received...)
//		missing keys:%v
//		extra keys:%v
//	(...)
//
//	(expected...)
//		missing keys:%v
//		extra keys:%v
//	(...)
func MapKeysDiff(tab indent.Tab, missingKeysA, extraKeysA, missingKeysB, extraKeysB []string) string {
	t := tab.String()
	result := t + "(received...):\n"

	if len(missingKeysA) > 0 {
		result += tab.Indent(1, "missing keys:", strings.Join(missingKeysA, ", "), "\n")
	}

	if len(extraKeysA) > 0 {
		result += tab.Indent(1, "extra keys:", strings.Join(extraKeysA, ", "), "\n")
	}

	result += t + "(...)\n\n" +
		t + "(expected...):\n"

	if len(missingKeysB) > 0 {
		result += tab.Indent(1, "missing keys:", strings.Join(missingKeysB, ", "), "\n")
	}

	if len(extraKeysB) > 0 {
		result += tab.Indent(1, "extra keys: ", strings.Join(extraKeysB, ", "), "\n")
	}

	result += t + "(...)"
	return result
}

// MapValue sprints as:
//
//	map values are different for key '%s':
func MapValue(tab indent.Indentor, key string) string {
	return tab.Sprintf(MapValueFmt, key)
}

const MapValueFmt = "map values are different for key '%s':"

// MissingField sprints as:
//
//	Field '%s' is missing
func MissingField(tab indent.Indentor, value string) string {
	return tab.Sprintf(MissingFieldFmt, value)
}

const MissingFieldFmt = "field '%s' is missing"

// NilChanDirMismatch sprints as:
//
//	'both channels are nil but their directions mismatch:
//	  → Received: %s
//	  → Expected: %s'
func NilChanDirMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(NilChanDirectionMismatchFmt, recExp(tab, a, b))
}

const NilChanDirectionMismatchFmt = "both channels are nil but their directions mismatch:%s"

// NilChanTypesMismatch sprints as:
//
//	'both channels are nil but their types mismatch:
//	  → Received: %s
//	  → Expected: %s'
func NilChanTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(NilChanTypesMismatchFmt, recExp(tab, a, b))
}

const NilChanTypesMismatchFmt = "both channels are nil but their types mismatch:%s"

// NilExpected sprint as:
//
//		'The expected '%s' is nil:
//	   → Received: %s
//		  → Expected: %s'
func NilExpected(tab indent.Indentor, kind, typ string) string {
	return tab.Sprintf(NilExpectedFmt, kind, recExp(tab, typ, "<nil>"))
}

const NilExpectedFmt = "The expected '%s' is nil:%s"

// NilFuncTypesMismatch sprint as:
//
//		'both functions are nil but their types mismatch:
//	   → Received: %s
//		  → Expected: %s'
func NilFuncTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(NilFuncTypesMismatchFmt, recExp(tab, a, b))
}

const NilFuncTypesMismatchFmt = "both functions are nil but their types mismatch:%s"

// NilFuncSignMismatch sprint as:
//
//		'both functions are nil but their signatures mismatch:
//	   → Received: %s
//		  → Expected: %s'
func NilFuncSignMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(NilFuncSignMismatchFmt, recExp(tab, a, b))
}

const NilFuncSignMismatchFmt = "both functions are nil but their signatures mismatch:%s"

// NilInterface sprint as:
//
//	The %s interface is nil
func NilInterface(tab indent.Indentor, id string) string {
	return tab.Sprintf(NilInterfaceFmt, id)
}

const NilInterfaceFmt = "The %s interface is nil"

// NilPointersSignMismatch sprints as:
//
//		'both pointers are nil but their signatures mismatch:
//	   → Received: %s
//		  → Expected: %s'
func NilPointersSignMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(NilPointersSignMismatchFmt, recExp(tab, a, b))
}

const NilPointersSignMismatchFmt = "both pointers are nil but their signatures mismatch:%s"

// NilReceived sprints as:
//
//	'The received '%s' is nil:
//	   → Received: %s
//	   → Expected: %s'
func NilReceived(tab indent.Indentor, kind, typ string) string {
	return tab.Sprintf(NilReceivedFmt, kind, recExp(tab, typ, "<nil>"))
}

const NilReceivedFmt = "The received '%s' is nil:%s"

// NilTypesMismatch sprints as:
//
//	'both %s are nil but their types mismatch:
//	   → Received: %s
//	   → Expected: %s'
func NilTypesMismatch(tab indent.Indentor, kind, a, b string) string {
	return tab.Sprintf(NilTypesMismatchFmt, kind, recExp(tab, a, b))
}

const NilTypesMismatchFmt = "both %s are nil but their types mismatch:%s"

// PointerValues sprint as:
//
//	'pointer values are different'
func PointerValues(tab indent.Indentor) string {
	return tab.Sprint(PointerValueFmt)
}

const PointerValueFmt = "pointer values are different"

func PointersAddrMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(PointersAddrMismatchFmt, recExp(tab, a, b))
}

const PointersAddrMismatchFmt = "pointer addresses mismatch:%s"

func Space(tab indent.Indentor, amount int) string {
	return tab.Sprint(strings.Repeat("␣", amount))
}

func SliceCapMismatch(tab indent.Indentor, a, b int) string {
	return tab.Sprintf(SliceCapMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const SliceCapMismatchFmt = "slice capacities mismatch:%s"

func SliceLenMismatch(tab indent.Indentor, a, b int) string {
	return tab.Sprintf(SliceLenMismatchFmt, recExp(tab, sprints.Digit(a), sprints.Digit(b)))
}

const SliceLenMismatchFmt = "slice lengths mismatch:%s"

func SliceTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(SliceTypesMismatchFmt, recExp(tab, a, b))
}

const SliceTypesMismatchFmt = "slice types mismatch:%s"

func SliceValues(tab indent.Indentor, idx int) string {
	return tab.Sprintf(SliceValuesFmt, idx)
}

const SliceValuesFmt = "slice values differ at index %d:"

func Strings(tab indent.Indentor, idx int) string {
	return tab.Sprintf(StringsFmt, idx)
}

const StringsFmt = "strings differ at index %d:"

// StructTypesMismatch sprint as:
//
//	'struct types mismatch:
//	   → Received: %s
//	   → Expected: %s'
func StructTypesMismatch(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(StructTypesMismatchFmt, recExp(tab, a, b))
}

const StructTypesMismatchFmt = "struct types mismatch:%s"

// StructFields sprint as:
//
//	'struct field '%s' mismatch:
func StructFields(tab indent.Indentor, field string) string {
	return tab.Sprintf(StructFieldsFmt, field)
}

const StructFieldsFmt = "struct field '%s' mismatch:"

// TypesMismatch returns a types mismatch Message
func TypesMismatch(tab indent.Indentor, a, b string) string {
	a = "<" + a + ">"
	b = "<" + b + ">"
	return tab.Sprintf(TypesMismatchFmt, recExp(tab, a, b))
}

const TypesMismatchFmt = "types mismatch:%s"

func UnsafePointers(tab indent.Indentor) string {
	return tab.Sprint(UnsafePointersFmt)
}

const UnsafePointersFmt = "unsafe ptrs values are different:"

// UnsafePointersAddr sprints as:
//
// unsafe pointer addresses mismatch:
//
//	→ Received: %s
//	→ Expected: %s'
func UnsafePointersAddr(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(UnsafePointersAddrFmt, recExp(tab, a, b))
}

const UnsafePointersAddrFmt = "unsafe pointer addresses mismatch:%s"

// Values sprint as:
//
//	values mismatch:%s
func Values(tab indent.Indentor, a, b string) string {
	return tab.Sprintf(ValuesFmt, recExp(tab, a, b))
}

const ValuesFmt = "values mismatch:%s"

func clean(s string) string {
	return strings.TrimRight(s, "\t\n ")
}

// recExp sprints as:
//
//	 "→ Received: %s"
//		"→ Expected: %s"
func recExp(tab indent.Indentor, rec, exp string) string {
	received := tab.Inc().Sprintf("→ Received: %s", rec)
	expected := tab.Inc().Sprintf("→ Expected: %s", exp)

	// Combine the results
	return printer.Sprintf("\n%s\n%s", received, expected)
}
