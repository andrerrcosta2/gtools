// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package term

import (
	"github.com/andrerrcosta2/gtools/core/format/custom"
)

type Color string

const (
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	White   Color = "\033[37m"
	Black   Color = "\033[30m"

	BrightRed     Color = "\033[91m"
	BrightGreen   Color = "\033[92m"
	BrightYellow  Color = "\033[93m"
	BrightBlue    Color = "\033[94m"
	BrightMagenta Color = "\033[95m"
	BrightCyan    Color = "\033[96m"
	BrightWhite   Color = "\033[97m"
)

func (c Color) Style(s ...any) string {
	return Sprint(c, s...)
}

func (c Color) Stylef(format string, args ...any) string {
	return Sprintf(c, format, args...)
}

func (c Color) Styleln(s ...any) string {
	return Sprintln(c, s...)
}

func (c Color) String() string {
	return string(c)
}

func (c Color) Equals(s custom.Style) bool {
	return string(c) == s.String()
}
