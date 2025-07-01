// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
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
		make(chan bool), make(chan int), make(chan int8), make(chan int16),
		make(chan int32), make(chan int64), make(chan uint), make(chan uint8),
		make(chan uint16), make(chan uint32), make(chan uint64), make(chan uintptr),
		make(chan float32), make(chan float64), make(chan complex64),
		make(chan complex128), make(chan string),
	}
}

func pointerToPrimitiveValues() []any {
	cb := make(chan *bool)
	ci := make(chan *int)
	ci8 := make(chan *int8)
	ci16 := make(chan *int16)
	ci32 := make(chan *int32)
	ci64 := make(chan *int64)
	cu := make(chan *uint)
	cu8 := make(chan *uint8)
	cu16 := make(chan *uint16)
	cu32 := make(chan *uint32)
	cu64 := make(chan *uint64)
	cuptr := make(chan *uintptr)
	cf32 := make(chan *float32)
	cf64 := make(chan *float64)
	cc64 := make(chan *complex64)
	cc128 := make(chan *complex128)
	cs := make(chan *string)
	return []any{
		cb, ci, ci8, ci16, ci32, ci64, cu, cu8, cu16, cu32, cu64, cuptr,
		cf32, cf64, cc64, cc128, cs,
	}
}

func sliceOfPrimitiveValues() []any {
	return []interface{}{
		make(chan []bool),
		make(chan []int),
		make(chan []int8),
		make(chan []int16),
		make(chan []int32),
		make(chan []int64),
		make(chan []uint),
		make(chan []uint8),
		make(chan []uint16),
		make(chan []uint32),
		make(chan []uint64),
		make(chan []uintptr),
		make(chan []float32),
		make(chan []float64),
		make(chan []complex64),
		make(chan []complex128),
		make(chan []string),
	}
}

func sliceOfPointerToPrimitiveValues() []any {
	return []interface{}{
		make(chan []*bool),
		make(chan []*int),
		make(chan []*int8),
		make(chan []*int16),
		make(chan []*int32),
		make(chan []*int64),
		make(chan []*uint),
		make(chan []*uint8),
		make(chan []*uint16),
		make(chan []*uint32),
		make(chan []*uint64),
		make(chan []*uintptr),
		make(chan []*float32),
		make(chan []*float64),
		make(chan []*complex64),
		make(chan []*complex128),
		make(chan []*string),
	}
}

func interfaceValues() []any {
	c1 := make(chan interface{})
	c2 := make(chan interf.Stringer)
	c3 := make(chan interf.Closer)
	c4 := make(chan interf.CloserReader)
	c5 := make(chan interf.CloserReaderWriter)
	c6 := make(chan interf.WithoutImplementation)
	c7 := make(chan interf.Simple)
	c8 := make(chan interf.Public)
	c9 := make(chan interf.TwoMethods)
	c10 := make(chan interf.OneMethod)
	c11 := make(chan interf.Error)
	c12 := make(chan interf.Seeder[any])
	return []interface{}{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12,
	}
}

func pointerToInterfaceValues() []any {
	c1 := make(chan *interface{})
	c2 := make(chan *interf.Stringer)
	c3 := make(chan *interf.Closer)
	c4 := make(chan *interf.CloserReader)
	c5 := make(chan *interf.CloserReaderWriter)
	c6 := make(chan *interf.WithoutImplementation)
	c7 := make(chan *interf.Simple)
	c8 := make(chan *interf.Public)
	c9 := make(chan *interf.TwoMethods)
	c10 := make(chan *interf.OneMethod)
	c11 := make(chan *interf.Error)
	c12 := make(chan *interf.Seeder[any])
	return []interface{}{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12,
	}
}

func sliceOfInterfaceValues() []any {
	c1 := make(chan []interface{})
	c2 := make(chan []interf.Stringer)
	c3 := make(chan []interf.Closer)
	c4 := make(chan []interf.CloserReader)
	c5 := make(chan []interf.CloserReaderWriter)
	c6 := make(chan []interf.WithoutImplementation)
	c7 := make(chan []interf.Simple)
	c8 := make(chan []interf.Public)
	c9 := make(chan []interf.TwoMethods)
	c10 := make(chan []interf.OneMethod)
	c11 := make(chan []interf.Error)
	c12 := make(chan []interf.Seeder[any])
	return []interface{}{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12,
	}
}

func sliceOfPointerToInterfaceValues() []any {
	c1 := make(chan []*interface{})
	c2 := make(chan []*interf.Stringer)
	c3 := make(chan []*interf.Closer)
	c4 := make(chan []*interf.CloserReader)
	c5 := make(chan []*interf.CloserReaderWriter)
	c6 := make(chan []*interf.WithoutImplementation)
	c7 := make(chan []*interf.Simple)
	c8 := make(chan []*interf.Public)
	c9 := make(chan []*interf.TwoMethods)
	c10 := make(chan []*interf.OneMethod)
	c11 := make(chan []*interf.Error)
	c12 := make(chan []*interf.Seeder[any])
	return []interface{}{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12,
	}
}

func structValues() []any {
	c1 := make(chan models.Boolean)
	c2 := make(chan models.Integer)
	c3 := make(chan models.String)
	c4 := make(chan models.Float)
	c5 := make(chan models.Complex)
	c6 := make(chan models.Uint)
	c7 := make(chan models.Channel[any])
	c8 := make(chan models.Slice[any])
	c9 := make(chan models.Map[any, any])
	c10 := make(chan models.NaturallyComparable)
	c11 := make(chan models.Stringer)
	c12 := make(chan models.CloserSuccess)
	c13 := make(chan models.CloserError)
	c14 := make(chan models.CloserReaderSuccess)
	c15 := make(chan models.CloserReaderError)
	c16 := make(chan models.CloserReaderWriterSuccess)
	c17 := make(chan models.CloserReaderWriterError)
	c18 := make(chan models.NotComparable)
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func sliceOfStructValues() []any {
	c1 := make(chan []models.Boolean)
	c2 := make(chan []models.Integer)
	c3 := make(chan []models.String)
	c4 := make(chan []models.Float)
	c5 := make(chan []models.Complex)
	c6 := make(chan []models.Uint)
	c7 := make(chan []models.Channel[any])
	c8 := make(chan []models.Slice[any])
	c9 := make(chan []models.Map[any, any])
	c10 := make(chan []models.NaturallyComparable)
	c11 := make(chan []models.Stringer)
	c12 := make(chan []models.CloserSuccess)
	c13 := make(chan []models.CloserError)
	c14 := make(chan []models.CloserReaderSuccess)
	c15 := make(chan []models.CloserReaderError)
	c16 := make(chan []models.CloserReaderWriterSuccess)
	c17 := make(chan []models.CloserReaderWriterError)
	c18 := make(chan []models.NotComparable)
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func pointerToSliceOfStructValues() []any {
	c1 := make(chan *[]models.Boolean)
	c2 := make(chan *[]models.Integer)
	c3 := make(chan *[]models.String)
	c4 := make(chan *[]models.Float)
	c5 := make(chan *[]models.Complex)
	c6 := make(chan *[]models.Uint)
	c7 := make(chan *[]models.Channel[any])
	c8 := make(chan *[]models.Slice[any])
	c9 := make(chan *[]models.Map[any, any])
	c10 := make(chan *[]models.NaturallyComparable)
	c11 := make(chan *[]models.Stringer)
	c12 := make(chan *[]models.CloserSuccess)
	c13 := make(chan *[]models.CloserError)
	c14 := make(chan *[]models.CloserReaderSuccess)
	c15 := make(chan *[]models.CloserReaderError)
	c16 := make(chan *[]models.CloserReaderWriterSuccess)
	c17 := make(chan *[]models.CloserReaderWriterError)
	c18 := make(chan *[]models.NotComparable)
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func senderValues() []any {
	return []any{
		make(chan<- bool), make(chan<- int), make(chan<- int8), make(chan<- int16),
		make(chan<- int32), make(chan<- int64), make(chan<- uint), make(chan<- uint8),
		make(chan<- uint16), make(chan<- uint32), make(chan<- uint64), make(chan<- uintptr),
		make(chan<- float32), make(chan<- float64), make(chan<- complex64),
		make(chan<- complex128), make(chan<- string),
		make(chan<- chan interface{}), make(chan<- *chan interface{}),
		make(chan<- *interf.DataBool[any]), make(chan<- *interf.Simple), make(chan<- *interf.Public),
		make(chan<- *interf.TwoMethods), make(chan<- *interf.OneMethod),
		make(chan<- *interf.Error), make(chan<- *interf.Seeder[any]),
		make(chan<- *models.Integer), make(chan<- *models.String), make(chan<- *models.Float),
		make(chan<- *models.Complex), make(chan<- *models.Uint), make(chan<- *models.Channel[any]),
		make(chan<- *models.Slice[any]), make(chan<- *models.Map[any, any]),
		make(chan<- *models.NaturallyComparable), make(chan<- *models.Stringer),
		make(chan<- *models.CloserSuccess), make(chan<- *models.CloserError),
		make(chan<- *models.CloserReaderSuccess), make(chan<- *models.CloserReaderError),
		make(chan<- *models.CloserReaderWriterSuccess), make(chan<- *models.CloserReaderWriterError),
		make(chan<- *models.NotComparable), make(chan<- *models.User), make(chan<- *models.Account),
		make(chan<- *models.Address),
	}
}

func receiverValues() []any {
	return []any{
		make(<-chan bool), make(<-chan int), make(<-chan int8), make(<-chan int16),
		make(<-chan int32), make(<-chan int64), make(<-chan uint), make(<-chan uint8),
		make(<-chan uint16), make(<-chan uint32), make(<-chan uint64), make(<-chan uintptr),
		make(<-chan float32), make(<-chan float64), make(<-chan complex64),
		make(<-chan complex128), make(<-chan string),
		make(<-chan chan interface{}), make(<-chan *chan interface{}),
		make(<-chan *interf.DataBool[any]), make(<-chan *interf.Simple), make(<-chan *interf.Public),
		make(<-chan *interf.TwoMethods), make(<-chan *interf.OneMethod),
		make(<-chan *interf.Error), make(<-chan *interf.Seeder[any]),
		make(<-chan *models.Integer), make(<-chan *models.String), make(<-chan *models.Float),
		make(<-chan *models.Complex), make(<-chan *models.Uint), make(<-chan *models.Channel[any]),
		make(<-chan *models.Slice[any]), make(<-chan *models.Map[any, any]),
		make(<-chan *models.NaturallyComparable), make(<-chan *models.Stringer),
		make(<-chan *models.CloserSuccess), make(<-chan *models.CloserError),
		make(<-chan *models.CloserReaderSuccess), make(<-chan *models.CloserReaderError),
		make(<-chan *models.CloserReaderWriterSuccess), make(<-chan *models.CloserReaderWriterError),
		make(<-chan *models.NotComparable), make(<-chan *models.User), make(<-chan *models.Account),
		make(<-chan *models.Address),
	}
}

func edgeValues() []any {
	return []interface{}{
		make(chan chan interface{}), make(chan *chan interface{}), make(chan *unsafe.Pointer), make(chan **unsafe.Pointer),
	}
}
