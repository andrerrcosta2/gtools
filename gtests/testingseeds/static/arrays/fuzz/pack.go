// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import "github.com/andrerrcosta2/gtools/core/data/str/iterables"

type (
	Pack      struct{}
	CatRefs   struct{}
	CatValues struct{}
)

func (c Pack) Refs() CatRefs {
	return CatRefs{}
}

func (c Pack) Values() CatValues {
	return CatValues{}
}

func (c CatRefs) All() *iterables.Slice[any] {
	return iterables.OfSlice[any](referencesSet()...)
}

func (c CatRefs) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice[any](primitiveRefs()...)
}

func (c CatRefs) pointerToPrimitive() *iterables.Slice[any] {
	return iterables.OfSlice[any](pointerToPrimitiveValues()...)
}

func (c CatRefs) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice[any](interfaceRefs()...)
}

func (c CatRefs) pointerToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice[any](pointerToInterfaceRefs()...)
}

func (c CatRefs) Structs() *iterables.Slice[any] {
	return iterables.OfSlice[any](structRefs()...)
}

func (c CatRefs) pointerToStructs() *iterables.Slice[any] {
	return iterables.OfSlice[any](pointerToStructRefs()...)
}

func (c CatRefs) Channels() *iterables.Slice[any] {
	return iterables.OfSlice[any](channelRefs()...)
}

func (c CatRefs) pointerToChannels() *iterables.Slice[any] {
	return iterables.OfSlice[any](pointerToChannelRefs()...)
}

func (c CatRefs) EdgeTypes() *iterables.Slice[any] {
	return iterables.OfSlice[any](EdgeTypesRefs()...)
}

func (c CatValues) All() *iterables.Slice[any] {
	return iterables.OfSlice[any](valuesSet()...)
}

func (c CatValues) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice[any](primitiveValues()...)
}

func (c CatValues) PointerToPrimitive() *iterables.Slice[any] {
	return iterables.OfSlice[any](pointerToPrimitiveValues()...)
}

func (c CatValues) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice[any](interfaceValues()...)
}

func (c CatValues) PointerToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice[any](pointerToInterfaceValues()...)
}

func (c CatValues) Structs() *iterables.Slice[any] {
	return iterables.OfSlice[any](structValues()...)
}

func (c CatValues) PointerToStructs() *iterables.Slice[any] {
	return iterables.OfSlice[any](pointerToStructValues()...)
}

func (c CatValues) Channels() *iterables.Slice[any] {
	return iterables.OfSlice[any](channelValues()...)
}

func (c CatValues) PointerToChannels() *iterables.Slice[any] {
	return iterables.OfSlice[any](pointerToChannelValues()...)
}
