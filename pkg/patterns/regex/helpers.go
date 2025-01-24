// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package regex

import (
	"github.com/andrerrcosta2/gtools/patterns/regex/tokens"
	"strings"
)

func ph(format string) string {
	return strings.ReplaceAll(format, "%s", string(tokens.PH))
}
