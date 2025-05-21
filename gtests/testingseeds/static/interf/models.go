// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interf

type Simple interface {
	Simple() int
}

type Public interface {
	PublicMethod() int
}

type Seeder[T any] interface {
	Seed() []T
}

type OneMethod interface {
	String() string
}

type TwoMethods interface {
	String() string
	Int() int
}

type Stringer interface {
	String() string
}

type Closer interface {
	Close() error
}

type CloserReader interface {
	Closer
	Read([]byte) (int, error)
}

type CloserReaderWriter interface {
	CloserReader
	Write([]byte) (int, error)
}

type WithoutImplementation interface {
	NoImplementation()
}

type Error interface {
	Error() string
}

type DataCalc[T any] interface {
	String() string
	Add(i2 T) T
	Sub(i2 T) T
	Mul(i2 T) T
	Div(i2 T) T
}

type DataBool[T any] interface {
	String() string
	And(i2 T) T
	Or(i2 T) T
	Xor(i2 T) T
	Not() T
}

type DataString[T any] interface {
	String() string
	Len() int
	Append(s T) T
	Prepend(s T) T
}

type WithNotComparableField[T any] interface {
	NotComparableField() []T
}

type TypeCastable interface {
	Type() string
}
