// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testcomparables

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
)

func Struct(intValue int, stringValue string) *ComparableStruct {
	return &ComparableStruct{Int: intValue, String: stringValue}
}

type ComparableStruct struct {
	Int    int
	String string
}

func RandomStructs(q int) *iterables.Slice[ComparableStruct] {
	out := make(iterables.Slice[ComparableStruct], q)

	for i := 0; i < q; i++ {
		out[i] = ComparableStruct{i, fmt.Sprintf("string_%d", i)}
	}
	return &out
}

func RandomReferenceStructs(q int) iterables.Slice[*ComparableStruct] {
	out := make(iterables.Slice[*ComparableStruct], q)

	for i := 0; i < q; i++ {
		out[i] = Struct(i, fmt.Sprintf("string_%d", i))
	}
	return out
}
