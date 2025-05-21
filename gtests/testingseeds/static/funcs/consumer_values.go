// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package funcs

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
	"unsafe"
)

func consumerOfPrimitives() []any {
	return []any{
		func(b bool) {}, func(i int) {}, func(i int8) {}, func(i int16) {}, func(i int32) {}, func(i int64) {},
		func(i uint) {}, func(i uint8) {}, func(i uint16) {}, func(i uint32) {}, func(i uint64) {},
		func(i uintptr) {}, func(f float32) {}, func(f float64) {}, func(c complex64) {}, func(c complex128) {},
		func(s string) {},
	}
}

func consumerOfVariadicPrimitives() []any {
	return []any{
		func(b ...bool) {}, func(i ...int) {}, func(i ...int8) {}, func(i ...int16) {}, func(i ...int32) {}, func(i ...int64) {},
		func(i ...uint) {}, func(i ...uint8) {}, func(i ...uint16) {}, func(i ...uint32) {}, func(i ...uint64) {},
		func(i ...uintptr) {}, func(f ...float32) {}, func(f ...float64) {}, func(c ...complex64) {}, func(c ...complex128) {},
		func(s ...string) {},
	}
}

func consumerOfInterfaces() []any {
	f1 := func(i interface{}) {}
	f2 := func(i interf.Stringer) {}
	f3 := func(i interf.Seeder[any]) {}
	f4 := func(i interf.Error) {}
	f5 := func(i interf.OneMethod) {}
	f6 := func(i interf.TwoMethods) {}
	f7 := func(i interf.Public) {}
	f8 := func(i interf.Simple) {}
	f9 := func(i interf.WithoutImplementation) {}
	f10 := func(i interf.Closer) {}
	f11 := func(i interf.CloserReader) {}
	f12 := func(i interf.CloserReaderWriter) {}

	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12,
	}
}

func consumerOfVariadicInterfaces() []any {
	f1 := func(i ...interface{}) {}
	f2 := func(i ...interf.Stringer) {}
	f3 := func(i ...interf.Seeder[any]) {}
	f4 := func(i ...interf.Error) {}
	f5 := func(i ...interf.OneMethod) {}
	f6 := func(i ...interf.TwoMethods) {}
	f7 := func(i ...interf.Public) {}
	f8 := func(i ...interf.Simple) {}
	f9 := func(i ...interf.WithoutImplementation) {}
	f10 := func(i ...interf.Closer) {}
	f11 := func(i ...interf.CloserReader) {}
	f12 := func(i ...interf.CloserReaderWriter) {}

	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12,
	}
}

func consumerStructs() []any {
	f1 := func(i structs.Boolean) {}
	f2 := func(i structs.Integer) {}
	f3 := func(i structs.Complex) {}
	f4 := func(i structs.String) {}
	f5 := func(i structs.Slice[any]) {}
	f6 := func(i structs.Uint) {}
	f7 := func(i structs.Map[any, any]) {}
	f8 := func(i static.Account) {}
	f9 := func(i static.Address) {}
	f10 := func(i static.Credential) {}
	f11 := func(i static.Profile) {}
	f12 := func(i structs.NotComparable) {}
	f13 := func(i structs.NotComparableWithMethods) {}
	f14 := func(i structs.NaturallyComparable) {}

	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8,
		f9, f10, f11, f12, f13, f14,
	}
}

func consumerVariadicStruct() []any {
	f1 := func(i ...structs.Boolean) {}
	f2 := func(i ...structs.Integer) {}
	f3 := func(i ...structs.Complex) {}
	f4 := func(i ...structs.String) {}
	f5 := func(i ...structs.Slice[any]) {}
	f6 := func(i ...structs.Uint) {}
	f7 := func(i ...structs.Map[any, any]) {}
	f8 := func(i ...static.Account) {}
	f9 := func(i ...static.Address) {}
	f10 := func(i ...static.Credential) {}
	f11 := func(i ...static.Profile) {}
	f12 := func(i ...structs.NotComparable) {}
	f13 := func(i ...structs.NotComparableWithMethods) {}
	f14 := func(i ...structs.NaturallyComparable) {}

	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8,
		f9, f10, f11, f12, f13, f14,
	}
}

func consumerChanOfPrimitives() []any {
	f1 := func(i chan int) {}
	f2 := func(i chan int8) {}
	f3 := func(i chan int16) {}
	f4 := func(i chan int32) {}
	f5 := func(i chan int64) {}
	f6 := func(i chan uint) {}
	f7 := func(i chan uint8) {}
	f8 := func(i chan uint16) {}
	f9 := func(i chan uint32) {}
	f10 := func(i chan uint64) {}
	f11 := func(i chan uintptr) {}
	f12 := func(i chan float32) {}
	f13 := func(i chan float64) {}
	f14 := func(i chan complex64) {}
	f15 := func(i chan complex128) {}
	f16 := func(i chan string) {}
	f17 := func(i chan bool) {}
	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17,
	}
}

func consumerVariadicChanOfPrimitives() []any {
	f1 := func(i ...chan int) {}
	f2 := func(i ...chan int8) {}
	f3 := func(i ...chan int16) {}
	f4 := func(i ...chan int32) {}
	f5 := func(i ...chan int64) {}
	f6 := func(i ...chan uint) {}
	f7 := func(i ...chan uint8) {}
	f8 := func(i ...chan uint16) {}
	f9 := func(i ...chan uint32) {}
	f10 := func(i ...chan uint64) {}
	f11 := func(i ...chan uintptr) {}
	f12 := func(i ...chan float32) {}
	f13 := func(i ...chan float64) {}
	f14 := func(i ...chan complex64) {}
	f15 := func(i ...chan complex128) {}
	f16 := func(i ...chan string) {}
	f17 := func(i ...chan bool) {}
	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17,
	}
}

func consumerMapPrimitives() []any {
	f1 := func(i map[string]int) {}
	f2 := func(i map[string]int8) {}
	f3 := func(i map[string]int16) {}
	f4 := func(i map[string]int32) {}
	f5 := func(i map[string]int64) {}
	f6 := func(i map[string]uint) {}
	f7 := func(i map[string]uint8) {}
	f8 := func(i map[string]uint16) {}
	f9 := func(i map[string]uint32) {}
	f10 := func(i map[string]uint64) {}
	f11 := func(i map[string]uintptr) {}
	f12 := func(i map[string]float32) {}
	f13 := func(i map[string]float64) {}
	f14 := func(i map[string]complex64) {}
	f15 := func(i map[string]complex128) {}
	f16 := func(i map[string]string) {}
	f17 := func(i map[string]bool) {}
	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17,
	}
}

func consumerVariadicMapOsPrimitives() []any {
	f1 := func(i ...map[string]int) {}
	f2 := func(i ...map[string]int8) {}
	f3 := func(i ...map[string]int16) {}
	f4 := func(i ...map[string]int32) {}
	f5 := func(i ...map[string]int64) {}
	f6 := func(i ...map[string]uint) {}
	f7 := func(i ...map[string]uint8) {}
	f8 := func(i ...map[string]uint16) {}
	f9 := func(i ...map[string]uint32) {}
	f10 := func(i ...map[string]uint64) {}
	f11 := func(i ...map[string]uintptr) {}
	f12 := func(i ...map[string]float32) {}
	f13 := func(i ...map[string]float64) {}
	f14 := func(i ...map[string]complex64) {}
	f15 := func(i ...map[string]complex128) {}
	f16 := func(i ...map[string]string) {}
	f17 := func(i ...map[string]bool) {}
	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17,
	}
}

func consumerSliceOfPrimitives() []any {
	f1 := func(i []int) {}
	f2 := func(i []int8) {}
	f3 := func(i []int16) {}
	f4 := func(i []int32) {}
	f5 := func(i []int64) {}
	f6 := func(i []uint) {}
	f7 := func(i []uint8) {}
	f8 := func(i []uint16) {}
	f9 := func(i []uint32) {}
	f10 := func(i []uint64) {}
	f11 := func(i []uintptr) {}
	f12 := func(i []float32) {}
	f13 := func(i []float64) {}
	f14 := func(i []complex64) {}
	f15 := func(i []complex128) {}
	f16 := func(i []string) {}
	f17 := func(i []bool) {}
	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17,
	}
}

func consumerVariadicSliceOfPrimitives() []any {
	f1 := func(i ...[]int) {}
	f2 := func(i ...[]int8) {}
	f3 := func(i ...[]int16) {}
	f4 := func(i ...[]int32) {}
	f5 := func(i ...[]int64) {}
	f6 := func(i ...[]uint) {}
	f7 := func(i ...[]uint8) {}
	f8 := func(i ...[]uint16) {}
	f9 := func(i ...[]uint32) {}
	f10 := func(i ...[]uint64) {}
	f11 := func(i ...[]uintptr) {}
	f12 := func(i ...[]float32) {}
	f13 := func(i ...[]float64) {}
	f14 := func(i ...[]complex64) {}
	f15 := func(i ...[]complex128) {}
	f16 := func(i ...[]string) {}
	f17 := func(i ...[]bool) {}
	return []any{
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17,
	}
}

func consumerEdgeValues() []any {
	return []any{
		func(i unsafe.Pointer) {}, func(i *unsafe.Pointer) {}, func(i **unsafe.Pointer) {},
		func(i **uintptr) {}, func(i **struct{}) {},
	}
}

func consumers() []any {
	return slices.Concat(
		consumerOfPrimitives(), consumerOfVariadicPrimitives(), consumerStructs(),
		consumerVariadicStruct(), consumerChanOfPrimitives(), consumerVariadicChanOfPrimitives(),
		consumerMapPrimitives(), consumerVariadicMapOsPrimitives(), consumerSliceOfPrimitives(),
		consumerVariadicSliceOfPrimitives(), consumerEdgeValues(), consumerOfInterfaces(),
		consumerOfVariadicInterfaces(),
	)
}
