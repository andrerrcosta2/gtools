// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/custom/term"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/stringutil"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
	"strings"
	"unicode"
)

// Text returns a diff between the received and expected strings.
func Text[O internal.Option](tab indent.Tab, value, expected string, before, after int, o ...O) (diff string, equals bool) {
	s := NewStrategy(o...)
	value, expected = formatStrings(s.stringFormat, value, expected)
	rec, exp, idx := s.strings(value, expected)
	if idx == -1 {
		return "", true
	}

	// Calculate the start and end indices for the received string
	startRec := max(0, idx-before)
	endRec := min(len(rec), idx+after)

	// Calculate the start and end indices for the expected string
	startExp := max(0, idx-before)
	endExp := min(len(exp), idx+after)

	// Construct the diff
	return differs.Message(differs.Strings(tab, idx), differs.Diff(tab,
		rec[startRec:endRec], exp[startExp:endExp])), false
}

func formatStrings(f op.Compare, a, b string) (string, string) {
	if f == 0 {
		return a, b
	}
	if f&compare.TrimSpace != 0 {
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)
	}
	if f&compare.IgnoreWhitespace != 0 {
		a = stringutil.NormWS(a)
		b = stringutil.NormWS(b)
	}
	if f&compare.IgnoreCase != 0 {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}
	if f&compare.IgnoreAccents != 0 {
		a = stringutil.RemoveAccents(a)
		b = stringutil.RemoveAccents(b)
	}
	return a, b
}

func defaultStringDiff(value, expected string) (string, string, int) {
	runesValue := []rune(value)
	runesExpected := []rune(expected)

	minLen := min(len(runesValue), len(runesExpected))

	// Find the first differing index
	diffIdx := -1
	for i := 0; i < minLen; i++ {
		if runesValue[i] != runesExpected[i] {
			diffIdx = i
			break
		}
	}

	// Handle length mismatch
	if diffIdx == -1 && len(runesValue) != len(runesExpected) {
		diffIdx = minLen
	}

	return markDiff(value, diffIdx), expected, diffIdx
}

func ignoreCaseStringDiff(value, expected string) (string, string, int) {
	runesValue := []rune(value)
	runesExpected := []rune(expected)

	minLen := min(len(runesValue), len(runesExpected))

	// Find the first differing index
	diffIdx := -1
	for i := 0; i < minLen; i++ {
		if unicode.ToLower(runesValue[i]) != unicode.ToLower(runesExpected[i]) {
			diffIdx = i
			break
		}
	}

	// Handle length mismatch
	if diffIdx == -1 && len(runesValue) != len(runesExpected) {
		diffIdx = minLen
	}

	return markDiff(value, diffIdx), expected, diffIdx
}

func markDiff(s string, idx int) string {
	runes := []rune(s)

	// Out of bounds → no diff
	if idx < 0 {
		return ""
	} else if idx >= len(runes) {
		return s
	}

	var result strings.Builder
	for i, r := range runes {
		inDiff := i == idx // only highlight the actual differing rune

		// Special case: space inside diff → red background
		if inDiff && r == ' ' {
			result.WriteString(fmx.SCustom(term.BgBrightRed, " "))
			continue
		}

		if inDiff {
			// Wrap invisibles/raw rune in red
			result.WriteString(fmx.SRed(formatInvisibleRune(r, true)))
		} else {
			// Outside diff: normal formatting
			result.WriteString(formatInvisibleRune(r, false))
		}
	}

	return result.String()
}

func formatInvisibleRune(r rune, inDiff bool) string {
	switch r {
	case '\t':
		if inDiff {
			return "\\t" // leave plain, will be wrapped red
		}
		return fmx.SCyan("\\t")
	case '\n':
		if inDiff {
			return "\\n"
		}
		return fmx.SCyan("\\n")
	case '\r':
		if inDiff {
			return "\\r"
		}
		return fmx.SCyan("\\r")
	default:
		return string(r)
	}
}
