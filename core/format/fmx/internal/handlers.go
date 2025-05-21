// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package internal

//func processFormat(builder *strings.Builder, format string, args []any) error {
//	argIndex := 0
//	for i := 0; i < len(format); {
//		if format[i] == '%' {
//			i++
//			if i >= len(format) {
//				return fmt.Errorf("incomplete format specifier")
//			}
//
//			// Parse the format specifier
//			specifier, consumed, err := parseSpecifier(format[i:])
//			if err != nil {
//				return err
//			}
//			i += consumed
//
//			// Apply the specifier to the current argument
//			if argIndex >= len(args) {
//				return fmt.Errorf("not enough arguments for format string")
//			}
//			if err := format.applySpecifier(builder, specifier, args[argIndex]); err != nil {
//				return err
//			}
//			argIndex++
//		} else {
//			// Append regular characters
//			builder.WriteByte(format[i])
//			i++
//		}
//	}
//	return nil
//}
//
//func parseSpecifier(format string) (printer.Spec, int, error) {
//	var spec format.Spec
//	i := 0
//
//	// Parse flags
//	for i < len(format) && (format[i] == '-' || format[i] == '0') {
//		if format[i] == '-' {
//			spec.LeftAlign = true
//		} else if format[i] == '0' {
//			spec.PadWithZero = true
//		}
//		i++
//	}
//
//	// Parse width
//	if i < len(format) && format[i] >= '0' && format[i] <= '9' {
//		width, consumed, err := parseNumber(format[i:])
//		if err != nil {
//			return format.Spec{}, 0, err
//		}
//		spec.Width = width
//		i += consumed
//	}
//
//	// Parse precision
//	if i < len(format) && format[i] == '.' {
//		i++
//		if i >= len(format) || format[i] < '0' || format[i] > '9' {
//			return format.Spec{}, 0, fmt.Errorf("invalid precision")
//		}
//		precision, consumed, err := parseNumber(format[i:])
//		if err != nil {
//			return format.Spec{}, 0, err
//		}
//		spec.Precision = &precision
//		i += consumed
//	}
//
//	// Parse verb
//	if i >= len(format) || !isValidVerb(format[i]) {
//		return format.Spec{}, 0, fmt.Errorf("invalid verb")
//	}
//	spec.Verb = rune(format[i])
//	i++
//
//	return spec, i, nil
//}
//
//func parseNumber(format string) (int, int, error) {
//	num := 0
//	i := 0
//	for i < len(format) && format[i] >= '0' && format[i] <= '9' {
//		num = num*10 + int(format[i]-'0')
//		i++
//	}
//	if i == 0 {
//		return 0, 0, fmt.Errorf("expected number")
//	}
//	return num, i, nil
//}
//
//func isValidVerb(c byte) bool {
//	switch c {
//	case 's', 'd', 'f', 'v', 'T', 't':
//		return true
//	default:
//		return false
//	}
//}
