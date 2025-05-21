// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package funcs

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"slices"
)

func SetOf() static.SeedSets[Categories, Categories] {
	return funcs{}
}

type funcs struct{}

func (f funcs) Values() Categories {
	return values{}
}

func (f funcs) Refs() Categories {
	return refs{}
}

type Categories interface {
	All() []any
	Consumers() Consumers

	Suppliers() Suppliers
}

type Consumers interface {
	All() []any
	Primitives() []any
	PrimitivesVariad() []any
	Interfaces() []any
	InterfacesVariad() []any
	Structs() []any
	StructsVariad() []any
	Channels() []any
	ChannelsVariad() []any
	Maps() []any
	MapsVariad() []any
	SlicesOfPrimitives() []any
	SlicesOfPrimitivesVariad() []any
	Edges() []any
}

type Suppliers interface {
	All() []any
	Primitives() []any
	Interfaces() []any
	Structs() []any
	Channels() []any
	Maps() []any
	SlicesOfPrimitives() []any
	Edges() []any
}

type values struct{}

func (v values) All() []any {
	return slices.Concat(suppliers(), consumers())
}

func (v values) Consumers() Consumers {
	return consumerValues{}
}

func (v values) Suppliers() Suppliers {
	return supplierValues{}
}

type refs struct{}

func (r refs) All() []any {
	return slices.Concat(suppliersAsRef(), consumersAsRef())
}

func (r refs) Consumers() Consumers {
	return consumerRefs{}
}

func (r refs) Suppliers() Suppliers {
	return supplierRefs{}
}

type consumerValues struct{}

func (c consumerValues) All() []any {
	return consumers()
}

func (c consumerValues) Primitives() []any {
	return consumerOfPrimitives()
}

func (c consumerValues) PrimitivesVariad() []any {
	return consumerOfVariadicPrimitives()
}

func (c consumerValues) Interfaces() []any {
	return consumerOfInterfaces()
}

func (c consumerValues) InterfacesVariad() []any {
	return consumerOfVariadicInterfaces()
}

func (c consumerValues) Structs() []any {
	return consumerStructs()
}

func (c consumerValues) StructsVariad() []any {
	return consumerVariadicStruct()
}

func (c consumerValues) Channels() []any {
	return consumerChanOfPrimitives()
}

func (c consumerValues) ChannelsVariad() []any {
	return consumerVariadicChanOfPrimitives()
}

func (c consumerValues) Maps() []any {
	return consumerMapPrimitives()
}

func (c consumerValues) MapsVariad() []any {
	return consumerVariadicMapOsPrimitives()
}

func (c consumerValues) SlicesOfPrimitives() []any {
	return consumerSliceOfPrimitives()
}

func (c consumerValues) SlicesOfPrimitivesVariad() []any {
	return consumerVariadicSliceOfPrimitives()
}

func (c consumerValues) Edges() []any {
	return consumerEdgeValues()
}

type supplierValues struct{}

func (s supplierValues) All() []any {
	return suppliers()
}

func (s supplierValues) Primitives() []any {
	return supplierOfPrimitives()
}

func (s supplierValues) Interfaces() []any {
	return supplierOfInterfaces()
}

func (s supplierValues) Structs() []any {
	return supplierStructs()
}

func (s supplierValues) Channels() []any {
	return supplierChanOfPrimitives()
}

func (s supplierValues) Maps() []any {
	return supplierMapPrimitives()
}

func (s supplierValues) SlicesOfPrimitives() []any {
	return supplierSliceOfPrimitives()
}

func (s supplierValues) Edges() []any {
	return supplierEdgeValues()
}

type consumerRefs struct{}

func (c consumerRefs) All() []any {
	return consumersAsRef()
}

func (c consumerRefs) Primitives() []any {
	return consumerPrimitiveAsRef()
}

func (c consumerRefs) PrimitivesVariad() []any {
	return consumerVariadicPrimitiveAsRef()
}

func (c consumerRefs) Interfaces() []any {
	return consumerInterfaceAsRef()
}

func (c consumerRefs) InterfacesVariad() []any {
	return consumerVariadicInterfaceAsRef()
}

func (c consumerRefs) Structs() []any {
	return consumerStructAsRef()
}

func (c consumerRefs) StructsVariad() []any {
	return consumerVariadicStructAsRef()
}

func (c consumerRefs) Channels() []any {
	return consumerChannelOfPrimitiveAsRef()
}

func (c consumerRefs) ChannelsVariad() []any {
	return consumerVariadicChannelOfPrimitiveAsRef()
}

func (c consumerRefs) Maps() []any {
	return consumerMapPrimitiveAsRef()
}

func (c consumerRefs) MapsVariad() []any {
	return consumerVariadicMapPrimitiveAsRef()
}

func (c consumerRefs) SlicesOfPrimitives() []any {
	return consumerSliceOfPrimitiveAsRef()
}

func (c consumerRefs) SlicesOfPrimitivesVariad() []any {
	return consumerVariadicSliceOfPrimitiveAsRef()
}

func (c consumerRefs) Edges() []any {
	return consumerEdgesAsRef()
}

type supplierRefs struct{}

func (s supplierRefs) All() []any {
	return suppliersAsRef()
}

func (s supplierRefs) Primitives() []any {
	return supplierOfPrimitivesAsRef()
}

func (s supplierRefs) Interfaces() []any {
	return supplierOfInterfacesAsRef()
}

func (s supplierRefs) Structs() []any {
	return supplierStructsAsRef()
}

func (s supplierRefs) Channels() []any {
	return supplierChanOfPrimitivesAsRef()
}

func (s supplierRefs) Maps() []any {
	return supplierMapPrimitivesAsRef()
}

func (s supplierRefs) SlicesOfPrimitives() []any {
	return supplierSliceOfPrimitivesAsRef()
}

func (s supplierRefs) Edges() []any {
	return supplierEdgeRefs()
}
