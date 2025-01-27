// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type Closeable interface {
	IsClosed() bool
	Close() error
}

type Taggable[T prim.Hashable] interface {
	// Tag adds the given tags to the Taggable and returns the tag string.
	// It returns a string in the format "tag1,tag2,...,tagN".
	Tag(tags ...T)
	// Tags returns the current tags associated with the Taggable.
	Tags() []T
}
