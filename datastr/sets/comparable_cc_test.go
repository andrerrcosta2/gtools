// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/keyer/keyers"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/testlite"
)

func TestConcCmp_GeneralTest_WriteOperations(t *testing.T) {
	counter := testlite.Map(t)
	flags := testlite.Flag(t)

	// Create a new concurrent backingMap.
	set := ConcComparable[int]()
	cast, ok := set.(*concCmp[int])
	if !ok {
		t.Errorf("Expected backingMap to be of type concCmp, but it is not")
	}

	// Test insertion in parallel
	it := random.Int(1000, 10, 1000)
	it.Parallel(func(i, v int) {
		set.Add(i)
		flags.Set(true, keyers.CmpToString(i))
	}, 6)

	// Assert the number of registered flags is the same of the size of the backingMap
	flags.AssertRegisteredFlags(set.Len())

	flags.AssertRegisteredFlags(len(cast.table))

	// Test deletion in parallel (only odd values)
	it.Filter(func(v int) bool { return v%2 != 0 }).
		Parallel(func(i, v int) {
			flags.AssertFlags(true, keyers.CmpToString(v))
			set.Remove(v)
			flags.Set(false, keyers.CmpToString(v))
			counter.Add(1, "remove")
		}, 6)

	// Assert the number of false flags is the same of the size of the backingMap
	flags.AssertRegisteredFlagsBy(func(k string, v bool) bool {
		return !v
	}, counter.Get("remove"))
}
