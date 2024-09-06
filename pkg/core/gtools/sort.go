// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtools

// ComparableOf is an interface for types that can be compared by a method call
type ComparableOf interface {
	Equal(o interface{}) bool
}

// SortableOf is an interface for types that can be sorted by a method call
type SortableOf interface {
	Less(o interface{}) bool
	ComparableOf
}

// PersistentSortableOf is an interface for sortable types that are meant to be persisted,
// so they must be able to be recognizable even when its memory address is changed
type PersistentSortableOf interface {
	SortableOf
	Unique() []byte
}

// AggregableOf is an interface for types that can be aggregated by a method call
type AggregableOf interface {
	ComparableOf
	Add(other AggregableOf) AggregableOf
}
