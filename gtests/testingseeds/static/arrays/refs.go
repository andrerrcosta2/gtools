// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
	"unsafe"
)

func primitivesZeroRefs() []any {
	bv := [1]bool{false}
	iv := [1]int{0}
	i8v := [1]int8{int8(0)}
	i16v := [1]int16{int16(0)}
	i32v := [1]int32{int32(0)}
	i64v := [1]int64{int64(0)}
	uiv := [1]uint{uint(0)}
	u8v := [1]uint8{uint8(0)}
	u16v := [1]uint16{uint16(0)}
	u32v := [1]uint32{uint32(0)}
	u64v := [1]uint64{uint64(0)}
	uptrv := [1]uintptr{uintptr(0)}
	f32v := [1]float32{float32(0)}
	f64v := [1]float64{float64(0)}
	c64v := [1]complex64{complex64(0)}
	c128v := [1]complex128{complex128(0)}
	sv := [1]string{""}
	return []any{&bv, &iv, &i8v, &i16v, &i32v, &i64v, &uiv, &u8v, &u16v,
		&u32v, &u64v, &uptrv, &f32v, &f64v, &c64v, &c128v, &sv,
	}
}

func primitivesEmptyRefs() []any {
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

func interfacesRefs() []any {
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

func pointerToInterfacesRefs() []interface{} {
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

func structsRefs() []any {
	return []any{
		&[1]structs.Boolean{}, &[1]structs.Integer{}, &[1]structs.Uint{}, &[1]structs.Float{},
		&[1]structs.Complex{}, &[1]structs.String{}, &[1]structs.Channel[any]{}, &[1]structs.Slice[any]{},
		&[1]structs.Map[any, any]{}, &[1]structs.Function[any]{},
		&[1]structs.Simple{}, &[1]structs.Public{}, &[1]structs.CloserSuccess{}, &[1]structs.CloserError{},
		&[1]structs.CloserReaderSuccess{}, &[1]structs.CloserReaderError{},
		&[1]structs.CloserReaderWriterSuccess{}, &[1]structs.CloserReaderWriterError{},
		&[1]structs.Stringer{}, &[1]structs.StringerBytes{}, &[1]structs.StringerString{},
		&[1]structs.NaturallyComparable{}, &[1]structs.NaturallyComparableWithMethods{},
		&[1]structs.NotComparable{}, &[1]structs.OneData{}, &[1]structs.TwoData{},
		&[1]static.User{}, &[1]static.Address{}, &[1]static.Credential{}, &[1]static.Account{},
		&[1]static.Product{}, &[1]static.Profile{},
	}
}

func refToStructsRefs() []interface{} {
	return []any{
		&[1]*structs.Boolean{}, &[1]*structs.Integer{}, &[1]*structs.Uint{}, &[1]*structs.Float{},
		&[1]*structs.Complex{}, &[1]*structs.String{}, &[1]*structs.Channel[any]{}, &[1]*structs.Slice[any]{},
		&[1]*structs.Map[any, any]{}, &[1]*structs.Function[any]{},
		&[1]*structs.Simple{}, &[1]*structs.Public{}, &[1]*structs.CloserSuccess{}, &[1]*structs.CloserError{},
		&[1]*structs.CloserReaderSuccess{}, &[1]*structs.CloserReaderError{},
		&[1]*structs.CloserReaderWriterSuccess{}, &[1]*structs.CloserReaderWriterError{},
		&[1]*structs.Stringer{}, &[1]*structs.StringerBytes{}, &[1]*structs.StringerString{},
		&[1]*structs.NaturallyComparable{}, &[1]*structs.NaturallyComparableWithMethods{},
		&[1]*structs.NotComparable{}, &[1]*structs.OneData{}, &[1]*structs.TwoData{},
		&[1]*static.User{}, &[1]*static.Address{}, &[1]*static.Credential{}, &[1]*static.Account{},
		&[1]*static.Product{}, &[1]*static.Profile{},
	}
}

func channelsRefs() []interface{} {
	return []any{
		&[1]chan structs.Boolean{}, &[1]chan structs.Integer{}, &[1]chan structs.Uint{}, &[1]chan structs.Float{},
		&[1]chan structs.Complex{}, &[1]chan structs.String{}, &[1]chan structs.Channel[any]{}, &[1]chan structs.Slice[any]{},
		&[1]chan structs.Map[any, any]{}, &[1]chan structs.Function[any]{},
		&[1]chan structs.Simple{}, &[1]chan structs.Public{}, &[1]chan structs.CloserSuccess{}, &[1]chan structs.CloserError{},
		&[1]chan structs.CloserReaderSuccess{}, &[1]chan structs.CloserReaderError{},
		&[1]chan static.Product{}, &[1]chan static.Credential{},

		&[1]chan bool{}, &[1]chan int{}, &[1]chan int8{}, &[1]chan int16{}, &[1]chan int32{}, &[1]chan int64{},
		&[1]chan uint{}, &[1]chan uint8{}, &[1]chan uint16{}, &[1]chan uint32{}, &[1]chan uint64{}, &[1]chan uintptr{},
		&[1]chan float32{}, &[1]chan float64{}, &[1]chan complex64{}, &[1]chan complex128{}, &[1]chan string{},
	}
}

func refToChannelRefs() []interface{} {
	return []any{
		&[1]*chan structs.Boolean{}, &[1]*chan structs.Integer{}, &[1]*chan structs.Uint{}, &[1]*chan structs.Float{},
		&[1]*chan structs.Complex{}, &[1]*chan structs.String{}, &[1]*chan structs.Channel[any]{}, &[1]*chan structs.Slice[any]{},
		&[1]*chan structs.Map[any, any]{}, &[1]*chan structs.Function[any]{},
		&[1]*chan structs.Simple{}, &[1]*chan structs.Public{}, &[1]*chan structs.CloserSuccess{}, &[1]*chan structs.CloserError{},
		&[1]*chan structs.CloserReaderSuccess{}, &[1]*chan structs.CloserReaderError{},
		&[1]*chan static.Product{}, &[1]*chan static.Credential{},

		&[1]*chan bool{}, &[1]*chan int{}, &[1]*chan int8{}, &[1]*chan int16{}, &[1]*chan int32{}, &[1]*chan int64{},
		&[1]*chan uint{}, &[1]*chan uint8{}, &[1]*chan uint16{}, &[1]*chan uint32{}, &[1]*chan uint64{}, &[1]*chan uintptr{},
		&[1]*chan float32{}, &[1]*chan float64{}, &[1]*chan complex64{}, &[1]*chan complex128{}, &[1]*chan string{},
	}
}

func AsRefsOfEdgeTypes() []interface{} {
	return []interface{}{
		&[1]unsafe.Pointer{}, &[1]*unsafe.Pointer{}, &[1]**unsafe.Pointer{}, &[1][][]*chan **unsafe.Pointer{},
	}
}

func ofRefs() []interface{} {
	return slices.Concat(AsRefsOfEdgeTypes(), refToChannelRefs(),
		channelsRefs(), refToStructsRefs(), pointerToInterfacesRefs(),
		interfacesRefs(), primitivesZeroRefs(), primitivesEmptyRefs(), pointerToPrimitiveRefs())
}
