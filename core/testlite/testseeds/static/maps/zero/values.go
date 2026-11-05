// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
	"slices"
	"unsafe"
)

func valuesSet() []interface{} {
	return slices.Concat(edgeValues(), primitiveWithPtrKeyAndValueValues(), primitiveWithPtrValueValues(),
		primitiveWithPtrKeyValues(), primitiveValues())
}

func primitiveValues() []interface{} {
	return []interface{}{
		// bool keys
		map[bool]bool{}, map[bool]int{}, map[bool]int8{}, map[bool]int16{}, map[bool]int32{}, map[bool]int64{},
		map[bool]uint{}, map[bool]uint8{}, map[bool]uint16{}, map[bool]uint32{}, map[bool]uint64{}, map[bool]uintptr{},
		map[bool]float32{}, map[bool]float64{}, map[bool]complex64{}, map[bool]complex128{}, map[bool]string{}, map[bool]any{},
		// int keys
		map[int]bool{}, map[int]int{}, map[int]int8{}, map[int]int16{}, map[int]int32{}, map[int]int64{}, map[int]uint{},
		map[int]uint8{}, map[int]uint16{}, map[int]uint32{}, map[int]uint64{}, map[int]uintptr{}, map[int]float32{},
		map[int]float64{}, map[int]complex64{}, map[int]complex128{}, map[int]string{}, map[int]any{},
		// int8 keys
		map[int8]bool{}, map[int8]int{}, map[int8]int8{}, map[int8]int16{}, map[int8]int32{}, map[int8]int64{},
		map[int8]uint{}, map[int8]uint8{}, map[int8]uint16{}, map[int8]uint32{}, map[int8]uint64{}, map[int8]uintptr{},
		map[int8]float32{}, map[int8]float64{}, map[int8]complex64{}, map[int8]complex128{}, map[int8]string{},
		map[int8]any{},
		// int16 keys
		map[int16]bool{}, map[int16]int{}, map[int16]int8{}, map[int16]int16{}, map[int16]int32{}, map[int16]int64{},
		map[int16]uint{}, map[int16]uint8{}, map[int16]uint16{}, map[int16]uint32{}, map[int16]uint64{}, map[int16]uintptr{},
		map[int16]float32{}, map[int16]float64{}, map[int16]complex64{}, map[int16]complex128{}, map[int16]string{},
		map[int16]any{},
		// int32 keys
		map[int32]bool{}, map[int32]int{}, map[int32]int8{}, map[int32]int16{}, map[int32]int32{}, map[int64]int64{},
		map[int32]uint{}, map[int32]uint8{}, map[int32]uint16{}, map[int32]uint32{}, map[int32]uint64{}, map[int32]uintptr{},
		map[int32]float32{}, map[int32]float64{}, map[int32]complex64{}, map[int32]complex128{}, map[int32]string{},
		map[int32]any{},
		// int64 keys
		map[int64]bool{}, map[int64]int{}, map[int64]int8{}, map[int64]int16{}, map[int64]int32{}, map[int64]int64{},
		map[int64]uint{}, map[int64]uint8{}, map[int64]uint16{}, map[int64]uint32{}, map[int64]uint64{}, map[int64]uintptr{},
		map[int64]float32{}, map[int64]float64{}, map[int64]complex64{}, map[int64]complex128{}, map[int64]string{},
		map[int64]any{},
		// uint keys
		map[uint]bool{}, map[uint]int{}, map[uint]int8{}, map[uint]int16{}, map[uint]int32{}, map[uint]int64{},
		map[uint]uint{}, map[uint]uint8{}, map[uint]uint16{}, map[uint]uint32{}, map[uint]uint64{}, map[uint]uintptr{},
		map[uint]float32{}, map[uint]float64{}, map[uint]complex64{}, map[uint]complex128{}, map[uint]string{},
		map[uint]any{},
		// uint8 keys
		map[uint8]bool{}, map[uint8]int{}, map[uint8]int8{}, map[uint8]int16{}, map[uint8]int32{}, map[uint8]int64{},
		map[uint8]uint{}, map[uint8]uint8{}, map[uint8]uint16{}, map[uint8]uint32{}, map[uint8]uint64{}, map[uint8]uintptr{},
		map[uint8]float32{}, map[uint8]float64{}, map[uint8]complex64{}, map[uint8]complex128{}, map[uint8]string{},
		map[uint8]any{},
		// uint16 keys
		map[uint16]bool{}, map[uint16]int{}, map[uint16]int8{}, map[uint16]int16{}, map[uint16]int32{},
		map[uint16]int64{}, map[uint16]uint{}, map[uint16]uint8{}, map[uint16]uint16{}, map[uint16]uint32{},
		map[uint16]uint64{}, map[uint16]uintptr{}, map[uint16]float32{}, map[uint16]float64{}, map[uint16]complex64{},
		map[uint16]complex128{}, map[uint16]string{}, map[uint16]any{},
		// uint32 keys
		map[uint32]bool{}, map[uint32]int{}, map[uint32]int8{}, map[uint32]int16{}, map[uint32]int32{}, map[uint32]int64{},
		map[uint32]uint{}, map[uint32]uint8{}, map[uint32]uint16{}, map[uint32]uint32{}, map[uint32]uint64{},
		map[uint32]uintptr{}, map[uint32]float32{}, map[uint32]float64{}, map[uint32]complex64{}, map[uint32]complex128{},
		map[uint32]string{}, map[uint32]any{},
		// uint64 keys
		map[uint64]bool{}, map[uint64]int{}, map[uint64]int8{}, map[uint64]int16{}, map[uint64]int32{}, map[uint64]int64{},
		map[uint64]uint{}, map[uint64]uint8{}, map[uint64]uint16{}, map[uint64]uint32{}, map[uint64]uint64{},
		map[uint64]uintptr{}, map[uint64]float32{}, map[uint64]float64{}, map[uint64]complex64{}, map[uint64]complex128{},
		map[uint64]string{}, map[uint64]any{},
		// uintptr keys
		map[uintptr]bool{}, map[uintptr]int{}, map[uintptr]int8{}, map[uintptr]int16{}, map[uintptr]int32{},
		map[uintptr]int64{}, map[uintptr]uint{}, map[uintptr]uint8{}, map[uintptr]uint16{}, map[uintptr]uint32{},
		map[uintptr]uint64{}, map[uintptr]uintptr{}, map[uintptr]float32{}, map[uintptr]float64{}, map[uintptr]complex64{},
		map[uintptr]complex128{}, map[uintptr]string{}, map[uintptr]any{},
		// float32 keys
		map[float32]bool{}, map[float32]int{}, map[float32]int8{}, map[float32]int16{}, map[float32]int32{},
		map[float32]int64{}, map[float32]uint{}, map[float32]uint8{}, map[float32]uint16{}, map[float32]uint32{},
		map[float32]uint64{}, map[float32]uintptr{}, map[float32]float32{}, map[float32]float64{}, map[float32]complex64{},
		map[float32]complex128{}, map[float32]string{}, map[float32]any{},
		// float64 keys
		map[float64]bool{}, map[float64]int{}, map[float64]int8{}, map[float64]int16{}, map[float64]int32{}, map[float64]int64{},
		map[float64]uint{}, map[float64]uint8{}, map[float64]uint16{}, map[float64]uint32{}, map[float64]uint64{},
		map[float64]uintptr{}, map[float64]float32{}, map[float64]float64{}, map[float64]complex64{}, map[float64]complex128{},
		map[float64]string{}, map[float64]any{},
		// complex64 keys
		map[complex64]bool{}, map[complex64]int{}, map[complex64]int8{}, map[complex64]int16{}, map[complex64]int32{},
		map[complex64]int64{}, map[complex64]uint{}, map[complex64]uint8{}, map[complex64]uint16{}, map[complex64]uint32{},
		map[complex64]uint64{}, map[complex64]uintptr{}, map[complex64]float32{}, map[complex64]float64{}, map[complex64]complex64{},
		map[complex64]complex128{}, map[complex64]string{}, map[complex64]any{},
		// complex128 keys
		map[complex128]bool{}, map[complex128]int{}, map[complex128]int8{}, map[complex128]int16{}, map[complex128]int32{},
		map[complex128]int64{}, map[complex128]uint{}, map[complex128]uint8{}, map[complex128]uint16{}, map[complex128]uint32{},
		map[complex128]uint64{}, map[complex128]uintptr{}, map[complex128]float32{}, map[complex128]float64{},
		map[complex128]complex64{}, map[complex128]complex128{}, map[complex128]string{}, map[complex128]any{},
		// string keys
		map[string]bool{}, map[string]int{}, map[string]int8{}, map[string]int16{}, map[string]int32{}, map[string]int64{},
		map[string]uint{}, map[string]uint8{}, map[string]uint16{}, map[string]uint32{}, map[string]uint64{},
		map[string]uintptr{}, map[string]float32{}, map[string]float64{}, map[string]complex64{}, map[string]complex128{},
		map[string]string{}, map[string]any{},
		// any keys
		map[any]bool{}, map[any]int{}, map[any]int8{}, map[any]int16{}, map[any]int32{}, map[any]int64{}, map[any]uint{},
		map[any]uint8{}, map[any]uint16{}, map[any]uint32{}, map[any]uint64{}, map[any]uintptr{}, map[any]float32{},
		map[any]float64{}, map[any]complex64{}, map[any]complex128{}, map[any]string{}, map[any]any{},
	}
}

func primitiveWithPtrKeyValues() []interface{} {
	return []interface{}{
		// bool keys
		map[*bool]bool{}, map[*bool]int{}, map[*bool]int8{}, map[*bool]int16{}, map[*bool]int32{}, map[*bool]int64{},
		map[*bool]uint{}, map[*bool]uint8{}, map[*bool]uint16{}, map[*bool]uint32{}, map[*bool]uint64{}, map[*bool]uintptr{},
		map[*bool]float32{}, map[*bool]float64{}, map[*bool]complex64{}, map[*bool]complex128{}, map[*bool]string{},
		map[*bool]any{},
		// int keys
		map[*int]bool{}, map[*int]int{}, map[*int]int8{}, map[*int]int16{}, map[*int]int32{}, map[*int]int64{},
		map[*int]uint{}, map[*int]uint8{}, map[*int]uint16{}, map[*int]uint32{}, map[*int]uint64{}, map[*int]uintptr{},
		map[*int]float32{}, map[*int]float64{}, map[*int]complex64{}, map[*int]complex128{}, map[*int]string{},
		map[*int]any{},
		// int8 keys
		map[*int8]bool{}, map[*int8]int{}, map[*int8]int8{}, map[*int8]int16{}, map[*int8]int32{}, map[*int8]int64{},
		map[*int8]uint{}, map[*int8]uint8{}, map[*int8]uint16{}, map[*int8]uint32{}, map[*int8]uint64{}, map[*int8]uintptr{},
		map[*int8]float32{}, map[*int8]float64{}, map[*int8]complex64{}, map[*int8]complex128{}, map[*int8]string{},
		map[*int8]any{},
		// int16 keys
		map[*int16]bool{}, map[*int16]int{}, map[*int16]int8{}, map[*int16]int16{}, map[*int16]int32{}, map[*int16]int64{},
		map[*int16]uint{}, map[*int16]uint8{}, map[*int16]uint16{}, map[*int16]uint32{}, map[*int16]uint64{}, map[*int16]uintptr{},
		map[*int16]float32{}, map[*int16]float64{}, map[*int16]complex64{}, map[*int16]complex128{}, map[*int16]string{},
		map[*int16]any{},
		// int32 keys
		map[*int32]bool{}, map[*int32]int{}, map[*int32]int8{}, map[*int32]int16{}, map[*int32]int32{}, map[*int32]int64{},
		map[*int32]uint{}, map[*int32]uint8{}, map[*int32]uint16{}, map[*int32]uint32{}, map[*int32]uint64{},
		map[*int32]uintptr{}, map[*int32]float32{}, map[*int32]float64{}, map[*int32]complex64{}, map[*int32]complex128{},
		map[*int32]string{}, map[*int32]any{},
		// int64 keys
		map[*int64]bool{}, map[*int64]int{}, map[*int64]int8{}, map[*int64]int16{}, map[*int64]int32{}, map[*int64]int64{},
		map[*int64]uint{}, map[*int64]uint8{}, map[*int64]uint16{}, map[*int64]uint32{}, map[*int64]uint64{},
		map[*int64]uintptr{}, map[*int64]float32{}, map[*int64]float64{}, map[*int64]complex64{}, map[*int64]complex128{},
		map[*int64]string{}, map[*int64]any{},
		// uint keys
		map[*uint]bool{}, map[*uint]int{}, map[*uint]int8{}, map[*uint]int16{}, map[*uint]int32{}, map[*uint]int64{},
		map[*uint]uint{}, map[*uint]uint8{}, map[*uint]uint16{}, map[*uint]uint32{}, map[*uint]uint64{}, map[*uint]uintptr{},
		map[*uint]float32{}, map[*uint]float64{}, map[*uint]complex64{}, map[*uint]complex128{}, map[*uint]string{},
		map[*uint]any{},
		// uint8 keys
		map[*uint8]bool{}, map[*uint8]int{}, map[*uint8]int8{}, map[*uint8]int16{}, map[*uint8]int32{}, map[*uint8]int64{},
		map[*uint8]uint{}, map[*uint8]uint8{}, map[*uint8]uint16{}, map[*uint8]uint32{}, map[*uint8]uint64{},
		map[*uint8]uintptr{}, map[*uint8]float32{}, map[*uint8]float64{}, map[*uint8]complex64{}, map[*uint8]complex128{},
		map[*uint8]string{}, map[*uint8]any{},
		// uint16 keys
		map[*uint16]bool{}, map[*uint16]int{}, map[*uint16]int8{}, map[*uint16]int16{}, map[*uint16]int32{}, map[*uint16]int64{},
		map[*uint16]uint{}, map[*uint16]uint8{}, map[*uint16]uint16{}, map[*uint16]uint32{}, map[*uint16]uint64{},
		map[*uint16]uintptr{}, map[*uint16]float32{}, map[*uint16]float64{}, map[*uint16]complex64{}, map[*uint16]complex128{},
		map[*uint16]string{}, map[*uint16]any{},
		// uint32 keys
		map[*uint32]bool{}, map[*uint32]int{}, map[*uint32]int8{}, map[*uint32]int16{}, map[*uint32]int32{},
		map[*uint32]int64{}, map[*uint32]uint{}, map[*uint32]uint8{}, map[*uint32]uint16{}, map[*uint32]uint32{},
		map[*uint32]uint64{}, map[*uint32]uintptr{}, map[*uint32]float32{}, map[*uint32]float64{}, map[*uint32]complex64{},
		map[*uint32]complex128{}, map[*uint32]string{}, map[*uint32]any{},
		// uint64 keys
		map[*uint64]bool{}, map[*uint64]int{}, map[*uint64]int8{}, map[*uint64]int16{}, map[*uint64]int32{},
		map[*uint64]int64{}, map[*uint64]uint{}, map[*uint64]uint8{}, map[*uint64]uint16{}, map[*uint64]uint32{},
		map[*uint64]uint64{}, map[*uint64]uintptr{}, map[*uint64]float32{}, map[*uint64]float64{}, map[*uint64]complex64{},
		map[*uint64]complex128{}, map[*uint64]string{}, map[*uint64]any{},
		// uintptr keys
		map[*uintptr]bool{}, map[*uintptr]int{}, map[*uintptr]int8{}, map[*uintptr]int16{}, map[*uintptr]int32{},
		map[*uintptr]int64{}, map[*uintptr]uint{}, map[*uintptr]uint8{}, map[*uintptr]uint16{}, map[*uintptr]uint32{},
		map[*uintptr]uint64{}, map[*uintptr]uintptr{}, map[*uintptr]float32{}, map[*uintptr]float64{}, map[*uintptr]complex64{},
		map[*uintptr]complex128{}, map[*uintptr]string{}, map[*uintptr]any{},
		// float32 keys
		map[*float32]bool{}, map[*float32]int{}, map[*float32]int8{}, map[*float32]int16{}, map[*float32]int32{},
		map[*float32]int64{}, map[*float32]uint{}, map[*float32]uint8{}, map[*float32]uint16{}, map[*float32]uint32{},
		map[*float32]uint64{}, map[*float32]uintptr{}, map[*float32]float32{}, map[*float32]float64{},
		map[*float32]complex64{}, map[*float32]complex128{}, map[*float32]string{}, map[*float32]any{},
		// float64 keys
		map[*float64]bool{}, map[*float64]int{}, map[*float64]int8{}, map[*float64]int16{}, map[*float64]int32{},
		map[*float64]int64{}, map[*float64]uint{}, map[*float64]uint8{}, map[*float64]uint16{}, map[*float64]uint32{},
		map[*float64]uint64{}, map[*float64]uintptr{}, map[*float64]float32{}, map[*float64]float64{}, map[*float64]complex64{},
		map[*float64]complex128{}, map[*float64]string{}, map[*float64]any{},
		// complex64 keys
		map[*complex64]bool{}, map[*complex64]int{}, map[*complex64]int8{}, map[*complex64]int16{}, map[*complex64]int32{},
		map[*complex64]int64{}, map[*complex64]uint{}, map[*complex64]uint8{}, map[*complex64]uint16{}, map[*complex64]uint32{},
		map[*complex64]uint64{}, map[*complex64]uintptr{}, map[*complex64]float32{}, map[*complex64]float64{},
		map[*complex64]complex64{}, map[*complex64]complex128{}, map[*complex64]string{}, map[*complex64]any{},
		// complex128 keys
		map[*complex128]bool{}, map[*complex128]int{}, map[*complex128]int8{}, map[*complex128]int16{}, map[*complex128]int32{},
		map[*complex128]int64{}, map[*complex128]uint{}, map[*complex128]uint8{}, map[*complex128]uint16{}, map[*complex128]uint32{},
		map[*complex128]uint64{}, map[*complex128]uintptr{}, map[*complex128]float32{}, map[*complex128]float64{},
		map[*complex128]complex64{}, map[*complex128]complex128{}, map[*complex128]string{}, map[*complex128]any{},
		// string keys
		map[*string]bool{}, map[*string]int{}, map[*string]int8{}, map[*string]int16{}, map[*string]int32{},
		map[*string]int64{}, map[*string]uint{}, map[*string]uint8{}, map[*string]uint16{}, map[*string]uint32{},
		map[*string]uint64{}, map[*string]uintptr{}, map[*string]float32{}, map[*string]float64{}, map[*string]complex64{},
		map[*string]complex128{}, map[*string]string{}, map[*string]any{},
		// any keys
		map[*any]bool{}, map[*any]int{}, map[*any]int8{}, map[*any]int16{}, map[*any]int32{}, map[*any]int64{},
		map[*any]uint{}, map[*any]uint8{}, map[*any]uint16{}, map[*any]uint32{}, map[*any]uint64{}, map[*any]uintptr{},
		map[*any]float32{}, map[*any]float64{}, map[*any]complex64{}, map[*any]complex128{}, map[*any]string{},
		map[*any]any{},
	}
}

func primitiveWithPtrValueValues() []interface{} {
	return []interface{}{
		// bool keys
		map[bool]*bool{}, map[bool]*int{}, map[bool]*int8{}, map[bool]*int16{}, map[bool]*int32{}, map[bool]*int64{},
		map[bool]*uint{}, map[bool]*uint8{}, map[bool]*uint16{}, map[bool]*uint32{}, map[bool]*uint64{}, map[bool]*uintptr{},
		map[bool]*float32{}, map[bool]*float64{}, map[bool]*complex64{}, map[bool]*complex128{}, map[bool]*string{},
		map[bool]*any{},
		// int keys
		map[int]*bool{}, map[int]int{}, map[int]int8{}, map[int]int16{}, map[int]int32{}, map[int]int64{}, map[int]*uint{},
		map[int]*uint8{}, map[int]*uint16{}, map[int]*uint32{}, map[int]*uint64{}, map[int]*uintptr{}, map[int]*float32{},
		map[int]*float64{}, map[int]*complex64{}, map[int]*complex128{}, map[int]*string{}, map[int]*any{},
		// int8 keys
		map[int8]*bool{}, map[int8]*int{}, map[int8]*int8{}, map[int8]*int16{}, map[int8]*int32{}, map[int8]*int64{},
		map[int8]*uint{}, map[int8]*uint8{}, map[int8]*uint16{}, map[int8]*uint32{}, map[int8]*uint64{}, map[int8]*uintptr{},
		map[int8]*float32{}, map[int8]*float64{}, map[int8]*complex64{}, map[int8]*complex128{}, map[int8]*string{},
		map[int8]*any{},
		// int16 keys
		map[int16]*bool{}, map[int16]*int{}, map[int16]*int8{}, map[int16]int16{}, map[int16]*int32{}, map[int16]*int64{},
		map[int16]*uint{}, map[int16]*uint8{}, map[int16]*uint16{}, map[int16]*uint32{}, map[int16]*uint64{},
		map[int16]*uintptr{}, map[int16]*float32{}, map[int16]*float64{}, map[int16]*complex64{}, map[int16]*complex128{},
		map[int16]*string{}, map[int16]*any{},
		// int32 keys
		map[int32]*bool{}, map[int32]*int{}, map[int32]*int8{}, map[int32]*int16{}, map[int32]*int32{}, map[int32]*int64{},
		map[int32]*uint{}, map[int32]*uint8{}, map[int32]*uint16{}, map[int32]*uint32{}, map[int32]*uint64{},
		map[int32]*uintptr{}, map[int32]*float32{}, map[int32]*float64{}, map[int32]*complex64{}, map[int32]*complex128{},
		map[int32]*string{}, map[int32]*any{},
		// int64 keys
		map[int64]*bool{}, map[int64]*int{}, map[int64]*int8{}, map[int64]*int16{}, map[int64]*int32{}, map[int64]*int64{},
		map[int64]*uint{}, map[int64]*uint8{}, map[int64]*uint16{}, map[int64]*uint32{}, map[int64]*uint64{},
		map[int64]*uintptr{}, map[int64]*float32{}, map[int64]*float64{}, map[int64]*complex64{}, map[int64]*complex128{},
		map[int64]*string{}, map[int64]*any{},
		// uint keys
		map[uint]*bool{}, map[uint]*int{}, map[uint]*int8{}, map[uint]*int16{}, map[uint]*int32{}, map[uint]*int64{},
		map[uint]*uint{}, map[uint]*uint8{}, map[uint]*uint16{}, map[uint]*uint32{}, map[uint]*uint64{}, map[uint]*uintptr{},
		map[uint]*float32{}, map[uint]*float64{}, map[uint]*complex64{}, map[uint]*complex128{}, map[uint]*string{},
		map[uint]*any{},
		// uint8 keys
		map[uint8]*bool{}, map[uint8]*int{}, map[uint8]*int8{}, map[uint8]*int16{}, map[uint8]*int32{}, map[uint8]*int64{},
		map[uint8]*uint{}, map[uint8]*uint8{}, map[uint8]*uint16{}, map[uint8]*uint32{}, map[uint8]*uint64{},
		map[uint8]*uintptr{}, map[uint8]*float32{}, map[uint8]*float64{}, map[uint8]*complex64{}, map[uint8]*complex128{},
		map[uint8]*string{}, map[uint8]*any{},
		// uint16 keys
		map[uint16]*bool{}, map[uint16]*int{}, map[uint16]*int8{}, map[uint16]*int16{}, map[uint16]*int32{}, map[uint16]*int64{},
		map[uint16]*uint{}, map[uint16]*uint8{}, map[uint16]*uint16{}, map[uint16]*uint32{}, map[uint16]*uint64{},
		map[uint16]*uintptr{}, map[uint16]*float32{}, map[uint16]*float64{}, map[uint16]*complex64{}, map[uint16]*complex128{},
		map[uint16]*string{}, map[uint16]*any{},
		// uint32 keys
		map[uint32]*bool{}, map[uint32]*int{}, map[uint32]*int8{}, map[uint32]*int16{}, map[uint32]*int32{}, map[uint32]*int64{},
		map[uint32]*uint{}, map[uint32]*uint8{}, map[uint32]*uint16{}, map[uint32]*uint32{}, map[uint32]*uint64{},
		map[uint32]*uintptr{}, map[uint32]*float32{}, map[uint32]*float64{}, map[uint32]*complex64{}, map[uint32]*complex128{},
		map[uint32]*string{}, map[uint32]*any{},
		// uint64 keys
		map[uint64]*bool{}, map[uint64]*int{}, map[uint64]*int8{}, map[uint64]*int16{}, map[uint64]*int32{}, map[uint64]*int64{},
		map[uint64]*uint{}, map[uint64]*uint8{}, map[uint64]*uint16{}, map[uint64]*uint32{}, map[uint64]*uint64{},
		map[uint64]*uintptr{}, map[uint64]*float32{}, map[uint64]*float64{}, map[uint64]*complex64{}, map[uint64]*complex128{},
		map[uint64]*string{}, map[uint64]*any{},
		// uintptr keys
		map[uintptr]*bool{}, map[uintptr]*int{}, map[uintptr]*int8{}, map[uintptr]*int16{}, map[uintptr]*int32{},
		map[uintptr]*int64{}, map[uintptr]*uint{}, map[uintptr]*uint8{}, map[uintptr]*uint16{}, map[uintptr]*uint32{},
		map[uintptr]*uint64{}, map[uintptr]*uintptr{}, map[uintptr]*float32{}, map[uintptr]*float64{}, map[uintptr]*complex64{},
		map[uintptr]*complex128{}, map[uintptr]*string{}, map[uintptr]*any{},
		// float32 keys
		map[float32]*bool{}, map[float32]*int{}, map[float32]*int8{}, map[float32]*int16{}, map[float32]*int32{},
		map[float32]*int64{}, map[float32]*uint{}, map[float32]*uint8{}, map[float32]*uint16{}, map[float32]*uint32{},
		map[float32]*uint64{}, map[float32]*uintptr{}, map[float32]*float32{}, map[float32]*float64{}, map[float32]*complex64{},
		map[float32]*complex128{}, map[float32]*string{}, map[float32]*any{},
		// float64 keys
		map[float64]*bool{}, map[float64]*int{}, map[float64]*int8{}, map[float64]*int16{}, map[float64]*int32{},
		map[float64]*int64{}, map[float64]*uint{}, map[float64]*uint8{}, map[float64]*uint16{}, map[float64]*uint32{},
		map[float64]*uint64{}, map[float64]*uintptr{}, map[float64]*float32{}, map[float64]*float64{}, map[float64]*complex64{},
		map[float64]*complex128{}, map[float64]*string{}, map[float64]*any{},
		// complex64 keys
		map[complex64]*bool{}, map[complex64]*int{}, map[complex64]*int8{}, map[complex64]*int16{}, map[complex64]*int32{},
		map[complex64]*int64{}, map[complex64]*uint{}, map[complex64]*uint8{}, map[complex64]*uint16{}, map[complex64]*uint32{},
		map[complex64]*uint64{}, map[complex64]*uintptr{}, map[complex64]*float32{}, map[complex64]*float64{}, map[complex64]*complex64{},
		map[complex64]*complex128{}, map[complex64]*string{}, map[complex64]*any{},
		// complex128 keys
		map[complex128]*bool{}, map[complex128]*int{}, map[complex128]*int8{}, map[complex128]*int16{}, map[complex128]*int32{},
		map[complex128]*int64{}, map[complex128]*uint{}, map[complex128]*uint8{}, map[complex128]*uint16{}, map[complex128]*uint32{},
		map[complex128]*uint64{}, map[complex128]*uintptr{}, map[complex128]*float32{}, map[complex128]*float64{}, map[complex128]*complex64{},
		map[complex128]*complex128{}, map[complex128]*string{}, map[complex128]*any{},
		// string keys
		map[string]*bool{}, map[string]*int{}, map[string]*int8{}, map[string]*int16{}, map[string]*int32{},
		map[string]*int64{}, map[string]*uint{}, map[string]*uint8{}, map[string]*uint16{}, map[string]*uint32{},
		map[string]*uint64{}, map[string]*uintptr{}, map[string]*float32{}, map[string]*float64{}, map[string]*complex64{},
		map[string]*complex128{}, map[string]*string{}, map[string]*any{},
		// any keys
		map[any]*bool{}, map[any]*int{}, map[any]*int8{}, map[any]*int16{}, map[any]*int32{},
		map[any]*int64{}, map[any]*uint{}, map[any]*uint8{}, map[any]*uint16{}, map[any]*uint32{},
		map[any]*uint64{}, map[any]*uintptr{}, map[any]*float32{}, map[any]*float64{}, map[any]*complex64{},
		map[any]*complex128{}, map[any]*string{}, map[any]*any{},
	}
}

func primitiveWithPtrKeyAndValueValues() []interface{} {
	return []interface{}{
		// bool keys
		map[*bool]*bool{}, map[*bool]*int{}, map[*bool]*int8{}, map[*bool]*int16{}, map[*bool]*int32{},
		map[*bool]*int64{}, map[*bool]*uint{}, map[*bool]*uint8{}, map[*bool]*uint16{}, map[*bool]*uint32{},
		map[*bool]*uint64{}, map[*bool]*uintptr{}, map[*bool]*float32{}, map[*bool]*float64{}, map[*bool]*complex64{},
		map[*bool]*complex128{}, map[*bool]*string{}, map[*bool]*any{},
		// int keys
		map[*int]*bool{}, map[*int]*int{}, map[*int]*int8{}, map[*int]*int16{}, map[*int]*int32{}, map[*int]*int64{},
		map[*int]*uint{}, map[*int]*uint8{}, map[*int]*uint16{}, map[*int]*uint32{}, map[*int]*uint64{}, map[*int]*uintptr{},
		map[*int]*float32{}, map[*int]*float64{}, map[*int]*complex64{}, map[*int]*complex128{}, map[*int]*string{},
		map[*int]*any{},
		// int8 keys
		map[*int8]*bool{}, map[*int8]*int{}, map[*int8]*int8{}, map[*int8]*int16{}, map[*int8]*int32{}, map[*int8]*int64{},
		map[*int8]*uint{}, map[*int8]*uint8{}, map[*int8]*uint16{}, map[*int8]*uint32{}, map[*int8]*uint64{},
		map[*int8]*uintptr{}, map[*int8]*float32{}, map[*int8]*float64{}, map[*int8]*complex64{}, map[*int8]*complex128{},
		map[*int8]*string{}, map[*int8]*any{},
		// int16 keys
		map[*int16]*bool{}, map[*int16]*int{}, map[*int16]*int8{}, map[*int16]*int16{}, map[*int16]*int32{}, map[*int16]*int64{},
		map[*int16]*uint{}, map[*int16]*uint8{}, map[*int16]*uint16{}, map[*int16]*uint32{}, map[*int16]*uint64{},
		map[*int16]*uintptr{}, map[*int16]*float32{}, map[*int16]*float64{}, map[*int16]*complex64{}, map[*int16]*complex128{},
		map[*int16]*string{}, map[*int16]*any{},
		// int32 keys
		map[*int32]*bool{}, map[*int32]*int{}, map[*int32]*int8{}, map[*int32]*int16{}, map[*int32]*int32{}, map[*int32]*int64{},
		map[*int32]*uint{}, map[*int32]*uint8{}, map[*int32]*uint16{}, map[*int32]*uint32{}, map[*int32]*uint64{},
		map[*int32]*uintptr{}, map[*int32]*float32{}, map[*int32]*float64{}, map[*int32]*complex64{}, map[*int32]*complex128{},
		map[*int32]*string{}, map[*int32]*any{},
		// int64 keys
		map[*int64]*bool{}, map[*int64]*int{}, map[*int64]*int8{}, map[*int64]*int16{}, map[*int64]*int32{}, map[*int64]*int64{},
		map[*int64]*uint{}, map[*int64]*uint8{}, map[*int64]*uint16{}, map[*int64]*uint32{}, map[*int64]*uint64{},
		map[*int64]*uintptr{}, map[*int64]*float32{}, map[*int64]*float64{}, map[*int64]*complex64{}, map[*int64]*complex128{},
		map[*int64]*string{}, map[*int64]*any{},
		// uint keys
		map[*uint]*bool{}, map[*uint]*int{}, map[*uint]*int8{}, map[*uint]*int16{}, map[*uint]*int32{}, map[*uint]*int64{},
		map[*uint]*uint{}, map[*uint]*uint8{}, map[*uint]*uint16{}, map[*uint]*uint32{}, map[*uint]*uint64{}, map[*uint]*uintptr{},
		map[*uint]*float32{}, map[*uint]*float64{}, map[*uint]*complex64{}, map[*uint]*complex128{}, map[*uint]*string{},
		map[*uint]*any{},
		// uint8 keys
		map[*uint8]*bool{}, map[*uint8]*int{}, map[*uint8]*int8{}, map[*uint8]*int16{}, map[*uint8]*int32{}, map[*uint8]*int64{},
		map[*uint8]*uint{}, map[*uint8]*uint8{}, map[*uint8]*uint16{}, map[*uint8]*uint32{}, map[*uint8]*uint64{},
		map[*uint8]*uintptr{}, map[*uint8]*float32{}, map[*uint8]*float64{}, map[*uint8]*complex64{}, map[*uint8]*complex128{},
		map[*uint8]*string{}, map[*uint8]*any{},
		// uint16 keys
		map[*uint16]*bool{}, map[*uint16]*int{}, map[*uint16]*int8{}, map[*uint16]*int16{}, map[*uint16]*int32{},
		map[*uint16]*int64{}, map[*uint16]*uint{}, map[*uint16]*uint8{}, map[*uint16]*uint16{}, map[*uint16]*uint32{},
		map[*uint16]*uint64{}, map[*uint16]*uintptr{}, map[*uint16]*float32{}, map[*uint16]*float64{}, map[*uint16]*complex64{},
		map[*uint16]*complex128{}, map[*uint16]*string{}, map[*uint16]*any{},
		// uint32 keys
		map[*uint32]*bool{}, map[*uint32]*int{}, map[*uint32]*int8{}, map[*uint32]*int16{}, map[*uint32]*int32{},
		map[*uint32]*int64{}, map[*uint32]*uint{}, map[*uint32]*uint8{}, map[*uint32]*uint16{}, map[*uint32]*uint32{},
		map[*uint32]*uint64{}, map[*uint32]*uintptr{}, map[*uint32]*float32{}, map[*uint32]*float64{}, map[*uint32]*complex64{},
		map[*uint32]*complex128{}, map[*uint32]*string{}, map[*uint32]*any{},
		// uint64 keys
		map[*uint64]*bool{}, map[*uint64]*int{}, map[*uint64]*int8{}, map[*uint64]*int16{}, map[*uint64]*int32{},
		map[*uint64]*int64{}, map[*uint64]*uint{}, map[*uint64]*uint8{}, map[*uint64]*uint16{}, map[*uint64]*uint32{},
		map[*uint64]*uint64{}, map[*uint64]*uintptr{}, map[*uint64]*float32{}, map[*uint64]*float64{}, map[*uint64]*complex64{},
		map[*uint64]*complex128{}, map[*uint64]*string{}, map[*uint64]*any{},
		// uintptr keys
		map[*uintptr]*bool{}, map[*uintptr]*int{}, map[*uintptr]*int8{}, map[*uintptr]*int16{}, map[*uintptr]*int32{},
		map[*uintptr]*int64{}, map[*uintptr]*uint{}, map[*uintptr]*uint8{}, map[*uintptr]*uint16{}, map[*uintptr]*uint32{},
		map[*uintptr]*uint64{}, map[*uintptr]*uintptr{}, map[*uintptr]*float32{}, map[*uintptr]*float64{}, map[*uintptr]*complex64{},
		map[*uintptr]*complex128{}, map[*uintptr]*string{}, map[*uintptr]*any{},
		// float32 keys
		map[*float32]*bool{}, map[*float32]*int{}, map[*float32]*int8{}, map[*float32]*int16{}, map[*float32]*int32{},
		map[*float32]*int64{}, map[*float32]*uint{}, map[*float32]*uint8{}, map[*float32]*uint16{}, map[*float32]*uint32{},
		map[*float32]*uint64{}, map[*float32]*uintptr{}, map[*float32]*float32{}, map[*float32]*float64{},
		map[*float32]*complex64{}, map[*float32]*complex128{}, map[*float32]*string{}, map[*float32]*any{},
		// float64 keys
		map[*float64]*bool{}, map[*float64]*int{}, map[*float64]*int8{}, map[*float64]*int16{}, map[*float64]*int32{},
		map[*float64]*int64{}, map[*float64]*uint{}, map[*float64]*uint8{}, map[*float64]*uint16{}, map[*float64]*uint32{},
		map[*float64]*uint64{}, map[*float64]*uintptr{}, map[*float64]*float32{}, map[*float64]*float64{},
		map[*float64]*complex64{}, map[*float64]*complex128{}, map[*float64]*string{}, map[*float64]*any{},
		// complex64 keys
		map[*complex64]*bool{}, map[*complex64]*int{}, map[*complex64]*int8{}, map[*complex64]*int16{}, map[*complex64]*int32{},
		map[*complex64]*int64{}, map[*complex64]*uint{}, map[*complex64]*uint8{}, map[*complex64]*uint16{}, map[*complex64]*uint32{},
		map[*complex64]*uint64{}, map[*complex64]*uintptr{}, map[*complex64]*float32{}, map[*complex64]*float64{},
		map[*complex64]*complex64{}, map[*complex64]*complex128{}, map[*complex64]*string{}, map[*complex64]*any{},
		// complex128 keys
		map[*complex128]*bool{}, map[*complex128]*int{}, map[*complex128]*int8{}, map[*complex128]*int16{}, map[*complex128]*int32{},
		map[*complex128]*int64{}, map[*complex128]*uint{}, map[*complex128]*uint8{}, map[*complex128]*uint16{}, map[*complex128]*uint32{},
		map[*complex128]*uint64{}, map[*complex128]*uintptr{}, map[*complex128]*float32{}, map[*complex128]*float64{},
		map[*complex128]*complex64{}, map[*complex128]*complex128{}, map[*complex128]*string{}, map[*complex128]*any{},
		// string keys
		map[*string]*bool{}, map[*string]*int{}, map[*string]*int8{}, map[*string]*int16{}, map[*string]*int32{},
		map[*string]*int64{}, map[*string]*uint{}, map[*string]*uint8{}, map[*string]*uint16{}, map[*string]*uint32{},
		map[*string]*uint64{}, map[*string]*uintptr{}, map[*string]*float32{}, map[*string]*float64{},
		map[*string]*complex64{}, map[*string]*complex128{}, map[*string]*string{}, map[*string]*any{},
		// any keys
		map[*any]*bool{}, map[*any]*int{}, map[*any]*int8{}, map[*any]*int16{}, map[*any]*int32{},
		map[*any]*int64{}, map[*any]*uint{}, map[*any]*uint8{}, map[*any]*uint16{}, map[*any]*uint32{},
		map[*any]*uint64{}, map[*any]*uintptr{}, map[*any]*float32{}, map[*any]*float64{}, map[*any]*complex64{},
		map[*any]*complex128{}, map[*any]*string{}, map[*any]*any{},
	}
}

func edgeValues() []interface{} {
	return []interface{}{
		map[unsafe.Pointer]unsafe.Pointer{}, map[*unsafe.Pointer]unsafe.Pointer{},
		map[unsafe.Pointer]*unsafe.Pointer{}, map[*unsafe.Pointer]*unsafe.Pointer{},
		map[**string]string{}, map[string]**string{}, map[*unsafe.Pointer]map[string]string{},
	}
}
