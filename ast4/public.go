// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ast4

import (
	"go/ast"
	"go/token"
)

type (
	CommentedNode struct {
		Node    ast.Node
		Comment *ast.CommentGroup
		Pos     token.Position
		Text    string
	}
)
