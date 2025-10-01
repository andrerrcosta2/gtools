// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package consumer

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

func (c CatRefs) All() *iterables.Slice[any] {
	return iterables.OfSlice(referencesSet()...)
}

func (c CatRefs) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveRefs()...)
}

func (c CatRefs) VariadicPrimitive() *iterables.Slice[any] {
	return iterables.OfSlice(variadicPrimitiveRefs()...)
}

func (c CatRefs) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceRefs()...)
}

func (c CatRefs) VariadicInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(variadicInterfaceRefs()...)
}

func (c CatRefs) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structRefs()...)
}

func (c CatRefs) VariadicStructs() *iterables.Slice[any] {
	return iterables.OfSlice(variadicStructRefs()...)
}

func (c CatRefs) ChannelOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(channelOfPrimitiveRefs()...)
}

func (c CatRefs) VariadicChannelOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(variadicChannelOfPrimitiveRefs()...)
}

func (c CatRefs) MapOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(mapOfPrimitiveRefs()...)
}

func (c CatRefs) VariadicMapOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(variadicMapOfPrimitiveRefs()...)
}

func (c CatRefs) SliceOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPrimitiveRefs()...)
}

func (c CatRefs) VariadicSliceOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(variadicSliceOfPrimitiveRefs()...)
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

func (c CatValues) VariadicPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(variadicPrimitiveValues()...)
}

func (c CatValues) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceValues()...)
}

func (c CatValues) VariadicInterfaces() *iterables.Slice[any] {
	return iterables.OfSlice(variadicInterfaceValues()...)
}

func (c CatValues) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structValues()...)
}

func (c CatValues) VariadicStructs() *iterables.Slice[any] {
	return iterables.OfSlice(variadicStructValues()...)
}

func (c CatValues) ChannelOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(chanOfPrimitiveValues()...)
}

func (c CatValues) VariadicChannelOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(variadicChanOfPrimitiveValues()...)
}

func (c CatValues) MapOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(mapOfPrimitiveValues()...)
}

func (c CatValues) VariadicMapOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(variadicMapOfPrimitiveValues()...)
}

func (c CatValues) SliceOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPrimitiveValues()...)
}

func (c CatValues) VariadicSliceOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(variadicSliceOfPrimitiveValues()...)
}

func (c CatValues) Edges() *iterables.Slice[any] {
	return iterables.OfSlice(edgeValues()...)
}
