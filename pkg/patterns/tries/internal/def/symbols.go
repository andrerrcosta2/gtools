// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package def

import (
	"github.com/andrerrcosta2/gtools/patterns/grammar"
)

var (
	Root        = grammar.ByteSymbol(`\x01`)
	Placeholder = grammar.ByteSymbol(`\x02`)
)
