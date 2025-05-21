// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package custom

import "io"

type Printer interface {
	Writer
	Print(style Style, s ...any)
	Printf(style Style, format string, a ...any)
	Println(style Style, s ...any)
	Sprint(style Style, s ...any) string
	Sprintf(style Style, format string, s ...any) string
	Sprintln(style Style, s ...any) string
}

type Writer interface {
	Fprint(w io.Writer, a ...any) (n int, err error)
	Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	Fprintln(w io.Writer, a ...any) (n int, err error)
}

type Buffer interface {
	Append(b []byte, a ...any) []byte
	Appendf(b []byte, format string, a ...any) []byte
	Appendln(b []byte, a ...any) []byte
}
