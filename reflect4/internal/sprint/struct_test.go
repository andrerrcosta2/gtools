// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/_testdata"
	"reflect"
	"testing"
)

func TestFields_BasicStructs(t *testing.T) {
	zero := indent.Zero()
	tests := []struct {
		name     string
		input    reflect.Value
		expected string
		err      error
	}{
		{
			name: "Test struct{ Name string, Age  }",
			input: reflect.ValueOf(struct {
				Name string
				Age  int
			}{
				Name: "Alice",
				Age:  30,
			}),
			expected: "\t" + sprints.TypeValue(zero, "Name", "string") +
				sprints.LtKeyValue(zero, "Age", "int"),
			err: nil,
		},
		{
			name: "Test unexported field struct{ Name string, age int }",
			input: reflect.ValueOf(struct {
				Name string
				age  int
			}{
				Name: "Alice",
				age:  30,
			}),
			expected: "\t" + sprints.TypeValue(zero, "Name", "string") +
				sprints.LtKeyValue(zero, "age", "int"),
			err: nil,
		},
		{
			name:  "Test outer unexported field struct Person{ Name string, age int }",
			input: reflect.ValueOf(_testdata.UnexportedField),
			expected: "\t" + sprints.TypeValue(zero, "Name", "string") +
				sprints.LtKeyValue(zero, "age", "int"),
			err: nil,
		},
		{
			name:  "Test pointer unexported field struct &Person{ Name string, age int }",
			input: reflect.ValueOf(&_testdata.UnexportedField),
			expected: "\t" + sprints.TypeValue(zero, "Name", "string") +
				sprints.LtKeyValue(zero, "age", "int"),
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Fields(zero, test.input)
			if err != nil {
				if test.err == nil {
					t.Errorf("unexpected error %s", err)
				} else if err.Error() != test.err.Error() {
					t.Errorf("mismatched errors. \nreceived:\n%q, \nexpected:\n%q",
						err, test.err,
					)
				}
			}
			if got != test.expected {
				t.Errorf(differs.Quick(got, test.expected))
			} else {
				t.Log(fmx.SGreenf("fields match"))
			}
		})
	}

}
