// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import "github.com/andrerrcosta2/gtools/conc/core/syncs"

type Sealable interface {
	Shutdown() syncs.Waiter
	Close() error
}
