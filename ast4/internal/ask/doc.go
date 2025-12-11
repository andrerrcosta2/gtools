// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ask

import (
	"go/ast"
)

func IsSimilarCommentGroup(a, b *ast.CommentGroup) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a.List) != len(b.List) {
		return false
	}
	for i, com := range a.List {
		if !IsSimilarComment(com, b.List[i]) {
			return false
		}
	}
	return true
}

func IsSimilarComment(a, b *ast.Comment) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Text == b.Text
}
