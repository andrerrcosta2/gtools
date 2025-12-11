// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
)

func Expr(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.ArrayType:
		return "*ast.ArrayType"
	case *ast.BadExpr:
		return "*ast.BadExpr"
	case *ast.BasicLit:
		return "*ast.BasicLit"
	case *ast.BinaryExpr:
		return "*ast.BinaryExpr"
	case *ast.CallExpr:
		return "*ast.CallExpr"
	case *ast.ChanType:
		return "*ast.ChanType"
	case *ast.CompositeLit:
		return "*ast.CompositeLit"
	case *ast.Ellipsis:
		return "*ast.Ellipsis"
	case *ast.FuncLit:
		return "*ast.FuncLit"
	case *ast.FuncType:
		return "*ast.FuncType"
	case *ast.Ident:
		return "*ast.Ident"
	case *ast.IndexExpr:
		return "*ast.IndexExpr"
	case *ast.IndexListExpr:
		return "*ast.IndexListExpr"
	case *ast.InterfaceType:
		return "*ast.InterfaceType"
	case *ast.KeyValueExpr:
		return "*ast.KeyValueExpr"
	case *ast.MapType:
		return "*ast.MapType"
	case *ast.ParenExpr:
		return "*ast.ParenExpr"
	case *ast.SelectorExpr:
		return "*ast.SelectorExpr"
	case *ast.SliceExpr:
		return "*ast.SliceExpr"
	case *ast.StarExpr:
		return "*ast.StarExpr"
	case *ast.StructType:
		return "*ast.StructType"
	case *ast.TypeAssertExpr:
		return "*ast.TypeAssertExpr"
	case *ast.UnaryExpr:
		return "*ast.UnaryExpr"
	default:
		return std.ExprTypeError(std.Zero, e)
	}
}
