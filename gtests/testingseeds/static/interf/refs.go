// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interf

import (
	"slices"
	"unsafe"
)

func implementablesOfRefs() []interface{} {
	return []interface{}{
		NewSimpleAsRef(), NewCloserErrorAsRef(), NewCloserReaderErrorAsRef(),
		NewCloserReaderWriterErrorAsRef(), NewAnyStringerAsRef("hello"),
		NewBytesStringerAsRef([]byte("hello")), NewStringStringerAsRef("hello"),
		NewOneMethodAsRef("hello"), NewTwoMethodsAsRef("hello", 250),
		NewOneMethodAsRef("hello"), NewTwoMethodsAsRef("hello", 250),
		NewPublicAsRef(), NewCloserSuccessAsRef(), NewCloserReaderSuccessAsRef(),
		NewCloserReaderWriterSuccessAsRef(),
	}
}

func valueAsImplementablesAsRefs() []any {
	i1 := NewIntegerCalc(0)
	i2 := NewUintCalc(0)
	i3 := NewFloatCalc(0)
	i4 := NewComplexCalc(0)
	i5 := NewBoolData(false)
	i6 := NewStringData("")
	i7 := NewValuedNotComparable(0)
	return []any{
		&i1, &i2, &i3, &i4, &i5, &i6, &i7,
	}

}

func edgesAsRefs() []interface{} {
	var i1 *interface{}
	var x = "data"
	var i2 interface{} = unsafe.Pointer(&x)
	return []interface{}{
		i1, i2,
	}
}

func ofRefs() []interface{} {
	return slices.Concat(implementablesOfRefs(), valueAsImplementablesAsRefs(),
		edgesAsRefs(),
	)
}
