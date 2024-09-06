// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sentinels

import (
	"fmt"
	"testing"
)

func TestXorRanged(t *testing.T) {
	xor, _ := XorRanged(0, 6, 3)

	fmt.Printf("%d: %v\n", 0, xor.Current())
	for count := 0; ; count++ {
		n, ok := xor.Next()
		if !ok {
			break
		}
		fmt.Printf("%d: %v\n", count, n)
	}
}
