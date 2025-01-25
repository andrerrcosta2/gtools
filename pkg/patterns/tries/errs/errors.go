// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package errs

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/gtools/gerrors"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
)

var ConsecutivePlaceholders = func(s grammar.Symbol) error { return fmt.Errorf("consecutive placeholders are meaningless: %s\n", s) }
var MismatchedPlaceholders = func(s, p grammar.Symbol) error {
	return fmt.Errorf("mismatched placeholders between pattern '%s' and symbol '%s'\n", s, p)
}
var RootSymbolOnPath = func(s grammar.Symbol) error {
	return fmt.Errorf("invalid path. the Root symbol '%s' is not allowed: %s\n", def.Root, s)
}

var EmptyEntry = errors.New("empty symbol or pattern\n")
var EmptySymbol = errors.New("empty symbol\n")
var PlaceholderSymbol = errors.New("a symbol cannot be composed only by a placeholder\n")

var RemoveNodeWithChildrenNotAllowed = errors.New("remove nodes with children is not allowed on this context\n")

var PlaceholderChildOfPlaceholder = gerrors.Tagged(errors.New("a placeholder can't be a child of another Placeholder\n"), "placeholder-child-of-placeholder")

var PlaceholderChildOfPlaceholderWithKey = func(key string) error {
	return gerrors.Tagged(fmt.Errorf(PlaceholderChildOfPlaceholder.Error()+
		"parent Placeholder key: %s\n", key), "placeholder-child-of-placeholder")
}
