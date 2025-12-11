// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package file

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/ask"
)

func IsSimilar(a, b *ast.File) bool {
	return ask.IsSimilarFile(a, b)
}
