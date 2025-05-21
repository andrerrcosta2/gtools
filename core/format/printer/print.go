// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package printer

import (
	"fmt"
	"io"
)

func Fprint(w io.Writer, s ...any) (n int, err error) {
	return fmt.Fprint(w, s...)
}

func Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

// Print is the same as fmt.Print
// TODO: the idea here is to implement the printer to call a format.Sprintable
// interface for '%v' and '%s' before calling fmt.Stringer
func Print(s ...any) {
	fmt.Print(s...)
}

// Println is the same as fmt.Println
// TODO: the idea here is to implement the printer to call a format.Sprintable
// interface for '%v' and '%s' before calling fmt.Stringer
func Println(s ...any) {
	fmt.Println(s...)
}

// Printf is the same as fmt.Printf
// TODO: the idea here is to implement the printer to call a format.Sprintable
// interface for '%v' and '%s' before calling fmt.Stringer
func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

// Sprintf returns a formatted string
// TODO: the idea here is to implement the printer to call a format.Sprintable
// interface for '%v' and '%s' before calling fmt.Stringer
func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

// Sprintln returns a formatted string
// TODO: the idea here is to implement the printer to call a format.Sprintable
// interface for '%v' and '%s' before calling fmt.Stringer
func Sprintln(s ...any) string {
	return fmt.Sprintln(s...)
}

// Sprint returns a formatted string
// TODO: the idea here is to implement the printer to call a format.Sprintable
// interface for '%v' and '%s' before calling fmt.Stringer
func Sprint(s ...any) string {
	return fmt.Sprint(s...)
}
