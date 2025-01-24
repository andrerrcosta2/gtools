// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package seeders

import "github.com/andrerrcosta2/gtools/core/gtools"

type Slice[T any] interface {
	Seed(int) []T
}

type Stream[T any] interface {
	Seed(int) gtools.Stream[T]
}
