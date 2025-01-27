// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package regex

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/patterns/regex/dictionaries"
	"github.com/andrerrcosta2/gtools/patterns/regex/tokens"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
)

// Dictionary returns a dictionary of regular expression tokens
func Dictionary() str.Dictionary[string, []symbols.Logical] {
	return str.MergeDictionaries[string, []symbols.Logical](dictionaries.Word, tokens.Regexp)
}
