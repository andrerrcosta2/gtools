// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package term

import "github.com/andrerrcosta2/gtools/core/format/custom"

type Category string

const (
	Title  Category = Category(Bold)
	Header Category = Category(Bold)
)

func (c Category) Style(s ...any) string {
	return Sprint(c, s)
}

func (c Category) Stylef(format string, args ...any) string {
	return Sprintf(c, format, args...)
}

func (c Category) Styleln(s ...any) string {
	return Sprintln(c, s...)
}

func (c Category) String() string {
	return string(c)
}

func (c Category) Equals(s custom.Style) bool {
	return string(c) == s.String()
}
