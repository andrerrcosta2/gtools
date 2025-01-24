// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package symbols

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/gtools/constraints/prim/nums/floats"
	"github.com/andrerrcosta2/gtools/core/gtools/constraints/prim/nums/ints"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"strings"
)

var (
	Empty = Of(``)
	Repl  = Of(`%s`)
)

// Of is a short for grammar.NewSymbol. It returns a grammar.ByteSymbol from a value.
// It takes any type K, and calls grammar.NewSymbol.
// If there is an error creating the symbol it panics.
func Of[T any](s T) grammar.Symbol {
	symbol, err := grammar.NewSymbol(s)
	if err != nil {
		panic(fmt.Sprintf("error creating symbol: %v\n", err))
	}
	return symbol
}

func String(s string) grammar.Symbol {
	return grammar.ByteSymbol(s)
}

func OfPlaceholder(format string, placeholder string) grammar.Symbol {
	// Count how many %s placeholders exist in the format string
	count := strings.Count(format, "%s")

	// Create a slice of the same size filled with the placeholder value
	args := make([]interface{}, count)
	for i := range args {
		args[i] = placeholder
	}

	// Use fmt.Sprintf to replace all "%s" with the placeholders
	processed := fmt.Sprintf(format, args...)
	return grammar.ByteSymbol(processed)
}

// From creates a slice of grammar.ByteSymbol from a given variadic string.
func From(s ...string) []grammar.Symbol {
	if len(s) == 0 {
		return nil
	}
	// Create a new slice with the same length as the input
	out := make([]grammar.Symbol, len(s))
	// Iterate over the input strings and convert them to grammar.ByteSymbol
	for i := range s {
		out[i] = grammar.ByteSymbol(s[i])
	}
	// Return the slice of grammar.ByteSymbol
	return out
}

// Integer returns a grammar.ByteSymbol from a given integer.
//
// Example:
//
//	Int(42) // returns a grammar.ByteSymbol with the value "42"
func Integer(i int) grammar.Symbol {
	return grammar.ByteSymbol(ints.ToBytes(i))
}

// Float returns a grammar.ByteSymbol from a given float64.
//
// Example:
//
//	Float(42.42) // returns a grammar.ByteSymbol with the value "42.42"
func Float(f float64) (grammar.Symbol, error) {
	bytes, err := floats.ToBytes(f)
	return grammar.ByteSymbol(bytes), err
}
