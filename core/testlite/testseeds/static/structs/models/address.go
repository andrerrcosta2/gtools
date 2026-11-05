// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import "github.com/andrerrcosta2/gtools/core/seeders/random"

// Address represents a nested struct within the main struct
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// AddressAsValue Constructor for creating a sample Address
func AddressAsValue(street, city, state, zipCode string) Address {
	return Address{
		Street:  street,
		City:    city,
		State:   state,
		ZipCode: zipCode,
	}
}

// AddressAsPointer Constructor for creating a sample Address
func AddressAsPointer(street, city, state, zipCode string) *Address {
	return &Address{
		Street:  street,
		City:    city,
		State:   state,
		ZipCode: zipCode,
	}
}

// AddressAsRandValue Constructor for creating a random Address
func AddressAsRandValue() Address {
	return AddressAsValue(
		random.Alphabet(1, 1, 20).At(0),
		random.Alphabet(1, 1, 20).At(0),
		random.Alphabet(1, 1, 20).At(0),
		random.Alphabet(1, 2, 4).At(0),
	)
}

// AddressAsRandRef Constructor for creating a random Address
func AddressAsRandRef() *Address {
	return AddressAsPointer(
		random.Alphabet(1, 1, 20).At(0),
		random.Alphabet(1, 1, 20).At(0),
		random.Alphabet(1, 1, 20).At(0),
		random.Alphabet(1, 2, 4).At(0),
	)
}
