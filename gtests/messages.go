// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtests

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
)

// ErrorDiff prints the error message and the received and expected values:
//
//	<message> {
//		received:
//			'<received>'
//		expected:
//			'<expected>'
//	}
func ErrorDiff(msg, rec, exp string) string {
	return fmx.Sprint(msg, " {") +
		"\n" + fmx.Sprint("received:", "\n") +
		"'" + indent.Soft(indent.Tab(1), stringMsg(rec)) + "'" +
		"\n" + fmx.Sprint("expected:", "\n") +
		"'" + indent.Soft(indent.Tab(1), stringMsg(exp)) + "'" + fmx.Sprint("\n", "}")
}

// ErrorMismatchValues reports a mismatch between received and expected values:
//
//	"<message>:
//		received: '<received>', expected: '<expected>'"
func ErrorMismatchValues(msg string, rec, exp any) string {
	return fmx.Sprintf("%s: \n\treceived: '%v', expected: '%v'", msg, rec, exp)
}

func stringMsg(s string) string {
	if s == "" {
		return "<empty>"
	}
	return s
}
