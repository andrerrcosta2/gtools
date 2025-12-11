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

func NewFunc(name string, receiver string, body string) *ast.FuncDecl {
	return create.Func(name, receiver, body)
}

func FindNodeByCommentPrefix(fset *token.FileSet, file *ast.File, prefix string) []CommentedNode {
	return ([]CommentedNode)(docs.FindNodeByCommentPrefix(fset, file, prefix))
}
