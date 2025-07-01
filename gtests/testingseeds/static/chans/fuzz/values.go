// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
	"unsafe"
)

func valuesSet() []any {
	return slices.Concat(edgeValues(), interfaceValues(), pointerToInterfaceValues(), primitiveValues(), sliceOfPrimitiveValues(),
		sliceOfPointerToPrimitiveValues(), structValues(), sliceOfStructValues(),
		sliceOfPointerToInterfaceValues(), sliceOfInterfaceValues(), senderValues(), receiverValues(),
		pointerToSliceOfStructValues(),
	)
}

func primitiveValues() []any {
	return []interface{}{
		random.SingleOf[chan bool](), random.SingleOf[chan int](), random.SingleOf[chan int8](), random.SingleOf[chan int16](),
		random.SingleOf[chan int32](), random.SingleOf[chan int64](), random.SingleOf[chan uint](), random.SingleOf[chan uint8](),
		random.SingleOf[chan uint16](), random.SingleOf[chan uint32](), random.SingleOf[chan uint64](), random.SingleOf[chan uintptr](),
		random.SingleOf[chan float32](), random.SingleOf[chan float64](), random.SingleOf[chan complex64](),
		random.SingleOf[chan complex128](), random.SingleOf[chan string](),
	}
}

func pointerToPrimitiveValues() []any {
	cb := random.SingleOf[chan *bool]()
	ci := random.SingleOf[chan *int]()
	ci8 := random.SingleOf[chan *int8]()
	ci16 := random.SingleOf[chan *int16]()
	ci32 := random.SingleOf[chan *int32]()
	ci64 := random.SingleOf[chan *int64]()
	cu := random.SingleOf[chan *uint]()
	cu8 := random.SingleOf[chan *uint8]()
	cu16 := random.SingleOf[chan *uint16]()
	cu32 := random.SingleOf[chan *uint32]()
	cu64 := random.SingleOf[chan *uint64]()
	cuptr := random.SingleOf[chan *uintptr]()
	cf32 := random.SingleOf[chan *float32]()
	cf64 := random.SingleOf[chan *float64]()
	cc64 := random.SingleOf[chan *complex64]()
	cc128 := random.SingleOf[chan *complex128]()
	cs := random.SingleOf[chan *string]()
	return []any{
		cb, ci, ci8, ci16, ci32, ci64, cu, cu8, cu16, cu32, cu64, cuptr,
		cf32, cf64, cc64, cc128, cs,
	}
}

func sliceOfPrimitiveValues() []any {
	return []interface{}{
		random.SingleOf[chan []bool](),
		random.SingleOf[chan []int](),
		random.SingleOf[chan []int8](),
		random.SingleOf[chan []int16](),
		random.SingleOf[chan []int32](),
		random.SingleOf[chan []int64](),
		random.SingleOf[chan []uint](),
		random.SingleOf[chan []uint8](),
		random.SingleOf[chan []uint16](),
		random.SingleOf[chan []uint32](),
		random.SingleOf[chan []uint64](),
		random.SingleOf[chan []uintptr](),
		random.SingleOf[chan []float32](),
		random.SingleOf[chan []float64](),
		random.SingleOf[chan []complex64](),
		random.SingleOf[chan []complex128](),
		random.SingleOf[chan []string](),
	}
}

func sliceOfPointerToPrimitiveValues() []any {
	return []interface{}{
		random.SingleOf[chan []*bool](),
		random.SingleOf[chan []*int](),
		random.SingleOf[chan []*int8](),
		random.SingleOf[chan []*int16](),
		random.SingleOf[chan []*int32](),
		random.SingleOf[chan []*int64](),
		random.SingleOf[chan []*uint](),
		random.SingleOf[chan []*uint8](),
		random.SingleOf[chan []*uint16](),
		random.SingleOf[chan []*uint32](),
		random.SingleOf[chan []*uint64](),
		random.SingleOf[chan []*uintptr](),
		random.SingleOf[chan []*float32](),
		random.SingleOf[chan []*float64](),
		random.SingleOf[chan []*complex64](),
		random.SingleOf[chan []*complex128](),
		random.SingleOf[chan []*string](),
	}
}

func interfaceValues() []any {
	c1 := random.SingleOf[chan interface{}]()
	c2 := random.SingleOf[chan interf.Stringer]()
	c3 := random.SingleOf[chan interf.Closer]()
	c4 := random.SingleOf[chan interf.CloserReader]()
	c5 := random.SingleOf[chan interf.CloserReaderWriter]()
	c6 := random.SingleOf[chan interf.WithoutImplementation]()
	c7 := random.SingleOf[chan interf.Simple]()
	c8 := random.SingleOf[chan interf.Public]()
	c9 := random.SingleOf[chan interf.TwoMethods]()
	c10 := random.SingleOf[chan interf.OneMethod]()
	c11 := random.SingleOf[chan interf.Error]()
	c12 := random.SingleOf[chan interf.Seeder[any]]()
	return []interface{}{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12,
	}
}

func pointerToInterfaceValues() []any {
	c1 := random.SingleOf[chan *interface{}]()
	c2 := random.SingleOf[chan *interf.Stringer]()
	c3 := random.SingleOf[chan *interf.Closer]()
	c4 := random.SingleOf[chan *interf.CloserReader]()
	c5 := random.SingleOf[chan *interf.CloserReaderWriter]()
	c6 := random.SingleOf[chan *interf.WithoutImplementation]()
	c7 := random.SingleOf[chan *interf.Simple]()
	c8 := random.SingleOf[chan *interf.Public]()
	c9 := random.SingleOf[chan *interf.TwoMethods]()
	c10 := random.SingleOf[chan *interf.OneMethod]()
	c11 := random.SingleOf[chan *interf.Error]()
	c12 := random.SingleOf[chan *interf.Seeder[any]]()
	return []interface{}{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12,
	}
}

func sliceOfInterfaceValues() []any {
	c1 := random.SingleOf[chan []interface{}]()
	c2 := random.SingleOf[chan []interf.Stringer]()
	c3 := random.SingleOf[chan []interf.Closer]()
	c4 := random.SingleOf[chan []interf.CloserReader]()
	c5 := random.SingleOf[chan []interf.CloserReaderWriter]()
	c6 := random.SingleOf[chan []interf.WithoutImplementation]()
	c7 := random.SingleOf[chan []interf.Simple]()
	c8 := random.SingleOf[chan []interf.Public]()
	c9 := random.SingleOf[chan []interf.TwoMethods]()
	c10 := random.SingleOf[chan []interf.OneMethod]()
	c11 := random.SingleOf[chan []interf.Error]()
	c12 := random.SingleOf[chan []interf.Seeder[any]]()
	return []interface{}{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12,
	}
}

func sliceOfPointerToInterfaceValues() []any {
	c1 := random.SingleOf[chan []*interface{}]()
	c2 := random.SingleOf[chan []*interf.Stringer]()
	c3 := random.SingleOf[chan []*interf.Closer]()
	c4 := random.SingleOf[chan []*interf.CloserReader]()
	c5 := random.SingleOf[chan []*interf.CloserReaderWriter]()
	c6 := random.SingleOf[chan []*interf.WithoutImplementation]()
	c7 := random.SingleOf[chan []*interf.Simple]()
	c8 := random.SingleOf[chan []*interf.Public]()
	c9 := random.SingleOf[chan []*interf.TwoMethods]()
	c10 := random.SingleOf[chan []*interf.OneMethod]()
	c11 := random.SingleOf[chan []*interf.Error]()
	c12 := random.SingleOf[chan []*interf.Seeder[any]]()
	return []interface{}{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12,
	}
}

func structValues() []any {
	c1 := random.SingleOf[chan models.Boolean]()
	c2 := random.SingleOf[chan models.Integer]()
	c3 := random.SingleOf[chan models.String]()
	c4 := random.SingleOf[chan models.Float]()
	c5 := random.SingleOf[chan models.Complex]()
	c6 := random.SingleOf[chan models.Uint]()
	c7 := random.SingleOf[chan models.Channel[any]]()
	c8 := random.SingleOf[chan models.Slice[any]]()
	c9 := random.SingleOf[chan models.Map[any, any]]()
	c10 := random.SingleOf[chan models.NaturallyComparable]()
	c11 := random.SingleOf[chan models.Stringer]()
	c12 := random.SingleOf[chan models.CloserSuccess]()
	c13 := random.SingleOf[chan models.CloserError]()
	c14 := random.SingleOf[chan models.CloserReaderSuccess]()
	c15 := random.SingleOf[chan models.CloserReaderError]()
	c16 := random.SingleOf[chan models.CloserReaderWriterSuccess]()
	c17 := random.SingleOf[chan models.CloserReaderWriterError]()
	c18 := random.SingleOf[chan models.NotComparable]()
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func sliceOfStructValues() []any {
	c1 := random.SingleOf[chan []models.Boolean]()
	c2 := random.SingleOf[chan []models.Integer]()
	c3 := random.SingleOf[chan []models.String]()
	c4 := random.SingleOf[chan []models.Float]()
	c5 := random.SingleOf[chan []models.Complex]()
	c6 := random.SingleOf[chan []models.Uint]()
	c7 := random.SingleOf[chan []models.Channel[any]]()
	c8 := random.SingleOf[chan []models.Slice[any]]()
	c9 := random.SingleOf[chan []models.Map[any, any]]()
	c10 := random.SingleOf[chan []models.NaturallyComparable]()
	c11 := random.SingleOf[chan []models.Stringer]()
	c12 := random.SingleOf[chan []models.CloserSuccess]()
	c13 := random.SingleOf[chan []models.CloserError]()
	c14 := random.SingleOf[chan []models.CloserReaderSuccess]()
	c15 := random.SingleOf[chan []models.CloserReaderError]()
	c16 := random.SingleOf[chan []models.CloserReaderWriterSuccess]()
	c17 := random.SingleOf[chan []models.CloserReaderWriterError]()
	c18 := random.SingleOf[chan []models.NotComparable]()
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func pointerToSliceOfStructValues() []any {
	c1 := random.SingleOf[chan *[]models.Boolean]()
	c2 := random.SingleOf[chan *[]models.Integer]()
	c3 := random.SingleOf[chan *[]models.String]()
	c4 := random.SingleOf[chan *[]models.Float]()
	c5 := random.SingleOf[chan *[]models.Complex]()
	c6 := random.SingleOf[chan *[]models.Uint]()
	c7 := random.SingleOf[chan *[]models.Channel[any]]()
	c8 := random.SingleOf[chan *[]models.Slice[any]]()
	c9 := random.SingleOf[chan *[]models.Map[any, any]]()
	c10 := random.SingleOf[chan *[]models.NaturallyComparable]()
	c11 := random.SingleOf[chan *[]models.Stringer]()
	c12 := random.SingleOf[chan *[]models.CloserSuccess]()
	c13 := random.SingleOf[chan *[]models.CloserError]()
	c14 := random.SingleOf[chan *[]models.CloserReaderSuccess]()
	c15 := random.SingleOf[chan *[]models.CloserReaderError]()
	c16 := random.SingleOf[chan *[]models.CloserReaderWriterSuccess]()
	c17 := random.SingleOf[chan *[]models.CloserReaderWriterError]()
	c18 := random.SingleOf[chan *[]models.NotComparable]()
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func senderValues() []any {
	return []any{
		random.SingleOf[chan<- bool](), random.SingleOf[chan<- int](), random.SingleOf[chan<- int8](), random.SingleOf[chan<- int16](),
		random.SingleOf[chan<- int32](), random.SingleOf[chan<- int64](), random.SingleOf[chan<- uint](), random.SingleOf[chan<- uint8](),
		random.SingleOf[chan<- uint16](), random.SingleOf[chan<- uint32](), random.SingleOf[chan<- uint64](), random.SingleOf[chan<- uintptr](),
		random.SingleOf[chan<- float32](), random.SingleOf[chan<- float64](), random.SingleOf[chan<- complex64](),
		random.SingleOf[chan<- complex128](), random.SingleOf[chan<- string](),
		random.SingleOf[chan<- chan interface{}](), random.SingleOf[chan<- *chan interface{}](),
		random.SingleOf[chan<- *interf.DataBool[any]](), random.SingleOf[chan<- *interf.Simple](), random.SingleOf[chan<- *interf.Public](),
		random.SingleOf[chan<- *interf.TwoMethods](), random.SingleOf[chan<- *interf.OneMethod](),
		random.SingleOf[chan<- *interf.Error](), random.SingleOf[chan<- *interf.Seeder[any]](),
		random.SingleOf[chan<- *models.Integer](), random.SingleOf[chan<- *models.String](), random.SingleOf[chan<- *models.Float](),
		random.SingleOf[chan<- *models.Complex](), random.SingleOf[chan<- *models.Uint](), random.SingleOf[chan<- *models.Channel[any]](),
		random.SingleOf[chan<- *models.Slice[any]](), random.SingleOf[chan<- *models.Map[any, any]](),
		random.SingleOf[chan<- *models.NaturallyComparable](), random.SingleOf[chan<- *models.Stringer](),
		random.SingleOf[chan<- *models.CloserSuccess](), random.SingleOf[chan<- *models.CloserError](),
		random.SingleOf[chan<- *models.CloserReaderSuccess](), random.SingleOf[chan<- *models.CloserReaderError](),
		random.SingleOf[chan<- *models.CloserReaderWriterSuccess](), random.SingleOf[chan<- *models.CloserReaderWriterError](),
		random.SingleOf[chan<- *models.NotComparable](), random.SingleOf[chan<- *models.User](), random.SingleOf[chan<- *models.Account](),
		random.SingleOf[chan<- *models.Address](),
	}
}

func receiverValues() []any {
	return []any{
		random.SingleOf[<-chan bool](), random.SingleOf[<-chan int](), random.SingleOf[<-chan int8](), random.SingleOf[<-chan int16](),
		random.SingleOf[<-chan int32](), random.SingleOf[<-chan int64](), random.SingleOf[<-chan uint](), random.SingleOf[<-chan uint8](),
		random.SingleOf[<-chan uint16](), random.SingleOf[<-chan uint32](), random.SingleOf[<-chan uint64](), random.SingleOf[<-chan uintptr](),
		random.SingleOf[<-chan float32](), random.SingleOf[<-chan float64](), random.SingleOf[<-chan complex64](),
		random.SingleOf[<-chan complex128](), random.SingleOf[<-chan string](),
		random.SingleOf[<-chan chan interface{}](), random.SingleOf[<-chan *chan interface{}](),
		random.SingleOf[<-chan *interf.DataBool[any]](), random.SingleOf[<-chan *interf.Simple](), random.SingleOf[<-chan *interf.Public](),
		random.SingleOf[<-chan *interf.TwoMethods](), random.SingleOf[<-chan *interf.OneMethod](),
		random.SingleOf[<-chan *interf.Error](), random.SingleOf[<-chan *interf.Seeder[any]](),
		random.SingleOf[<-chan *models.Integer](), random.SingleOf[<-chan *models.String](), random.SingleOf[<-chan *models.Float](),
		random.SingleOf[<-chan *models.Complex](), random.SingleOf[<-chan *models.Uint](), random.SingleOf[<-chan *models.Channel[any]](),
		random.SingleOf[<-chan *models.Slice[any]](), random.SingleOf[<-chan *models.Map[any, any]](),
		random.SingleOf[<-chan *models.NaturallyComparable](), random.SingleOf[<-chan *models.Stringer](),
		random.SingleOf[<-chan *models.CloserSuccess](), random.SingleOf[<-chan *models.CloserError](),
		random.SingleOf[<-chan *models.CloserReaderSuccess](), random.SingleOf[<-chan *models.CloserReaderError](),
		random.SingleOf[<-chan *models.CloserReaderWriterSuccess](), random.SingleOf[<-chan *models.CloserReaderWriterError](),
		random.SingleOf[<-chan *models.NotComparable](), random.SingleOf[<-chan *models.User](), random.SingleOf[<-chan *models.Account](),
		random.SingleOf[<-chan *models.Address](),
	}
}

func edgeValues() []any {
	return []interface{}{
		random.SingleOf[chan chan interface{}](), random.SingleOf[chan *chan interface{}](), random.SingleOf[chan *unsafe.Pointer](), random.SingleOf[chan **unsafe.Pointer](),
	}
}
