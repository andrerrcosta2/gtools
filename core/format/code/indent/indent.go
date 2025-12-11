// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package indent

import (
	"math"
	"strings"

	"github.com/andrerrcosta2/gtools/core/format/printer"
)

type Indentor interface {
	Dec() Indentor
	Inc() Indentor
	Indent(inc int, s ...any) string
	Indentf(inc int, format string, args ...any) string
	Plus(int) Indentor
	Sprint(str ...interface{}) string
	Sprintf(str string, args ...interface{}) string
	Sprintln(str ...interface{}) string
	Sprintlnf(str string, args ...interface{}) string
	String() string
	Trim(str string) string
	Value() uint16
}

// Root returns a new root Branch
func Root() Branch { return 0 }

// Zero returns a new zero Tab
func Zero() Tab {
	return 0
}

type Tab uint16

func (t Tab) Dec() Indentor {
	if t == 0 {
		return Zero()
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

func (t Tab) Inc() Indentor {
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

func (t Tab) Plus(n int) Indentor {
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

func (t Tab) Trim(str string) string {
	return strings.TrimPrefix(str, t.String())
}

func (t Tab) Value() uint16 {
	return uint16(t)
}

type Branch uint16

func (t Branch) Dec() Indentor {
	if t == 0 {
		return Branch(0)
	}
	return t - 1
}

func (t Branch) Indent(inc int, s ...any) string {
	return Soft(t.Plus(inc), printer.Sprint(s...))
}

func (t Branch) Indentf(inc int, format string, args ...any) string {
	s := printer.Sprintf(format, args...)
	s = strings.TrimLeftFunc(s, func(r rune) bool { // Assign trimmed string back
		return r == '\t' || r == '\n' || r == ' '
	})
	return Soft(t.Plus(inc), s)
}

func (t Branch) Inc() Indentor {
	return t + 1
}

func (t Branch) Join(sep string, s ...string) string {
	sb := strings.Builder{}
	for i, ss := range s {
		sb.WriteString(t.Sprint(ss))
		if i < len(s)-1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}

func (t Branch) Joinln(s ...string) string {
	sb := strings.Builder{}
	for _, ss := range s {
		sb.WriteString(t.Sprint(ss) + "\n")
	}
	return sb.String()
}

func (t Branch) Plus(n int) Indentor {
	return Branch(math.Max(0, float64(t+Branch(n))))
}

func (t Branch) Size() int {
	return int(t)
}

func (t Branch) Sprint(data ...any) string {
	return t.String() + printer.Sprint(data...)
}

func (t Branch) Sprintf(format string, s ...any) string {
	return t.Sprint(printer.Sprintf(format, s...))
}

func (t Branch) Sprintfln(format string, data ...any) string {
	return t.Sprint(printer.Sprintf(format, data...) + "\n")
}

func (t Branch) Sprintln(data ...any) string {
	return t.Sprint(printer.Sprintln(data...))
}

func (t Branch) Sprintlnf(format string, data ...any) string {
	return "\n" + t.Sprint(printer.Sprintf(format, data...))
}

func (t Branch) String() string {
	if t == 0 {
		return ""
	}
	return strings.Repeat("\t", int(t-1)) + "└── "
}

func (t Branch) Trim(str string) string {
	return strings.TrimPrefix(str, t.String())
}

func (t Branch) Value() uint16 {
	return uint16(t)
}

var _ Indentor = (Branch)(0)
