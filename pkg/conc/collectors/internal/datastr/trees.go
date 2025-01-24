// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package datastr

import "github.com/andrerrcosta2/gtools/core/gtools/functions"

// TreeAsyncFlattenedData This structure represents the
// asynchronously flattened data from a tree as a slice
type TreeAsyncFlattenedData[K comparable, I, O any] struct {
	stages map[K]*AsyncValue[I, O]
	o      []O
}

func (t *TreeAsyncFlattenedData[K, I, O]) AddAsync(stage K, f functions.Function[I, O]) {

}
