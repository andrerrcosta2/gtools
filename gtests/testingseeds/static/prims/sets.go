// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prims

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static"

func SetOf() static.SeedSets[Categories, Categories] {
	return &prims{}
}

type prims struct{}

func (p *prims) Values() Categories {
	return values{}
}

func (p *prims) Refs() Categories {
	return refs{}
}

type Categories interface {
	All() []any
	Zero() []any
	Rand() []any
}

type values struct{}

func (v values) All() []any {
	return ofValues()
}

func (v values) Zero() []any {
	return zero()
}

func (v values) Rand() []any {
	return random()
}

type refs struct{}

func (r refs) All() []any {
	return ofRefs()
}

func (r refs) Zero() []any {
	return zeroAsRef()
}

func (r refs) Rand() []any {
	return randomAsRef()
}
