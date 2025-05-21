// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package xtr

import (
	"fmt"
	"testing"
)

func TestExtractFields(t *testing.T) {
	// Test Struct
	type MyStruct struct {
		Exported   string
		unexported int
	}

	obj := MyStruct{
		Exported:   "hello",
		unexported: 42,
	}

	fields := Fields(obj)
	for _, f := range fields {
		fmt.Printf("%s: %v\n", f.Name, f.Value)
	}
}
