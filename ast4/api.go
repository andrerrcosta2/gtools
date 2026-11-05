// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ast4

import (
	"go/ast"
	"go/token"

	"github.com/andrerrcosta2/gtools/ast4/internal/create"
	"github.com/andrerrcosta2/gtools/ast4/internal/docs"
	"github.com/andrerrcosta2/gtools/ast4/internal/xtr"
)

func FindStructs(file *ast.File) []*ast.TypeSpec {
	return xtr.Structs(file)
}

func FindFields(structNode *ast.StructType) []*ast.Field {
	return xtr.Fields(structNode)
}

func FindMethods(file *ast.File) []*ast.FuncDecl {
	return xtr.Methods(file)
}

func NewField(name string, typeName string) *ast.Field {
	return create.Field(name, typeName)
}

func NewFunction(fn string) (*ast.FuncDecl, error) {
	return create.Function(fn)
}

// RangeCommentedNodesByPrefix ranges through the file and calls the given function for each node
// that has a comment starting with the given prefix
func RangeCommentedNodesByPrefix(
	fset *token.FileSet,
	file *ast.File,
	prefix string,
	fn func(cn CommentedNode) bool,
) {
	docs.RangeNodeByCommentPrefix(fset, file, prefix, func(n ast.Node, doc *ast.CommentGroup, pos token.Position, text string) bool {
		return fn(CommentedNode{
			Node:    n,
			Comment: doc,
			Pos:     pos,
			Text:    text,
		})
	})
}

// FindCommentedNodesByPrefix returns all nodes that have a comment starting with the given prefix
func FindCommentedNodesByPrefix(fset *token.FileSet, file *ast.File, prefix string) []CommentedNode {
	var results []CommentedNode

	RangeCommentedNodesByPrefix(fset, file, prefix, func(cn CommentedNode) bool {
		results = append(results, cn)
		return true
	})

	return results
}
