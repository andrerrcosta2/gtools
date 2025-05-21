// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

//go:build tt

package gstrings

import (
	"github.com/andrerrcosta2/gtools/core/internal/assertlite"
	"testing"
)

func TestTreeCreation(t *testing.T) {
	tb := TreeBuilder[string](2, true)
	assertlite.NotNil(t, tb)
}

func TestStructAndValueNodes(t *testing.T) {
	tb := TreeBuilder[string](2, true)

	root, err := tb.NewStructNode("Root")
	assertlite.NoError(t, err)
	assertlite.NotNil(t, root)

	err = root.AddStructChild("Child1")
	assertlite.NoError(t, err)

	err = root.AddValueChild("Value1")
	assertlite.NoError(t, err)

	// Ensure retrieval works
	child, err := root.GetStructChild("Child1")
	assertlite.NoError(t, err)
	assertlite.NotNil(t, child)

	value, err := root.GetValueChild("Value1")
	assertlite.NoError(t, err)
	assertlite.NotNil(t, value)

	// Ensure duplicate nodes are rejected
	err = root.AddStructChild("Child1")
	assertlite.ErrorIs(t, err, NE)

	err = root.AddValueChild("Value1")
	assertlite.ErrorIs(t, err, NE)
}

func TestTreeStringOutput(t *testing.T) {
	tb := TreeBuilder[string](2, true)

	root, _ := tb.NewStructNode("Root")
	root.AddStructChild("Child1")
	root.AddValueChild("Value1")

	child, _ := root.GetStructChild("Child1")
	child.AddValueChild("SubValue")

	value, _ := root.GetValueChild("Value1")
	value.SetValue("Data1")

	subValue, _ := child.GetValueChild("SubValue")
	subValue.SetValue("Data2")

	expected := `Root {
  Child1 {
    SubValue: Data2
  }
  Value1: Data1
}
`
	assertlite.Equal(t, expected, tb.String())
}
