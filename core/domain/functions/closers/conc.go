// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package closers

import (
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"sync"
)

func SyncThread(wg *sync.WaitGroup, sem gtools.Semaphore) {
	wg.Done()
	sem.Rls()
}
