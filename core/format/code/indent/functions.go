// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package indent

import (
	"fmt"
	"strings"

	"github.com/andrerrcosta2/gtools/core/util/typeutil/stringutil"
)

// Hard enforces a minimum indentation level for each line in the input string.
//
// The function ensures that every line in the input string has at least the specified
// indentation level. If a line already has sufficient indentation, it remains unchanged;
// otherwise, the line is adjusted to meet the required indentation level.
// Additionally, all lines are further indented by the specified value, ensuring consistent
// alignment across the entire block.
//
// Parameters:
//   - value: The number of indentation levels to enforce. Each level corresponds to one tab (`\t`)
//     or a configurable number of spaces, depending on the implementation of `Tab`.
//   - s:     The input string to be processed. This string may contain multiple lines separated by `\n`.
//
// Returns:
// A new string where each line meets the minimum indentation requirements and is further
// indented by the specified value.
//
// Example:
//
//	input := "func main() {\nfmt.Println(\"Hello, World!\")\n}"
//	fmt.Print(Hard(2, input))
//
// Output:
//
//	func main() {
//	    fmt.Println("Hello, World!")
//	}
func Hard[I Indentor](indentor I, s string) string {
	// Split the input string into lines
	lines := strings.Split(s, "\n")

	// Build the formatted result
	var result strings.Builder
	for i, line := range lines {
		// Trim only trailing whitespace (not leading) to preserve existing indentation
		line = strings.TrimRightFunc(line, func(r rune) bool {
			return r == ' '
		})

		// Add the required indentation to the line
		result.WriteString(indentor.Sprint(line))

		// Add a newline character unless it's the last line
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

// Soft applies a uniform indentation to the beginning of each line in the input string.
//
// The function splits the input string into lines, applies the specified indentation level
// (in terms of tabs or spaces) to each line, and returns the indented result as a single string.
// Existing internal formatting (e.g., tabs or spaces within the lines) is preserved.
//
// Parameters:
//   - value: The number of indentation levels to apply. Each level corresponds to one tab (`\t`)
//     or a configurable number of spaces, depending on the implementation of `Tab`.
//   - s:     The input string to be indented. This string may contain multiple lines separated by `\n`.
//
// Returns:
// A new string where each line starts with the specified indentation level.
//
// Example:
//
//	input := "func main() {\n\tfmt.Println(\"Hello, World!\")\n}"
//	fmt.Print(Soft(1, input))
//
// Output:
//
//	func main() {
//		fmt.Println("Hello, World!")
//	}
func Soft[I Indentor](indentor I, s string) string {
	// Split the input string into lines
	lines := strings.Split(s, "\n")

	// Build the formatted result
	var result strings.Builder
	for i, line := range lines {
		if line == "" || line == "\n" || line == "\r" || line == "\t" {
			result.WriteString(line)
		} else {
			result.WriteString(indentor.Sprint(line))
		}

		// Add a newline character unless it's the last line
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

func Smarklt(s ...any) string {
	sb := strings.Builder{}
	for _, v := range s {
		sb.WriteString(stringutil.ReplaceLast(fmt.Sprint(v), "\t", "└── "))
	}
	return sb.String()
}

func Smark(indentor Indentor, s ...any) string {
	ss := indentor.Sprint(s...)
	lines := strings.Split(ss, "\n")
	for i, line := range lines {
		lines[i] = Smarkln(line)
	}
	return strings.Join(lines, "\n")
}

func Smarkf(indentor Indentor, format string, args ...any) string {
	s := indentor.Sprintf(format, args...)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = Smarkln(line)
	}
	return strings.Join(lines, "\n")
}

func Smarkln(line string) string {
	tabs := -1
	for i, r := range line {
		if r == '\t' {
			tabs++
		} else {
			if strings.HasPrefix(strings.TrimSpace(line[i:]), "└──") {
				return line
			}
			if tabs == -1 {
				return line
			} else {
				return line[:tabs] + "└── " + line[tabs+1:]
			}
		}
	}
	return line
}
