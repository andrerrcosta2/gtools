// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/domain/gerrors"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/pointers"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/sprint"
	"github.com/andrerrcosta2/gtools/reflect4/internal/standards"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/primitives"
	"reflect"
	"sort"
	"unsafe"
)

// Values returns the difference between two values
func Values(tab indent.Tab, value, expected reflect.Value) (diff string, equals bool, err error) {
	if msg, has, eq := invalids(value, expected); has {
		return tab.Smark(msg), eq, nil
	}

	if value.Kind() != expected.Kind() {
		return tab.Smark(differs.TypesMismatch(tab, types.Name(value.Type()), types.Name(expected.Type()))),
			false, err
	}

	if value.Kind() <= reflect.Complex128 {
		d := pr(tab, value, expected)
		if !d.Equals {
			return differs.Values(tab, d.Received, d.Expected), false, d.Err
		}
		return d.Message, d.Equals, d.Err
	}

	if value.Kind() == reflect.String {
		diff, equals = Text(tab, value.String(), expected.String(),
			standards.DifferStartIndex, standards.DifferMaxSize)
		return diff, equals, nil
	}

	var d differs.Difference
	switch value.Type().Kind() {
	case reflect.Array:
		d = ar(tab, value, expected, tracker.Diff())
	case reflect.Chan:
		d = ch(value, expected, tracker.Diff())
	case reflect.Func:
		d = fn(value, expected, tracker.Diff())
	case reflect.Interface:
		d = in(tab, value, expected, tracker.Diff())
	case reflect.Map:
		d = mp(tab, value, expected, tracker.Diff())
	case reflect.Ptr:
		d = pt(tab, value, expected, tracker.Diff())
	case reflect.Slice:
		d = sl(tab, value, expected, tracker.Diff())
	case reflect.Struct:
		d = st(tab, value, expected, tracker.Diff())
	case reflect.UnsafePointer:
		d = up(value, expected, tracker.Diff())
	default:
		d = differs.DiffError(fmt.Sprintf("unsupported kind: %s", value.Kind().String()),
			fmt.Errorf("unsupported kind: %s", value.Kind().String()))
	}
	if !d.Equals {
		return differs.Message(tab.Smark(d.Message), d.Diff), false, nil
	}
	return "", true, nil
}

func vl(tab indent.Tab, a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if msg, has, eq := invalids(a, b); has {
		return differs.Difference{
			Message: msg,
			Equals:  eq,
		}
	}

	if a.Kind() != b.Kind() {
		if msg, has, eq := invalids(a, b); has {
			return t.Mark(a, b, differs.Difference{
				Message: msg,
				Equals:  eq,
			})
		}
		return t.Mark(a, b, differs.NotEquals(differs.TypesMismatch(indent.Zero(), a.Kind().String(),
			b.Kind().String()), "", types.Name(a.Type()), types.Name(b.Type())))
	}

	if a.Kind() <= reflect.Complex128 || a.Kind() == reflect.String {
		return pr(tab, a, b)
	}

	switch a.Type().Kind() {
	case reflect.Array:
		return ar(tab, a, b, t)
	case reflect.Chan:
		return ch(a, b, t)
	case reflect.Func:
		return fn(a, b, t)
	case reflect.Interface:
		return in(tab, a, b, t)
	case reflect.Map:
		return mp(tab, a, b, t)
	case reflect.Ptr:
		return pt(tab, a, b, t)
	case reflect.Slice:
		return sl(tab, a, b, t)
	case reflect.Struct:
		return st(tab, a, b, t)
	case reflect.UnsafePointer:
		return up(a, b, t)
	default:
		return differs.DiffError(fmt.Sprintf("unsupported kind: %s", a.Kind().String()),
			fmt.Errorf("unsupported kind: %s", a.Kind().String()))
	}
}

func pr(tab indent.Tab, a, b reflect.Value) differs.Difference {
	name, _ := primitives.Name(a.Type())
	switch a.Kind() {
	case reflect.Bool:
		if a.Bool() != b.Bool() {
			return differs.NotEquals(differs.Values(indent.Zero(), sprints.Typed(name, a.Bool()),
				sprints.Typed(name, b.Bool())), "", indent.Zero().Sprint(sprints.Typed(name, a.Bool())),
				indent.Zero().Sprint(sprints.Typed(name, b.Bool())))
		}
		return differs.Equals("", sprints.Typed(name, a.Bool()), sprints.Typed(name, b.Bool()))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if a.Int() != b.Int() {
			message := differs.Values(indent.Zero(), sprints.Typed(name, a.Int()), sprints.Typed(name, b.Int()))
			return differs.NotEquals(message,
				"", indent.Zero().Sprint(sprints.TypedDigit(name, a.Int())),
				indent.Zero().Sprint(sprints.TypedDigit(name, b.Int())))
		}
		return differs.Equals("", tab.Sprint(sprints.TypedDigit(name, a.Int())),
			indent.Zero().Sprint(sprints.TypedDigit(name, b.Int())))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if a.Uint() != b.Uint() {
			return differs.NotEquals(differs.Values(indent.Zero(), sprints.Typed(name, a.Uint()),
				sprints.Typed(name, b.Uint())), "", indent.Zero().Sprint(sprints.TypedDigit(name, a.Uint())),
				indent.Zero().Sprint(sprints.TypedDigit(name, b.Uint())))
		}
		return differs.Equals("", indent.Zero().Sprint(sprints.TypedDigit(name, a.Uint())),
			indent.Zero().Sprint(sprints.TypedDigit(name, b.Uint())))

	case reflect.Float32, reflect.Float64:
		if a.Float() != b.Float() {
			return differs.NotEquals(differs.Values(indent.Zero(), sprints.Typed(name, a.Float()),
				sprints.Typed(name, b.Float())), "", indent.Zero().Sprint(sprints.TypedRoundFloat(name, a.Float())),
				indent.Zero().Sprint(sprints.TypedRoundFloat(name, b.Float())))
		}
		return differs.Equals("", indent.Zero().Sprint(sprints.Typed(name, a.Float())),
			indent.Zero().Sprint(sprints.Typed(name, b.Float())))

	case reflect.Complex64, reflect.Complex128:
		if a.Complex() != b.Complex() {
			fmt.Printf("%s: %v != %v\n", name, a.Complex(), b.Complex())
			return differs.NotEquals(differs.Values(indent.Zero(), sprints.Typed(name, a.Complex()),
				sprints.Typed(name, b.Complex())), "", indent.Zero().Sprint(sprints.Typed(name, a.Complex())),
				indent.Zero().Sprint(sprints.Typed(name, b.Complex())))
		}
		return differs.Equals("", indent.Zero().Sprint(sprints.Typed(name, a.Complex())),
			indent.Zero().Sprint(sprints.Typed(name, b.Complex())))

	case reflect.String:
		return sg(tab, a.String(), b.String())
	default:
		return differs.DiffError("no default values for primitives", reflect4.ErrInvalidValue)
	}
}

func ar(tab indent.Tab, a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	// Golang is weird. It considers two arrays of different sizes as not having the same type...
	if a.Type().Elem() != b.Type().Elem() {
		return t.Mark(a, b, differs.NotEquals(differs.ArrayTypesMismatch(indent.Zero(), a.Type().String(),
			b.Type().String()), "", a.Type().String(), b.Type().String()))
	}
	if a.Len() != b.Len() {
		aa, err := sprint.Array(tab, a)
		if err != nil {
			aa = sprints.Error(indent.Zero(), err.Error())
		}
		bb, err := sprint.Array(tab, b)
		if err != nil {
			bb = sprints.Error(indent.Zero(), err.Error())
		}
		return t.Mark(a, b, differs.NotEquals(differs.ArrayLenMismatch(indent.Zero(), a.Len(), b.Len()),
			"", aa, bb))
	}

	for i := 0; i < a.Len(); i++ {
		elemA := a.Index(i)
		elemB := b.Index(i)

		diff := vl(tab.Inc(), elemA, elemB, t)
		if !diff.Equals {
			message := differs.Append(indent.Zero(), differs.ArrayElem(indent.Zero(), i), diff.Message)
			return t.Mark(a, b, differs.NotEquals(message, diff.Diff, diff.Received,
				diff.Expected),
			)
		}
	}
	return t.Mark(a, b, differs.Equals("", a.String(), b.String()))
}

func ch(a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	msg, has, equals := nullables(a, b)
	if has {
		if equals {
			return t.Mark(a, b, differs.Equals(msg, a.Type().String(), b.Type().String()))
		}
		return t.Mark(a, b, differs.NotEquals(msg, "", a.Type().String(), b.Type().String()))
	}

	if a.Type() != b.Type() {
		return t.Mark(a, b, differs.NotEquals(differs.ChanTypesMismatch(indent.Zero(), a.Type().String(), b.Type().String()),
			"", a.Type().String(), b.Type().String()))
	}

	if a.Type().ChanDir() != b.Type().ChanDir() {
		return t.Mark(a, b, differs.NotEquals(differs.ChanDirMismatch(indent.Zero(),
			a.Type().String(), b.Type().String()), "", a.Type().String(), b.Type().String()))
	}

	bufSizeA := a.Cap()
	bufSizeB := b.Cap()

	if bufSizeA != bufSizeB {
		return t.Mark(a, b, differs.NotEquals(differs.ChanBufSizeMismatch(indent.Zero(), bufSizeA, bufSizeB),
			"", fmt.Sprintf("Buffer size: %d", bufSizeA), fmt.Sprintf("Buffer size: %d", bufSizeB)))
	}

	return t.Mark(a, b, differs.Equals("", a.String(), b.String()))
}

func fn(a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	msg, has, equals := nullables(a, b)
	if has {
		if equals {
			return t.Mark(a, b, differs.Equals(msg, a.Type().String(), b.Type().String()))
		}
		return t.Mark(a, b, differs.NotEquals(msg, "", a.Type().String(), b.Type().String()))
	}

	if a.Type() != b.Type() {
		aa, err := sprint.Func(indent.Zero(), a)
		if err != nil {
			aa = sprints.Error(indent.Zero(), err.Error())
		}
		bb, err := sprint.Func(indent.Zero(), b)
		if err != nil {
			bb = sprints.Error(indent.Zero(), err.Error())
		}
		return t.Mark(a, b, differs.NotEquals(differs.FuncTypesMismatch(indent.Zero(), aa,
			bb), "", aa, bb))
	}

	if a.String() != b.String() {
		return t.Mark(a, b, differs.NotEquals(differs.FuncSignMismatch(indent.Zero(),
			a.Type().String(), b.Type().String()), "", a.Type().String(), b.Type().String()))
	}

	return t.Mark(a, b, differs.Equals("", a.String(), b.String()))
}

func in(tab indent.Tab, a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	msg, has, equals := nullables(a, b)
	if has {
		if equals {
			return t.Mark(a, b, differs.Equals(msg, a.Type().String(), b.Type().String()))
		}
		return t.Mark(a, b, differs.NotEquals(msg, "", a.Type().String(), b.Type().String()))
	}

	if a.Type() != b.Type() {
		return t.Mark(a, b, differs.NotEquals(differs.InterfaceTypesMismatch(indent.Zero(), a.Type().String(),
			b.Type().String()), "", a.Type().String(), b.Type().String()))
	}

	differ := vl(tab.Inc(), a.Elem(), b.Elem(), t)
	if !differ.Equals {
		message := differs.Append(indent.Zero(), differs.InterfaceImpl(indent.Zero()), differ.Message)
		//fmx.Redf("\nBefore Differ: Tab: %d\n", tab)
		return t.Mark(a, b, differs.NotEquals(message, differ.Diff, differ.Received, differ.Expected))
	}

	return t.Mark(a, b, differs.Equals("", a.String(), b.String()))
}

func mp(tab indent.Tab, a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	msg, has, equals := nullables(a, b)
	if has {
		if equals {
			return t.Mark(a, b, differs.Equals(msg, a.Type().String(), b.Type().String()))
		}
		return t.Mark(a, b, differs.NotEquals(msg, "", a.Type().String(), b.Type().String()))
	}

	if a.Type() != b.Type() {
		return t.Mark(a, b, differs.NotEquals(differs.MapTypesMismatch(tab, a.Type().String(), b.Type().String()),
			"", a.Type().String(), b.Type().String()))
	}

	keysA := a.MapKeys()
	keysB := b.MapKeys()

	setA := make(map[interface{}]struct{})
	setB := make(map[interface{}]struct{})

	for _, key := range keysA {
		setA[key.Interface()] = struct{}{}
	}
	for _, key := range keysB {
		setB[key.Interface()] = struct{}{}
	}

	var missingKeysA, extraKeysA, missingKeysB, extraKeysB []string
	for key := range setA {
		if _, exists := setB[key]; !exists {
			extraKeysA = append(extraKeysA, fmt.Sprintf("%v", key))
			missingKeysB = append(missingKeysB, fmt.Sprintf("%v", key))
		}
	}
	for key := range setB {
		if _, exists := setA[key]; !exists {
			extraKeysB = append(extraKeysB, fmt.Sprintf("%v", key))
			missingKeysA = append(missingKeysA, fmt.Sprintf("%v", key))
		}
	}

	if len(missingKeysA) > 0 || len(extraKeysA) > 0 || len(missingKeysB) > 0 || len(extraKeysB) > 0 {
		diff := differs.MapKeysDiff(tab, missingKeysA, extraKeysA, missingKeysB, extraKeysB)
		return t.Mark(a, b, differs.NotEquals(differs.MapKeys(indent.Zero()), diff, "", ""))
	}

	// 🔧 Sort keys alphabetically before comparison
	sort.Slice(keysA, func(i, j int) bool {
		return fmt.Sprint(keysA[i].Interface()) < fmt.Sprint(keysA[j].Interface())
	})

	// Now compare values in sorted order
	for _, key := range keysA {
		valA := a.MapIndex(key)
		valB := b.MapIndex(key)

		diff := vl(tab.Inc(), valA, valB, t)
		if !diff.Equals {
			message := differs.Append(indent.Zero(), differs.MapValue(indent.Zero(), key.String()), diff.Message)
			return t.Mark(a, b, differs.NotEquals(message, diff.Diff, diff.Received, diff.Expected))
		}
	}

	return t.Mark(a, b, differs.Equals("", a.String(), b.String()))
}

func pt(tab indent.Tab, a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	msg, has, equals := nullables(a, b)
	if has {
		if equals {
			return t.Mark(a, b, differs.Equals(msg, a.Type().String(), b.Type().String()))
		}
		return t.Mark(a, b, differs.NotEquals(msg, "", a.Type().String(), b.Type().String()))
	}

	differ := vl(tab.Inc(), a.Elem(), b.Elem(), t)
	if !differ.Equals {
		message := differs.Append(indent.Zero(), differs.Pointers(indent.Zero()), differ.Message)
		//fmx.Redf("\nBefore Differ: Tab: %d\n", tab)
		return t.Mark(a, b, differs.NotEquals(message, differ.Diff, differ.Received, differ.Expected))
	}
	return t.Mark(a, b, differs.Equals("", a.String(), b.String()))
}

func sg(tab indent.Tab, value, expected string) differs.Difference {
	received, expected, idx := difference(value, expected, standards.DifferStartIndex, standards.DifferMaxSize)
	if idx == -1 {
		return differs.Equals("", value, expected)
	}

	startReceived := max(0, idx-standards.DifferStartIndex)
	endReceived := min(len(received), idx+standards.DifferMaxSize)

	// Ensure valid bounds for the received string
	if startReceived > len(received) {
		startReceived = len(received)
	}
	if endReceived < startReceived {
		endReceived = startReceived
	}

	// Calculate the start and end indices for the expected string
	startExpected := max(0, idx-standards.DifferStartIndex)
	endExpected := min(len(expected), idx+standards.DifferMaxSize)

	// Ensure valid bounds for the expected string
	if startExpected > len(expected) {
		startExpected = len(expected)
	}
	if endExpected < startExpected {
		endExpected = startExpected
	}

	receivedSubstring := received[startReceived:endReceived]
	expectedSubstring := expected[startExpected:endExpected]

	diff := differs.Diff(tab, receivedSubstring, expectedSubstring)
	// leave the diff to be resolved on the higher level
	return differs.NotEquals(differs.Strings(indent.Zero(), idx), diff, receivedSubstring, expectedSubstring)
}

func sl(tab indent.Tab, a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	msg, has, equals := nullables(a, b)
	if has {
		if equals {
			return t.Mark(a, b, differs.Equals(msg, a.Type().String(), b.Type().String()))
		}
		return t.Mark(a, b, differs.NotEquals(msg, "", a.Type().String(), b.Type().String()))
	}

	if a.Type() != b.Type() {
		aa := types.Name(a.Type())
		bb := types.Name(b.Type())
		return t.Mark(a, b, differs.NotEquals(differs.SliceTypesMismatch(indent.Zero(), aa, bb), "", aa, bb))
	}

	if a.Len() != b.Len() {
		aa := types.Name(a.Type())
		bb := types.Name(b.Type())
		return t.Mark(a, b, differs.NotEquals(differs.SliceLenMismatch(indent.Zero(), a.Len(), b.Len()), "", aa, bb))
	}

	if a.Cap() != b.Cap() {
		aa := types.Name(a.Type())
		bb := types.Name(b.Type())
		return t.Mark(a, b, differs.NotEquals(differs.SliceCapMismatch(indent.Zero(), a.Cap(), b.Cap()),
			"", aa, bb))
	}

	for i := 0; i < a.Len(); i++ {
		valA := a.Index(i)
		valB := b.Index(i)
		differ := vl(tab.Inc(), valA, valB, t)
		if !differ.Equals {
			message := differs.Append(indent.Zero(), differs.SliceValues(indent.Zero(), i), differ.Message)
			return t.Mark(a, b, differs.NotEquals(message, differ.Diff, differ.Received, differ.Expected))
		}
	}

	return t.Mark(a, b, differs.Equals("", a.String(), b.String()))
}

func st(tab indent.Tab, a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	// Structs can't be nil
	if a.Type() != b.Type() {
		return t.Mark(a, b, differs.NotEquals(differs.StructTypesMismatch(indent.Zero(),
			types.Name(a.Type()), types.Name(b.Type())), "", types.Name(a.Type()), types.Name(b.Type())))
	}

	for i := 0; i < a.NumField(); i++ {
		field := a.Type().Field(i)
		valA := a.Field(i)
		valB := b.Field(i)
		differ := vl(tab.Inc(), valA, valB, t)
		if !differ.Equals {
			message := differs.Append(indent.Zero(), differs.StructFields(indent.Zero(), field.Name), differ.Message)
			return t.Mark(a, b, differs.NotEquals(message, differ.Diff, differ.Received, differ.Expected))
		}
	}

	return t.Mark(a, b, differs.Equals("", a.String(), b.String()))
}

func up(a, b reflect.Value, t *tracker.DiffTracker) differs.Difference {
	if diff, ok := flagCycl(a, b, t); ok {
		return diff
	}

	// Handle nil values
	msg, has, equals := nullables(a, b)
	if has {
		if equals {
			return t.Mark(a, b, differs.Equals(msg, a.Type().String(), b.Type().String()))
		}
		return t.Mark(a, b, differs.NotEquals(msg, "", a.Type().String(), b.Type().String()))
	}

	// Get the raw memory addresses
	ptrA := unsafe.Pointer(a.Pointer())
	ptrB := unsafe.Pointer(b.Pointer())

	// Compare raw addresses
	if ptrA == ptrB {
		return t.Mark(a, b, differs.Equals("", fmt.Sprintf("0x%x", ptrA), fmt.Sprintf("0x%x", ptrB)))
	}

	// Report address mismatch
	return t.Mark(a, b, differs.NotEquals(differs.UnsafePointersAddr(indent.Zero(),
		sprints.UnsafeAddrf(ptrA), sprints.UnsafeAddrf(ptrB)),
		"", fmt.Sprintf("0x%x", ptrA), fmt.Sprintf("0x%x", ptrB)))
}

func flagCycl(a, b reflect.Value, t *tracker.DiffTracker) (differs.Difference, bool) {
	if diff, ok := t.Get(a, b); ok { // Avoiding cycles
		return diff, true
	}
	addrA, errA := pointers.Of(a)
	addrB, errB := pointers.Of(b)
	if errA != nil || errB != nil {
		return differs.DiffError("invalid pointer(s)", gerrors.StackOf(errA, errB)), true
	}
	return t.Mark(a, b, differs.Equals("cyclic reference",
		sprints.Uintptrf(addrA), sprints.Uintptrf(addrB))), false
}
