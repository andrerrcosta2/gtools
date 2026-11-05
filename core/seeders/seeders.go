// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package seeders

import (
	"github.com/andrerrcosta2/gtools/core/domain/contracts/validator"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
)

func Random[T any](validator validator.Typed[T]) Slice[T] {
	return &rand[T]{
		validator: validator,
	}
}

type rand[T any] struct {
	validator validator.Typed[T]
}

func (r *rand[T]) Seed(n int) []T {
	return random.Validated[T](r.validator, n).Values()
}
