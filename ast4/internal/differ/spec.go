// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/ast4/internal/types"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func Spec(a, b ast.Spec) (equals bool, diff string) {
	return spec(std.Zero, a, b)
}

func spec(t indent.Indentor, a, b ast.Spec) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Spec", "ast.Spec"); has {
		return eq, d
	}
	switch s := a.(type) {
	case *ast.ImportSpec:
		if other, ok := b.(*ast.ImportSpec); ok {
			return importSpec(t, s, other)
		}
		return false, differs.TypesMismatch(t, "*ast.ImportSpec", types.Spec(b))
	case *ast.TypeSpec:
		if other, ok := b.(*ast.TypeSpec); ok {
			return typeSpec(t, s, other)
		}
		return false, differs.TypesMismatch(t, "*ast.TypeSpec", types.Spec(b))
	case *ast.ValueSpec:
		if other, ok := b.(*ast.ValueSpec); ok {
			return valueSpec(t, s, other)
		}
		return false, differs.TypesMismatch(t, "*ast.ValueSpec", types.Spec(b))
	default:
		return false, std.SpecTypeDiffError(t, a, b)
	}
}

func ImportSpec(a, b *ast.ImportSpec) (equals bool, diff string) {
	return importSpec(std.Zero, a, b)
}

func importSpec(t indent.Indentor, a, b *ast.ImportSpec) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "ImportSpec", "*ast.ImportSpec"); has {
		return eq, d
	}
	if a.EndPos != b.EndPos {
		return false, differs.ValuesOf(t, "EndPos",
			sprints.Field(std.Zero, "EndPos", sprints.Digit(a.EndPos)),
			sprints.Field(std.Zero, "EndPos", sprints.Digit(b.EndPos)),
		)
	}
	if eq, dif := commentGroup(t.Inc(), a.Doc, b.Doc); !eq {
		return false, t.Sprintf("'Doc' values mismatch:\n%s", dif)
	}
	if eq, dif := ident(t.Inc(), a.Name, b.Name); !eq {
		return false, t.Sprintf("'Name' values mismatch:\n%s", dif)
	}
	if eq, dif := basicLit(t.Inc(), a.Path, b.Path); !eq {
		return false, t.Sprintf("'Path' values mismatch:\n%s", dif)
	}
	if eq, dif := commentGroup(t.Inc(), a.Comment, b.Comment); !eq {
		return false, t.Sprintf("'Comment' values mismatch:\n%s", dif)
	}
	return true, diff
}

func TypeSpec(a, b *ast.TypeSpec) (equals bool, diff string) {
	return typeSpec(std.Zero, a, b)
}

func typeSpec(t indent.Indentor, a, b *ast.TypeSpec) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "TypeSpec", "*ast.TypeSpec"); has {
		return eq, d
	}
	if a.Assign != b.Assign {
		return false, differs.ValuesOf(t, "Assign",
			sprints.Field(std.Zero, "Assign", sprints.Digit(a.Assign)),
			sprints.Field(std.Zero, "Assign", sprints.Digit(b.Assign)),
		)
	}
	if eq, dif := commentGroup(t.Inc(), a.Doc, b.Doc); !eq {
		return false, t.Sprintf("'Doc' values mismatch:\n%s", dif)
	}
	if eq, dif := ident(t.Inc(), a.Name, b.Name); !eq {
		return false, t.Sprintf("'Name' values mismatch:\n%s", dif)
	}
	if eq, dif := fieldList(t.Inc(), a.TypeParams, b.TypeParams); !eq {
		return false, t.Sprintf("'TypeParams' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Type, b.Type); !eq {
		return false, t.Sprintf("'Type' values mismatch:\n%s", dif)
	}
	if eq, dif := commentGroup(t.Inc(), a.Comment, b.Comment); !eq {
		return false, t.Sprintf("'Comment' values mismatch:\n%s", dif)
	}
	return true, diff
}

func ValueSpec(a, b *ast.ValueSpec) (equals bool, diff string) {
	return valueSpec(std.Zero, a, b)
}

func valueSpec(t indent.Indentor, a, b *ast.ValueSpec) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "ValueSpec", "*ast.ValueSpec"); has {
		return eq, d
	}
	if eq, dif := commentGroup(t.Inc(), a.Doc, b.Doc); !eq {
		return false, t.Sprintf("'Doc' values mismatch:\n%s", dif)
	}
	if len(a.Names) != len(b.Names) {
		return false, differs.SizesOf(t, "Names", sprints.Digit(len(a.Names)), sprints.Digit(len(b.Names)))
	}
	for i := 0; i < len(a.Names); i++ {
		if eq, dif := ident(t.Inc(), a.Names[i], b.Names[i]); !eq {
			return false, t.Sprintf("'Names[%d]' values mismatch:\n%s", i, dif)
		}
	}
	if eq, dif := expr(t.Inc(), a.Type, b.Type); !eq {
		return false, t.Sprintf("'Type' values mismatch:\n%s", dif)
	}
	for i := 0; i < len(a.Values); i++ {
		if eq, dif := expr(t.Inc(), a.Values[i], b.Values[i]); !eq {
			return false, t.Sprintf("Values[%d] values mismatch:\n%s", i, dif)
		}
	}
	if eq, dif := commentGroup(t.Inc(), a.Comment, b.Comment); !eq {
		return false, t.Sprintf("Comment values mismatch:\n%s", dif)
	}
	return true, diff
}
