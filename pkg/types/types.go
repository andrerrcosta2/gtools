// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import (
	"github.com/andrerrcosta2/gtools/pkg/structs/sets"
	"strings"
)

var Builtins = sets.NewString(
	"string", "bool", "int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	"byte", "rune", "float32", "float64", "complex64", "complex128",
	"error", "interface{}", "struct{}", "map", "chan", "func", "any",

	// Pointer variations
	"*string", "*bool", "*int", "*int8", "*int16", "*int32", "*int64",
	"*uint", "*uint8", "*uint16", "*uint32", "*uint64", "*uintptr",
	"*byte", "*rune", "*float32", "*float64", "*complex64", "*complex128",
	"*error", "*interface{}", "*struct{}", "*map", "*chan", "*func", "*any",

	// Reference variations
	"&string", "&bool", "&int", "&int8", "&int16", "&int32", "&int64",
	"&uint", "&uint8", "&uint16", "&uint32", "&uint64", "&uintptr",
	"&byte", "&rune", "&float32", "&float64", "&complex64", "&complex128",
	"&error", "&interface{}", "&struct{}", "&map", "&chan", "&func", "&any",
)

func IsBuiltinType(typ string) bool {
	// Remove leading pointer/reference symbols
	trimmedType := strings.TrimPrefix(strings.TrimPrefix(typ, "*"), "&")
	return Builtins.Has(typ) || Builtins.Has(trimmedType)
}
