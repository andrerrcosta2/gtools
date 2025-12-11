// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"fmt"
	"testing"

	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
)

func TestPrint(t *testing.T) {
	tc := []struct {
		name  string
		value any
	}{
		{
			name:  "models.Map[any, any]",
			value: random.Of[models.Map[any, any]](10).Values(),
		},
	}

	t.Run("print slices", func(t *testing.T) {
		for _, tt := range tc {
			assertlite.NoPanic(t, func() {
				fmt.Println(tt.value)
			})
		}
	})
}
