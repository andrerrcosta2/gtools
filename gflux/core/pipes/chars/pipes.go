// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chars

import "strings"

func SplitPrefix(s string, prefix rune) (rune, string) {
	if strings.HasPrefix(s, string(prefix)) {
		return prefix, strings.TrimPrefix(s, string(prefix))
	}
	return 0, s
}
