// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package seeders

import (
	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
)

type Slice[T any] interface {
	Seed(int) []T
}

type Stream[T any] interface {
	Seed(int) stream.Readable[T]
}
