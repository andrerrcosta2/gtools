// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package charsets

const (
	Alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	AlphaNumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// ASCIIPrintable Full ASCII-safe printable set (excluding whitespace)
	ASCIIPrintable = AlphaNumeric + Symbols
	Digits         = "0123456789"
	Lower          = "abcdefghijklmnopqrstuvwxyz"
	// SafeSymbols Safer symbols (less likely to cause encoding issues or input bugs)
	SafeSymbols = "-_@#%+=:.,"
	// Symbols based on OWASP recommendations
	Symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~"
	Upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Validate reports whether the input contains any characters not in the allowed charset.
// It returns (false, firstInvalidRune) if it finds one; otherwise (true, 0).
func Validate(input, charset string) (bool, rune) {
	allowed := make(map[rune]struct{}, len(charset))
	for _, r := range charset {
		allowed[r] = struct{}{}
	}

	for _, r := range input {
		if _, ok := allowed[r]; !ok {
			return false, r
		}
	}
	return true, 0
}
