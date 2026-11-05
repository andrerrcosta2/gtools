// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package obs

import (
	"fmt"
	"time"
)

type TimeoutError struct {
	After time.Duration
}

func (e TimeoutError) Error() string {
	return fmt.Sprintf("timeout after %s", e.After)
}
