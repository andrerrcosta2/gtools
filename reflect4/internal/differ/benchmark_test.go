// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/_testdata"
	"reflect"
	"testing"
)

func BenchmarkBetween(b *testing.B) {
	for _, bb := range _testdata.BetweenPerfTests {
		b.Run(reflect.ValueOf(bb.X).Type().String(), func(b *testing.B) {
			va := reflect.ValueOf(bb.X)
			vb := reflect.ValueOf(bb.Y)
			b.ReportAllocs()
			b.ResetTimer() // shouldn't account the reflect op
			for i := 0; i < b.N; i++ {
				_, _ = Between[internal.Option](va, vb)
			}
		})
	}
}
