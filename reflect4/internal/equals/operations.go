// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package equals

import (
	"reflect"
	"strings"

	"github.com/andrerrcosta2/gtools/core/util/typeutil/stringutil"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
)

func bytealg(a, b []byte) bool {
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	// There is a test for this in package bytes.
	return string(a) == string(b)
}

// defaultArray evaluates array equality as:
//  1. Same type
//  2. Same length
//  3. equality between elements
func defaultArray(a, b reflect.Value, differ *Strategy) bool {
	al, bl := a.Len(), b.Len()
	if al != bl {
		return false
	}
	for i := 0; i < al; i++ {
		if !deep(a.Index(i), b.Index(i), differ) {
			return false
		}
	}
	return true
}

// defaultChan evaluates channel equality as:
//  1. Same length
//  2. Same capacity
func defaultChan(a, b reflect.Value) bool {
	if !a.IsNil() && !b.IsNil() {
		return a.Len() == b.Len() &&
			a.Cap() == b.Cap()
	}
	return a.IsNil() && b.IsNil()
}

// defaultFunc evaluates function equality as:
//  1. Same type
//  2. Same nullability state
func defaultFunc(a, b reflect.Value) bool {
	return a.IsNil() == b.IsNil()
}

// defaultInterface evaluates interface equality as:
//  1. Same type
//  2. Same nullability state
//  3. equality between elements
func defaultInterface(a, b reflect.Value, differ *Strategy) bool {
	if a.IsNil() || b.IsNil() {
		return a.IsNil() == b.IsNil()
	}
	return deep(a.Elem(), b.Elem(), differ)
}

// defaultMap evaluates msp equality as:
//  1. Same type
//  2. Same nullability state
//  3. Same length
//  4. equality between elements
func defaultMap(a, b reflect.Value, differ *Strategy) bool {
	if a.IsNil() != b.IsNil() {
		return false
	}
	if a.Len() != b.Len() {
		return false
	}
	if a.UnsafePointer() == b.UnsafePointer() {
		return true
	}
	if differ.check(a, b) {
		return true
	}
	differ.mark(a, b)
	// Detect whether we need deep key comparison
	if differ.shouldDeepCompareKeys(a.Type().Key()) {
		return deepKeyedMapEquals(a, b, differ)
	}
	return shallowKeyedMapEquals(a, b, differ)
}

// shallowKeyedMapEquals performs a shallow comparison for values whose key equality is met
// directly. it performs a O(n) comparisons.
func shallowKeyedMapEquals(a, b reflect.Value, differ *Strategy) bool {
	for _, keyA := range a.MapKeys() {
		valA := a.MapIndex(keyA)
		valB := b.MapIndex(keyA)
		if !valB.IsValid() {
			return false // missing key in b
		}
		if !deep(valA, valB, differ) {
			return false
		}
	}
	return true
}

// deepKeyedMapEquals performs a deep comparison for maps whose key equality
// cannot be determined by Go's built-in == operator. It performs O(n²)
// comparisons but ensures one-to-one matching between keys.
func deepKeyedMapEquals(a, b reflect.Value, differ *Strategy) bool {
	// Collect keys of b once
	keysB := b.MapKeys()
	used := make([]bool, len(keysB))

	// For each keyA in a, find an unused keyB such that:
	//   deep(keyA, keyB) == true  AND  deep(valA, valB) == true
	for _, keyA := range a.MapKeys() {
		valA := a.MapIndex(keyA)
		found := false

		for i, keyB := range keysB {
			if used[i] {
				continue
			}

			// If keys are semantically equal, try the value
			if !deep(keyA, keyB, differ) {
				continue
			}

			valB := b.MapIndex(keyB)
			if deep(valA, valB, differ) {
				// only mark used when we found a matching key+value pair
				used[i] = true
				found = true
				break
			}
			// otherwise: key matches but value doesn't — keep searching other keyB
		}

		if !found {
			// no keyB had both equal key and equal value
			return false
		}
	}

	// ensure there are no unmatched keys left in b
	for _, u := range used {
		if !u {
			return false
		}
	}

	return true
}

// defaultPtr evaluates pointer equality as:
//  1. Same type
//  2. equality between elements
func defaultPtr(a, b reflect.Value, differ *Strategy) bool {
	if a.UnsafePointer() == b.UnsafePointer() {
		return true
	}
	if differ.check(a, b) {
		return true
	}
	differ.mark(a, b)
	return deep(a.Elem(), b.Elem(), differ)
}

// defaultSlice evaluates slice equality as:
//  1. Same type
//  2. Same length
//  2. equality between elements
func defaultSlice(a, b reflect.Value, differ *Strategy) bool {
	if a.IsNil() != b.IsNil() {
		return false
	}
	if a.Len() != b.Len() {
		return false
	}
	if a.UnsafePointer() == b.UnsafePointer() {
		return true
	}
	if differ.check(a, b) {
		return true
	}
	differ.mark(a, b)
	// Special case for []byte, which is common.
	if a.Type().Elem().Kind() == reflect.Uint8 {
		return bytealg(a.Bytes(), b.Bytes())
	}
	for i := 0; i < a.Len(); i++ {
		if !deep(a.Index(i), b.Index(i), differ) {
			return false
		}
	}
	return true
}

// defaultStruct evaluates slice equality as:
//  1. Same type
//  2. equality between fields
func defaultStruct(a, b reflect.Value, differ *Strategy) (eq bool) {
	eq = true
	differ.rideFields(a, func(i int, value reflect.Value) bool {
		if !deep(a.Field(i), b.Field(i), differ) {
			eq = false
			return false // short-circuit
		}
		return true
	})
	return
}

// defaultUnsafePointer evaluates unsafe pointer equality as:
//  1. Same address
func defaultUnsafePointer(a, b reflect.Value) bool {
	return a.UnsafePointer() == b.UnsafePointer()
}

func differStrings(a, b reflect.Value, differ *Strategy) bool {
	return differ.strings(formatString(a.String(), differ.stringFormat),
		formatString(b.String(), differ.stringFormat))
}

func formatString(s string, f op.Compare) string {
	if f == 0 {
		return s
	}
	if f&compare.TrimSpace != 0 {
		s = strings.TrimSpace(s)
	}
	if f&compare.IgnoreWhitespace != 0 {
		s = stringutil.NormWS(s)
	}
	if f&compare.IgnoreCase != 0 {
		s = strings.ToLower(s)
	}
	if f&compare.IgnoreAccents != 0 {
		s = stringutil.RemoveAccents(s)
	}
	return s
}

// identityShallow evaluates any shallow type equality by:
//  1. same type
//  2. same nullable state
//  3. same memory address
func identityShallow(a, b reflect.Value) bool {
	if !a.IsNil() && !b.IsNil() {
		return a.Pointer() == b.Pointer()
	}
	return a.IsNil() && b.IsNil()
}

// identitySlice evaluates slice equality by:
//  1. same type
//  2. same nullable state
//  3. same memory address
//  4. same length (contiguous fetching space)
func identitySlice(a reflect.Value, b reflect.Value, _ *Strategy) bool {
	if !a.IsNil() && !b.IsNil() {
		// slices might share index addresses not being the same
		return a.Pointer() == b.Pointer() &&
			a.Len() == b.Len()
	}
	return a.IsNil() && b.IsNil()
}

// identityShallow evaluates any deep type equality by:
//  1. same type
//  2. same nullable state
//  3. same memory address
func identityDeep(a, b reflect.Value, _ *Strategy) bool {
	if !a.IsNil() && !b.IsNil() {
		return a.Pointer() == b.Pointer()
	}
	return a.IsNil() && b.IsNil()
}

// identityShallow evaluates string equality by:
//  1. equal values ignoring case
func ignoreCaseString(a, b reflect.Value) bool {
	return stringutil.EqualsIgnoreCase(a.String(), b.String())
}

// nativeChan evaluates channel equality by:
//  1. same types
//  2. same nullable state
//  3. same memory address
func nativeChan(a, b reflect.Value) bool {
	return a.IsNil() && b.IsNil() || a.Pointer() == b.Pointer()
}

// nativeFunc evaluates func equality by:
//  1. same type
//  2. both are nil
func nativeFunc(a, b reflect.Value) bool {
	return a.IsNil() && b.IsNil()
}

// serialChan evaluates channel equality by:
//  1. same type
//  2. same size if not nil
//  3. empty and nil are equals
func serialChan(a, b reflect.Value) bool {
	if a.IsNil() || b.IsNil() {
		return (a.IsNil() && (b.IsNil() || b.Len() == 0)) ||
			(b.IsNil() && a.Len() == 0)
	}
	return a.Len() == b.Len()
}

// serialMap evaluates map equality by:
//  1. same type
//  2. same size and equal elements if not nil
//  3. empty and nil are equals
func serialMap(a, b reflect.Value, differ *Strategy) bool {
	if a.IsNil() || b.IsNil() {
		return (a.IsNil() && (b.IsNil() || b.Len() == 0)) ||
			(b.IsNil() && a.Len() == 0)
	}
	if a.Len() != b.Len() {
		return false
	}
	if a.UnsafePointer() == b.UnsafePointer() {
		return true
	}
	if differ.check(a, b) {
		return true
	}
	differ.mark(a, b)

	for _, keyA := range a.MapKeys() {
		valA := a.MapIndex(keyA)
		valB := b.MapIndex(keyA)
		if !valB.IsValid() {
			return false // missing key in b
		}
		if !deep(valA, valB, differ) {
			return false
		}
	}
	return true
}

// serialMap evaluates slice equality by:
//  1. same type
//  2. same size and equal elements if not nil
//  3. empty and nil are equals
func serialSlice(a, b reflect.Value, differ *Strategy) bool {
	if a.IsNil() || b.IsNil() {
		return a.IsNil() && (b.IsNil() || b.Len() == 0) || b.IsNil() && a.Len() == 0
	}
	if a.Len() != b.Len() {
		return false
	}
	if a.UnsafePointer() == b.UnsafePointer() {
		return true
	}
	if differ.check(a, b) {
		return true
	}
	differ.mark(a, b)
	// Special case for []byte, which is common.
	if a.Type().Elem().Kind() == reflect.Uint8 {
		return bytealg(a.Bytes(), b.Bytes())
	}
	for i := 0; i < a.Len(); i++ {
		if !deep(a.Index(i), b.Index(i), differ) {
			return false
		}
	}
	return true
}

func deepSkip(_, _ reflect.Value, _ *Strategy) bool {
	return true
}

func shallowSkip(_, _ reflect.Value) bool {
	return true
}

//func typeMightContainPointers(t reflect.Type) bool {
//	switch t.Kind() {
//	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface, reflect.Chan, reflect.Func:
//		return true
//	case reflect.Array:
//		return typeMightContainPointers(t.Elem())
//	case reflect.Struct:
//		for i := 0; i < t.NumField(); i++ {
//			if typeMightContainPointers(t.Field(i).Type) {
//				return true
//			}
//		}
//	}
//	return false
//}
