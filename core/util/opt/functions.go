// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package opt

// EmptyString returns the non-empty string between `a` and `b`.
//
// If `a` is empty, it returns `b`. Otherwise, it returns `a`.
//
// Parameters:
// - a: The first string to compare.
// - b: The second string to compare.
//
// Returns:
// - string: The non-empty string between `a` and `b`.
func EmptyString(a, b string) string {
	if a == "" {
		return b
	}
	return a
}
