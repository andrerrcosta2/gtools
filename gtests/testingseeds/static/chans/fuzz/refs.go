// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
	"unsafe"
)

func referencesSets() []interface{} {
	return slices.Concat(
		primitiveRefs(), pointerToPrimitiveRefs(),
		sliceOfPrimitiveRefs(), sliceOfPointerToPrimitiveRefs(), interfaceRefs(),
		pointerToInterfaceRefs(), sliceOfInterfaceRefs(), sliceOfPointerToInterfaceRefs(),
		structRefs(), sliceOfStructRefs(), pointerToSliceOfStructRefs(), edgeRefs(),
		senderRefs(), receiverRefs(),
	)
}

func primitiveRefs() []interface{} {
	cb := random.SingleOf[chan bool]()
	ci := random.SingleOf[chan int]()
	ci8 := random.SingleOf[chan int8]()
	ci16 := random.SingleOf[chan int16]()
	ci32 := random.SingleOf[chan int32]()
	ci64 := random.SingleOf[chan int64]()
	cu := random.SingleOf[chan uint]()
	cu8 := random.SingleOf[chan uint8]()
	cu16 := random.SingleOf[chan uint16]()
	cu32 := random.SingleOf[chan uint32]()
	cu64 := random.SingleOf[chan uint64]()
	cuptr := random.SingleOf[chan uintptr]()
	cf32 := random.SingleOf[chan float32]()
	cf64 := random.SingleOf[chan float64]()
	cc64 := random.SingleOf[chan complex64]()
	cc128 := random.SingleOf[chan complex128]()
	cs := random.SingleOf[chan string]()
	return []interface{}{
		&cb, &ci, &ci8, &ci16, &ci32, &ci64, &cu, &cu8, &cu16, &cu32, &cu64,
		&cuptr, &cf32, &cf64, &cc64, &cc128, &cs,
	}
}

func pointerToPrimitiveRefs() []interface{} {
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
	return []interface{}{
		&cb, &ci, &ci8, &ci16, &ci32, &ci64, &cu, &cu8, &cu16, &cu32, &cu64, &cuptr, &cf32, &cf64, &cc64, &cc128, &cs,
	}
}

func sliceOfPrimitiveRefs() []interface{} {
	cb := random.SingleOf[chan []bool]()
	ci := random.SingleOf[chan []int]()
	ci8 := random.SingleOf[chan []int8]()
	ci16 := random.SingleOf[chan []int16]()
	ci32 := random.SingleOf[chan []int32]()
	ci64 := random.SingleOf[chan []int64]()
	cu := random.SingleOf[chan []uint]()
	cu8 := random.SingleOf[chan []uint8]()
	cu16 := random.SingleOf[chan []uint16]()
	cu32 := random.SingleOf[chan []uint32]()
	cu64 := random.SingleOf[chan []uint64]()
	cuptr := random.SingleOf[chan []uintptr]()
	cf32 := random.SingleOf[chan []float32]()
	cf64 := random.SingleOf[chan []float64]()
	cc64 := random.SingleOf[chan []complex64]()
	cc128 := random.SingleOf[chan []complex128]()
	cs := random.SingleOf[chan []string]()
	return []interface{}{
		&cb, &ci, &ci8, &ci16, &ci32, &ci64, &cu, &cu8, &cu16, &cu32, &cu64,
		&cuptr, &cf32, &cf64, &cc64, &cc128, &cs,
	}
}

func sliceOfPointerToPrimitiveRefs() []interface{} {
	cb := random.SingleOf[chan []*bool]()
	ci := random.SingleOf[chan []*int]()
	ci8 := random.SingleOf[chan []*int8]()
	ci16 := random.SingleOf[chan []*int16]()
	ci32 := random.SingleOf[chan []*int32]()
	ci64 := random.SingleOf[chan []*int64]()
	cu := random.SingleOf[chan []*uint]()
	cu8 := random.SingleOf[chan []*uint8]()
	cu16 := random.SingleOf[chan []*uint16]()
	cu32 := random.SingleOf[chan []*uint32]()
	cu64 := random.SingleOf[chan []*uint64]()
	cuptr := random.SingleOf[chan []*uintptr]()
	cf32 := random.SingleOf[chan []*float32]()
	cf64 := random.SingleOf[chan []*float64]()
	cc64 := random.SingleOf[chan []*complex64]()
	cc128 := random.SingleOf[chan []*complex128]()
	cs := random.SingleOf[chan []*string]()
	return []interface{}{
		&cb, &ci, &ci8, &ci16, &ci32, &ci64, &cu, &cu8, &cu16, &cu32, &cu64,
		&cuptr, &cf32, &cf64, &cc64, &cc128, &cs,
	}
}

func interfaceRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12,
	}
}

func sliceOfInterfaceRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12,
	}
}

func sliceOfPointerToInterfaceRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12,
	}
}

func pointerToInterfaceRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12,
	}
}

func structRefs() []interface{} {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18,
	}
}

func sliceOfStructRefs() []any {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18,
	}
}

func pointerToSliceOfStructRefs() []any {
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
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18,
	}
}

func senderRefs() []any {
	c1 := random.SingleOf[chan<- bool]()
	c2 := random.SingleOf[chan<- int]()
	c3 := random.SingleOf[chan<- int8]()
	c4 := random.SingleOf[chan<- int16]()
	c5 := random.SingleOf[chan<- int32]()
	c6 := random.SingleOf[chan<- int64]()
	c7 := random.SingleOf[chan<- uint]()
	c8 := random.SingleOf[chan<- uint8]()
	c9 := random.SingleOf[chan<- uint16]()
	c10 := random.SingleOf[chan<- uint32]()
	c11 := random.SingleOf[chan<- uint64]()
	c12 := random.SingleOf[chan<- uintptr]()
	c13 := random.SingleOf[chan<- float32]()
	c14 := random.SingleOf[chan<- float64]()
	c15 := random.SingleOf[chan<- complex64]()
	c16 := random.SingleOf[chan<- complex128]()
	c17 := random.SingleOf[chan<- string]()
	c18 := random.SingleOf[chan<- chan interface{}]()
	c19 := random.SingleOf[chan<- *chan interface{}]()
	c20 := random.SingleOf[chan<- *interf.DataBool[any]]()
	c21 := random.SingleOf[chan<- *interf.Simple]()
	c22 := random.SingleOf[chan<- *interf.Public]()
	c23 := random.SingleOf[chan<- *interf.TwoMethods]()
	c24 := random.SingleOf[chan<- *interf.OneMethod]()
	c25 := random.SingleOf[chan<- *interf.Error]()
	c26 := random.SingleOf[chan<- *interf.Seeder[any]]()
	c27 := random.SingleOf[chan<- *models.Integer]()
	c28 := random.SingleOf[chan<- *models.String]()
	c29 := random.SingleOf[chan<- *models.Float]()
	c30 := random.SingleOf[chan<- *models.Complex]()
	c31 := random.SingleOf[chan<- *models.Uint]()
	c32 := random.SingleOf[chan<- *models.Channel[any]]()
	c33 := random.SingleOf[chan<- *models.Slice[any]]()
	c34 := random.SingleOf[chan<- *models.Map[any, any]]()
	c35 := random.SingleOf[chan<- *models.NaturallyComparable]()
	c36 := random.SingleOf[chan<- *models.Stringer]()
	c37 := random.SingleOf[chan<- *models.CloserSuccess]()
	c38 := random.SingleOf[chan<- *models.CloserError]()
	c39 := random.SingleOf[chan<- *models.CloserReaderSuccess]()
	c40 := random.SingleOf[chan<- *models.CloserReaderError]()
	c41 := random.SingleOf[chan<- *models.CloserReaderWriterSuccess]()
	c42 := random.SingleOf[chan<- *models.CloserReaderWriterError]()
	c43 := random.SingleOf[chan<- *models.NotComparable]()
	c44 := random.SingleOf[chan<- *models.User]()
	c45 := random.SingleOf[chan<- *models.Account]()
	c46 := random.SingleOf[chan<- *models.Address]()
	return []any{
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18, &c19, &c20, &c21, &c22, &c23, &c24, &c25, &c26, &c27, &c28, &c29, &c30,
		&c31, &c32, &c33, &c34, &c35, &c36, &c37, &c38, &c39, &c40, &c41, &c42, &c43, &c44, &c45, &c46,
	}
}

func receiverRefs() []any {
	c1 := random.SingleOf[<-chan bool]()
	c2 := random.SingleOf[<-chan int]()
	c3 := random.SingleOf[<-chan int8]()
	c4 := random.SingleOf[<-chan int16]()
	c5 := random.SingleOf[<-chan int32]()
	c6 := random.SingleOf[<-chan int64]()
	c7 := random.SingleOf[<-chan uint]()
	c8 := random.SingleOf[<-chan uint8]()
	c9 := random.SingleOf[<-chan uint16]()
	c10 := random.SingleOf[<-chan uint32]()
	c11 := random.SingleOf[<-chan uint64]()
	c12 := random.SingleOf[<-chan uintptr]()
	c13 := random.SingleOf[<-chan float32]()
	c14 := random.SingleOf[<-chan float64]()
	c15 := random.SingleOf[<-chan complex64]()
	c16 := random.SingleOf[<-chan complex128]()
	c17 := random.SingleOf[<-chan string]()
	c18 := random.SingleOf[<-chan chan interface{}]()
	c19 := random.SingleOf[<-chan *chan interface{}]()
	c20 := random.SingleOf[<-chan *interf.DataBool[any]]()
	c21 := random.SingleOf[<-chan *interf.Simple]()
	c22 := random.SingleOf[<-chan *interf.Public]()
	c23 := random.SingleOf[<-chan *interf.TwoMethods]()
	c24 := random.SingleOf[<-chan *interf.OneMethod]()
	c25 := random.SingleOf[<-chan *interf.Error]()
	c26 := random.SingleOf[<-chan *interf.Seeder[any]]()
	c27 := random.SingleOf[<-chan *models.Integer]()
	c28 := random.SingleOf[<-chan *models.String]()
	c29 := random.SingleOf[<-chan *models.Float]()
	c30 := random.SingleOf[<-chan *models.Complex]()
	c31 := random.SingleOf[<-chan *models.Uint]()
	c32 := random.SingleOf[<-chan *models.Channel[any]]()
	c33 := random.SingleOf[<-chan *models.Slice[any]]()
	c34 := random.SingleOf[<-chan *models.Map[any, any]]()
	c35 := random.SingleOf[<-chan *models.NaturallyComparable]()
	c36 := random.SingleOf[<-chan *models.Stringer]()
	c37 := random.SingleOf[<-chan *models.CloserSuccess]()
	c38 := random.SingleOf[<-chan *models.CloserError]()
	c39 := random.SingleOf[<-chan *models.CloserReaderSuccess]()
	c40 := random.SingleOf[<-chan *models.CloserReaderError]()
	c41 := random.SingleOf[<-chan *models.CloserReaderWriterSuccess]()
	c42 := random.SingleOf[<-chan *models.CloserReaderWriterError]()
	c43 := random.SingleOf[<-chan *models.NotComparable]()
	c44 := random.SingleOf[<-chan *models.User]()
	c45 := random.SingleOf[<-chan *models.Account]()
	c46 := random.SingleOf[<-chan *models.Address]()
	return []any{
		&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15,
		&c16, &c17, &c18, &c19, &c20, &c21, &c22, &c23, &c24, &c25, &c26, &c27, &c28, &c29, &c30,
		&c31, &c32, &c33, &c34, &c35, &c36, &c37, &c38, &c39, &c40, &c41, &c42, &c43, &c44, &c45, &c46,
	}
}

func edgeRefs() []interface{} {
	c1 := random.SingleOf[chan chan interface{}]()
	c2 := random.SingleOf[chan *chan interface{}]()
	c3 := random.SingleOf[chan *unsafe.Pointer]()
	c4 := random.SingleOf[chan **unsafe.Pointer]()
	return []interface{}{
		&c1, &c2, &c3, &c4,
	}
}
