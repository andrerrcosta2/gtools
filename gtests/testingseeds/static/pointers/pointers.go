// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package pointers

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"math/rand"
	"unsafe"
)

func Invalid() *any {
	return (*interface{})(nil)
}

var ZeroBoolInst = Bool(true)
var RandBoolInst = Bool(rand.Intn(2) == 1)

func Bool(b bool) *bool {
	return &b
}

var ZeroByteInst = Byte(byte(0))
var RandByteInst = Byte(byte(rand.Intn(256)))

func Byte(i byte) *byte {
	return &i
}

func Chan[C ~chan T, T any](i C) *C {
	return &i
}

var ZeroComplex64Inst = Complex64(complex64(0))
var RandComplex64Inst = Complex64(random.Complex64(1, 256, 256, 256, 256).At(0))

func Complex64(i complex64) *complex64 {
	return &i
}

var ZeroComplex128Inst = Complex128(complex128(0))
var RandComplex128Inst = Complex128(random.Complex128(1, 256, 256, 256, 256).At(0))

func Complex128(i complex128) *complex128 {
	return &i
}

var ZeroErrorInst = Error(errors.New(""))
var RandErrorInst = Error(errors.New(random.String(1, 256).At(0)))

func Error(i error) *error {
	return &i
}

var ZeroFloat32Inst = Float32(float32(0))
var RandFloat32Inst = Float32(rand.Float32())

func Float32(i float32) *float32 {
	return &i
}

var ZeroFloat64Inst = Float64(float64(0))
var RandFloat64Inst = Float64(rand.Float64())

func Float64(i float64) *float64 {
	return &i
}

var ZeroIntInst = Int(0)
var RandIntInst = Int(rand.Int())

func Int(i int) *int {
	return &i
}

var ZeroInt8Inst = Int8(int8(0))
var RandInt8Inst = Int8(random.Int8(1).At(0))

func Int8(i int8) *int8 {
	return &i
}

var ZeroInt16Inst = Int16(int16(0))
var RandInt16Inst = Int16(random.Int16(1).At(0))

func Int16(i int16) *int16 {
	return &i
}

var ZeroInt32Inst = Int32(int32(0))
var RandInt32Inst = Int32(random.Int32(1).At(0))

func Int32(i int32) *int32 {
	return &i
}

var ZeroInt64Inst = Int64(int64(0))
var RandInt64Inst = Int64(random.Int64(1).At(0))

func Int64(i int64) *int64 {
	return &i
}

func Map[K ~map[string]T, T any](i K) *K {
	return &i
}

func Nil[T any]() *T {
	return (*T)(nil)
}

var ZeroRuneInst = Rune(rune(0))
var RandRuneInst = Rune(random.Rune(1).At(0))

func Rune(i rune) *rune {
	return &i
}

func Slice[S ~[]T, T any](i S) *S {
	return &i
}

var ZeroStringInst = String("")
var RandStringInst = String(random.String(1, 256).At(0))

func String(i string) *string {
	return &i
}

var ZeroUintInst = Uint(0)
var RandUintInst = Uint(random.Uint(1).At(0))

func Uint(i uint) *uint {
	return &i
}

var ZeroUint8Inst = Uint8(uint8(0))
var RandUint8Inst = Uint8(random.Uint8(1).At(0))

func Uint8(i uint8) *uint8 {
	return &i
}

var ZeroUint16Inst = Uint16(uint16(0))
var RandUint16Inst = Uint16(random.Uint16(1).At(0))

func Uint16(i uint16) *uint16 {
	return &i
}

var ZeroUint32Inst = Uint32(uint32(0))
var RandUint32Inst = Uint32(random.Uint32(1).At(0))

func Uint32(i uint32) *uint32 {
	return &i
}

var ZeroUint64Inst = Uint64(uint64(0))
var RandUint64Inst = Uint64(random.Uint64(1).At(0))

func Uint64(i uint64) *uint64 {
	return &i
}

var ZeroUintptrInst = Uintptr(uintptr(0))
var RandUintptrInst = Uintptr(uintptr(random.Uint32(1).At(0)))

func Uintptr(i uintptr) *uintptr {
	return &i
}

var UnsafeOfZeroSimpleInst = Unsafe(structs.ZeroSimpleInst)
var UnsafeOfZeroPublicInst = Unsafe(structs.ZeroPublicInst)
var UnsafeOfZeroOneDataInst = Unsafe(structs.ZeroOneDataInst)
var UnsafeOfZeroTwoDataInst = Unsafe(structs.ZeroTwoDataInst)
var UnsafeOfZeroStringerInst = Unsafe(structs.ZeroStringerInst)
var UnsafeOfZeroStringerStringInst = Unsafe(structs.ZeroStringerStringInst)
var UnsafeOfZeroStringerBytesInst = Unsafe(structs.ZeroStringerBytesInst)
var UnsafeOfZeroCloserSuccessInst = Unsafe(structs.ZeroCloserSuccessInst)
var UnsafeOfZeroCloserErrorInst = Unsafe(structs.ZeroCloserErrorInst)
var UnsafeOfZeroCloserReaderSuccessInst = Unsafe(structs.ZeroCloserReaderSuccessInst)
var UnsafeOfZeroZeroCloserReaderErrorInst = Unsafe(structs.ZeroCloserReaderErrorInst)
var UnsafeOfZeroZeroCloserReaderWriterSuccessInst = Unsafe(structs.ZeroCloserReaderWriterSuccessInst)
var UnsafeOfZeroCloserReaderWriterErrorInst = Unsafe(structs.ZeroCloserReaderWriterErrorInst)
var UnsafeOfZeroNaturallyComparableInst = Unsafe(structs.ZeroNaturallyComparableInst)
var UnsafeOfZeroNaturallyComparableWithMethodsInst = Unsafe(structs.ZeroNaturallyComparableWithMethodsInst)
var UnsafeOfZeroNotComparableInst = Unsafe(structs.ZeroNotComparableInst)
var UnsafeOfZeroNotComparableWithMethodsInst = Unsafe(structs.ZeroNotComparableWithMethodsInst)

func Unsafe[T any](i *T) unsafe.Pointer {
	return unsafe.Pointer(i)
}

func Value[T any](i T) *T {
	return &i
}
