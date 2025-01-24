// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package edges

import "github.com/andrerrcosta2/gtools/core/gtools"

type Edge[F any, T any] interface {
	From() F
	To() T
}

type Weighted[F any, T any, W any] interface {
	From() F
	To() T
	Weight() W
}

type SortableEdge[F any, T any] interface {
	Edge[F, T]
	gtools.SortableOf
}

type SortableWeighted[F any, T any, W any] interface {
	Weighted[F, T, W]
	gtools.SortableOf
}

type UniqueSortableEdge[F any, T any] interface {
	Edge[F, T]
	gtools.UniqueOf
	gtools.SortableOf
}

type UniqueSortableWeighted[F any, T any, W any] interface {
	Weighted[F, T, W]
	gtools.UniqueOf
	gtools.SortableOf
}

type SingleTyped[G any] interface {
	Edge[G, G]
}

type SortableSingleTyped[G any] interface {
	SortableEdge[G, G]
}

type UniqueSortableSingleTyped[G any] interface {
	UniqueSortableEdge[G, G]
}

type SingleTypedWeighted[G any, W any] interface {
	Weighted[G, G, W]
}

type SortableSingleTypedWeighted[G any, W any] interface {
	SortableWeighted[G, G, W]
}

type UniqueSortableSingleTypedWeighted[G any, W any] interface {
	UniqueSortableWeighted[G, G, W]
}
