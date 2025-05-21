// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interf

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static"

func SetOf() static.SeedSets[Categories, Categories] {
	return interfs{}
}

type interfs struct{}

func (i interfs) Values() Categories {
	return values{}
}

func (i interfs) Refs() Categories {
	return refs{}
}

type Categories interface {
	All() []any
	Impls() []any
	Edges() []any
}

type values struct{}

func (v values) All() []any {
	return ofValues()
}

func (v values) Impls() []any {
	return implementables()
}

func (v values) Edges() []any {
	return edgeValues()
}

type refs struct{}

func (r refs) All() []any {
	return ofRefs()
}

func (r refs) Impls() []any {
	return implementablesOfRefs()
}

func (r refs) Edges() []any {
	return edgesAsRefs()
}
