// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interf

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"

func NewSimple() Simple {
	return &structs.Simple{}
}

func NewSimpleAsRef() *Simple {
	var s Simple = &structs.Simple{}
	return &s
}

func NewPublic() Public {
	return &structs.Public{}
}

func NewPublicAsRef() *Public {
	var p Public = &structs.Public{}
	return &p
}

func NewOneMethod(name string) OneMethod {
	return structs.OneDataAsRef(name)
}

func NewOneMethodAsRef(name string) *OneMethod {
	var o OneMethod = structs.OneDataAsRef(name)
	return &o
}

func NewTwoMethods(name string, age int) TwoMethods {
	return structs.TwoDataAsRef(name, age)
}

func NewTwoMethodsAsRef(name string, age int) *TwoMethods {
	var t TwoMethods = structs.TwoDataAsRef(name, age)
	return &t
}

func NewAnyStringer(data any) Stringer {
	return structs.StringerAsRef(data)
}

func NewAnyStringerAsRef(data any) *Stringer {
	var s Stringer = structs.StringerAsRef(data)
	return &s
}

func NewStringStringer(data string) Stringer {
	return structs.StringerStringAsRef(data)
}

func NewStringStringerAsRef(data string) *Stringer {
	var s Stringer = structs.StringerStringAsRef(data)
	return &s
}

func NewBytesStringer(data []byte) Stringer {
	return structs.StringerBytesAsRef(data)
}

func NewBytesStringerAsRef(data []byte) *Stringer {
	var s Stringer = structs.StringerBytesAsRef(data)
	return &s
}

func NewCloserSuccess() Closer {
	return structs.CloserSuccessAsRef()
}

func NewCloserSuccessAsRef() *Closer {
	var c Closer = structs.CloserSuccessAsRef()
	return &c
}

func NewCloserError() Closer {
	return structs.CloserErrorAsRef()
}

func NewCloserErrorAsRef() *Closer {
	var c Closer = structs.CloserErrorAsRef()
	return &c
}

func NewCloserReaderSuccess() CloserReader {
	return structs.CloserReaderSuccessAsRef()
}

func NewCloserReaderSuccessAsRef() *CloserReader {
	var c CloserReader = structs.CloserReaderSuccessAsRef()
	return &c
}

func NewCloserReaderError() CloserReader {
	return structs.CloserReaderErrorAsRef()
}

func NewCloserReaderErrorAsRef() *CloserReader {
	var c CloserReader = structs.CloserReaderErrorAsRef()
	return &c
}

func NewCloserReaderWriterSuccess() CloserReaderWriter {
	return structs.CloserReaderWriterSuccessAsRef()
}

func NewCloserReaderWriterSuccessAsRef() *CloserReaderWriter {
	var c CloserReaderWriter = structs.CloserReaderWriterSuccessAsRef()
	return &c
}

func NewCloserReaderWriterError() CloserReaderWriter {
	return structs.CloserReaderWriterErrorAsRef()
}

func NewCloserReaderWriterErrorAsRef() *CloserReaderWriter {
	var c CloserReaderWriter = structs.CloserReaderWriterErrorAsRef()
	return &c
}

func NewIntegerCalc(val int) DataCalc[structs.Integer] {
	return structs.Integer(val)
}

func NewUintCalc(val uint) DataCalc[structs.Uint] {
	return structs.Uint(val)
}

func NewFloatCalc(val float64) DataCalc[structs.Float] {
	return structs.Float(val)
}

func NewComplexCalc(val complex128) DataCalc[structs.Complex] {
	return structs.Complex(val)
}

func NewBoolData(val bool) DataBool[structs.Boolean] {
	return structs.Boolean(val)
}

func NewStringData(val string) DataString[structs.String] {
	return structs.String(val)
}

func NewValuedNotComparable(val int) WithNotComparableField[string] {
	return structs.ValuedNotComparableAsValue(val)
}

func StringTypeCastable(val string) TypeCastable {
	return structs.SimpleUnsafeCastableStringAsRef(val)
}

func FloatTypeCastable(val float64) TypeCastable {
	return structs.SimpleUnsafeCastableFloatAsRef(val)
}

func IntTypeCastable(val int) TypeCastable {
	return structs.SimpleUnsafeCastableIntAsRef(val)
}

func BoolTypeCastable(val bool) TypeCastable {
	return structs.SimpleUnsafeCastableBoolAsRef(val)
}
