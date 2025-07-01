// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
	"unsafe"
)

func referencesSet() []interface{} {
	return slices.Concat(
		edgeRefs(), pointerToChannelRefs(), channelRefs(), pointerToStructRefs(),
		structRefs(), pointerToInterfaceRefs(), interfaceRefs(), nilPrimitiveRefs(),
		nilPointerToPrimitiveRefs(), primitiveRefs(),
	)
}

func primitiveRefs() []interface{} {
	return []interface{}{
		&[]bool{}, &[]int{}, &[]int8{}, &[]int16{}, &[]int32{}, &[]int64{}, &[]uint{}, &[]uint8{},
		&[]uint16{}, &[]uint32{}, &[]uint64{}, &[]uintptr{}, &[]float32{}, &[]float64{}, &[]complex64{},
		&[]complex128{}, &[]string{},
	}
}

func nilPrimitiveRefs() []interface{} {
	var s1 []bool
	var s2 []int
	var s3 []int8
	var s4 []int16
	var s5 []int32
	var s6 []int64
	var s7 []uint
	var s8 []uint8
	var s9 []uint16
	var s10 []uint32
	var s11 []uint64
	var s12 []uintptr
	var s13 []float32
	var s14 []float64
	var s15 []complex64
	var s16 []complex128
	var s17 []string
	return []interface{}{
		&s1, &s2, &s3, &s4, &s5, &s6, &s7, &s8, &s9, &s10, &s11, &s12, &s13, &s14, &s15, &s16, &s17,
	}
}

func nilPointerToPrimitiveRefs() []interface{} {
	var s1 *[]bool
	var s2 *[]int
	var s3 *[]int8
	var s4 *[]int16
	var s5 *[]int32
	var s6 *[]int64
	var s7 *[]uint
	var s8 *[]uint8
	var s9 *[]uint16
	var s10 *[]uint32
	var s11 *[]uint64
	var s12 *[]uintptr
	var s13 *[]float32
	var s14 *[]float64
	var s15 *[]complex64
	var s16 *[]complex128
	var s17 *[]string
	return []interface{}{
		&s1, &s2, &s3, &s4, &s5, &s6, &s7, &s8, &s9, &s10, &s11,
		&s12, &s13, &s14, &s15, &s16, &s17,
	}
}

func interfaceRefs() []interface{} {
	return []interface{}{
		&[]interf.Stringer{}, &[]interf.Closer{}, &[]interf.CloserReader{}, &[]interf.CloserReaderWriter{},
		&[]interf.WithoutImplementation{}, &[]interf.Error{}, &[]interf.Simple{}, &[]interf.OneMethod{},
		&[]interf.TwoMethods{}, &[]interf.Public{}, &[]interf.Seeder[any]{}, &[]interf.Simple{},
		&[]interf.OneMethod{}, &[]interf.TwoMethods{},
	}
}

func pointerToInterfaceRefs() []interface{} {
	return []interface{}{
		&[]*interf.Stringer{}, &[]*interf.Closer{}, &[]*interf.CloserReader{}, &[]*interf.CloserReaderWriter{},
		&[]*interf.WithoutImplementation{}, &[]*interf.Error{}, &[]*interf.Simple{}, &[]*interf.OneMethod{},
		&[]*interf.TwoMethods{}, &[]*interf.Public{}, &[]*interf.Seeder[any]{}, &[]*interf.Simple{},
		&[]*interf.OneMethod{}, &[]*interf.TwoMethods{},
	}
}

func structRefs() []interface{} {
	return []interface{}{
		&[]models.Boolean{}, &[]models.Integer{}, &[]models.String{}, &[]models.Float{},
		&[]models.Complex{}, &[]models.Map[string, int]{}, &[]models.Slice[any]{},
		&[]models.Simple{}, &[]models.Public{}, &[]models.Channel[any]{},
		&[]models.CloserSuccess{}, &[]models.CloserError{}, &[]models.CloserReaderSuccess{},
		&[]models.CloserReaderError{}, &[]models.CloserReaderWriterSuccess{},
		&[]models.CloserReaderWriterError{}, &[]models.NotComparable{},
		&[]models.NaturallyComparable{}, &[]models.NaturallyComparableWithMethods{},
		&[]models.Stringer{}, &[]models.StringerBytes{}, &[]models.StringerString{},
		&[]models.OneData{}, &[]models.TwoData{},
		&[]models.Profile{}, &[]models.Credential{}, &[]models.Product{}, &[]models.Address{},
		&[]models.User{}, &[]models.Account{},
	}
}

func pointerToStructRefs() []interface{} {
	return []interface{}{
		&[]*models.Boolean{}, &[]*models.Integer{}, &[]*models.String{}, &[]*models.Float{},
		&[]*models.Complex{}, &[]*models.Map[string, int]{}, &[]*models.Slice[any]{},
		&[]*models.Simple{}, &[]*models.Public{}, &[]*models.Channel[any]{},
		&[]*models.CloserSuccess{}, &[]*models.CloserError{}, &[]*models.CloserReaderSuccess{},
		&[]*models.CloserReaderError{}, &[]*models.CloserReaderWriterSuccess{},
		&[]*models.CloserReaderWriterError{}, &[]*models.NotComparable{},
		&[]*models.NaturallyComparable{}, &[]*models.NaturallyComparableWithMethods{},
		&[]*models.Stringer{}, &[]*models.StringerBytes{}, &[]*models.StringerString{},
		&[]*models.OneData{}, &[]*models.TwoData{},
		&[]*models.Profile{}, &[]*models.Credential{}, &[]*models.Product{}, &[]*models.Address{},
		&[]*models.User{}, &[]*models.Account{},
	}
}

func channelRefs() []interface{} {
	return []interface{}{
		&[]chan models.Public{}, &[]chan models.String{}, &[]chan models.Simple{}, &[]chan models.CloserSuccess{},
		&[]chan models.CloserError{}, &[]chan models.CloserReaderSuccess{}, &[]chan models.CloserReaderError{},
		&[]chan models.CloserReaderWriterSuccess{}, &[]chan models.CloserReaderWriterError{},
		&[]chan models.Product{}, &[]chan models.Credential{},
		&[]chan bool{}, &[]chan int{}, &[]chan int8{}, &[]chan int16{}, &[]chan int32{}, &[]chan int64{},
		&[]chan uint{}, &[]chan uint8{}, &[]chan uint16{}, &[]chan uint32{}, &[]chan uint64{}, &[]chan uintptr{},
		&[]chan float32{}, &[]chan float64{}, &[]chan complex64{}, &[]chan complex128{}, &[]chan string{},
	}
}

func pointerToChannelRefs() []interface{} {
	return []interface{}{
		&[]*chan models.Public{}, &[]*chan models.String{}, &[]*chan models.Simple{}, &[]*chan models.CloserSuccess{},
		&[]*chan models.CloserError{}, &[]*chan models.CloserReaderSuccess{}, &[]*chan models.CloserReaderError{},
		&[]*chan models.CloserReaderWriterSuccess{}, &[]*chan models.CloserReaderWriterError{},
		&[]*chan models.Product{}, &[]*chan models.Credential{},
		&[]*chan bool{}, &[]*chan int{}, &[]*chan int8{}, &[]*chan int16{}, &[]*chan int32{}, &[]*chan int64{},
		&[]*chan uint{}, &[]*chan uint8{}, &[]*chan uint16{}, &[]*chan uint32{}, &[]*chan uint64{}, &[]*chan uintptr{},
		&[]*chan float32{}, &[]*chan float64{}, &[]*chan complex64{}, &[]*chan complex128{}, &[]*chan string{},
	}
}

func edgeRefs() []interface{} {
	return []interface{}{
		&[]**interface{}{},
		&[]unsafe.Pointer{}, &[]*unsafe.Pointer{}, &[]**unsafe.Pointer{}, &[][][]*chan **unsafe.Pointer{},
	}
}
