// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testseed

func ZeroValues() []any {
	return []any{
		Account{}, Address{},
		ComparableValue{}, ComplexUnsafeCastableBase{},
		ComplexUnsafeCastableOf{}, ComplexUnsafeCastableA{},
		ComplexUnsafeCastableB{}, NestedStruct{}, Profile{},
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
		&Account{}, &Address{},
		&ComparableValue{}, &ComplexUnsafeCastableBase{},
		&ComplexUnsafeCastableOf{}, &ComplexUnsafeCastableA{},
		&ComplexUnsafeCastableB{}, &NestedStruct{}, &Profile{},
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
