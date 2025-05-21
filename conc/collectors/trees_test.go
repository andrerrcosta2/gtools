// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/conc/collectors/internal/test/mocks"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"testing"
)

func TestBranchableCollector_collect_AddSequentially(t *testing.T) {
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	collector := Branchable[*mocks.Branchable, [][]*mocks.Branchable]().(*branchableCollector[*mocks.Branchable, [][]*mocks.Branchable])

	iterables.OfSlice(mocks.SmallBranchableTestData...).
		// Shuffle the order of the data
		Some(len(mocks.SmallBranchableTestData)).
		Each(func(data *mocks.Branchable) {
			shouldAddToBranchableSuccessfully[*mocks.Branchable](tt, collector, data)
		})

	tt.PrintLogStack()
	branches := collector.get()
	for i, branch := range branches {
		fmt.Printf("Branch %d:\n", i)
		for j, value := range branch {
			fmt.Printf("  [%d]: {id:%d data:%s branch:%+v}\n", j, value.Id, value.Data, dereferenceBranch(value.BranchRef))
		}
	}
}

func dereferenceBranch(branch *mocks.Branchable) interface{} {
	if branch == nil {
		return nil
	}
	return *branch
}
