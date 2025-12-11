// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"fmt"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/ast4/internal/types"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func Expr(a, b ast.Expr) (equals bool, diff string) {
	return expr(std.Zero, a, b)
}

func expr(t indent.Indentor, a, b ast.Expr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Expr", "ast.Expr"); has {
		return eq, d
	}
	switch expr := a.(type) {
	case *ast.ArrayType:
		if other, ok := b.(*ast.ArrayType); ok {
			return arrayType(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.ArrayType", types.Expr(b))
	case *ast.BadExpr:
		if other, ok := b.(*ast.BadExpr); ok {
			return badExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.BadExpr", types.Expr(b))
	case *ast.BasicLit:
		if other, ok := b.(*ast.BasicLit); ok {
			return basicLit(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.BasicLit", types.Expr(b))
	case *ast.BinaryExpr:
		if other, ok := b.(*ast.BinaryExpr); ok {
			return binaryExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.BinaryExpr", types.Expr(b))
	case *ast.CallExpr:
		if other, ok := b.(*ast.CallExpr); ok {
			return callExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.CallExpr", types.Expr(b))
	case *ast.ChanType:
		if other, ok := b.(*ast.ChanType); ok {
			return chanType(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.ChanType", types.Expr(b))
	case *ast.CompositeLit:
		if other, ok := b.(*ast.CompositeLit); ok {
			return compositeLit(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.CompositeLit", types.Expr(b))
	case *ast.Ellipsis:
		if other, ok := b.(*ast.Ellipsis); ok {
			return ellipsis(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.Ellipsis", types.Expr(b))
	case *ast.FuncLit:
		if other, ok := b.(*ast.FuncLit); ok {
			return funcLit(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.FuncLit", types.Expr(b))
	case *ast.FuncType:
		if other, ok := b.(*ast.FuncType); ok {
			return funcType(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.FuncType", types.Expr(b))
	case *ast.Ident:
		if other, ok := b.(*ast.Ident); ok {
			return ident(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.Ident", types.Expr(b))
	case *ast.IndexExpr:
		if other, ok := b.(*ast.IndexExpr); ok {
			return indexExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.IndexExpr", types.Expr(b))
	case *ast.IndexListExpr:
		if other, ok := b.(*ast.IndexListExpr); ok {
			return indexListExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.IndexListExpr", types.Expr(b))
	case *ast.InterfaceType:
		if other, ok := b.(*ast.InterfaceType); ok {
			return interfaceType(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.InterfaceType", types.Expr(b))
	case *ast.KeyValueExpr:
		if other, ok := b.(*ast.KeyValueExpr); ok {
			return keyValueExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.KeyValueExpr", types.Expr(b))
	case *ast.MapType:
		if other, ok := b.(*ast.MapType); ok {
			return mapType(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.MapType", types.Expr(b))
	case *ast.ParenExpr:
		if other, ok := b.(*ast.ParenExpr); ok {
			return parenExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.ParenExpr", types.Expr(b))
	case *ast.SelectorExpr:
		if other, ok := b.(*ast.SelectorExpr); ok {
			return selectorExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.SelectorExpr", types.Expr(b))
	case *ast.SliceExpr:
		if other, ok := b.(*ast.SliceExpr); ok {
			return sliceExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.SliceExpr", types.Expr(b))
	case *ast.StarExpr:
		if other, ok := b.(*ast.StarExpr); ok {
			return starExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.StarExpr", types.Expr(b))
	case *ast.StructType:
		if other, ok := b.(*ast.StructType); ok {
			return structType(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.StructType", types.Expr(b))
	case *ast.TypeAssertExpr:
		if other, ok := b.(*ast.TypeAssertExpr); ok {
			return typeAssertExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.TypeAssertExpr", types.Expr(b))
	case *ast.UnaryExpr:
		if other, ok := b.(*ast.UnaryExpr); ok {
			return unaryExpr(t, expr, other)
		}
		return false, differs.TypesMismatch(t, "*ast.UnaryExpr", types.Expr(b))
	default:
		return false, std.ExprTypeDiffError(t, a, b)
	}
}

func ArrayType(expr, other *ast.ArrayType) (equals bool, diff string) {
	return arrayType(std.Zero, expr, other)
}

func arrayType(t indent.Indentor, a, b *ast.ArrayType) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "ArrayType", "*ast.ArrayType"); has {
		return eq, d
	}
	if eq, dif := expr(t.Inc(), a.Elt, b.Elt); !eq {
		return false, t.Sprintf("'Elt' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Len, b.Len); !eq {
		return false, t.Sprintf("'Len' values mismatch:\n%s", dif)
	}
	if a.Lbrack != b.Lbrack {
		return false, differs.ValuesOf(t, "Lbrack",
			sprints.Field(std.Zero, "Lbrack", sprints.Digit(a.Lbrack)),
			sprints.Field(std.Zero, "Lbrack", sprints.Digit(b.Lbrack)))
	}
	return true, diff
}

func BadExpr(expr, other *ast.BadExpr) (equals bool, diff string) {
	return badExpr(std.Zero, expr, other)
}

func badExpr(t indent.Indentor, a, b *ast.BadExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "BadExpr", "*ast.BadExpr"); has {
		return eq, d
	}
	if a.From != b.From {
		return false, differs.ValuesOf(t, "From",
			sprints.Field(std.Zero, "From", sprints.Digit(a.From)),
			sprints.Field(std.Zero, "From", sprints.Digit(b.From)))
	}
	if a.To != b.To {
		return false, differs.ValuesOf(t, "To",
			sprints.Field(std.Zero, "To", sprints.Digit(a.To)),
			sprints.Field(std.Zero, "To", sprints.Digit(b.To)))
	}
	return true, diff
}

func BasicLit(a, b *ast.BasicLit) (equals bool, diff string) {
	return basicLit(std.Zero, a, b)
}

func basicLit(t indent.Indentor, a, b *ast.BasicLit) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "BasicLit", "*ast.BasicLit"); has {
		return eq, d
	}
	if a.Kind != b.Kind {
		return false, differs.ValuesOf(t, "Kind",
			sprints.Field(std.Zero, "Kind", sprints.Digit(a.Kind)),
			sprints.Field(std.Zero, "Kind", sprints.Digit(b.Kind)))
	}
	if a.Value != b.Value {
		return false, differs.ValuesOf(t, "Value",
			sprints.Field(std.Zero, "Value", a.Value),
			sprints.Field(std.Zero, "Value", b.Value))
	}

	if a.ValuePos != b.ValuePos {
		return false, differs.ValuesOf(t, "ValuePos",
			sprints.Field(std.Zero, "ValuePos", sprints.Digit(a.ValuePos)),
			sprints.Field(std.Zero, "ValuePos", sprints.Digit(b.ValuePos)))
	}
	return true, diff
}

func BinaryExpr(a, b *ast.BinaryExpr) (equals bool, diff string) {
	return binaryExpr(std.Zero, a, b)
}

func binaryExpr(t indent.Indentor, a, b *ast.BinaryExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "BinaryExpr", "*ast.BinaryExpr"); has {
		return eq, d
	}
	if a.OpPos != b.OpPos {
		return false, differs.ValuesOf(t, "OpPos",
			sprints.Field(std.Zero, "OpPos", sprints.Digit(a.OpPos)),
			sprints.Field(std.Zero, "OpPos", sprints.Digit(b.OpPos)))
	}
	if a.Op != b.Op {
		return false, differs.ValuesOf(t, "Op",
			sprints.Field(std.Zero, "Op", sprints.Digit(a.Op)),
			sprints.Field(std.Zero, "Op", sprints.Digit(b.Op)))
	}

	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Y, b.Y); !eq {
		return false, t.Sprintf("'Y' values mismatch:\n%s", dif)
	}
	return true, diff
}

func CallExpr(a, b *ast.CallExpr) (equals bool, diff string) {
	return callExpr(std.Zero, a, b)
}

func callExpr(t indent.Indentor, a, b *ast.CallExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "CallExpr", "*ast.CallExpr"); has {
		return eq, d
	}
	if a.Ellipsis != b.Ellipsis {
		return false, differs.ValuesOf(t, "Ellipsis",
			sprints.Field(std.Zero, "Ellipsis", sprints.Digit(a.Ellipsis)),
			sprints.Field(std.Zero, "Ellipsis", sprints.Digit(b.Ellipsis)))
	}
	if a.Lparen != b.Lparen {
		return false, differs.ValuesOf(t, "Lparen",
			sprints.Field(std.Zero, "Lparen", sprints.Digit(a.Lparen)),
			sprints.Field(std.Zero, "Lparen", sprints.Digit(b.Lparen)))
	}
	if a.Rparen != b.Rparen {
		return false, differs.ValuesOf(t, "Rparen",
			sprints.Field(std.Zero, "Rparen", sprints.Digit(a.Rparen)),
			sprints.Field(std.Zero, "Rparen", sprints.Digit(b.Rparen)))
	}
	if eq, dif := expr(t.Inc(), a.Fun, b.Fun); !eq {
		return false, t.Sprintf("Fun values mismatch:\n%s", dif)
	}
	for i := range a.Args {
		if eq, dif := expr(t.Inc(), a.Args[i], b.Args[i]); !eq {
			return false, t.Sprintf("'Args[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff
}

func ChanType(a, b *ast.ChanType) (equals bool, diff string) {
	return chanType(std.Zero, a, b)
}

func chanType(t indent.Indentor, a, b *ast.ChanType) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "ChanType", "*ast.ChanType"); has {
		return eq, d
	}
	if a.Arrow != b.Arrow {
		return false, differs.ValuesOf(t, "Arrow",
			sprints.Field(std.Zero, "Arrow", sprints.Digit(a.Arrow)),
			sprints.Field(std.Zero, "Arrow", sprints.Digit(b.Arrow)))
	}
	if a.Begin != b.Begin {
		return false, differs.ValuesOf(t, "Begin",
			sprints.Field(std.Zero, "Begin", sprints.Digit(a.Begin)),
			sprints.Field(std.Zero, "Begin", sprints.Digit(b.Begin)))
	}
	if a.Dir != b.Dir {
		return false, differs.ValuesOf(t, "Dir",
			sprints.Field(std.Zero, "Dir", sprints.Digit(a.Dir)),
			sprints.Field(std.Zero, "Dir", sprints.Digit(b.Dir)))
	}
	if eq, dif := expr(t.Inc(), a.Value, b.Value); !eq {
		return false, t.Sprintf("Value values mismatch:\n%s", dif)
	}
	return true, diff
}

func CompositeLit(a, b *ast.CompositeLit) (equals bool, diff string) {
	return compositeLit(std.Zero, a, b)
}

func compositeLit(t indent.Indentor, a, b *ast.CompositeLit) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "CompositeLit", "*ast.CompositeLit"); has {
		return eq, d
	}
	if a.Lbrace != b.Lbrace {
		return false, differs.ValuesOf(t, "Lbrace",
			sprints.Field(std.Zero, "Lbrace", sprints.Digit(a.Lbrace)),
			sprints.Field(std.Zero, "Lbrace", sprints.Digit(b.Lbrace)))
	}
	if a.Rbrace != b.Rbrace {
		return false, differs.ValuesOf(t, "Rbrace",
			sprints.Field(std.Zero, "Rbrace", sprints.Digit(a.Rbrace)),
			sprints.Field(std.Zero, "Rbrace", sprints.Digit(b.Rbrace)))
	}
	if eq, dif := expr(t.Inc(), a.Type, b.Type); !eq {
		return false, fmt.Sprintf("'Type' values mismatch:\n%s", dif)
	}
	for i := range a.Elts {
		if eq, dif := Expr(a.Elts[i], b.Elts[i]); !eq {
			return false, fmt.Sprintf("'Elts[%d]' values mismatch:\n%s", i, dif)
		}
	}
	if a.Incomplete != b.Incomplete {
		return false, differs.ValuesOf(t, "Incomplete",
			sprints.Field(std.Zero, "Incomplete", sprints.Bool(std.Zero, a.Incomplete)),
			sprints.Field(std.Zero, "Incomplete", sprints.Bool(std.Zero, b.Incomplete)))
	}
	return true, diff
}

func Ellipsis(a, b *ast.Ellipsis) (equals bool, diff string) {
	return ellipsis(std.Zero, a, b)
}

func ellipsis(t indent.Indentor, a, b *ast.Ellipsis) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Ellipsis", "*ast.Ellipsis"); has {
		return eq, d
	}
	if a.Ellipsis != b.Ellipsis {
		return false, differs.ValuesOf(t, "Ellipsis",
			sprints.Field(std.Zero, "Ellipsis", sprints.Digit(a.Ellipsis)),
			sprints.Field(std.Zero, "Ellipsis", sprints.Digit(b.Ellipsis)))
	}
	if eq, dif := expr(t.Inc(), a.Elt, b.Elt); !eq {
		return false, t.Sprintf("'Elt' values mismatch:\n%s", dif)
	}
	return true, diff
}

func FuncLit(a, b *ast.FuncLit) (equals bool, diff string) {
	return funcLit(std.Zero, a, b)
}

func funcLit(t indent.Indentor, a, b *ast.FuncLit) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "FuncLit", "*ast.FuncLit"); has {
		return eq, d
	}
	if eq, dif := expr(t, a.Type, b.Type); !eq {
		return false, t.Sprintf("'Type' values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t, a.Body, b.Body); !eq {
		return false, t.Sprintf("'Body' values mismatch:\n%s", dif)
	}
	return true, diff
}

func FuncType(a, b *ast.FuncType) (equals bool, diff string) {
	return funcType(std.Zero, a, b)
}

func funcType(t indent.Indentor, a, b *ast.FuncType) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "FuncType", "*ast.FuncType"); has {
		return eq, d
	}
	if a.Func != b.Func {
		return false, differs.ValuesOf(t, "Func",
			sprints.Field(std.Zero, "Func", sprints.Digit(a.Func)),
			sprints.Field(std.Zero, "Func", sprints.Digit(b.Func)))
	}
	if eq, dif := fieldList(t.Inc(), a.TypeParams, b.TypeParams); !eq {
		return false, t.Sprintf("'TypeParams' values mismatch: \n%s", dif)
	}
	if eq, dif := fieldList(t.Inc(), a.Params, b.Params); !eq {
		return false, t.Sprintf("'Params' values mismatch:\n%s", dif)
	}
	if eq, dif := fieldList(t.Inc(), a.Results, b.Results); !eq {
		return false, t.Sprintf("'Results' values mismatch:\n%s", dif)
	}
	return true, diff
}

func Ident(a, b *ast.Ident) (equals bool, diff string) {
	return ident(std.Zero, a, b)
}

func ident(t indent.Indentor, a, b *ast.Ident) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Ident", "*ast.Ident"); has {
		return eq, d
	}
	if a.Name != b.Name {
		return false, differs.ValuesOf(t, "Name",
			sprints.Field(std.Zero, "Name", a.Name),
			sprints.Field(std.Zero, "Name", b.Name))
	}
	if a.NamePos != b.NamePos {
		return false, differs.ValuesOf(t, "NamePos",
			sprints.Field(std.Zero, "NamePos", sprints.Digit(a.NamePos)),
			sprints.Field(std.Zero, "NamePos", sprints.Digit(b.NamePos)))
	}
	return true, diff
}

func IndexExpr(a, b *ast.IndexExpr) (equals bool, diff string) {
	return indexExpr(std.Zero, a, b)
}

func indexExpr(t indent.Indentor, a, b *ast.IndexExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "IndexExpr", "*ast.IndexExpr"); has {
		return eq, d
	}
	if a.Lbrack != b.Lbrack {
		return false, differs.ValuesOf(t, "Lbrack",
			sprints.Field(std.Zero, "Lbrack", sprints.Digit(a.Lbrack)),
			sprints.Field(std.Zero, "Lbrack", sprints.Digit(b.Lbrack)))
	}
	if a.Rbrack != b.Rbrack {
		return false, differs.ValuesOf(t, "Rbrack",
			sprints.Field(std.Zero, "Rbrack", sprints.Digit(a.Rbrack)),
			sprints.Field(std.Zero, "Rbrack", sprints.Digit(b.Rbrack)),
		)
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch: \n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Index, b.Index); !eq {
		return false, t.Sprintf("'Index' values mismatch:\n%s", dif)
	}
	return true, diff
}

func IndexListExpr(a, b *ast.IndexListExpr) (equals bool, diff string) {
	return indexListExpr(std.Zero, a, b)
}

func indexListExpr(t indent.Indentor, a, b *ast.IndexListExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "IndexListExpr", "*ast.IndexListExpr"); has {
		return eq, d
	}
	if a.Lbrack != b.Lbrack {
		return false, differs.ValuesOf(t, "Lbrack",
			sprints.Field(std.Zero, "Lbrack", sprints.Digit(a.Lbrack)),
			sprints.Field(std.Zero, "Lbrack", sprints.Digit(b.Lbrack)))
	}
	if a.Rbrack != b.Rbrack {
		return false, differs.ValuesOf(t, "Rbrack",
			sprints.Field(std.Zero, "Rbrack", sprints.Digit(a.Rbrack)),
			sprints.Field(std.Zero, "Rbrack", sprints.Digit(b.Rbrack)))
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch:\n%s", dif)
	}
	for i := range a.Indices {
		if eq, dif := expr(t.Inc(), a.Indices[i], b.Indices[i]); !eq {
			return false, t.Sprintf("'Indices[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff
}

func InterfaceType(a, b *ast.InterfaceType) (equals bool, diff string) {
	return interfaceType(std.Zero, a, b)
}

func interfaceType(t indent.Indentor, a, b *ast.InterfaceType) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "InterfaceType", "*ast.InterfaceType"); has {
		return eq, d
	}
	if a.Interface != b.Interface {
		return false, differs.ValuesOf(t, "Interface",
			sprints.Field(std.Zero, "Interface", sprints.Digit(a.Interface)),
			sprints.Field(std.Zero, "Interface", sprints.Digit(b.Interface)))
	}
	if a.Incomplete != b.Incomplete {
		return false, differs.ValuesOf(t, "Incomplete",
			sprints.Field(std.Zero, "Incomplete", sprints.Bool(std.Zero, a.Incomplete)),
			sprints.Field(std.Zero, "Incomplete", sprints.Bool(std.Zero, b.Incomplete)))
	}
	if eq, dif := fieldList(t.Inc(), a.Methods, b.Methods); !eq {
		return false, t.Sprintf("'Methods' values mismatch:\n%s", dif)
	}
	return true, diff
}

func KeyValueExpr(a, b *ast.KeyValueExpr) (equals bool, diff string) {
	return keyValueExpr(std.Zero, a, b)
}

func keyValueExpr(t indent.Indentor, a, b *ast.KeyValueExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "KeyValueExpr", "*ast.KeyValueExpr"); has {
		return eq, d
	}
	if eq, dif := expr(t.Inc(), a.Key, b.Key); !eq {
		return false, t.Sprintf("'Key' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Value, b.Value); !eq {
		return false, t.Sprintf("'Value' values mismatch:\n%s", dif)
	}
	return true, diff
}

func MapType(a, b *ast.MapType) (equals bool, diff string) {
	return mapType(std.Zero, a, b)
}

func mapType(t indent.Indentor, a, b *ast.MapType) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "KeyValueExpr", "*ast.KeyValueExpr"); has {
		return eq, d
	}
	if eq, dif := expr(t.Inc(), a.Key, b.Key); !eq {
		return false, t.Sprintf("'Key' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Key, b.Key); !eq {
		return false, t.Sprintf("'Key' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Value, b.Value); !eq {
		return false, t.Sprintf("'Value' values mismatch:\n%s", dif)
	}
	return true, diff
}

func ParenExpr(a, b *ast.ParenExpr) (equals bool, diff string) {
	return parenExpr(std.Zero, a, b)
}

func parenExpr(t indent.Indentor, a, b *ast.ParenExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "ParenExpr", "*ast.ParenExpr"); has {
		return eq, d
	}
	if a.Lparen != b.Lparen {
		return false, differs.ValuesOf(t, "Lparen",
			sprints.Field(std.Zero, "Lparen", sprints.Digit(a.Lparen)),
			sprints.Field(std.Zero, "Lparen", sprints.Digit(b.Lparen)))
	}
	if a.Rparen != b.Rparen {
		return false, differs.ValuesOf(t, "Rparen",
			sprints.Field(std.Zero, "Rparen", sprints.Digit(a.Rparen)),
			sprints.Field(std.Zero, "Rparen", sprints.Digit(b.Rparen)),
		)
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch: \n%s", dif)
	}
	return true, diff
}

func SelectorExpr(a, b *ast.SelectorExpr) (equals bool, diff string) {
	return selectorExpr(std.Zero, a, b)
}

func selectorExpr(t indent.Indentor, a, b *ast.SelectorExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "SelectorExpr", "*ast.SelectorExpr"); has {
		return eq, d
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch: \n%s", dif)
	}
	if eq, dif := ident(t.Inc(), a.Sel, b.Sel); !eq {
		return false, t.Sprintf("'Sel' values mismatch: \n%s", dif)
	}
	return true, diff
}

func SliceExpr(a, b *ast.SliceExpr) (equals bool, diff string) {
	return sliceExpr(std.Zero, a, b)
}

func sliceExpr(t indent.Indentor, a, b *ast.SliceExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "SliceExpr", "*ast.SliceExpr"); has {
		return eq, d
	}
	if a.Lbrack != b.Lbrack {
		return false, differs.ValuesOf(t, "Lbrack",
			sprints.Field(std.Zero, "Lbrack", sprints.Digit(a.Lbrack)),
			sprints.Field(std.Zero, "Lbrack", sprints.Digit(b.Lbrack)),
		)
	}
	if a.Slice3 != b.Slice3 {
		return false, differs.ValuesOf(t, "Slice3",
			sprints.Field(std.Zero, "Slice3", sprints.Bool(std.Zero, a.Slice3)),
			sprints.Field(std.Zero, "Slice3", sprints.Bool(std.Zero, b.Slice3)),
		)
	}
	if a.Rbrack != b.Rbrack {
		return false, differs.ValuesOf(t, "Rbrack",
			sprints.Field(std.Zero, "Rbrack", sprints.Digit(a.Rbrack)),
			sprints.Field(std.Zero, "Rbrack", sprints.Digit(b.Rbrack)),
		)
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch: \n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Low, b.Low); !eq {
		return false, t.Sprintf("'Low' values mismatch: \n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.High, b.High); !eq {
		return false, t.Sprintf("'High' values mismatch: \n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Max, b.Max); !eq {
		return false, t.Sprintf("'Max' values mismatch: \n%s", dif)
	}
	return true, diff
}

func StarExpr(a, b *ast.StarExpr) (equals bool, diff string) {
	return starExpr(std.Zero, a, b)
}

func starExpr(t indent.Indentor, a, b *ast.StarExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "StarExpr", "*ast.StarExpr"); has {
		return eq, d
	}
	if a.Star != b.Star {
		return false, differs.ValuesOf(t, "Star",
			sprints.Field(std.Zero, "Star", sprints.Digit(a.Star)),
			sprints.Field(std.Zero, "Star", sprints.Digit(b.Star)),
		)
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch:\n%s", dif)
	}
	return true, diff
}

func StructType(a, b *ast.StructType) (equals bool, diff string) {
	return structType(std.Zero, a, b)
}

func structType(t indent.Indentor, a, b *ast.StructType) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "StructType", "*ast.StructType"); has {
		return eq, d
	}
	if a.Struct != b.Struct {
		return false, differs.ValuesOf(t, "Struct",
			sprints.Field(std.Zero, "Struct", sprints.Digit(a.Struct)),
			sprints.Field(std.Zero, "Struct", sprints.Digit(b.Struct)),
		)
	}
	if a.Incomplete != b.Incomplete {
		return false, differs.ValuesOf(t, "Incomplete",
			sprints.Field(std.Zero, "Incomplete", sprints.Bool(std.Zero, a.Incomplete)),
			sprints.Field(std.Zero, "Incomplete", sprints.Bool(std.Zero, b.Incomplete)),
		)
	}
	if eq, dif := fieldList(t.Inc(), a.Fields, b.Fields); !eq {
		return false, t.Sprintf("'Fields' values mismatch:\n%s", dif)
	}
	return true, diff
}

func TypeAssertExpr(a, b *ast.TypeAssertExpr) (equals bool, diff string) {
	return typeAssertExpr(std.Zero, a, b)
}

func typeAssertExpr(t indent.Indentor, a, b *ast.TypeAssertExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "TypeAssertExpr", "*ast.TypeAssertExpr"); has {
		return eq, d
	}
	if eq, dif := Expr(a.X, b.X); !eq {
		return false, fmt.Sprintf("X values mismatch: \n%s", dif)
	}
	if a.Lparen != b.Lparen {
		return false, differs.ValuesOf(t, "Lparen",
			sprints.Field(std.Zero, "Lparen", sprints.Digit(a.Lparen)),
			sprints.Field(std.Zero, "Lparen", sprints.Digit(b.Lparen)),
		)
	}
	if a.Rparen != b.Rparen {
		return false, differs.ValuesOf(t, "Rparen",
			sprints.Field(std.Zero, "Rparen", sprints.Digit(a.Rparen)),
			sprints.Field(std.Zero, "Rparen", sprints.Digit(b.Rparen)),
		)
	}
	if eq, dif := expr(t.Inc(), a.Type, b.Type); !eq {
		return false, t.Sprintf("'Type' values mismatch:\n%s", dif)
	}
	return true, diff
}

func UnaryExpr(a, b *ast.UnaryExpr) (equals bool, diff string) {
	return unaryExpr(std.Zero, a, b)
}

func unaryExpr(t indent.Indentor, a, b *ast.UnaryExpr) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "UnaryExpr", "*ast.UnaryExpr"); has {
		return eq, d
	}
	if a.OpPos != b.OpPos {
		return false, differs.ValuesOf(t, "OpPos",
			sprints.Field(std.Zero, "OpPos", sprints.Digit(a.OpPos)),
			sprints.Field(std.Zero, "OpPos", sprints.Digit(b.OpPos)),
		)
	}
	if a.Op != b.Op {
		return false, differs.ValuesOf(t, "Op",
			sprints.Field(std.Zero, "Op", sprints.Digit(a.Op)),
			sprints.Field(std.Zero, "Op", sprints.Digit(b.Op)),
		)
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, fmt.Sprintf("'X' values mismatch:\n%s", dif)
	}
	return true, diff
}
