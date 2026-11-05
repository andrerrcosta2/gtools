// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package predicates

// True is a constant predicate that always returns true
func True[T any](_ T) bool {
	return true
}

// False is a constant predicate that always returns false
func False[T any](_ T) bool {
	return false
}
