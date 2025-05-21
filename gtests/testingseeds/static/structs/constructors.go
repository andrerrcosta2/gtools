// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

var ZeroSimpleInst = SimpleAsRef()

func SimpleAsRef() *Simple {
	return &Simple{}
}

func SimpleAsValue() Simple {
	return Simple{}
}

var ZeroPublicInst = PublicAsRef()

func PublicAsRef() *Public {
	return &Public{}
}

func PublicAsValue() Public {
	return Public{}
}

var ZeroOneDataInst = OneDataAsRef("")

func OneDataAsRef(name string) *OneData {
	return &OneData{
		Name: name,
	}
}

func OneDataAsValue(name string) OneData {
	return OneData{
		Name: name,
	}
}

var ZeroTwoDataInst = TwoDataAsRef("", 0)

func TwoDataAsRef(name string, age int) *TwoData {
	return &TwoData{
		Name: name,
		Age:  age,
	}
}

func TwoDataAsValue(name string, age int) TwoData {
	return TwoData{
		Name: name,
		Age:  age,
	}
}

var ZeroStringerInst = StringerAsRef("")

func StringerAsRef(data any) *Stringer {
	return &Stringer{
		data: data,
	}
}

func StringerAsValue(data any) Stringer {
	return Stringer{
		data: data,
	}
}

var ZeroStringerStringInst = StringerStringAsRef("")

func StringerStringAsValue(data string) StringerString {
	return StringerString{data}
}

func StringerStringAsRef(data string) *StringerString {
	return &StringerString{data}
}

var ZeroStringerBytesInst = StringerBytesAsRef([]byte{})

func StringerBytesAsValue(data []byte) StringerBytes {
	return StringerBytes{data}
}

func StringerBytesAsRef(data []byte) *StringerBytes {
	return &StringerBytes{data}
}

var ZeroCloserSuccessInst = CloserSuccessAsRef()

func CloserSuccessAsRef() *CloserSuccess {
	return &CloserSuccess{}
}

func CloserSuccessAsValue() CloserSuccess {
	return CloserSuccess{}
}

var ZeroCloserErrorInst = CloserErrorAsRef()

func CloserErrorAsRef() *CloserError {
	return &CloserError{}
}

func CloserErrorAsValue() CloserError {
	return CloserError{}
}

var ZeroCloserReaderSuccessInst = CloserReaderSuccessAsRef()

func CloserReaderSuccessAsRef() *CloserReaderSuccess {
	return &CloserReaderSuccess{}
}

func CloserReaderSuccessAsValue() CloserReaderSuccess {
	return CloserReaderSuccess{}
}

var ZeroCloserReaderErrorInst = CloserReaderErrorAsRef()

func CloserReaderErrorAsRef() *CloserReaderError {
	return &CloserReaderError{}
}

func CloserReaderErrorAsValue() CloserReaderError {
	return CloserReaderError{}
}

var ZeroCloserReaderWriterSuccessInst = CloserReaderWriterSuccessAsRef()

func CloserReaderWriterSuccessAsRef() *CloserReaderWriterSuccess {
	return &CloserReaderWriterSuccess{}
}

func CloserReaderWriterSuccessAsValue() CloserReaderWriterSuccess {
	return CloserReaderWriterSuccess{}
}

var ZeroCloserReaderWriterErrorInst = CloserReaderWriterErrorAsRef()

func CloserReaderWriterErrorAsRef() *CloserReaderWriterError {
	return &CloserReaderWriterError{}
}

func CloserReaderWriterErrorAsValue() CloserReaderWriterError {
	return CloserReaderWriterError{}
}

var ZeroNaturallyComparableInst = NaturallyComparableAsRef(0, "")

// NaturallyComparableAsValue is a struct that can be compared using the == operator
// Structs are naturally comparable if they do not contain:
//   - Slices: Since slices are inherently non-comparable (they are reference types
//     with varying lengths and capacities), any struct containing a slice is also non-comparable.
//   - Maps: Maps are reference types as well and cannot be compared with == (except for nil comparisons).
//   - Functions: Function fields are non-comparable because they refer to code blocks and may hold
//     unique addresses or closures, making them inherently unique per instance.
//   - Arrays with Non-Comparable Elements: If a struct contains an array where elements themselves are
//     non-comparable, the array (and hence the struct) becomes non-comparable.
//   - Channels: Channels are reference types and cannot be compared with ==.
//   - Structs and interfaces with non-comparable fields: If a struct contains a field that is non-comparable,
//     the struct itself becomes non-comparable.
//   - unsafe.Pointer: Since pointers are inherently non-comparable, any struct containing a pointer
//     is also non-comparable.
func NaturallyComparableAsValue(a int, b string) NaturallyComparable {
	return NaturallyComparable{a, b}
}

// NaturallyComparableAsRef is a struct that can be compared using the == operator
// Structs are naturally comparable if they do not contain:
//   - Slices: Since slices are inherently non-comparable (they are reference types
//     with varying lengths and capacities), any struct containing a slice is also non-comparable.
//   - Maps: Maps are reference types as well and cannot be compared with == (except for nil comparisons).
//   - Functions: Function fields are non-comparable because they refer to code blocks and may hold
//     unique addresses or closures, making them inherently unique per instance.
//   - Arrays with Non-Comparable Elements: If a struct contains an array where elements themselves are
//     non-comparable, the array (and hence the struct) becomes non-comparable.
//   - Channels: Channels are reference types and cannot be compared with ==.
//   - Structs and interfaces with non-comparable fields: If a struct contains a field that is non-comparable,
//     the struct itself becomes non-comparable.
//   - unsafe.Pointer: Since pointers are inherently non-comparable, any struct containing a pointer
//     is also non-comparable.
func NaturallyComparableAsRef(a int, b string) *NaturallyComparable {
	return &NaturallyComparable{a, b}
}

var ZeroNaturallyComparableWithMethodsInst = NaturallyComparableWithMethodsAsRef(0, "")

func NaturallyComparableWithMethodsAsValue(a int, b string) NaturallyComparableWithMethods {
	return NaturallyComparableWithMethods{a, b}
}

func NaturallyComparableWithMethodsAsRef(a int, b string) *NaturallyComparableWithMethods {
	return &NaturallyComparableWithMethods{a, b}
}

var ZeroNotComparableInst = NotComparableAsRef(0, nil)

// NotComparableAsValue is a struct that cannot be compared using the == operator
// Structs are non-comparable if they contain:
//   - Slices: Since slices are inherently non-comparable (they are reference types
//     with varying lengths and capacities), any struct containing a slice is also non-comparable.
//   - Maps: Maps are reference types as well and cannot be compared with == (except for nil comparisons).
//   - Functions: Function fields are non-comparable because they refer to code blocks and may hold
//     unique addresses or closures, making them inherently unique per instance.
//   - Arrays with Non-Comparable Elements: If a struct contains an array where elements themselves are
//     non-comparable, the array (and hence the struct) becomes non-comparable.
//   - Channels: Channels are reference types and cannot be compared with ==.
//   - Structs and interfaces with non-comparable fields: If a struct contains a field that is non-comparable,
//     the struct itself becomes non-comparable.
//   - unsafe.Pointer: Since pointers are inherently non-comparable, any struct containing a pointer
//     is also non-comparable.
func NotComparableAsValue(a int, b []string) NotComparable {
	return NotComparable{a, b}
}

// NotComparableAsRef is a struct that cannot be compared using the == operator
// Structs are non-comparable if they contain:
//   - Slices: Since slices are inherently non-comparable (they are reference types
//     with varying lengths and capacities), any struct containing a slice is also non-comparable.
//   - Maps: Maps are reference types as well and cannot be compared with == (except for nil comparisons).
//   - Functions: Function fields are non-comparable because they refer to code blocks and may hold
//     unique addresses or closures, making them inherently unique per instance.
//   - Arrays with Non-Comparable Elements: If a struct contains an array where elements themselves are
//     non-comparable, the array (and hence the struct) becomes non-comparable.
//   - Channels: Channels are reference types and cannot be compared with ==.
//   - Structs and interfaces with non-comparable fields: If a struct contains a field that is non-comparable,
//     the struct itself becomes non-comparable.
//   - unsafe.Pointer: Since pointers are inherently non-comparable, any struct containing a pointer
//     is also non-comparable.
func NotComparableAsRef(a int, b []string) *NotComparable {
	return &NotComparable{a, b}
}

var ZeroNotComparableWithMethodsInst = NotComparableWithMethodsAsRef(0, nil)

func NotComparableWithMethodsAsValue(a int, b []string) NotComparableWithMethods {
	return NotComparableWithMethods{a, b}
}

func NotComparableWithMethodsAsRef(a int, b []string) *NotComparableWithMethods {
	return &NotComparableWithMethods{a, b}
}

var ZeroBooleanInst = BoolAsRef()

func BoolAsValue() Boolean {
	return false
}

func BoolAsRef() *Boolean {
	return new(Boolean)
}

var ZeroIntegerInst = IntAsRef()

func IntAsValue() Integer {
	return 0
}

func IntAsRef() *Integer {
	return new(Integer)
}

var ZeroUintInst = UintAsRef()

func UintAsValue() Uint {
	return 0
}

func UintAsRef() *Uint {
	return new(Uint)
}

var ZeroFloatInst = FloatAsRef()

func FloatAsValue() Float {
	return 0.0
}

func FloatAsRef() *Float {
	return new(Float)
}

var ZeroComplexInst = ComplexAsRef()

func ComplexAsValue() Complex {
	return 0 + 0i
}

func ComplexAsRef() *Complex {
	return new(Complex)
}

var ZeroStringInst = StringAsRef()

func StringAsValue() String {
	return ""
}

func StringAsRef() *String {
	return new(String)
}

var ZeroChannelInst = ChannelAsRef()

func ChannelAsValue() Channel[int] {
	return make(Channel[int])
}

func ChannelAsRef() *Channel[int] {
	return new(Channel[int])
}

var ZeroSliceInst = SliceAsRef[int]()

func SliceAsValue[T any]() Slice[T] {
	return make([]T, 0)
}

func SliceAsRef[T any]() *Slice[T] {
	return new(Slice[T])
}

var ZeroMapInst = MapAsRef[int, int]()

func MapAsValue[K comparable, V any]() Map[K, V] {
	return make(Map[K, V])
}

func MapAsRef[K comparable, V any]() *Map[K, V] {
	return new(Map[K, V])
}

var ZeroValueNotComparableInst = ValuedNotComparableAsRef(0)

func ValuedNotComparableAsValue(value int) ValuedNotComparable {
	return ValuedNotComparable{value, []string{}}
}

func ValuedNotComparableAsRef(value int) *ValuedNotComparable {
	return &ValuedNotComparable{value, []string{}}
}

var ZeroUnsafeCastableStringInst = SimpleUnsafeCastableStringAsRef("")

func SimpleUnsafeCastableStringAsValue(value string) SimpleUnsafeCastableString {
	return SimpleUnsafeCastableString{ValueA: value}
}

func SimpleUnsafeCastableStringAsRef(value string) *SimpleUnsafeCastableString {
	return &SimpleUnsafeCastableString{ValueA: value}
}

var ZeroUnsafeCastableIntInst = SimpleUnsafeCastableIntAsRef(0)

func SimpleUnsafeCastableIntAsValue(value int) SimpleUnsafeCastableInt {
	return SimpleUnsafeCastableInt{ValueB: value}
}

func SimpleUnsafeCastableIntAsRef(value int) *SimpleUnsafeCastableInt {
	return &SimpleUnsafeCastableInt{ValueB: value}
}

var ZeroUnsafeCastableFloatInst = SimpleUnsafeCastableFloatAsRef(0.0)

func SimpleUnsafeCastableFloatAsValue(value float64) SimpleUnsafeCastableFloat {
	return SimpleUnsafeCastableFloat{ValueC: value}
}

func SimpleUnsafeCastableFloatAsRef(value float64) *SimpleUnsafeCastableFloat {
	return &SimpleUnsafeCastableFloat{ValueC: value}
}

var ZeroUnsafeCastableBoolInst = SimpleUnsafeCastableBoolAsRef(false)

func SimpleUnsafeCastableBoolAsValue(value bool) SimpleUnsafeCastableBool {
	return SimpleUnsafeCastableBool{ValueD: value}
}

func SimpleUnsafeCastableBoolAsRef(value bool) *SimpleUnsafeCastableBool {
	return &SimpleUnsafeCastableBool{ValueD: value}
}

var ZeroComplexUnsafeCastableInst = ComplexUnsafeCastableOfAsRef(NestedStruct{}, 0, map[string]any{}, 0.0)

func ComplexUnsafeCastableOfAsValue(nested NestedStruct, id int, metadata map[string]any, score float64) ComplexUnsafeCastableOf {
	return ComplexUnsafeCastableOf{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
	}
}

func ComplexUnsafeCastableOfAsRef(nested NestedStruct, id int, metadata map[string]any, score float64) *ComplexUnsafeCastableOf {
	return &ComplexUnsafeCastableOf{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
	}
}

func ComplexUnsafeCastableDescriptionAsValue(nested NestedStruct, id int, metadata map[string]any, score float64, description string) ComplexUnsafeCastableDescription {
	return ComplexUnsafeCastableDescription{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
		Description: description,
	}
}

// CyclicSimpleNode creates a linked list with a cycle
func CyclicSimpleNode(a, b any) *SimpleNode {
	n1 := &SimpleNode{
		Value: a,
	}
	n2 := &SimpleNode{
		Next:  n1,
		Value: b,
	}
	n1.Next = n2
	return n1
}

var CyclicSimpleNodeInst = CyclicSimpleNode(1, 2)

// CyclicLinkedList creates a double linked list with a cycle
func CyclicLinkedList(values ...any) *LinkedList {
	list := &LinkedList{}
	var previous *BinaryNode
	for _, value := range values {
		node := &BinaryNode{Value: value}
		if list.Head == nil {
			list.Head = node
		}
		if previous != nil {
			previous.Right = node
			node.Left = previous
		}
		previous = node
	}

	list.Tail = previous
	list.Tail.Right = list.Head
	list.Head.Left = list.Tail
	return list
}
