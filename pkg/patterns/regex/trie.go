// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package regex

import (
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/patterns/tries"
)

// Trie creates a new OfPatterns from regex.Dictionary.
// It supports both building concurrently using the given semaphore or sequentially.
// If the strict flag is true, the OfPatterns is built in strict mode.
func Trie(semaphore gtools.Semaphore, strict bool) (tries.Pattern, error) {
	return tries.OfPatternsFrom(Dictionary(), semaphore, strict)
}
