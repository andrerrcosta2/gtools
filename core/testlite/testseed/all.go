// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testseed

func ZeroValues() []any {
	return []any{
		ComparableValue{}, ComplexUnsafeCastableBase{},
		ComplexUnsafeCastableOf{}, ComplexUnsafeCastableA{},
		ComplexUnsafeCastableB{}, NestedStruct{},
		SortableValue{}, SortableRef{}, StructCloserError{},
		StructCloserReaderError{}, StructCloserReaderSuccess{},
		StructCloserReaderWriterError{}, StructEmpty{},
		StructNaturallyComparable{}, StructNaturallyComparableWithMethods{},
		StructNestedStruct{}, StructNotComparable{},
		StructNotComparableWithMethods{}, StructOneData{}, StructPublic{},
		StructSimple{}, StructSimpleUnsafeCastableBase{},
		StructSimpleUnsafeCastableBool{}, StructSimpleUnsafeCastableFloat{},
		StructSimpleUnsafeCastableInt{}, StructSimpleUnsafeCastableString{},
		StructStringer{}, StructTwoData{}, StructValuedNotComparable{},
		StructWithPointers{}, StructWithPointersAndUnexportedFields{},
	}
}

func ZeroRef() []any {
	return []any{
		&ComparableValue{}, &ComplexUnsafeCastableBase{},
		&ComplexUnsafeCastableOf{}, &ComplexUnsafeCastableA{},
		&ComplexUnsafeCastableB{}, &NestedStruct{},
		&SortableValue{}, &SortableRef{}, &StructCloserError{},
		&StructCloserReaderError{}, &StructCloserReaderSuccess{},
		&StructCloserReaderWriterError{}, &StructEmpty{},
		&StructNaturallyComparable{}, &StructNaturallyComparableWithMethods{},
		&StructNestedStruct{}, &StructNotComparable{},
		&StructNotComparableWithMethods{}, &StructOneData{}, &StructPublic{},
		&StructSimple{}, &StructSimpleUnsafeCastableBase{},
		&StructSimpleUnsafeCastableBool{}, &StructSimpleUnsafeCastableFloat{},
		&StructSimpleUnsafeCastableInt{}, &StructSimpleUnsafeCastableString{},
		&StructStringer{}, &StructTwoData{}, &StructValuedNotComparable{},
		&StructWithPointers{}, &StructWithPointersAndUnexportedFields{},
	}
}
