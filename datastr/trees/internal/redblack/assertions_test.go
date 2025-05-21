// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package redblack

import (
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
	"testing"
)

func assertValidRoot[V any](t *testing.T, tag string, a, b *Node[V]) {
	t.Helper()
	if a != b {
		t.Errorf("Expected root to be '%s(%v)', but got\n%v", tag, a.value, b)
	}
	if a.parent != nil {
		t.Errorf("Expected root parent to be 'nil', but got\n%v", a.parent)
	}
	if a.color != nodes.Black {
		t.Errorf("Expected root color to be 'black', but got '%v'", a.color)
	}
}
