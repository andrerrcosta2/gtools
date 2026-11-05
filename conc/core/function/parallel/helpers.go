// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package parallel

import (
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/semaph"
)

func sem(limit int) syncs.Semaphore {
	return semaph.Limiter(limit)
}
