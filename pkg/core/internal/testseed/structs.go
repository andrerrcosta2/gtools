// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

//go:build tt

package testseed

import (
	"errors"
	"fmt"
	"math/rand"
)

type StructEmpty struct{}

type StructSimple struct{}

func (s *StructSimple) Simple() int {
	return rand.Intn(1000)
}

type StructPublic struct{}

func (p *StructPublic) PublicMethod() int {
	return rand.Intn(1000)
}

type StructOneData struct {
	Name string
}

func (o *StructOneData) String() string {
	return o.Name
}

type StructTwoData struct {
	Name string
	Age  int
}

func (t *StructTwoData) String() string {
	return t.Name
}

func (t *StructTwoData) Int() int {
	return t.Age
}

type StructStringer struct {
	data any
}

func (s *StructStringer) String() string {
	return fmt.Sprintf("%v", s.data)
}

type StructStringerString struct{ data string }

func (s *StructStringerString) String() string {
	return s.data
}

type StructStringerBytes struct{ data []byte }

func (s *StructStringerBytes) String() string {
	return string(s.data)
}

type StructCloserSuccess struct{}

func (c *StructCloserSuccess) Close() error { return nil }

type StructCloserError struct{}

func (c *StructCloserError) Close() error { return errors.New("error") }

type StructCloserReaderSuccess struct {
	data []byte
}

func (c *StructCloserReaderSuccess) Close() error { return nil }

func (c *StructCloserReaderSuccess) Read(data []byte) (int, error) {
	c.data = data
	return len(c.data), nil
}

type StructCloserReaderError struct {
	data []byte
}

func (c *StructCloserReaderError) Close() error { return errors.New("error closing") }

func (c *StructCloserReaderError) Read(data []byte) (int, error) {
	return len(c.data), errors.New("error reading")
}

type StructCloserReaderWriterSuccess struct {
	read  []byte
	write []byte
}

func (c *StructCloserReaderWriterSuccess) Close() error { return nil }

func (c *StructCloserReaderWriterSuccess) Read(data []byte) (int, error) {
	c.read = data
	return len(c.read), nil
}

func (c *StructCloserReaderWriterSuccess) Write(data []byte) (int, error) {
	c.write = data
	return len(c.write), nil
}

type StructCloserReaderWriterError struct {
	read  []byte
	write []byte
}

func (c *StructCloserReaderWriterError) Close() error { return errors.New("error closing") }

func (c *StructCloserReaderWriterError) Read(data []byte) (int, error) {
	return len(c.read), errors.New("error reading")
}

func (c *StructCloserReaderWriterError) Write(data []byte) (int, error) {
	return len(c.write), errors.New("error writing")
}

type StructNaturallyComparable struct {
	ComparableValueA int
	ComparableValueB string
}

type StructNaturallyComparableWithMethods struct {
	comparableValueA int
	comparableValueB string
}

func (n *StructNaturallyComparableWithMethods) ComparableValueA() int {
	return n.comparableValueA
}

func (n *StructNaturallyComparableWithMethods) ComparableValueB() string {
	return n.comparableValueB
}

type StructNotComparable struct {
	ComparableValueA    int
	NotComparableValueB []string
}

type StructNotComparableWithMethods struct {
	comparableValueA    int
	notComparableValueB []string
}

func (n *StructNotComparableWithMethods) ComparableValueA() int {
	return n.comparableValueA
}

func (n *StructNotComparableWithMethods) NotComparableValueB() []string {
	return n.notComparableValueB
}

type StructInteger int

func (i StructInteger) String() string {
	return fmt.Sprintf("%d", i)
}

func (i StructInteger) Add(i2 StructInteger) StructInteger {
	return i + i2
}

func (i StructInteger) Sub(i2 StructInteger) StructInteger {
	return i - i2
}

func (i StructInteger) Mul(i2 StructInteger) StructInteger {
	return i * i2
}

func (i StructInteger) Div(i2 StructInteger) StructInteger {
	return i / i2
}

type StructFloat float64

func (f StructFloat) String() string {
	return fmt.Sprintf("%f", f)
}

func (f StructFloat) Add(f2 StructFloat) StructFloat {
	return f + f2
}

func (f StructFloat) Sub(f2 StructFloat) StructFloat {
	return f - f2
}

func (f StructFloat) Mul(f2 StructFloat) StructFloat {
	return f * f2
}

func (f StructFloat) Div(f2 StructFloat) StructFloat {
	return f / f2
}

type StructBoolean bool

func (b StructBoolean) String() string {
	return fmt.Sprintf("%t", b)
}

func (b StructBoolean) Not() StructBoolean {
	return !b
}

func (b StructBoolean) And(b2 StructBoolean) StructBoolean {
	return b && b2
}

func (b StructBoolean) Or(b2 StructBoolean) StructBoolean {
	return b || b2
}

func (b StructBoolean) Xor(b2 StructBoolean) StructBoolean {
	return b != b2
}

type StructString string

func (s StructString) String() string {
	return string(s)
}

func (s StructString) Append(s2 StructString) StructString {
	return s + s2
}

func (s StructString) Prepend(s2 StructString) StructString {
	return s2 + s
}

func (s StructString) Len() int {
	return len(s)
}

type StructUint uint

func (u StructUint) String() string {
	return fmt.Sprintf("%d", u)
}

func (u StructUint) Add(u2 StructUint) StructUint {
	return u + u2
}

func (u StructUint) Sub(u2 StructUint) StructUint {
	return u - u2
}

func (u StructUint) Mul(u2 StructUint) StructUint {
	return u * u2
}

func (u StructUint) Div(u2 StructUint) StructUint {
	return u / u2
}

type StructComplex complex128

func (c StructComplex) String() string {
	return fmt.Sprintf("%v", complex128(c))
}

func (c StructComplex) Add(c2 StructComplex) StructComplex {
	return c + c2
}

func (c StructComplex) Sub(c2 StructComplex) StructComplex {
	return c - c2
}

func (c StructComplex) Mul(c2 StructComplex) StructComplex {
	return c * c2
}

func (c StructComplex) Div(c2 StructComplex) StructComplex {
	return c / c2
}

type StructChannel[T any] chan T

func (c *StructChannel[T]) String() string {
	var zero T
	return fmt.Sprintf("chan %T", zero)
}

func (c *StructChannel[T]) Add(send T) {
	*c <- send
}

func (c *StructChannel[T]) Sub() T {
	return <-*c
}

func (c *StructChannel[T]) Close() {
	close(*c)
}

type StructSlice[T any] []T

func (s *StructSlice[T]) String() string {
	var zero T
	return fmt.Sprintf("[]%T", zero)
}

func (s *StructSlice[T]) Add(send T) {
	*s = append(*s, send)
}

func (s *StructSlice[T]) Get(i int) T {
	return (*s)[i]
}

type StructMap[K comparable, V any] map[K]V

func (m *StructMap[K, V]) String() string {
	var key K
	var val V
	return fmt.Sprintf("map[%T]%T", key, val)
}

func (m *StructMap[K, V]) Add(key K, val V) {
	(*m)[key] = val
}

func (m *StructMap[K, V]) Get(key K) V {
	return (*m)[key]
}

type StructFunction[T any] func() T

func (f *StructFunction[T]) String() string {
	var zero T
	return fmt.Sprintf("func() %T", zero)
}

func (f *StructFunction[T]) Call() {
	(*f)()
}

type StructValuedNotComparable struct {
	value              int
	notComparableField []string
}

func (v StructValuedNotComparable) Value() int {
	return v.value
}

func (v StructValuedNotComparable) NotComparableField() []string {
	return v.notComparableField
}

// StructSimpleUnsafeCastableBase is a base struct for unsafe casts
// Unsafe casts are used to modify the underlying memory of unsafe pointers
type StructSimpleUnsafeCastableBase struct {
	ID int
}

type StructSimpleUnsafeCastableString struct {
	StructSimpleUnsafeCastableBase
	ValueA string
}

type StructSimpleUnsafeCastableInt struct {
	StructSimpleUnsafeCastableBase
	ValueB int
}

type StructSimpleUnsafeCastableFloat struct {
	StructSimpleUnsafeCastableBase
	ValueC float64
}

type StructSimpleUnsafeCastableBool struct {
	StructSimpleUnsafeCastableBase
	ValueD bool
}

func (a *StructSimpleUnsafeCastableString) Type() string { return "String" }
func (b *StructSimpleUnsafeCastableInt) Type() string    { return "Int" }
func (c *StructSimpleUnsafeCastableFloat) Type() string  { return "Float" }
func (d *StructSimpleUnsafeCastableBool) Type() string   { return "Bool" }

// StructNestedStruct specifications:
//   - size 24 bytes
//   - alignment 8
//   - no padding
//
// A string in Go is composed of an address and a length, both of which are 8 bytes each.
// Therefore, the total size of the string field is 16 bytes, but it aligns to 8 bytes.
type StructNestedStruct struct {
	NestedField1 int    // Size 8 bytes. Offset 0
	NestedField2 string // Size 16 bytes. Offset 8
}

func (a *StructNestedStruct) String() string {
	return fmt.Sprintf("%d %s", a.NestedField1, a.NestedField2)
}

// StructComplexUnsafeCastableBase is a base struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type StructComplexUnsafeCastableBase struct {
	Nested   StructNestedStruct // Size 24 bytes. Offset 0
	ID       int                // Size 8 bytes. Offset 24
	Metadata map[string]any     // Size 16 bytes. Offset 32
	Score    float64            // Size 8 bytes. Offset 48
}

// StructComplexUnsafeCastableOf is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type StructComplexUnsafeCastableOf struct {
	StructComplexUnsafeCastableBase
}

// StructComplexUnsafeCastableDescription is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type StructComplexUnsafeCastableDescription struct {
	StructComplexUnsafeCastableBase        // Size 56 bytes. Offset 0
	Value                           int    // Size 8 bytes. Offset 56
	Description                     string // Size 16 bytes. Offset 64
}

// StructComplexUnsafeCastableB is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type StructComplexUnsafeCastableB struct {
	StructComplexUnsafeCastableBase                // Size 56 bytes. Offset 0
	Score                           float64        // Size 8 bytes. Offset 56
	Metadata                        map[string]any // Size 16 bytes. Offset 64
}

// StructComplexUnsafeCastableC is a struct for unsafe casts
// Specifications:
//   - size 60 bytes
//   - alignment 8
//   - empty padding 0
type StructComplexUnsafeCastableC struct {
	StructComplexUnsafeCastableBase        // Size 56 bytes. Offset 0
	Less                            uint32 // Size 4 bytes. Offset 56
}

// StructComplexUnsafeCastableDeref is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type StructComplexUnsafeCastableDeref struct {
	NestedField1 int            // Size 8 bytes. Offset 0
	NestedField2 string         // Size 16 bytes. Offset 8
	ID           int            // Size 8 bytes. Offset 24
	Metadata     map[string]any // Size 16 bytes. Offset 32
	Score        float64        // Size 8 bytes. Offset 48
}

// StructComplexUnsafeCastableReinterpreted is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type StructComplexUnsafeCastableReinterpreted struct {
	FieldA float64 // Size 8 bytes. Offset 0
	FieldB any     // Size 16 bytes. Offset 8
	FieldC uint    // Size 8 bytes. Offset 24
	FieldD string  // Size 16 bytes. Offset 32
	FieldE int     // Size 8 bytes. Offset 48
}
