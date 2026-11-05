// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"reflect"
	"testing"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
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
		testlite.ErrorDiff("different differ messages", diff.Message, expMsg))
	assertlite.True(t, diff.Diff == expDiff,
		testlite.ErrorDiff("different differ diffs", diff.Diff, expDiff))
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

func testDifferValues(t testlite.LoggerTesting, test TestCase) {
	t.Helper()
	t.Logf(test.name+"a: %v, b: %v", test.a, test.b)
	diff, equals := Between(test.a, test.b, test.options...)
	if equals != test.equals {
		t.Errorf("received compare '%t', expected '%t'", equals, test.equals)
	} else {
		t.Logf("Expected compare '%t', got '%t'", test.equals, equals)
	}
	if diff != test.diff {
		t.Log(differs.Quick(diff, test.diff))
		t.Errorf(testlite.ErrorDiff("different diff messages", diff, test.diff))
		t.Log(testlite.ErrorDiff("Full defaultStringDiff between messages:", diff, test.diff))
	} else {
		t.Logf(testlite.ErrorDiff("compare diff messages", diff, test.diff))
	}
}

type DifferFunc func(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference

func runPrimTests[T any](t *testing.T, cases []primTableTest[T], differFunc DifferFunc) {
	t.Helper()
	var zero indent.Branch
	s := DefaultStrat()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			va := reflect.ValueOf(tc.aType)
			vb := reflect.ValueOf(tc.bType)

			received := differFunc(zero, va, vb, s)

			if !tc.equal {
				assertlite.True(t, !received.Equals, "expected values to differ but got "+
					"'received.Equals = true'")
				expected := differs.Values(zero, sprints.Typed(tc.typeName, tc.a), sprints.Typed(tc.typeName, tc.b))
				assertlite.True(t, received.Message == expected, testlite.ErrorDiff("different received messages",
					received.Message, expected))
			} else {
				assertlite.True(t, received.Equals, "expected values to be equals but got "+
					"'received.Equals = false'")
				assertlite.True(t, received.Message == "")
			}
		})
	}
}
