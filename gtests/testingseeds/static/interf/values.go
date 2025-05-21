// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interf

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
	"unsafe"
)

func implementables() []any {
	return []any{
		NewSimple(), NewCloserError(), NewCloserReaderError(), NewCloserReaderWriterError(),
		NewAnyStringer("hello"), NewBytesStringer([]byte("hello")), NewStringStringer("hello"),
		NewOneMethod("hello"), NewPublic(), NewCloserSuccess(),
		NewCloserReaderSuccess(), NewCloserReaderWriterSuccess(), NewTwoMethods("hello", 250),
	}
}

func valueAsImplementables() []any {
	return []any{
		NewIntegerCalc(0), NewUintCalc(0), NewFloatCalc(0),
		NewComplexCalc(0), NewBoolData(false), NewStringData(""),
		NewValuedNotComparable(0),
	}
}

func edgeValues() []any {
	var i1 interface{} = &structs.Empty{}
	var data = 0
	var i2 interface{} = unsafe.Pointer(&data)
	return []any{i1, i2}
}

func ofValues() []interface{} {
	return slices.Concat(implementables(), valueAsImplementables(), edgeValues())
}
