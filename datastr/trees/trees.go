// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package trees

import (
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
)

type abstract[N nodes.Typed[V], V any] struct {
	_root N
	_size int
}

func (t *abstract[N, V]) Root() N {
	return t._root
}

func (t *abstract[N, V]) SetRoot(root N) {
	t._root = root
}

func (t *abstract[N, V]) Size() int {
	return t._size
}

func (t *abstract[N, V]) SetSize(size int) {
	t._size = size
}
