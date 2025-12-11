// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package loop

import (
	"context"
	"testing"

	"github.com/andrerrcosta2/gtools/core/format/fmx"
)

// TestBF order matters, repetition is allowed
func TestBF(t *testing.T) {
	t.Run("[1,2,3,4,5]: range 0:5", func(t *testing.T) {
		_ = BF(context.Background(), []int{1, 2, 3, 4, 5}, 0, 5, 1, func(comb []int) error {
			fmx.Println(comb)
			return nil
		})
	})
}
