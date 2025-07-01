// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package supplier

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

func (c CatRefs) Interfaces() []any {
	return interfaceRefs()
}

func (c CatRefs) Structs() []any {
	return structRefs()
}

func (c CatRefs) ChannelsOfPrimitives() []any {
	return chanOfPrimitiveRefs()
}

func (c CatRefs) MapOfPrimitives() []any {
	return mapOfPrimitiveRefs()
}

func (c CatRefs) SliceOfPrimitives() []any {
	return sliceOfPrimitiveRefs()
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

func (c CatValues) Interfaces() []any {
	return interfaceValues()
}

func (c CatValues) Structs() []any {
	return structValues()
}

func (c CatValues) ChannelsOfPrimitives() []any {
	return chanOfPrimitiveValues()
}

func (c CatValues) MapOfPrimitives() []any {
	return mapOfPrimitiveValues()
}

func (c CatValues) SliceOfPrimitives() []any {
	return sliceOfPrimitiveValues()
}

func (c CatValues) Edges() []any {
	return edgeValues()
}
