// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fmx

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/custom"
	"io"
)

type Sprinter interface {
	Sprint(tab indent.Indentor) string
}

func Printer() custom.Printer {
	return &fmx{}
}

type fmx struct {
}

func (p *fmx) Fprint(w io.Writer, s ...any) (n int, err error) {
	return fmt.Fprint(w, s...)
}

func (p *fmx) Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

func (p *fmx) Fprintln(w io.Writer, s ...any) (n int, err error) {
	return fmt.Fprintln(w, s...)
}

func (p *fmx) Print(style custom.Style, s ...any) {
	Custom(style, s...)
}

func (p *fmx) Println(style custom.Style, s ...any) {
	Customln(style, s...)
}

func (p *fmx) Printf(style custom.Style, format string, a ...any) {
	Customf(style, format, a...)
}

func (p *fmx) Sprint(style custom.Style, s ...any) string {
	return SCustom(style, s...)
}

func (p *fmx) Sprintln(style custom.Style, s ...any) string {
	return SCustomln(style, s...)
}

func (p *fmx) Sprintf(style custom.Style, format string, s ...any) string {
	return SCustomf(style, format, s...)
}

var _ custom.Printer = (*fmx)(nil)

func Native() custom.Printer {
	return &native{}
}

type native struct {
}

func (p *native) Fprint(w io.Writer, s ...any) (n int, err error) {
	return fmt.Fprint(w, s...)
}

func (p *native) Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

func (p *native) Fprintln(w io.Writer, s ...any) (n int, err error) {
	return fmt.Fprintln(w, s...)
}

func (p *native) Print(_ custom.Style, s ...any) {
	fmt.Print(s...)
}

func (p *native) Println(_ custom.Style, s ...any) {
	fmt.Println(s...)
}

func (p *native) Printf(_ custom.Style, format string, a ...any) {
	fmt.Printf(format, a...)
}

func (p *native) Sprint(_ custom.Style, s ...any) string {
	return fmt.Sprint(s...)
}

func (p *native) Sprintln(_ custom.Style, s ...any) string {
	return fmt.Sprintln(s...)
}

func (p *native) Sprintf(_ custom.Style, format string, s ...any) string {
	return fmt.Sprintf(format, s...)
}

var _ custom.Printer = (*native)(nil)
