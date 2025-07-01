// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/gtests"
	"reflect"
)

type TestCase struct {
	name   string
	a      reflect.Value
	b      reflect.Value
	diff   string
	equals bool
	err    error
}

func testDifferValues(tt gtests.Loggable, test TestCase) {
	tt.Helper()
	tt.StackTitlef(test.name, "a: %v, b: %v", test.a, test.b)
	diff, equals, err := Between(indent.Zero(), test.a, test.b)
	if equals != test.equals {
		tt.Errorf("received equals '%t', expected '%t'", equals, test.equals)
		tt.StackErrorf("expected equals '%t', got '%t'", test.equals, equals)
	} else {
		tt.StackSuccessf("Expected equals '%t', got '%t'", test.equals, equals)
	}
	if diff != test.diff {
		tt.StackError(differs.Quick(diff, test.diff))
		tt.Errorf(gtests.ErrorDiff("different diff messages", diff, test.diff))
		tt.StackError(gtests.ErrorDiff("Full difference between messages:", diff, test.diff))

	} else {
		tt.StackSuccessf(gtests.ErrorDiff("equal diff messages", diff, test.diff))
	}
	if test.err != nil && !errors.Is(err, test.err) {
		tt.StackErrorf("Expected error '%v', got '%v'", test.err, err)
		tt.Errorf("Expected error '%v', got '%v'", test.err, err)
	} else {
		tt.StackSuccessf("Expected error '%v', got '%v'", test.err, err)
	}
	tt.StackLn()
}
