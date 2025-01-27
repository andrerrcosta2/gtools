// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package random

import (
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/domain/validators"
	"reflect"
)

// Validated TODO: maybe create a new randOf which uses validators
func Validated[T any](v validators.Typed[T], q int) *iterables.Slice[T] {
	if q <= 0 {
		return iterables.OfSlice[T]()
	}
	var zero T
	typx := reflect.TypeOf(zero)
	result := make(iterables.Slice[T], q)
	for i := 0; i < q; i++ {
		// That seems a little bit overcautious from golang compiler to me.
		// Fortunately the optimizations of golang compiler reduces the overhead
		// here close to zero.
		if rnd, ok := randOf(typx).(T); ok {
			result[i] = rnd
		}
	}
	return &result
}
