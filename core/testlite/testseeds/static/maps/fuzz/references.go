// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"slices"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/seeders/random"
)

func referencesSet() []interface{} {
	return slices.Concat(primitiveRefs(), primitiveWithPtrKeyRefs(), primitiveWithPtrValueRefs(),
		primitiveWithPtrKeyAndValueRefs(), edgeRefs())
}

func primitiveRefs() []interface{} {
	return []interface{}{
		// bool keys
		random.SingleOf[*map[bool]bool](), random.SingleOf[*map[bool]int](), random.SingleOf[*map[bool]int8](),
		random.SingleOf[*map[bool]int16](), random.SingleOf[*map[bool]int32](), random.SingleOf[*map[bool]int64](),
		random.SingleOf[*map[bool]uint](), random.SingleOf[*map[bool]uint8](), random.SingleOf[*map[bool]uint16](),
		random.SingleOf[*map[bool]uint32](), random.SingleOf[*map[bool]uint64](), random.SingleOf[*map[bool]uintptr](),
		random.SingleOf[*map[bool]float32](), random.SingleOf[*map[bool]float64](),
		random.SingleOf[*map[bool]complex64](), random.SingleOf[*map[bool]complex128](), random.SingleOf[*map[bool]string](),
		random.SingleOf[*map[bool]any](),
		// int keys
		random.SingleOf[*map[int]bool](), random.SingleOf[*map[int]int](), random.SingleOf[*map[int]int8](), random.SingleOf[*map[int]int16](), random.SingleOf[*map[int]int32](), random.SingleOf[*map[int]int64](),
		random.SingleOf[*map[int]uint](), random.SingleOf[*map[int]uint8](), random.SingleOf[*map[int]uint16](), random.SingleOf[*map[int]uint32](), random.SingleOf[*map[int]uint64](), random.SingleOf[*map[int]uintptr](),
		random.SingleOf[*map[int]float32](), random.SingleOf[*map[int]float64](), random.SingleOf[*map[int]complex64](), random.SingleOf[*map[int]complex128](), random.SingleOf[*map[int]string](),
		random.SingleOf[*map[int]any](),
		// int8 keys
		random.SingleOf[*map[int8]bool](), random.SingleOf[*map[int8]int](), random.SingleOf[*map[int8]int8](), random.SingleOf[*map[int8]int16](), random.SingleOf[*map[int8]int32](), random.SingleOf[*map[int8]int64](),
		random.SingleOf[*map[int8]uint](), random.SingleOf[*map[int8]uint8](), random.SingleOf[*map[int8]uint16](), random.SingleOf[*map[int8]uint32](), random.SingleOf[*map[int8]uint64](),
		random.SingleOf[*map[int8]uintptr](), random.SingleOf[*map[int8]float32](), random.SingleOf[*map[int8]float64](), random.SingleOf[*map[int8]complex64](), random.SingleOf[*map[int8]complex128](),
		random.SingleOf[*map[int8]string](), random.SingleOf[*map[int8]any](),
		// int16 keys
		random.SingleOf[*map[int16]bool](), random.SingleOf[*map[int16]int](), random.SingleOf[*map[int16]int8](), random.SingleOf[*map[int16]int16](), random.SingleOf[*map[int16]int32](), random.SingleOf[*map[int16]int64](),
		random.SingleOf[*map[int16]uint](), random.SingleOf[*map[int16]uint8](), random.SingleOf[*map[int16]uint16](), random.SingleOf[*map[int16]uint32](), random.SingleOf[*map[int16]uint64](),
		random.SingleOf[*map[int16]uintptr](), random.SingleOf[*map[int16]float32](), random.SingleOf[*map[int16]float64](), random.SingleOf[*map[int16]complex64](), random.SingleOf[*map[int16]complex128](),
		random.SingleOf[*map[int16]string](), random.SingleOf[*map[int16]any](),
		// int32 keys
		random.SingleOf[*map[int32]bool](), random.SingleOf[*map[int32]int](), random.SingleOf[*map[int32]int8](), random.SingleOf[*map[int32]int16](), random.SingleOf[*map[int32]int32](), random.SingleOf[*map[int64]int64](),
		random.SingleOf[*map[int32]uint](), random.SingleOf[*map[int32]uint8](), random.SingleOf[*map[int32]uint16](), random.SingleOf[*map[int32]uint32](), random.SingleOf[*map[int32]uint64](),
		random.SingleOf[*map[int32]uintptr](), random.SingleOf[*map[int32]float32](), random.SingleOf[*map[int32]float64](), random.SingleOf[*map[int32]complex64](), random.SingleOf[*map[int32]complex128](),
		random.SingleOf[*map[int32]string](), random.SingleOf[*map[int32]any](),
		// int64 keys
		random.SingleOf[*map[int64]bool](), random.SingleOf[*map[int64]int](), random.SingleOf[*map[int64]int8](), random.SingleOf[*map[int64]int16](), random.SingleOf[*map[int64]int32](), random.SingleOf[*map[int64]int64](),
		random.SingleOf[*map[int64]uint](), random.SingleOf[*map[int64]uint8](), random.SingleOf[*map[int64]uint16](), random.SingleOf[*map[int64]uint32](), random.SingleOf[*map[int64]uint64](),
		random.SingleOf[*map[int64]uintptr](), random.SingleOf[*map[int64]float32](), random.SingleOf[*map[int64]float64](), random.SingleOf[*map[int64]complex64](), random.SingleOf[*map[int64]complex128](),
		random.SingleOf[*map[int64]string](), random.SingleOf[*map[int64]any](),
		// uint keys
		random.SingleOf[*map[uint]bool](), random.SingleOf[*map[uint]int](), random.SingleOf[*map[uint]int8](), random.SingleOf[*map[uint]int16](), random.SingleOf[*map[uint]int32](), random.SingleOf[*map[uint]int64](),
		random.SingleOf[*map[uint]uint](), random.SingleOf[*map[uint]uint8](), random.SingleOf[*map[uint]uint16](), random.SingleOf[*map[uint]uint32](), random.SingleOf[*map[uint]uint64](),
		random.SingleOf[*map[uint]uintptr](), random.SingleOf[*map[uint]float32](), random.SingleOf[*map[uint]float64](), random.SingleOf[*map[uint]complex64](), random.SingleOf[*map[uint]complex128](),
		random.SingleOf[*map[uint]string](), random.SingleOf[*map[uint]any](),
		// uint8 keys
		random.SingleOf[*map[uint8]bool](), random.SingleOf[*map[uint8]int](), random.SingleOf[*map[uint8]int8](), random.SingleOf[*map[uint8]int16](), random.SingleOf[*map[uint8]int32](), random.SingleOf[*map[uint8]int64](),
		random.SingleOf[*map[uint8]uint](), random.SingleOf[*map[uint8]uint8](), random.SingleOf[*map[uint8]uint16](), random.SingleOf[*map[uint8]uint32](), random.SingleOf[*map[uint8]uint64](),
		random.SingleOf[*map[uint8]uintptr](), random.SingleOf[*map[uint8]float32](), random.SingleOf[*map[uint8]float64](), random.SingleOf[*map[uint8]complex64](), random.SingleOf[*map[uint8]complex128](),
		random.SingleOf[*map[uint8]string](), random.SingleOf[*map[uint8]any](),
		// uint16 keys
		random.SingleOf[*map[uint16]bool](), random.SingleOf[*map[uint16]int](), random.SingleOf[*map[uint16]int8](), random.SingleOf[*map[uint16]int16](), random.SingleOf[*map[uint16]int32](),
		random.SingleOf[*map[uint16]int64](), random.SingleOf[*map[uint16]uint](), random.SingleOf[*map[uint16]uint8](), random.SingleOf[*map[uint16]uint16](), random.SingleOf[*map[uint16]uint32](),
		random.SingleOf[*map[uint16]uint64](), random.SingleOf[*map[uint16]uintptr](), random.SingleOf[*map[uint16]float32](), random.SingleOf[*map[uint16]float64](), random.SingleOf[*map[uint16]complex64](),
		random.SingleOf[*map[uint16]complex128](), random.SingleOf[*map[uint16]string](), random.SingleOf[*map[uint16]any](),
		// uint32 keys
		random.SingleOf[*map[uint32]bool](), random.SingleOf[*map[uint32]int](), random.SingleOf[*map[uint32]int8](), random.SingleOf[*map[uint32]int16](), random.SingleOf[*map[uint32]int32](),
		random.SingleOf[*map[uint32]int64](), random.SingleOf[*map[uint32]uint](), random.SingleOf[*map[uint32]uint8](), random.SingleOf[*map[uint32]uint16](), random.SingleOf[*map[uint32]uint32](),
		random.SingleOf[*map[uint32]uint64](), random.SingleOf[*map[uint32]uintptr](), random.SingleOf[*map[uint32]float32](), random.SingleOf[*map[uint32]float64](), random.SingleOf[*map[uint32]complex64](),
		random.SingleOf[*map[uint32]complex128](), random.SingleOf[*map[uint32]string](), random.SingleOf[*map[uint32]any](),
		// uint64 keys
		random.SingleOf[*map[uint64]bool](), random.SingleOf[*map[uint64]int](), random.SingleOf[*map[uint64]int8](), random.SingleOf[*map[uint64]int16](), random.SingleOf[*map[uint64]int32](),
		random.SingleOf[*map[uint64]int64](), random.SingleOf[*map[uint64]uint](), random.SingleOf[*map[uint64]uint8](), random.SingleOf[*map[uint64]uint16](), random.SingleOf[*map[uint64]uint32](),
		random.SingleOf[*map[uint64]uint64](), random.SingleOf[*map[uint64]uintptr](), random.SingleOf[*map[uint64]float32](), random.SingleOf[*map[uint64]float64](), random.SingleOf[*map[uint64]complex64](),
		random.SingleOf[*map[uint64]complex128](), random.SingleOf[*map[uint64]string](), random.SingleOf[*map[uint64]any](),
		// uintptr keys
		random.SingleOf[*map[uintptr]bool](), random.SingleOf[*map[uintptr]int](), random.SingleOf[*map[uintptr]int8](), random.SingleOf[*map[uintptr]int16](), random.SingleOf[*map[uintptr]int32](),
		random.SingleOf[*map[uintptr]int64](), random.SingleOf[*map[uintptr]uint](), random.SingleOf[*map[uintptr]uint8](), random.SingleOf[*map[uintptr]uint16](), random.SingleOf[*map[uintptr]uint32](),
		random.SingleOf[*map[uintptr]uint64](), random.SingleOf[*map[uintptr]uintptr](), random.SingleOf[*map[uintptr]float32](), random.SingleOf[*map[uintptr]float64](),
		random.SingleOf[*map[uintptr]complex64](), random.SingleOf[*map[uintptr]complex128](), random.SingleOf[*map[uintptr]string](), random.SingleOf[*map[uintptr]any](),
		// float32 keys
		random.SingleOf[*map[float32]bool](), random.SingleOf[*map[float32]int](), random.SingleOf[*map[float32]int8](), random.SingleOf[*map[float32]int16](), random.SingleOf[*map[float32]int32](),
		random.SingleOf[*map[float32]int64](), random.SingleOf[*map[float32]uint](), random.SingleOf[*map[float32]uint8](), random.SingleOf[*map[float32]uint16](), random.SingleOf[*map[float32]uint32](),
		random.SingleOf[*map[float32]uint64](), random.SingleOf[*map[float32]uintptr](), random.SingleOf[*map[float32]float32](), random.SingleOf[*map[float32]float64](),
		random.SingleOf[*map[float32]complex64](), random.SingleOf[*map[float32]complex128](), random.SingleOf[*map[float32]string](), random.SingleOf[*map[float32]any](),
		// float64 keys
		random.SingleOf[*map[float64]bool](), random.SingleOf[*map[float64]int](), random.SingleOf[*map[float64]int8](), random.SingleOf[*map[float64]int16](), random.SingleOf[*map[float64]int32](),
		random.SingleOf[*map[float64]int64](), random.SingleOf[*map[float64]uint](), random.SingleOf[*map[float64]uint8](), random.SingleOf[*map[float64]uint16](), random.SingleOf[*map[float64]uint32](),
		random.SingleOf[*map[float64]uint64](), random.SingleOf[*map[float64]uintptr](), random.SingleOf[*map[float64]float32](), random.SingleOf[*map[float64]float64](),
		random.SingleOf[*map[float64]complex64](), random.SingleOf[*map[float64]complex128](), random.SingleOf[*map[float64]string](), random.SingleOf[*map[float64]any](),
		// complex64 keys
		random.SingleOf[*map[complex64]bool](), random.SingleOf[*map[complex64]int](), random.SingleOf[*map[complex64]int8](), random.SingleOf[*map[complex64]int16](), random.SingleOf[*map[complex64]int32](),
		random.SingleOf[*map[complex64]int64](), random.SingleOf[*map[complex64]uint](), random.SingleOf[*map[complex64]uint8](), random.SingleOf[*map[complex64]uint16](),
		random.SingleOf[*map[complex64]uint32](), random.SingleOf[*map[complex64]uint64](), random.SingleOf[*map[complex64]uintptr](), random.SingleOf[*map[complex64]float32](),
		random.SingleOf[*map[complex64]float64](), random.SingleOf[*map[complex64]complex64](), random.SingleOf[*map[complex64]complex128](), random.SingleOf[*map[complex64]string](),
		random.SingleOf[*map[complex64]any](),
		// complex128 keys
		random.SingleOf[*map[complex128]bool](), random.SingleOf[*map[complex128]int](), random.SingleOf[*map[complex128]int8](), random.SingleOf[*map[complex128]int16](),
		random.SingleOf[*map[complex128]int32](), random.SingleOf[*map[complex128]int64](), random.SingleOf[*map[complex128]uint](), random.SingleOf[*map[complex128]uint8](),
		random.SingleOf[*map[complex128]uint16](), random.SingleOf[*map[complex128]uint32](), random.SingleOf[*map[complex128]uint64](), random.SingleOf[*map[complex128]uintptr](),
		random.SingleOf[*map[complex128]float32](), random.SingleOf[*map[complex128]float64](), random.SingleOf[*map[complex128]complex64](),
		random.SingleOf[*map[complex128]complex128](), random.SingleOf[*map[complex128]string](), random.SingleOf[*map[complex128]any](),
		// string keys
		random.SingleOf[*map[string]bool](), random.SingleOf[*map[string]int](), random.SingleOf[*map[string]int8](), random.SingleOf[*map[string]int16](), random.SingleOf[*map[string]int32](),
		random.SingleOf[*map[string]int64](), random.SingleOf[*map[string]uint](), random.SingleOf[*map[string]uint8](), random.SingleOf[*map[string]uint16](), random.SingleOf[*map[string]uint32](),
		random.SingleOf[*map[string]uint64](), random.SingleOf[*map[string]uintptr](), random.SingleOf[*map[string]float32](), random.SingleOf[*map[string]float64](), random.SingleOf[*map[string]complex64](),
		random.SingleOf[*map[string]complex128](), random.SingleOf[*map[string]string](), random.SingleOf[*map[string]any](),
		// any keys
		random.SingleOf[*map[any]bool](), random.SingleOf[*map[any]int](), random.SingleOf[*map[any]int8](), random.SingleOf[*map[any]int16](), random.SingleOf[*map[any]int32](), random.SingleOf[*map[any]int64](),
		random.SingleOf[*map[any]uint](), random.SingleOf[*map[any]uint8](), random.SingleOf[*map[any]uint16](), random.SingleOf[*map[any]uint32](), random.SingleOf[*map[any]uint64](), random.SingleOf[*map[any]uintptr](),
		random.SingleOf[*map[any]float32](), random.SingleOf[*map[any]float64](), random.SingleOf[*map[any]complex64](), random.SingleOf[*map[any]complex128](), random.SingleOf[*map[any]string](),
		random.SingleOf[*map[any]any](),
	}
}

func primitiveWithPtrKeyRefs() []interface{} {
	return []interface{}{
		// bool keys
		random.SingleOf[*map[*bool]bool](), random.SingleOf[*map[*bool]int](), random.SingleOf[*map[*bool]int8](), random.SingleOf[*map[*bool]int16](), random.SingleOf[*map[*bool]int32](), random.SingleOf[*map[*bool]int64](),
		random.SingleOf[*map[*bool]uint](), random.SingleOf[*map[*bool]uint8](), random.SingleOf[*map[*bool]uint16](), random.SingleOf[*map[*bool]uint32](), random.SingleOf[*map[*bool]uint64](),
		random.SingleOf[*map[*bool]uintptr](), random.SingleOf[*map[*bool]float32](), random.SingleOf[*map[*bool]float64](), random.SingleOf[*map[*bool]complex64](), random.SingleOf[*map[*bool]complex128](),
		random.SingleOf[*map[*bool]string](), random.SingleOf[*map[*bool]any](),
		// int keys
		random.SingleOf[*map[*int]bool](), random.SingleOf[*map[*int]int](), random.SingleOf[*map[*int]int8](), random.SingleOf[*map[*int]int16](), random.SingleOf[*map[*int]int32](), random.SingleOf[*map[*int]int64](),
		random.SingleOf[*map[*int]uint](), random.SingleOf[*map[*int]uint8](), random.SingleOf[*map[*int]uint16](), random.SingleOf[*map[*int]uint32](), random.SingleOf[*map[*int]uint64](), random.SingleOf[*map[*int]uintptr](),
		random.SingleOf[*map[*int]float32](), random.SingleOf[*map[*int]float64](), random.SingleOf[*map[*int]complex64](), random.SingleOf[*map[*int]complex128](), random.SingleOf[*map[*int]string](),
		random.SingleOf[*map[*int]any](),
		// int8 keys
		random.SingleOf[*map[*int8]bool](), random.SingleOf[*map[*int8]int](), random.SingleOf[*map[*int8]int8](), random.SingleOf[*map[*int8]int16](), random.SingleOf[*map[*int8]int32](), random.SingleOf[*map[*int8]int64](),
		random.SingleOf[*map[*int8]uint](), random.SingleOf[*map[*int8]uint8](), random.SingleOf[*map[*int8]uint16](), random.SingleOf[*map[*int8]uint32](), random.SingleOf[*map[*int8]uint64](),
		random.SingleOf[*map[*int8]uintptr](), random.SingleOf[*map[*int8]float32](), random.SingleOf[*map[*int8]float64](), random.SingleOf[*map[*int8]complex64](), random.SingleOf[*map[*int8]complex128](),
		random.SingleOf[*map[*int8]string](), random.SingleOf[*map[*int8]any](),
		// int16 keys
		random.SingleOf[*map[*int16]bool](), random.SingleOf[*map[*int16]int](), random.SingleOf[*map[*int16]int8](), random.SingleOf[*map[*int16]int16](), random.SingleOf[*map[*int16]int32](),
		random.SingleOf[*map[*int16]int64](), random.SingleOf[*map[*int16]uint](), random.SingleOf[*map[*int16]uint8](), random.SingleOf[*map[*int16]uint16](), random.SingleOf[*map[*int16]uint32](),
		random.SingleOf[*map[*int16]uint64](), random.SingleOf[*map[*int16]uintptr](), random.SingleOf[*map[*int16]float32](), random.SingleOf[*map[*int16]float64](), random.SingleOf[*map[*int16]complex64](),
		random.SingleOf[*map[*int16]complex128](), random.SingleOf[*map[*int16]string](), random.SingleOf[*map[*int16]any](),
		// int32 keys
		random.SingleOf[*map[*int32]bool](), random.SingleOf[*map[*int32]int](), random.SingleOf[*map[*int32]int8](), random.SingleOf[*map[*int32]int16](), random.SingleOf[*map[*int32]int32](),
		random.SingleOf[*map[*int32]int64](), random.SingleOf[*map[*int32]uint](), random.SingleOf[*map[*int32]uint8](), random.SingleOf[*map[*int32]uint16](), random.SingleOf[*map[*int32]uint32](),
		random.SingleOf[*map[*int32]uint64](), random.SingleOf[*map[*int32]uintptr](), random.SingleOf[*map[*int32]float32](), random.SingleOf[*map[*int32]float64](), random.SingleOf[*map[*int32]complex64](),
		random.SingleOf[*map[*int32]complex128](), random.SingleOf[*map[*int32]string](), random.SingleOf[*map[*int32]any](),
		// int64 keys
		random.SingleOf[*map[*int64]bool](), random.SingleOf[*map[*int64]int](), random.SingleOf[*map[*int64]int8](), random.SingleOf[*map[*int64]int16](), random.SingleOf[*map[*int64]int32](),
		random.SingleOf[*map[*int64]int64](), random.SingleOf[*map[*int64]uint](), random.SingleOf[*map[*int64]uint8](), random.SingleOf[*map[*int64]uint16](), random.SingleOf[*map[*int64]uint32](),
		random.SingleOf[*map[*int64]uint64](), random.SingleOf[*map[*int64]uintptr](), random.SingleOf[*map[*int64]float32](), random.SingleOf[*map[*int64]float64](), random.SingleOf[*map[*int64]complex64](),
		random.SingleOf[*map[*int64]complex128](), random.SingleOf[*map[*int64]string](), random.SingleOf[*map[*int64]any](),
		// uint keys
		random.SingleOf[*map[*uint]bool](), random.SingleOf[*map[*uint]int](), random.SingleOf[*map[*uint]int8](), random.SingleOf[*map[*uint]int16](), random.SingleOf[*map[*uint]int32](), random.SingleOf[*map[*uint]int64](),
		random.SingleOf[*map[*uint]uint](), random.SingleOf[*map[*uint]uint8](), random.SingleOf[*map[*uint]uint16](), random.SingleOf[*map[*uint]uint32](), random.SingleOf[*map[*uint]uint64](),
		random.SingleOf[*map[*uint]uintptr](), random.SingleOf[*map[*uint]float32](), random.SingleOf[*map[*uint]float64](), random.SingleOf[*map[*uint]complex64](),
		random.SingleOf[*map[*uint]complex128](), random.SingleOf[*map[*uint]string](), random.SingleOf[*map[*uint]any](),
		// uint8 keys
		random.SingleOf[*map[*uint8]bool](), random.SingleOf[*map[*uint8]int](), random.SingleOf[*map[*uint8]int8](), random.SingleOf[*map[*uint8]int16](), random.SingleOf[*map[*uint8]int32](),
		random.SingleOf[*map[*uint8]int64](), random.SingleOf[*map[*uint8]uint](), random.SingleOf[*map[*uint8]uint8](), random.SingleOf[*map[*uint8]uint16](), random.SingleOf[*map[*uint8]uint32](),
		random.SingleOf[*map[*uint8]uint64](), random.SingleOf[*map[*uint8]uintptr](), random.SingleOf[*map[*uint8]float32](), random.SingleOf[*map[*uint8]float64](), random.SingleOf[*map[*uint8]complex64](),
		random.SingleOf[*map[*uint8]complex128](), random.SingleOf[*map[*uint8]string](), random.SingleOf[*map[*uint8]any](), // uint16 keys
		random.SingleOf[*map[*uint16]bool](), random.SingleOf[*map[*uint16]int](), random.SingleOf[*map[*uint16]int8](), random.SingleOf[*map[*uint16]int16](), random.SingleOf[*map[*uint16]int32](),
		random.SingleOf[*map[*uint16]int64](), random.SingleOf[*map[*uint16]uint](), random.SingleOf[*map[*uint16]uint8](), random.SingleOf[*map[*uint16]uint16](), random.SingleOf[*map[*uint16]uint32](),
		random.SingleOf[*map[*uint16]uint64](), random.SingleOf[*map[*uint16]uintptr](), random.SingleOf[*map[*uint16]float32](), random.SingleOf[*map[*uint16]float64](),
		random.SingleOf[*map[*uint16]complex64](), random.SingleOf[*map[*uint16]complex128](), random.SingleOf[*map[*uint16]string](), random.SingleOf[*map[*uint16]any](),
		// uint32 keys
		random.SingleOf[*map[*uint32]bool](), random.SingleOf[*map[*uint32]int](), random.SingleOf[*map[*uint32]int8](), random.SingleOf[*map[*uint32]int16](), random.SingleOf[*map[*uint32]int32](),
		random.SingleOf[*map[*uint32]int64](), random.SingleOf[*map[*uint32]uint](), random.SingleOf[*map[*uint32]uint8](), random.SingleOf[*map[*uint32]uint16](), random.SingleOf[*map[*uint32]uint32](),
		random.SingleOf[*map[*uint32]uint64](), random.SingleOf[*map[*uint32]uintptr](), random.SingleOf[*map[*uint32]float32](), random.SingleOf[*map[*uint32]float64](),
		random.SingleOf[*map[*uint32]complex64](), random.SingleOf[*map[*uint32]complex128](), random.SingleOf[*map[*uint32]string](), random.SingleOf[*map[*uint32]any](),
		// uint64 keys
		random.SingleOf[*map[*uint64]bool](), random.SingleOf[*map[*uint64]int](), random.SingleOf[*map[*uint64]int8](), random.SingleOf[*map[*uint64]int16](), random.SingleOf[*map[*uint64]int32](),
		random.SingleOf[*map[*uint64]int64](), random.SingleOf[*map[*uint64]uint](), random.SingleOf[*map[*uint64]uint8](), random.SingleOf[*map[*uint64]uint16](), random.SingleOf[*map[*uint64]uint32](),
		random.SingleOf[*map[*uint64]uint64](), random.SingleOf[*map[*uint64]uintptr](), random.SingleOf[*map[*uint64]float32](), random.SingleOf[*map[*uint64]float64](), random.SingleOf[*map[*uint64]complex64](),
		random.SingleOf[*map[*uint64]complex128](), random.SingleOf[*map[*uint64]string](), random.SingleOf[*map[*uint64]any](),
		// uintptr keys
		random.SingleOf[*map[*uintptr]bool](), random.SingleOf[*map[*uintptr]int](), random.SingleOf[*map[*uintptr]int8](), random.SingleOf[*map[*uintptr]int16](), random.SingleOf[*map[*uintptr]int32](),
		random.SingleOf[*map[*uintptr]int64](), random.SingleOf[*map[*uintptr]uint](), random.SingleOf[*map[*uintptr]uint8](), random.SingleOf[*map[*uintptr]uint16](),
		random.SingleOf[*map[*uintptr]uint32](), random.SingleOf[*map[*uintptr]uint64](), random.SingleOf[*map[*uintptr]uintptr](), random.SingleOf[*map[*uintptr]float32](),
		random.SingleOf[*map[*uintptr]float64](), random.SingleOf[*map[*uintptr]complex64](), random.SingleOf[*map[*uintptr]complex128](), random.SingleOf[*map[*uintptr]string](),
		random.SingleOf[*map[*uintptr]any](),
		// float32 keys
		random.SingleOf[*map[*float32]bool](), random.SingleOf[*map[*float32]int](), random.SingleOf[*map[*float32]int8](), random.SingleOf[*map[*float32]int16](), random.SingleOf[*map[*float32]int32](),
		random.SingleOf[*map[*float32]int64](), random.SingleOf[*map[*float32]uint](), random.SingleOf[*map[*float32]uint8](), random.SingleOf[*map[*float32]uint16](),
		random.SingleOf[*map[*float32]uint32](), random.SingleOf[*map[*float32]uint64](), random.SingleOf[*map[*float32]uintptr](), random.SingleOf[*map[*float32]float32](),
		random.SingleOf[*map[*float32]float64](), random.SingleOf[*map[*float32]complex64](), random.SingleOf[*map[*float32]complex128](), random.SingleOf[*map[*float32]string](),
		random.SingleOf[*map[*float32]any](),
		// float64 keys
		random.SingleOf[*map[*float64]bool](), random.SingleOf[*map[*float64]int](), random.SingleOf[*map[*float64]int8](), random.SingleOf[*map[*float64]int16](), random.SingleOf[*map[*float64]int32](),
		random.SingleOf[*map[*float64]int64](), random.SingleOf[*map[*float64]uint](), random.SingleOf[*map[*float64]uint8](), random.SingleOf[*map[*float64]uint16](),
		random.SingleOf[*map[*float64]uint32](), random.SingleOf[*map[*float64]uint64](), random.SingleOf[*map[*float64]uintptr](), random.SingleOf[*map[*float64]float32](),
		random.SingleOf[*map[*float64]float64](), random.SingleOf[*map[*float64]complex64](), random.SingleOf[*map[*float64]complex128](), random.SingleOf[*map[*float64]string](),
		random.SingleOf[*map[*float64]any](),
		// complex64 keys
		random.SingleOf[*map[*complex64]bool](), random.SingleOf[*map[*complex64]int](), random.SingleOf[*map[*complex64]int8](), random.SingleOf[*map[*complex64]int16](),
		random.SingleOf[*map[*complex64]int32](), random.SingleOf[*map[*complex64]int64](), random.SingleOf[*map[*complex64]uint](), random.SingleOf[*map[*complex64]uint8](),
		random.SingleOf[*map[*complex64]uint16](), random.SingleOf[*map[*complex64]uint32](), random.SingleOf[*map[*complex64]uint64](), random.SingleOf[*map[*complex64]uintptr](),
		random.SingleOf[*map[*complex64]float32](), random.SingleOf[*map[*complex64]float64](), random.SingleOf[*map[*complex64]complex64](), random.SingleOf[*map[*complex64]complex128](),
		random.SingleOf[*map[*complex64]string](), random.SingleOf[*map[*complex64]any](),
		// complex128 keys
		random.SingleOf[*map[*complex128]bool](), random.SingleOf[*map[*complex128]int](), random.SingleOf[*map[*complex128]int8](), random.SingleOf[*map[*complex128]int16](),
		random.SingleOf[*map[*complex128]int32](), random.SingleOf[*map[*complex128]int64](), random.SingleOf[*map[*complex128]uint](), random.SingleOf[*map[*complex128]uint8](),
		random.SingleOf[*map[*complex128]uint16](), random.SingleOf[*map[*complex128]uint32](), random.SingleOf[*map[*complex128]uint64](), random.SingleOf[*map[*complex128]uintptr](),
		random.SingleOf[*map[*complex128]float32](), random.SingleOf[*map[*complex128]float64](), random.SingleOf[*map[*complex128]complex64](),
		random.SingleOf[*map[*complex128]complex128](), random.SingleOf[*map[*complex128]string](), random.SingleOf[*map[*complex128]any](),
		// string keys
		random.SingleOf[*map[*string]bool](), random.SingleOf[*map[*string]int](), random.SingleOf[*map[*string]int8](), random.SingleOf[*map[*string]int16](), random.SingleOf[*map[*string]int32](),
		random.SingleOf[*map[*string]int64](), random.SingleOf[*map[*string]uint](), random.SingleOf[*map[*string]uint8](), random.SingleOf[*map[*string]uint16](), random.SingleOf[*map[*string]uint32](),
		random.SingleOf[*map[*string]uint64](), random.SingleOf[*map[*string]uintptr](), random.SingleOf[*map[*string]float32](), random.SingleOf[*map[*string]float64](),
		random.SingleOf[*map[*string]complex64](), random.SingleOf[*map[*string]complex128](), random.SingleOf[*map[*string]string](), random.SingleOf[*map[*string]any](),
		// any keys
		random.SingleOf[*map[*any]bool](), random.SingleOf[*map[*any]int](), random.SingleOf[*map[*any]int8](), random.SingleOf[*map[*any]int16](), random.SingleOf[*map[*any]int32](), random.SingleOf[*map[*any]int64](),
		random.SingleOf[*map[*any]uint](), random.SingleOf[*map[*any]uint8](), random.SingleOf[*map[*any]uint16](), random.SingleOf[*map[*any]uint32](), random.SingleOf[*map[*any]uint64](),
		random.SingleOf[*map[*any]uintptr](), random.SingleOf[*map[*any]float32](), random.SingleOf[*map[*any]float64](), random.SingleOf[*map[*any]complex64](), random.SingleOf[*map[*any]complex128](),
		random.SingleOf[*map[*any]string](), random.SingleOf[*map[*any]any](),
	}
}

func primitiveWithPtrValueRefs() []interface{} {
	return []interface{}{
		// bool keys
		random.SingleOf[*map[bool]*bool](), random.SingleOf[*map[bool]*int](), random.SingleOf[*map[bool]*int8](), random.SingleOf[*map[bool]*int16](), random.SingleOf[*map[bool]*int32](), random.SingleOf[*map[bool]*int64](),
		random.SingleOf[*map[bool]*uint](), random.SingleOf[*map[bool]*uint8](), random.SingleOf[*map[bool]*uint16](), random.SingleOf[*map[bool]*uint32](), random.SingleOf[*map[bool]*uint64](),
		random.SingleOf[*map[bool]*uintptr](), random.SingleOf[*map[bool]*float32](), random.SingleOf[*map[bool]*float64](), random.SingleOf[*map[bool]*complex64](), random.SingleOf[*map[bool]*complex128](),
		random.SingleOf[*map[bool]*string](), random.SingleOf[*map[bool]*any](),
		// int keys
		random.SingleOf[*map[int]*bool](), random.SingleOf[*map[int]int](), random.SingleOf[*map[int]int8](), random.SingleOf[*map[int]int16](), random.SingleOf[*map[int]int32](), random.SingleOf[*map[int]int64](),
		random.SingleOf[*map[int]*uint](), random.SingleOf[*map[int]*uint8](), random.SingleOf[*map[int]*uint16](), random.SingleOf[*map[int]*uint32](), random.SingleOf[*map[int]*uint64](), random.SingleOf[*map[int]*uintptr](),
		random.SingleOf[*map[int]*float32](), random.SingleOf[*map[int]*float64](), random.SingleOf[*map[int]*complex64](), random.SingleOf[*map[int]*complex128](), random.SingleOf[*map[int]*string](),
		random.SingleOf[*map[int]*any](),
		// int8 keys
		random.SingleOf[*map[int8]*bool](), random.SingleOf[*map[int8]*int](), random.SingleOf[*map[int8]*int8](), random.SingleOf[*map[int8]*int16](), random.SingleOf[*map[int8]*int32](), random.SingleOf[*map[int8]*int64](),
		random.SingleOf[*map[int8]*uint](), random.SingleOf[*map[int8]*uint8](), random.SingleOf[*map[int8]*uint16](), random.SingleOf[*map[int8]*uint32](), random.SingleOf[*map[int8]*uint64](),
		random.SingleOf[*map[int8]*uintptr](), random.SingleOf[*map[int8]*float32](), random.SingleOf[*map[int8]*float64](), random.SingleOf[*map[int8]*complex64](), random.SingleOf[*map[int8]*complex128](),
		random.SingleOf[*map[int8]*string](), random.SingleOf[*map[int8]*any](),
		// int16 keys
		random.SingleOf[*map[int16]*bool](), random.SingleOf[*map[int16]*int](), random.SingleOf[*map[int16]*int8](), random.SingleOf[*map[int16]int16](), random.SingleOf[*map[int16]*int32](),
		random.SingleOf[*map[int16]*int64](), random.SingleOf[*map[int16]*uint](), random.SingleOf[*map[int16]*uint8](), random.SingleOf[*map[int16]*uint16](), random.SingleOf[*map[int16]*uint32](),
		random.SingleOf[*map[int16]*uint64](), random.SingleOf[*map[int16]*uintptr](), random.SingleOf[*map[int16]*float32](), random.SingleOf[*map[int16]*float64](), random.SingleOf[*map[int16]*complex64](),
		random.SingleOf[*map[int16]*complex128](), random.SingleOf[*map[int16]*string](), random.SingleOf[*map[int16]*any](),
		// int32 keys
		random.SingleOf[*map[int32]*bool](), random.SingleOf[*map[int32]*int](), random.SingleOf[*map[int32]*int8](), random.SingleOf[*map[int32]*int16](), random.SingleOf[*map[int32]*int32](),
		random.SingleOf[*map[int32]*int64](), random.SingleOf[*map[int32]*uint](), random.SingleOf[*map[int32]*uint8](), random.SingleOf[*map[int32]*uint16](), random.SingleOf[*map[int32]*uint32](),
		random.SingleOf[*map[int32]*uint64](), random.SingleOf[*map[int32]*uintptr](), random.SingleOf[*map[int32]*float32](), random.SingleOf[*map[int32]*float64](), random.SingleOf[*map[int32]*complex64](),
		random.SingleOf[*map[int32]*complex128](), random.SingleOf[*map[int32]*string](), random.SingleOf[*map[int32]*any](),
		// int64 keys
		random.SingleOf[*map[int64]*bool](), random.SingleOf[*map[int64]*int](), random.SingleOf[*map[int64]*int8](), random.SingleOf[*map[int64]*int16](), random.SingleOf[*map[int64]*int32](),
		random.SingleOf[*map[int64]*int64](), random.SingleOf[*map[int64]*uint](), random.SingleOf[*map[int64]*uint8](), random.SingleOf[*map[int64]*uint16](), random.SingleOf[*map[int64]*uint32](),
		random.SingleOf[*map[int64]*uint64](), random.SingleOf[*map[int64]*uintptr](), random.SingleOf[*map[int64]*float32](), random.SingleOf[*map[int64]*float64](), random.SingleOf[*map[int64]*complex64](),
		random.SingleOf[*map[int64]*complex128](), random.SingleOf[*map[int64]*string](), random.SingleOf[*map[int64]*any](),
		// uint keys
		random.SingleOf[*map[uint]*bool](), random.SingleOf[*map[uint]*int](), random.SingleOf[*map[uint]*int8](), random.SingleOf[*map[uint]*int16](), random.SingleOf[*map[uint]*int32](), random.SingleOf[*map[uint]*int64](),
		random.SingleOf[*map[uint]*uint](), random.SingleOf[*map[uint]*uint8](), random.SingleOf[*map[uint]*uint16](), random.SingleOf[*map[uint]*uint32](), random.SingleOf[*map[uint]*uint64](),
		random.SingleOf[*map[uint]*uintptr](), random.SingleOf[*map[uint]*float32](), random.SingleOf[*map[uint]*float64](), random.SingleOf[*map[uint]*complex64](), random.SingleOf[*map[uint]*complex128](),
		random.SingleOf[*map[uint]*string](), random.SingleOf[*map[uint]*any](),
		// uint8 keys
		random.SingleOf[*map[uint8]*bool](), random.SingleOf[*map[uint8]*int](), random.SingleOf[*map[uint8]*int8](), random.SingleOf[*map[uint8]*int16](), random.SingleOf[*map[uint8]*int32](), random.SingleOf[*map[uint8]*int64](),
		random.SingleOf[*map[uint8]*uint](), random.SingleOf[*map[uint8]*uint8](), random.SingleOf[*map[uint8]*uint16](), random.SingleOf[*map[uint8]*uint32](), random.SingleOf[*map[uint8]*uint64](),
		random.SingleOf[*map[uint8]*uintptr](), random.SingleOf[*map[uint8]*float32](), random.SingleOf[*map[uint8]*float64](), random.SingleOf[*map[uint8]*complex64](), random.SingleOf[*map[uint8]*complex128](),
		random.SingleOf[*map[uint8]*string](), random.SingleOf[*map[uint8]*any](),
		// uint16 keys
		random.SingleOf[*map[uint16]*bool](), random.SingleOf[*map[uint16]*int](), random.SingleOf[*map[uint16]*int8](), random.SingleOf[*map[uint16]*int16](), random.SingleOf[*map[uint16]*int32](), random.SingleOf[*map[uint16]*int64](),
		random.SingleOf[*map[uint16]*uint](), random.SingleOf[*map[uint16]*uint8](), random.SingleOf[*map[uint16]*uint16](), random.SingleOf[*map[uint16]*uint32](), random.SingleOf[*map[uint16]*uint64](),
		random.SingleOf[*map[uint16]*uintptr](), random.SingleOf[*map[uint16]*float32](), random.SingleOf[*map[uint16]*float64](), random.SingleOf[*map[uint16]*complex64](), random.SingleOf[*map[uint16]*complex128](),
		random.SingleOf[*map[uint16]*string](), random.SingleOf[*map[uint16]*any](),
		// uint32 keys
		random.SingleOf[*map[uint32]*bool](), random.SingleOf[*map[uint32]*int](), random.SingleOf[*map[uint32]*int8](), random.SingleOf[*map[uint32]*int16](), random.SingleOf[*map[uint32]*int32](), random.SingleOf[*map[uint32]*int64](),
		random.SingleOf[*map[uint32]*uint](), random.SingleOf[*map[uint32]*uint8](), random.SingleOf[*map[uint32]*uint16](), random.SingleOf[*map[uint32]*uint32](), random.SingleOf[*map[uint32]*uint64](),
		random.SingleOf[*map[uint32]*uintptr](), random.SingleOf[*map[uint32]*float32](), random.SingleOf[*map[uint32]*float64](), random.SingleOf[*map[uint32]*complex64](), random.SingleOf[*map[uint32]*complex128](),
		random.SingleOf[*map[uint32]*string](), random.SingleOf[*map[uint32]*any](),
		// uint64 keys
		random.SingleOf[*map[uint64]*bool](), random.SingleOf[*map[uint64]*int](), random.SingleOf[*map[uint64]*int8](), random.SingleOf[*map[uint64]*int16](), random.SingleOf[*map[uint64]*int32](), random.SingleOf[*map[uint64]*int64](),
		random.SingleOf[*map[uint64]*uint](), random.SingleOf[*map[uint64]*uint8](), random.SingleOf[*map[uint64]*uint16](), random.SingleOf[*map[uint64]*uint32](), random.SingleOf[*map[uint64]*uint64](), random.SingleOf[*map[uint64]*uintptr](),
		random.SingleOf[*map[uint64]*float32](), random.SingleOf[*map[uint64]*float64](), random.SingleOf[*map[uint64]*complex64](), random.SingleOf[*map[uint64]*complex128](), random.SingleOf[*map[uint64]*string](),
		random.SingleOf[*map[uint64]*any](),
		// uintptr keys
		random.SingleOf[*map[uintptr]*bool](), random.SingleOf[*map[uintptr]*int](), random.SingleOf[*map[uintptr]*int8](), random.SingleOf[*map[uintptr]*int16](), random.SingleOf[*map[uintptr]*int32](),
		random.SingleOf[*map[uintptr]*int64](), random.SingleOf[*map[uintptr]*uint](), random.SingleOf[*map[uintptr]*uint8](), random.SingleOf[*map[uintptr]*uint16](), random.SingleOf[*map[uintptr]*uint32](),
		random.SingleOf[*map[uintptr]*uint64](), random.SingleOf[*map[uintptr]*uintptr](), random.SingleOf[*map[uintptr]*float32](), random.SingleOf[*map[uintptr]*float64](),
		random.SingleOf[*map[uintptr]*complex64](), random.SingleOf[*map[uintptr]*complex128](), random.SingleOf[*map[uintptr]*string](), random.SingleOf[*map[uintptr]*any](),
		// float32 keys
		random.SingleOf[*map[float32]*bool](), random.SingleOf[*map[float32]*int](), random.SingleOf[*map[float32]*int8](), random.SingleOf[*map[float32]*int16](), random.SingleOf[*map[float32]*int32](),
		random.SingleOf[*map[float32]*int64](), random.SingleOf[*map[float32]*uint](), random.SingleOf[*map[float32]*uint8](), random.SingleOf[*map[float32]*uint16](), random.SingleOf[*map[float32]*uint32](),
		random.SingleOf[*map[float32]*uint64](), random.SingleOf[*map[float32]*uintptr](), random.SingleOf[*map[float32]*float32](), random.SingleOf[*map[float32]*float64](), random.SingleOf[*map[float32]*complex64](),
		random.SingleOf[*map[float32]*complex128](), random.SingleOf[*map[float32]*string](), random.SingleOf[*map[float32]*any](),
		// float64 keys
		random.SingleOf[*map[float64]*bool](), random.SingleOf[*map[float64]*int](), random.SingleOf[*map[float64]*int8](), random.SingleOf[*map[float64]*int16](), random.SingleOf[*map[float64]*int32](),
		random.SingleOf[*map[float64]*int64](), random.SingleOf[*map[float64]*uint](), random.SingleOf[*map[float64]*uint8](), random.SingleOf[*map[float64]*uint16](), random.SingleOf[*map[float64]*uint32](),
		random.SingleOf[*map[float64]*uint64](), random.SingleOf[*map[float64]*uintptr](), random.SingleOf[*map[float64]*float32](), random.SingleOf[*map[float64]*float64](),
		random.SingleOf[*map[float64]*complex64](), random.SingleOf[*map[float64]*complex128](), random.SingleOf[*map[float64]*string](), random.SingleOf[*map[float64]*any](),
		// complex64 keys
		random.SingleOf[*map[complex64]*bool](), random.SingleOf[*map[complex64]*int](), random.SingleOf[*map[complex64]*int8](), random.SingleOf[*map[complex64]*int16](), random.SingleOf[*map[complex64]*int32](),
		random.SingleOf[*map[complex64]*int64](), random.SingleOf[*map[complex64]*uint](), random.SingleOf[*map[complex64]*uint8](), random.SingleOf[*map[complex64]*uint16](),
		random.SingleOf[*map[complex64]*uint32](), random.SingleOf[*map[complex64]*uint64](), random.SingleOf[*map[complex64]*uintptr](), random.SingleOf[*map[complex64]*float32](),
		random.SingleOf[*map[complex64]*float64](), random.SingleOf[*map[complex64]*complex64](), random.SingleOf[*map[complex64]*complex128](), random.SingleOf[*map[complex64]*string](),
		random.SingleOf[*map[complex64]*any](),
		// complex128 keys
		random.SingleOf[*map[complex128]*bool](), random.SingleOf[*map[complex128]*int](), random.SingleOf[*map[complex128]*int8](), random.SingleOf[*map[complex128]*int16](),
		random.SingleOf[*map[complex128]*int32](), random.SingleOf[*map[complex128]*int64](), random.SingleOf[*map[complex128]*uint](), random.SingleOf[*map[complex128]*uint8](),
		random.SingleOf[*map[complex128]*uint16](), random.SingleOf[*map[complex128]*uint32](), random.SingleOf[*map[complex128]*uint64](), random.SingleOf[*map[complex128]*uintptr](),
		random.SingleOf[*map[complex128]*float32](), random.SingleOf[*map[complex128]*float64](), random.SingleOf[*map[complex128]*complex64](), random.SingleOf[*map[complex128]*complex128](),
		random.SingleOf[*map[complex128]*string](), random.SingleOf[*map[complex128]*any](),
		// string keys
		random.SingleOf[*map[string]*bool](), random.SingleOf[*map[string]*int](), random.SingleOf[*map[string]*int8](), random.SingleOf[*map[string]*int16](), random.SingleOf[*map[string]*int32](),
		random.SingleOf[*map[string]*int64](), random.SingleOf[*map[string]*uint](), random.SingleOf[*map[string]*uint8](), random.SingleOf[*map[string]*uint16](), random.SingleOf[*map[string]*uint32](),
		random.SingleOf[*map[string]*uint64](), random.SingleOf[*map[string]*uintptr](), random.SingleOf[*map[string]*float32](), random.SingleOf[*map[string]*float64](),
		random.SingleOf[*map[string]*complex64](), random.SingleOf[*map[string]*complex128](), random.SingleOf[*map[string]*string](), random.SingleOf[*map[string]*any](),
		// any keys
		random.SingleOf[*map[any]*bool](), random.SingleOf[*map[any]*int](), random.SingleOf[*map[any]*int8](), random.SingleOf[*map[any]*int16](), random.SingleOf[*map[any]*int32](),
		random.SingleOf[*map[any]*int64](), random.SingleOf[*map[any]*uint](), random.SingleOf[*map[any]*uint8](), random.SingleOf[*map[any]*uint16](), random.SingleOf[*map[any]*uint32](),
		random.SingleOf[*map[any]*uint64](), random.SingleOf[*map[any]*uintptr](), random.SingleOf[*map[any]*float32](), random.SingleOf[*map[any]*float64](), random.SingleOf[*map[any]*complex64](),
		random.SingleOf[*map[any]*complex128](), random.SingleOf[*map[any]*string](), random.SingleOf[*map[any]*any](),
	}
}

func primitiveWithPtrKeyAndValueRefs() []interface{} {
	return []interface{}{
		// bool keys
		random.SingleOf[*map[*bool]*bool](), random.SingleOf[*map[*bool]*int](), random.SingleOf[*map[*bool]*int8](), random.SingleOf[*map[*bool]*int16](), random.SingleOf[*map[*bool]*int32](),
		random.SingleOf[*map[*bool]*int64](), random.SingleOf[*map[*bool]*uint](), random.SingleOf[*map[*bool]*uint8](), random.SingleOf[*map[*bool]*uint16](), random.SingleOf[*map[*bool]*uint32](),
		random.SingleOf[*map[*bool]*uint64](), random.SingleOf[*map[*bool]*uintptr](), random.SingleOf[*map[*bool]*float32](), random.SingleOf[*map[*bool]*float64](), random.SingleOf[*map[*bool]*complex64](),
		random.SingleOf[*map[*bool]*complex128](), random.SingleOf[*map[*bool]*string](), random.SingleOf[*map[*bool]*any](),
		// int keys
		random.SingleOf[*map[*int]*bool](), random.SingleOf[*map[*int]*int](), random.SingleOf[*map[*int]*int8](), random.SingleOf[*map[*int]*int16](), random.SingleOf[*map[*int]*int32](), random.SingleOf[*map[*int]*int64](),
		random.SingleOf[*map[*int]*uint](), random.SingleOf[*map[*int]*uint8](), random.SingleOf[*map[*int]*uint16](), random.SingleOf[*map[*int]*uint32](), random.SingleOf[*map[*int]*uint64](),
		random.SingleOf[*map[*int]*uintptr](), random.SingleOf[*map[*int]*float32](), random.SingleOf[*map[*int]*float64](), random.SingleOf[*map[*int]*complex64](), random.SingleOf[*map[*int]*complex128](),
		random.SingleOf[*map[*int]*string](), random.SingleOf[*map[*int]*any](),
		// int8 keys
		random.SingleOf[*map[*int8]*bool](), random.SingleOf[*map[*int8]*int](), random.SingleOf[*map[*int8]*int8](), random.SingleOf[*map[*int8]*int16](), random.SingleOf[*map[*int8]*int32](),
		random.SingleOf[*map[*int8]*int64](), random.SingleOf[*map[*int8]*uint](), random.SingleOf[*map[*int8]*uint8](), random.SingleOf[*map[*int8]*uint16](), random.SingleOf[*map[*int8]*uint32](),
		random.SingleOf[*map[*int8]*uint64](), random.SingleOf[*map[*int8]*uintptr](), random.SingleOf[*map[*int8]*float32](), random.SingleOf[*map[*int8]*float64](), random.SingleOf[*map[*int8]*complex64](),
		random.SingleOf[*map[*int8]*complex128](), random.SingleOf[*map[*int8]*string](), random.SingleOf[*map[*int8]*any](),
		// int16 keys
		random.SingleOf[*map[*int16]*bool](), random.SingleOf[*map[*int16]*int](), random.SingleOf[*map[*int16]*int8](), random.SingleOf[*map[*int16]*int16](), random.SingleOf[*map[*int16]*int32](),
		random.SingleOf[*map[*int16]*int64](), random.SingleOf[*map[*int16]*uint](), random.SingleOf[*map[*int16]*uint8](), random.SingleOf[*map[*int16]*uint16](), random.SingleOf[*map[*int16]*uint32](),
		random.SingleOf[*map[*int16]*uint64](), random.SingleOf[*map[*int16]*uintptr](), random.SingleOf[*map[*int16]*float32](), random.SingleOf[*map[*int16]*float64](),
		random.SingleOf[*map[*int16]*complex64](), random.SingleOf[*map[*int16]*complex128](), random.SingleOf[*map[*int16]*string](), random.SingleOf[*map[*int16]*any](),
		// int32 keys
		random.SingleOf[*map[*int32]*bool](), random.SingleOf[*map[*int32]*int](), random.SingleOf[*map[*int32]*int8](), random.SingleOf[*map[*int32]*int16](), random.SingleOf[*map[*int32]*int32](),
		random.SingleOf[*map[*int32]*int64](), random.SingleOf[*map[*int32]*uint](), random.SingleOf[*map[*int32]*uint8](), random.SingleOf[*map[*int32]*uint16](), random.SingleOf[*map[*int32]*uint32](),
		random.SingleOf[*map[*int32]*uint64](), random.SingleOf[*map[*int32]*uintptr](), random.SingleOf[*map[*int32]*float32](), random.SingleOf[*map[*int32]*float64](),
		random.SingleOf[*map[*int32]*complex64](), random.SingleOf[*map[*int32]*complex128](), random.SingleOf[*map[*int32]*string](), random.SingleOf[*map[*int32]*any](),
		// int64 keys
		random.SingleOf[*map[*int64]*bool](), random.SingleOf[*map[*int64]*int](), random.SingleOf[*map[*int64]*int8](), random.SingleOf[*map[*int64]*int16](), random.SingleOf[*map[*int64]*int32](),
		random.SingleOf[*map[*int64]*int64](), random.SingleOf[*map[*int64]*uint](), random.SingleOf[*map[*int64]*uint8](), random.SingleOf[*map[*int64]*uint16](),
		random.SingleOf[*map[*int64]*uint32](), random.SingleOf[*map[*int64]*uint64](), random.SingleOf[*map[*int64]*uintptr](), random.SingleOf[*map[*int64]*float32](),
		random.SingleOf[*map[*int64]*float64](), random.SingleOf[*map[*int64]*complex64](), random.SingleOf[*map[*int64]*complex128](), random.SingleOf[*map[*int64]*string](),
		random.SingleOf[*map[*int64]*any](),
		// uint keys
		random.SingleOf[*map[*uint]*bool](), random.SingleOf[*map[*uint]*int](), random.SingleOf[*map[*uint]*int8](), random.SingleOf[*map[*uint]*int16](), random.SingleOf[*map[*uint]*int32](),
		random.SingleOf[*map[*uint]*int64](), random.SingleOf[*map[*uint]*uint](), random.SingleOf[*map[*uint]*uint8](), random.SingleOf[*map[*uint]*uint16](), random.SingleOf[*map[*uint]*uint32](),
		random.SingleOf[*map[*uint]*uint64](), random.SingleOf[*map[*uint]*uintptr](), random.SingleOf[*map[*uint]*float32](), random.SingleOf[*map[*uint]*float64](),
		random.SingleOf[*map[*uint]*complex64](), random.SingleOf[*map[*uint]*complex128](), random.SingleOf[*map[*uint]*string](), random.SingleOf[*map[*uint]*any](),
		// uint8 keys
		random.SingleOf[*map[*uint8]*bool](), random.SingleOf[*map[*uint8]*int](), random.SingleOf[*map[*uint8]*int8](), random.SingleOf[*map[*uint8]*int16](), random.SingleOf[*map[*uint8]*int32](),
		random.SingleOf[*map[*uint8]*int64](), random.SingleOf[*map[*uint8]*uint](), random.SingleOf[*map[*uint8]*uint8](), random.SingleOf[*map[*uint8]*uint16](), random.SingleOf[*map[*uint8]*uint32](),
		random.SingleOf[*map[*uint8]*uint64](), random.SingleOf[*map[*uint8]*uintptr](), random.SingleOf[*map[*uint8]*float32](), random.SingleOf[*map[*uint8]*float64](),
		random.SingleOf[*map[*uint8]*complex64](), random.SingleOf[*map[*uint8]*complex128](), random.SingleOf[*map[*uint8]*string](), random.SingleOf[*map[*uint8]*any](),
		// uint16 keys
		random.SingleOf[*map[*uint16]*bool](), random.SingleOf[*map[*uint16]*int](), random.SingleOf[*map[*uint16]*int8](), random.SingleOf[*map[*uint16]*int16](), random.SingleOf[*map[*uint16]*int32](),
		random.SingleOf[*map[*uint16]*int64](), random.SingleOf[*map[*uint16]*uint](), random.SingleOf[*map[*uint16]*uint8](), random.SingleOf[*map[*uint16]*uint16](),
		random.SingleOf[*map[*uint16]*uint32](), random.SingleOf[*map[*uint16]*uint64](), random.SingleOf[*map[*uint16]*uintptr](), random.SingleOf[*map[*uint16]*float32](),
		random.SingleOf[*map[*uint16]*float64](), random.SingleOf[*map[*uint16]*complex64](), random.SingleOf[*map[*uint16]*complex128](), random.SingleOf[*map[*uint16]*string](),
		random.SingleOf[*map[*uint16]*any](),
		// uint32 keys
		random.SingleOf[*map[*uint32]*bool](), random.SingleOf[*map[*uint32]*int](), random.SingleOf[*map[*uint32]*int8](), random.SingleOf[*map[*uint32]*int16](), random.SingleOf[*map[*uint32]*int32](),
		random.SingleOf[*map[*uint32]*int64](), random.SingleOf[*map[*uint32]*uint](), random.SingleOf[*map[*uint32]*uint8](), random.SingleOf[*map[*uint32]*uint16](),
		random.SingleOf[*map[*uint32]*uint32](), random.SingleOf[*map[*uint32]*uint64](), random.SingleOf[*map[*uint32]*uintptr](), random.SingleOf[*map[*uint32]*float32](), random.SingleOf[*map[*uint32]*float64](),
		random.SingleOf[*map[*uint32]*complex64](), random.SingleOf[*map[*uint32]*complex128](), random.SingleOf[*map[*uint32]*string](), random.SingleOf[*map[*uint32]*any](),
		// uint64 keys
		random.SingleOf[*map[*uint64]*bool](), random.SingleOf[*map[*uint64]*int](), random.SingleOf[*map[*uint64]*int8](), random.SingleOf[*map[*uint64]*int16](), random.SingleOf[*map[*uint64]*int32](),
		random.SingleOf[*map[*uint64]*int64](), random.SingleOf[*map[*uint64]*uint](), random.SingleOf[*map[*uint64]*uint8](), random.SingleOf[*map[*uint64]*uint16](), random.SingleOf[*map[*uint64]*uint32](),
		random.SingleOf[*map[*uint64]*uint64](), random.SingleOf[*map[*uint64]*uintptr](), random.SingleOf[*map[*uint64]*float32](), random.SingleOf[*map[*uint64]*float64](),
		random.SingleOf[*map[*uint64]*complex64](), random.SingleOf[*map[*uint64]*complex128](), random.SingleOf[*map[*uint64]*string](), random.SingleOf[*map[*uint64]*any](),
		// uintptr keys
		random.SingleOf[*map[*uintptr]*bool](), random.SingleOf[*map[*uintptr]*int](), random.SingleOf[*map[*uintptr]*int8](), random.SingleOf[*map[*uintptr]*int16](), random.SingleOf[*map[*uintptr]*int32](),
		random.SingleOf[*map[*uintptr]*int64](), random.SingleOf[*map[*uintptr]*uint](), random.SingleOf[*map[*uintptr]*uint8](), random.SingleOf[*map[*uintptr]*uint16](),
		random.SingleOf[*map[*uintptr]*uint32](), random.SingleOf[*map[*uintptr]*uint64](), random.SingleOf[*map[*uintptr]*uintptr](), random.SingleOf[*map[*uintptr]*float32](),
		random.SingleOf[*map[*uintptr]*float64](), random.SingleOf[*map[*uintptr]*complex64](), random.SingleOf[*map[*uintptr]*complex128](), random.SingleOf[*map[*uintptr]*string](),
		random.SingleOf[*map[*uintptr]*any](),
		// float32 keys
		random.SingleOf[*map[*float32]*bool](), random.SingleOf[*map[*float32]*int](), random.SingleOf[*map[*float32]*int8](), random.SingleOf[*map[*float32]*int16](),
		random.SingleOf[*map[*float32]*int32](), random.SingleOf[*map[*float32]*int64](), random.SingleOf[*map[*float32]*uint](), random.SingleOf[*map[*float32]*uint8](),
		random.SingleOf[*map[*float32]*uint16](), random.SingleOf[*map[*float32]*uint32](), random.SingleOf[*map[*float32]*uint64](), random.SingleOf[*map[*float32]*uintptr](),
		random.SingleOf[*map[*float32]*float32](), random.SingleOf[*map[*float32]*float64](), random.SingleOf[*map[*float32]*complex64](), random.SingleOf[*map[*float32]*complex128](),
		random.SingleOf[*map[*float32]*string](), random.SingleOf[*map[*float32]*any](),
		// float64 keys
		random.SingleOf[*map[*float64]*bool](), random.SingleOf[*map[*float64]*int](), random.SingleOf[*map[*float64]*int8](), random.SingleOf[*map[*float64]*int16](),
		random.SingleOf[*map[*float64]*int32](), random.SingleOf[*map[*float64]*int64](), random.SingleOf[*map[*float64]*uint](), random.SingleOf[*map[*float64]*uint8](),
		random.SingleOf[*map[*float64]*uint16](), random.SingleOf[*map[*float64]*uint32](), random.SingleOf[*map[*float64]*uint64](), random.SingleOf[*map[*float64]*uintptr](),
		random.SingleOf[*map[*float64]*float32](), random.SingleOf[*map[*float64]*float64](), random.SingleOf[*map[*float64]*complex64](), random.SingleOf[*map[*float64]*complex128](),
		random.SingleOf[*map[*float64]*string](), random.SingleOf[*map[*float64]*any](),
		// complex64 keys
		random.SingleOf[*map[*complex64]*bool](), random.SingleOf[*map[*complex64]*int](), random.SingleOf[*map[*complex64]*int8](), random.SingleOf[*map[*complex64]*int16](),
		random.SingleOf[*map[*complex64]*int32](), random.SingleOf[*map[*complex64]*int64](), random.SingleOf[*map[*complex64]*uint](), random.SingleOf[*map[*complex64]*uint8](),
		random.SingleOf[*map[*complex64]*uint16](), random.SingleOf[*map[*complex64]*uint32](), random.SingleOf[*map[*complex64]*uint64](), random.SingleOf[*map[*complex64]*uintptr](),
		random.SingleOf[*map[*complex64]*float32](), random.SingleOf[*map[*complex64]*float64](), random.SingleOf[*map[*complex64]*complex64](), random.SingleOf[*map[*complex64]*complex128](),
		random.SingleOf[*map[*complex64]*string](), random.SingleOf[*map[*complex64]*any](),
		// complex128 keys
		random.SingleOf[*map[*complex128]*bool](), random.SingleOf[*map[*complex128]*int](), random.SingleOf[*map[*complex128]*int8](), random.SingleOf[*map[*complex128]*int16](),
		random.SingleOf[*map[*complex128]*int32](), random.SingleOf[*map[*complex128]*int64](), random.SingleOf[*map[*complex128]*uint](), random.SingleOf[*map[*complex128]*uint8](),
		random.SingleOf[*map[*complex128]*uint16](), random.SingleOf[*map[*complex128]*uint32](), random.SingleOf[*map[*complex128]*uint64](), random.SingleOf[*map[*complex128]*uintptr](),
		random.SingleOf[*map[*complex128]*float32](), random.SingleOf[*map[*complex128]*float64](), random.SingleOf[*map[*complex128]*complex64](),
		random.SingleOf[*map[*complex128]*complex128](), random.SingleOf[*map[*complex128]*string](), random.SingleOf[*map[*complex128]*any](),
		// string keys
		random.SingleOf[*map[*string]*bool](), random.SingleOf[*map[*string]*int](), random.SingleOf[*map[*string]*int8](), random.SingleOf[*map[*string]*int16](),
		random.SingleOf[*map[*string]*int32](), random.SingleOf[*map[*string]*int64](), random.SingleOf[*map[*string]*uint](), random.SingleOf[*map[*string]*uint8](),
		random.SingleOf[*map[*string]*uint16](), random.SingleOf[*map[*string]*uint32](), random.SingleOf[*map[*string]*uint64](), random.SingleOf[*map[*string]*uintptr](),
		random.SingleOf[*map[*string]*float32](), random.SingleOf[*map[*string]*float64](), random.SingleOf[*map[*string]*complex64](), random.SingleOf[*map[*string]*complex128](),
		random.SingleOf[*map[*string]*string](), random.SingleOf[*map[*string]*any](),
		// any keys
		random.SingleOf[*map[*any]*bool](), random.SingleOf[*map[*any]*int](), random.SingleOf[*map[*any]*int8](), random.SingleOf[*map[*any]*int16](), random.SingleOf[*map[*any]*int32](),
		random.SingleOf[*map[*any]*int64](), random.SingleOf[*map[*any]*uint](), random.SingleOf[*map[*any]*uint8](), random.SingleOf[*map[*any]*uint16](), random.SingleOf[*map[*any]*uint32](),
		random.SingleOf[*map[*any]*uint64](), random.SingleOf[*map[*any]*uintptr](), random.SingleOf[*map[*any]*float32](), random.SingleOf[*map[*any]*float64](), random.SingleOf[*map[*any]*complex64](),
		random.SingleOf[*map[*any]*complex128](), random.SingleOf[*map[*any]*string](), random.SingleOf[*map[*any]*any](),
	}
}

func edgeRefs() []interface{} {
	return []interface{}{
		random.SingleOf[*map[unsafe.Pointer]unsafe.Pointer](), random.SingleOf[*map[*unsafe.Pointer]unsafe.Pointer](),
		random.SingleOf[*map[unsafe.Pointer]*unsafe.Pointer](), random.SingleOf[*map[*unsafe.Pointer]*unsafe.Pointer](),
		random.SingleOf[*map[**string]string](), random.SingleOf[*map[string]**string](), random.SingleOf[*map[*unsafe.Pointer]map[string]string](),
	}
}
