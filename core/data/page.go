// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

// Page interface to paginate data
type Page[T any] interface {
	// PageNumber Returns the current page number (0-based index)
	PageNumber() int

	// PageSize Returns the size of the page
	PageSize() int

	// TotalElements Returns the total number of elements across all pages
	TotalElements() int

	// TotalPages Returns the total number of pages
	TotalPages() int

	// HasNext Returns true if there are more pages
	HasNext() bool

	// IsLast Returns true if this is the last page
	IsLast() bool

	// Content Returns the content of the current page
	Content() []T

	// Iterable to allow iterating through the elements in the current page
	Iterable[T]
}

func IsLastPage[T any](p Page[T]) bool {
	return p.PageNumber() == p.TotalPages()-1
}

func HasNextPage[T any](p Page[T]) bool {
	return p.PageNumber() < p.TotalPages()-1
}
