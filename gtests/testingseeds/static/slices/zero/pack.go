// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

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

func (r CatRefs) All() []any {
	return referencesSet()
}

func (r CatRefs) Primitives() []any {
	return primitiveRefs()
}

func (r CatRefs) NilPrimitives() []any {
	return nilPrimitiveRefs()
}

func (r CatRefs) NilPointerToPrimitives() []any {
	return nilPointerToPrimitiveRefs()
}

func (r CatRefs) Interfaces() []any {
	return interfaceRefs()
}

func (r CatRefs) PointerToInterfaces() []any {
	return pointerToInterfaceRefs()
}

func (r CatRefs) Structs() []any {
	return structRefs()
}

func (r CatRefs) PointerToStructs() []any {
	return pointerToStructRefs()
}

func (r CatRefs) Channels() []any {
	return channelRefs()
}

func (r CatRefs) PointerToChannels() []any {
	return pointerToChannelRefs()
}

func (r CatRefs) Edges() []any {
	return edgeRefs()
}

func (v CatValues) All() []any {
	return valuesSet()
}

func (v CatValues) Primitives() []any {
	return primitiveValues()
}

func (v CatValues) Interfaces() []any {
	return interfaceValues()
}

func (v CatValues) PointerToInterfaces() []any {
	return pointerToInterfaceValues()
}

func (v CatValues) Structs() []any {
	return structValues()
}

func (v CatValues) PointerToStructs() []any {
	return pointerToStructValues()
}

func (v CatValues) Channels() []any {
	return channelValues()
}

func (v CatValues) PointerToChannels() []any {
	return pointerToChannelValues()
}

func (v CatValues) Edges() []any {
	return edgeValues()
}
