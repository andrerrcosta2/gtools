// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"fmt"
	"testing"
)

func TestGen_NewGraph(t *testing.T) {
	g, err := New[int, int](Directed, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Graph: %v\n", g)
}
