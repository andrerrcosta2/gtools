// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"testing"
)

func TestConcCmp_GeneralTest_WriteOperations(t *testing.T) {
	tt := testingtools.ConcLite(t, testlogs.Default)

	// Create a new concurrent set.
	set := ConcComparable[int]()
	cast, ok := set.(*concCmp[int])
	if !ok {
		t.Errorf("Expected set to be of type concCmp, but it is not")
	}

	// Test insertion in parallel
	it := random.Int(1000, 10, 1000)
	it.Parallel(func(i, v int) {
		set.Add(i)
		tt.Flag(true, comparators.StringHash(i))
	}, 6)

	// Assert the number of registered flags is the same of the size of the set
	tt.AssertRegisteredFlags(set.Len(), "Expected '%d' registered flags, but got %d",
		set.Len(), tt.FlagsSize())

	tt.AssertRegisteredFlags(len(cast.set), "Expected '%d' registered flags, but got %d",
		len(cast.set), tt.FlagsSize())

	// Test deletion in parallel (only odd values)
	it.Filter(func(v int) bool { return v%2 != 0 }).
		Parallel(func(i, v int) {
			tt.AssertFlag(true, comparators.StringHash(v))
			set.Remove(v)
			tt.Flag(false, comparators.StringHash(v))
			tt.RegisterCalls(1, "remove")
		}, 6)

	// Assert the number of false flags is the same of the size of the set
	tt.AssertRegisteredFlagsBy(func(k string, v bool) bool {
		return !v
	}, tt.CallsTo("remove"), "Expected '%d' false flags, but got %d")
}
