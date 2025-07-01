// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

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

func (c CatRefs) All() []any {
	return referencesSet()
}

func (c CatRefs) Primitives() []any {
	return primitiveRefs()
}

func (c CatRefs) pointerToPrimitive() []any {
	return pointerToPrimitiveValues()
}

func (c CatRefs) Interfaces() []any {
	return interfaceRefs()
}

func (c CatRefs) pointerToInterfaces() []any {
	return pointerToInterfaceRefs()
}

func (c CatRefs) Structs() []any {
	return structRefs()
}

func (c CatRefs) pointerToStructs() []any {
	return pointerToStructRefs()
}

func (c CatRefs) Channels() []any {
	return channelRefs()
}

func (c CatRefs) pointerToChannels() []any {
	return pointerToChannelRefs()
}

func (c CatRefs) EdgeTypes() []any {
	return EdgeTypesRefs()
}

func (c CatValues) All() []any {
	return valuesSet()
}

func (c CatValues) Primitives() []any {
	return primitiveValues()
}

func (c CatValues) PointerToPrimitive() []any {
	return pointerToPrimitiveValues()
}

func (c CatValues) Interfaces() []any {
	return interfaceValues()
}

func (c CatValues) PointerToInterfaces() []any {
	return pointersToInterfaceValues()
}

func (c CatValues) Structs() []any {
	return structValues()
}

func (c CatValues) PointerToStructs() []any {
	return pointersToStructValues()
}

func (c CatValues) Channels() []any {
	return channelValues()
}

func (c CatValues) PointerToChannels() []any {
	return pointersToChannelValues()
}
