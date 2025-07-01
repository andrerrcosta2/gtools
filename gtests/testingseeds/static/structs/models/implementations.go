// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"reflect"
)

var BooleanZeroInst = new(Boolean)

func BoolAsValue() Boolean {
	return Boolean{}
}

func BoolAsRef() *Boolean {
	return new(Boolean)
}

func BooleanAsRandRef() *Boolean {
	b := Boolean{random.SingleOf[bool]()}
	return &b
}

func BooleanAsRandValue() Boolean {
	return Boolean{random.SingleOf[bool]()}
}

type Boolean struct {
	value bool
}

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b.value)
}

func (b Boolean) Not() Boolean {
	return Boolean{!b.value}
}

func (b Boolean) And(b2 Boolean) Boolean {
	return Boolean{b.value && b2.value}
}

func (b Boolean) Or(b2 Boolean) Boolean {
	return Boolean{b.value || b2.value}
}

func (b Boolean) Xor(b2 Boolean) Boolean {
	return Boolean{b.value != b2.value}
}

var ChannelZeroInst = new(Channel[any])

func ChannelAsValue[T any]() Channel[T] {
	return Channel[T]{
		flow: make(chan T),
	}
}

func ChannelAsRef[T any]() *Channel[T] {
	return &Channel[T]{
		flow: make(chan T),
	}
}

func ChannelAsRandRef[T any]() *Channel[T] {
	return &Channel[T]{random.SingleOf[chan T]()}
}

func ChannelAsRandValue[T any]() Channel[T] {
	c := Channel[T]{random.SingleOf[chan T]()}
	return c
}

type Channel[T any] struct {
	flow chan T
}

func (c *Channel[T]) String() string {
	var zero T
	return fmt.Sprintf("chan %T", zero)
}

func (c *Channel[T]) Add(send T) {
	c.flow <- send
}

func (c *Channel[T]) Sub() T {
	return <-c.flow
}

func (c *Channel[T]) Close() {
	close(c.flow)
}

var CloserErrorZeroInst = new(CloserError)

func CloserErrorAsRef() *CloserError {
	return &CloserError{}
}

func CloserErrorAsValue() CloserError {
	return CloserError{}
}

func CloserErrorAsRandRef() *CloserError {
	return &CloserError{}
}

func CloserErrorAsRandValue() CloserError {
	return CloserError{}
}

type CloserError struct{}

func (c *CloserError) Close() error { return errors.New("error") }

var CloserReaderErrorZeroInst = new(CloserReaderError)

func CloserReaderErrorAsRef() *CloserReaderError {
	return &CloserReaderError{}
}

func CloserReaderErrorAsValue() CloserReaderError {
	return CloserReaderError{}
}

func CloserReaderErrorAsRandRef() *CloserReaderError {
	return &CloserReaderError{
		data: random.Bytes(1, 0, 100).At(0),
	}
}

func CloserReaderErrorAsRandValue() CloserReaderError {
	return CloserReaderError{
		data: random.Bytes(1, 0, 100).At(0),
	}
}

type CloserReaderError struct {
	data []byte
}

func (c *CloserReaderError) Close() error { return errors.New("error closing") }

func (c *CloserReaderError) Read(_ []byte) (int, error) {
	return len(c.data), errors.New("error reading")
}

var CloserReaderSuccessZeroInst = new(CloserReaderSuccess)

func CloserReaderSuccessAsRef() *CloserReaderSuccess {
	return &CloserReaderSuccess{}
}

func CloserReaderSuccessAsValue() CloserReaderSuccess {
	return CloserReaderSuccess{}
}

func CloserReaderSuccessAsRandRef() *CloserReaderSuccess {
	return &CloserReaderSuccess{
		data: random.Bytes(1, 0, 100).At(0),
	}
}

func CloserReaderSuccessAsRandValue() CloserReaderSuccess {
	return CloserReaderSuccess{
		data: random.Bytes(1, 0, 100).At(0),
	}
}

type CloserReaderSuccess struct {
	data []byte
}

func (c *CloserReaderSuccess) Close() error { return nil }

func (c *CloserReaderSuccess) Read(data []byte) (int, error) {
	c.data = data
	return len(c.data), nil
}

var CloserReaderWriterErrorZeroInst = new(CloserReaderWriterError)

func CloserReaderWriterErrorAsRef() *CloserReaderWriterError {
	return &CloserReaderWriterError{}
}

func CloserReaderWriterErrorAsValue() CloserReaderWriterError {
	return CloserReaderWriterError{}
}

func CloserReaderWriterErrorAsRandRef() *CloserReaderWriterError {
	return &CloserReaderWriterError{
		read:  random.Bytes(1, 0, 100).At(0),
		write: random.Bytes(1, 0, 100).At(0),
	}
}

func CloserReaderWriterErrorAsRandValue() CloserReaderWriterError {
	return CloserReaderWriterError{
		read:  random.Bytes(1, 0, 100).At(0),
		write: random.Bytes(1, 0, 100).At(0),
	}
}

type CloserReaderWriterError struct {
	read  []byte
	write []byte
}

func (c *CloserReaderWriterError) Close() error { return errors.New("error closing") }

func (c *CloserReaderWriterError) Read(_ []byte) (int, error) {
	return len(c.read), errors.New("error reading")
}

func (c *CloserReaderWriterError) Write(_ []byte) (int, error) {
	return len(c.write), errors.New("error writing")
}

var CloserReaderWriterSuccessZeroInst = new(CloserReaderWriterSuccess)

func CloserReaderWriterSuccessAsRef() *CloserReaderWriterSuccess {
	return &CloserReaderWriterSuccess{}
}

func CloserReaderWriterSuccessAsValue() CloserReaderWriterSuccess {
	return CloserReaderWriterSuccess{}
}

func CloserReaderWriterSuccessAsRandRef() *CloserReaderWriterSuccess {
	return &CloserReaderWriterSuccess{
		read:  random.Bytes(1, 0, 100).At(0),
		write: random.Bytes(1, 0, 100).At(0),
	}
}

func CloserReaderWriterSuccessAsRandValue() CloserReaderWriterSuccess {
	return CloserReaderWriterSuccess{
		read:  random.Bytes(1, 0, 100).At(0),
		write: random.Bytes(1, 0, 100).At(0),
	}
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

var CloserSuccessZeroInst = new(CloserSuccess)

func CloserSuccessAsRef() *CloserSuccess {
	return &CloserSuccess{}
}

func CloserSuccessAsValue() CloserSuccess {
	return CloserSuccess{}
}

func CloserSuccessAsRandRef() *CloserSuccess {
	return &CloserSuccess{}
}

func CloserSuccessAsRandValue() CloserSuccess {
	return CloserSuccess{}
}

type CloserSuccess struct{}

func (c *CloserSuccess) Close() error { return nil }

var ComplexZeroInst = new(Complex)

func ComplexAsValue() Complex {
	return Complex{}
}

func ComplexAsRef() *Complex {
	return new(Complex)
}

func ComplexAsRandRef() *Complex {
	return &Complex{random.SingleOf[complex128]()}
}

func ComplexAsRandValue() Complex {
	return Complex{random.SingleOf[complex128]()}
}

type Complex struct {
	value complex128
}

func (c Complex) String() string {
	return fmt.Sprintf("%v", c.value)
}

func (c Complex) Add(c2 Complex) Complex {
	return Complex{c.value + c2.value}
}

func (c Complex) Sub(c2 Complex) Complex {
	return Complex{c.value - c2.value}
}

func (c Complex) Mul(c2 Complex) Complex {
	return Complex{c.value * c2.value}
}

func (c Complex) Div(c2 Complex) Complex {
	return Complex{c.value / c2.value}
}

var FloatZeroInst = new(Float)

func FloatAsValue() Float {
	return Float{}
}

func FloatAsRef() *Float {
	return new(Float)
}

func FloatAsRandRef() *Float {
	return &Float{random.SingleOf[float64]()}
}

func FloatAsRandValue() Float {
	return Float{random.SingleOf[float64]()}
}

type Float struct {
	value float64
}

func (f Float) String() string {
	return fmt.Sprintf("%f", f.value)
}

func (f Float) Add(f2 Float) Float {
	return Float{f.value + f2.value}
}

func (f Float) Sub(f2 Float) Float {
	return Float{f.value - f2.value}
}

func (f Float) Mul(f2 Float) Float {
	return Float{f.value * f2.value}
}

func (f Float) Div(f2 Float) Float {
	return Float{f.value / f2.value}
}

var IntegerZeroInst = new(Integer)

func IntAsValue() Integer {
	return Integer{}
}

func IntAsRef() *Integer {
	return new(Integer)
}

func IntegerAsRandRef() *Integer {
	return &Integer{random.SingleOf[int]()}
}

func IntegerAsRandValue() Integer {
	return Integer{random.SingleOf[int]()}
}

type Integer struct {
	value int
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

func (i Integer) Add(i2 Integer) Integer {
	return Integer{i.value + i2.value}
}

func (i Integer) Sub(i2 Integer) Integer {
	return Integer{i.value - i2.value}
}

func (i Integer) Mul(i2 Integer) Integer {
	return Integer{i.value * i2.value}
}

func (i Integer) Div(i2 Integer) Integer {
	return Integer{i.value / i2.value}
}

var MapZeroInst = new(Map[string, any])

func MapAsValue[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		data: make(map[K]V),
	}
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V),
	}
}

func MapAsRandRef[K comparable, V any]() *Map[K, V] {
	size := random.Int(1, 0, 10).At(0)
	m := MapAsValue[K, V]()
	for i := 0; i < size; i++ {
		m.Add(random.SingleOf[K](), random.SingleOf[V]())
	}
	return &m
}

func MapAsRandValue[K comparable, V any]() Map[K, V] {
	size := random.Int(1, 0, 10).At(0)
	m := MapAsValue[K, V]()
	for i := 0; i < size; i++ {
		m.Add(random.SingleOf[K](), random.SingleOf[V]())
	}
	return m
}

type Map[K comparable, V any] struct {
	data map[K]V
}

func (m *Map[K, V]) String() string {
	var key K
	var val V
	return fmt.Sprintf("map[%T]%T%v", key, val, m.data)
}

func (m *Map[K, V]) Add(key K, val V) {
	m.data[key] = val
}

func (m *Map[K, V]) Get(key K) V {
	return m.data[key]
}

var SliceZeroInst = new(Slice[any])

func SliceAsValue[T any]() Slice[T] {
	return Slice[T]{}
}

func SliceAsRef[T any]() *Slice[T] {
	return new(Slice[T])
}

func SliceAsRandRef[T any]() *Slice[T] {
	size := random.Int(1, 0, 10).At(0)
	return &Slice[T]{random.Of[T](size).Values()}
}

func SliceAsRandValue[T any]() Slice[T] {
	size := random.Int(1, 0, 10).At(0)
	return Slice[T]{random.Of[T](size).Values()}
}

type Slice[T any] struct {
	data []T
}

func (s *Slice[T]) String() string {
	var zero T
	t := reflect.TypeOf(zero).String()
	return fmt.Sprintf("[]%s%v", t, s.data)
}

func (s *Slice[T]) Add(send T) {
	s.data = append(s.data, send)
}

func (s *Slice[T]) Get(i int) T {
	return s.data[i]
}

var StringZeroInst = new(String)

func StringAsValue() String {
	return String{}
}

func StringAsRef() *String {
	return new(String)
}

func StringAsRandRef() *String {
	return &String{random.SingleOf[string]()}
}

func StringAsRandValue() String {
	return String{random.SingleOf[string]()}
}

type String struct {
	value string
}

func (s String) String() string {
	return s.value
}

func (s String) Append(s2 String) String {
	return String{s.value + s2.value}
}

func (s String) Prepend(s2 String) String {
	return String{s2.value + s.value}
}

func (s String) Len() int {
	return len(s.value)
}

var StringerZeroInst = new(Stringer)

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

func StringerAsRandRef() *Stringer {
	return &Stringer{
		data: random.SingleOf[any](),
	}
}

func StringerAsRandValue() Stringer {
	return Stringer{
		data: random.SingleOf[any](),
	}
}

type Stringer struct {
	data any
}

func (s *Stringer) String() string {
	return fmt.Sprintf("%v", s.data)
}

var StringerBytesZeroInst = new(StringerBytes)

func StringerBytesAsValue(data []byte) StringerBytes {
	return StringerBytes{data}
}

func StringerBytesAsRef(data []byte) *StringerBytes {
	return &StringerBytes{data}
}

func StringerBytesAsRandRef() *StringerBytes {
	return &StringerBytes{
		data: random.SingleOf[[]byte](),
	}
}

func StringerBytesAsRandValue() StringerBytes {
	return StringerBytes{
		data: random.SingleOf[[]byte](),
	}
}

type StringerBytes struct{ data []byte }

func (s *StringerBytes) String() string {
	return string(s.data)
}

var StringerStringZeroInst = new(StringerString)

func StringerStringAsValue(data string) StringerString {
	return StringerString{data}
}

func StringerStringAsRef(data string) *StringerString {
	return &StringerString{data}
}

func StringerStringAsRandRef() *StringerString {
	return &StringerString{
		data: random.SingleOf[string](),
	}
}

func StringerStringAsRandValue() StringerString {
	return StringerString{
		data: random.SingleOf[string](),
	}
}

type StringerString struct{ data string }

func (s *StringerString) String() string {
	return s.data
}

var TwoDataZeroInst = new(TwoData)

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

func TwoDataAsRandRef() *TwoData {
	return &TwoData{
		Name: random.SingleOf[string](),
		Age:  random.Int(1, 0, 100).At(0),
	}
}

func TwoDataAsRandValue() TwoData {
	return TwoData{
		Name: random.SingleOf[string](),
		Age:  random.Int(1, 0, 100).At(0),
	}
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

var UintZeroInst = new(Uint)

func UintAsValue() Uint {
	return Uint{}
}

func UintAsRef() *Uint {
	return new(Uint)
}

func UintAsRandRef() *Uint {
	return &Uint{random.SingleOf[uint]()}
}

func UintAsRandValue() Uint {
	return Uint{random.SingleOf[uint]()}
}

type Uint struct {
	value uint
}

func (u Uint) String() string {
	return fmt.Sprintf("%d", u.value)
}

func (u Uint) Add(u2 Uint) Uint {
	return Uint{u.value + u2.value}
}

func (u Uint) Sub(u2 Uint) Uint {
	return Uint{u.value - u2.value}

}

func (u Uint) Mul(u2 Uint) Uint {
	return Uint{u.value * u2.value}

}

func (u Uint) Div(u2 Uint) Uint {
	return Uint{u.value / u2.value}

}
