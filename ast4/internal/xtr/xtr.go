// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package xtr

import (
	"go/ast"
)

// Structs extracts all structs from an *ast.File.
func Structs(file *ast.File) []*ast.TypeSpec {
	var structs []*ast.TypeSpec

	ast.Inspect(file, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if _, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
			structs = append(structs, typeSpec)
		}

		return true
	})

	return structs
}

// Fields extracts all fields from a *ast.StructType.
func Fields(structNode *ast.StructType) []*ast.Field {
	if structNode == nil || structNode.Fields == nil {
		return nil
	}

	// structNode.Fields.List já é um []*ast.Field com todos os campos da struct
	return structNode.Fields.List
}

// Methods extracts all methods from an *ast.File.
func Methods(file *ast.File) []*ast.FuncDecl {
	var methods []*ast.FuncDecl

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if fn.Recv != nil && len(fn.Recv.List) > 0 {
				methods = append(methods, fn)
			}
		}
	}

	return methods
}
