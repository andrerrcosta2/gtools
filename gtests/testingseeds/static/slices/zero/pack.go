// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import "github.com/andrerrcosta2/gtools/core/data/str/iterables"

type (
	Pack      struct{}
	CatRefs   struct{}
	CatValues struct{}
)

func (p Pack) Refs() CatRefs {
	return CatRefs{}
}

func (p Pack) Values() CatValues {
	return CatValues{}
}

func (r CatRefs) All() *iterables.Slice[any] {
	return iterables.OfSlice(referencesSet()...)
}

func (r CatRefs) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveRefs()...)
}

func (r CatRefs) NilPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(nilPrimitiveRefs()...)
}

func (r CatRefs) NilPointerToPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(nilPointerToPrimitiveRefs()...)
}

func (r CatRefs) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceRefs()...)
}

func (r CatRefs) PointerToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToInterfaceRefs()...)
}

func (r CatRefs) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structRefs()...)
}

func (r CatRefs) PointerToStructs() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToStructRefs()...)
}

func (r CatRefs) Channels() *iterables.Slice[any] {
	return iterables.OfSlice(channelRefs()...)
}

func (r CatRefs) PointerToChannels() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToChannelRefs()...)
}

func (r CatRefs) Edges() *iterables.Slice[any] {
	return iterables.OfSlice(edgeRefs()...)
}

func (v CatValues) All() *iterables.Slice[any] {
	return iterables.OfSlice(valuesSet()...)
}

func (v CatValues) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveValues()...)
}

func (v CatValues) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceValues()...)
}

func (v CatValues) PointerToInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToInterfaceValues()...)
}

func (v CatValues) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structValues()...)
}

func (v CatValues) PointerToStructs() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToStructValues()...)
}

func (v CatValues) Channels() *iterables.Slice[any] {
	return iterables.OfSlice(channelValues()...)
}

func (v CatValues) PointerToChannels() *iterables.Slice[any] {
	return iterables.OfSlice(pointerToChannelValues()...)
}

func (v CatValues) Edges() *iterables.Slice[any] {
	return iterables.OfSlice(edgeValues()...)
}
