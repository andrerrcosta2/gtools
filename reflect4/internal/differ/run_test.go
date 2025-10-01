// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"reflect"
	"testing"
)

func assertEquals(t *testing.T, diff differs.Difference) {
	t.Helper()
	assertlite.True(t, diff.Equals, "expected values to be equals but got 'diff.Equals = false',\n%s\n%s",
		diff.Message, diff.Diff)
	assertlite.True(t, diff.Message == "", "expected empty diff message but got %s", diff.Message)
	assertlite.True(t, diff.Diff == "", "expected empty diff diff but got %s", diff.Diff)
	assertlite.NoError(t, diff.Err, "expected no error but got %v", diff.Err)
}

func assertNotEquals(t *testing.T, diff differs.Difference, expMsg, expDiff string) {
	t.Helper()
	assertlite.True(t, diff.Message == expMsg,
		gtests.ErrorDiff("different differ messages", diff.Message, expMsg))
	assertlite.True(t, diff.Diff == expDiff,
		gtests.ErrorDiff("different differ diffs", diff.Diff, expDiff))
	assertlite.False(t, diff.Equals, "expected both channels to be different but got 'diff.Equals'")
}

type TestCase struct {
	name    string
	a       reflect.Value
	b       reflect.Value
	options []internal.Option
	diff    string
	equals  bool
}

func testDifferValues(tt gtests.Loggable, test TestCase) {
	tt.Helper()
	tt.StackTitlef(test.name, "a: %v, b: %v", test.a, test.b)
	diff, equals := Between(test.a, test.b, test.options...)
	if equals != test.equals {
		tt.Errorf("received compare '%t', expected '%t'", equals, test.equals)
		tt.StackErrorf("expected compare '%t', got '%t'", test.equals, equals)
	} else {
		tt.StackSuccessf("Expected compare '%t', got '%t'", test.equals, equals)
	}
	if diff != test.diff {
		tt.StackError(differs.Quick(diff, test.diff))
		tt.Errorf(gtests.ErrorDiff("different diff messages", diff, test.diff))
		tt.StackError(gtests.ErrorDiff("Full defaultStringDiff between messages:", diff, test.diff))
	} else {
		tt.StackSuccessf(gtests.ErrorDiff("compare diff messages", diff, test.diff))
	}
	tt.StackLn()
}

type DifferFunc func(tab indent.Indentor, a, b reflect.Value) differs.Difference

func runPrimTests[T any](t *testing.T, cases []primTableTest[T], differFunc DifferFunc) {
	t.Helper()
	var zero indent.Branch
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			va := reflect.ValueOf(tc.aType)
			vb := reflect.ValueOf(tc.bType)

			received := differFunc(zero, va, vb)

			//typedA := sprints.Typed(tc.typeName, tc.a)
			//typedB := sprints.Typed(tc.typeName, tc.b)
			//
			//assertlite.Equals(t, received.Received, zero.Sprint(typedA))
			//assertlite.Equals(t, received.Expected, zero.Sprint(typedB))
			//assertlite.NoError(t, received.Err)

			if !tc.equal {
				assertlite.True(t, !received.Equals, "expected values to differ but got "+
					"'received.Equals = true'")
				expected := differs.Values(zero, sprints.Typed(tc.typeName, tc.a), sprints.Typed(tc.typeName, tc.b))
				assertlite.True(t, received.Message == expected, gtests.ErrorDiff("different received messages",
					received.Message, expected))
			} else {
				assertlite.True(t, received.Equals, "expected values to be equals but got "+
					"'received.Equals = false'")
				assertlite.True(t, received.Message == "")
			}
		})
	}
}
