// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"fmt"

	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
)

func Entry(key string, value symbols.Logical) dict.Entry[string, symbols.Logical] {
	return &patternTrieEntry{
		key:   key,
		value: value,
	}
}

type patternTrieEntry struct {
	key   string
	value symbols.Logical
}

func (e *patternTrieEntry) Key() string {
	return e.key
}

func (e *patternTrieEntry) Value() symbols.Logical {
	return e.value
}

func (e *patternTrieEntry) String() string {
	return fmt.Sprintf("%v: %v", e.key, e.value)
}
