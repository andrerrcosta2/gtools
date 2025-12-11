// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
)

func nilDiff(t indent.Indentor, a, b ast.Node, kind, typ string) (hasNil, equals bool, diff string) {
	if a == nil && b == nil {
		return true, true, diff
	}
	if a == nil {
		return true, false, differs.NilReceived(t, kind, typ)
	}
	if b == nil {
		return true, false, differs.NilExpected(t, kind, typ)
	}
	return
}
