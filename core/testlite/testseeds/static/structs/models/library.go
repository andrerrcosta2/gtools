// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import "github.com/andrerrcosta2/gtools/core/seeders/random"

// DeepBook Representing a library system.
type DeepBook struct {
	Title     string // Title of the book
	Author    string // Author of the book
	ISBN      string // ISBN number
	PageCount int    // Total pages
}

func DeepBookAsValue(title, author, isbn string, pageCount int) DeepBook {
	return DeepBook{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		PageCount: pageCount,
	}
}

func DeepBookAsRef(title, author, isbn string, pageCount int) *DeepBook {
	return &DeepBook{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		PageCount: pageCount,
	}
}

func DeepBookAsRandValue() DeepBook {
	return DeepBookAsValue(
		random.Alphabet(1, 3, 30).At(0),
		random.Alphabet(1, 3, 15).At(0)+" "+random.Alphabet(1, 3, 15).At(0),
		random.Alphanumeric(1, 3, 13).At(0),
		random.Int(100, 2500).At(0),
	)
}

func DeepBookAsRandRef() *DeepBook {
	return DeepBookAsRef(
		random.Alphabet(1, 3, 30).At(0),
		random.Alphabet(1, 3, 15).At(0)+" "+random.Alphabet(1, 3, 15).At(0),
		random.Alphanumeric(1, 3, 13).At(0),
		random.Int(100, 2500).At(0),
	)
}

type DeepShelf struct {
	Label string     // Shelf label or identifier
	Books []DeepBook // Books on the shelf
}

func DeepShelfAsValue(label string, books []DeepBook) DeepShelf {
	return DeepShelf{
		Label: label,
		Books: books,
	}
}

func DeepShelfAsRef(label string, books []DeepBook) DeepShelf {
	return DeepShelf{
		Label: label,
		Books: books,
	}
}

func DeepShelfAsRandValue() DeepShelf {
	shelf := DeepShelf{
		Label: random.Alphabet(1, 3, 15).At(0),
	}
	size := random.Int(1, 1, 50).At(0)
	for i := 0; i < size; i++ {
		shelf.Books = append(shelf.Books, DeepBookAsRandValue())
	}
	return shelf
}

func DeepShelfAsRandRef() *DeepShelf {
	shelf := &DeepShelf{
		Label: random.Alphabet(1, 3, 15).At(0),
	}
	size := random.Int(1, 1, 50).At(0)
	for i := 0; i < size; i++ {
		shelf.Books = append(shelf.Books, DeepBookAsRandValue())
	}
	return shelf
}

type DeepSection struct {
	Name    string      // Section name
	Shelves []DeepShelf // Shelves in the section
}

func DeepSectionAsValue(name string, shelves []DeepShelf) DeepSection {
	return DeepSection{
		Name:    name,
		Shelves: shelves,
	}
}

func DeepSectionAsRef(name string, shelves []DeepShelf) DeepSection {
	return DeepSection{
		Name:    name,
		Shelves: shelves,
	}
}

func DeepSectionAsRandValue() DeepSection {
	section := DeepSection{
		Name: random.Alphabet(1, 3, 15).At(0),
	}
	size := random.Int(1, 1, 10).At(0)
	for i := 0; i < size; i++ {
		section.Shelves = append(section.Shelves, DeepShelfAsRandValue())
	}
	return section
}

func DeepSectionAsRandRef() *DeepSection {
	section := &DeepSection{
		Name: random.Alphabet(1, 3, 15).At(0),
	}
	size := random.Int(1, 1, 10).At(0)
	for i := 0; i < size; i++ {
		section.Shelves = append(section.Shelves, DeepShelfAsRandValue())
	}
	return section
}

type DeepLibrary struct {
	Name     string        // Library name
	Address  string        // Location of the library
	Sections []DeepSection // Sections in the library
}

func DeepLibraryAsValue(name, address string, sections []DeepSection) DeepLibrary {
	return DeepLibrary{
		Name:     name,
		Address:  address,
		Sections: sections,
	}
}

func DeepLibraryAsRef(name, address string, sections []DeepSection) DeepLibrary {
	return DeepLibrary{
		Name:     name,
		Address:  address,
		Sections: sections,
	}
}

func DeepLibraryAsRandValue() DeepLibrary {
	library := DeepLibrary{
		Name:    random.Alphabet(1, 3, 15).At(0),
		Address: random.Alphabet(1, 3, 30).At(0),
	}
	size := random.Int(1, 1, 5).At(0)
	for i := 0; i < size; i++ {
		library.Sections = append(library.Sections, DeepSectionAsRandValue())
	}
	return library
}

func DeepLibraryAsRandRef() *DeepLibrary {
	library := &DeepLibrary{
		Name:    random.Alphabet(1, 3, 15).At(0),
		Address: random.Alphabet(1, 3, 30).At(0),
	}
	size := random.Int(1, 1, 5).At(0)
	for i := 0; i < size; i++ {
		library.Sections = append(library.Sections, DeepSectionAsRandValue())
	}
	return library
}
