// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"slices"
)

func primitivesZero() []interface{} {
	return []interface{}{
		[1]bool{false}, [1]int{0}, [1]int8{int8(0)}, [1]int16{int16(0)}, [1]int32{int32(0)},
		[1]int64{int64(0)}, [1]uint{uint(0)}, [1]uint8{uint8(0)}, [1]uint16{uint16(0)}, [1]uint32{uint32(0)},
		[1]uint64{uint64(0)}, [1]uintptr{uintptr(0)}, [1]float32{float32(0)}, [1]float64{float64(0)},
		[1]complex64{complex64(0)}, [1]complex128{complex128(0)}, [1]string{""},
	}
}

func primitivesEmpty() []interface{} {
	return []interface{}{
		[1]bool{}, [1]int{}, [1]int8{}, [1]int16{}, [1]int32{}, [1]int64{}, [1]uint{}, [1]uint8{},
		[1]uint16{}, [1]uint32{}, [1]uint64{}, [1]uintptr{}, [1]float32{}, [1]float64{},
		[1]complex64{}, [1]complex128{}, [1]string{},
	}
}

func pointerToPrimitive() []interface{} {
	return []interface{}{
		[1]*bool{}, [1]*int{}, [1]*int8{}, [1]*int16{}, [1]*int32{}, [1]*int64{}, [1]*uint{}, [1]*uint8{},
		[1]*uint16{}, [1]*uint32{}, [1]*uint64{}, [1]*uintptr{}, [1]*float32{}, [1]*float64{},
		[1]*complex64{}, [1]*complex128{}, [1]*string{},
	}
}

func ofInterfaces() []interface{} {
	return []any{
		[1]interf.Closer{}, [1]interf.CloserReader{}, [1]interf.CloserReaderWriter{},
		[1]interf.Error{}, [1]interf.OneMethod{}, [1]interf.Public{}, [1]interf.Seeder[any]{},
		[1]interf.Simple{}, [1]interf.Stringer{}, [1]interf.TwoMethods{},
		[1]interf.WithoutImplementation{},
	}
}

func ofPointersToInterfaces() []any {
	return []any{
		[1]*interf.Closer{}, [1]*interf.CloserReader{}, [1]*interf.CloserReaderWriter{},
		[1]*interf.Error{}, [1]*interf.OneMethod{}, [1]*interf.Public{}, [1]*interf.Seeder[any]{},
		[1]*interf.Simple{}, [1]*interf.Stringer{}, [1]*interf.TwoMethods{},
		[1]*interf.WithoutImplementation{},
	}
}

func ofStructs() []any {
	return []any{
		[1]structs.Boolean{}, [1]structs.Integer{}, [1]structs.Uint{}, [1]structs.Float{},
		[1]structs.Complex{}, [1]structs.String{}, [1]structs.Channel[any]{}, [1]structs.Slice[any]{},
		[1]structs.Map[any, any]{}, [1]structs.Function[any]{},
		[1]structs.Simple{}, [1]structs.Public{}, [1]structs.CloserSuccess{}, [1]structs.CloserError{},
		[1]structs.CloserReaderSuccess{}, [1]structs.CloserReaderError{},
		[1]structs.CloserReaderWriterSuccess{}, [1]structs.CloserReaderWriterError{},
		[1]structs.Stringer{}, [1]structs.StringerBytes{}, [1]structs.StringerString{},
		[1]structs.NaturallyComparable{}, [1]structs.NaturallyComparableWithMethods{},
		[1]structs.NotComparable{}, [1]structs.OneData{}, [1]structs.TwoData{},
		[1]static.User{}, [1]static.Address{}, [1]static.Credential{}, [1]static.Account{},
		[1]static.Product{}, [1]static.Profile{},
	}
}

func ofPointersToStructs() []interface{} {
	return []interface{}{
		[1]*structs.Boolean{}, [1]*structs.Integer{}, [1]*structs.Uint{}, [1]*structs.Float{},
		[1]*structs.Complex{}, [1]*structs.String{}, [1]*structs.Channel[any]{}, [1]*structs.Slice[any]{},
		[1]*structs.Map[any, any]{}, [1]*structs.Function[any]{},
		[1]*structs.Simple{}, [1]*structs.Public{}, [1]*structs.CloserSuccess{}, [1]*structs.CloserError{},
		[1]*structs.CloserReaderSuccess{}, [1]*structs.CloserReaderError{},
		[1]*structs.CloserReaderWriterSuccess{}, [1]*structs.CloserReaderWriterError{},
		[1]*structs.Stringer{}, [1]*structs.StringerBytes{}, [1]*structs.StringerString{},
		[1]*structs.NaturallyComparable{}, [1]*structs.NaturallyComparableWithMethods{},
		[1]*structs.NotComparable{}, [1]*structs.OneData{}, [1]*structs.TwoData{},
		[1]*static.User{}, [1]*static.Address{}, [1]*static.Credential{}, [1]*static.Account{},
		[1]*static.Product{}, [1]*static.Profile{},
	}
}

func ofChannels() []any {
	return []any{
		[1]chan structs.Boolean{}, [1]chan structs.Integer{}, [1]chan structs.Uint{}, [1]chan structs.Float{},
		[1]chan structs.Complex{}, [1]chan structs.String{}, [1]chan structs.Channel[any]{}, [1]chan structs.Slice[any]{},
		[1]chan structs.Map[any, any]{}, [1]chan structs.Function[any]{},
		[1]chan structs.Simple{}, [1]chan structs.Public{}, [1]chan structs.CloserSuccess{}, [1]chan structs.CloserError{},
		[1]chan structs.CloserReaderSuccess{}, [1]chan structs.CloserReaderError{},
		[1]chan static.Product{}, [1]chan static.Credential{},

		[1]chan bool{}, [1]chan int{}, [1]chan int8{}, [1]chan int16{}, [1]chan int32{}, [1]chan int64{},
		[1]chan uint{}, [1]chan uint8{}, [1]chan uint16{}, [1]chan uint32{}, [1]chan uint64{}, [1]chan uintptr{},
		[1]chan float32{}, [1]chan float64{}, [1]chan complex64{}, [1]chan complex128{}, [1]chan string{},
	}
}

func ofPointersToChannels() []any {
	return []any{
		[1]*chan structs.Boolean{}, [1]*chan structs.Integer{}, [1]*chan structs.Uint{}, [1]*chan structs.Float{},
		[1]*chan structs.Complex{}, [1]*chan structs.String{}, [1]*chan structs.Channel[any]{}, [1]*chan structs.Slice[any]{},
		[1]*chan structs.Map[any, any]{}, [1]*chan structs.Function[any]{},
		[1]*chan structs.Simple{}, [1]*chan structs.Public{}, [1]*chan structs.CloserSuccess{}, [1]*chan structs.CloserError{},
		[1]*chan structs.CloserReaderSuccess{}, [1]*chan structs.CloserReaderError{},
		[1]*chan static.Product{}, [1]*chan static.Credential{},

		[1]*chan bool{}, [1]*chan int{}, [1]*chan int8{}, [1]*chan int16{}, [1]*chan int32{}, [1]*chan int64{},
		[1]*chan uint{}, [1]*chan uint8{}, [1]*chan uint16{}, [1]*chan uint32{}, [1]*chan uint64{}, [1]*chan uintptr{},
		[1]*chan float32{}, [1]*chan float64{}, [1]*chan complex64{}, [1]*chan complex128{}, [1]*chan string{},
	}
}

func ofValues() []interface{} {
	return slices.Concat(ofPointersToChannels(), ofChannels(), ofStructs(),
		ofPointersToStructs(), ofPointersToInterfaces(), ofInterfaces(),
		primitivesEmpty(), pointerToPrimitive(), primitivesZero())
}
