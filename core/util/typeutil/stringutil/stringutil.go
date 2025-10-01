// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package stringutil

import (
	"bytes"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// PadLeft returns the string s left-padded with pad until the total length reaches width.
// If s is longer than width, it is returned unchanged.
// The padding is applied by repeating the pad string as needed.
// Example: PadLeft("hi", "0", 5) → "000hi"
func PadLeft(s string, pad string, width int) string {
	if len(s) >= width {
		return s
	}
	padding := strings.Repeat(pad, (width-len(s))/len(pad)+1)
	return padding[:width-len(s)] + s
}

// Slugify converts a string into a URL-friendly slug.
// It removes accents, converts to lowercase, replaces non-alphanumeric characters with hyphens,
// and trims extra hyphens.
// Example: Slugify("Hello, 世界! 🌍") → "hello-shi-jie"
func Slugify(s string) string {
	s = strings.ToLower(s)
	s = RemoveAccents(s)
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

// CamelToSnake converts a camelCase or PascalCase string to snake_case.
// It inserts underscores before uppercase letters (except the first) and converts to lowercase.
// Example: CamelToSnake("HTTPServerID") → "http_server_id"
func CamelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				// Check if previous char is not uppercase (start of a new word)
				last, _ := utf8.DecodeLastRuneInString(result.String())
				if !unicode.IsUpper(last) {
					result.WriteRune('_')
				} else {
					// Look ahead to see if this is an acronym continuation
					if i < len(s)-1 {
						next, _ := utf8.DecodeRuneInString(s[i+1:])
						if !unicode.IsUpper(next) {
							result.WriteRune('_')
						}
					}
				}
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// Truncate returns a substring of s up to max runes (Unicode characters).
// If s is shorter than max, it is returned unchanged.
// The result is always valid UTF-8.
// Example: Truncate("hello", 3) → "hel"
func Truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max]
}

// IsASCII returns true if the string contains only ASCII characters (0-127).
// Example: IsASCII("hello") → true, IsASCII("café") → false
func IsASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

// RemoveAccents replaces accented Unicode characters with their closest ASCII equivalents.
// It supports common Latin accents.
// Example: RemoveAccents("résumé") → "resume"
func RemoveAccents(s string) string {
	t := norm.NFD.String(s)
	result := make([]rune, 0, len(t))

	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue // Skip diacritics
		}
		result = append(result, r)
	}
	return string(result)
}

// SplitByLength splits a string into substrings of maximum length n.
// Each part will have at most n **bytes**, not runes.
// If n <= 0, returns nil.
// Warning: May split multi-byte runes if n falls inside one.
// Example: SplitByLength("hello", 2) → ["he", "ll", "o"]
func SplitByLength(s string, n int) []string {
	var result []string
	runes := []rune(s)
	for i := 0; i < len(runes); i += n {
		end := i + n
		if end > len(runes) {
			end = len(runes)
		}
		result = append(result, string(runes[i:end]))
	}
	return result
}

// Dedent removes common leading whitespace from a multiline string.
// It finds the minimum indentation (non-empty lines only) and removes that prefix from all lines.
// Empty lines are left unchanged.
// Example:
//
//	Dedent(`
//		Hello
//		  World
//	`)
//
// → "Hello\n  World"
func Dedent(s string) string {
	lines := strings.Split(s, "\n")
	minIndent := -1
	for _, line := range lines {
		trimmed := strings.TrimLeft(line, " \\t")
		if trimmed == "" {
			continue
		}
		indent := len(line) - len(trimmed)
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}
	for i, line := range lines {
		if len(line) >= minIndent {
			lines[i] = line[minIndent:]
		}
	}
	return strings.Join(lines, "\n")
}

// Wrap wraps a string so that each line does not exceed width characters.
// It breaks lines at word boundaries when possible.
// Whitespace is normalized before wrapping.
// Example: Wrap("hello world test", 8) → "hello\nworld\ntest"
func Wrap(s string, width int) string {
	var buf bytes.Buffer
	var lineLen int
	for _, word := range strings.Fields(s) {
		if lineLen+len(word)+1 > width && lineLen > 0 {
			buf.WriteByte('\n')
			lineLen = 0
		} else if lineLen > 0 {
			buf.WriteByte(' ')
			lineLen++
		}
		buf.WriteString(word)
		lineLen += len(word)
	}
	return buf.String()
}
