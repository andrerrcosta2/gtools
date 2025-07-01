// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package supplier

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
	"unsafe"
)

func referencesSet() []any {
	return slices.Concat(
		primitiveRefs(), interfaceRefs(), structRefs(),
		chanOfPrimitiveRefs(),
		sliceOfPrimitiveRefs(), mapOfPrimitiveRefs(), edgeRefs(),
	)
}

func primitiveRefs() []any {
	f1 := func() bool { return false }
	f2 := func() int { return 0 }
	f3 := func() int8 { return int8(0) }
	f4 := func() int16 { return int16(0) }
	f5 := func() int32 { return int32(0) }
	f6 := func() int64 { return int64(0) }
	f7 := func() uint { return uint(0) }
	f8 := func() uint8 { return uint8(0) }
	f9 := func() uint16 { return uint16(0) }
	f10 := func() uint32 { return uint32(0) }
	f11 := func() uint64 { return uint64(0) }
	f12 := func() uintptr { return uintptr(0) }
	f13 := func() float32 { return float32(0) }
	f14 := func() float64 { return float64(0) }
	f15 := func() complex64 { return complex64(0) }
	f16 := func() complex128 { return complex128(0) }
	f17 := func() string { return "" }
	return []any{
		&f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9,
		&f10, &f11, &f12, &f13, &f14, &f15, &f16, &f17,
	}
}

func interfaceRefs() []any {
	f1 := func() interf.Stringer { return nil }
	f2 := func() interf.Seeder[any] { return nil }
	f3 := func() interf.Error { return nil }
	f4 := func() interf.OneMethod { return nil }
	f5 := func() interf.TwoMethods { return nil }
	f6 := func() interf.Public { return nil }
	f7 := func() interf.Simple { return nil }
	f8 := func() interf.WithoutImplementation { return nil }
	f9 := func() interf.Closer { return nil }
	f10 := func() interf.CloserReader { return nil }
	f11 := func() interf.CloserReaderWriter { return nil }
	return []any{
		&f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9, &f10, &f11,
	}
}

func structRefs() []any {
	f1 := func() models.Boolean { return models.Boolean{} }
	f2 := func() models.Integer { return models.Integer{} }
	f3 := func() models.Complex { return models.Complex{} }
	f4 := func() models.String { return models.String{} }
	f5 := func() models.Slice[any] { return models.Slice[any]{} }
	f6 := func() models.Uint { return models.Uint{} }
	f7 := func() models.Map[string, any] { return models.Map[string, any]{} }
	f8 := func() models.Account { return models.Account{} }
	f9 := func() models.Address { return models.Address{} }
	f10 := func() models.Credential { return models.Credential{} }
	f11 := func() models.Profile { return models.Profile{} }
	f12 := func() models.NotComparable { return models.NotComparable{} }
	f13 := func() models.NotComparableWithMethods { return models.NotComparableWithMethods{} }
	f14 := func() models.NaturallyComparable { return models.NaturallyComparable{} }
	f15 := func() models.NaturallyComparableWithMethods { return models.NaturallyComparableWithMethods{} }
	f16 := func() models.CloserSuccess { return models.CloserSuccess{} }
	f17 := func() models.CloserError { return models.CloserError{} }
	f18 := func() models.CloserReaderSuccess { return models.CloserReaderSuccess{} }
	f19 := func() models.CloserReaderError { return models.CloserReaderError{} }
	f20 := func() models.CloserReaderWriterSuccess { return models.CloserReaderWriterSuccess{} }
	f21 := func() models.CloserReaderWriterError { return models.CloserReaderWriterError{} }
	return []any{
		&f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9, &f10, &f11, &f12,
		&f13, &f14, &f15, &f16, &f17, &f18, &f19, &f20, &f21,
	}
}

func chanOfPrimitiveRefs() []any {
	f1 := func() chan int { return nil }
	f2 := func() chan int8 { return nil }
	f3 := func() chan int16 { return nil }
	f4 := func() chan int32 { return nil }
	f5 := func() chan int64 { return nil }
	f6 := func() chan uint { return nil }
	f7 := func() chan uint8 { return nil }
	f8 := func() chan uint16 { return nil }
	f9 := func() chan uint32 { return nil }
	f10 := func() chan uint64 { return nil }
	f11 := func() chan uintptr { return nil }
	f12 := func() chan float32 { return nil }
	f13 := func() chan float64 { return nil }
	f14 := func() chan complex64 { return nil }
	f15 := func() chan complex128 { return nil }
	f16 := func() chan string { return nil }
	f17 := func() chan bool { return nil }
	return []any{
		&f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9,
		&f10, &f11, &f12, &f13, &f14, &f15, &f16, &f17,
	}
}

func mapOfPrimitiveRefs() []any {
	f1 := func() map[int]int { return nil }
	f2 := func() map[int8]int8 { return nil }
	f3 := func() map[int16]int16 { return nil }
	f4 := func() map[int32]int32 { return nil }
	f5 := func() map[int64]int64 { return nil }
	f6 := func() map[uint]uint { return nil }
	f7 := func() map[uint8]uint8 { return nil }
	f8 := func() map[uint16]uint16 { return nil }
	f9 := func() map[uint32]uint32 { return nil }
	f10 := func() map[uint64]uint64 { return nil }
	f11 := func() map[uintptr]uintptr { return nil }
	f12 := func() map[float32]float32 { return nil }
	f13 := func() map[float64]float64 { return nil }
	f14 := func() map[complex64]complex64 { return nil }
	f15 := func() map[complex128]complex128 { return nil }
	f16 := func() map[string]string { return nil }
	f17 := func() map[bool]bool { return nil }
	return []any{
		&f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9,
		&f10, &f11, &f12, &f13, &f14, &f15, &f16, &f17,
	}
}

func sliceOfPrimitiveRefs() []any {
	f1 := func() []int { return []int{} }
	f2 := func() []int8 { return []int8{} }
	f3 := func() []int16 { return []int16{} }
	f4 := func() []int32 { return []int32{} }
	f5 := func() []int64 { return []int64{} }
	f6 := func() []uint { return []uint{} }
	f7 := func() []uint8 { return []uint8{} }
	f8 := func() []uint16 { return []uint16{} }
	f9 := func() []uint32 { return []uint32{} }
	f10 := func() []uint64 { return []uint64{} }
	f11 := func() []uintptr { return []uintptr{} }
	f12 := func() []float32 { return []float32{} }
	f13 := func() []float64 { return []float64{} }
	f14 := func() []complex64 { return []complex64{} }
	f15 := func() []complex128 { return []complex128{} }
	f16 := func() []string { return []string{} }
	f17 := func() []bool { return []bool{} }
	return []any{
		&f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9, &f10,
		&f11, &f12, &f13, &f14, &f15, &f16, &f17,
	}
}

func edgeRefs() []any {
	f1 := func() unsafe.Pointer { return nil }
	f2 := func() *unsafe.Pointer { return nil }
	f3 := func() **unsafe.Pointer { return nil }
	return []any{
		&f1, &f2, &f3,
	}
}
