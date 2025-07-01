// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
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
	return []any{
		random.SingleOf[*[]bool](), random.SingleOf[*[]int](), random.SingleOf[*[]int8](),
		random.SingleOf[*[]int16](), random.SingleOf[*[]int32](), random.SingleOf[*[]int64](),
		random.SingleOf[*[]uint](), random.SingleOf[*[]uint8](), random.SingleOf[*[]uint16](),
		random.SingleOf[*[]uint32](), random.SingleOf[*[]uint64](), random.SingleOf[*[]uintptr](),
		random.SingleOf[*[]float32](), random.SingleOf[*[]float64](), random.SingleOf[*[]complex64](),
		random.SingleOf[*[]complex128](), random.SingleOf[*[]string](),
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
	var s1 []*bool
	var s2 []*int
	var s3 []*int8
	var s4 []*int16
	var s5 []*int32
	var s6 []*int64
	var s7 []*uint
	var s8 []*uint8
	var s9 []*uint16
	var s10 []*uint32
	var s11 []*uint64
	var s12 []*uintptr
	var s13 []*float32
	var s14 []*float64
	var s15 []*complex64
	var s16 []*complex128
	var s17 []*string
	return []interface{}{
		&s1, &s2, &s3, &s4, &s5, &s6, &s7, &s8, &s9, &s10, &s11,
		&s12, &s13, &s14, &s15, &s16, &s17,
	}
}

func interfaceRefs() []interface{} {
	return []interface{}{
		random.SingleOf[*[]interf.Stringer](), random.SingleOf[*[]interf.Closer](),
		random.SingleOf[*[]interf.CloserReader](), random.SingleOf[*[]interf.CloserReaderWriter](),
		random.SingleOf[*[]interf.WithoutImplementation](), random.SingleOf[*[]interf.Error](),
		random.SingleOf[*[]interf.Simple](), random.SingleOf[*[]interf.OneMethod](),
		random.SingleOf[*[]interf.TwoMethods](), random.SingleOf[*[]interf.Public](),
		random.SingleOf[*[]interf.Seeder[any]](), random.SingleOf[*[]interf.Simple](),
		random.SingleOf[*[]interf.OneMethod](), random.SingleOf[*[]interf.TwoMethods](),
	}
}

func pointerToInterfaceRefs() []interface{} {
	return []interface{}{
		random.SingleOf[*[]*interf.Stringer](), random.SingleOf[*[]*interf.Closer](),
		random.SingleOf[*[]*interf.CloserReader](), random.SingleOf[*[]*interf.CloserReaderWriter](),
		random.SingleOf[*[]*interf.WithoutImplementation](), random.SingleOf[*[]*interf.Error](),
		random.SingleOf[*[]*interf.Simple](), random.SingleOf[*[]*interf.OneMethod](),
		random.SingleOf[*[]*interf.TwoMethods](), random.SingleOf[*[]*interf.Public](),
		random.SingleOf[*[]*interf.Seeder[any]](), random.SingleOf[*[]*interf.Simple](),
		random.SingleOf[*[]*interf.OneMethod](), random.SingleOf[*[]*interf.TwoMethods](),
	}
}

func structRefs() []interface{} {
	return []interface{}{
		random.SingleOf[*[]models.Boolean](), random.SingleOf[*[]models.Integer](),
		random.SingleOf[*[]models.String](), random.SingleOf[*[]models.Float](),
		random.SingleOf[*[]models.Complex](), random.SingleOf[*[]models.Map[string, int]](),
		random.SingleOf[*[]models.Slice[any]](), random.SingleOf[*[]models.Simple](),
		random.SingleOf[*[]models.Public](), random.SingleOf[*[]models.Channel[any]](),
		random.SingleOf[*[]models.CloserSuccess](), random.SingleOf[*[]models.CloserError](),
		random.SingleOf[*[]models.CloserReaderSuccess](), random.SingleOf[*[]models.CloserReaderError](),
		random.SingleOf[*[]models.CloserReaderWriterSuccess](), random.SingleOf[*[]models.CloserReaderWriterError](),
		random.SingleOf[*[]models.NotComparable](), random.SingleOf[*[]models.NaturallyComparable](),
		random.SingleOf[*[]models.NaturallyComparableWithMethods](), random.SingleOf[*[]models.Stringer](),
		random.SingleOf[*[]models.StringerBytes](), random.SingleOf[*[]models.StringerString](),
		random.SingleOf[*[]models.OneData](), random.SingleOf[*[]models.TwoData](),
		random.SingleOf[*[]models.Profile](), random.SingleOf[*[]models.Credential](),
		random.SingleOf[*[]models.Product](), random.SingleOf[*[]models.Address](),
		random.SingleOf[*[]models.User](), random.SingleOf[*[]models.Account](),
	}
}

func pointerToStructRefs() []interface{} {
	return []interface{}{
		random.SingleOf[*[]*models.Boolean](), random.SingleOf[*[]*models.Integer](),
		random.SingleOf[*[]*models.String](), random.SingleOf[*[]*models.Float](),
		random.SingleOf[*[]*models.Complex](), random.SingleOf[*[]*models.Map[string, int]](),
		random.SingleOf[*[]*models.Slice[any]](), random.SingleOf[*[]*models.Simple](),
		random.SingleOf[*[]*models.Public](), random.SingleOf[*[]*models.Channel[any]](),
		random.SingleOf[*[]*models.CloserSuccess](), random.SingleOf[*[]*models.CloserError](),
		random.SingleOf[*[]*models.CloserReaderSuccess](), random.SingleOf[*[]*models.CloserReaderError](),
		random.SingleOf[*[]*models.CloserReaderWriterSuccess](), random.SingleOf[*[]*models.CloserReaderWriterError](),
		random.SingleOf[*[]*models.NotComparable](), random.SingleOf[*[]*models.NaturallyComparable](),
		random.SingleOf[*[]*models.NaturallyComparableWithMethods](), random.SingleOf[*[]*models.Stringer](),
		random.SingleOf[*[]*models.StringerBytes](), random.SingleOf[*[]*models.StringerString](),
		random.SingleOf[*[]*models.OneData](), random.SingleOf[*[]*models.TwoData](),
		random.SingleOf[*[]*models.Profile](), random.SingleOf[*[]*models.Credential](),
		random.SingleOf[*[]*models.Product](), random.SingleOf[*[]*models.Address](),
		random.SingleOf[*[]*models.User](), random.SingleOf[*[]*models.Account](),
	}
}

func channelRefs() []interface{} {
	return []interface{}{
		random.SingleOf[*[]chan models.Public](), random.SingleOf[*[]chan models.String](),
		random.SingleOf[*[]chan models.Simple](), random.SingleOf[*[]chan models.CloserSuccess](),
		random.SingleOf[*[]chan models.CloserError](), random.SingleOf[*[]chan models.CloserReaderSuccess](),
		random.SingleOf[*[]chan models.CloserReaderError](), random.SingleOf[*[]chan models.CloserReaderWriterSuccess](),
		random.SingleOf[*[]chan models.CloserReaderWriterError](), random.SingleOf[*[]chan models.Product](),
		random.SingleOf[*[]chan models.Credential](), random.SingleOf[*[]chan bool](),
		random.SingleOf[*[]chan int](), random.SingleOf[*[]chan int8](), random.SingleOf[*[]chan int16](),
		random.SingleOf[*[]chan int32](), random.SingleOf[*[]chan int64](), random.SingleOf[*[]chan uint](),
		random.SingleOf[*[]chan uint8](), random.SingleOf[*[]chan uint16](), random.SingleOf[*[]chan uint32](),
		random.SingleOf[*[]chan uint64](), random.SingleOf[*[]chan uintptr](), random.SingleOf[*[]chan float32](),
		random.SingleOf[*[]chan float64](), random.SingleOf[*[]chan complex64](), random.SingleOf[*[]chan complex128](), random.SingleOf[*[]chan string](),
	}
}

func pointerToChannelRefs() []interface{} {
	return []interface{}{
		random.SingleOf[*[]*chan models.Public](), random.SingleOf[*[]*chan models.String](),
		random.SingleOf[*[]*chan models.Simple](), random.SingleOf[*[]*chan models.CloserSuccess](),
		random.SingleOf[*[]*chan models.CloserError](), random.SingleOf[*[]*chan models.CloserReaderSuccess](),
		random.SingleOf[*[]*chan models.CloserReaderError](), random.SingleOf[*[]*chan models.CloserReaderWriterSuccess](),
		random.SingleOf[*[]*chan models.CloserReaderWriterError](), random.SingleOf[*[]*chan models.Product](),
		random.SingleOf[*[]*chan models.Credential](), random.SingleOf[*[]*chan bool](),
		random.SingleOf[*[]*chan int](), random.SingleOf[*[]*chan int8](), random.SingleOf[*[]*chan int16](),
		random.SingleOf[*[]*chan int32](), random.SingleOf[*[]*chan int64](), random.SingleOf[*[]*chan uint](),
		random.SingleOf[*[]*chan uint8](), random.SingleOf[*[]*chan uint16](), random.SingleOf[*[]*chan uint32](),
		random.SingleOf[*[]*chan uint64](), random.SingleOf[*[]*chan uintptr](), random.SingleOf[*[]*chan float32](),
		random.SingleOf[*[]*chan float64](), random.SingleOf[*[]*chan complex64](), random.SingleOf[*[]*chan complex128](),
		random.SingleOf[*[]*chan string](),
	}
}

func edgeRefs() []interface{} {
	return []interface{}{
		random.SingleOf[**[]*interface{}](),
		random.SingleOf[unsafe.Pointer](), random.SingleOf[*unsafe.Pointer](),
		random.SingleOf[**unsafe.Pointer](), random.SingleOf[[][]*chan **unsafe.Pointer](),
	}
}
