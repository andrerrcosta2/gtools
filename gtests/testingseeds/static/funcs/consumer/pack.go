// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package consumer

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

func (c CatRefs) All() []any {
	return referencesSet()
}

func (c CatRefs) Primitives() []any {
	return primitiveRefs()
}

func (c CatRefs) VariadicPrimitive() []any {
	return variadicPrimitiveRefs()
}

func (c CatRefs) Interfaces() []any {
	return interfaceRefs()
}

func (c CatRefs) VariadicInterfaces() []any {
	return variadicInterfaceRefs()
}

func (c CatRefs) Structs() []any {
	return structRefs()
}

func (c CatRefs) VariadicStructs() []any {
	return variadicStructRefs()
}

func (c CatRefs) ChannelOfPrimitives() []any {
	return channelOfPrimitiveRefs()
}

func (c CatRefs) VariadicChannelOfPrimitives() []any {
	return variadicChannelOfPrimitiveRefs()
}

func (c CatRefs) MapOfPrimitives() []any {
	return mapOfPrimitiveRefs()
}

func (c CatRefs) VariadicMapOfPrimitives() []any {
	return variadicMapOfPrimitiveRefs()
}

func (c CatRefs) SliceOfPrimitives() []any {
	return sliceOfPrimitiveRefs()
}

func (c CatRefs) VariadicSliceOfPrimitives() []any {
	return variadicSliceOfPrimitiveRefs()
}

func (c CatRefs) Edges() []any {
	return edgeRefs()
}

func (c CatValues) All() []any {
	return valuesSet()
}

func (c CatValues) Primitives() []any {
	return primitiveValues()
}

func (c CatValues) VariadicPrimitives() []any {
	return variadicPrimitiveValues()
}

func (c CatValues) Interfaces() []any {
	return interfaceValues()
}

func (c CatValues) VariadicInterfaces() []any {
	return variadicInterfaceValues()
}

func (c CatValues) Structs() []any {
	return structValues()
}

func (c CatValues) VariadicStructs() []any {
	return variadicStructValues()
}

func (c CatValues) ChannelOfPrimitives() []any {
	return chanOfPrimitiveValues()
}

func (c CatValues) VariadicChannelOfPrimitives() []any {
	return variadicChanOfPrimitiveValues()
}

func (c CatValues) MapOfPrimitives() []any {
	return mapOfPrimitiveValues()
}

func (c CatValues) VariadicMapOfPrimitives() []any {
	return variadicMapOfPrimitiveValues()
}

func (c CatValues) SliceOfPrimitives() []any {
	return sliceOfPrimitiveValues()
}

func (c CatValues) VariadicSliceOfPrimitives() []any {
	return variadicSliceOfPrimitiveValues()
}

func (c CatValues) Edges() []any {
	return edgeValues()
}
