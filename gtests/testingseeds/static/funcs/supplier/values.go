// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package supplier

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
	"unsafe"
)

func valuesSet() []any {
	return slices.Concat(
		primitiveValues(), structValues(), chanOfPrimitiveValues(),
		mapOfPrimitiveValues(), sliceOfPrimitiveValues(), edgeValues(), interfaceValues(),
	)
}

func primitiveValues() []any {
	return []any{
		func() bool { return false }, func() int { return 0 }, func() int8 { return int8(0) },
		func() int16 { return int16(0) }, func() int32 { return int32(0) }, func() int64 { return int64(0) },
		func() uint { return uint(0) }, func() uint8 { return uint8(0) }, func() uint16 { return uint16(0) },
		func() uint32 { return uint32(0) }, func() uint64 { return uint64(0) }, func() uintptr { return uintptr(0) },
		func() float32 { return float32(0) }, func() float64 { return float64(0) }, func() complex64 { return complex64(0) },
		func() complex128 { return complex128(0) }, func() string { return "" },
	}
}

func interfaceValues() []any {
	return []any{
		func() interf.Stringer { return nil }, func() interf.Seeder[any] { return nil },
		func() interf.Error { return nil }, func() interf.OneMethod { return nil },
		func() interf.TwoMethods { return nil }, func() interf.Public { return nil },
		func() interf.Simple { return nil }, func() interf.WithoutImplementation { return nil },
		func() interf.Closer { return nil }, func() interf.CloserReader { return nil },
		func() interf.CloserReaderWriter { return nil },
	}
}

func structValues() []any {
	return []any{
		func() models.Boolean { return models.Boolean{} }, func() models.Integer { return models.Integer{} },
		func() models.Complex { return models.Complex{} }, func() models.String { return models.String{} },
		func() models.Slice[any] { return models.Slice[any]{} }, func() models.Uint { return models.Uint{} },
		func() models.Map[string, any] { return models.Map[string, any]{} },
		func() models.Account { return models.Account{} }, func() models.Address { return models.Address{} },
		func() models.Credential { return models.Credential{} }, func() models.Profile { return models.Profile{} },
		func() models.NotComparable { return models.NotComparable{} },
		func() models.NotComparableWithMethods { return models.NotComparableWithMethods{} },
		func() models.NaturallyComparable { return models.NaturallyComparable{} },
		func() models.NaturallyComparableWithMethods { return models.NaturallyComparableWithMethods{} },
		func() models.CloserSuccess { return models.CloserSuccess{} }, func() models.CloserError { return models.CloserError{} },
		func() models.CloserReaderSuccess { return models.CloserReaderSuccess{} }, func() models.CloserReaderError { return models.CloserReaderError{} },
		func() models.CloserReaderWriterSuccess { return models.CloserReaderWriterSuccess{} }, func() models.CloserReaderWriterError { return models.CloserReaderWriterError{} },
		func() models.Stringer { return models.Stringer{} }, func() models.StringerBytes { return models.StringerBytes{} },
		func() models.StringerString { return models.StringerString{} },
	}
}

func chanOfPrimitiveValues() []any {
	return []any{
		func() chan int { return nil }, func() chan int8 { return nil }, func() chan int16 { return nil },
		func() chan int32 { return nil }, func() chan int64 { return nil }, func() chan uint { return nil },
		func() chan uint8 { return nil }, func() chan uint16 { return nil }, func() chan uint32 { return nil },
		func() chan uint64 { return nil }, func() chan uintptr { return nil }, func() chan float32 { return nil },
		func() chan float64 { return nil }, func() chan complex64 { return nil }, func() chan complex128 { return nil },
		func() chan string { return nil }, func() chan bool { return nil },
	}
}

func mapOfPrimitiveValues() []any {
	return []any{
		func() map[string]int { return nil }, func() map[string]int8 { return nil },
		func() map[string]int16 { return nil }, func() map[string]int32 { return nil },
		func() map[string]int64 { return nil }, func() map[string]uint { return nil },
		func() map[string]uint8 { return nil }, func() map[string]uint16 { return nil },
		func() map[string]uint32 { return nil }, func() map[string]uint64 { return nil },
		func() map[string]uintptr { return nil }, func() map[string]float32 { return nil },
		func() map[string]float64 { return nil }, func() map[string]complex64 { return nil },
		func() map[string]complex128 { return nil }, func() map[string]string { return nil },
		func() map[string]bool { return nil },
	}
}

func sliceOfPrimitiveValues() []any {
	return []any{
		func() []int { return []int{} }, func() []int8 { return []int8{} }, func() []int16 { return []int16{} },
		func() []int32 { return []int32{} }, func() []int64 { return []int64{} }, func() []uint { return []uint{} },
		func() []uint8 { return []uint8{} }, func() []uint16 { return []uint16{} }, func() []uint32 { return []uint32{} },
		func() []uint64 { return []uint64{} }, func() []uintptr { return []uintptr{} }, func() []float32 { return []float32{} },
		func() []float64 { return []float64{} }, func() []complex64 { return []complex64{} }, func() []complex128 { return []complex128{} },
		func() []string { return []string{} }, func() []bool { return []bool{} },
	}
}

func edgeValues() []any {
	return []any{
		func() unsafe.Pointer { return nil },
		func() *unsafe.Pointer { return nil },
		func() **unsafe.Pointer { return nil },
	}
}
