// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package regex

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries"
)

// Trie creates a new OfPatterns from regex.Dictionary.
// It supports both building concurrently using the given semaphore or sequentially.
// If the strict flag is true, the OfPatterns is built in strict mode.
func Trie(strict bool, maxParallelOps, buildInParallel int) (tries.Pattern, error) {
	var dictionary str.Dictionary[string, []symbols.Logical] = Dictionary()
	return tries.OfPatternsFrom(dictionary, strict, maxParallelOps, buildInParallel)
}
