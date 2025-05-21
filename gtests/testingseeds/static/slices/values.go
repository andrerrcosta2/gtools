// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package slices

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
	"unsafe"
)

func primitivesEmpty() []interface{} {
	return []interface{}{
		[]bool{}, []int{}, []int8{}, []int16{}, []int32{}, []int64{}, []uint{}, []uint8{},
		[]uint16{}, []uint32{}, []uint64{}, []uintptr{}, []float32{}, []float64{}, []complex64{},
		[]complex128{}, []string{},
	}
}

func primitivesNil() []interface{} {
	var bv []bool
	var iv []int
	var i8v []int8
	var i16v []int16
	var i32v []int32
	var i64v []int64
	var uiv []uint
	var u8v []uint8
	var u16v []uint16
	var u32v []uint32
	var u64v []uint64
	var uptrv []uintptr
	var f32v []float32
	var f64v []float64
	var c64v []complex64
	var c128v []complex128
	var sv []string
	return []interface{}{bv, iv, i8v, i16v, i32v, i64v, uiv, u8v, u16v, u32v, u64v, uptrv, f32v, f64v, c64v, c128v, sv}
}

func pointerToPrimitiveNil() []interface{} {
	var bv []*bool
	var iv []*int
	var i8v []*int8
	var i16v []*int16
	var i32v []*int32
	var i64v []*int64
	var uiv []*uint
	var u8v []*uint8
	var u16v []*uint16
	var u32v []*uint32
	var u64v []*uint64
	var uptrv []*uintptr
	var f32v []*float32
	var f64v []*float64
	var c64v []*complex64
	var c128v []*complex128
	var sv []*string
	return []interface{}{&bv, &iv, &i8v, &i16v, &i32v, &i64v, &uiv, &u8v, &u16v, &u32v, &u64v, &uptrv, &f32v, &f64v, &c64v, &c128v, &sv}
}

func ofInterfaces() []interface{} {
	return []interface{}{
		[]interf.Stringer{}, []interf.Closer{}, []interf.CloserReader{}, []interf.CloserReaderWriter{},
		[]interf.WithoutImplementation{}, []interf.Error{}, []interf.Simple{}, []interf.OneMethod{},
		[]interf.TwoMethods{}, []interf.Public{}, []interf.Seeder[any]{}, []interf.Simple{},
		[]interf.OneMethod{}, []interf.TwoMethods{},
	}
}

func ofPointersToInterfaces() []interface{} {
	return []interface{}{
		[]*interf.Stringer{}, []*interf.Closer{}, []*interf.CloserReader{}, []*interf.CloserReaderWriter{},
		[]*interf.WithoutImplementation{}, []*interf.Error{}, []*interf.Simple{}, []*interf.OneMethod{},
		[]*interf.TwoMethods{}, []*interf.Public{}, []*interf.Seeder[any]{}, []*interf.Simple{},
		[]*interf.OneMethod{}, []*interf.TwoMethods{},
	}
}

func ofStructs() []interface{} {
	return []interface{}{
		[]structs.Boolean{}, []structs.Integer{}, []structs.String{}, []structs.Float{},
		[]structs.Complex{}, []structs.Map[any, any]{}, []structs.Slice[any]{},
		[]structs.Simple{}, []structs.Public{}, []structs.Channel[any]{},
		[]structs.CloserSuccess{}, []structs.CloserError{}, []structs.CloserReaderSuccess{},
		[]structs.CloserReaderError{}, []structs.CloserReaderWriterSuccess{},
		[]structs.CloserReaderWriterError{}, []structs.NotComparable{},
		[]structs.NaturallyComparable{}, []structs.NaturallyComparableWithMethods{},
		[]structs.Stringer{}, []structs.StringerBytes{}, []structs.StringerString{},
		[]structs.OneData{}, []structs.TwoData{},
		[]static.Profile{}, []static.Credential{}, []static.Product{}, []static.Address{},
		[]static.User{}, []static.Account{},
	}
}

func ofPointersToStructs() []interface{} {
	return []any{
		[]*structs.Boolean{}, []*structs.Integer{}, []*structs.String{}, []*structs.Float{},
		[]*structs.Complex{}, []*structs.Map[any, any]{}, []*structs.Slice[any]{},
		[]*structs.Simple{}, []*structs.Public{}, []*structs.Channel[any]{},
		[]*structs.CloserSuccess{}, []*structs.CloserError{}, []*structs.CloserReaderSuccess{},
		[]*structs.CloserReaderError{}, []*structs.CloserReaderWriterSuccess{},
		[]*structs.CloserReaderWriterError{}, []*structs.NotComparable{},
		[]*structs.NaturallyComparable{}, []*structs.NaturallyComparableWithMethods{},
		[]*structs.Stringer{}, []*structs.StringerBytes{}, []*structs.StringerString{},
		[]*structs.OneData{}, []*structs.TwoData{},
		[]*static.Profile{}, []*static.Credential{}, []*static.Product{}, []*static.Address{},
		[]*static.User{}, []*static.Account{},
	}
}

func ofChannels() []interface{} {
	return []interface{}{
		[]chan structs.Public{}, []chan structs.String{}, []chan structs.Simple{}, []chan structs.CloserSuccess{},
		[]chan structs.CloserError{}, []chan structs.CloserReaderSuccess{}, []chan structs.CloserReaderError{},
		[]chan structs.CloserReaderWriterSuccess{}, []chan structs.CloserReaderWriterError{},
		[]chan static.Product{}, []chan static.Credential{},
		[]chan bool{}, []chan int{}, []chan int8{}, []chan int16{}, []chan int32{}, []chan int64{},
		[]chan uint{}, []chan uint8{}, []chan uint16{}, []chan uint32{}, []chan uint64{}, []chan uintptr{},
		[]chan float32{}, []chan float64{}, []chan complex64{}, []chan complex128{}, []chan string{},
	}
}

func ofPointersOfChannels() []interface{} {
	return []interface{}{
		[]*chan structs.Public{}, []*chan structs.String{}, []*chan structs.Simple{}, []*chan structs.CloserSuccess{},
		[]*chan structs.CloserError{}, []*chan structs.CloserReaderSuccess{}, []*chan structs.CloserReaderError{},
		[]*chan structs.CloserReaderWriterSuccess{}, []*chan structs.CloserReaderWriterError{},
		[]*chan static.Product{}, []*chan static.Credential{},
		[]*chan bool{}, []*chan int{}, []*chan int8{}, []*chan int16{}, []*chan int32{}, []*chan int64{},
		[]*chan uint{}, []*chan uint8{}, []*chan uint16{}, []*chan uint32{}, []*chan uint64{}, []*chan uintptr{},
		[]*chan float32{}, []*chan float64{}, []*chan complex64{}, []*chan complex128{}, []*chan string{},
	}
}

func edgeValues() []interface{} {
	return []interface{}{
		[]**interface{}{}, []*chan interface{}{}, []*unsafe.Pointer{},
		[][][]*chan **unsafe.Pointer{},
	}
}

func ofValues() []interface{} {
	return slices.Concat(ofPointersOfChannels(), ofChannels(), ofStructs(),
		ofPointersToStructs(), ofPointersToInterfaces(), ofInterfaces(),
		pointerToPrimitiveNil(), primitivesNil(), primitivesEmpty(),
		edgeValues(),
	)
}
