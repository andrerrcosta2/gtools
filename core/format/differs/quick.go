// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"github.com/andrerrcosta2/gtools/core/format/printer"
	"strings"
)

func Quick(received, expected string) string {
	if received == expected {
		return "Strings match."
	}

	for i := 0; i < len(received) || i < len(expected); i++ {
		// Check if one string is shorter than the other
		if i >= len(received) {
			return buildDiff(i, printer.Sprintf("Received is shorter. Difference at index %d:", i), received, expected)
		}
		if i >= len(expected) {
			return printer.Sprintf("Expected is shorter. Extra: %q at index %d\n", received[i:], i)
		}

		if received[i] != expected[i] {
			rs, re := max(0, i-10), min(len(received), i+10)
			es, ee := max(0, i-10), min(len(expected), i+10)

			// Build the output
			sb := strings.Builder{}
			sb.WriteString(printer.Sprintf("Difference at index %d:\n", i))
			sb.WriteString(printer.Sprintf("  Received: %q...\n", received[rs:re]))
			sb.WriteString(printer.Sprintf("            %s^\n", strings.Repeat(" ", i-rs)))
			sb.WriteString(printer.Sprintf("  Expected: %q...\n", expected[es:ee]))
			sb.WriteString(printer.Sprintf("            %s^\n", strings.Repeat(" ", i-es)))
			return sb.String()
		}
	}

	return "Strings match."
}

func buildDiff(idx int, message, received, expected string) string {
	rs, re := max(0, idx-10), min(len(received), idx+10)
	es, ee := max(0, idx-10), min(len(expected), idx+10)

	// Build the output
	sb := strings.Builder{}
	sb.WriteString(message)
	sb.WriteString(printer.Sprintf("  Received: %q...\n", received[rs:re]))
	sb.WriteString(printer.Sprintf("            %s^\n", strings.Repeat(" ", idx-rs)))
	sb.WriteString(printer.Sprintf("  Expected: %q...\n", expected[es:ee]))
	sb.WriteString(printer.Sprintf("            %s^\n", strings.Repeat(" ", idx-es)))
	return sb.String()
}
