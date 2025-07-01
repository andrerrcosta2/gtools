// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testseed

import "errors"

func NilInterfaces() []any {
	return []any{
		InterfaceCloserZeroImpl(), InterfaceCloserReaderZeroImpl(), InterfaceCloserReaderWriterZeroImpl(),
		InterfaceDataBoolZeroImpl(), InterfaceDataCalcZeroImpl(), InterfaceDataStringZeroImpl(),
		InterfaceErrorZeroImpl(), InterfaceOneMethodZeroImpl(), InterfacePublicZeroImpl(), InterfaceSeederZeroImpl(),
		InterfaceSimpleZeroImpl(), InterfaceStringerZeroImpl(), InterfaceTwoMethodsZeroImpl(), InterfaceTypeCastableZeroImpl(),
		InterfaceWithNotComparableFieldZeroImpl(), interface{}(0),
	}
}

func ZeroInterfaces() []any {
	return []any{}
}

func InterfaceCloserZeroImpl() InterfaceCloser {
	return &StructCloserSuccess{}
}

type InterfaceCloser interface {
	Close() error
}

func InterfaceCloserReaderZeroImpl() InterfaceCloserReader {
	return &StructCloserReaderSuccess{}
}

type InterfaceCloserReader interface {
	InterfaceCloser
	Read([]byte) (int, error)
}

func InterfaceCloserReaderWriterZeroImpl() InterfaceCloserReaderWriter {
	return &StructCloserReaderWriterSuccess{}
}

type InterfaceCloserReaderWriter interface {
	InterfaceCloserReader
	Write([]byte) (int, error)
}

func InterfaceDataBoolZeroImpl() InterfaceDataBool[StructBoolean] {
	return StructBoolean(false)
}

type InterfaceDataBool[T any] interface {
	String() string
	And(i2 T) T
	Or(i2 T) T
	Xor(i2 T) T
	Not() T
}

func InterfaceDataCalcZeroImpl() InterfaceDataCalc[StructInteger] {
	return StructInteger(0)
}

type InterfaceDataCalc[T any] interface {
	String() string
	Add(i2 T) T
	Sub(i2 T) T
	Mul(i2 T) T
	Div(i2 T) T
}

func InterfaceDataStringZeroImpl() InterfaceDataString[StructString] {
	return StructString("")
}

type InterfaceDataString[T any] interface {
	String() string
	Len() int
	Append(s T) T
	Prepend(s T) T
}

func InterfaceErrorZeroImpl() InterfaceError {
	return errors.New("")
}

type InterfaceError interface {
	Error() string
}

func InterfaceOneMethodZeroImpl() InterfaceOneMethod {
	return &StructOneData{}
}

type InterfaceOneMethod interface {
	String() string
}

func InterfacePublicZeroImpl() InterfacePublic {
	return &StructPublic{}
}

type InterfacePublic interface {
	PublicMethod() int
}

func InterfaceSeederZeroImpl() InterfaceSeeder[int] {
	return StructIntSeeder{}
}

type InterfaceSeeder[T any] interface {
	Seed() []T
}

func InterfaceSimpleZeroImpl() InterfaceSimple {
	return &StructSimple{}
}

type InterfaceSimple interface {
	Simple() int
}

func InterfaceStringerZeroImpl() InterfaceStringer {
	return &StructStringer{}
}

type InterfaceStringer interface {
	String() string
}

func InterfaceTwoMethodsZeroImpl() InterfaceTwoMethods {
	return &StructTwoData{}
}

type InterfaceTwoMethods interface {
	String() string
	Int() int
}

func InterfaceTypeCastableZeroImpl() InterfaceTypeCastable {
	return &StructSimpleUnsafeCastableBool{}
}

type InterfaceTypeCastable interface {
	Type() string
}

func InterfaceWithNotComparableFieldZeroImpl() InterfaceWithNotComparableField[string] {
	return &StructValuedNotComparable{}
}

type InterfaceWithNotComparableField[T any] interface {
	NotComparableField() []T
}

type InterfaceWithoutImplementation interface {
	NoImplementation()
}
