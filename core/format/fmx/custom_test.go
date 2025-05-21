// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fmx

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/custom/term"
	"testing"
)

func TestCustom(t *testing.T) {
	Code(indent.Tab(1), 10, "func DiffError(message string, err error) Difference "+
		"{\n\treturn Difference{\n\t\tMessage: message,\n\t\tEquals:  "+
		"false,\n\t\tErr:     err,\n\t}\n}", term.BgDarkCharcoal)
}

func TestCustomSprint(t *testing.T) {
	value := SRed("test")
	fmt.Printf("%T: %#v\n", value, value)
	Print(value)
}
