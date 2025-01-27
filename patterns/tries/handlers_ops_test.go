// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
)

func HANDLERS_shouldFindNSymbols(t gtests.Loggable, from nodes.PatternTrie, s str.Entry[string, symbols.Logical], n int) {
	t.StackLogf("searchSymbolsByPattern: '%s'\n", s.Key())
	// Assert only one entry is found
	ss, ok := searchSymbolsByPattern(from, s.Key(), 6)
	if !ok {
		t.Errorf("expected to find '%v' but was not found\n", s.Key())
	}
	if len(ss) != n {
		t.Errorf("expected '%d' value(s) for '%v' but found %v\n", n, s.Key(), ss)
	}
	if !s.Value().Equal(ss[0]) {
		t.Errorf("expected %v but found %v\n", s.Value(), ss[0])
	}
}
