// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// Package stringutil provides utility functions for manipulating and formatting strings.
package stringutil

import (
	"strings"
	"unicode"
)

// Equals evaluates whether two strings are strictly equals
func Equals(str1, str2 string) bool {
	return str1 == str2
}

// EqualsIgnoreCase evaluates whether two strings are equals ignoring cases
func EqualsIgnoreCase(str1, str2 string) bool {
	return strings.ToLower(str1) == strings.ToLower(str2)
}

// IsEmpty checks if the string is exactly "" (zero length).
// Example: IsEmpty("") → true, IsEmpty(" ") → false
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsBlank checks if the string is empty or contains only whitespace (spaces, tabs, newlines, etc.).
// Uses unicode.IsSpace for full Unicode support.
// Example: IsBlank("   \t\n") → true, IsBlank(" a ") → false
func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// IsUpper checks if all letter characters in the string are uppercase.
// Non-letter characters (digits, symbols, whitespace) are ignored.
// Returns false if there are no letters.
// Example: IsUpper("HELLO") → true, IsUpper("Hello") → false, IsUpper("123") → false
func IsUpper(s string) bool {
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetter
}

// IsLower checks if all letter characters in the string are lowercase.
// Non-letter characters are ignored.
// Returns false if there are no letters.
// Example: IsLower("hello") → true, IsLower("Hello") → false, IsLower("123") → false
func IsLower(s string) bool {
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsLower(r) {
				return false
			}
		}
	}
	return hasLetter
}

// IsAlpha checks if the string contains only Unicode letters (no digits, symbols, or whitespace).
// Returns false if string is empty.
// Example: IsAlpha("hello") → true, IsAlpha("hello1") → false, IsAlpha("café") → true
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsNumeric checks if the string contains only Unicode decimal digits (0-9 and equivalents).
// Returns false if string is empty or contains any non-digit.
// Example: IsNumeric("123") → true, IsNumeric("12.3") → false, IsNumeric("١٢٣") → true (Arabic digits)
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsAlnum checks if the string contains only alphanumeric characters (letters or digits).
// Returns false if string is empty or contains whitespace/symbols.
// Example: IsAlnum("abc123") → true, IsAlnum("abc 123") → false, IsAlnum("café123") → true
func IsAlnum(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsPrintable checks if all characters in the string are printable (including space).
// Uses unicode.IsPrint for Unicode-aware checking.
// Returns false if string is empty.
// Example: IsPrintable("hello") → true, IsPrintable("hello\t") → false
func IsPrintable(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsPrint(r) && r != '\u00AD' /* optional hyphen */ {
			return false
		}
	}
	return true
}

// HasPrefixCI performs a case-insensitive prefix check.
// Uses strings.EqualFold for Unicode-safe comparison.
// Example: HasPrefixCI("Hello World", "he") → true
func HasPrefixCI(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return strings.EqualFold(s[:len(prefix)], prefix)
}

// HasSuffixCI performs a case-insensitive suffix check.
// Uses strings.EqualFold for correctness with Unicode.
// Example: HasSuffixCI("Hello World", "WORLD") → true
func HasSuffixCI(s, suffix string) bool {
	if len(suffix) > len(s) {
		return false
	}
	return strings.EqualFold(s[len(s)-len(suffix):], suffix)
}

// ReplaceFirst replaces the first occurrence of old with new in str.
func ReplaceFirst(str, sub, new string) string {
	if sub == "" || str == "" {
		return str
	}
	i := strings.Index(str, sub)
	if i == -1 {
		return str
	}
	return str[:i] + new + str[i+len(sub):]
}

// ReplaceLast replaces the last occurrence of old with new in str.
func ReplaceLast(str, sub, new string) string {
	if sub == "" || str == "" {
		return str
	}
	i := strings.LastIndex(str, sub)
	if i == -1 {
		return str
	}
	return str[:i] + new + str[i+len(sub):]
}

// ReplaceAt replaces only the specified occurrences (1-indexed) of old with new.
// Example: ReplaceAt("foo foo foo", "foo", "bar", 1, 3) → "bar foo bar"
func ReplaceAt(str, old, new string, occurrences ...int) string {
	if old == "" || str == "" || len(occurrences) == 0 {
		return str
	}

	// Sorting and deduplicating occurrences
	occSet := make(map[int]bool)
	for _, n := range occurrences {
		if n > 0 {
			occSet[n] = true
		}
	}

	if len(occSet) == 0 {
		return str
	}

	var result strings.Builder
	start := 0
	currentOccurrence := 0

	for {
		i := strings.Index(str[start:], old)
		if i == -1 {
			break
		}
		i += start
		currentOccurrence++

		result.WriteString(str[start:i])

		if occSet[currentOccurrence] {
			result.WriteString(new)
		} else {
			result.WriteString(old)
		}

		start = i + len(old)
	}

	// Add remaining part of string
	result.WriteString(str[start:])
	return result.String()
}
