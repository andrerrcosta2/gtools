// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
	"unsafe"
)

func valuesSet() []interface{} {
	return slices.Concat(pointerToChannelValues(), channelValues(), structValues(),
		pointerToStructValues(), pointerToInterfaceValues(), interfaceValues(),
		pointerToPrimitiveValues(), primitiveValues(),
		edgeValues(),
	)
}

func primitiveValues() []interface{} {
	return []interface{}{
		random.Of[bool](10).Values(), random.Of[int](10).Values(), random.Of[int8](10).Values(),
		random.Of[int16](10).Values(), random.Of[int32](10).Values(), random.Of[int64](10).Values(),
		random.Of[uint](10).Values(), random.Of[uint8](10).Values(),
		random.Of[uint16](10).Values(), random.Of[uint32](10).Values(), random.Of[uint64](10).Values(),
		random.Of[uintptr](10).Values(), random.Of[float32](10).Values(), random.Of[float64](10).Values(),
		random.Of[complex64](10).Values(), random.Of[complex128](10).Values(), random.Of[string](10).Values(),
	}
}

func pointerToPrimitiveValues() []interface{} {
	var bv = random.Of[*bool](10).Values()
	var iv = random.Of[*int](10).Values()
	var i8v = random.Of[*int8](10).Values()
	var i16v = random.Of[*int16](10).Values()
	var i32v = random.Of[*int32](10).Values()
	var i64v = random.Of[*int64](10).Values()
	var uiv = random.Of[*uint](10).Values()
	var u8v = random.Of[*uint8](10).Values()
	var u16v = random.Of[*uint16](10).Values()
	var u32v = random.Of[*uint32](10).Values()
	var u64v = random.Of[*uint64](10).Values()
	var uptrv = random.Of[*uintptr](10).Values()
	var f32v = random.Of[*float32](10).Values()
	var f64v = random.Of[*float64](10).Values()
	var c64v = random.Of[*complex64](10).Values()
	var c128v = random.Of[*complex128](10).Values()
	var sv = random.Of[*string](10).Values()
	return []interface{}{bv, iv, i8v, i16v, i32v, i64v, uiv, u8v, u16v, u32v, u64v, uptrv, f32v, f64v, c64v, c128v, sv}
}

func interfaceValues() []interface{} {
	return []interface{}{
		random.Of[interf.Stringer](10).Values(), random.Of[interf.Closer](10).Values(), random.Of[interf.CloserReader](10).Values(), random.Of[interf.CloserReaderWriter](10).Values(),
		random.Of[interf.WithoutImplementation](10).Values(), random.Of[interf.Error](10).Values(), random.Of[interf.Simple](10).Values(), random.Of[interf.OneMethod](10).Values(),
		random.Of[interf.TwoMethods](10).Values(), random.Of[interf.Public](10).Values(), random.Of[interf.Seeder[any]](10).Values(), random.Of[interf.Simple](10).Values(),
		random.Of[interf.OneMethod](10).Values(), random.Of[interf.TwoMethods](10).Values(),
	}
}

func pointerToInterfaceValues() []interface{} {
	return []interface{}{
		random.Of[*interf.Stringer](10).Values(), random.Of[*interf.Closer](10).Values(), random.Of[*interf.CloserReader](10).Values(), random.Of[*interf.CloserReaderWriter](10).Values(),
		random.Of[*interf.WithoutImplementation](10).Values(), random.Of[*interf.Error](10).Values(), random.Of[*interf.Simple](10).Values(), random.Of[*interf.OneMethod](10).Values(),
		random.Of[*interf.TwoMethods](10).Values(), random.Of[*interf.Public](10).Values(), random.Of[*interf.Seeder[any]](10).Values(), random.Of[*interf.Simple](10).Values(),
		random.Of[*interf.OneMethod](10).Values(), random.Of[*interf.TwoMethods](10).Values(),
	}
}

func structValues() []interface{} {
	return []interface{}{
		random.Of[models.Boolean](10).Values(), random.Of[models.Integer](10).Values(), random.Of[models.String](10).Values(), random.Of[models.Float](10).Values(),
		random.Of[models.Complex](10).Values(), random.Of[models.Map[any, any]](10).Values(), random.Of[models.Slice[any]](10).Values(),
		random.Of[models.Simple](10).Values(), random.Of[models.Public](10).Values(), random.Of[models.Channel[any]](10).Values(),
		random.Of[models.CloserSuccess](10).Values(), random.Of[models.CloserError](10).Values(), random.Of[models.CloserReaderSuccess](10).Values(),
		random.Of[models.CloserReaderError](10).Values(), random.Of[models.CloserReaderWriterSuccess](10).Values(),
		random.Of[models.CloserReaderWriterError](10).Values(), random.Of[models.NotComparable](10).Values(),
		random.Of[models.NaturallyComparable](10).Values(), random.Of[models.NaturallyComparableWithMethods](10).Values(),
		random.Of[models.Stringer](10).Values(), random.Of[models.StringerBytes](10).Values(), random.Of[models.StringerString](10).Values(),
		random.Of[models.OneData](10).Values(), random.Of[models.TwoData](10).Values(),
		random.Of[models.Profile](10).Values(), random.Of[models.Credential](10).Values(), random.Of[models.Product](10).Values(), random.Of[models.Address](10).Values(),
		random.Of[models.User](10).Values(), random.Of[models.Account](10).Values(),
	}
}

func pointerToStructValues() []interface{} {
	return []any{
		random.Of[*models.Boolean](10).Values(), random.Of[*models.Integer](10).Values(), random.Of[*models.String](10).Values(), random.Of[*models.Float](10).Values(),
		random.Of[*models.Complex](10).Values(), random.Of[*models.Map[any, any]](10).Values(), random.Of[*models.Slice[any]](10).Values(),
		random.Of[*models.Simple](10).Values(), random.Of[*models.Public](10).Values(), random.Of[*models.Channel[any]](10).Values(),
		random.Of[*models.CloserSuccess](10).Values(), random.Of[*models.CloserError](10).Values(), random.Of[*models.CloserReaderSuccess](10).Values(),
		random.Of[*models.CloserReaderError](10).Values(), random.Of[*models.CloserReaderWriterSuccess](10).Values(),
		random.Of[*models.CloserReaderWriterError](10).Values(), random.Of[*models.NotComparable](10).Values(),
		random.Of[*models.NaturallyComparable](10).Values(), random.Of[*models.NaturallyComparableWithMethods](10).Values(),
		random.Of[*models.Stringer](10).Values(), random.Of[*models.StringerBytes](10).Values(), random.Of[*models.StringerString](10).Values(),
		random.Of[*models.OneData](10).Values(), random.Of[*models.TwoData](10).Values(),
		random.Of[*models.Profile](10).Values(), random.Of[*models.Credential](10).Values(), random.Of[*models.Product](10).Values(), random.Of[*models.Address](10).Values(),
		random.Of[*models.User](10).Values(), random.Of[*models.Account](10).Values(),
	}
}

func channelValues() []interface{} {
	return []interface{}{
		random.Of[chan models.Public](10).Values(), random.Of[chan models.String](10).Values(), random.Of[chan models.Simple](10).Values(), random.Of[chan models.CloserSuccess](10).Values(),
		random.Of[chan models.CloserError](10).Values(), random.Of[chan models.CloserReaderSuccess](10).Values(), random.Of[chan models.CloserReaderError](10).Values(),
		random.Of[chan models.CloserReaderWriterSuccess](10).Values(), random.Of[chan models.CloserReaderWriterError](10).Values(),
		random.Of[chan models.Product](10).Values(), random.Of[chan models.Credential](10).Values(),
		random.Of[chan bool](10).Values(), random.Of[chan int](10).Values(), random.Of[chan int8](10).Values(), random.Of[chan int16](10).Values(), random.Of[chan int32](10).Values(), random.Of[chan int64](10).Values(),
		random.Of[chan uint](10).Values(), random.Of[chan uint8](10).Values(), random.Of[chan uint16](10).Values(), random.Of[chan uint32](10).Values(), random.Of[chan uint64](10).Values(), random.Of[chan uintptr](10).Values(),
		random.Of[chan float32](10).Values(), random.Of[chan float64](10).Values(), random.Of[chan complex64](10).Values(), random.Of[chan complex128](10).Values(), random.Of[chan string](10).Values(),
	}
}

func pointerToChannelValues() []interface{} {
	return []interface{}{
		random.Of[*chan models.Public](10).Values(), random.Of[*chan models.String](10).Values(), random.Of[*chan models.Simple](10).Values(), random.Of[*chan models.CloserSuccess](10).Values(),
		random.Of[*chan models.CloserError](10).Values(), random.Of[*chan models.CloserReaderSuccess](10).Values(), random.Of[*chan models.CloserReaderError](10).Values(),
		random.Of[*chan models.CloserReaderWriterSuccess](10).Values(), random.Of[*chan models.CloserReaderWriterError](10).Values(),
		random.Of[*chan models.Product](10).Values(), random.Of[*chan models.Credential](10).Values(),
		random.Of[*chan bool](10).Values(), random.Of[*chan int](10).Values(), random.Of[*chan int8](10).Values(), random.Of[*chan int16](10).Values(), random.Of[*chan int32](10).Values(), random.Of[*chan int64](10).Values(),
		random.Of[*chan uint](10).Values(), random.Of[*chan uint8](10).Values(), random.Of[*chan uint16](10).Values(), random.Of[*chan uint32](10).Values(), random.Of[*chan uint64](10).Values(), random.Of[*chan uintptr](10).Values(),
		random.Of[*chan float32](10).Values(), random.Of[*chan float64](10).Values(), random.Of[*chan complex64](10).Values(), random.Of[*chan complex128](10).Values(), random.Of[*chan string](10).Values(),
	}
}

func edgeValues() []interface{} {
	return []interface{}{
		random.Of[**interface{}](10).Values(), random.Of[*chan interface{}](10).Values(),
		random.Of[*unsafe.Pointer](10).Values(), random.Of[*chan **unsafe.Pointer](10).Values(),
	}
}
