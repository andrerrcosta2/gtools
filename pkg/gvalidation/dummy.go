// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gvalidation

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/generics"
)

type DummyData[T any] struct {
	data T
}

func UseCore() {
	var x generics.Typed[int] = &DummyData[int]{
		data: 1,
	}
	fmt.Printf("x: %v\n", x)
}
