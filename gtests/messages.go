// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtests

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
)

func ErrorMsg(msg, rec, exp string) string {
	return fmx.Sprint(msg, " {") +
		"\n" + fmx.Sprint("received:", "\n") +
		"'" + indent.Soft(1, stringMsg(rec)) + "'" +
		"\n" + fmx.Sprint("expected:", "\n") +
		"'" + indent.Soft(1, stringMsg(exp)) + "'" + fmx.Sprint("\n", "}")
}

func stringMsg(s string) string {
	if s == "" {
		return "<empty>"
	}
	return s
}
