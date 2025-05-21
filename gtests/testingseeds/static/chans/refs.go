// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chans

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
	"unsafe"
)

func primitivesRefs() []interface{} {
	cb := make(chan bool)
	ci := make(chan int)
	ci8 := make(chan int8)
	ci16 := make(chan int16)
	ci32 := make(chan int32)
	ci64 := make(chan int64)
	cu := make(chan uint)
	cu8 := make(chan uint8)
	cu16 := make(chan uint16)
	cu32 := make(chan uint32)
	cu64 := make(chan uint64)
	cuptr := make(chan uintptr)
	cf32 := make(chan float32)
	cf64 := make(chan float64)
	cc64 := make(chan complex64)
	cc128 := make(chan complex128)
	cs := make(chan string)
	return []interface{}{
		&cb, &ci, &ci8, &ci16, &ci32, &ci64, &cu, &cu8, &cu16, &cu32, &cu64,
		&cuptr, &cf32, &cf64, &cc64, &cc128, &cs,
	}
}

func pointerToPrimitiveRefs() []interface{} {
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
	return []interface{}{
		&cb, &ci, &ci8, &ci16, &ci32, &ci64, &cu, &cu8, &cu16, &cu32, &cu64, &cuptr, &cf32, &cf64, &cc64, &cc128, &cs,
	}
}

func sliceOfPrimitivesRefs() []interface{} {
	cb := make(chan []bool)
	ci := make(chan []int)
	ci8 := make(chan []int8)
	ci16 := make(chan []int16)
	ci32 := make(chan []int32)
	ci64 := make(chan []int64)
	cu := make(chan []uint)
	cu8 := make(chan []uint8)
	cu16 := make(chan []uint16)
	cu32 := make(chan []uint32)
	cu64 := make(chan []uint64)
	cuptr := make(chan []uintptr)
	cf32 := make(chan []float32)
	cf64 := make(chan []float64)
	cc64 := make(chan []complex64)
	cc128 := make(chan []complex128)
	cs := make(chan []string)
	return []interface{}{
		&cb, &ci, &ci8, &ci16, &ci32, &ci64, &cu, &cu8, &cu16, &cu32, &cu64,
		&cuptr, &cf32, &cf64, &cc64, &cc128, &cs,
	}
}

func sliceOfRefToPrimitivesRefs() []interface{} {
	cb := make(chan []*bool)
	ci := make(chan []*int)
	ci8 := make(chan []*int8)
	ci16 := make(chan []*int16)
	ci32 := make(chan []*int32)
	ci64 := make(chan []*int64)
	cu := make(chan []*uint)
	cu8 := make(chan []*uint8)
	cu16 := make(chan []*uint16)
	cu32 := make(chan []*uint32)
	cu64 := make(chan []*uint64)
	cuptr := make(chan []*uintptr)
	cf32 := make(chan []*float32)
	cf64 := make(chan []*float64)
	cc64 := make(chan []*complex64)
	cc128 := make(chan []*complex128)
	cs := make(chan []*string)
	return []interface{}{
		&cb, &ci, &ci8, &ci16, &ci32, &ci64, &cu, &cu8, &cu16, &cu32, &cu64,
		&cuptr, &cf32, &cf64, &cc64, &cc128, &cs,
	}
}

func interfaceRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12,
	}
}

func sliceOfInterfaceRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12,
	}
}

func sliceOfRefToInterfaceRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12,
	}
}

func pointerToInterfaceRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12,
	}
}

func ofStructsRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18,
	}
}

func sliceOfStructRefs() []any {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18,
	}
}

func pointerToSliceOfStructRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18,
	}
}

func sendersAsRefs() []any {
	c1 := make(chan<- bool)
	c2 := make(chan<- int)
	c3 := make(chan<- int8)
	c4 := make(chan<- int16)
	c5 := make(chan<- int32)
	c6 := make(chan<- int64)
	c7 := make(chan<- uint)
	c8 := make(chan<- uint8)
	c9 := make(chan<- uint16)
	c10 := make(chan<- uint32)
	c11 := make(chan<- uint64)
	c12 := make(chan<- uintptr)
	c13 := make(chan<- float32)
	c14 := make(chan<- float64)
	c15 := make(chan<- complex64)
	c16 := make(chan<- complex128)
	c17 := make(chan<- string)
	c18 := make(chan<- chan interface{})
	c19 := make(chan<- *chan interface{})
	c20 := make(chan<- *interf.DataBool[any])
	c21 := make(chan<- *interf.Simple)
	c22 := make(chan<- *interf.Public)
	c23 := make(chan<- *interf.TwoMethods)
	c24 := make(chan<- *interf.OneMethod)
	c25 := make(chan<- *interf.Error)
	c26 := make(chan<- *interf.Seeder[any])
	c27 := make(chan<- *structs.Integer)
	c28 := make(chan<- *structs.String)
	c29 := make(chan<- *structs.Float)
	c30 := make(chan<- *structs.Complex)
	c31 := make(chan<- *structs.Uint)
	c32 := make(chan<- *structs.Channel[any])
	c33 := make(chan<- *structs.Slice[any])
	c34 := make(chan<- *structs.Map[any, any])
	c35 := make(chan<- *structs.NaturallyComparable)
	c36 := make(chan<- *structs.Stringer)
	c37 := make(chan<- *structs.CloserSuccess)
	c38 := make(chan<- *structs.CloserError)
	c39 := make(chan<- *structs.CloserReaderSuccess)
	c40 := make(chan<- *structs.CloserReaderError)
	c41 := make(chan<- *structs.CloserReaderWriterSuccess)
	c42 := make(chan<- *structs.CloserReaderWriterError)
	c43 := make(chan<- *structs.NotComparable)
	c44 := make(chan<- *static.User)
	c45 := make(chan<- *static.Account)
	c46 := make(chan<- *static.Address)
	return []any{
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18, &c19, &c20, &c21, &c22, &c23, &c24, &c25, &c26, &c27, &c28, &c29, &c30,
		&c31, &c32, &c33, &c34, &c35, &c36, &c37, &c38, &c39, &c40, &c41, &c42, &c43, &c44, &c45, &c46,
	}
}

func receiversAsRefs() []any {
	c1 := make(<-chan bool)
	c2 := make(<-chan int)
	c3 := make(<-chan int8)
	c4 := make(<-chan int16)
	c5 := make(<-chan int32)
	c6 := make(<-chan int64)
	c7 := make(<-chan uint)
	c8 := make(<-chan uint8)
	c9 := make(<-chan uint16)
	c10 := make(<-chan uint32)
	c11 := make(<-chan uint64)
	c12 := make(<-chan uintptr)
	c13 := make(<-chan float32)
	c14 := make(<-chan float64)
	c15 := make(<-chan complex64)
	c16 := make(<-chan complex128)
	c17 := make(<-chan string)
	c18 := make(<-chan chan interface{})
	c19 := make(<-chan *chan interface{})
	c20 := make(<-chan *interf.DataBool[any])
	c21 := make(<-chan *interf.Simple)
	c22 := make(<-chan *interf.Public)
	c23 := make(<-chan *interf.TwoMethods)
	c24 := make(<-chan *interf.OneMethod)
	c25 := make(<-chan *interf.Error)
	c26 := make(<-chan *interf.Seeder[any])
	c27 := make(<-chan *structs.Integer)
	c28 := make(<-chan *structs.String)
	c29 := make(<-chan *structs.Float)
	c30 := make(<-chan *structs.Complex)
	c31 := make(<-chan *structs.Uint)
	c32 := make(<-chan *structs.Channel[any])
	c33 := make(<-chan *structs.Slice[any])
	c34 := make(<-chan *structs.Map[any, any])
	c35 := make(<-chan *structs.NaturallyComparable)
	c36 := make(<-chan *structs.Stringer)
	c37 := make(<-chan *structs.CloserSuccess)
	c38 := make(<-chan *structs.CloserError)
	c39 := make(<-chan *structs.CloserReaderSuccess)
	c40 := make(<-chan *structs.CloserReaderError)
	c41 := make(<-chan *structs.CloserReaderWriterSuccess)
	c42 := make(<-chan *structs.CloserReaderWriterError)
	c43 := make(<-chan *structs.NotComparable)
	c44 := make(<-chan *static.User)
	c45 := make(<-chan *static.Account)
	c46 := make(<-chan *static.Address)
	return []any{
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18, &c19, &c20, &c21, &c22, &c23, &c24, &c25, &c26, &c27, &c28, &c29, &c30,
		&c31, &c32, &c33, &c34, &c35, &c36, &c37, &c38, &c39, &c40, &c41, &c42, &c43, &c44, &c45, &c46,
	}
}

func edgeRefs() []interface{} {
	c1 := make(chan chan interface{})
	c2 := make(chan *chan interface{})
	c3 := make(chan *unsafe.Pointer)
	c4 := make(chan **unsafe.Pointer)
	return []interface{}{
		&c1, &c2, &c3, &c4,
	}
}

func ofRefs() []interface{} {
	return slices.Concat(
		primitivesRefs(), pointerToPrimitiveRefs(),
		sliceOfPrimitivesRefs(), sliceOfRefToPrimitivesRefs(), interfaceRefs(),
		pointerToInterfaceRefs(), sliceOfInterfaceRefs(), sliceOfRefToInterfaceRefs(),
		ofStructsRefs(), sliceOfStructRefs(), pointerToSliceOfStructRefs(), edgeRefs(),
		sendersAsRefs(), receiversAsRefs(),
	)
}
