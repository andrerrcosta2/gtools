// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package indent

import (
	"github.com/andrerrcosta2/gtools/core/format/printer"
	"math"
	"strings"
)

func Zero() Tab {
	return 0
}

type Tab uint16

func (t Tab) Dec() Tab {
	if t == 0 {
		return 0
	}
	return t - 1
}

// Indent indents the given string by the specified number of tabs
func (t Tab) Indent(inc int, s ...any) string {
	return Soft(t.Plus(inc), printer.Sprint(s...))
}

func (t Tab) Indentf(inc int, format string, args ...any) string {
	s := printer.Sprintf(format, args...)
	s = strings.TrimLeftFunc(s, func(r rune) bool { // Assign trimmed string back
		return r == '\t' || r == '\n' || r == ' '
	})
	return Soft(t.Plus(inc), s)
}

func (t Tab) Inc() Tab {
	return t + 1
}

func (t Tab) Join(sep string, s ...string) string {
	sb := strings.Builder{}
	for i, ss := range s {
		sb.WriteString(t.Sprint(ss))
		if i < len(s)-1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}

func (t Tab) Joinln(s ...string) string {
	sb := strings.Builder{}
	for _, ss := range s {
		sb.WriteString(t.Sprint(ss) + "\n")
	}
	return sb.String()
}

func (t Tab) Smark(s ...any) string {
	return Smark(t, s...)
}

func (t Tab) Smarkf(format string, args ...any) string {
	return Smarkf(t, format, args...)
}

func (t Tab) Plus(n int) Tab {
	return Tab(math.Max(0, float64(t+Tab(n))))
}

func (t Tab) Size() int {
	return int(t)
}

func (t Tab) Sprint(data ...any) string {
	return t.String() + printer.Sprint(data...)
}

func (t Tab) Sprintf(format string, s ...any) string {
	return t.Sprint(printer.Sprintf(format, s...))
}

func (t Tab) Sprintfln(format string, data ...any) string {
	return t.Sprint(printer.Sprintf(format, data...) + "\n")
}

func (t Tab) Sprintln(data ...any) string {
	return t.Sprint(printer.Sprintln(data...))
}

func (t Tab) Sprintlnf(format string, data ...any) string {
	return "\n" + t.Sprint(printer.Sprintf(format, data...))
}

func (t Tab) String() string {
	return strings.Repeat("\t", int(t))
}
