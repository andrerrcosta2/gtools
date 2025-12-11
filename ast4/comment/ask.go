// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comment

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/ask"
)

func IsSimilarCommentGroup(a, b *ast.CommentGroup) bool {
	return ask.IsSimilarCommentGroup(a, b)
}

func IsSimilarComment(a, b *ast.Comment) bool {
	return ask.IsSimilarComment(a, b)
}
