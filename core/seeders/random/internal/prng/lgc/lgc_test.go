// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lgc

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/seeds"
	"math"
	"testing"
)

const (
	bucketSize    = 50
	distributions = 1_000_000
	deviation     = 0.15
)

func TestFloat32(t *testing.T) {
	t.Run("lgc.Float32 distribution", func(t *testing.T) {
		var (
			expected = float64(distributions) / float64(bucketSize)
			upper    = expected * (1 + deviation)
			lower    = expected * (1 - deviation)
			ba       = make([]int, bucketSize)
			bb       = make([]int, bucketSize)
			width    = math.MaxUint64 / float64(bucketSize)
		)

		for i := 0; i < distributions; i++ {
			a, b := Float32(seeds.Clock(), seeds.Clock())
			idx := int(math.Floor(float64(a) / width))
			ba[idx]++
			idx = int(math.Floor(float64(b) / width))
			bb[idx]++
		}

		for i, count := range ba {
			c := float64(count)
			if c < lower || c > upper {
				l := float64(i) * width
				u := lower + width
				t.Errorf("bucket %2d [%.3f, %.3f): expected ~%.0f ±%.0f (%.1f%%), got %d",
					i, l, u, expected, expected*deviation, deviation*100, count)
			}
		}

		for i, count := range bb {
			c := float64(count)
			if c < lower || c > upper {
				l := float64(i) * width
				u := lower + width
				t.Errorf("bucket %2d [%.3f, %.3f): expected ~%.0f ±%.0f (%.1f%%), got %d",
					i, l, u, expected, expected*deviation, deviation*100, count)
			}
		}
	})
}
