// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package term

import (
	"github.com/andrerrcosta2/gtools/core/format/custom"
)

type Styles string

const (
	Bold             Styles = "\033[1m"
	BoldBlack        Styles = "\033[1;30m"
	BoldBrightBlack  Styles = "\033[1;90m"
	BoldRed          Styles = "\033[1;31m"
	BoldBrightRed    Styles = "\033[1;91m"
	BoldGreen        Styles = "\033[1;32m"
	BoldBrightGreen  Styles = "\033[1;92m"
	BoldYellow       Styles = "\033[1;33m"
	BoldBrightYellow Styles = "\033[1;93m"
	BoldBlue         Styles = "\033[1;34m"
	BoldBrightBlue   Styles = "\033[1;94m"
	BoldPurple       Styles = "\033[1;35m"
	BoldBrightPurple Styles = "\033[1;95m"
	BoldCyan         Styles = "\033[1;36m"
	BoldBrightCyan   Styles = "\033[1;96m"
	BoldWhite        Styles = "\033[1;37m"
	BoldBrightWhite  Styles = "\033[1;97m"
	Italic           Styles = "\033[3m"
	Underline        Styles = "\033[4m"
	UnderlineBlack   Styles = "\033[4;30m"
	UnderlineRed     Styles = "\033[4;31m"
	UnderlineGreen   Styles = "\033[4;32m"
	UnderlineYellow  Styles = "\033[4;33m"
	UnderlineBlue    Styles = "\033[4;34m"
	UnderlinePurple  Styles = "\033[4;35m"
	UnderlineCyan    Styles = "\033[4;36m"
	UnderlineWhite   Styles = "\033[4;37m"
	Strikethrough    Styles = "\033[9m"
)

func (c Styles) Style(s ...any) string {
	return Sprint(c, s...)
}

func (c Styles) Stylef(format string, args ...any) string {
	return Sprintf(c, format, args...)
}

func (c Styles) Styleln(s ...any) string {
	return Sprintln(c, s...)
}

func (c Styles) String() string {
	return string(c)
}

func (c Styles) Equals(s custom.Style) bool {
	return string(c) == s.String()
}
