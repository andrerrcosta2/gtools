// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

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
	return iterables.OfSlice(referencesSet()...)
}

func (c CatRefs) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveRefs()...)
}

func (c CatRefs) pointerToPrimitive() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToPrimitiveValues()...)
}

func (c CatRefs) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceRefs()...)
}

func (c CatRefs) pointerToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToInterfaceRefs()...)
}

func (c CatRefs) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structRefs()...)
}

func (c CatRefs) pointerToStructs() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToStructRefs()...)
}

func (c CatRefs) Channels() *iterables.Slice[any] {
	return iterables.OfSlice(channelRefs()...)
}

func (c CatRefs) pointerToChannels() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToChannelRefs()...)
}

func (c CatRefs) EdgeTypes() *iterables.Slice[any] {
	return iterables.OfSlice(EdgeTypesRefs()...)
}

func (c CatValues) All() *iterables.Slice[any] {
	return iterables.OfSlice(valuesSet()...)
}

func (c CatValues) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveValues()...)
}

func (c CatValues) PointerToPrimitive() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToPrimitiveValues()...)
}

func (c CatValues) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceValues()...)
}

func (c CatValues) PointerToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(pointersToInterfaceValues()...)
}

func (c CatValues) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structValues()...)
}

func (c CatValues) PointerToStructs() *iterables.Slice[any] {
	return iterables.OfSlice(pointersToStructValues()...)
}

func (c CatValues) Channels() *iterables.Slice[any] {
	return iterables.OfSlice(channelValues()...)
}

func (c CatValues) PointerToChannels() *iterables.Slice[any] {
	return iterables.OfSlice(pointersToChannelValues()...)
}
