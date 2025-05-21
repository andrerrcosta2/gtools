// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package format

//
//import (
//	"fmt"
//	"reflect"
//	"strconv"
//	"strings"
//	"unicode/utf8"
//)
//
//type buffer []byte
//
//func (b *buffer) write(p []byte) {
//	*b = append(*b, p...)
//}
//
//func (b *buffer) writeString(s string) {
//	*b = append(*b, s...)
//}
//
//func (b *buffer) writeByte(c byte) {
//	*b = append(*b, c)
//}
//
//func (b *buffer) writeRune(r rune) {
//	*b = utf8.AppendRune(*b, r)
//}
//
//func String(v string, verb rune) error {
//	switch verb {
//	case 'v':
//		if p.fmt.sharpV {
//			p.fmt.fmtQ(v)
//		} else {
//			p.fmt.fmtS(v)
//		}
//	case 's':
//		p.fmt.fmtS(v)
//	case 'x':
//		p.fmt.fmtSx(v, ldigits)
//	case 'X':
//		p.fmt.fmtSx(v, udigits)
//	case 'q':
//		p.fmt.fmtQ(v)
//	default:
//		p.badVerb(verb)
//	}
//}
//
//func formatInt(builder *strings.Builder, spec Spec, arg any) error {
//	num, ok := arg.(int)
//	if !ok {
//		return fmt.Errorf("expected int, got %T", arg)
//	}
//	formatValue(builder, spec, strconv.Itoa(num))
//	return nil
//}
//
//func formatFloat(builder *strings.Builder, spec Spec, arg any) error {
//	num, ok := arg.(float64)
//	if !ok {
//		return fmt.Errorf("expected float64, got %T", arg)
//	}
//	formatValue(builder, spec, strconv.FormatFloat(num, 'f', -1, 64))
//	return nil
//}
//
//func formatDefault(builder *strings.Builder, spec Spec, arg any) error {
//	formatValue(builder, spec, fmt.Sprintf("%v", arg))
//	return nil
//}
//
//func formatType(builder *strings.Builder, spec Spec, arg any) error {
//	formatValue(builder, spec, reflect.TypeOf(arg).String())
//	return nil
//}
//
//func formatBool(builder *strings.Builder, spec Spec, arg any) error {
//	b, ok := arg.(bool)
//	if !ok {
//		return fmt.Errorf("expected bool, got %T", arg)
//	}
//	formatValue(builder, spec, strconv.FormatBool(b))
//	return nil
//}
//
//func formatValue(builder *strings.Builder, spec Spec, value string) {
//	padding := spec.Width - utf8.RuneCountInString(value)
//	if padding > 0 {
//		padChar := ' '
//		if spec.PadWithZero {
//			padChar = '0'
//		}
//		if spec.LeftAlign {
//			builder.WriteString(value)
//			builder.WriteString(strings.Repeat(string(padChar), padding))
//		} else {
//			builder.WriteString(strings.Repeat(string(padChar), padding))
//			builder.WriteString(value)
//		}
//	} else {
//		builder.WriteString(value)
//	}
//}
