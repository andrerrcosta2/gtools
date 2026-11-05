// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package searcher

// Linear is a minimal searcher for "flat" collections
type Linear[N any] interface {
	Has(N) bool              // check if collection contains the element
	Find(target N) (N, bool) // retrieve the actual stored value
}

// Graph extends Linear with traversal and path-finding
type Graph[N any] interface {
	Linear[N]
	Traverse(yield func(N) bool) // visit each node
	Path(from, to N) []N         // return path from->to
}
