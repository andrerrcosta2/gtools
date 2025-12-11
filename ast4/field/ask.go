// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package field

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/ask"
)

func IsSimilar(a, b *ast.Field) bool {
	return ask.IsSimilarField(a, b)
}

func IsSimilarList(a, b *ast.FieldList) bool {
	return ask.IsSimilarFieldList(a, b)
}
