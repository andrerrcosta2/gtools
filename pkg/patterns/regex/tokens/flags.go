// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tokens

import (
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
)

type Flag struct {
	symbols.Logical
}

var (
	// Global g: Finds all matches within the input string, instead of stopping after the first match.
	Global = Flag{Logical: symbols.OpenOf(grammar.ByteSymbol([]byte{'.'}), true)}
	// IgnoreCase i: Performs case-insensitive matching.
	IgnoreCase = Flag{Logical: symbols.OpenOf(grammar.ByteSymbol([]byte{'i'}), true)}
	// Multiline m: Enables ^ and $ to match the beginning and end of each line, respectively.
	Multiline = Flag{Logical: symbols.OpenOf(grammar.ByteSymbol([]byte{'m'}), true)}
	// Singleline s: Disables ^ and $ from matching the beginning and end of each line, respectively.
)
