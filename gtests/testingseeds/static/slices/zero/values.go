// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
	"unsafe"
)

func valuesSet() []interface{} {
	return slices.Concat(pointerToChannelValues(), channelValues(), structValues(),
		pointerToStructValues(), pointerToInterfaceValues(), interfaceValues(),
		pointerToPrimitiveValues(), primitiveValues(),
		edgeValues(),
	)
}

func primitiveValues() []interface{} {
	return []interface{}{
		[]bool{}, []int{}, []int8{}, []int16{}, []int32{}, []int64{}, []uint{}, []uint8{},
		[]uint16{}, []uint32{}, []uint64{}, []uintptr{}, []float32{}, []float64{}, []complex64{},
		[]complex128{}, []string{},
	}
}

func pointerToPrimitiveValues() []interface{} {
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
	return []interface{}{bv, iv, i8v, i16v, i32v, i64v, uiv, u8v, u16v, u32v, u64v, uptrv, f32v, f64v, c64v, c128v, sv}
}

func interfaceValues() []interface{} {
	return []interface{}{
		[]interf.Stringer{}, []interf.Closer{}, []interf.CloserReader{}, []interf.CloserReaderWriter{},
		[]interf.WithoutImplementation{}, []interf.Error{}, []interf.Simple{}, []interf.OneMethod{},
		[]interf.TwoMethods{}, []interf.Public{}, []interf.Seeder[any]{}, []interf.Simple{},
		[]interf.OneMethod{}, []interf.TwoMethods{},
	}
}

func pointerToInterfaceValues() []interface{} {
	return []interface{}{
		[]*interf.Stringer{}, []*interf.Closer{}, []*interf.CloserReader{}, []*interf.CloserReaderWriter{},
		[]*interf.WithoutImplementation{}, []*interf.Error{}, []*interf.Simple{}, []*interf.OneMethod{},
		[]*interf.TwoMethods{}, []*interf.Public{}, []*interf.Seeder[any]{}, []*interf.Simple{},
		[]*interf.OneMethod{}, []*interf.TwoMethods{},
	}
}

func structValues() []interface{} {
	return []interface{}{
		[]models.Boolean{}, []models.Integer{}, []models.String{}, []models.Float{},
		[]models.Complex{}, []models.Map[any, any]{}, []models.Slice[any]{},
		[]models.Simple{}, []models.Public{}, []models.Channel[any]{},
		[]models.CloserSuccess{}, []models.CloserError{}, []models.CloserReaderSuccess{},
		[]models.CloserReaderError{}, []models.CloserReaderWriterSuccess{},
		[]models.CloserReaderWriterError{}, []models.NotComparable{},
		[]models.NaturallyComparable{}, []models.NaturallyComparableWithMethods{},
		[]models.Stringer{}, []models.StringerBytes{}, []models.StringerString{},
		[]models.OneData{}, []models.TwoData{},
		[]models.Profile{}, []models.Credential{}, []models.Product{}, []models.Address{},
		[]models.User{}, []models.Account{},
	}
}

func pointerToStructValues() []interface{} {
	return []any{
		[]*models.Boolean{}, []*models.Integer{}, []*models.String{}, []*models.Float{},
		[]*models.Complex{}, []*models.Map[any, any]{}, []*models.Slice[any]{},
		[]*models.Simple{}, []*models.Public{}, []*models.Channel[any]{},
		[]*models.CloserSuccess{}, []*models.CloserError{}, []*models.CloserReaderSuccess{},
		[]*models.CloserReaderError{}, []*models.CloserReaderWriterSuccess{},
		[]*models.CloserReaderWriterError{}, []*models.NotComparable{},
		[]*models.NaturallyComparable{}, []*models.NaturallyComparableWithMethods{},
		[]*models.Stringer{}, []*models.StringerBytes{}, []*models.StringerString{},
		[]*models.OneData{}, []*models.TwoData{},
		[]*models.Profile{}, []*models.Credential{}, []*models.Product{}, []*models.Address{},
		[]*models.User{}, []*models.Account{},
	}
}

func channelValues() []interface{} {
	return []interface{}{
		[]chan models.Public{}, []chan models.String{}, []chan models.Simple{}, []chan models.CloserSuccess{},
		[]chan models.CloserError{}, []chan models.CloserReaderSuccess{}, []chan models.CloserReaderError{},
		[]chan models.CloserReaderWriterSuccess{}, []chan models.CloserReaderWriterError{},
		[]chan models.Product{}, []chan models.Credential{},
		[]chan bool{}, []chan int{}, []chan int8{}, []chan int16{}, []chan int32{}, []chan int64{},
		[]chan uint{}, []chan uint8{}, []chan uint16{}, []chan uint32{}, []chan uint64{}, []chan uintptr{},
		[]chan float32{}, []chan float64{}, []chan complex64{}, []chan complex128{}, []chan string{},
	}
}

func pointerToChannelValues() []interface{} {
	return []interface{}{
		[]*chan models.Public{}, []*chan models.String{}, []*chan models.Simple{}, []*chan models.CloserSuccess{},
		[]*chan models.CloserError{}, []*chan models.CloserReaderSuccess{}, []*chan models.CloserReaderError{},
		[]*chan models.CloserReaderWriterSuccess{}, []*chan models.CloserReaderWriterError{},
		[]*chan models.Product{}, []*chan models.Credential{},
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
