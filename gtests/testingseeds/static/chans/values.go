// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chans

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
	"unsafe"
)

func primitives() []interface{} {
	return []interface{}{
		make(chan bool), make(chan int), make(chan int8), make(chan int16),
		make(chan int32), make(chan int64), make(chan uint), make(chan uint8),
		make(chan uint16), make(chan uint32), make(chan uint64), make(chan uintptr),
		make(chan float32), make(chan float64), make(chan complex64),
		make(chan complex128), make(chan string),
	}
}

func pointerToPrimitives() []any {
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

func sliceOfPrimitives() []interface{} {
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

func sliceOfPointerToPrimitives() []interface{} {
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

func interfaces() []interface{} {
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

func pointerToInterfaces() []interface{} {
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

func sliceOfInterfaces() []interface{} {
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

func sliceOfPointerToInterfaces() []any {
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

func ofStructs() []interface{} {
	c1 := make(chan structs.Boolean)
	c2 := make(chan structs.Integer)
	c3 := make(chan structs.String)
	c4 := make(chan structs.Float)
	c5 := make(chan structs.Complex)
	c6 := make(chan structs.Uint)
	c7 := make(chan structs.Channel[any])
	c8 := make(chan structs.Slice[any])
	c9 := make(chan structs.Map[any, any])
	c10 := make(chan structs.NaturallyComparable)
	c11 := make(chan structs.Stringer)
	c12 := make(chan structs.CloserSuccess)
	c13 := make(chan structs.CloserError)
	c14 := make(chan structs.CloserReaderSuccess)
	c15 := make(chan structs.CloserReaderError)
	c16 := make(chan structs.CloserReaderWriterSuccess)
	c17 := make(chan structs.CloserReaderWriterError)
	c18 := make(chan structs.NotComparable)
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func sliceOfStructs() []interface{} {
	c1 := make(chan []structs.Boolean)
	c2 := make(chan []structs.Integer)
	c3 := make(chan []structs.String)
	c4 := make(chan []structs.Float)
	c5 := make(chan []structs.Complex)
	c6 := make(chan []structs.Uint)
	c7 := make(chan []structs.Channel[any])
	c8 := make(chan []structs.Slice[any])
	c9 := make(chan []structs.Map[any, any])
	c10 := make(chan []structs.NaturallyComparable)
	c11 := make(chan []structs.Stringer)
	c12 := make(chan []structs.CloserSuccess)
	c13 := make(chan []structs.CloserError)
	c14 := make(chan []structs.CloserReaderSuccess)
	c15 := make(chan []structs.CloserReaderError)
	c16 := make(chan []structs.CloserReaderWriterSuccess)
	c17 := make(chan []structs.CloserReaderWriterError)
	c18 := make(chan []structs.NotComparable)
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func pointerToSliceOfStructs() []interface{} {
	c1 := make(chan *[]structs.Boolean)
	c2 := make(chan *[]structs.Integer)
	c3 := make(chan *[]structs.String)
	c4 := make(chan *[]structs.Float)
	c5 := make(chan *[]structs.Complex)
	c6 := make(chan *[]structs.Uint)
	c7 := make(chan *[]structs.Channel[any])
	c8 := make(chan *[]structs.Slice[any])
	c9 := make(chan *[]structs.Map[any, any])
	c10 := make(chan *[]structs.NaturallyComparable)
	c11 := make(chan *[]structs.Stringer)
	c12 := make(chan *[]structs.CloserSuccess)
	c13 := make(chan *[]structs.CloserError)
	c14 := make(chan *[]structs.CloserReaderSuccess)
	c15 := make(chan *[]structs.CloserReaderError)
	c16 := make(chan *[]structs.CloserReaderWriterSuccess)
	c17 := make(chan *[]structs.CloserReaderWriterError)
	c18 := make(chan *[]structs.NotComparable)
	return []any{
		c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15,
		c16, c17, c18,
	}
}

func senders() []any {
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
		make(chan<- *structs.Integer), make(chan<- *structs.String), make(chan<- *structs.Float),
		make(chan<- *structs.Complex), make(chan<- *structs.Uint), make(chan<- *structs.Channel[any]),
		make(chan<- *structs.Slice[any]), make(chan<- *structs.Map[any, any]),
		make(chan<- *structs.NaturallyComparable), make(chan<- *structs.Stringer),
		make(chan<- *structs.CloserSuccess), make(chan<- *structs.CloserError),
		make(chan<- *structs.CloserReaderSuccess), make(chan<- *structs.CloserReaderError),
		make(chan<- *structs.CloserReaderWriterSuccess), make(chan<- *structs.CloserReaderWriterError),
		make(chan<- *structs.NotComparable), make(chan<- *static.User), make(chan<- *static.Account),
		make(chan<- *static.Address),
	}
}

func receivers() []any {
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
		make(<-chan *structs.Integer), make(<-chan *structs.String), make(<-chan *structs.Float),
		make(<-chan *structs.Complex), make(<-chan *structs.Uint), make(<-chan *structs.Channel[any]),
		make(<-chan *structs.Slice[any]), make(<-chan *structs.Map[any, any]),
		make(<-chan *structs.NaturallyComparable), make(<-chan *structs.Stringer),
		make(<-chan *structs.CloserSuccess), make(<-chan *structs.CloserError),
		make(<-chan *structs.CloserReaderSuccess), make(<-chan *structs.CloserReaderError),
		make(<-chan *structs.CloserReaderWriterSuccess), make(<-chan *structs.CloserReaderWriterError),
		make(<-chan *structs.NotComparable), make(<-chan *static.User), make(<-chan *static.Account),
		make(<-chan *static.Address),
	}
}

func edgeValues() []interface{} {
	return []interface{}{
		make(chan chan interface{}), make(chan *chan interface{}), make(chan *unsafe.Pointer), make(chan **unsafe.Pointer),
	}
}

func ofValues() []interface{} {
	return slices.Concat(edgeValues(), interfaces(), pointerToInterfaces(), primitives(), sliceOfPrimitives(),
		sliceOfPointerToPrimitives(), ofStructs(), sliceOfStructs(),
		sliceOfPointerToInterfaces(), sliceOfInterfaces(), senders(), receivers(),
	)
}
