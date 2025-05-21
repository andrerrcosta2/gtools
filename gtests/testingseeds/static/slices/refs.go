// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package slices

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
	"unsafe"
)

func primitivesEmptyAsRef() []interface{} {
	return []interface{}{
		&[]bool{}, &[]int{}, &[]int8{}, &[]int16{}, &[]int32{}, &[]int64{}, &[]uint{}, &[]uint8{},
		&[]uint16{}, &[]uint32{}, &[]uint64{}, &[]uintptr{}, &[]float32{}, &[]float64{}, &[]complex64{},
		&[]complex128{}, &[]string{},
	}
}

func primitivesNilAsRef() []interface{} {
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

func pointerToPrimitiveNilAsRef() []interface{} {
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

func ofInterfacesAsRef() []interface{} {
	return []interface{}{
		&[]interf.Stringer{}, &[]interf.Closer{}, &[]interf.CloserReader{}, &[]interf.CloserReaderWriter{},
		&[]interf.WithoutImplementation{}, &[]interf.Error{}, &[]interf.Simple{}, &[]interf.OneMethod{},
		&[]interf.TwoMethods{}, &[]interf.Public{}, &[]interf.Seeder[any]{}, &[]interf.Simple{},
		&[]interf.OneMethod{}, &[]interf.TwoMethods{},
	}
}

func ofPointersToInterfacesAsRef() []interface{} {
	return []interface{}{
		&[]*interf.Stringer{}, &[]*interf.Closer{}, &[]*interf.CloserReader{}, &[]*interf.CloserReaderWriter{},
		&[]*interf.WithoutImplementation{}, &[]*interf.Error{}, &[]*interf.Simple{}, &[]*interf.OneMethod{},
		&[]*interf.TwoMethods{}, &[]*interf.Public{}, &[]*interf.Seeder[any]{}, &[]*interf.Simple{},
		&[]*interf.OneMethod{}, &[]*interf.TwoMethods{},
	}
}

func ofStructsAsRef() []interface{} {
	return []interface{}{
		&[]structs.Boolean{}, &[]structs.Integer{}, &[]structs.String{}, &[]structs.Float{},
		&[]structs.Complex{}, &[]structs.Map[string, int]{}, &[]structs.Slice[any]{},
		&[]structs.Simple{}, &[]structs.Public{}, &[]structs.Channel[any]{},
		&[]structs.CloserSuccess{}, &[]structs.CloserError{}, &[]structs.CloserReaderSuccess{},
		&[]structs.CloserReaderError{}, &[]structs.CloserReaderWriterSuccess{},
		&[]structs.CloserReaderWriterError{}, &[]structs.NotComparable{},
		&[]structs.NaturallyComparable{}, &[]structs.NaturallyComparableWithMethods{},
		&[]structs.Stringer{}, &[]structs.StringerBytes{}, &[]structs.StringerString{},
		&[]structs.OneData{}, &[]structs.TwoData{},
		&[]static.Profile{}, &[]static.Credential{}, &[]static.Product{}, &[]static.Address{},
		&[]static.User{}, &[]static.Account{},
	}
}

func ofPointersToStructsAsRef() []interface{} {
	return []interface{}{
		&[]*structs.Boolean{}, &[]*structs.Integer{}, &[]*structs.String{}, &[]*structs.Float{},
		&[]*structs.Complex{}, &[]*structs.Map[string, int]{}, &[]*structs.Slice[any]{},
		&[]*structs.Simple{}, &[]*structs.Public{}, &[]*structs.Channel[any]{},
		&[]*structs.CloserSuccess{}, &[]*structs.CloserError{}, &[]*structs.CloserReaderSuccess{},
		&[]*structs.CloserReaderError{}, &[]*structs.CloserReaderWriterSuccess{},
		&[]*structs.CloserReaderWriterError{}, &[]*structs.NotComparable{},
		&[]*structs.NaturallyComparable{}, &[]*structs.NaturallyComparableWithMethods{},
		&[]*structs.Stringer{}, &[]*structs.StringerBytes{}, &[]*structs.StringerString{},
		&[]*structs.OneData{}, &[]*structs.TwoData{},
		&[]*static.Profile{}, &[]*static.Credential{}, &[]*static.Product{}, &[]*static.Address{},
		&[]*static.User{}, &[]*static.Account{},
	}
}

func ofChannelsAsRef() []interface{} {
	return []interface{}{
		&[]chan structs.Public{}, &[]chan structs.String{}, &[]chan structs.Simple{}, &[]chan structs.CloserSuccess{},
		&[]chan structs.CloserError{}, &[]chan structs.CloserReaderSuccess{}, &[]chan structs.CloserReaderError{},
		&[]chan structs.CloserReaderWriterSuccess{}, &[]chan structs.CloserReaderWriterError{},
		&[]chan static.Product{}, &[]chan static.Credential{},
		&[]chan bool{}, &[]chan int{}, &[]chan int8{}, &[]chan int16{}, &[]chan int32{}, &[]chan int64{},
		&[]chan uint{}, &[]chan uint8{}, &[]chan uint16{}, &[]chan uint32{}, &[]chan uint64{}, &[]chan uintptr{},
		&[]chan float32{}, &[]chan float64{}, &[]chan complex64{}, &[]chan complex128{}, &[]chan string{},
	}
}

func ofPointersToChannelsAsRef() []interface{} {
	return []interface{}{
		&[]*chan structs.Public{}, &[]*chan structs.String{}, &[]*chan structs.Simple{}, &[]*chan structs.CloserSuccess{},
		&[]*chan structs.CloserError{}, &[]*chan structs.CloserReaderSuccess{}, &[]*chan structs.CloserReaderError{},
		&[]*chan structs.CloserReaderWriterSuccess{}, &[]*chan structs.CloserReaderWriterError{},
		&[]*chan static.Product{}, &[]*chan static.Credential{},
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

func ofRefs() []interface{} {
	return slices.Concat(
		edgeRefs(), ofPointersToChannelsAsRef(), ofChannelsAsRef(), ofPointersToStructsAsRef(),
		ofStructsAsRef(), ofPointersToInterfacesAsRef(), ofInterfacesAsRef(), primitivesNilAsRef(),
		pointerToPrimitiveNilAsRef(), primitivesEmptyAsRef(),
	)
}
