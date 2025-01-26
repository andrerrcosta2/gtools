// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package subsets

import (
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	sbs := Xor(2, 8, 5)
	fmt.Printf("Subsets: %v\n", sbs)
}

func TestOr(t *testing.T) {
	sbs := Or(0, 6, 3)
	fmt.Printf("Subsets: %v\n", sbs)
}
