// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/gtests"
)

func shouldAddToBranchableSuccessfully[B data.Branchable[B]](t gtests.Loggable, collector *branchableCollector[B, [][]B], e B) {
	t.Helper()
	collector.collect(e)
	errs := collector.Errors()
	if len(errs) > 0 {
		t.Errorf("error while adding '%+v': %v\n", e, errs)
		return
	}
	t.StackLogf("branchable '%+v' added\n", e)
}
