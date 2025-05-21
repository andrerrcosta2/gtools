// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package term

import (
	"github.com/andrerrcosta2/gtools/core/format/custom"
)

type Bg string

const (
	BgRed     Bg = "\033[41m"
	BgGreen   Bg = "\033[42m"
	BgYellow  Bg = "\033[43m"
	BgBlue    Bg = "\033[44m"
	BgMagenta Bg = "\033[45m"
	BgCyan    Bg = "\033[46m"
	BgWhite   Bg = "\033[47m"
	BgBlack   Bg = "\033[40m"

	BgDarkGrey      Bg = "\033[48;5;236m"
	BgVeryDarkGrey  Bg = "\033[48;2;10;10;10m"
	BgCharcoal      Bg = "\033[48;5;236m"
	BgDarkCharcoal  Bg = "\033[48;2;25;25;25m"
	BgDarkSlateGrey Bg = "\033[48;2;47;79;79m"
	BgSlateGrey     Bg = "\033[48;5;237m"
	BgMidnightBlue  Bg = "\033[48;2;25;25;112m"
	BgNavyBlue      Bg = "\033[48;5;18m"

	BgBrightBlack   Bg = "\033[100m"
	BgBrightRed     Bg = "\033[101m"
	BgBrightGreen   Bg = "\033[102m"
	BgBrightYellow  Bg = "\033[103m"
	BgBrightBlue    Bg = "\033[104m"
	BgBrightMagenta Bg = "\033[105m"
	Bgs             Bg = "\033[106m"
	BgBrightWhite   Bg = "\033[107m"
)

func (c Bg) Style(s ...any) string {
	return Sprint(c, s)
}

func (c Bg) Stylef(format string, args ...any) string {
	return Sprintf(c, format, args...)
}

func (c Bg) Styleln(s ...any) string {
	return Sprintln(c, s...)
}

func (c Bg) String() string {
	return string(c)
}

func (c Bg) Equals(s custom.Style) bool {
	return string(c) == s.String()
}
