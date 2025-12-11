// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ask

import (
	"go/ast"
	"log"
)

func IsSimilarExpr(a, b ast.Expr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	switch expr := a.(type) {
	case *ast.ArrayType:
		other, ok := b.(*ast.ArrayType)
		if !ok {
			return false
		}
		return IsSimilarArrayType(expr, other)
	case *ast.BadExpr:
		_, ok := b.(*ast.BadExpr)
		return ok
	case *ast.BasicLit:
		other, ok := b.(*ast.BasicLit)
		if !ok {
			return false
		}
		return IsSimilarBasicLit(expr, other)
	case *ast.BinaryExpr:
		other, ok := b.(*ast.BinaryExpr)
		if !ok {
			return false
		}
		return IsSimilarBinaryExpr(expr, other)
	case *ast.CallExpr:
		other, ok := b.(*ast.CallExpr)
		if !ok {
			return false
		}
		return IsSimilarCallExpr(expr, other)
	case *ast.ChanType:
		other, ok := b.(*ast.ChanType)
		if !ok {
			return false
		}
		return IsSimilarChanType(expr, other)
	case *ast.CompositeLit:
		other, ok := b.(*ast.CompositeLit)
		if !ok {
			return false
		}
		return IsSimilarCompositeLit(expr, other)
	case *ast.Ellipsis:
		other, ok := b.(*ast.Ellipsis)
		if !ok {
			return false
		}
		return IsSimilarEllipsis(expr, other)
	case *ast.FuncLit:
		other, ok := b.(*ast.FuncLit)
		if !ok {
			return false
		}
		return IsSimilarFuncLit(expr, other)
	case *ast.FuncType:
		other, ok := b.(*ast.FuncType)
		if !ok {
			return false
		}
		return IsSimilarFuncType(expr, other)
	case *ast.Ident:
		other, ok := b.(*ast.Ident)
		if !ok {
			return false
		}
		return IsSimilarIdent(expr, other)
	case *ast.IndexExpr:
		other, ok := b.(*ast.IndexExpr)
		if !ok {
			return false
		}
		return IsSimilarIndexExpr(expr, other)
	case *ast.IndexListExpr:
		other, ok := b.(*ast.IndexListExpr)
		if !ok {
			return false
		}
		return IsSimilarIndexListExpr(expr, other)
	case *ast.InterfaceType:
		other, ok := b.(*ast.InterfaceType)
		if !ok {
			return false
		}
		return IsSimilarInterfaceType(expr, other)
	case *ast.KeyValueExpr:
		other, ok := b.(*ast.KeyValueExpr)
		if !ok {
			return false
		}
		return IsSimilarKeyValueExpr(expr, other)
	case *ast.MapType:
		other, ok := b.(*ast.MapType)
		if !ok {
			return false
		}
		return IsSimilarMapType(expr, other)
	case *ast.ParenExpr:
		other, ok := b.(*ast.ParenExpr)
		if !ok {
			return false
		}
		return IsSimilarParenExpr(expr, other)
	case *ast.SelectorExpr:
		other, ok := b.(*ast.SelectorExpr)
		if !ok {
			return false
		}
		return IsSimilarSelectorExpr(expr, other)
	case *ast.SliceExpr:
		other, ok := b.(*ast.SliceExpr)
		if !ok {
			return false
		}
		return IsSimilarSliceExpr(expr, other)
	case *ast.StarExpr:
		other, ok := b.(*ast.StarExpr)
		if !ok {
			return false
		}
		return IsSimilarStarExpr(expr, other)
	case *ast.StructType:
		other, ok := b.(*ast.StructType)
		if !ok {
			return false
		}
		return IsSimilarStructType(expr, other)
	case *ast.TypeAssertExpr:
		other, ok := b.(*ast.TypeAssertExpr)
		if !ok {
			return false
		}
		return IsSimilarTypeAssertExpr(expr, other)
	case *ast.UnaryExpr:
		other, ok := b.(*ast.UnaryExpr)
		if !ok {
			return false
		}
		return IsSimilarUnaryExpr(expr, other)
	default:
		return false
	}
}

func IsSimilarArrayType(a, b *ast.ArrayType) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.Elt, b.Elt) && IsSimilarExpr(a.Len, b.Len)
}

func IsSimilarBasicLit(a, b *ast.BasicLit) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Kind == b.Kind && a.Value == b.Value
}

func IsSimilarBinaryExpr(a, b *ast.BinaryExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.X, b.X) && IsSimilarExpr(a.Y, b.Y) && a.Op == b.Op
}

func IsSimilarCallExpr(a, b *ast.CallExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a.Args) != len(b.Args) {
		return false
	}
	for i := range a.Args {
		if !IsSimilarExpr(a.Args[i], b.Args[i]) {
			return false
		}
	}
	return IsSimilarExpr(a.Fun, b.Fun)
}

func IsSimilarChanType(a, b *ast.ChanType) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Dir == b.Dir && IsSimilarExpr(a.Value, b.Value)
}

func IsSimilarCompositeLit(a, b *ast.CompositeLit) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	if len(a.Elts) != len(b.Elts) {
		return false
	}
	for i := range a.Elts {
		if !IsSimilarExpr(a.Elts[i], b.Elts[i]) {
			return false
		}
	}
	return IsSimilarExpr(a.Type, b.Type)
}

func IsSimilarEllipsis(a, b *ast.Ellipsis) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.Elt, b.Elt)
}

func IsSimilarFuncLit(a, b *ast.FuncLit) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarFuncType(a.Type, b.Type) && IsSimilarBlockStmt(a.Body, b.Body)
}

func IsSimilarFuncType(a, b *ast.FuncType) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	tp := IsSimilarFieldList(a.TypeParams, b.TypeParams)
	log.Printf("Is Similar Type Params: %v\n", tp)
	pr := IsSimilarFieldList(a.Params, b.Params)
	log.Printf("Is Similar Params: %v\n", pr)
	res := IsSimilarFieldList(a.Results, b.Results)
	log.Printf("Is Similar Results: %v\n", res)
	return tp && pr && res
}

func IsSimilarIdent(a, b *ast.Ident) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Name == b.Name
}

func IsSimilarIndexExpr(a, b *ast.IndexExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.X, b.X) && IsSimilarExpr(a.Index, b.Index)
}

func IsSimilarIndexListExpr(a, b *ast.IndexListExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a.Indices) != len(b.Indices) {
		return false
	}
	for i := range a.Indices {
		if !IsSimilarExpr(a.Indices[i], b.Indices[i]) {
			return false
		}
	}
	return IsSimilarExpr(a.X, b.X)
}

func IsSimilarInterfaceType(a, b *ast.InterfaceType) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Incomplete == b.Incomplete && IsSimilarFieldList(a.Methods, b.Methods)
}

func IsSimilarKeyValueExpr(a, b *ast.KeyValueExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.Key, b.Key) && IsSimilarExpr(a.Value, b.Value)
}

func IsSimilarMapType(a, b *ast.MapType) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.Key, b.Key) && IsSimilarExpr(a.Value, b.Value)
}

func IsSimilarParenExpr(a, b *ast.ParenExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.X, b.X)
}

func IsSimilarSelectorExpr(a, b *ast.SelectorExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.X, b.X) && a.Sel.Name == b.Sel.Name
}

func IsSimilarSliceExpr(a, b *ast.SliceExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.X, b.X) &&
		IsSimilarExpr(a.Low, b.Low) &&
		IsSimilarExpr(a.High, b.High) &&
		IsSimilarExpr(a.Max, b.Max) &&
		a.Slice3 == b.Slice3
}

func IsSimilarStarExpr(a, b *ast.StarExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.X, b.X)
}

func IsSimilarStructType(a, b *ast.StructType) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Incomplete == b.Incomplete && IsSimilarFieldList(a.Fields, b.Fields)
}

func IsSimilarTypeAssertExpr(a, b *ast.TypeAssertExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarExpr(a.X, b.X) && IsSimilarExpr(a.Type, b.Type)
}

func IsSimilarUnaryExpr(a, b *ast.UnaryExpr) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Op == b.Op && IsSimilarExpr(a.X, b.X)
}

func IsExpr(expr ast.Expr) bool {
	switch expr.(type) {
	case *ast.ArrayType, *ast.BadExpr, *ast.BasicLit, *ast.BinaryExpr, *ast.CallExpr,
		*ast.ChanType, *ast.CompositeLit, *ast.Ellipsis, *ast.FuncLit, *ast.FuncType,
		*ast.Ident, *ast.IndexExpr, *ast.IndexListExpr, *ast.InterfaceType, *ast.KeyValueExpr,
		*ast.MapType, *ast.ParenExpr, *ast.SelectorExpr, *ast.SliceExpr, *ast.StarExpr,
		*ast.StructType, *ast.TypeAssertExpr, *ast.UnaryExpr:
		return true
	default:
		return false
	}
}

func IsArrayType(expr ast.Expr) bool {
	_, ok := expr.(*ast.ArrayType)
	return ok
}

func IsBadExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.BadExpr)
	return ok
}

func IsBasicLit(expr ast.Expr) bool {
	_, ok := expr.(*ast.BasicLit)
	return ok
}

func IsCallExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.CallExpr)
	return ok
}

func IsChanType(expr ast.Expr) bool {
	_, ok := expr.(*ast.ChanType)
	return ok
}

func IsCompositeLit(expr ast.Expr) bool {
	_, ok := expr.(*ast.CompositeLit)
	return ok
}

func IsEllipsis(expr ast.Expr) bool {
	_, ok := expr.(*ast.Ellipsis)
	return ok
}

func IsFuncLit(expr ast.Expr) bool {
	_, ok := expr.(*ast.FuncLit)
	return ok
}

func IsFuncType(expr ast.Expr) bool {
	_, ok := expr.(*ast.FuncType)
	return ok
}

func IsIdent(expr ast.Expr) bool {
	_, ok := expr.(*ast.Ident)
	return ok
}

func IsInterfaceType(expr ast.Expr) bool {
	_, ok := expr.(*ast.InterfaceType)
	return ok
}

func IsIndexExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.IndexExpr)
	return ok
}

func IsIndexListExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.IndexListExpr)
	return ok
}

func IsKeyValueExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.KeyValueExpr)
	return ok
}

func IsMapType(expr ast.Expr) bool {
	_, ok := expr.(*ast.MapType)
	return ok
}

func IsParenExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.ParenExpr)
	return ok
}

func IsSelectorExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.SelectorExpr)
	return ok
}

func IsSliceExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.SliceExpr)
	return ok
}

func IsStarExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.StarExpr)
	return ok
}

func IsStructType(expr ast.Expr) bool {
	_, ok := expr.(*ast.StructType)
	return ok
}

func IsTypeAssertExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.TypeAssertExpr)
	return ok
}

func IsUnaryExpr(expr ast.Expr) bool {
	_, ok := expr.(*ast.UnaryExpr)
	return ok
}
