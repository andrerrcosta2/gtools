// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package term

import (
	"github.com/andrerrcosta2/gtools/core/format/custom"
	"github.com/andrerrcosta2/gtools/core/format/printer"
	"strings"
)

var reset = Reset{}

type Reset struct{}

func (r Reset) String() string {
	return "\033[0m"
}

func (r Reset) Equals(s custom.Style) bool {
	return "\033[0m" == s.String()
}

func (r Reset) Style(s string) string {
	return s + "\033[0m"
}

func (r Reset) Stylef(format string, args ...any) string {
	return printer.Sprintf(format, args...) + "\033[0m"
}

func Sprint(s custom.Style, str ...any) string {
	if custom.None.Equals(s) {
		return printer.Sprint(str...)
	}
	sentence := printer.Sprint(str...)
	lines := strings.Split(sentence, "\n")
	for i := range lines {
		lines[i] = s.String() + lines[i] + "\033[0m"
	}
	return strings.Join(lines, "\n")
}

func Sprintf(s custom.Style, format string, args ...any) string {
	if custom.None.Equals(s) {
		return printer.Sprintf(format, args...)
	}
	lines := strings.Split(printer.Sprintf(format, args...), "\n")
	for i := range lines {
		lines[i] = s.String() + lines[i] + "\033[0m"
	}
	return strings.Join(lines, "\n")
}

func Sprintln(s custom.Style, str ...any) string {
	if custom.None.Equals(s) {
		return printer.Sprintln(str...)
	}
	lines := strings.Split(printer.Sprintln(str...), "\n")
	for i := range lines {
		lines[i] = s.String() + lines[i] + "\033[0m"
	}
	return strings.Join(lines, "\n")
}
