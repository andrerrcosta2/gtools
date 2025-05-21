// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"strings"
)

// Text returns a diff between the received and expected strings.
func Text(tab indent.Tab, value, expected string, before, after int) (diff string, equals bool) {
	received, expected, idx := difference(value, expected, before, after)
	if idx == -1 {
		return "", true
	}

	// Calculate the start and end indices for the received string
	startReceived := max(0, idx-before)
	endReceived := min(len(received), idx+after)

	// Calculate the start and end indices for the expected string
	startExpected := max(0, idx-before)
	endExpected := min(len(expected), idx+after)

	// Extract the substrings to display
	receivedSubstring := received[startReceived:endReceived]
	expectedSubstring := expected[startExpected:endExpected]

	// Construct the diff
	return differs.Message(differs.Strings(tab, idx), differs.Diff(tab, receivedSubstring, expectedSubstring)), false
}

func difference(value, expected string, before, after int) (string, string, int) {
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

	return markDiff(value, diffIdx, before, after), expected, diffIdx
}

func markDiff(s string, idx int, before, after int) string {
	runes := []rune(s)

	visualized := formatInvisibleChars(s)
	visualRunes := []rune(visualized)

	// Ensure idx is within bounds
	if idx < 0 || idx >= len(runes) {
		// If idx is out of bounds, highlight the entire string
		coloredPart := fmx.SRed(string(visualRunes))
		return coloredPart
	}

	// Calculate the start and end indices for highlighting
	start := max(0, idx-before)
	end := min(len(runes), idx+after)

	// Ensure valid bounds
	if start > len(runes) {
		start = len(runes)
	}
	if end < start {
		end = start
	}

	// Apply red color to the differing portion
	coloredPart := fmx.SRed(string(visualRunes[idx:min(end, len(visualRunes))]))

	// Reconstruct the string with the colored part
	result := string(visualRunes[:start]) + string(visualRunes[start:idx]) + coloredPart + string(visualRunes[end:])

	return result
}

func formatInvisibleChars(s string) string {
	var result strings.Builder
	for _, r := range s {
		switch r {
		case ' ':
			result.WriteRune('␣') // Replace space with middle dot
		case '\t':
			result.WriteString("→") // Replace tab with arrow
		case '\n':
			result.WriteString("⏎\n") // Replace newline with return symbol
		case '\r':
			result.WriteString("↵") // Replace carriage return with return symbol
		default:
			result.WriteRune(r)
		}
	}
	return result.String()
}
