// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/ast4/internal/types"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
)

func Decl(a, b ast.Decl) (equals bool, diff string) {
	return decl(std.Zero, a, b)
}

func decl(t indent.Indentor, a, b ast.Decl) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Decl", "ast.Decl"); has {
		return eq, d
	}
	switch decl := a.(type) {
	case *ast.GenDecl:
		if other, ok := b.(*ast.GenDecl); ok {
			return genDecl(t, decl, other)
		}
		return false, differs.TypesMismatch(t, "*ast.GenDecl", types.Decl(b))
	case *ast.FuncDecl:
		if other, ok := b.(*ast.FuncDecl); ok {
			return funcDecl(t, decl, other)
		}
		return false, differs.TypesMismatch(t, "*ast.FuncDecl", types.Decl(b))
	default:
		return false, std.DeclTypeDiffError(t, a, b)
	}
}

func GenDecl(a, b *ast.GenDecl) (equals bool, diff string) {
	return genDecl(std.Zero, a, b)
}

func genDecl(t indent.Indentor, a, b *ast.GenDecl) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Decl", "ast.Decl"); has {
		return eq, d
	}
	if a.Lparen != b.Lparen {
		return false, differs.ValuesOf(t, "Lparen",
			sprints.Field(t, "Lparen", sprints.Digit(a.Lparen)),
			sprints.Field(t, "Lparen", sprints.Digit(b.Lparen)))
	}
	if a.Rparen != b.Rparen {
		return false, differs.ValuesOf(t, "Rparen",
			sprints.Field(t, "Rparen", sprints.Digit(a.Rparen)),
			sprints.Field(t, "Rparen", sprints.Digit(b.Rparen)))
	}
	if a.TokPos != b.TokPos {
		return false, differs.ValuesOf(t, "TokPos",
			sprints.Field(t, "TokPos", sprints.Digit(a.TokPos)),
			sprints.Field(t, "TokPos", sprints.Digit(b.TokPos)))
	}
	if a.Tok != b.Tok {
		return false, differs.ValuesOf(t, "Tok",
			sprints.Field(t, "Tok", sprints.Digit(a.Tok)),
			sprints.Field(t, "Tok", sprints.Digit(b.Tok)))
	}
	if eq, dif := commentGroup(t.Inc(), a.Doc, b.Doc); !eq {
		return false, t.Sprintf("Doc values mismatch: %s", dif)
	}
	for i := 0; i < len(a.Specs); i++ {
		if eq, dif := spec(t.Inc(), a.Specs[i], b.Specs[i]); !eq {
			return false, t.Sprintf("Spec[%d] values mismatch: %s", i, dif)
		}
	}
	return true, ""
}

func FuncDecl(decl, other *ast.FuncDecl) (equals bool, diff string) {
	return funcDecl(std.Zero, decl, other)
}

func funcDecl(t indent.Indentor, a, b *ast.FuncDecl) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "FuncDecl", "ast.FuncDecl"); has {
		return eq, d
	}
	if eq, dif := commentGroup(t.Inc(), a.Doc, b.Doc); !eq {
		return false, t.Sprintf("Doc values mismatch: %s", dif)
	}
	if eq, dif := ident(t.Inc(), a.Name, b.Name); !eq {
		return false, t.Sprintf("Name values mismatch: %s", dif)
	}
	if eq, dif := fieldList(t.Inc(), a.Recv, b.Recv); !eq {
		return false, t.Sprintf("Recv values mismatch: %s", dif)
	}
	if eq, dif := funcType(t.Inc(), a.Type, b.Type); !eq {
		return false, t.Sprintf("Type values mismatch: %s", dif)
	}
	if eq, dif := blockStmt(t.Inc(), a.Body, b.Body); !eq {
		return false, t.Sprintf("Body values mismatch: %s", dif)
	}
	return true, ""
}
