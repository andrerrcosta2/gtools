// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package nameds

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
)

type NamedBool bool
type NamedInt int
type NamedInt8 int8
type NamedInt16 int16
type NamedInt32 int32
type NamedInt64 int64
type NamedUint uint
type NamedUint8 uint8
type NamedUint16 uint16
type NamedUint32 uint32
type NamedUint64 uint64
type NamedUintptr uintptr
type NamedFloat32 float32
type NamedFloat64 float64
type NamedComplex64 complex64
type NamedComplex128 complex128
type NamedString string
type NamedBoolSlice []bool
type NamedIntSlice []int
type NamedInt8Slice []int8
type NamedInt16Slice []int16
type NamedInt32Slice []int32
type NamedInt64Slice []int64
type NamedUintSlice []uint
type NamedUint8Slice []uint8
type NamedUint16Slice []uint16
type NamedUint32Slice []uint32
type NamedUint64Slice []uint64
type NamedUintptrSlice []uintptr
type NamedFloat32Slice []float32
type NamedFloat64Slice []float64
type NamedComplex64Slice []complex64
type NamedComplex128Slice []complex128
type NamedStringSlice []string
type NamedStruct struct{}
type NamedInterface interface{}

type NamedIntStruct structs.Integer
type NamedFloatStruct structs.Float
type NamedBoolStruct structs.Boolean
type NamedUintStruct structs.Uint
type NamedComplexStruct structs.Complex

type NamedBoolChannel chan bool
type NamedBoolPtrChannel chan *bool
type NamedIntChannel chan int
type NamedIntPtrChannel chan *int
type NamedInt8Channel chan int8
type NamedInt8PtrChannel chan *int8
type NamedInt16Channel chan int16
type NamedInt16PtrChannel chan *int16
type NamedInt32Channel chan int32
type NamedInt32PtrChannel chan *int32
type NamedInt64Channel chan int64
type NamedInt64PtrChannel chan *int64
type NamedUintChannel chan uint
type NamedUintPtrChannel chan *uint
type NamedUint8Channel chan uint8
type NamedUint8PtrChannel chan *uint8
type NamedUint16Channel chan uint16
type NamedUint16PtrChannel chan *uint16
type NamedUint32Channel chan uint32
type NamedUint32PtrChannel chan *uint32
type NamedUint64Channel chan uint64
type NamedUint64PtrChannel chan *uint64
type NamedUintptrChannel chan uintptr
type NamedUintptrPtrChannel chan *uintptr
type NamedFloat32Channel chan float32
type NamedFloat32PtrChannel chan *float32
type NamedFloat64Channel chan float64
type NamedFloat64PtrChannel chan *float64
type NamedComplex64Channel chan complex64
type NamedComplex64PtrChannel chan *complex64
type NamedComplex128Channel chan complex128
type NamedComplex128PtrChannel chan *complex128
type NamedStructChannel chan struct{}
type NamedStructPtrChannel chan *struct{}
type NamedInterfaceChannel chan interface{}
type NamedInterfacePtrChannel chan *interface{}
type NamedStringChannel chan string
type NamedStringPtrChannel chan *string
type NamedBoolSliceChannel chan []bool
type NamedBoolSlicePtrChannel chan []*bool
type NamedIntSliceChannel chan []int
type NamedIntSlicePtrChannel chan []*int
type NamedInt8SliceChannel chan []int8
type NamedInt8SlicePtrChannel chan []*int8
type NamedInt16SliceChannel chan []int16
type NamedInt16SlicePtrChannel chan []*int16
type NamedInt32SliceChannel chan []int32
type NamedInt32SlicePtrChannel chan []*int32
type NamedInt64SliceChannel chan []int64
type NamedInt64SlicePtrChannel chan []*int64
type NamedUintSliceChannel chan []uint
type NamedUintSlicePtrChannel chan []*uint
type NamedUint8SliceChannel chan []uint8
type NamedUint8SlicePtrChannel chan []*uint8
type NamedUint16SliceChannel chan []uint16
type NamedUint16SlicePtrChannel chan []*uint16
type NamedUint32SliceChannel chan []uint32
type NamedUint32SlicePtrChannel chan []*uint32
type NamedUint64SliceChannel chan []uint64
type NamedUint64SlicePtrChannel chan []*uint64
type NamedUintptrSliceChannel chan []uintptr
type NamedUintptrSlicePtrChannel chan []*uintptr
type NamedFloat32SliceChannel chan []float32
type NamedFloat32SlicePtrChannel chan []*float32
type NamedFloat64SliceChannel chan []float64
type NamedFloat64SlicePtrChannel chan []*float64
type NamedComplex64SliceChannel chan []complex64
type NamedComplex64SlicePtrChannel chan []*complex64
type NamedComplex128SliceChannel chan []complex128
type NamedComplex128SlicePtrChannel chan []*complex128
type NamedStringSliceChannel chan []string
type NamedStringSlicePtrChannel chan []*string
type NamedStructSliceChannel chan []struct{}
type NamedStructSlicePtrChannel chan []*struct{}
type NamedInterfaceSliceChannel chan []interface{}
type NamedInterfaceSlicePtrChannel chan []*interface{}
type NamedBoolBoolMapChannel chan map[bool]bool
type NamedBoolIntMapChannel chan map[bool]int
type NamedBoolInt8MapChannel chan map[bool]int8
type NamedBoolInt16MapChannel chan map[bool]int16
type NamedBoolInt32MapChannel chan map[bool]int32
type NamedBoolInt64MapChannel chan map[bool]int64
type NamedBoolUintMapChannel chan map[bool]uint
type NamedBoolUint8MapChannel chan map[bool]uint8
type NamedBoolUint16MapChannel chan map[bool]uint16
type NamedBoolUint32MapChannel chan map[bool]uint32
type NamedBoolUint64MapChannel chan map[bool]uint64
type NamedBoolUintptrMapChannel chan map[bool]uintptr
type NamedBoolFloat32MapChannel chan map[bool]float32
type NamedBoolFloat64MapChannel chan map[bool]float64
type NamedBoolComplex64MapChannel chan map[bool]complex64
type NamedBoolComplex128MapChannel chan map[bool]complex128
type NamedBoolStringMapChannel chan map[bool]string
type NamedBoolStructMapChannel chan map[bool]struct{}
type NamedBoolInterfaceMapChannel chan map[bool]interface{}
type NamedIntBoolMapChannel chan map[int]bool
type NamedIntIntMapChannel chan map[int]int
type NamedIntInt8MapChannel chan map[int]int8
type NamedIntInt16MapChannel chan map[int]int16
type NamedIntInt32MapChannel chan map[int]int32
type NamedIntInt64MapChannel chan map[int]int64
type NamedIntUintMapChannel chan map[int]uint
type NamedIntUint8MapChannel chan map[int]uint8
type NamedIntUint16MapChannel chan map[int]uint16
type NamedIntUint32MapChannel chan map[int]uint32
type NamedIntUint64MapChannel chan map[int]uint64
type NamedIntUintptrMapChannel chan map[int]uintptr
type NamedIntFloat32MapChannel chan map[int]float32
type NamedIntFloat64MapChannel chan map[int]float64
type NamedIntComplex64MapChannel chan map[int]complex64
type NamedIntComplex128MapChannel chan map[int]complex128
type NamedIntStringMapChannel chan map[int]string
type NamedIntStructMapChannel chan map[int]struct{}
type NamedIntInterfaceMapChannel chan map[int]interface{}
type NamedInt8BoolMapChannel chan map[int8]bool
type NamedInt8IntMapChannel chan map[int8]int
type NamedInt8Int8MapChannel chan map[int8]int8
type NamedInt8Int16MapChannel chan map[int8]int16
type NamedInt8Int32MapChannel chan map[int8]int32
type NamedInt8Int64MapChannel chan map[int8]int64
type NamedInt8UintMapChannel chan map[int8]uint
type NamedInt8Uint8MapChannel chan map[int8]uint8
type NamedInt8Uint16MapChannel chan map[int8]uint16
type NamedInt8Uint32MapChannel chan map[int8]uint32
type NamedInt8Uint64MapChannel chan map[int8]uint64
type NamedInt8UintptrMapChannel chan map[int8]uintptr
type NamedInt8Float32MapChannel chan map[int8]float32
type NamedInt8Float64MapChannel chan map[int8]float64
type NamedInt8Complex64MapChannel chan map[int8]complex64
type NamedInt8Complex128MapChannel chan map[int8]complex128
type NamedInt8StringMapChannel chan map[int8]string
type NamedInt8StructMapChannel chan map[int8]struct{}
type NamedInt8InterfaceMapChannel chan map[int8]interface{}
type NamedInt16BoolMapChannel chan map[int16]bool
type NamedInt16IntMapChannel chan map[int16]int
type NamedInt16Int8MapChannel chan map[int16]int8
type NamedInt16Int16MapChannel chan map[int16]int16
type NamedInt16Int32MapChannel chan map[int16]int32
type NamedInt16Int64MapChannel chan map[int16]int64
type NamedInt16UintMapChannel chan map[int16]uint
type NamedInt16Uint8MapChannel chan map[int16]uint8
type NamedInt16Uint16MapChannel chan map[int16]uint16
type NamedInt16Uint32MapChannel chan map[int16]uint32
type NamedInt16Uint64MapChannel chan map[int16]uint64
type NamedInt16UintptrMapChannel chan map[int16]uintptr
type NamedInt16Float32MapChannel chan map[int16]float32
type NamedInt16Float64MapChannel chan map[int16]float64
type NamedInt16Complex64MapChannel chan map[int16]complex64
type NamedInt16Complex128MapChannel chan map[int16]complex128
type NamedInt16StringMapChannel chan map[int16]string
type NamedInt16StructMapChannel chan map[int16]struct{}
type NamedInt16InterfaceMapChannel chan map[int16]interface{}
type NamedInt32BoolMapChannel chan map[int32]bool
type NamedInt32IntMapChannel chan map[int32]int
type NamedInt32Int8MapChannel chan map[int32]int8
type NamedInt32Int16MapChannel chan map[int32]int16
type NamedInt32Int32MapChannel chan map[int32]int32
type NamedInt32Int64MapChannel chan map[int32]int64
type NamedInt32UintMapChannel chan map[int32]uint
type NamedInt32Uint8MapChannel chan map[int32]uint8
type NamedInt32Uint16MapChannel chan map[int32]uint16
type NamedInt32Uint32MapChannel chan map[int32]uint32
type NamedInt32Uint64MapChannel chan map[int32]uint64
type NamedInt32UintptrMapChannel chan map[int32]uintptr
type NamedInt32Float32MapChannel chan map[int32]float32
type NamedInt32Float64MapChannel chan map[int32]float64
type NamedInt32Complex64MapChannel chan map[int32]complex64
type NamedInt32Complex128MapChannel chan map[int32]complex128
type NamedInt32StringMapChannel chan map[int32]string
type NamedInt32StructMapChannel chan map[int32]struct{}
type NamedInt32InterfaceMapChannel chan map[int32]interface{}
type NamedInt64BoolMapChannel chan map[int64]bool
type NamedInt64IntMapChannel chan map[int64]int
type NamedInt64Int8MapChannel chan map[int64]int8
type NamedInt64Int16MapChannel chan map[int64]int16
type NamedInt64Int32MapChannel chan map[int64]int32
type NamedInt64Int64MapChannel chan map[int64]int64
type NamedInt64UintMapChannel chan map[int64]uint
type NamedInt64Uint8MapChannel chan map[int64]uint8
type NamedInt64Uint16MapChannel chan map[int64]uint16
type NamedInt64Uint32MapChannel chan map[int64]uint32
type NamedInt64Uint64MapChannel chan map[int64]uint64
type NamedInt64UintptrMapChannel chan map[int64]uintptr
type NamedInt64Float32MapChannel chan map[int64]float32
type NamedInt64Float64MapChannel chan map[int64]float64
type NamedInt64Complex64MapChannel chan map[int64]complex64
type NamedInt64Complex128MapChannel chan map[int64]complex128
type NamedInt64StringMapChannel chan map[int64]string
type NamedInt64StructMapChannel chan map[int64]struct{}
type NamedInt64InterfaceMapChannel chan map[int64]interface{}
type NamedUintBoolMapChannel chan map[uint]bool
type NamedUintIntMapChannel chan map[uint]int
type NamedUintInt8MapChannel chan map[uint]int8
type NamedUintInt16MapChannel chan map[uint]int16
type NamedUintInt32MapChannel chan map[uint]int32
type NamedUintInt64MapChannel chan map[uint]int64
type NamedUintUintMapChannel chan map[uint]uint
type NamedUintUint8MapChannel chan map[uint]uint8
type NamedUintUint16MapChannel chan map[uint]uint16
type NamedUintUint32MapChannel chan map[uint]uint32
type NamedUintUint64MapChannel chan map[uint]uint64
type NamedUintUintptrMapChannel chan map[uint]uintptr
type NamedUintFloat32MapChannel chan map[uint]float32
type NamedUintFloat64MapChannel chan map[uint]float64
type NamedUintComplex64MapChannel chan map[uint]complex64
type NamedUintComplex128MapChannel chan map[uint]complex128
type NamedUintStringMapChannel chan map[uint]string
type NamedUintStructMapChannel chan map[uint]struct{}
type NamedUintInterfaceMapChannel chan map[uint]interface{}
type NamedUint8BoolMapChannel chan map[uint8]bool
type NamedUint8IntMapChannel chan map[uint8]int
type NamedUint8Int8MapChannel chan map[uint8]int8
type NamedUint8Int16MapChannel chan map[uint8]int16
type NamedUint8Int32MapChannel chan map[uint8]int32
type NamedUint8Int64MapChannel chan map[uint8]int64
type NamedUint8UintMapChannel chan map[uint8]uint
type NamedUint8Uint8MapChannel chan map[uint8]uint8
type NamedUint8Uint16MapChannel chan map[uint8]uint16
type NamedUint8Uint32MapChannel chan map[uint8]uint32
type NamedUint8Uint64MapChannel chan map[uint8]uint64
type NamedUint8UintptrMapChannel chan map[uint8]uintptr
type NamedUint8Float32MapChannel chan map[uint8]float32
type NamedUint8Float64MapChannel chan map[uint8]float64
type NamedUint8Complex64MapChannel chan map[uint8]complex64
type NamedUint8Complex128MapChannel chan map[uint8]complex128
type NamedUint8StringMapChannel chan map[uint8]string
type NamedUint8StructMapChannel chan map[uint8]struct{}
type NamedUint8InterfaceMapChannel chan map[uint8]interface{}
type NamedUint16BoolMapChannel chan map[uint16]bool
type NamedUint16IntMapChannel chan map[uint16]int
type NamedUint16Int8MapChannel chan map[uint16]int8
type NamedUint16Int16MapChannel chan map[uint16]int16
type NamedUint16Int32MapChannel chan map[uint16]int32
type NamedUint16Int64MapChannel chan map[uint16]int64
type NamedUint16UintMapChannel chan map[uint16]uint
type NamedUint16Uint8MapChannel chan map[uint16]uint8
type NamedUint16Uint16MapChannel chan map[uint16]uint16
type NamedUint16Uint32MapChannel chan map[uint16]uint32
type NamedUint16Uint64MapChannel chan map[uint16]uint64
type NamedUint16UintptrMapChannel chan map[uint16]uintptr
type NamedUint16Float32MapChannel chan map[uint16]float32
type NamedUint16Float64MapChannel chan map[uint16]float64
type NamedUint16Complex64MapChannel chan map[uint16]complex64
type NamedUint16Complex128MapChannel chan map[uint16]complex128
type NamedUint16StringMapChannel chan map[uint16]string
type NamedUint16StructMapChannel chan map[uint16]struct{}
type NamedUint16InterfaceMapChannel chan map[uint16]interface{}
type NamedUint32BoolMapChannel chan map[uint32]bool
type NamedUint32IntMapChannel chan map[uint32]int
type NamedUint32Int8MapChannel chan map[uint32]int8
type NamedUint32Int16MapChannel chan map[uint32]int16
type NamedUint32Int32MapChannel chan map[uint32]int32
type NamedUint32Int64MapChannel chan map[uint32]int64
type NamedUint32UintMapChannel chan map[uint32]uint
type NamedUint32Uint8MapChannel chan map[uint32]uint8
type NamedUint32Uint16MapChannel chan map[uint32]uint16
type NamedUint32Uint32MapChannel chan map[uint32]uint32
type NamedUint32Uint64MapChannel chan map[uint32]uint64
type NamedUint32UintptrMapChannel chan map[uint32]uintptr
type NamedUint32Float32MapChannel chan map[uint32]float32
type NamedUint32Float64MapChannel chan map[uint32]float64
type NamedUint32Complex64MapChannel chan map[uint32]complex64
type NamedUint32Complex128MapChannel chan map[uint32]complex128
type NamedUint32StringMapChannel chan map[uint32]string
type NamedUint32StructMapChannel chan map[uint32]struct{}
type NamedUint32InterfaceMapChannel chan map[uint32]interface{}
type NamedUint64BoolMapChannel chan map[uint64]bool
type NamedUint64IntMapChannel chan map[uint64]int
type NamedUint64Int8MapChannel chan map[uint64]int8
type NamedUint64Int16MapChannel chan map[uint64]int16
type NamedUint64Int32MapChannel chan map[uint64]int32
type NamedUint64Int64MapChannel chan map[uint64]int64
type NamedUint64UintMapChannel chan map[uint64]uint
type NamedUint64Uint8MapChannel chan map[uint64]uint8
type NamedUint64Uint16MapChannel chan map[uint64]uint16
type NamedUint64Uint32MapChannel chan map[uint64]uint32
type NamedUint64Uint64MapChannel chan map[uint64]uint64
type NamedUint64UintptrMapChannel chan map[uint64]uintptr
type NamedUint64Float32MapChannel chan map[uint64]float32
type NamedUint64Float64MapChannel chan map[uint64]float64
type NamedUint64Complex64MapChannel chan map[uint64]complex64
type NamedUint64Complex128MapChannel chan map[uint64]complex128
type NamedUint64StringMapChannel chan map[uint64]string
type NamedUint64StructMapChannel chan map[uint64]struct{}
type NamedUint64InterfaceMapChannel chan map[uint64]interface{}
type NamedUintptrBoolMapChannel chan map[uintptr]bool
type NamedUintptrIntMapChannel chan map[uintptr]int
type NamedUintptrInt8MapChannel chan map[uintptr]int8
type NamedUintptrInt16MapChannel chan map[uintptr]int16
type NamedUintptrInt32MapChannel chan map[uintptr]int32
type NamedUintptrInt64MapChannel chan map[uintptr]int64
type NamedUintptrUintMapChannel chan map[uintptr]uint
type NamedUintptrUint8MapChannel chan map[uintptr]uint8
type NamedUintptrUint16MapChannel chan map[uintptr]uint16
type NamedUintptrUint32MapChannel chan map[uintptr]uint32
type NamedUintptrUint64MapChannel chan map[uintptr]uint64
type NamedUintptrUintptrMapChannel chan map[uintptr]uintptr
type NamedUintptrFloat32MapChannel chan map[uintptr]float32
type NamedUintptrFloat64MapChannel chan map[uintptr]float64
type NamedUintptrComplex64MapChannel chan map[uintptr]complex64
type NamedUintptrComplex128MapChannel chan map[uintptr]complex128
type NamedUintptrStringMapChannel chan map[uintptr]string
type NamedUintptrStructMapChannel chan map[uintptr]struct{}
type NamedUintptrInterfaceMapChannel chan map[uintptr]interface{}
type NamedFloat32BoolMapChannel chan map[float32]bool
type NamedFloat32IntMapChannel chan map[float32]int
type NamedFloat32Int8MapChannel chan map[float32]int8
type NamedFloat32Int16MapChannel chan map[float32]int16
type NamedFloat32Int32MapChannel chan map[float32]int32
type NamedFloat32Int64MapChannel chan map[float32]int64
type NamedFloat32UintMapChannel chan map[float32]uint
type NamedFloat32Uint8MapChannel chan map[float32]uint8
type NamedFloat32Uint16MapChannel chan map[float32]uint16
type NamedFloat32Uint32MapChannel chan map[float32]uint32
type NamedFloat32Uint64MapChannel chan map[float32]uint64
type NamedFloat32UintptrMapChannel chan map[float32]uintptr
type NamedFloat32Float32MapChannel chan map[float32]float32
type NamedFloat32Float64MapChannel chan map[float32]float64
type NamedFloat32Complex64MapChannel chan map[float32]complex64
type NamedFloat32Complex128MapChannel chan map[float32]complex128
type NamedFloat32StringMapChannel chan map[float32]string
type NamedFloat32StructMapChannel chan map[float32]struct{}
type NamedFloat32InterfaceMapChannel chan map[float32]interface{}
type NamedFloat64BoolMapChannel chan map[float64]bool
type NamedFloat64IntMapChannel chan map[float64]int
type NamedFloat64Int8MapChannel chan map[float64]int8
type NamedFloat64Int16MapChannel chan map[float64]int16
type NamedFloat64Int32MapChannel chan map[float64]int32
type NamedFloat64Int64MapChannel chan map[float64]int64
type NamedFloat64UintMapChannel chan map[float64]uint
type NamedFloat64Uint8MapChannel chan map[float64]uint8
type NamedFloat64Uint16MapChannel chan map[float64]uint16
type NamedFloat64Uint32MapChannel chan map[float64]uint32
type NamedFloat64Uint64MapChannel chan map[float64]uint64
type NamedFloat64UintptrMapChannel chan map[float64]uintptr
type NamedFloat64Float32MapChannel chan map[float64]float32
type NamedFloat64Float64MapChannel chan map[float64]float64
type NamedFloat64Complex64MapChannel chan map[float64]complex64
type NamedFloat64Complex128MapChannel chan map[float64]complex128
type NamedFloat64StringMapChannel chan map[float64]string
type NamedFloat64StructMapChannel chan map[float64]struct{}
type NamedFloat64InterfaceMapChannel chan map[float64]interface{}
type NamedComplex64BoolMapChannel chan map[complex64]bool
type NamedComplex64IntMapChannel chan map[complex64]int
type NamedComplex64Int8MapChannel chan map[complex64]int8
type NamedComplex64Int16MapChannel chan map[complex64]int16
type NamedComplex64Int32MapChannel chan map[complex64]int32
type NamedComplex64Int64MapChannel chan map[complex64]int64
type NamedComplex64UintMapChannel chan map[complex64]uint
type NamedComplex64Uint8MapChannel chan map[complex64]uint8
type NamedComplex64Uint16MapChannel chan map[complex64]uint16
type NamedComplex64Uint32MapChannel chan map[complex64]uint32
type NamedComplex64Uint64MapChannel chan map[complex64]uint64
type NamedComplex64UintptrMapChannel chan map[complex64]uintptr
type NamedComplex64Float32MapChannel chan map[complex64]float32
type NamedComplex64Float64MapChannel chan map[complex64]float64
type NamedComplex64Complex64MapChannel chan map[complex64]complex64
type NamedComplex64Complex128MapChannel chan map[complex64]complex128
type NamedComplex64StringMapChannel chan map[complex64]string
type NamedComplex64StructMapChannel chan map[complex64]struct{}
type NamedComplex64InterfaceMapChannel chan map[complex64]interface{}
type NamedComplex128BoolMapChannel chan map[complex128]bool
type NamedComplex128IntMapChannel chan map[complex128]int
type NamedComplex128Int8MapChannel chan map[complex128]int8
type NamedComplex128Int16MapChannel chan map[complex128]int16
type NamedComplex128Int32MapChannel chan map[complex128]int32
type NamedComplex128Int64MapChannel chan map[complex128]int64
type NamedComplex128UintMapChannel chan map[complex128]uint
type NamedComplex128Uint8MapChannel chan map[complex128]uint8
type NamedComplex128Uint16MapChannel chan map[complex128]uint16
type NamedComplex128Uint32MapChannel chan map[complex128]uint32
type NamedComplex128Uint64MapChannel chan map[complex128]uint64
type NamedComplex128UintptrMapChannel chan map[complex128]uintptr
type NamedComplex128Float32MapChannel chan map[complex128]float32
type NamedComplex128Float64MapChannel chan map[complex128]float64
type NamedComplex128Complex64MapChannel chan map[complex128]complex64
type NamedComplex128Complex128MapChannel chan map[complex128]complex128
type NamedComplex128StringMapChannel chan map[complex128]string
type NamedComplex128StructMapChannel chan map[complex128]struct{}
type NamedComplex128InterfaceMapChannel chan map[complex128]interface{}
type NamedStringBoolMapChannel chan map[string]bool
type NamedStringIntMapChannel chan map[string]int
type NamedStringInt8MapChannel chan map[string]int8
type NamedStringInt16MapChannel chan map[string]int16
type NamedStringInt32MapChannel chan map[string]int32
type NamedStringInt64MapChannel chan map[string]int64
type NamedStringUintMapChannel chan map[string]uint
type NamedStringUint8MapChannel chan map[string]uint8
type NamedStringUint16MapChannel chan map[string]uint16
type NamedStringUint32MapChannel chan map[string]uint32
type NamedStringUint64MapChannel chan map[string]uint64
type NamedStringUintptrMapChannel chan map[string]uintptr
type NamedStringFloat32MapChannel chan map[string]float32
type NamedStringFloat64MapChannel chan map[string]float64
type NamedStringComplex64MapChannel chan map[string]complex64
type NamedStringComplex128MapChannel chan map[string]complex128
type NamedStringStringMapChannel chan map[string]string
type NamedStringStructMapChannel chan map[string]struct{}
type NamedStringInterfaceMapChannel chan map[string]interface{}
type NamedStructBoolMapChannel chan map[struct{}]bool
type NamedStructIntMapChannel chan map[struct{}]int
type NamedStructInt8MapChannel chan map[struct{}]int8
type NamedStructInt16MapChannel chan map[struct{}]int16
type NamedStructInt32MapChannel chan map[struct{}]int32
type NamedStructInt64MapChannel chan map[struct{}]int64
type NamedStructUintMapChannel chan map[struct{}]uint
type NamedStructUint8MapChannel chan map[struct{}]uint8
type NamedStructUint16MapChannel chan map[struct{}]uint16
type NamedStructUint32MapChannel chan map[struct{}]uint32
type NamedStructUint64MapChannel chan map[struct{}]uint64
type NamedStructUintptrMapChannel chan map[struct{}]uintptr
type NamedStructFloat32MapChannel chan map[struct{}]float32
type NamedStructFloat64MapChannel chan map[struct{}]float64
type NamedStructComplex64MapChannel chan map[struct{}]complex64
type NamedStructComplex128MapChannel chan map[struct{}]complex128
type NamedStructStringMapChannel chan map[struct{}]string
type NamedStructStructMapChannel chan map[struct{}]struct{}
type NamedStructInterfaceMapChannel chan map[struct{}]interface{}
type NamedInterfaceBoolMapChannel chan map[interface{}]bool
type NamedInterfaceIntMapChannel chan map[interface{}]int
type NamedInterfaceInt8MapChannel chan map[interface{}]int8
type NamedInterfaceInt16MapChannel chan map[interface{}]int16
type NamedInterfaceInt32MapChannel chan map[interface{}]int32
type NamedInterfaceInt64MapChannel chan map[interface{}]int64
type NamedInterfaceUintMapChannel chan map[interface{}]uint
type NamedInterfaceUint8MapChannel chan map[interface{}]uint8
type NamedInterfaceUint16MapChannel chan map[interface{}]uint16
type NamedInterfaceUint32MapChannel chan map[interface{}]uint32
type NamedInterfaceUint64MapChannel chan map[interface{}]uint64
type NamedInterfaceUintptrMapChannel chan map[interface{}]uintptr
type NamedInterfaceFloat32MapChannel chan map[interface{}]float32
type NamedInterfaceFloat64MapChannel chan map[interface{}]float64
type NamedInterfaceComplex64MapChannel chan map[interface{}]complex64
type NamedInterfaceComplex128MapChannel chan map[interface{}]complex128
type NamedInterfaceStringMapChannel chan map[interface{}]string
type NamedInterfaceStructMapChannel chan map[interface{}]struct{}
type NamedInterfaceInterfaceMapChannel chan map[interface{}]interface{}

type NamedBoolBoolMap map[bool]bool
type NamedBoolIntMap map[bool]int
type NamedBoolInt8Map map[bool]int8
type NamedBoolInt16Map map[bool]int16
type NamedBoolInt32Map map[bool]int32
type NamedBoolInt64Map map[bool]int64
type NamedBoolUintMap map[bool]uint
type NamedBoolUint8Map map[bool]uint8
type NamedBoolUint16Map map[bool]uint16
type NamedBoolUint32Map map[bool]uint32
type NamedBoolUint64Map map[bool]uint64
type NamedBoolUintptrMap map[bool]uintptr
type NamedBoolFloat32Map map[bool]float32
type NamedBoolFloat64Map map[bool]float64
type NamedBoolComplex64Map map[bool]complex64
type NamedBoolComplex128Map map[bool]complex128
type NamedBoolStringMap map[bool]string
type NamedBoolStructMap map[bool]struct{}
type NamedBoolInterfaceMap map[bool]interface{}
type NamedIntBoolMap map[int]bool
type NamedIntIntMap map[int]int
type NamedIntInt8Map map[int]int8
type NamedIntInt16Map map[int]int16
type NamedIntInt32Map map[int]int32
type NamedIntInt64Map map[int]int64
type NamedIntUintMap map[int]uint
type NamedIntUint8Map map[int]uint8
type NamedIntUint16Map map[int]uint16
type NamedIntUint32Map map[int]uint32
type NamedIntUint64Map map[int]uint64
type NamedIntUintptrMap map[int]uintptr
type NamedIntFloat32Map map[int]float32
type NamedIntFloat64Map map[int]float64
type NamedIntComplex64Map map[int]complex64
type NamedIntComplex128Map map[int]complex128
type NamedIntStringMap map[int]string
type NamedIntStructMap map[int]struct{}
type NamedIntInterfaceMap map[int]interface{}
type NamedInt8BoolMap map[int8]bool
type NamedInt8IntMap map[int8]int
type NamedInt8Int8Map map[int8]int8
type NamedInt8Int16Map map[int8]int16
type NamedInt8Int32Map map[int8]int32
type NamedInt8Int64Map map[int8]int64
type NamedInt8UintMap map[int8]uint
type NamedInt8Uint8Map map[int8]uint8
type NamedInt8Uint16Map map[int8]uint16
type NamedInt8Uint32Map map[int8]uint32
type NamedInt8Uint64Map map[int8]uint64
type NamedInt8UintptrMap map[int8]uintptr
type NamedInt8Float32Map map[int8]float32
type NamedInt8Float64Map map[int8]float64
type NamedInt8Complex64Map map[int8]complex64
type NamedInt8Complex128Map map[int8]complex128
type NamedInt8StringMap map[int8]string
type NamedInt8StructMap map[int8]struct{}
type NamedInt8InterfaceMap map[int8]interface{}
type NamedInt16BoolMap map[int16]bool
type NamedInt16IntMap map[int16]int
type NamedInt16Int8Map map[int16]int8
type NamedInt16Int16Map map[int16]int16
type NamedInt16Int32Map map[int16]int32
type NamedInt16Int64Map map[int16]int64
type NamedInt16UintMap map[int16]uint
type NamedInt16Uint8Map map[int16]uint8
type NamedInt16Uint16Map map[int16]uint16
type NamedInt16Uint32Map map[int16]uint32
type NamedInt16Uint64Map map[int16]uint64
type NamedInt16UintptrMap map[int16]uintptr
type NamedInt16Float32Map map[int16]float32
type NamedInt16Float64Map map[int16]float64
type NamedInt16Complex64Map map[int16]complex64
type NamedInt16Complex128Map map[int16]complex128
type NamedInt16StringMap map[int16]string
type NamedInt16StructMap map[int16]struct{}
type NamedInt16InterfaceMap map[int16]interface{}
type NamedInt32BoolMap map[int32]bool
type NamedInt32IntMap map[int32]int
type NamedInt32Int8Map map[int32]int8
type NamedInt32Int16Map map[int32]int16
type NamedInt32Int32Map map[int32]int32
type NamedInt32Int64Map map[int32]int64
type NamedInt32UintMap map[int32]uint
type NamedInt32Uint8Map map[int32]uint8
type NamedInt32Uint16Map map[int32]uint16
type NamedInt32Uint32Map map[int32]uint32
type NamedInt32Uint64Map map[int32]uint64
type NamedInt32UintptrMap map[int32]uintptr
type NamedInt32Float32Map map[int32]float32
type NamedInt32Float64Map map[int32]float64
type NamedInt32Complex64Map map[int32]complex64
type NamedInt32Complex128Map map[int32]complex128
type NamedInt32StringMap map[int32]string
type NamedInt32StructMap map[int32]struct{}
type NamedInt32InterfaceMap map[int32]interface{}
type NamedInt64BoolMap map[int64]bool
type NamedInt64IntMap map[int64]int
type NamedInt64Int8Map map[int64]int8
type NamedInt64Int16Map map[int64]int16
type NamedInt64Int32Map map[int64]int32
type NamedInt64Int64Map map[int64]int64
type NamedInt64UintMap map[int64]uint
type NamedInt64Uint8Map map[int64]uint8
type NamedInt64Uint16Map map[int64]uint16
type NamedInt64Uint32Map map[int64]uint32
type NamedInt64Uint64Map map[int64]uint64
type NamedInt64UintptrMap map[int64]uintptr
type NamedInt64Float32Map map[int64]float32
type NamedInt64Float64Map map[int64]float64
type NamedInt64Complex64Map map[int64]complex64
type NamedInt64Complex128Map map[int64]complex128
type NamedInt64StringMap map[int64]string
type NamedInt64StructMap map[int64]struct{}
type NamedInt64InterfaceMap map[int64]interface{}
type NamedUintBoolMap map[uint]bool
type NamedUintIntMap map[uint]int
type NamedUintInt8Map map[uint]int8
type NamedUintInt16Map map[uint]int16
type NamedUintInt32Map map[uint]int32
type NamedUintInt64Map map[uint]int64
type NamedUintUintMap map[uint]uint
type NamedUintUint8Map map[uint]uint8
type NamedUintUint16Map map[uint]uint16
type NamedUintUint32Map map[uint]uint32
type NamedUintUint64Map map[uint]uint64
type NamedUintUintptrMap map[uint]uintptr
type NamedUintFloat32Map map[uint]float32
type NamedUintFloat64Map map[uint]float64
type NamedUintComplex64Map map[uint]complex64
type NamedUintComplex128Map map[uint]complex128
type NamedUintStringMap map[uint]string
type NamedUintStructMap map[uint]struct{}
type NamedUintInterfaceMap map[uint]interface{}
type NamedUintptrBoolMap map[uintptr]bool
type NamedUintptrIntMap map[uintptr]int
type NamedUintptrInt8Map map[uintptr]int8
type NamedUintptrInt16Map map[uintptr]int16
type NamedUintptrInt32Map map[uintptr]int32
type NamedUintptrInt64Map map[uintptr]int64
type NamedUintptrUintMap map[uintptr]uint
type NamedUintptrUint8Map map[uintptr]uint8
type NamedUintptrUint16Map map[uintptr]uint16
type NamedUintptrUint32Map map[uintptr]uint32
type NamedUintptrUint64Map map[uintptr]uint64
type NamedUintptrUintptrMap map[uintptr]uintptr
type NamedUintptrFloat32Map map[uintptr]float32
type NamedUintptrFloat64Map map[uintptr]float64
type NamedUintptrComplex64Map map[uintptr]complex64
type NamedUintptrComplex128Map map[uintptr]complex128
type NamedUintptrStringMap map[uintptr]string
type NamedUintptrStructMap map[uintptr]struct{}
type NamedUintptrInterfaceMap map[uintptr]interface{}
type NamedFloat32BoolMap map[float32]bool
type NamedFloat32IntMap map[float32]int
type NamedFloat32Int8Map map[float32]int8
type NamedFloat32Int16Map map[float32]int16
type NamedFloat32Int32Map map[float32]int32
type NamedFloat32Int64Map map[float32]int64
type NamedFloat32UintMap map[float32]uint
type NamedFloat32Uint8Map map[float32]uint8
type NamedFloat32Uint16Map map[float32]uint16
type NamedFloat32Uint32Map map[float32]uint32
type NamedFloat32Uint64Map map[float32]uint64
type NamedFloat32UintptrMap map[float32]uintptr
type NamedFloat32Float32Map map[float32]float32
type NamedFloat32Float64Map map[float32]float64
type NamedFloat32Complex64Map map[float32]complex64
type NamedFloat32Complex128Map map[float32]complex128
type NamedFloat32StringMap map[float32]string
type NamedFloat32StructMap map[float32]struct{}
type NamedFloat32InterfaceMap map[float32]interface{}
type NamedFloat64BoolMap map[float64]bool
type NamedFloat64IntMap map[float64]int
type NamedFloat64Int8Map map[float64]int8
type NamedFloat64Int16Map map[float64]int16
type NamedFloat64Int32Map map[float64]int32
type NamedFloat64Int64Map map[float64]int64
type NamedFloat64UintMap map[float64]uint
type NamedFloat64Uint8Map map[float64]uint8
type NamedFloat64Uint16Map map[float64]uint16
type NamedFloat64Uint32Map map[float64]uint32
type NamedFloat64Uint64Map map[float64]uint64
type NamedFloat64UintptrMap map[float64]uintptr
type NamedFloat64Float32Map map[float64]float32
type NamedFloat64Float64Map map[float64]float64
type NamedFloat64Complex64Map map[float64]complex64
type NamedFloat64Complex128Map map[float64]complex128
type NamedFloat64StringMap map[float64]string
type NamedFloat64StructMap map[float64]struct{}
type NamedFloat64InterfaceMap map[float64]interface{}
type NamedComplex64BoolMap map[complex64]bool
type NamedComplex64IntMap map[complex64]int
type NamedComplex64Int8Map map[complex64]int8
type NamedComplex64Int16Map map[complex64]int16
type NamedComplex64Int32Map map[complex64]int32
type NamedComplex64Int64Map map[complex64]int64
type NamedComplex64UintMap map[complex64]uint
type NamedComplex64Uint8Map map[complex64]uint8
type NamedComplex64Uint16Map map[complex64]uint16
type NamedComplex64Uint32Map map[complex64]uint32
type NamedComplex64Uint64Map map[complex64]uint64
type NamedComplex64UintptrMap map[complex64]uintptr
type NamedComplex64Float32Map map[complex64]float32
type NamedComplex64Float64Map map[complex64]float64
type NamedComplex64Complex64Map map[complex64]complex64
type NamedComplex64Complex128Map map[complex64]complex128
type NamedComplex64StringMap map[complex64]string
type NamedComplex64StructMap map[complex64]struct{}
type NamedComplex64InterfaceMap map[complex64]interface{}
type NamedStringBoolMap map[string]bool
type NamedStringIntMap map[string]int
type NamedStringInt8Map map[string]int8
type NamedStringInt16Map map[string]int16
type NamedStringInt32Map map[string]int32
type NamedStringInt64Map map[string]int64
type NamedStringUintMap map[string]uint
type NamedStringUint8Map map[string]uint8
type NamedStringUint16Map map[string]uint16
type NamedStringUint32Map map[string]uint32
type NamedStringUint64Map map[string]uint64
type NamedStringUintptrMap map[string]uintptr
type NamedStringFloat32Map map[string]float32
type NamedStringFloat64Map map[string]float64
type NamedStringComplex64Map map[string]complex64
type NamedStringComplex128Map map[string]complex128
type NamedStringStringMap map[string]string
type NamedStringStructMap map[string]struct{}
type NamedStringInterfaceMap map[string]interface{}
type NamedStructBoolMap map[struct{}]bool
type NamedStructIntMap map[struct{}]int
type NamedStructInt8Map map[struct{}]int8
type NamedStructInt16Map map[struct{}]int16
type NamedStructInt32Map map[struct{}]int32
type NamedStructInt64Map map[struct{}]int64
type NamedStructUintMap map[struct{}]uint
type NamedStructUint8Map map[struct{}]uint8
type NamedStructUint16Map map[struct{}]uint16
type NamedStructUint32Map map[struct{}]uint32
type NamedStructUint64Map map[struct{}]uint64
type NamedStructUintptrMap map[struct{}]uintptr
type NamedStructFloat32Map map[struct{}]float32
type NamedStructFloat64Map map[struct{}]float64
type NamedStructComplex64Map map[struct{}]complex64
type NamedStructComplex128Map map[struct{}]complex128
type NamedStructStringMap map[struct{}]string
type NamedStructStructMap map[struct{}]struct{}
type NamedStructInterfaceMap map[struct{}]interface{}
type NamedInterfaceBoolMap map[interface{}]bool
type NamedInterfaceIntMap map[interface{}]int
type NamedInterfaceInt8Map map[interface{}]int8
type NamedInterfaceInt16Map map[interface{}]int16
type NamedInterfaceInt32Map map[interface{}]int32
type NamedInterfaceInt64Map map[interface{}]int64
type NamedInterfaceUintMap map[interface{}]uint
type NamedInterfaceUint8Map map[interface{}]uint8
type NamedInterfaceUint16Map map[interface{}]uint16
type NamedInterfaceUint32Map map[interface{}]uint32
type NamedInterfaceUint64Map map[interface{}]uint64
type NamedInterfaceUintptrMap map[interface{}]uintptr
type NamedInterfaceFloat32Map map[interface{}]float32
type NamedInterfaceFloat64Map map[interface{}]float64
type NamedInterfaceComplex64Map map[interface{}]complex64
type NamedInterfaceComplex128Map map[interface{}]complex128
type NamedInterfaceStringMap map[interface{}]string
type NamedInterfaceStructMap map[interface{}]struct{}
type NamedInterfaceInterfaceMap map[interface{}]interface{}

type NamedPrimitiveWithMethods int

func (p NamedPrimitiveWithMethods) String() string {
	return string(rune(p))
}
