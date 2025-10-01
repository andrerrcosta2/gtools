// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fmx

import (
	"github.com/andrerrcosta2/gtools/core/format/custom/term"
	"github.com/andrerrcosta2/gtools/core/format/printer"
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

func Cyan(s ...any) {
	printer.Print(term.Cyan.Style(printer.Sprint(s...)))
}

func Cyanf(format string, args ...any) {
	printer.Print(term.Cyan.Stylef(format, args...))
}

func Cyanfc(format string, args ...any) {
	Call(term.Cyan.Stylef(format, args...))
}

func Cyanln(s ...any) {
	printer.Print(term.Cyan.Style(printer.Sprintln(s...)))
}

func SCyan(s ...any) string {
	return term.Cyan.Style(printer.Sprint(s...))
}

func SCyanf(format string, args ...any) string {
	return term.Cyan.Stylef(format, args...)
}

func SCyanln(s ...any) string {
	return printer.Sprint(term.Cyan.Style(printer.Sprintln(s...)))
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

func Purple(s ...any) {
	printer.Print(term.BoldBrightPurple.Style(printer.Sprint(s...)))
}

func Purplef(format string, args ...any) {
	printer.Print(term.BoldBrightPurple.Stylef(format, args...))
}

func Purplefc(format string, args ...any) {
	Call(term.BoldBrightPurple.Stylef(format, args...))
}

func Purpleln(s ...any) {
	printer.Print(term.BoldBrightPurple.Style(printer.Sprintln(s...)))
}

func SPurple(s ...any) string {
	return term.BoldBrightPurple.Style(printer.Sprint(s...))
}

func SPurplef(format string, args ...any) string {
	return term.BoldBrightPurple.Stylef(format, args...)
}

func SPurpleln(s ...any) string {
	return printer.Sprint(term.BoldBrightPurple.Style(printer.Sprintln(s...)))
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

func Yellow(s ...any) {
	printer.Print(term.Yellow.Style(printer.Sprint(s...)))
}

func Yellowf(format string, args ...any) {
	printer.Print(term.Yellow.Stylef(format, args...))
}

func Yellowfc(format string, args ...any) {
	Call(term.Yellow.Stylef(format, args...))
}

func Yellowln(s ...any) {
	printer.Print(term.Yellow.Style(printer.Sprintln(s...)))
}

func SYellow(s ...any) string {
	return term.Yellow.Style(s...)
}

func SYellowf(format string, args ...any) string {
	return term.Yellow.Stylef(format, args...)
}

func SYellowln(s ...any) string {
	return term.Yellow.Styleln(s...)
}
