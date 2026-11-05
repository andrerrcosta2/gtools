package constraints

import "github.com/andrerrcosta2/gtools/core/domain/constraints/prim"

type Comparable[T any] interface {
	Equal(other T) bool
}

type Ordered[T any] interface {
	Less(other T) bool
}

type Sortable[T any] interface {
	Comparable[T]
	Ordered[T]
}

type Hashable[T prim.Hashable] interface {
	Hash() T
}

type Identifiable interface {
	ID() string
}

type Entity[T any] interface {
	Identifiable
	Sortable[T]
}

type Aggregable[T any] interface {
	Add(other T) T
}
