// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

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
