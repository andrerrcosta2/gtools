// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
)

func valuesSet() []interface{} {
	return slices.Concat(pointersToChannelValues(), channelValues(), structValues(),
		pointersToStructValues(), pointersToInterfaceValues(), interfaceValues(),
		primitiveValues(), pointerToPrimitiveValues())
}

func primitiveValues() []interface{} {
	return []interface{}{
		[1]bool{}, [1]int{}, [1]int8{}, [1]int16{}, [1]int32{}, [1]int64{}, [1]uint{}, [1]uint8{},
		[1]uint16{}, [1]uint32{}, [1]uint64{}, [1]uintptr{}, [1]float32{}, [1]float64{},
		[1]complex64{}, [1]complex128{}, [1]string{},
	}
}

func pointerToPrimitiveValues() []interface{} {
	return []interface{}{
		[1]*bool{}, [1]*int{}, [1]*int8{}, [1]*int16{}, [1]*int32{}, [1]*int64{}, [1]*uint{}, [1]*uint8{},
		[1]*uint16{}, [1]*uint32{}, [1]*uint64{}, [1]*uintptr{}, [1]*float32{}, [1]*float64{},
		[1]*complex64{}, [1]*complex128{}, [1]*string{},
	}
}

func interfaceValues() []interface{} {
	return []any{
		[1]interf.Closer{}, [1]interf.CloserReader{}, [1]interf.CloserReaderWriter{},
		[1]interf.Error{}, [1]interf.OneMethod{}, [1]interf.Public{}, [1]interf.Seeder[any]{},
		[1]interf.Simple{}, [1]interf.Stringer{}, [1]interf.TwoMethods{},
		[1]interf.WithoutImplementation{},
	}
}

func pointersToInterfaceValues() []any {
	return []any{
		[1]*interf.Closer{}, [1]*interf.CloserReader{}, [1]*interf.CloserReaderWriter{},
		[1]*interf.Error{}, [1]*interf.OneMethod{}, [1]*interf.Public{}, [1]*interf.Seeder[any]{},
		[1]*interf.Simple{}, [1]*interf.Stringer{}, [1]*interf.TwoMethods{},
		[1]*interf.WithoutImplementation{},
	}
}

func structValues() []any {
	return []any{
		[1]models.Boolean{}, [1]models.Integer{}, [1]models.Uint{}, [1]models.Float{},
		[1]models.Complex{}, [1]models.String{}, [1]models.Channel[any]{}, [1]models.Slice[any]{},
		[1]models.Map[any, any]{}, [1]models.Function[any]{},
		[1]models.Simple{}, [1]models.Public{}, [1]models.CloserSuccess{}, [1]models.CloserError{},
		[1]models.CloserReaderSuccess{}, [1]models.CloserReaderError{},
		[1]models.CloserReaderWriterSuccess{}, [1]models.CloserReaderWriterError{},
		[1]models.Stringer{}, [1]models.StringerBytes{}, [1]models.StringerString{},
		[1]models.NaturallyComparable{}, [1]models.NaturallyComparableWithMethods{},
		[1]models.NotComparable{}, [1]models.OneData{}, [1]models.TwoData{},
		[1]models.User{}, [1]models.Address{}, [1]models.Credential{}, [1]models.Account{},
		[1]models.Product{}, [1]models.Profile{},
	}
}

func pointersToStructValues() []interface{} {
	return []interface{}{
		[1]*models.Boolean{}, [1]*models.Integer{}, [1]*models.Uint{}, [1]*models.Float{},
		[1]*models.Complex{}, [1]*models.String{}, [1]*models.Channel[any]{}, [1]*models.Slice[any]{},
		[1]*models.Map[any, any]{}, [1]*models.Function[any]{},
		[1]*models.Simple{}, [1]*models.Public{}, [1]*models.CloserSuccess{}, [1]*models.CloserError{},
		[1]*models.CloserReaderSuccess{}, [1]*models.CloserReaderError{},
		[1]*models.CloserReaderWriterSuccess{}, [1]*models.CloserReaderWriterError{},
		[1]*models.Stringer{}, [1]*models.StringerBytes{}, [1]*models.StringerString{},
		[1]*models.NaturallyComparable{}, [1]*models.NaturallyComparableWithMethods{},
		[1]*models.NotComparable{}, [1]*models.OneData{}, [1]*models.TwoData{},
		[1]*models.User{}, [1]*models.Address{}, [1]*models.Credential{}, [1]*models.Account{},
		[1]*models.Product{}, [1]*models.Profile{},
	}
}

func channelValues() []any {
	return []any{
		[1]chan models.Boolean{}, [1]chan models.Integer{}, [1]chan models.Uint{}, [1]chan models.Float{},
		[1]chan models.Complex{}, [1]chan models.String{}, [1]chan models.Channel[any]{}, [1]chan models.Slice[any]{},
		[1]chan models.Map[any, any]{}, [1]chan models.Function[any]{},
		[1]chan models.Simple{}, [1]chan models.Public{}, [1]chan models.CloserSuccess{}, [1]chan models.CloserError{},
		[1]chan models.CloserReaderSuccess{}, [1]chan models.CloserReaderError{},
		[1]chan models.Product{}, [1]chan models.Credential{},

		[1]chan bool{}, [1]chan int{}, [1]chan int8{}, [1]chan int16{}, [1]chan int32{}, [1]chan int64{},
		[1]chan uint{}, [1]chan uint8{}, [1]chan uint16{}, [1]chan uint32{}, [1]chan uint64{}, [1]chan uintptr{},
		[1]chan float32{}, [1]chan float64{}, [1]chan complex64{}, [1]chan complex128{}, [1]chan string{},
	}
}

func pointersToChannelValues() []any {
	return []any{
		[1]*chan models.Boolean{}, [1]*chan models.Integer{}, [1]*chan models.Uint{}, [1]*chan models.Float{},
		[1]*chan models.Complex{}, [1]*chan models.String{}, [1]*chan models.Channel[any]{}, [1]*chan models.Slice[any]{},
		[1]*chan models.Map[any, any]{}, [1]*chan models.Function[any]{},
		[1]*chan models.Simple{}, [1]*chan models.Public{}, [1]*chan models.CloserSuccess{}, [1]*chan models.CloserError{},
		[1]*chan models.CloserReaderSuccess{}, [1]*chan models.CloserReaderError{},
		[1]*chan models.Product{}, [1]*chan models.Credential{},

		[1]*chan bool{}, [1]*chan int{}, [1]*chan int8{}, [1]*chan int16{}, [1]*chan int32{}, [1]*chan int64{},
		[1]*chan uint{}, [1]*chan uint8{}, [1]*chan uint16{}, [1]*chan uint32{}, [1]*chan uint64{}, [1]*chan uintptr{},
		[1]*chan float32{}, [1]*chan float64{}, [1]*chan complex64{}, [1]*chan complex128{}, [1]*chan string{},
	}
}
