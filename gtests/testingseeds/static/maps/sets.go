// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static"

func SetOf() static.SeedSets[Categories, Categories] {
	return &maps{}
}

type maps struct{}

func (m maps) Values() Categories {
	return values{}
}

func (m maps) Refs() Categories {
	return refs{}
}

type Categories interface {
	All() []any
	Primitives() []any
	PrimitivesWithRefKey() []any
	PrimitivesWithRefValue() []any
	PrimitiveWithRefKeyAndValue() []any
	Edges() []any
}

type values struct{}

func (v values) All() []any {
	return ofValues()
}

func (v values) Primitives() []any {
	return primitives()
}

func (v values) PrimitivesWithRefKey() []any {
	return primitivesPtrKeys()
}

func (v values) PrimitivesWithRefValue() []any {
	return primitivesPtrValues()
}

func (v values) PrimitiveWithRefKeyAndValue() []any {
	return primitivesPtrKeysAndValues()
}

func (v values) Edges() []any {
	return edgeValues()
}

type refs struct{}

func (r refs) All() []any {
	return ofRefs()
}

func (r refs) Primitives() []any {
	return primitivesAsRef()
}

func (r refs) PrimitivesWithRefKey() []any {
	return primitiveWithPtrKeysAndValuesAsRef()
}

func (r refs) PrimitivesWithRefValue() []any {
	return primitiveWithPtrValuesAsRef()
}

func (r refs) PrimitiveWithRefKeyAndValue() []any {
	return primitiveWithPtrKeysAsRef()
}

func (r refs) Edges() []any {
	return edgeAsRefs()
}
