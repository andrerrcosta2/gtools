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

// XrtSymbol extracts the symbol and variable from a type string.
// The type string can contain '*' or '&' symbols, which indicate a pointer or reference.
// The function returns the symbol and variable as separate strings.
// If there are no symbols, it returns an empty symbol and the full type as the variable.
func XrtSymbol(typ string) (symbol string, variable string) {
	// Iterate over the type string
	for i, r := range typ {
		// If the character is not '*' or '&', we've found the end of the symbol
		if r != '*' && r != '&' {
			// Return the symbol and the rest of the string as the variable
			return typ[:i], typ[i:]
		}
	}
	// If there are no symbols, return an empty symbol and the full type as the variable
	return "", typ
}
