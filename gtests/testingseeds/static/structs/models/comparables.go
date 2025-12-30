// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import "github.com/andrerrcosta2/gtools/core/seeders/random"

var NaturallyComparableZeroInst = new(NaturallyComparable)

// NaturallyComparableAsValue is a struct that can be compared using the == operator
// StructsPtr are naturally comparable if they do not contain:
//   - Slices: Since slices are inherently non-comparable (they are reference type4
//     with varying lengths and capacities), any struct containing a slice is also non-comparable.
//   - Maps: Maps are reference type4 as well and cannot be compared with == (except for nil comparisons).
//   - Functions: Function field4 are non-comparable because they refer to code blocks and may hold
//     unique address or closures, making them inherently unique per instance.
//   - Arrays with Non-Comparable Elements: If a struct contains an array where elements themselves are
//     non-comparable, the array (and hence the struct) becomes non-comparable.
//   - Channels: Channels are reference type4 and cannot be compared with ==.
//   - StructsPtr and interfaces with non-comparable field4: If a struct contains a field that is non-comparable,
//     the struct itself becomes non-comparable.
//   - unsafe.Pointer: Since ptrs are inherently non-comparable, any struct containing a pointer
//     is also non-comparable.
func NaturallyComparableAsValue(a int, b string) NaturallyComparable {
	return NaturallyComparable{a, b}
}

// NaturallyComparableAsRef is a struct that can be compared using the == operator
// StructsPtr are naturally comparable if they do not contain:
//   - Slices: Since slices are inherently non-comparable (they are reference type4
//     with varying lengths and capacities), any struct containing a slice is also non-comparable.
//   - Maps: Maps are reference type4 as well and cannot be compared with == (except for nil comparisons).
//   - Functions: Function field4 are non-comparable because they refer to code blocks and may hold
//     unique address or closures, making them inherently unique per instance.
//   - Arrays with Non-Comparable Elements: If a struct contains an array where elements themselves are
//     non-comparable, the array (and hence the struct) becomes non-comparable.
//   - Channels: Channels are reference type4 and cannot be compared with ==.
//   - StructsPtr and interfaces with non-comparable field4: If a struct contains a field that is non-comparable,
//     the struct itself becomes non-comparable.
//   - unsafe.Pointer: Since ptrs are inherently non-comparable, any struct containing a pointer
//     is also non-comparable.
func NaturallyComparableAsRef(a int, b string) *NaturallyComparable {
	return &NaturallyComparable{a, b}
}

func NaturallyComparableAsRandRef() *NaturallyComparable {
	return &NaturallyComparable{
		ComparableValueA: random.SingleOf[int](),
		ComparableValueB: random.SingleOf[string](),
	}
}

func NaturallyComparableAsRandValue() NaturallyComparable {
	return NaturallyComparable{
		ComparableValueA: random.SingleOf[int](),
		ComparableValueB: random.SingleOf[string](),
	}
}

type NaturallyComparable struct {
	ComparableValueA int
	ComparableValueB string
}

var NaturallyComparableWithMethodsZeroInst = new(NaturallyComparableWithMethods)

func NaturallyComparableWithMethodsAsValue(a int, b string) NaturallyComparableWithMethods {
	return NaturallyComparableWithMethods{a, b}
}

func NaturallyComparableWithMethodsAsRef(a int, b string) *NaturallyComparableWithMethods {
	return &NaturallyComparableWithMethods{a, b}
}

func NaturallyComparableWithMethodsAsRandRef() *NaturallyComparableWithMethods {
	return &NaturallyComparableWithMethods{
		comparableValueA: random.SingleOf[int](),
		comparableValueB: random.SingleOf[string](),
	}
}

func NaturallyComparableWithMethodsAsRandValue() NaturallyComparableWithMethods {
	return NaturallyComparableWithMethods{
		comparableValueA: random.SingleOf[int](),
		comparableValueB: random.SingleOf[string](),
	}
}

type NaturallyComparableWithMethods struct {
	comparableValueA int
	comparableValueB string
}

func (n *NaturallyComparableWithMethods) ComparableValueA() int {
	return n.comparableValueA
}

func (n *NaturallyComparableWithMethods) ComparableValueB() string {
	return n.comparableValueB
}

// NotComparableZeroInst is a struct that cannot be compared using the == operator
// StructsPtr are non-comparable if they contain:
//   - Slices: Since slices are inherently non-comparable (they are reference type4
//     with varying lengths and capacities), any struct containing a slice is also non-comparable.
//   - Maps: Maps are reference type4 as well and cannot be compared with == (except for nil comparisons).
//   - Functions: Function field4 are non-comparable because they refer to code blocks and may hold
//     unique address or closures, making them inherently unique per instance.
//   - Arrays with Non-Comparable Elements: If a struct contains an array where elements themselves are
//     non-comparable, the array (and hence the struct) becomes non-comparable.
//   - Channels: Channels are reference type4 and cannot be compared with ==.
//   - StructsPtr and interfaces with non-comparable field4: If a struct contains a field that is non-comparable,
//     the struct itself becomes non-comparable.
//   - unsafe.Pointer: Since ptrs are inherently non-comparable, any struct containing a pointer
//     is also non-comparable.
var NotComparableZeroInst = new(NotComparable)

func NotComparableAsValue(a int, b []string) NotComparable {
	return NotComparable{a, b}
}

func NotComparableAsRef(a int, b []string) *NotComparable {
	return &NotComparable{a, b}
}

func NotComparableAsRandRef() *NotComparable {
	return &NotComparable{
		ComparableValueA:    random.SingleOf[int](),
		NotComparableValueB: random.SingleOf[[]string](),
	}
}

func NotComparableAsRandValue() NotComparable {
	return NotComparable{
		ComparableValueA:    random.SingleOf[int](),
		NotComparableValueB: random.SingleOf[[]string](),
	}
}

type NotComparable struct {
	ComparableValueA    int
	NotComparableValueB []string
}

var NotComparableWithMethodsZeroInst = new(NotComparableWithMethods)

func NotComparableWithMethodsAsValue(a int, b []string) NotComparableWithMethods {
	return NotComparableWithMethods{a, b}
}

func NotComparableWithMethodsAsRef(a int, b []string) *NotComparableWithMethods {
	return &NotComparableWithMethods{a, b}
}

func NotComparableWithMethodsAsRandRef() *NotComparableWithMethods {
	return &NotComparableWithMethods{
		comparableValueA:    random.SingleOf[int](),
		notComparableValueB: random.SingleOf[[]string](),
	}
}

func NotComparableWithMethodsAsRandValue() NotComparableWithMethods {
	return NotComparableWithMethods{
		comparableValueA:    random.SingleOf[int](),
		notComparableValueB: random.SingleOf[[]string](),
	}
}

type NotComparableWithMethods struct {
	comparableValueA    int
	notComparableValueB []string
}

func (n *NotComparableWithMethods) ComparableValueA() int {
	return n.comparableValueA
}

func (n *NotComparableWithMethods) NotComparableValueB() []string {
	return n.notComparableValueB
}

var ValuedNotComparableZeroInst = new(ValuedNotComparable)

func ValuedNotComparableAsValue(value int) ValuedNotComparable {
	return ValuedNotComparable{value, []string{}}
}

func ValuedNotComparableAsRef(value int) *ValuedNotComparable {
	return &ValuedNotComparable{value, []string{}}
}

func ValuedNotComparableAsRandRef() *ValuedNotComparable {
	return &ValuedNotComparable{
		value:              random.SingleOf[int](),
		notComparableField: random.String(10, 5, 20).Values(),
	}
}

func ValuedNotComparableAsRandValue() ValuedNotComparable {
	return ValuedNotComparable{
		value:              random.SingleOf[int](),
		notComparableField: random.String(10, 5, 20).Values(),
	}
}

type ValuedNotComparable struct {
	value              int
	notComparableField []string
}

func (v ValuedNotComparable) Value() int {
	return v.value
}

func (v ValuedNotComparable) NotComparableField() []string {
	return v.notComparableField
}
