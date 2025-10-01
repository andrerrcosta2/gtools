// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package term

import (
	"fmt"
	"testing"
)

func TestTernBg(t *testing.T) {
	data := "hello world"
	styled := "hello" + BgRed.Style(data[5:6]) + "world"
	fmt.Printf("%s\n", styled)
}
