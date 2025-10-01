// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package supplier

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

func (c CatRefs) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceRefs()...)
}

func (c CatRefs) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structRefs()...)
}

func (c CatRefs) ChannelsOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(chanOfPrimitiveRefs()...)
}

func (c CatRefs) MapOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(mapOfPrimitiveRefs()...)
}

func (c CatRefs) SliceOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPrimitiveRefs()...)
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

func (c CatValues) Interfaces() *iterables.Slice[any] {
	return iterables.OfSlice(interfaceValues()...)
}

func (c CatValues) Structs() *iterables.Slice[any] {
	return iterables.OfSlice(structValues()...)
}

func (c CatValues) ChannelsOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(chanOfPrimitiveValues()...)
}

func (c CatValues) MapOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(mapOfPrimitiveValues()...)
}

func (c CatValues) SliceOfPrimitives() *iterables.Slice[any] {
	return iterables.OfSlice(sliceOfPrimitiveValues()...)
}

func (c CatValues) Edges() *iterables.Slice[any] {
	return iterables.OfSlice(edgeValues()...)
}
