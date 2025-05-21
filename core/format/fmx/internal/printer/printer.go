// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package printer

//import (
//	"fmt"
//	"strings"
//)
//
//type Spec struct {
//	Precision        *int // Precision (optional)
//	Width            int  // Minimum width
//	Verb             rune // e.g., 's', 'd', 'f'
//	LeftAlign        bool // LeftChild alignment flag
//	PadWithZero      bool // Zero padding flag
//	ForceSign        bool // Force '+' sign for numbers
//	PadSignWithSpace bool // Pad positive numbers with space instead of '+'
//	Alternative      bool // '#' flag (alternate form)
//	UseGrouping      bool // ',' flag (group digits with thousands separator)
//	DynamicWidth     bool // '*' flag (width from argument)
//	DynamicPrecision bool // '.' flag with '*' (precision from argument)
//}
//
//func applySpecifier(builder *strings.Builder, spec Spec, arg any) error {
//	switch spec.Verb {
//	case 's':
//		return formatString(builder, spec, arg)
//	case 'd':
//		return formatInt(builder, spec, arg)
//	case 'f':
//		return formatFloat(builder, spec, arg)
//	case 'v':
//		return formatDefault(builder, spec, arg)
//	case 'T':
//		return formatType(builder, spec, arg)
//	case 't':
//		return formatBool(builder, spec, arg)
//	default:
//		return fmt.Errorf("unsupported verb: %c", spec.Verb)
//	}
//}
