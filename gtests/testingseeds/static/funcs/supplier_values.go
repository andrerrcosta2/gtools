// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package funcs

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
	"unsafe"
)

func supplierOfPrimitives() []any {
	return []any{
		func() bool { return false }, func() int { return 0 }, func() int8 { return int8(0) },
		func() int16 { return int16(0) }, func() int32 { return int32(0) }, func() int64 { return int64(0) },
		func() uint { return uint(0) }, func() uint8 { return uint8(0) }, func() uint16 { return uint16(0) },
		func() uint32 { return uint32(0) }, func() uint64 { return uint64(0) }, func() uintptr { return uintptr(0) },
		func() float32 { return float32(0) }, func() float64 { return float64(0) }, func() complex64 { return complex64(0) },
		func() complex128 { return complex128(0) }, func() string { return "" },
	}
}

func supplierOfInterfaces() []any {
	return []any{
		func() interf.Stringer { return nil }, func() interf.Seeder[any] { return nil },
		func() interf.Error { return nil }, func() interf.OneMethod { return nil },
		func() interf.TwoMethods { return nil }, func() interf.Public { return nil },
		func() interf.Simple { return nil }, func() interf.WithoutImplementation { return nil },
		func() interf.Closer { return nil }, func() interf.CloserReader { return nil },
		func() interf.CloserReaderWriter { return nil },
	}
}

func supplierStructs() []any {
	return []any{
		func() structs.Boolean { return structs.Boolean(false) }, func() structs.Integer { return structs.Integer(0) },
		func() structs.Complex { return structs.Complex(0) }, func() structs.String { return structs.String("") },
		func() structs.Slice[any] { return structs.Slice[any](nil) }, func() structs.Uint { return structs.Uint(0) },
		func() structs.Map[any, any] { return structs.Map[any, any](nil) },
		func() static.Account { return static.Account{} }, func() static.Address { return static.Address{} },
		func() static.Credential { return static.Credential{} }, func() static.Profile { return static.Profile{} },
		func() structs.NotComparable { return structs.NotComparable{} },
		func() structs.NotComparableWithMethods { return structs.NotComparableWithMethods{} },
		func() structs.NaturallyComparable { return structs.NaturallyComparable{} },
		func() structs.NaturallyComparableWithMethods { return structs.NaturallyComparableWithMethods{} },
		func() structs.CloserSuccess { return structs.CloserSuccess{} }, func() structs.CloserError { return structs.CloserError{} },
		func() structs.CloserReaderSuccess { return structs.CloserReaderSuccess{} }, func() structs.CloserReaderError { return structs.CloserReaderError{} },
		func() structs.CloserReaderWriterSuccess { return structs.CloserReaderWriterSuccess{} }, func() structs.CloserReaderWriterError { return structs.CloserReaderWriterError{} },
		func() structs.Stringer { return structs.Stringer{} }, func() structs.StringerBytes { return structs.StringerBytes{} },
		func() structs.StringerString { return structs.StringerString{} },
	}
}

func supplierChanOfPrimitives() []any {
	return []any{
		func() chan int { return nil }, func() chan int8 { return nil }, func() chan int16 { return nil },
		func() chan int32 { return nil }, func() chan int64 { return nil }, func() chan uint { return nil },
		func() chan uint8 { return nil }, func() chan uint16 { return nil }, func() chan uint32 { return nil },
		func() chan uint64 { return nil }, func() chan uintptr { return nil }, func() chan float32 { return nil },
		func() chan float64 { return nil }, func() chan complex64 { return nil }, func() chan complex128 { return nil },
		func() chan string { return nil }, func() chan bool { return nil },
	}
}

func supplierMapPrimitives() []any {
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

func supplierSliceOfPrimitives() []any {
	return []any{
		func() []int { return []int{} }, func() []int8 { return []int8{} }, func() []int16 { return []int16{} },
		func() []int32 { return []int32{} }, func() []int64 { return []int64{} }, func() []uint { return []uint{} },
		func() []uint8 { return []uint8{} }, func() []uint16 { return []uint16{} }, func() []uint32 { return []uint32{} },
		func() []uint64 { return []uint64{} }, func() []uintptr { return []uintptr{} }, func() []float32 { return []float32{} },
		func() []float64 { return []float64{} }, func() []complex64 { return []complex64{} }, func() []complex128 { return []complex128{} },
		func() []string { return []string{} }, func() []bool { return []bool{} },
	}
}

func supplierEdgeValues() []any {
	return []any{
		func() unsafe.Pointer { return nil },
		func() *unsafe.Pointer { return nil },
		func() **unsafe.Pointer { return nil },
	}
}

func suppliers() []any {
	return slices.Concat(
		supplierOfPrimitives(), supplierStructs(), supplierChanOfPrimitives(),
		supplierMapPrimitives(), supplierSliceOfPrimitives(), supplierEdgeValues(), supplierOfInterfaces(),
	)
}
