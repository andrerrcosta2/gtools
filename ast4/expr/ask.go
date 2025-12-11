// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package expr

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/ask"
)

func IsSimilar(a, b ast.Expr) bool {
	return ask.IsSimilarExpr(a, b)
}

func IsSimilarArrayType(a, b *ast.ArrayType) bool {
	return ask.IsSimilarArrayType(a, b)
}

func IsSimilarBasicLit(a, b *ast.BasicLit) bool {
	return ask.IsSimilarBasicLit(a, b)
}

func IsSimilarBinaryExpr(a, b *ast.BinaryExpr) bool {
	return ask.IsSimilarBinaryExpr(a, b)
}

func IsSimilarCallExpr(a, b *ast.CallExpr) bool {
	return ask.IsSimilarCallExpr(a, b)
}

func IsSimilarChanType(a, b *ast.ChanType) bool {
	return ask.IsSimilarChanType(a, b)
}

func IsSimilarCompositeLit(a, b *ast.CompositeLit) bool {
	return ask.IsSimilarCompositeLit(a, b)
}

func IsSimilarEllipsis(a, b *ast.Ellipsis) bool {
	return ask.IsSimilarEllipsis(a, b)
}

func IsSimilarFuncLit(a, b *ast.FuncLit) bool {
	return ask.IsSimilarFuncLit(a, b)
}

func IsSimilarFuncType(a, b *ast.FuncType) bool {
	return ask.IsSimilarFuncType(a, b)
}

func IsSimilarIdent(a, b *ast.Ident) bool {
	return ask.IsSimilarIdent(a, b)
}

func IsSimilarIndexExpr(a, b *ast.IndexExpr) bool {
	return ask.IsSimilarIndexExpr(a, b)
}

func IsSimilarIndexListExpr(a, b *ast.IndexListExpr) bool {
	return ask.IsSimilarIndexListExpr(a, b)
}

func IsSimilarInterfaceType(a, b *ast.InterfaceType) bool {
	return ask.IsSimilarInterfaceType(a, b)
}

func IsSimilarKeyValueExpr(a, b *ast.KeyValueExpr) bool {
	return ask.IsSimilarKeyValueExpr(a, b)
}

func IsSimilarMapType(a, b *ast.MapType) bool {
	return ask.IsSimilarMapType(a, b)
}

func IsSimilarParenExpr(a, b *ast.ParenExpr) bool {
	return ask.IsSimilarParenExpr(a, b)
}

func IsSimilarSelectorExpr(a, b *ast.SelectorExpr) bool {
	return ask.IsSimilarSelectorExpr(a, b)
}

func IsSimilarSliceExpr(a, b *ast.SliceExpr) bool {
	return ask.IsSimilarSliceExpr(a, b)
}

func IsSimilarStarExpr(a, b *ast.StarExpr) bool {
	return ask.IsSimilarStarExpr(a, b)
}

func IsSimilarStructType(a, b *ast.StructType) bool {
	return ask.IsSimilarStructType(a, b)
}

func IsSimilarTypeAssertExpr(a, b *ast.TypeAssertExpr) bool {
	return ask.IsSimilarTypeAssertExpr(a, b)
}

func IsSimilarUnaryExpr(a, b *ast.UnaryExpr) bool {
	return ask.IsSimilarUnaryExpr(a, b)
}

func IsExpr(expr ast.Expr) bool {
	return ask.IsExpr(expr)
}

func IsArrayType(expr ast.Expr) bool {
	return ask.IsArrayType(expr)
}

func IsBadExpr(expr ast.Expr) bool {
	return ask.IsBadExpr(expr)
}

func IsBasicLit(expr ast.Expr) bool {
	return ask.IsBasicLit(expr)
}

func IsCallExpr(expr ast.Expr) bool {
	return ask.IsCallExpr(expr)
}

func IsChanType(expr ast.Expr) bool {
	return ask.IsChanType(expr)
}

func IsCompositeLit(expr ast.Expr) bool {
	return ask.IsCompositeLit(expr)
}

func IsEllipsis(expr ast.Expr) bool {
	return ask.IsEllipsis(expr)
}

func IsFuncLit(expr ast.Expr) bool {
	return ask.IsFuncLit(expr)
}

func IsFuncType(expr ast.Expr) bool {
	return ask.IsFuncType(expr)
}

func IsIdent(expr ast.Expr) bool {
	return ask.IsIdent(expr)
}

func IsInterfaceType(expr ast.Expr) bool {
	return ask.IsInterfaceType(expr)
}

func IsIndexExpr(expr ast.Expr) bool {
	return ask.IsIndexExpr(expr)
}

func IsIndexListExpr(expr ast.Expr) bool {
	return ask.IsIndexListExpr(expr)
}

func IsKeyValueExpr(expr ast.Expr) bool {
	return ask.IsKeyValueExpr(expr)
}

func IsMapType(expr ast.Expr) bool {
	return ask.IsMapType(expr)
}

func IsParenExpr(expr ast.Expr) bool {
	return ask.IsParenExpr(expr)
}

func IsSelectorExpr(expr ast.Expr) bool {
	return ask.IsSelectorExpr(expr)
}

func IsSliceExpr(expr ast.Expr) bool {
	return ask.IsSliceExpr(expr)
}

func IsStarExpr(expr ast.Expr) bool {
	return ask.IsStarExpr(expr)
}

func IsStructType(expr ast.Expr) bool {
	return ask.IsStructType(expr)
}

func IsTypeAssertExpr(expr ast.Expr) bool {
	return ask.IsTypeAssertExpr(expr)
}

func IsUnaryExpr(expr ast.Expr) bool {
	return ask.IsUnaryExpr(expr)
}
