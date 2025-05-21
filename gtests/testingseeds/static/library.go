// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

// DeepBook Representing a library system.
type DeepBook struct {
	Title     string // Title of the book
	Author    string // Author of the book
	ISBN      string // ISBN number
	PageCount int    // Total pages
}

type DeepShelf struct {
	Label string     // Shelf label or identifier
	Books []DeepBook // Books on the shelf
}

type DeepSection struct {
	Name    string      // Section name
	Shelves []DeepShelf // Shelves in the section
}

type DeepLibrary struct {
	Name     string        // Library name
	Address  string        // Location of the library
	Sections []DeepSection // Sections in the library
}
