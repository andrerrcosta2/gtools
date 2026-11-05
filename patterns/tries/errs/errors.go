// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package errs

import (
	"errors"
	"fmt"

	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
)

var (
	ConsecutivePlaceholders = func(s grammar.Symbol) error { return fmt.Errorf("consecutive placeholders are meaningless: %s\n", s) }
	MismatchedPlaceholders  = func(s, p grammar.Symbol) error {
		return fmt.Errorf("mismatched placeholders between pattern '%s' and symbol '%s'\n", s, p)
	}
	RootSymbolOnPath = func(s grammar.Symbol) error {
		return fmt.Errorf("invalid path. the Root symbol '%s' is not allowed: %s\n", def.Root, s)
	}

	EmptyEntry        = errors.New("empty symbol or pattern\n")
	EmptySymbol       = errors.New("empty symbol\n")
	PlaceholderSymbol = errors.New("a symbol cannot be composed only by a placeholder\n")

	RemoveNodeWithChildrenNotAllowed = errors.New("remove node with children is not allowed on this context\n")

	PlaceholderChildOfPlaceholder = errors.New("a placeholder can't be a child of another Placeholder\n")

	PlaceholderChildOfPlaceholderWithKey = func(key string) error {
		return fmt.Errorf(PlaceholderChildOfPlaceholder.Error()+"parent Placeholder key: %s\n", key)
	}
)
