// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fmx

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/custom"
	"github.com/andrerrcosta2/gtools/core/format/custom/term"
	"github.com/andrerrcosta2/gtools/core/format/printer"
	"strings"
)

func Blueln(s ...any) {
	printer.Print(term.Blue.Style(printer.Sprintln(s...)))
}

func Blue(s ...any) {
	printer.Print(term.Blue.Style(printer.Sprintln(s...)))
}

func Bluef(format string, args ...any) {
	printer.Print(term.Blue.Stylef(format, args...))
}

func Bluefc(format string, args ...any) {
	Call(term.Blue.Stylef(format, args...))
}

func SBlue(s ...any) string {
	return term.Blue.Style(printer.Sprint(s...))
}

func SBluef(format string, args ...any) string {
	return term.Blue.Stylef(format, args...)
}

func SBlueln(s ...any) string {
	return printer.Sprintln(term.Blue.Style(printer.Sprintln(s...)))
}

func Bold(s ...any) {
	printer.Print(term.Bold.Style(printer.Sprint(s...)))
}

func Boldf(format string, args ...any) {
	printer.Print(term.Bold.Stylef(format, args...))
}

func Boldfc(format string, args ...any) {
	Call(term.Bold.Stylef(format, args...))
}

func Boldln(s ...any) {
	printer.Print(term.Bold.Style(printer.Sprintln(s...)))
}

func SBold(s ...any) string {
	return term.Bold.Style(printer.Sprint(s...))
}

func SBoldf(format string, args ...any) string {
	value := printer.Sprintf(format, args...)
	if len(value) == 0 {
		return ""
	}
	return term.Bold.Stylef(format, args...)
}

func SBoldln(s ...any) string {
	return printer.Sprintln(term.Bold.Style(printer.Sprintln(s...)))
}

func Code(tab indent.Tab, line int, code string, bg term.Bg) {
	printer.Print(SCode(tab, line, code, bg))
}

func SCode(tab indent.Tab, line int, code string, bg term.Bg) string {
	// Split the code into lines
	lines := strings.Split(code, "\n")

	// Build the formatted code block
	var result strings.Builder
	for i, codeLine := range lines {
		// Format the line number without leading spaces
		lineNumber := printer.Sprintf("%d", line+i)           // No padding here
		coloredLineNumber := bg.Style(" " + lineNumber + " ") // Add 1 space before and after

		// Add padding manually after applying the background color
		paddedLineNumber := printer.Sprintf("%6s", coloredLineNumber)

		// Append the indented code line
		result.WriteString(paddedLineNumber + tab.Sprint(codeLine) + "\n")
	}

	return result.String()
}

func Color(color term.Color, s ...any) {
	printer.Print(color.Style(printer.Sprint(s...)))
}

func Colorf(color term.Color, format string, args ...any) {
	printer.Print(color.Stylef(format, args...))
}

func Colorfc(color term.Color, format string, args ...any) {
	Call(color.Stylef(format, args...))
}

func Colorln(color term.Color, s ...any) {
	printer.Print(color.Style(printer.Sprintln(s...)))
}

func SColor(color term.Color, s ...any) string {
	return color.Style(printer.Sprint(s...))
}

func SColorf(color term.Color, format string, args ...any) string {
	return color.Stylef(format, args...)
}

func SColorln(color term.Color, s ...any) string {
	return printer.Sprintln(color.Style(printer.Sprintln(s...)))
}

func Custom(style custom.Style, s ...any) {
	printer.Print(style.Style(printer.Sprint(s...)))
}

func Customf(style custom.Style, format string, args ...any) {
	printer.Print(style.Stylef(format, args...))
}

func Customfc(style custom.Style, format string, args ...any) {
	Call(style.Stylef(format, args...))
}

func Customln(style custom.Style, s ...any) {
	printer.Print(style.Style(printer.Sprint(s, " ")))
}

func SCustom(style custom.Style, s ...any) string {
	return style.Style(printer.Sprint(s...))
}

func SCustomf(style custom.Style, format string, args ...any) string {
	return style.Stylef(format, args...)
}

func SCustomln(style custom.Style, s ...any) string {
	return printer.Sprint(style.Style(printer.Sprintln(s...)))
}

func Errorf(format string, args ...any) error {
	return errors.New(printer.Sprintf(format, args...))
}

func Green(s ...any) {
	printer.Print(term.Green.Style(printer.Sprint(s...)))
}

func Greenf(format string, args ...any) {
	printer.Print(term.Green.Stylef(format, args...))
}

func Greenfc(format string, args ...any) {
	Call(term.Green.Stylef(format, args...))
}

func Greenln(s ...any) {
	printer.Print(term.Green.Style(printer.Sprintln(s...)))
}

func SGreen(s ...any) string {
	return term.Green.Style(printer.Sprint(s...))
}

func SGreenf(format string, args ...any) string {
	return term.Green.Stylef(format, args...)
}

func SGreenln(s ...any) string {
	return printer.Sprint(term.Green.Style(printer.Sprintln(s...)))
}

func Italic(s ...any) {
	printer.Print(term.Italic.Style(printer.Sprint(s...)))
}

func Italicf(format string, args ...any) {
	printer.Print(term.Italic.Stylef(format, args...))
}

func Italicfc(format string, args ...any) {
	Call(term.Italic.Stylef(format, args...))
}

func Italicln(s ...any) {
	printer.Print(term.Italic.Style(printer.Sprintln(s...)))
}

func SItalic(s ...any) string {
	return term.Italic.Style(printer.Sprint(s...))
}

func SItalicf(format string, args ...any) string {
	return term.Italic.Stylef(format, args...)
}

func SItalicln(s ...any) string {
	return printer.Sprint(term.Italic.Style(printer.Sprintln(s...)))
}

func Print(s ...any) {
	printer.Print(s...)
}

func Printf(format string, s ...any) {
	printer.Printf(format, s...)
}

func Printfln(format string, s ...any) {
	printer.Printf(format+"\n", s...)
}

func Println(s ...any) {
	printer.Println(s...)
}

func Printlnf(format string, s ...any) {
	printer.Printf("\n"+format, s...)
}

func Sprint(s ...any) string {
	return printer.Sprint(s...)
}

func Sprintf(format string, s ...any) string {
	return printer.Sprintf(format, s...)
}

func Sprintln(s ...any) string {
	return printer.Sprintln(s...)
}

func CPrintln(s ...any) {
	Call(printer.Sprintln(s...))
}

func Red(s ...any) {
	printer.Print(term.Red.Style(printer.Sprint(s...)))
}

func Redf(format string, args ...any) {
	printer.Print(term.Red.Stylef(format, args...))
}

func Redfc(format string, args ...any) {
	Call(term.Red.Stylef(format, args...))
}

func Redln(s ...any) {
	printer.Print(term.Red.Style(printer.Sprintln(s...)))
}

func SRed(s ...any) string {
	return term.Red.Style(s...)
}

func SRedf(format string, args ...any) string {
	return term.Red.Stylef(format, args...)
}

func SRedln(s ...any) string {
	return term.Red.Styleln(s...)
}

func Spaced(s ...any) {
	printer.Print("\n\n" + printer.Sprint(s...) + "\n\n")
}

func Spacedf(format string, args ...any) {
	printer.Printf("\n\n"+format+"\n\n", args...)
}

func Spacedfc(format string, args ...any) {
	Callf("\n\n"+format+"\n\n", args...)
}

func Spacedln(s ...any) {
	printer.Println("\n\n" + printer.Sprint(s...) + "\n\n")
}

func SSpaced(s ...any) string {
	return "\n\n" + printer.Sprint(s...) + "\n\n"
}

func SSpacedf(format string, args ...any) string {
	return "\n\n" + printer.Sprintf(format, args...) + "\n\n"
}

func SSpacedfc(format string, args ...any) string {
	return SCallf("\n\n"+format+"\n\n", args...)
}

func SSpacedln(s ...any) string {
	return "\n\n" + printer.Sprintln(printer.Sprint(s...)) + "\n\n"
}

func Strike(s ...any) {
	printer.Print(term.Strikethrough.Style(printer.Sprint(s...)))
}

func Strikef(format string, args ...any) {
	printer.Print(term.Strikethrough.Stylef(format, args...))
}

func Strikefc(format string, args ...any) {
	term.Strikethrough.Stylef(format, args...)
}

func Strikeln(s ...any) {
	printer.Print(term.Strikethrough.Style(printer.Sprintln(s...)))
}

func SStrike(s ...any) string {
	return term.Strikethrough.Style(printer.Sprint(s...))
}

func SStrikef(format string, args ...any) string {
	return term.Strikethrough.Stylef(format, args...)
}

func SStrikeln(s ...any) string {
	return printer.Sprint(term.Strikethrough.Style(printer.Sprintln(s...)))
}

func Strip(message, s string) {
	printer.Printf("%s%q", message, s)
}

func Sstrip(message, s string) string {
	return printer.Sprintf("%s%q", message, s)
}

func Underline(s ...any) {
	printer.Print(term.Underline.Style(printer.Sprint(s...)))
}

func Underlinef(format string, args ...any) {
	printer.Print(term.Underline.Stylef(format, args...))
}

func Underlinefc(format string, args ...any) {
	Call(term.Underline.Stylef(format, args...))
}

func Underlineln(s ...any) {
	printer.Print(term.Underline.Style(printer.Sprintln(s...)))
}

func SUnderline(s ...any) string {
	return term.Underline.Style(printer.Sprint(s...))
}

func SUnderlinef(format string, args ...any) string {
	return term.Underline.Stylef(format, args...)
}

func SUnderlineln(s ...any) string {
	return printer.Sprint(term.Underline.Style(printer.Sprintln(s...)))
}
