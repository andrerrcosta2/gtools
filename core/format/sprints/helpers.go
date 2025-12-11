// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprints

import (
	"strings"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
)

func Anonymous(tab indent.Tab, typx string, fields ...string) string {
	if len(fields) == 0 {
		return tab.Sprintf("%s {}", typx)
	}

	s := strings.Join(fields, "; ")

	return tab.Sprintf("%s { %s }", typx, s)
}

func Append(whole, data string) string {
	return fmx.Sprintf("%s%s", whole, data)
}

func AppendErrors(errs ...error) string {
	var sb strings.Builder

	for _, err := range errs {
		sb.WriteString(err.Error() + "\n")
	}
	return sb.String()
}

func KeyValue(tab indent.Tab, key, value string) string {
	return tab.Sprint(key + " " + value)
}

func LKeyValue(tab indent.Tab, typx, value string) string {
	return "\n" + tab.Sprint(typx+" "+value)
}

func LtKeyValue(tab indent.Tab, typx, value string) string {
	return "\n" + tab.Sprint("\t"+typx+" "+value)
}

func MaxCharsLeft(tab indent.Tab, s string, max int) string {
	runes := []rune(s)
	if len(runes) <= max {
		return tab.Sprintf("%s", s)
	}
	return tab.Sprintf("%s...", string(runes[:max]))
}

func MaxCharsRight(tab indent.Tab, s string, max int) string {
	runes := []rune(s)
	if len(runes) <= max {
		return tab.Sprintf("%s", s)
	}
	return tab.Sprintf("...%s", string(runes[len(runes)-max:]))
}
