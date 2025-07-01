// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
	"unsafe"
)

func referencesSet() []any {
	return slices.Concat(EdgeTypesRefs(), pointerToChannelRefs(),
		channelRefs(), pointerToStructRefs(), pointerToInterfaceRefs(),
		interfaceRefs(), primitiveRefs(), pointerToPrimitiveRefs(),
		structRefs())
}

func primitiveRefs() []any {
	// The only way to generate random sizes randomically for arrays
	// is through reflection. But the output should be a "reflect.Value"
	// to be converted inline.
	// I will work here a defined size.
	var bv = random.SingleOf[[4]bool]()
	var iv = random.SingleOf[[4]int]()
	var i8v = random.SingleOf[[4]int8]()
	var i16v = random.SingleOf[[4]int16]()
	var i32v = random.SingleOf[[4]int32]()
	var i64v = random.SingleOf[[4]int64]()
	var uiv = random.SingleOf[[4]int]()
	var u8v = random.SingleOf[[4]int8]()
	var u16v = random.SingleOf[[4]int16]()
	var u32v = random.SingleOf[[4]int32]()
	var u64v = random.SingleOf[[4]int64]()
	var f32v = random.SingleOf[[4]int]()
	var f64v = random.SingleOf[[4]float64]()
	var c64v = random.SingleOf[[4]complex64]()
	var c128v = random.SingleOf[[4]complex128]()
	var sv = random.SingleOf[[4]string]()
	return []interface{}{&bv, &iv, &i8v, &i16v, &i32v, &i64v, &uiv, &u8v, &u16v,
		&u32v, &u64v, &f32v, &f64v, &c64v, &c128v, &sv,
	}
}

func pointerToPrimitiveRefs() []any {
	// The only way to generate random sizes randomically for arrays
	// is through reflection. But the output should be a "reflect.Value"
	// to be converted inline.
	// I will work here a defined size.
	var bv = random.SingleOf[[4]*bool]()
	var iv = random.SingleOf[[4]*int]()
	var i8v = random.SingleOf[[4]*int8]()
	var i16v = random.SingleOf[[4]*int16]()
	var i32v = random.SingleOf[[4]*int32]()
	var i64v = random.SingleOf[[4]*int64]()
	var uiv = random.SingleOf[[4]*int]()
	var u8v = random.SingleOf[[4]*int8]()
	var u16v = random.SingleOf[[4]*int16]()
	var u32v = random.SingleOf[[4]*int32]()
	var u64v = random.SingleOf[[4]*int64]()
	var f32v = random.SingleOf[[4]*int]()
	var f64v = random.SingleOf[[4]*float64]()
	var c64v = random.SingleOf[[4]*complex64]()
	var c128v = random.SingleOf[[4]*complex128]()
	var sv = random.SingleOf[[4]*string]()
	return []interface{}{&bv, &iv, &i8v, &i16v, &i32v, &i64v, &uiv, &u8v, &u16v,
		&u32v, &u64v, &f32v, &f64v, &c64v, &c128v, &sv,
	}
}

func interfaceRefs() []any {
	var i1 = random.SingleOf[[4]interf.Closer]()
	var i2 = random.SingleOf[[4]interf.CloserReader]()
	var i3 = random.SingleOf[[4]interf.CloserReaderWriter]()
	var i4 = random.SingleOf[[4]interf.Error]()
	var i5 = random.SingleOf[[4]interf.OneMethod]()
	var i6 = random.SingleOf[[4]interf.Public]()
	var i7 = random.SingleOf[[4]interf.Seeder[any]]()
	var i8 = random.SingleOf[[4]interf.Simple]()
	var i9 = random.SingleOf[[4]interf.Stringer]()
	var i10 = random.SingleOf[[4]interf.TwoMethods]()
	var i11 = random.SingleOf[[4]interf.WithoutImplementation]()
	return []any{
		&i1, &i2, &i3, &i4, &i5, &i6, &i7, &i8, &i9, &i10, &i11,
	}
}

func pointerToInterfaceRefs() []interface{} {
	var i1 = random.SingleOf[[4]*interf.Closer]()
	var i2 = random.SingleOf[[4]*interf.CloserReader]()
	var i3 = random.SingleOf[[4]*interf.CloserReaderWriter]()
	var i4 = random.SingleOf[[4]*interf.Error]()
	var i5 = random.SingleOf[[4]*interf.OneMethod]()
	var i6 = random.SingleOf[[4]*interf.Public]()
	var i7 = random.SingleOf[[4]*interf.Seeder[any]]()
	var i8 = random.SingleOf[[4]*interf.Simple]()
	var i9 = random.SingleOf[[4]*interf.Stringer]()
	var i10 = random.SingleOf[[4]*interf.TwoMethods]()
	var i11 = random.SingleOf[[4]*interf.WithoutImplementation]()
	return []any{
		&i1, &i2, &i3, &i4, &i5, &i6, &i7, &i8, &i9, &i10, &i11,
	}
}

func structRefs() []any {
	a1 := random.SingleOf[[4]models.Boolean]()
	a2 := random.SingleOf[[4]models.Integer]()
	a3 := random.SingleOf[[4]models.Uint]()
	a4 := random.SingleOf[[4]models.Float]()
	a5 := random.SingleOf[[4]models.Complex]()
	a6 := random.SingleOf[[4]models.String]()
	a7 := random.SingleOf[[4]models.Channel[any]]()
	a8 := random.SingleOf[[4]models.Slice[any]]()
	a9 := random.SingleOf[[4]models.Map[any, any]]()
	a10 := random.SingleOf[[4]models.Function[any]]()
	a11 := random.SingleOf[[4]models.Simple]()
	a12 := random.SingleOf[[4]models.Public]()
	a13 := random.SingleOf[[4]models.CloserSuccess]()
	a14 := random.SingleOf[[4]models.CloserError]()
	a15 := random.SingleOf[[4]models.CloserReaderSuccess]()
	a16 := random.SingleOf[[4]models.CloserReaderError]()
	a17 := random.SingleOf[[4]models.CloserReaderWriterSuccess]()
	a18 := random.SingleOf[[4]models.CloserReaderWriterError]()
	a19 := random.SingleOf[[4]models.Stringer]()
	a20 := random.SingleOf[[4]models.StringerBytes]()
	a21 := random.SingleOf[[4]models.StringerString]()
	a22 := random.SingleOf[[4]models.NaturallyComparable]()
	a23 := random.SingleOf[[4]models.NaturallyComparableWithMethods]()
	a24 := random.SingleOf[[4]models.NotComparable]()
	a25 := random.SingleOf[[4]models.OneData]()
	a26 := random.SingleOf[[4]models.TwoData]()
	a27 := random.SingleOf[[4]models.User]()
	a28 := random.SingleOf[[4]models.Address]()
	a29 := random.SingleOf[[4]models.Credential]()
	a30 := random.SingleOf[[4]models.Account]()
	a31 := random.SingleOf[[4]models.Product]()
	a32 := random.SingleOf[[4]models.Profile]()
	return []any{
		&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8, &a9, &a10,
		&a11, &a12, &a13, &a14, &a15, &a16, &a17, &a18,
		&a19, &a20, &a21, &a22, &a23, &a24, &a25, &a26,
		&a27, &a28, &a29, &a30, &a31, &a32,
	}
}

func pointerToStructRefs() []interface{} {
	a1 := random.SingleOf[[4]*models.Boolean]()
	a2 := random.SingleOf[[4]*models.Integer]()
	a3 := random.SingleOf[[4]*models.Uint]()
	a4 := random.SingleOf[[4]*models.Float]()
	a5 := random.SingleOf[[4]*models.Complex]()
	a6 := random.SingleOf[[4]*models.String]()
	a7 := random.SingleOf[[4]*models.Channel[any]]()
	a8 := random.SingleOf[[4]*models.Slice[any]]()
	a9 := random.SingleOf[[4]*models.Map[any, any]]()
	a10 := random.SingleOf[[4]*models.Function[any]]()
	a11 := random.SingleOf[[4]*models.Simple]()
	a12 := random.SingleOf[[4]*models.Public]()
	a13 := random.SingleOf[[4]*models.CloserSuccess]()
	a14 := random.SingleOf[[4]*models.CloserError]()
	a15 := random.SingleOf[[4]*models.CloserReaderSuccess]()
	a16 := random.SingleOf[[4]*models.CloserReaderError]()
	a17 := random.SingleOf[[4]*models.CloserReaderWriterSuccess]()
	a18 := random.SingleOf[[4]*models.CloserReaderWriterError]()
	a19 := random.SingleOf[[4]*models.Stringer]()
	a20 := random.SingleOf[[4]*models.StringerBytes]()
	a21 := random.SingleOf[[4]*models.StringerString]()
	a22 := random.SingleOf[[4]*models.NaturallyComparable]()
	a23 := random.SingleOf[[4]*models.NaturallyComparableWithMethods]()
	a24 := random.SingleOf[[4]*models.NotComparable]()
	a25 := random.SingleOf[[4]*models.OneData]()
	a26 := random.SingleOf[[4]*models.TwoData]()
	a27 := random.SingleOf[[4]*models.User]()
	a28 := random.SingleOf[[4]*models.Address]()
	a29 := random.SingleOf[[4]*models.Credential]()
	a30 := random.SingleOf[[4]*models.Account]()
	a31 := random.SingleOf[[4]*models.Product]()
	a32 := random.SingleOf[[4]*models.Profile]()
	return []any{
		&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8, &a9, &a10,
		&a11, &a12, &a13, &a14, &a15, &a16, &a17, &a18,
		&a19, &a20, &a21, &a22, &a23, &a24, &a25, &a26,
		&a27, &a28, &a29, &a30, &a31, &a32,
	}
}

func channelRefs() []interface{} {
	a1 := random.SingleOf[[4]chan models.Boolean]()
	a2 := random.SingleOf[[4]chan models.Integer]()
	a3 := random.SingleOf[[4]chan models.Uint]()
	a4 := random.SingleOf[[4]chan models.Float]()
	a5 := random.SingleOf[[4]chan models.Complex]()
	a6 := random.SingleOf[[4]chan models.String]()
	a7 := random.SingleOf[[4]chan models.Channel[any]]()
	a8 := random.SingleOf[[4]chan models.Slice[any]]()
	a9 := random.SingleOf[[4]chan models.Map[any, any]]()
	a10 := random.SingleOf[[4]chan models.Function[any]]()
	a11 := random.SingleOf[[4]chan models.Simple]()
	a12 := random.SingleOf[[4]chan models.Public]()
	a13 := random.SingleOf[[4]chan models.CloserSuccess]()
	a14 := random.SingleOf[[4]chan models.CloserError]()
	a15 := random.SingleOf[[4]chan models.CloserReaderSuccess]()
	a16 := random.SingleOf[[4]chan models.CloserReaderError]()
	a17 := random.SingleOf[[4]chan models.Product]()
	a18 := random.SingleOf[[4]chan models.Credential]()
	a19 := random.SingleOf[[4]chan bool]()
	a20 := random.SingleOf[[4]chan int]()
	a21 := random.SingleOf[[4]chan int8]()
	a22 := random.SingleOf[[4]chan int16]()
	a23 := random.SingleOf[[4]chan int32]()
	a24 := random.SingleOf[[4]chan int64]()
	a25 := random.SingleOf[[4]chan uint]()
	a26 := random.SingleOf[[4]chan uint8]()
	a27 := random.SingleOf[[4]chan uint16]()
	a28 := random.SingleOf[[4]chan uint32]()
	a29 := random.SingleOf[[4]chan uint64]()
	a30 := random.SingleOf[[4]chan uintptr]()
	a31 := random.SingleOf[[4]chan float32]()
	a32 := random.SingleOf[[4]chan float64]()
	a33 := random.SingleOf[[4]chan complex64]()
	a34 := random.SingleOf[[4]chan complex128]()
	a35 := random.SingleOf[[4]chan string]()

	return []any{
		&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8, &a9, &a10,
		&a11, &a12, &a13, &a14, &a15, &a16, &a17, &a18,
		&a19, &a20, &a21, &a22, &a23, &a24, &a25, &a26,
		&a27, &a28, &a29, &a30, &a31, &a32, &a33, &a34, &a35,
	}
}

func pointerToChannelRefs() []interface{} {
	a1 := random.SingleOf[[4]*chan models.Boolean]()
	a2 := random.SingleOf[[4]*chan models.Integer]()
	a3 := random.SingleOf[[4]*chan models.Uint]()
	a4 := random.SingleOf[[4]*chan models.Float]()
	a5 := random.SingleOf[[4]*chan models.Complex]()
	a6 := random.SingleOf[[4]*chan models.String]()
	a7 := random.SingleOf[[4]*chan models.Channel[any]]()
	a8 := random.SingleOf[[4]*chan models.Slice[any]]()
	a9 := random.SingleOf[[4]*chan models.Map[any, any]]()
	a10 := random.SingleOf[[4]*chan models.Function[any]]()
	a11 := random.SingleOf[[4]*chan models.Simple]()
	a12 := random.SingleOf[[4]*chan models.Public]()
	a13 := random.SingleOf[[4]*chan models.CloserSuccess]()
	a14 := random.SingleOf[[4]*chan models.CloserError]()
	a15 := random.SingleOf[[4]*chan models.CloserReaderSuccess]()
	a16 := random.SingleOf[[4]*chan models.CloserReaderError]()
	a17 := random.SingleOf[[4]*chan models.Product]()
	a18 := random.SingleOf[[4]*chan models.Credential]()
	a19 := random.SingleOf[[4]*chan bool]()
	a20 := random.SingleOf[[4]*chan int]()
	a21 := random.SingleOf[[4]*chan int8]()
	a22 := random.SingleOf[[4]*chan int16]()
	a23 := random.SingleOf[[4]*chan int32]()
	a24 := random.SingleOf[[4]*chan int64]()
	a25 := random.SingleOf[[4]*chan uint]()
	a26 := random.SingleOf[[4]*chan uint8]()
	a27 := random.SingleOf[[4]*chan uint16]()
	a28 := random.SingleOf[[4]*chan uint32]()
	a29 := random.SingleOf[[4]*chan uint64]()
	a30 := random.SingleOf[[4]*chan uintptr]()
	a31 := random.SingleOf[[4]*chan float32]()
	a32 := random.SingleOf[[4]*chan float64]()
	a33 := random.SingleOf[[4]*chan complex64]()
	a34 := random.SingleOf[[4]*chan complex128]()
	a35 := random.SingleOf[[4]*chan string]()

	return []any{
		&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8, &a9, &a10,
		&a11, &a12, &a13, &a14, &a15, &a16, &a17, &a18,
		&a19, &a20, &a21, &a22, &a23, &a24, &a25, &a26,
		&a27, &a28, &a29, &a30, &a31, &a32, &a33, &a34, &a35,
	}
}

func EdgeTypesRefs() []interface{} {
	return []interface{}{
		&[4]unsafe.Pointer{}, &[4]*unsafe.Pointer{}, &[4]**unsafe.Pointer{}, &[4][][]*chan **unsafe.Pointer{},
	}
}
