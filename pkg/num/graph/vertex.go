// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/functions"
)

func NewVertex[F any](node F) *Vertex[F] {
	return &Vertex[F]{
		node: node,
	}
}

type Vertex[F any] struct {
	node F
}

func (v *Vertex[F]) Node() F {
	return v.node
}

func NewVertices[F any](vertex ...*Vertex[F]) *Vertices[F] {
	return &Vertices[F]{
		vts: vertex,
	}
}

type Vertices[F any] struct {
	vts []*Vertex[F]
}

func (v *Vertices[F]) Range(consumer functions.BiPredicate[int, *Vertex[F]]) {
	for i, vtx := range v.vts {
		if !consumer(i, vtx) {
			break
		}
	}
}

func (v *Vertices[F]) Len() int {
	return len(v.vts)
}

func (v *Vertices[F]) Add(vtx ...*Vertex[F]) {
	v.vts = append(v.vts, vtx...)
}
