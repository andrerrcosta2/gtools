// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package spec

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/ask"
)

func IsSimilar(a, b ast.Spec) bool {
	return ask.IsSimilarSpec(a, b)
}

func IsSimilarImportSpec(a, b *ast.ImportSpec) bool {
	return ask.IsSimilarImportSpec(a, b)
}

func IsSimilarTypeSpec(a, b *ast.TypeSpec) bool {
	return ask.IsSimilarTypeSpec(a, b)
}

func IsSimilarValueSpec(a, b *ast.ValueSpec) bool {
	return ask.IsSimilarValueSpec(a, b)
}

func IsSpec(s ast.Node) bool {
	return ask.IsSpec(s)
}

func IsImportSpec(s ast.Spec) bool {
	return ask.IsImportSpec(s)
}

func IsTypeSpec(s ast.Spec) bool {
	return ask.IsTypeSpec(s)
}

func IsValueSpec(s ast.Spec) bool {
	return ask.IsValueSpec(s)
}
