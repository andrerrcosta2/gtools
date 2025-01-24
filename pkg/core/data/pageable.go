// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

type Pageable interface {
	// PageNumber returns the current page number (0-based).
	PageNumber() int

	// PageSize returns the number of elements per page.
	PageSize() int
	// Offset returns the offset of the first element on the current page.
	Offset() int
}

// NewPageRequest creates a new PageRequest
//
// It takes two parameters: the page number and the page size.
//
// Returns a pointer to the created PageRequest.
func NewPageRequest(pageNumber, pageSize int) *PageRequest {
	return &PageRequest{pageNumber: pageNumber, pageSize: pageSize}
}

// PageRequest represents a request for a specific page with a size
type PageRequest struct {
	pageNumber int
	pageSize   int
}

// PageNumber returns the current page number
func (p *PageRequest) PageNumber() int {
	return p.pageNumber
}

// PageSize returns the page size
func (p *PageRequest) PageSize() int {
	return p.pageSize
}

// Offset returns the offset for the current page
func (p *PageRequest) Offset() int {
	return p.pageNumber * p.pageSize
}
