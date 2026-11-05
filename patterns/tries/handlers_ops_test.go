// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
)

func HANDLERS_shouldFindNSymbols(t testlite.LoggerTesting, from nodes.PatternTrie, s dict.Entry[string, symbols.Logical], n int) {
	t.Logf("searchSymbolsByPattern: '%s'\n", s.Key())
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
