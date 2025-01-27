// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package seeds

import (
	"fmt"
	"math"
	"testing"
)

func TestClockDistribution(t *testing.T) {
	clocks := make(map[uint64]uint64)
	buckets := uint64(50)
	bucketSize := math.MaxUint64 / buckets
	for i := 0; i < 100000; i++ {
		c := Clock()
		bucket := c / bucketSize
		clocks[bucket]++
	}
	var hid, higher uint64
	var lid, lower uint64 = 0, math.MaxUint64
	for k, v := range clocks {
		if v > higher {
			higher = v
			hid = k
		}
		if v < lower {
			lower = v
			lid = k
		}
	}
	fmt.Printf("Higher[%d]: %d\n", hid, higher)
	fmt.Printf("Lower[%d]: %d\n", lid, lower)
	fmt.Printf("Clocks: %v\n", clocks)
}

func TestClockOddEvenDistribution(t *testing.T) {
	t.Run("Whole Function", func(t *testing.T) {
		clocks := make([]int16, 2)
		for i := 0; i < 1000; i++ {
			clock := Clock()
			if clock%2 == 0 {
				clocks[0]++
			} else {
				clocks[1]++
			}
		}
		fmt.Printf("Odd: %d, Even: %d\n", clocks[0], clocks[1])
	})
}
