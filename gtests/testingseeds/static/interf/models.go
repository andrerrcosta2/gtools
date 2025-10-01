// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interf

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"

type Simple interface {
	Simple() int
}

type Public interface {
	PublicMethod() int
}

type Seeder[T any] interface {
	Seed() []T
}

func OneDataAsRef(s string) OneMethod {
	return models.OneDataAsRef(s)
}

type OneMethod interface {
	String() string
}

func TwoDataAsTwoMethodsRef(s string, i int) TwoMethods {
	return models.TwoDataAsRef(s, i)
}

func ThreeDataAsTwoMethodsRef(s string, i int, f float64) TwoMethods {
	return models.ThreeDataAsRef(s, i, f)
}

type TwoMethods interface {
	String() string
	Int() int
}

func ThreeDataAsThreeMethodsRef(s string, i int, f float64) ThreeMethods {
	return models.ThreeDataAsRef(s, i, f)
}

type ThreeMethods interface {
	String() string
	Int() int
	Float() float64
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

func FloatTypeCastable(id int, v float64) TypeCastable {
	return models.SimpleUnsafeCastableFloatAsRef(id, v)
}

func BoolTypeCastable(id int, v bool) TypeCastable {
	return models.SimpleUnsafeCastableBoolAsRef(id, v)
}
