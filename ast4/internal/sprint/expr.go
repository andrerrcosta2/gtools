// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func Expr(tab indent.Indentor, expr ast.Expr) string {
	if expr == nil {
		return "ast.Expr<nil>"
	}
	switch e := expr.(type) {
	case *ast.ArrayType:
		return ArrayType(tab, e)
	case *ast.BadExpr:
		return BadExpr(tab, e)
	case *ast.BasicLit:
		return BasicLit(tab, e)
	case *ast.BinaryExpr:
		return BinaryExpr(tab, e)
	case *ast.CallExpr:
		return CallExpr(tab, e)
	case *ast.ChanType:
		return ChanType(tab, e)
	case *ast.CompositeLit:
		return CompositeLit(tab, e)
	case *ast.Ellipsis:
		return Ellipsis(tab, e)
	case *ast.FuncLit:
		return FuncLit(tab, e)
	case *ast.FuncType:
		return FuncType(tab, e)
	case *ast.Ident:
		return Ident(tab, e)
	case *ast.IndexExpr:
		return IndexExpr(tab, e)
	case *ast.IndexListExpr:
		return IndexListExpr(tab, e)
	case *ast.InterfaceType:
		return InterfaceType(tab, e)
	case *ast.KeyValueExpr:
		return KeyValueExpr(tab, e)
	case *ast.MapType:
		return MapType(tab, e)
	case *ast.ParenExpr:
		return ParenExpr(tab, e)
	case *ast.SelectorExpr:
		return SelectorExpr(tab, e)
	case *ast.SliceExpr:
		return SliceExpr(tab, e)
	case *ast.StarExpr:
		return StarExpr(tab, e)
	case *ast.StructType:
		return StructType(tab, e)
	case *ast.TypeAssertExpr:
		return TypeAssertExpr(tab, e)
	case *ast.UnaryExpr:
		return UnaryExpr(tab, e)
	default:
		return std.ExprTypeError(tab, e)
	}
}

func ArrayType(tab indent.Indentor, expr *ast.ArrayType) string {
	if expr == nil {
		return "ArrayType<nil>"
	}
	return sprints.Closedobj(tab, "ArrayType",
		sprints.Ltfield(tab, "Lbrack", sprints.Value(std.Zero, expr.Lbrack)),
		sprints.Ltfield(tab, "Elt", Expr(tab, expr.Elt)),
		sprints.Ltfield(tab, "Len", Expr(tab, expr.Len)))
}

func BadExpr(tab indent.Indentor, expr *ast.BadExpr) string {
	if expr == nil {
		return fmx.SRedf("BadExpr<nil>")
	}
	return sprints.Error(std.Zero, sprints.Closedobj(tab, "BadExpr",
		sprints.Ltfield(tab, "From", sprints.Value(std.Zero, expr.From)),
		sprints.Ltfield(tab, "To", sprints.Value(std.Zero, expr.To))))
}

func BasicLit(tab indent.Indentor, expr *ast.BasicLit) string {
	if expr == nil {
		return "BasicLit<nil>"
	}
	return sprints.Closedobj(tab, "BasicLit",
		sprints.Ltfield(tab, "Kind", sprints.Value(std.Zero, expr.Kind)),
		sprints.Ltfield(tab, "ValuePos", sprints.Value(std.Zero, expr.ValuePos)),
		sprints.Ltfield(tab, "Value", expr.Value))
}

func BinaryExpr(tab indent.Indentor, expr *ast.BinaryExpr) string {
	if expr == nil {
		return "BinaryExpr<nil>"
	}
	return sprints.Closedobj(tab, "BinaryExpr",
		sprints.Ltfield(tab, "OpPos", sprints.Value(std.Zero, expr.OpPos)),
		sprints.Ltfield(tab, "Op", sprints.Value(std.Zero, expr.Op)),
		sprints.Ltfield(tab, "X", Expr(tab.Inc(), expr.X)),
		sprints.Ltfield(tab, "Y", Expr(tab.Inc(), expr.Y)))
}

func CallExpr(tab indent.Indentor, expr *ast.CallExpr) string {
	if expr == nil {
		return "CallExpr<nil>"
	}
	return sprints.Closedobj(tab, "CallExpr",
		sprints.Ltfield(tab, "Fun", Expr(tab.Inc(), expr.Fun)),
		sprints.Ltfield(tab, "Lparen", sprints.Value(std.Zero, expr.Lparen)),
		sprints.Ltfield(tab, "Args", sprints.Slicef(tab.Inc(), expr.Args, Expr)),
		sprints.Ltfield(tab, "Ellipsis", sprints.Value(std.Zero, expr.Ellipsis)),
		sprints.Ltfield(tab, "Rparen", sprints.Value(std.Zero, expr.Rparen)))
}

func ChanType(tab indent.Indentor, expr *ast.ChanType) string {
	if expr == nil {
		return "ChanType<nil>"
	}
	return sprints.Closedobj(tab, "ChanType",
		sprints.Ltfield(tab, "Begin", sprints.Value(std.Zero, expr.Begin)),
		sprints.Ltfield(tab, "Arrow", sprints.Value(std.Zero, expr.Arrow)),
		sprints.Ltfield(tab, "Dir", sprints.Value(std.Zero, expr.Dir)),
		sprints.Ltfield(tab, "Value", Expr(tab.Inc(), expr.Value)))
}

func CompositeLit(tab indent.Indentor, expr *ast.CompositeLit) string {
	if expr == nil {
		return "CompositeLit<nil>"
	}
	return sprints.Closedobj(tab, "CompositeLit",
		sprints.Ltfield(tab, "Type", Expr(tab.Inc(), expr.Type)),
		sprints.Ltfield(tab, "Lbrace", sprints.Value(std.Zero, expr.Lbrace)),
		sprints.Ltfield(tab, "Elts", sprints.Slicef(tab.Inc(), expr.Elts, Expr)),
		sprints.Ltfield(tab, "Rbrace", sprints.Value(std.Zero, expr.Rbrace)))
}

func Ellipsis(tab indent.Indentor, expr *ast.Ellipsis) string {
	if expr == nil {
		return "Ellipsis<nil>"
	}
	return sprints.Closedobj(tab, "Ellipsis",
		sprints.Ltfield(tab, "Ellipsis", sprints.Value(std.Zero, expr.Ellipsis)),
		sprints.Ltfield(tab, "Elt", Expr(tab.Inc(), expr.Elt)))
}

func FuncLit(tab indent.Indentor, expr *ast.FuncLit) string {
	if expr == nil {
		return "FuncLit<nil>"
	}
	return sprints.Closedobj(tab, "FuncLit",
		sprints.Ltfield(tab, "Type", Expr(tab.Inc(), expr.Type)),
		sprints.Ltfield(tab, "Body", Stmt(tab.Inc(), expr.Body)))
}

func FuncType(tab indent.Indentor, e *ast.FuncType) string {
	if e == nil {
		return "FuncType<nil>"
	}
	return sprints.Closedobj(tab, "FuncType",
		sprints.Ltfield(tab, "Func", sprints.Value(std.Zero, e.Func)),
		sprints.Ltfield(tab, "TypeParams", FieldList(tab.Inc(), e.TypeParams)),
		sprints.Ltfield(tab, "Params", FieldList(tab.Inc(), e.Params)),
		sprints.Ltfield(tab, "Results", FieldList(tab.Inc(), e.Results)))
}

func Ident(tab indent.Indentor, expr *ast.Ident) string {
	if expr == nil {
		return "Ident<nil>"
	}
	return sprints.Closedobj(tab, "Ident",
		sprints.Ltfield(tab, "NamePos", sprints.Value(std.Zero, expr.NamePos)),
		sprints.Ltfield(tab, "Name", expr.Name),
		sprints.Ltfield(tab, sprints.Deprecated(std.Zero, "Obj"), sprints.Deprecated(std.Zero, "expr.Obj")))
}

func IndexExpr(tab indent.Indentor, expr *ast.IndexExpr) string {
	if expr == nil {
		return "IndexExpr<nil>"
	}
	return sprints.Closedobj(tab, "IndexExpr",
		sprints.Ltfield(tab, "Lbrack", sprints.Value(std.Zero, expr.Lbrack)),
		sprints.Ltfield(tab, "X", Expr(tab.Inc(), expr.X)),
		sprints.Ltfield(tab, "Index", Expr(tab.Inc(), expr.Index)),
		sprints.Ltfield(tab, "Rbrack", sprints.Value(std.Zero, expr.Rbrack)))
}

func IndexListExpr(tab indent.Indentor, expr *ast.IndexListExpr) string {
	if expr == nil {
		return "IndexListExpr<nil>"
	}
	return sprints.Closedobj(tab, "IndexListExpr",
		sprints.Ltfield(tab, "Lbrack", sprints.Value(std.Zero, expr.Lbrack)),
		sprints.Ltfield(tab, "X", Expr(tab, expr.X)),
		sprints.Ltfield(tab, "Indices", sprints.Slicef(tab.Inc(), expr.Indices, Expr)),
		sprints.Ltfield(tab, "Rbrack", sprints.Value(std.Zero, expr.Rbrack)))
}

func InterfaceType(tab indent.Indentor, expr *ast.InterfaceType) string {
	if expr == nil {
		return "InterfaceType<nil>"
	}
	return sprints.Closedobj(tab, "InterfaceType",
		sprints.Ltfield(tab, "Interface", sprints.Value(std.Zero, expr.Interface)),
		sprints.Ltfield(tab, "Methods", FieldList(tab, expr.Methods)),
		sprints.Ltfield(tab, "Incomplete", sprints.Value(std.Zero, expr.Incomplete)))
}

func KeyValueExpr(tab indent.Indentor, expr *ast.KeyValueExpr) string {
	if expr == nil {
		return "KeyValueExpr<nil>"
	}
	return sprints.Closedobj(tab, "KeyValueExpr",
		sprints.Ltfield(tab, "Key", Expr(tab, expr.Key)),
		sprints.Ltfield(tab, "Colon", sprints.Value(tab, expr.Colon)),
		sprints.Ltfield(tab, "Value", Expr(tab, expr.Value)))
}

func MapType(tab indent.Indentor, expr *ast.MapType) string {
	if expr == nil {
		return "MapType<nil>"
	}
	return sprints.Closedobj(tab, "MapType",
		sprints.Ltfield(tab, "Map", sprints.Value(tab, expr.Map)),
		sprints.Ltfield(tab, "Key", Expr(tab, expr.Key)),
		sprints.Ltfield(tab, "Value", Expr(tab, expr.Value)))
}

func ParenExpr(tab indent.Indentor, expr *ast.ParenExpr) string {
	if expr == nil {
		return "ParenExpr<nil>"
	}
	return sprints.Closedobj(tab, "ParenExpr",
		sprints.Ltfield(tab, "Lparen", sprints.Value(tab, expr.Lparen)),
		sprints.Ltfield(tab, "X", Expr(tab, expr.X)),
		sprints.Ltfield(tab, "Rparen", sprints.Value(tab, expr.Rparen)))
}

func SelectorExpr(tab indent.Indentor, expr *ast.SelectorExpr) string {
	if expr == nil {
		return "SelectorExpr<nil>"
	}
	return sprints.Closedobj(tab, "SelectorExpr",
		sprints.Ltfield(tab, "X", Expr(tab, expr.X)),
		sprints.Ltfield(tab, "Sel", Ident(tab, expr.Sel)))
}

func SliceExpr(tab indent.Indentor, expr *ast.SliceExpr) string {
	if expr == nil {
		return "SliceExpr<nil>"
	}
	return sprints.Closedobj(tab, "SliceExpr",
		sprints.Ltfield(tab, "X", Expr(tab, expr.X)),
		sprints.Ltfield(tab, "Lbrack", sprints.Value(tab, expr.Lbrack)),
		sprints.Ltfield(tab, "Low", Expr(tab, expr.Low)),
		sprints.Ltfield(tab, "High", Expr(tab, expr.High)),
		sprints.Ltfield(tab, "Max", Expr(tab, expr.Max)),
		sprints.Ltfield(tab, "Slice3", sprints.Value(tab, expr.Slice3)),
		sprints.Ltfield(tab, "Rbrack", sprints.Value(tab, expr.Rbrack)))
}

func StarExpr(tab indent.Indentor, expr *ast.StarExpr) string {
	if expr == nil {
		return "StarExpr<nil>"
	}
	return sprints.Closedobj(tab, "StarExpr",
		sprints.Ltfield(tab, "StarPos", sprints.Value(tab, expr.Star)),
		sprints.Ltfield(tab, "Star", Expr(tab, expr.X)))
}

func StructType(tab indent.Indentor, expr *ast.StructType) string {
	if expr == nil {
		return "StructType<nil>"
	}
	return sprints.Closedobj(tab, "StructType",
		sprints.Ltfield(tab, "Struct", sprints.Value(tab, expr.Struct)),
		sprints.Ltfield(tab, "Fields", FieldList(tab, expr.Fields)))
}

func TypeAssertExpr(tab indent.Indentor, expr *ast.TypeAssertExpr) string {
	if expr == nil {
		return "TypeAssertExpr<nil>"
	}
	return sprints.Closedobj(tab, "TypeAssertExpr",
		sprints.Ltfield(tab, "Lparen", sprints.Value(tab, expr.Lparen)),
		sprints.Ltfield(tab, "X", Expr(tab, expr.X)),
		sprints.Ltfield(tab, "Type", Expr(tab, expr.Type)),
		sprints.Ltfield(tab, "Rparen", sprints.Value(tab, expr.Rparen)))
}

func UnaryExpr(tab indent.Indentor, expr *ast.UnaryExpr) string {
	if expr == nil {
		return "UnaryExpr<nil>"
	}
	return sprints.Closedobj(tab, "UnaryExpr",
		sprints.Ltfield(tab, "OpPos", sprints.Value(tab, expr.OpPos)),
		sprints.Ltfield(tab, "Op", sprints.Value(tab, expr.Op)),
		sprints.Ltfield(tab, "X", Expr(tab, expr.X)))
}
