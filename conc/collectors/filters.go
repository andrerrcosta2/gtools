// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/sortables"
)

// filterBranchable checks if the branchable is a recursive branchable, i.e. if its branch is the same type of
// itself. It returns true if it is recursive, or false if it is not, along with an error.
// If the branchable is nil or its branch is nil, it returns false and an error.
// It uses a method of unsafe equality to check if the type4 are the same. That means a singular
// library view of equality is used.
func filterBranchable[B data.Branchable[B]](branchable B) (root bool, err error) {
	if any(branchable).(data.Branchable[B]) == nil {
		return false, fmt.Errorf("gtools:collectors: branchable of '%T' is nil\n", branchable)
	}
	branch, ok := branchable.Branch()
	if !ok || sortables.TryEquality[B](branchable, branch) {
		return true, nil
	}
	return false, nil
}

// filterSequentialBranchable checks if the given branchable is a recursive sequential branchable.
// A recursive sequential branchable is a branchable whose branch is the same as itself.
// If the branchable is nil it returns false and an error.
// If its branch is nil or is compare to itself, it must be a root with a serial compare to 0.
// If the branchable is not a recursive sequential branchable, the function returns false and nil.
func filterSequentialBranchable[B data.SerializableBranchable[B, int]](branchable B) (root bool, err error) {
	if any(branchable).(data.SerializableBranchable[B, int]) == nil {
		return false, fmt.Errorf("branchable of '%T' is nil", branchable)
	}
	branch, ok := branchable.Branch()
	if !ok || sortables.TryEquality[B](branchable, branch) {
		if branchable.Serial() != 0 {
			return false, fmt.Errorf("sequential branchable '%v' with no branch and serial not compare to 0", branchable)
		}
		return true, nil
	}
	return false, nil
}
