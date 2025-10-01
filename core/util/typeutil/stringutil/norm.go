// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package stringutil

import (
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// NormNFC returns the canonical composed form of the string.
func NormNFC(s string) string {
	if norm.NFC.IsNormalString(s) {
		return s
	}
	return norm.NFC.String(s)
}

// NormNFD returns the canonical decomposed form of the string.
func NormNFD(s string) string {
	if norm.NFD.IsNormalString(s) {
		return s
	}
	return norm.NFD.String(s)
}

// NormWS replaces all sequences of whitespace (spaces, tabs, newlines)
// with a single space and trims leading/trailing whitespace.
// Example: NormWS(" a  \t b\n\nc ") → "a b c"
func NormWS(s string) string {
	var b strings.Builder
	wasWhitespace := false

	for _, r := range s {
		if unicode.IsSpace(r) {
			if !wasWhitespace {
				b.WriteRune(' ')
				wasWhitespace = true
			}
		} else {
			b.WriteRune(r)
			wasWhitespace = false
		}
	}
	return b.String()
}
