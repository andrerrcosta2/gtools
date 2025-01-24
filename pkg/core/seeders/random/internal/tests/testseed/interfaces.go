// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

//go:build testseed

package testseed

type InterfaceSimple interface {
	Simple() int
}

type InterfacePublic interface {
	PublicMethod() int
}

type InterfaceSeeder[T any] interface {
	Seed() []T
}

type InterfaceOneMethod interface {
	String() string
}

type InterfaceTwoMethods interface {
	String() string
	Int() int
}

type InterfaceStringer interface {
	String() string
}

type InterfaceCloser interface {
	Close() error
}

type InterfaceCloserReader interface {
	InterfaceCloser
	Read([]byte) (int, error)
}

type InterfaceCloserReaderWriter interface {
	InterfaceCloserReader
	Write([]byte) (int, error)
}

type InterfaceWithoutImplementation interface {
	NoImplementation()
}

type InterfaceError interface {
	Error() string
}

type InterfaceDataCalc[T any] interface {
	String() string
	Add(i2 T) T
	Sub(i2 T) T
	Mul(i2 T) T
	Div(i2 T) T
}

type InterfaceDataBool[T any] interface {
	String() string
	And(i2 T) T
	Or(i2 T) T
	Xor(i2 T) T
	Not() T
}

type InterfaceDataString[T any] interface {
	String() string
	Len() int
	Append(s T) T
	Prepend(s T) T
}

type InterfaceWithNotComparableField[T any] interface {
	NotComparableField() []T
}

type InterfaceTypeCastable interface {
	Type() string
}
