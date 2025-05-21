// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package parse

import (
	"strconv"
	"unicode/utf8"
)

const (
	commaSpaceString  = ", "
	nilAngleString    = "<nil>"
	nilParenString    = "(nil)"
	nilString         = "nil"
	mapString         = "map["
	percentBangString = "%!"
	missingString     = "(MISSING)"
	badIndexString    = "(BADINDEX)"
	panicString       = "(PANIC="
	extraString       = "%!(EXTRA "
	badWidthString    = "%!(BADWIDTH)"
	badPrecString     = "%!(BADPREC)"
	noVerbString      = "%!(NOVERB)"
	invReflectString  = "<invalid reflect.Value>"
)

// Spec holds format specifier details
type Spec struct {
	Index       int  // Explicit argument index (e.g., %[2]d)
	Width       int  // Minimum width
	Precision   int  // Precision
	Verb        rune // Format verb ('s', 'd', etc.)
	LeftAlign   bool // '-' flag
	PadWithZero bool // '0' flag
	AltFormat   bool // '#' flag
	PlusSign    bool // '+' flag
	SpaceSign   bool // ' ' flag
}

// Format extracts format specifiers from a format string
func Format(format string) ([]Spec, error) {
	var specs []Spec
	end := len(format)
	i := 0

	for i < end {
		if format[i] != '%' {
			i++
			continue
		}
		i++ // Move past '%'

		// Handle '%%' escape
		if i < end && format[i] == '%' {
			i++
			continue
		}

		spec := Spec{}

		// Parse argument index %[n] (optional)
		if i < end && format[i] >= '0' && format[i] <= '9' {
			start := i
			for i < end && format[i] >= '0' && format[i] <= '9' {
				i++
			}
			if i < end && format[i] == ']' {
				num, _ := strconv.Atoi(format[start:i])
				spec.Index = num
				i++ // Move past ']'
			}
		}

		// Parse flags
		for i < end {
			switch format[i] {
			case '-':
				spec.LeftAlign = true
			case '0':
				spec.PadWithZero = true
			case '#':
				spec.AltFormat = true
			case '+':
				spec.PlusSign = true
			case ' ':
				spec.SpaceSign = true
			default:
				goto EndFlags
			}
			i++
		}
	EndFlags:

		// Parse width (optional)
		if i < end && format[i] >= '0' && format[i] <= '9' {
			start := i
			for i < end && format[i] >= '0' && format[i] <= '9' {
				i++
			}
			num, _ := strconv.Atoi(format[start:i])
			spec.Width = num
		}

		// Parse precision (optional)
		if i < end && format[i] == '.' {
			i++
			start := i
			for i < end && format[i] >= '0' && format[i] <= '9' {
				i++
			}
			if start != i {
				num, _ := strconv.Atoi(format[start:i])
				spec.Precision = num
			}
		}

		// Parse verb
		if i < end {
			verb, size := utf8.DecodeRuneInString(format[i:])
			spec.Verb = verb
			i += size
		}

		specs = append(specs, spec)
	}

	return specs, nil
}
