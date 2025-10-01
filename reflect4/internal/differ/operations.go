// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"reflect"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/standards"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/maps"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
)

func between(tab indent.Indentor, a, b reflect.Value, diff *Strategy) differs.Difference {
	if diff, invalid := handleInvalidValues(tab, a, b); invalid {
		return diff
	}

	// we don't need to mark NotEquals results since it will cut the flow immediately,
	// no traversing can happen
	if a.Kind() != b.Kind() {
		msg := differs.TypesMismatch(tab, a.Kind().String(), b.Kind().String())
		return differs.NotEquals(msg, "", types.Name(a.Type()), types.Name(b.Type()))
	}

	if a.Kind() <= reflect.Complex128 {
		return differPrimitives(tab, a, b)
	}

	switch a.Type().Kind() {
	case reflect.Array:
		return diffArrays(tab, a, b, diff)
	case reflect.Chan:
		return diff.channels(tab, a, b, diff)
	case reflect.Func:
		return diff.functions(tab, a, b, diff)
	case reflect.Interface:
		return differInterfaces(tab, a, b, diff)
	case reflect.Map:
		return diff.maps(tab, a, b, diff)
	case reflect.Ptr:
		return diff.pointers(tab, a, b, diff)
	case reflect.Slice:
		return diff.slices(tab, a, b, diff)
	case reflect.Struct:
		return diffStructs(tab, a, b, diff)
	case reflect.String:
		return diffStrings(tab, a.String(), b.String(), diff)
	case reflect.UnsafePointer:
		return differUnsafePointers(tab, a, b, diff)
	default:
		return differs.DiffError(tab.Sprintf("unsupported kind: '%s' diff", a.Kind().String()),
			fmx.Errorf("unsupported kind: '%s' diff", a.Kind().String()))
	}
}

// defaultChanDiff computes the defaultStringDiff between two channels
//
//	It considers two channels as equals if they:
//	- have the same signature and direction
//	- if they have the same buffer capacity
func defaultChanDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[differChans] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))

	if a.IsNil() && b.IsNil() {
		return nilChanSignatureDiff(tab, a, b)
	}

	if !a.IsNil() && !b.IsNil() {
		if diff, ok := s.check(a, b); ok {
			return diff
		}
		if a.Pointer() == b.Pointer() {
			return differs.Equals()
		}
		return s.cache(a, b, defaultChanSignatureDiff(tab, a, b))
	}

	if a.IsNil() {
		name := types.Name(b.Type())
		msg := differs.NilReceived(tab, "chan", name)
		return differs.NotEquals(msg, "", name, name)
	}

	name := types.Name(a.Type())
	msg := differs.NilExpected(tab, "chan", name)
	return differs.NotEquals(msg, "", name, name)
}

// defaultChanSignatureDiff this function compares two channel signatures
// of relevant differences
func defaultChanSignatureDiff(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	ta, tb := a.Type(), b.Type()

	if ta.Elem() != tb.Elem() {
		an, bn := types.ValidValueName(ta.Elem()), types.ValidValueName(tb.Elem())
		msg := differs.ChanElemTypesMismatch(tab, an, bn)
		return differs.NotEquals(msg, "", an, bn)
	}

	if ta.ChanDir() != tb.ChanDir() {
		msg := differs.ChanDirMismatch(tab, ta.ChanDir().String(), tb.ChanDir().String())
		return differs.NotEquals(msg, "", ta.ChanDir().String(), tb.ChanDir().String())
	}

	capA := a.Cap()
	capB := b.Cap()

	if capA != capB {
		msg := differs.ChanBufSizeMismatch(tab, capA, capB)
		return differs.NotEquals(msg, "",
			fmx.Sprintf("Buffer size: %d", capA), fmx.Sprintf("Buffer size: %d", capB))
	}

	// All channel signatures match
	return differs.Equals()
}

func nilChanSignatureDiff(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	// Nil channels capacity is always zero
	ta, tb := a.Type(), b.Type()

	if ta.Elem() != tb.Elem() {
		an, bn := types.ValidValueName(ta.Elem()), types.ValidValueName(tb.Elem())
		msg := differs.NilChanTypesMismatch(tab, an, bn)
		return differs.NotEquals(msg, "", an, bn)
	}

	if ta.ChanDir() != tb.ChanDir() {
		msg := differs.NilChanDirMismatch(tab, ta.ChanDir().String(), tb.ChanDir().String())
		return differs.NotEquals(msg, "", ta.ChanDir().String(), tb.ChanDir().String())
	}
	return differs.Equals()
}

func sameChanDiff(tab indent.Indentor, a, b reflect.Value, _ *Strategy) differs.Difference {
	//fmx.Purplef("[differChans] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))

	if !a.IsNil() && !b.IsNil() {
		if a.Pointer() == b.Pointer() {
			return differs.Equals()
		}
		name := types.ValidValueName(a.Type())
		msg := differs.ChanAddressMismatch(tab,
			sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer()))
		return differs.NotEquals(msg, "", name, name)
	}

	if a.IsNil() && b.IsNil() {
		return nilChanSignatureDiff(tab, a, b)
	}

	if a.IsNil() {
		msg := differs.NilReceived(tab, "chan", b.Type().String())
		return differs.NotEquals(msg, "", a.Type().String(), b.Type().String())
	}

	msg := differs.NilExpected(tab, "chan", a.Type().String())
	return differs.NotEquals(msg, "", a.Type().String(), b.Type().String())
}

// defaultFuncDiff computes the defaultStringDiff between two functions.
//
//	it considers as compare functions if they:
//	- both addresses are the same; or
//	- have the same signature - even if the implementations are different
func defaultFuncDiff(tab indent.Indentor, a, b reflect.Value, _ *Strategy) differs.Difference {
	//fmx.Purplef("[differFunctions] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))

	if !a.IsNil() && !b.IsNil() {
		return funcSignatureDiff(tab, a, b)
	}

	if a.IsNil() && b.IsNil() {
		return nilFuncSignatureDiff(tab, a, b)
	}

	if a.IsNil() {
		bb := types.ValidValueName(b.Type())
		msg := differs.NilReceived(tab, "func", bb)
		return differs.NotEquals(msg, "", a.Type().String(), b.Type().String())
	}

	aa := types.ValidValueName(a.Type())
	msg := differs.NilExpected(tab, "func", aa)
	return differs.NotEquals(msg, "", a.Type().String(), b.Type().String())
}

func funcSignatureDiff(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	ta, tb := a.Type(), b.Type()
	aa := types.ValidValueName(ta)
	bb := types.ValidValueName(tb)
	if ta != tb {
		msg := differs.FuncTypesMismatch(tab, aa, bb)
		return differs.NotEquals(msg, "", aa, bb)
	}
	return differs.Equals()
}

func nilFuncSignatureDiff(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	ta, tb := a.Type(), b.Type()
	aa := types.ValidValueName(ta)
	bb := types.ValidValueName(tb)
	if ta != tb {
		msg := differs.NilFuncTypesMismatch(tab, aa, bb)
		return differs.NotEquals(msg, "", aa, bb)
	}
	return differs.Equals()
}

// defaultFuncDiff computes the defaultStringDiff between two functions.
//
//	it considers as compare functions if they:
//	- have the same type
//	- have the same signature - even if the implementations are different
func sameFuncDiff(tab indent.Indentor, a, b reflect.Value, _ *Strategy) differs.Difference {
	//fmx.Purplef("[differFunctions] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))

	if !a.IsNil() && !b.IsNil() {
		if a.Pointer() != b.Pointer() {
			msg := differs.FuncAddressMismatch(tab,
				sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer()))
			return differs.NotEquals(msg, "", a.Type().String(), b.Type().String())
		}
		return differs.Equals()
	}

	if a.IsNil() && b.IsNil() {
		return nilFuncSignatureDiff(tab, a, b)
	}

	if a.IsNil() {
		bb := types.ValidValueName(b.Type())
		msg := differs.NilReceived(tab, "func", bb)
		return differs.NotEquals(msg, "", bb, bb)
	}

	aa := types.ValidValueName(a.Type())
	msg := differs.NilExpected(tab, "func", aa)
	return differs.NotEquals(msg, "", aa, aa)
}

// emptyOrNullableIter reports whether two values
func emptyOrNullableIter(tab indent.Indentor, a, b reflect.Value) (diff differs.Difference, done bool) {
	if a.IsNil() || b.IsNil() {
		aa := types.Name(a.Type())
		bb := types.Name(b.Type())
		if !a.IsNil() {
			if a.Len() == 0 {
				return differs.Equals(), true
			}
			msg := differs.NilExpected(tab, b.Kind().String(), bb)
			return differs.NotEquals(msg, "", a.Type().String(), b.Type().String()), true
		}
		if !b.IsNil() {
			if b.Len() == 0 {
				return differs.Equals(), true
			}
			msg := differs.NilReceived(tab, a.Kind().String(), aa)
			return differs.NotEquals(msg, "", a.Type().String(), b.Type().String()), true
		}
		return differs.Equals(), true
	}
	return differs.Difference{}, false
}

func diffArrays(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[diffArrays] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))
	ta, tb := a.Type(), b.Type()
	aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)

	if ta.Elem() != tb.Elem() {
		msg := differs.ArrayTypesMismatch(tab, aa, bb)
		return differs.NotEquals(msg, "", aa, bb)
	}

	if a.Len() != b.Len() {
		// fmx.Purplef("[diffArrays] '%s'.Len() is different from '%s'.Len()\n", a.String(), b.String())

		msg := differs.ArrayLenMismatch(tab, a.Len(), b.Len())
		return differs.NotEquals(msg, "", aa, bb)
	}

	for i := 0; i < a.Len(); i++ {
		elemA := a.Index(i)
		elemB := b.Index(i)
		//fmx.Purplef("[diffArrays] '%s'.Index(%d) compare '%s'.Index(%d): %t\n", a.String(), i, b.String(), i,
		//	reflect.DeepEqual(a.Interface(), b.Interface()))

		diff := between(tab.Inc(), elemA, elemB, s)
		if !diff.Equals {
			msg := differs.ArrayElem(tab, i) + "\n" + diff.Message
			return differs.NotEquals(msg, diff.Diff, diff.Received, diff.Expected)
		}
	}
	return differs.Equals()
}

func differBool(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	name := types.ValidValueName(a.Type())
	if a.Bool() != b.Bool() {
		return differs.NotEquals(differs.Values(tab, sprints.Typed(name, a.Bool()),
			sprints.Typed(name, b.Bool())), "", tab.Sprint(sprints.Typed(name, a.Bool())),
			tab.Sprint(sprints.Typed(name, b.Bool())))
	}
	return differs.Equals()
}

func differComplex(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	name := types.ValidValueName(a.Type())
	if a.Complex() != b.Complex() {
		fmx.Printf("%s: %v != %v\n", name, a.Complex(), b.Complex())
		return differs.NotEquals(differs.Values(tab, sprints.Typed(name, a.Complex()),
			sprints.Typed(name, b.Complex())), "", tab.Sprint(sprints.Typed(name, a.Complex())),
			tab.Sprint(sprints.Typed(name, b.Complex())))
	}
	return differs.Equals()
}

func differFloat(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	name := types.ValidValueName(a.Type())
	if a.Float() != b.Float() {
		return differs.NotEquals(differs.Values(tab, sprints.Typed(name, a.Float()),
			sprints.Typed(name, b.Float())), "", tab.Sprint(sprints.TypedRoundFloat(name, a.Float())),
			tab.Sprint(sprints.TypedRoundFloat(name, b.Float())))
	}
	return differs.Equals()
}

func differInt(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	name := types.ValidValueName(a.Type())
	if a.Int() != b.Int() {
		msg := differs.Values(tab, sprints.Typed(name, a.Int()), sprints.Typed(name, b.Int()))
		return differs.NotEquals(msg,
			"", tab.Sprint(sprints.TypedDigit(name, a.Int())),
			tab.Sprint(sprints.TypedDigit(name, b.Int())))
	}
	return differs.Equals()
}

// differInterfaces interfaces are equals if:
//
//	A: both interface types are equals; and
//	A: both implementations are equals; or
//	B: both interfaces are nil; and
//	B: both types are equals
func differInterfaces(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[differInterfaces] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))

	ta, tb := a.Type(), b.Type()
	aa, bb := types.ValidValueName(a.Type()), types.ValidValueName(b.Type())

	if !a.IsNil() && !b.IsNil() {
		if ta != tb {
			msg := differs.InterfaceTypesMismatch(tab, aa, bb)
			return differs.NotEquals(msg, "", aa, bb)
		}

		diff := between(tab.Inc(), a.Elem(), b.Elem(), s)
		if !diff.Equals {
			msg := differs.InterfaceImpl(tab) + "\n" + diff.Message
			return differs.NotEquals(msg, diff.Diff, diff.Received, diff.Expected)
		}
		return differs.Equals()
	}

	if a.IsNil() && b.IsNil() {
		if ta != tb {
			msg := differs.NilTypesMismatch(tab, "interfaces", aa, bb)
			return differs.NotEquals(msg, "", aa, bb)
		}
		return differs.Equals()
	}

	if a.IsNil() {
		msg := differs.NilReceived(tab, "interface", bb)
		return differs.NotEquals(msg, "", aa, bb)
	}
	msg := differs.NilExpected(tab, "interface", aa)
	return differs.NotEquals(msg, "", aa, bb)
}

func differPrimitives(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	switch a.Kind() {
	case reflect.Bool:
		return differBool(tab, a, b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return differInt(tab, a, b)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return differUint(tab, a, b)
	case reflect.Float32, reflect.Float64:
		return differFloat(tab, a, b)
	case reflect.Complex64, reflect.Complex128:
		return differComplex(tab, a, b)
	default:
		return differs.DiffError("no default values for primitives", reflect4.ErrInvalidValue)
	}
}

// defaultMapDiff evaluates the defaultStringDiff between two maps given the following constraints:
//
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals; or
//	C: both types are nil: and
//	C: both types have the same signature
func defaultMapDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[defaultMapDiff] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))

	if !a.IsNil() && !b.IsNil() {
		return notNilMapDiffer(tab, a, b, s)
	}

	aa, bb := types.ValidValueName(a.Type()), types.ValidValueName(b.Type())
	if a.IsNil() && b.IsNil() {
		if a.Type() != b.Type() {
			msg := differs.NilTypesMismatch(tab, "maps", aa, bb)
			return differs.NotEquals(msg, "", aa, bb)
		}
		return differs.Equals()
	}

	if a.IsNil() {
		msg := differs.NilReceived(tab, "map", bb)
		return differs.NotEquals(msg, "", aa, bb)
	}
	msg := differs.NilExpected(tab, "map", aa)
	return differs.NotEquals(msg, "", aa, bb)
}

// serializableMapDiff evaluates the defaultStringDiff between two maps given the following constraints:
//
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals; or
//	C: both types are nil or empty: and
//	C: both types have the same signature
func serializableMapDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[defaultMapDiff] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))

	if !a.IsNil() && !b.IsNil() {
		return notNilMapDiffer(tab, a, b, s)
	}

	aa, bb := types.ValidValueName(a.Type()), types.ValidValueName(b.Type())
	if (a.IsNil() && b.IsNil()) ||
		(a.IsNil() && b.Len() == 0) ||
		(b.IsNil() && a.Len() == 0) {
		if a.Type() != b.Type() {
			msg := differs.NilTypesMismatch(tab, "maps", aa, bb)
			return differs.NotEquals(msg, "", aa, bb)
		}
		return differs.Equals()
	}

	if a.IsNil() {
		msg := differs.NilReceived(tab, "map", bb)
		return differs.NotEquals(msg, "", aa, bb)
	}
	msg := differs.NilExpected(tab, "map", aa)
	return differs.NotEquals(msg, "", aa, bb)
}

// notNilMapDiffer evaluates the defaultStringDiff between two maps given the following constraints:
//
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals;
func notNilMapDiffer(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	if diff, ok := s.check(a, b); ok {
		return diff
	}

	aa, bb := types.ValidValueName(a.Type()), types.ValidValueName(b.Type())
	if a.Pointer() == b.Pointer() {
		return differs.Equals()
	}

	if a.Type() != b.Type() {
		msg := differs.MapTypesMismatch(tab, aa, bb)
		return differs.NotEquals(msg, "", aa, bb)
	}

	if a.Len() != b.Len() {
		msg := differs.MapLenMismatch(tab, a.Len(), b.Len())
		diff, _ := mapLengthDiffer(tab, a, b)
		return s.cache(a, b, differs.NotEquals(msg, diff, aa, bb))
	}

	var missingKeysA, extraKeysA, missingKeysB, extraKeysB []string

	iterA := a.MapRange()
	for iterA.Next() {
		key := iterA.Key()
		valA := iterA.Value()
		valB := b.MapIndex(key)

		if !valB.IsValid() {
			missingKeysB = append(missingKeysB, fmx.Sprintf("%v", key.Interface()))
			extraKeysA = append(extraKeysA, fmx.Sprintf("%v", key.Interface()))
			continue
		}

		diff := between(tab.Inc(), valA, valB, s)
		if !diff.Equals {
			message := differs.MapValue(tab, fmx.Sprint(key.Interface())) + "\n" + diff.Message
			return s.cache(a, b, differs.NotEquals(message, diff.Diff, diff.Received, diff.Expected))
		}
	}

	// Check for keys in b but missing in a
	iterB := b.MapRange()
	for iterB.Next() {
		key := iterB.Key()
		if !a.MapIndex(key).IsValid() {
			missingKeysA = append(missingKeysA, fmx.Sprintf("%v", key.Interface()))
			extraKeysB = append(extraKeysB, fmx.Sprintf("%v", key.Interface()))
		}
	}

	if len(missingKeysA) > 0 || len(extraKeysA) > 0 || len(missingKeysB) > 0 || len(extraKeysB) > 0 {
		diff := differs.MapKeysDiff(indent.Tab(0), missingKeysA, extraKeysA, missingKeysB, extraKeysB)
		return s.cache(a, b, differs.NotEquals(differs.MapKeys(tab), diff, aa, bb))
	}

	return s.cache(a, b, differs.Equals())
}

func mapLengthDiffer(tab indent.Indentor, a, b reflect.Value) (diff string, equals bool) {
	var missingKeysA, extraKeysA, missingKeysB, extraKeysB []string
	equals = true

	iterA := a.MapRange()
	for iterA.Next() {
		key := iterA.Key()
		if !maps.HasKey(b, key) {
			equals = false
			keyStr := fmx.Sprintf("%v", key.Interface())
			missingKeysB = append(missingKeysB, keyStr) // b is missing this key
			extraKeysA = append(extraKeysA, keyStr)     // a has extra key
		}
	}

	iterB := b.MapRange()
	for iterB.Next() {
		key := iterB.Key()
		if !maps.HasKey(a, key) { // ✅ Fixed: check against `a`
			equals = false
			keyStr := fmx.Sprintf("%v", key.Interface())
			missingKeysA = append(missingKeysA, keyStr) // a is missing this key
			extraKeysB = append(extraKeysB, keyStr)     // b has extra key
		}
	}

	if !equals {
		diff = differs.MapKeysDiff(indent.Tab(tab.Value()), missingKeysA, extraKeysA, missingKeysB, extraKeysB)
	}
	return
}

// defaultPtrDiff evaluates the defaultStringDiff between two ptrs given the following constraints:
//
//	A: both addresses are the same; or
//	B: both types are equals; and
//	B: both elements are equals; or
//	C: both ptrs are nil; and
//	C: both ptrs have the same type
func defaultPtrDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[defaultPtrDiff] comparing ptrs %s vs %s: DeepEqual=%t\n",
	//	a.String(), b.String(), reflect.DeepEqual(readableValue(a), readableValue(b)))
	if !a.IsNil() && !b.IsNil() {
		// Cycle check
		if diff, done := s.check(a, b); done {
			//fmx.Yellowf("[defaultPtrDiff] flagCycl hit cache for %s vs %s: %s (cycle)\n",
			//	a.String(), b.String(), diff.Message)
			return diff
		}
		// If it doesn't exist, mark true to avoid infinite recursion
		s.cache(a, b, differs.Equals())

		if a.Pointer() == b.Pointer() { // same instance
			return s.cache(a, b, differs.Equals())
		}

		return s.cache(a, b, ptrElemDiff(tab, a, b, s))
	}

	if a.IsNil() && b.IsNil() {
		if a.Type() != b.Type() {
			aa := types.ValidValueName(a.Type())
			bb := types.ValidValueName(b.Type())
			msg := differs.NilPointersSignMismatch(tab, aa, bb)
			return differs.NotEquals(msg, "", aa, bb)
		}
		return differs.Equals()
	}

	if a.IsNil() {
		aa := types.ValidValueName(a.Type())
		bb := types.ValidValueName(b.Type())
		msg := differs.NilReceived(tab, "ptr", bb)
		return differs.NotEquals(msg, "", aa, bb)
	}

	aa := types.ValidValueName(a.Type())
	bb := types.ValidValueName(b.Type())
	msg := differs.NilExpected(tab, "ptr", aa)
	return differs.NotEquals(msg, "", aa, bb)
}

// ptrElemDiff evaluates the defaultStringDiff between a not nil pointer element
func ptrElemDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	// Recurse into the underlying element
	//fmx.Cyanf("[defaultPtrDiff] recursing into Elem() for %s vs %s\n",
	//	a.String(), b.String())
	differ := between(tab.Inc(), a.Elem(), b.Elem(), s)

	if !differ.Equals {
		//fmx.Redf("[defaultPtrDiff] Elem() differ for %s vs %s: %s\n",
		//	a.String(), b.String(), differ.Message)
		message := differs.PointerValues(tab) + "\n" + differ.Message
		return s.cache(a, b,
			differs.NotEquals(message, differ.Diff, differ.Received, differ.Expected))
	}

	//fmx.Greenf("[defaultPtrDiff] ptrs compare after recursion: %s vs %s\n",
	//	a.String(), b.String())
	return differs.Equals()
}

// samePtrDiff evaluates the defaultStringDiff between two ptrs based on the following constraints:
//
//	A: both ptrs hold the same address; or
//	B: both ptrs are nil; and
//	B: both ptrs have the same type
func samePtrDiff(tab indent.Indentor, a, b reflect.Value, _ *Strategy) differs.Difference {
	if !a.IsNil() && !b.IsNil() {
		pa, pb := a.Pointer(), b.Pointer()
		if pa != pb {
			spa, spb := sprints.Uintptrf(pa), sprints.Uintptrf(pb)
			msg := differs.PointersAddrMismatch(tab, spa, spb)
			return differs.NotEquals(msg, "", spa, spb)
		}
		return differs.Equals()
	}

	ta, tb := a.Type(), b.Type()
	aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)
	if a.IsNil() && b.IsNil() {
		if ta != tb {
			msg := differs.NilPointersSignMismatch(tab, aa, bb)
			return differs.NotEquals(msg, "", aa, bb)
		}
		return differs.Equals()
	}

	if a.IsNil() {
		msg := differs.NilReceived(tab, "ptr", bb)
		return differs.NotEquals(msg, "", aa, bb)
	}
	msg := differs.NilExpected(tab, "ptr", aa)
	return differs.NotEquals(msg, "", aa, bb)
}

func diffStrings(tab indent.Indentor, value, expected string, s *Strategy) differs.Difference {
	//fmx.Purplef("[diffStrings] '%s' compare '%s': %t\n", value, expected,
	//	reflect.DeepEqual(value, expected))
	value, expected = formatStrings(s.stringFormat, value, expected)
	received, expected, idx := s.strings(value, expected)
	if idx == -1 {
		return differs.Equals()
	}

	startReceived := max(0, idx-standards.DifferStartIndex)
	endReceived := min(len(received), idx+standards.DifferMaxSize)

	// Ensuring bounds
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

	receivedSubstring := "'" + received[startReceived:endReceived] + "'"
	expectedSubstring := "'" + expected[startExpected:endExpected] + "'"

	diff := differs.Diff(indent.Tab(tab.Value()), receivedSubstring, expectedSubstring)
	// leave the diff to be resolved on the higher level
	return differs.NotEquals(differs.Strings(tab, idx), diff, receivedSubstring, expectedSubstring)
}

// defaultSliceDiff evaluates the difference between two slices based on the constraints:
//
//	A: Both slices must have the same type; and
//	A: Both slices must have the same length and capacity
//	A: Both slices have equals elements; or
//	B: Both slices have the same type, length, capacity and backing array pointer; or
//	C: Both slices are nil and have the same type.
func defaultSliceDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[defaultSliceDiff] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(readableValue(a), readableValue(b)))

	if !a.IsNil() && !b.IsNil() {
		return defaultNotNilSliceDiff(tab, a, b, s)
	}

	ta, tb := a.Type(), b.Type()
	aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)

	if a.IsNil() && b.IsNil() {
		if ta.Elem() != tb.Elem() {
			return differs.NotEquals(differs.NilTypesMismatch(tab, "slices", aa, bb), "", aa, bb)
		}
		return differs.Equals()
	}

	if a.IsNil() {
		msg := differs.NilReceived(tab, "slice", bb)
		return differs.NotEquals(msg, "", aa, bb)
	}
	msg := differs.NilExpected(tab, "slice", aa)
	return differs.NotEquals(msg, "", aa, bb)
}

func defaultNotNilSliceDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	ta, tb := a.Type(), b.Type()
	if ta.Elem() != tb.Elem() {
		aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)
		return differs.NotEquals(differs.SliceTypesMismatch(tab, aa, bb), "", aa, bb)
	}

	if a.Len() != b.Len() {
		aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)
		return differs.NotEquals(differs.SliceLenMismatch(tab, a.Len(), b.Len()), "", aa, bb)
	}

	if a.Cap() != b.Cap() {
		aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)
		//fmx.Purplef("[differSlice] '%s'.Cap() different from '%s'.Cap()\n", a.String(), b.String())
		return differs.NotEquals(differs.SliceCapMismatch(tab, a.Cap(), b.Cap()),
			"", aa, bb)
	}

	// According to Go's slice header layout, a slice is a struct of:
	//     type slice struct {
	//         Data uintptr
	//         Len  int
	//         Cap  int
	//     }
	//
	// When two slices have the same:
	//   - Type (already checked earlier),
	//   - Length (a.Len() == b.Len()),
	//   - Capacity (a.Cap() == b.Cap()),
	//   - Data pointer (a.pointers() == b.pointers()),
	//
	// they reference the same portion of the same underlying array.
	// This means they are semantically the same instance of the slice.
	// Since Go slices are views into contiguous memory, there's no legal way
	// for two different slice instances with equals type/len/cap and the same
	// backing array pointer to represent different data.
	if a.Pointer() == b.Pointer() {
		return differs.Equals()
	}
	return sliceElemDiff(tab, a, b, s)
}

// sliceElemDiff evaluates the difference between each slice element
// this method doesn't check for type equalities directly, so it isn't reliable
// to be called before assert equality between the slices elements
func sliceElemDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	for i := 0; i < a.Len(); i++ {
		valA := a.Index(i)
		valB := b.Index(i)
		//fmx.Purplef("[defaultSliceDiff] '%s'.Index(%d) compare '%s'.Index(%d): %t\n",
		//	a.String(), i, b.String(), i,
		//	reflect.DeepEqual(readableValue(valA), readableValue(valB)))
		differ := between(tab.Inc(), valA, valB, s)
		if !differ.Equals {
			msg := differs.SliceValues(tab, i) + "\n" + differ.Message
			return differs.NotEquals(msg, differ.Diff, differ.Received, differ.Expected)
		}
	}
	return differs.Equals()
}

// serializableSliceDiff evaluates the difference between two slices based on the constraints:
//
//	A: Both slices must have the same type; and
//	A: Both slices must have the same length
//	A: Both slices have equals elements; or
//	B: Both slices have the same type, length, and backing array pointer; or
//	C: Both slices are nil or empty and have the same type.
func serializableSliceDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[serializableSliceDiff] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(readableValue(a), readableValue(b)))

	if !a.IsNil() && !b.IsNil() {
		return serializableNotNilSliceDiff(tab, a, b, s)
	}

	ta, tb := a.Type(), b.Type()
	aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)

	if (a.IsNil() && b.IsNil()) ||
		(a.IsNil() && b.Len() == 0) ||
		(a.Len() == 0 && b.IsNil()) {
		if ta.Elem() != tb.Elem() {
			return differs.NotEquals(differs.NilTypesMismatch(tab, "slices", aa, bb), "", aa, bb)
		}
		return differs.Equals()
	}

	if a.IsNil() {
		msg := differs.NilReceived(tab, "slice", bb)
		return differs.NotEquals(msg, "", aa, bb)
	}
	msg := differs.NilExpected(tab, "slice", aa)
	return differs.NotEquals(msg, "", aa, bb)
}

func serializableNotNilSliceDiff(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	ta, tb := a.Type(), b.Type()
	aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)

	if ta.Elem() != tb.Elem() {
		return differs.NotEquals(differs.SliceTypesMismatch(tab, aa, bb), "", aa, bb)
	}

	if a.Len() != b.Len() {
		return differs.NotEquals(differs.SliceLenMismatch(tab, a.Len(), b.Len()), "", aa, bb)
	}

	if a.Pointer() == b.Pointer() {
		return differs.Equals()
	}
	return sliceElemDiff(tab, a, b, s)
}

// diffStructs evaluates the difference between two structs based on the constraints:
//
//	A: both structs have the same type; and
//	A: both struct fields are equals;
func diffStructs(tab indent.Indentor, a, b reflect.Value, s *Strategy) (diff differs.Difference) {
	//fmx.Purplef("[diffStructs] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(readableValue(a), readableValue(b)))
	ta, tb := a.Type(), b.Type()
	aa, bb := types.ValidValueName(ta), types.ValidValueName(tb)
	// Structs can't be nil
	if ta != tb {
		return differs.NotEquals(differs.StructTypesMismatch(tab, aa, bb), "", aa, bb)
	}

	if !a.CanAddr() {
		a = values.UnsafeOfUnaddr(a)
		b = values.UnsafeOfUnaddr(b)
	}
	diff = differs.Equals()

	s.rideFields(a, b, func(i int, f1, f2 reflect.Value) bool {
		field := a.Type().Field(i)
		//fmx.Purplef("[diffStructs] '%s'.'%s' compare '%s'.'%s': %t\n",
		//	a.String(), field.Name, b.String(), field.Name,
		//	reflect.DeepEqual(readableValue(f1), readableValue(f2)))

		differ := between(tab.Inc(), f1, f2, s)
		if !differ.Equals {
			message := differs.StructFields(tab, field.Name) + "\n" + differ.Message
			diff = differs.NotEquals(message, differ.Diff, differ.Received, differ.Expected)
			return false // stop iterating
		}
		return true // continue
	})

	return
}

func differUint(tab indent.Indentor, a, b reflect.Value) differs.Difference {
	name := types.ValidValueName(a.Type())
	if a.Uint() != b.Uint() {
		return differs.NotEquals(differs.Values(tab, sprints.Typed(name, a.Uint()),
			sprints.Typed(name, b.Uint())), "", tab.Sprint(sprints.TypedDigit(name, a.Uint())),
			tab.Sprint(sprints.TypedDigit(name, b.Uint())))
	}
	return differs.Equals()
}

// differUnsafePointers evaluates the difference between two unsafe ptrs by the constraints:
//
//	A: both ptrs hold the same address; or
//	B: both ptrs are nil
func differUnsafePointers(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference {
	//fmx.Purplef("[differUnsafePointers] '%s' compare '%s': %t\n", a.String(), b.String(),
	//	reflect.DeepEqual(a.Interface(), b.Interface()))

	if a.IsNil() && b.IsNil() {
		return differs.Equals()
	}

	if !a.IsNil() && !b.IsNil() {
		ptrA := unsafe.Pointer(a.Pointer())
		ptrB := unsafe.Pointer(b.Pointer())

		spa, spb := sprints.Uintptrf(uintptr(ptrA)), sprints.Uintptrf(uintptr(ptrB))
		if ptrA == ptrB {
			return s.cache(a, b, differs.Equals())
		}

		msg := differs.UnsafePointersAddr(tab, spa, spb)
		return s.cache(a, b, differs.NotEquals(msg, "", spa, spb))
	}

	if a.IsNil() {
		ptrB := unsafe.Pointer(b.Pointer())
		msg := differs.NilReceived(tab, "unsafe.pointer", sprints.Uintptrf(uintptr(ptrB)))
		return differs.NotEquals(msg, "", a.String(), b.String())
	}

	ptrA := unsafe.Pointer(a.Pointer())
	msg := differs.NilExpected(tab, "unsafe.pointer", sprints.Uintptrf(uintptr(ptrA)))
	return differs.NotEquals(msg, "", a.String(), b.String())
}

func readableValue(v reflect.Value) any {
	if v.CanInterface() {
		return v.Interface()
	}
	// Use unsafe to bypass unexported access
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
	return v.Interface()
}

func readAnyField(v reflect.Value, i int) (field reflect.Value, canRead bool) {
	return v.Field(i), true
}

func readSettableField(v reflect.Value, i int) (field reflect.Value, canRead bool) {
	f := v.Field(i)
	return f, f.CanSet()
}

func rideAllFields(a, b reflect.Value, fn functions.TriPredicate[int, reflect.Value, reflect.Value]) {
	for i := 0; i < a.NumField(); i++ {
		f1 := a.Field(i)
		f2 := b.Field(i)
		if !fn(i, f1, f2) {
			return
		}
	}
}

func rideSettableFields(a, b reflect.Value, fn functions.TriPredicate[int, reflect.Value, reflect.Value]) {
	for i := 0; i < a.NumField(); i++ {
		f1 := a.Field(i)
		if !f1.CanSet() {
			continue
		}
		f2 := b.Field(i)
		if !fn(i, f1, f2) {
			return
		}
	}
}

func skip(_ indent.Indentor, _, _ reflect.Value, _ *Strategy) differs.Difference {
	return differs.Equals()
}
