// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package indent

import (
	"strings"
	"testing"

	"github.com/andrerrcosta2/gtools/core/format/printer"
)

func TestSmarkf(t *testing.T) {
	tests := []struct {
		name     string
		tab      Branch
		format   string
		args     []any
		expected string
	}{
		{
			name:     "No indentation with Smarkf",
			tab:      Branch(0),
			format:   "Hello, %s!",
			args:     []any{"World"},
			expected: "Hello, World!", // Root level: no marker
		},
		{
			name:     "One level of indentation with Smarkf",
			tab:      Branch(Zero().Inc().Value()),
			format:   "Value: %d",
			args:     []any{42},
			expected: "└── Value: 42", // First child of the root
		},
		{
			name:     "Two levels of indentation with Smarkf",
			tab:      Branch(Zero().Inc().Inc().Value()),
			format:   "Key: %s, Value: %v",
			args:     []any{"example", true},
			expected: "\t└── Key: example, Value: true", // Second level
		},
		{
			name:     "Multi-line string with Smarkf",
			tab:      Branch(Zero().Inc().Value()),
			format:   "Line 1\n\t\t%s",
			args:     []any{"test"},
			expected: "└── Line 1\n\t└── test", // Multi-line with tree structure
		},
		{
			name:   "Nested levels with Smarkf",
			tab:    Branch(0),
			format: "Outer:\n\t%s\n\t\tInner:\n\t\t\t%s\n\t\t\t\tValue: %d",
			args:   []any{"Level 1", "Level 2", 42},
			expected: `Outer:
└── Level 1
	└── Inner:
		└── Level 2
			└── Value: 42`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Smarkf(tt.tab, tt.format, tt.args...)
			if result != tt.expected {
				t.Errorf("Smarkf(%q, %v) = %q, expected %q", tt.format, tt.args, result, tt.expected)
				printer.Print(quickDiff(result, tt.expected))
			}
			t.Log("Result: =====================")
			t.Logf("\n%s", result)
			t.Log("Expected: ===================")
			t.Logf("\n%s", tt.expected)
		})
	}
}

func TestIndentWithTabs(t *testing.T) {
	tests := []struct {
		name     string
		tab      Tab
		inc      int
		input    string
		expected string
	}{
		{
			name:     "Indent single line with tabs",
			tab:      Zero(),
			inc:      1,
			input:    "Hello, World!",
			expected: "\tHello, World!",
		},
		{
			name:     "Indent multiple lines with tabs",
			tab:      Tab(Zero().Inc().Value()),
			inc:      1,
			input:    "Line 1\nLine 2\nLine 3",
			expected: "\t\tLine 1\n\t\tLine 2\n\t\tLine 3",
		},
		{
			name:     "Indent with existing tabs",
			tab:      Zero(),
			inc:      2,
			input:    "\tIndented Line",
			expected: "\t\t\tIndented Line",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tab.Plus(tt.inc).Indent(0, tt.input)
			if result != tt.expected {
				t.Errorf("Indent(%d, %q) = %q, expected %q", tt.inc, tt.input, result, tt.expected)
			}
		})
	}
}

func quickDiff(received, expected string) string {
	if received != expected {
		for i := 0; i < len(received) || i < len(expected); i++ {
			if i >= len(received) {
				return printer.Sprintf("Received is shorter. Missing: %q\n", expected[i:])
			}
			if i >= len(expected) {
				return printer.Sprintf("Expected is shorter. Extra: %q\n", received[i:])
			}
			if received[i] != expected[i] {
				sb := strings.Builder{}
				sb.WriteString(printer.Sprintf("Difference at index %d:\n", i))
				sb.WriteString(printer.Sprintf("  Received: %q\n", received[i]))
				sb.WriteString(printer.Sprintf("  Expected: %q\n", expected[i]))
				return sb.String()
			}
		}
	}
	return "Strings match."
}
