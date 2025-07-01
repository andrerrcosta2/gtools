// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package random

import (
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/domain/validators"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/cat"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/lite"
	"github.com/andrerrcosta2/gtools/core/util/casters"
	"reflect"
)

// Validated TODO: maybe create a new randOf which uses validators
func Validated[T any](v validators.Typed[T], q int) *iterables.Slice[T] {
	if q <= 0 {
		return iterables.OfSlice[T]()
	}
	t := reflect.TypeOf((*T)(nil)).Elem()
	result := make(iterables.Slice[T], q)
	switch cat.CastMethod(t) {
	case cat.Any:
		for i := 0; i < q; i++ {
			rdn := lite.RandAny()
			result[i] = rdn.(T)
		}
		break
	case cat.Injectable:
		for i := 0; i < q; i++ {
			var zero T
			result[i] = zero
		}
		break
	case cat.Reference:
		for i := 0; i < q; i++ {
			rdn := lite.RandOf(t)
			result[i] = casters.UnsafeReferenceOf[T](rdn)
		}
		break
	default:
		for i := 0; i < q; i++ {
			rdn := lite.RandOf(t)
			result[i] = casters.UnsafeValueOf[T](rdn)
		}
		break
	}
	return &result
}
