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
	return iterables.OfSlice(referencesSets()...)
}

func (c CatRefs) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveRefs()...)
}

func (c CatRefs) PointerToPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToPrimitiveValues()...)
}

func (c CatRefs) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceRefs()...)
}

func (c CatRefs) PointerToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToInterfaceRefs()...)
}

func (c CatRefs) SliceOfInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfInterfaceRefs()...)
}

func (c CatRefs) SliceOfPointersToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPointerToInterfaceRefs()...)
}

func (c CatRefs) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structRefs()...)
}

func (c CatRefs) SliceOfStructs() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfStructRefs()...)
}

func (c CatRefs) SliceOfPointersToStructs() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToSliceOfStructRefs()...)
}

func (c CatRefs) SliceOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPrimitiveRefs()...)
}

func (c CatRefs) SliceOfPointersToPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPointerToPrimitiveRefs()...)
}

func (c CatRefs) Senders() *iterables.Slice[any] {
	return iterables.OfSlice(senderRefs()...)
}

func (c CatRefs) Receivers() *iterables.Slice[any] {
	return iterables.OfSlice(receiverRefs()...)
}

func (c CatRefs) Edges() *iterables.Slice[any] {
	return iterables.OfSlice(edgeRefs()...)
}

func (c CatValues) All() *iterables.Slice[any] {
	return iterables.OfSlice(valuesSet()...)
}

func (c CatValues) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveValues()...)
}

func (c CatValues) PointerToPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToPrimitiveValues()...)
}

func (c CatValues) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceValues()...)
}

func (c CatValues) PointerToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToInterfaceValues()...)
}

func (c CatValues) SliceOfInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfInterfaceValues()...)
}

func (c CatValues) SliceOfPointersToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPointerToInterfaceValues()...)
}

func (c CatValues) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structValues()...)
}

func (c CatValues) SliceOfStructs() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfStructValues()...)
}

func (c CatValues) SliceOfPointersToStructs() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToSliceOfStructValues()...)
}

func (c CatValues) SliceOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPrimitiveValues()...)
}

func (c CatValues) SliceOfPointersToPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPointerToPrimitiveValues()...)
}

func (c CatValues) Senders() *iterables.Slice[any] {
	return iterables.OfSlice(senderValues()...)
}

func (c CatValues) Receivers() *iterables.Slice[any] {
	return iterables.OfSlice(receiverValues()...)
}

func (c CatValues) Edges() *iterables.Slice[any] {
	return iterables.OfSlice(edgeValues()...)
}
