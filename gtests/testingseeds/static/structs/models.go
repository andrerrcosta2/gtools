// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

import (
	"errors"
	"fmt"
	"math/rand"
)

type Empty struct{}

type Simple struct{}

func (s *Simple) Simple() int {
	return rand.Intn(1000)
}

type Public struct{}

func (p *Public) PublicMethod() int {
	return rand.Intn(1000)
}

type OneData struct {
	Name string
}

func (o *OneData) String() string {
	return o.Name
}

type TwoData struct {
	Name string
	Age  int
}

func (t *TwoData) String() string {
	return t.Name
}

func (t *TwoData) Int() int {
	return t.Age
}

type Stringer struct {
	data any
}

func (s *Stringer) String() string {
	return fmt.Sprintf("%v", s.data)
}

type StringerString struct{ data string }

func (s *StringerString) String() string {
	return s.data
}

type StringerBytes struct{ data []byte }

func (s *StringerBytes) String() string {
	return string(s.data)
}

type CloserSuccess struct{}

func (c *CloserSuccess) Close() error { return nil }

type CloserError struct{}

func (c *CloserError) Close() error { return errors.New("error") }

type CloserReaderSuccess struct {
	data []byte
}

func (c *CloserReaderSuccess) Close() error { return nil }

func (c *CloserReaderSuccess) Read(data []byte) (int, error) {
	c.data = data
	return len(c.data), nil
}

type CloserReaderError struct {
	data []byte
}

func (c *CloserReaderError) Close() error { return errors.New("error closing") }

func (c *CloserReaderError) Read(data []byte) (int, error) {
	return len(c.data), errors.New("error reading")
}

type CloserReaderWriterSuccess struct {
	read  []byte
	write []byte
}

func (c *CloserReaderWriterSuccess) Close() error { return nil }

func (c *CloserReaderWriterSuccess) Read(data []byte) (int, error) {
	c.read = data
	return len(c.read), nil
}

func (c *CloserReaderWriterSuccess) Write(data []byte) (int, error) {
	c.write = data
	return len(c.write), nil
}

type CloserReaderWriterError struct {
	read  []byte
	write []byte
}

func (c *CloserReaderWriterError) Close() error { return errors.New("error closing") }

func (c *CloserReaderWriterError) Read(data []byte) (int, error) {
	return len(c.read), errors.New("error reading")
}

func (c *CloserReaderWriterError) Write(data []byte) (int, error) {
	return len(c.write), errors.New("error writing")
}

type NaturallyComparable struct {
	ComparableValueA int
	ComparableValueB string
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

type NotComparable struct {
	ComparableValueA    int
	NotComparableValueB []string
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

type SimpleNode struct {
	Next  *SimpleNode
	Value any
}

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Value any
}

type NTreeNode struct {
	Nodes map[string]*NTreeNode
}

type LinkedList struct {
	Head *BinaryNode
	Tail *BinaryNode
}

func (l *LinkedList) Add(value any) {
	newNode := &BinaryNode{Value: value}
	
	if l.Tail != nil {
		l.Tail.Right = nil
		l.Head.Left = nil
	}
	if l.Tail != nil {
		l.Tail.Right = newNode
		newNode.Left = l.Tail
	} else {
		l.Head = newNode
	}
	l.Tail = newNode
}

func (l *LinkedList) Remove(value any) {
	if l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		if current.Value == value {
			if current == l.Head && current == l.Tail {
				l.Head, l.Tail = nil, nil
				return
			}
			if current == l.Head {
				l.Head = l.Head.Right
				l.Head.Left = nil
			}
			if current == l.Tail {
				l.Tail = l.Tail.Left
				l.Tail.Right = nil
			}
			if current.Left != nil {
				current.Left.Right = current.Right
			}
			if current.Right != nil {
				current.Right.Left = current.Left
			}
			return
		}
		current = current.Right
	}
}

type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

func (i Integer) Add(i2 Integer) Integer {
	return i + i2
}

func (i Integer) Sub(i2 Integer) Integer {
	return i - i2
}

func (i Integer) Mul(i2 Integer) Integer {
	return i * i2
}

func (i Integer) Div(i2 Integer) Integer {
	return i / i2
}

type Float float64

func (f Float) String() string {
	return fmt.Sprintf("%f", f)
}

func (f Float) Add(f2 Float) Float {
	return f + f2
}

func (f Float) Sub(f2 Float) Float {
	return f - f2
}

func (f Float) Mul(f2 Float) Float {
	return f * f2
}

func (f Float) Div(f2 Float) Float {
	return f / f2
}

type Boolean bool

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b)
}

func (b Boolean) Not() Boolean {
	return !b
}

func (b Boolean) And(b2 Boolean) Boolean {
	return b && b2
}

func (b Boolean) Or(b2 Boolean) Boolean {
	return b || b2
}

func (b Boolean) Xor(b2 Boolean) Boolean {
	return b != b2
}

type String string

func (s String) String() string {
	return string(s)
}

func (s String) Append(s2 String) String {
	return s + s2
}

func (s String) Prepend(s2 String) String {
	return s2 + s
}

func (s String) Len() int {
	return len(s)
}

type Uint uint

func (u Uint) String() string {
	return fmt.Sprintf("%d", u)
}

func (u Uint) Add(u2 Uint) Uint {
	return u + u2
}

func (u Uint) Sub(u2 Uint) Uint {
	return u - u2
}

func (u Uint) Mul(u2 Uint) Uint {
	return u * u2
}

func (u Uint) Div(u2 Uint) Uint {
	return u / u2
}

type Complex complex128

func (c Complex) String() string {
	return fmt.Sprintf("%v", c)
}

func (c Complex) Add(c2 Complex) Complex {
	return c + c2
}

func (c Complex) Sub(c2 Complex) Complex {
	return c - c2
}

func (c Complex) Mul(c2 Complex) Complex {
	return c * c2
}

func (c Complex) Div(c2 Complex) Complex {
	return c / c2
}

type Channel[T any] chan T

func (c *Channel[T]) String() string {
	var zero T
	return fmt.Sprintf("chan %T", zero)
}

func (c *Channel[T]) Add(send T) {
	*c <- send
}

func (c *Channel[T]) Sub() T {
	return <-*c
}

func (c *Channel[T]) Close() {
	close(*c)
}

type Slice[T any] []T

func (s *Slice[T]) String() string {
	var zero T
	return fmt.Sprintf("[]%T", zero)
}

func (s *Slice[T]) Add(send T) {
	*s = append(*s, send)
}

func (s *Slice[T]) Get(i int) T {
	return (*s)[i]
}

type Map[K comparable, V any] map[K]V

func (m *Map[K, V]) String() string {
	var key K
	var val V
	return fmt.Sprintf("map[%T]%T", key, val)
}

func (m *Map[K, V]) Add(key K, val V) {
	(*m)[key] = val
}

func (m *Map[K, V]) Get(key K) V {
	return (*m)[key]
}

type Function[T any] func() T

func (f *Function[T]) String() string {
	var zero T
	return fmt.Sprintf("func() %T", zero)
}

func (f *Function[T]) Call() {
	(*f)()
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

// SimpleUnsafeCastableBase is a base struct for unsafe casts
// Unsafe casts are used to modify the underlying memory of unsafe pointers
type SimpleUnsafeCastableBase struct {
	ID int
}

type SimpleUnsafeCastableString struct {
	SimpleUnsafeCastableBase
	ValueA string
}

type SimpleUnsafeCastableInt struct {
	SimpleUnsafeCastableBase
	ValueB int
}

type SimpleUnsafeCastableFloat struct {
	SimpleUnsafeCastableBase
	ValueC float64
}

type SimpleUnsafeCastableBool struct {
	SimpleUnsafeCastableBase
	ValueD bool
}

func (a *SimpleUnsafeCastableString) Type() string { return "String" }
func (b *SimpleUnsafeCastableInt) Type() string    { return "Int" }
func (c *SimpleUnsafeCastableFloat) Type() string  { return "Float" }
func (d *SimpleUnsafeCastableBool) Type() string   { return "Bool" }

// NestedStruct specifications:
//   - size 24 bytes
//   - alignment 8
//   - no padding
//
// A string in Go is composed of an address and a length, both of which are 8 bytes each.
// Therefore, the total size of the string field is 16 bytes, but it aligns to 8 bytes.
type NestedStruct struct {
	NestedField1 int    // Size 8 bytes. Offset 0
	NestedField2 string // Size 16 bytes. Offset 8
}

func (a *NestedStruct) String() string {
	return fmt.Sprintf("%d %s", a.NestedField1, a.NestedField2)
}

// ComplexUnsafeCastableBase is a base struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableBase struct {
	Nested   NestedStruct   // Size 24 bytes. Offset 0
	ID       int            // Size 8 bytes. Offset 24
	Metadata map[string]any // Size 16 bytes. Offset 32
	Score    float64        // Size 8 bytes. Offset 48
}

// ComplexUnsafeCastableOf is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableOf struct {
	ComplexUnsafeCastableBase
}

// ComplexUnsafeCastableDescription is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableDescription struct {
	ComplexUnsafeCastableBase        // Size 56 bytes. Offset 0
	Value                     int    // Size 8 bytes. Offset 56
	Description               string // Size 16 bytes. Offset 64
}

// ComplexUnsafeCastableB is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableB struct {
	ComplexUnsafeCastableBase                // Size 56 bytes. Offset 0
	Score                     float64        // Size 8 bytes. Offset 56
	Metadata                  map[string]any // Size 16 bytes. Offset 64
}

// ComplexUnsafeCastableC is a struct for unsafe casts
// Specifications:
//   - size 60 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableC struct {
	ComplexUnsafeCastableBase        // Size 56 bytes. Offset 0
	Less                      uint32 // Size 4 bytes. Offset 56
}

// ComplexUnsafeCastableDeref is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableDeref struct {
	NestedField1 int            // Size 8 bytes. Offset 0
	NestedField2 string         // Size 16 bytes. Offset 8
	ID           int            // Size 8 bytes. Offset 24
	Metadata     map[string]any // Size 16 bytes. Offset 32
	Score        float64        // Size 8 bytes. Offset 48
}

// ComplexUnsafeCastableReinterpreted is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableReinterpreted struct {
	FieldA float64 // Size 8 bytes. Offset 0
	FieldB any     // Size 16 bytes. Offset 8
	FieldC uint    // Size 8 bytes. Offset 24
	FieldD String  // Size 16 bytes. Offset 32
	FieldE int     // Size 8 bytes. Offset 48
}
