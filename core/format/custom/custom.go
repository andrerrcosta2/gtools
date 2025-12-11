// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package custom

import (
	"github.com/andrerrcosta2/gtools/core/format/printer"
)

type Style interface {
	Style(...any) string
	Styleln(...any) string
	Stylef(string, ...any) string
	String() string
	Equals(style Style) bool
}

var None Style = &none{}

type none struct{}

func (s none) String() string {
	return printer.Sprintf("%p", s)
}

func (s none) Equals(style Style) bool {
	return s == style
}

func (s none) Style(str ...any) string {
	return printer.Sprint(str...)
}

func (s none) Styleln(str ...any) string {
	return printer.Sprintln(str...)
}

func (s none) Stylef(format string, args ...any) string {
	return printer.Sprintf(format, args...)
}

var (
	Title     Style
	Header    Style
	Value     Style
	Separator Style
)
