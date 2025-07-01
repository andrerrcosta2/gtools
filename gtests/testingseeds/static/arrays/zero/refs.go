// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
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
	var bv [1]bool
	var iv [1]int
	var i8v [1]int8
	var i16v [1]int16
	var i32v [1]int32
	var i64v [1]int64
	var uiv [1]uint
	var u8v [1]uint8
	var u16v [1]uint16
	var u32v [1]uint32
	var u64v [1]uint64
	var uptrv [1]uintptr
	var f32v [1]float32
	var f64v [1]float64
	var c64v [1]complex64
	var c128v [1]complex128
	var sv [1]string
	return []interface{}{&bv, &iv, &i8v, &i16v, &i32v, &i64v, &uiv, &u8v, &u16v,
		&u32v, &u64v, &uptrv, &f32v, &f64v, &c64v, &c128v, &sv,
	}
}

func pointerToPrimitiveRefs() []any {
	var bv [1]*bool
	var iv [1]*int
	var i8v [1]*int8
	var i16v [1]*int16
	var i32v [1]*int32
	var i64v [1]*int64
	var uiv [1]*uint
	var u8v [1]*uint8
	var u16v [1]*uint16
	var u32v [1]*uint32
	var u64v [1]*uint64
	var uptrv [1]*uintptr
	var f32v [1]*float32
	var f64v [1]*float64
	var c64v [1]*complex64
	var c128v [1]*complex128
	var sv [1]*string
	return []any{&bv, &iv, &i8v, &i16v, &i32v, &i64v, &uiv, &u8v, &u16v,
		&u32v, &u64v, &uptrv, &f32v, &f64v, &c64v, &c128v, &sv,
	}
}

func interfaceRefs() []any {
	var i1 = [1]interf.Closer{}
	var i2 = [1]interf.CloserReader{}
	var i3 = [1]interf.CloserReaderWriter{}
	var i4 = [1]interf.Error{}
	var i5 = [1]interf.OneMethod{}
	var i6 = [1]interf.Public{}
	var i7 = [1]interf.Seeder[any]{}
	var i8 = [1]interf.Simple{}
	var i9 = [1]interf.Stringer{}
	var i10 = [1]interf.TwoMethods{}
	var i11 = [1]interf.WithoutImplementation{}
	return []any{
		&i1, &i2, &i3, &i4, &i5, &i6, &i7, &i8, &i9, &i10, &i11,
	}
}

func pointerToInterfaceRefs() []interface{} {
	var i1 = [1]*interf.Closer{}
	var i2 = [1]*interf.CloserReader{}
	var i3 = [1]*interf.CloserReaderWriter{}
	var i4 = [1]*interf.Error{}
	var i5 = [1]*interf.OneMethod{}
	var i6 = [1]*interf.Public{}
	var i7 = [1]*interf.Seeder[any]{}
	var i8 = [1]*interf.Simple{}
	var i9 = [1]*interf.Stringer{}
	var i10 = [1]*interf.TwoMethods{}
	var i11 = [1]*interf.WithoutImplementation{}
	return []any{
		&i1, &i2, &i3, &i4, &i5, &i6, &i7, &i8, &i9, &i10, &i11,
	}
}

func structRefs() []any {
	return []any{
		&[1]models.Boolean{}, &[1]models.Integer{}, &[1]models.Uint{}, &[1]models.Float{},
		&[1]models.Complex{}, &[1]models.String{}, &[1]models.Channel[any]{}, &[1]models.Slice[any]{},
		&[1]models.Map[any, any]{}, &[1]models.Function[any]{},
		&[1]models.Simple{}, &[1]models.Public{}, &[1]models.CloserSuccess{}, &[1]models.CloserError{},
		&[1]models.CloserReaderSuccess{}, &[1]models.CloserReaderError{},
		&[1]models.CloserReaderWriterSuccess{}, &[1]models.CloserReaderWriterError{},
		&[1]models.Stringer{}, &[1]models.StringerBytes{}, &[1]models.StringerString{},
		&[1]models.NaturallyComparable{}, &[1]models.NaturallyComparableWithMethods{},
		&[1]models.NotComparable{}, &[1]models.OneData{}, &[1]models.TwoData{},
		&[1]models.User{}, &[1]models.Address{}, &[1]models.Credential{}, &[1]models.Account{},
		&[1]models.Product{}, &[1]models.Profile{},
	}
}

func pointerToStructRefs() []interface{} {
	return []any{
		&[1]*models.Boolean{}, &[1]*models.Integer{}, &[1]*models.Uint{}, &[1]*models.Float{},
		&[1]*models.Complex{}, &[1]*models.String{}, &[1]*models.Channel[any]{}, &[1]*models.Slice[any]{},
		&[1]*models.Map[any, any]{}, &[1]*models.Function[any]{},
		&[1]*models.Simple{}, &[1]*models.Public{}, &[1]*models.CloserSuccess{}, &[1]*models.CloserError{},
		&[1]*models.CloserReaderSuccess{}, &[1]*models.CloserReaderError{},
		&[1]*models.CloserReaderWriterSuccess{}, &[1]*models.CloserReaderWriterError{},
		&[1]*models.Stringer{}, &[1]*models.StringerBytes{}, &[1]*models.StringerString{},
		&[1]*models.NaturallyComparable{}, &[1]*models.NaturallyComparableWithMethods{},
		&[1]*models.NotComparable{}, &[1]*models.OneData{}, &[1]*models.TwoData{},
		&[1]*models.User{}, &[1]*models.Address{}, &[1]*models.Credential{}, &[1]*models.Account{},
		&[1]*models.Product{}, &[1]*models.Profile{},
	}
}

func channelRefs() []interface{} {
	return []any{
		&[1]chan models.Boolean{}, &[1]chan models.Integer{}, &[1]chan models.Uint{}, &[1]chan models.Float{},
		&[1]chan models.Complex{}, &[1]chan models.String{}, &[1]chan models.Channel[any]{}, &[1]chan models.Slice[any]{},
		&[1]chan models.Map[any, any]{}, &[1]chan models.Function[any]{},
		&[1]chan models.Simple{}, &[1]chan models.Public{}, &[1]chan models.CloserSuccess{}, &[1]chan models.CloserError{},
		&[1]chan models.CloserReaderSuccess{}, &[1]chan models.CloserReaderError{},
		&[1]chan models.Product{}, &[1]chan models.Credential{},

		&[1]chan bool{}, &[1]chan int{}, &[1]chan int8{}, &[1]chan int16{}, &[1]chan int32{}, &[1]chan int64{},
		&[1]chan uint{}, &[1]chan uint8{}, &[1]chan uint16{}, &[1]chan uint32{}, &[1]chan uint64{}, &[1]chan uintptr{},
		&[1]chan float32{}, &[1]chan float64{}, &[1]chan complex64{}, &[1]chan complex128{}, &[1]chan string{},
	}
}

func pointerToChannelRefs() []interface{} {
	return []any{
		&[1]*chan models.Boolean{}, &[1]*chan models.Integer{}, &[1]*chan models.Uint{}, &[1]*chan models.Float{},
		&[1]*chan models.Complex{}, &[1]*chan models.String{}, &[1]*chan models.Channel[any]{}, &[1]*chan models.Slice[any]{},
		&[1]*chan models.Map[any, any]{}, &[1]*chan models.Function[any]{},
		&[1]*chan models.Simple{}, &[1]*chan models.Public{}, &[1]*chan models.CloserSuccess{}, &[1]*chan models.CloserError{},
		&[1]*chan models.CloserReaderSuccess{}, &[1]*chan models.CloserReaderError{},
		&[1]*chan models.Product{}, &[1]*chan models.Credential{},

		&[1]*chan bool{}, &[1]*chan int{}, &[1]*chan int8{}, &[1]*chan int16{}, &[1]*chan int32{}, &[1]*chan int64{},
		&[1]*chan uint{}, &[1]*chan uint8{}, &[1]*chan uint16{}, &[1]*chan uint32{}, &[1]*chan uint64{}, &[1]*chan uintptr{},
		&[1]*chan float32{}, &[1]*chan float64{}, &[1]*chan complex64{}, &[1]*chan complex128{}, &[1]*chan string{},
	}
}

func EdgeTypesRefs() []interface{} {
	return []interface{}{
		&[1]unsafe.Pointer{}, &[1]*unsafe.Pointer{}, &[1]**unsafe.Pointer{}, &[1][][]*chan **unsafe.Pointer{},
	}
}
