// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/datastr/internal/hash"
)

type hashable[H comparable, T any] interface {
	hash(t T) H
	table() hash.Store[H, T]
}

type insertable[T any, H comparable] interface {
	hashable[H, T]
	contains(t T) bool
	size() int
	setHash(h H, idx int)
	append(t T)
}

type mapProvider[K comparable, V any] interface {
	backingMap() map[K]V
}

type snapshotter[T any] interface {
	snap() []T
}
