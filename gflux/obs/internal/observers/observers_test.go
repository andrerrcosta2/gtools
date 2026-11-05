// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observers

import (
	"errors"
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/core/lifecycle/gates"
)

var TestError = errors.New("test error")

func TestCreate_NextIgnoredAfterClose(t *testing.T) {
	gate := gates.Atomic()

	called := 0
	o := NewNext[int](func(v int) {
		called++
	})

	o = Create(o, gate)

	o.Next(1)
	gate.Close()
	o.Next(2)

	assertlite.True(t, 1 == called)
}

func TestCreate_ErrorClosesAndRunsOnce(t *testing.T) {
	gate := gates.Atomic()

	called := 0
	o := NewError[int](func(err error) {
		called++
	})

	o = Create(o, gate)

	o.Error(TestError)
	o.Error(TestError)

	assertlite.True(t, gate.Closed())
	assertlite.True(t, 1 == called)
}

func TestCreate_CompleteClosesAndRunsOnce(t *testing.T) {
	gate := gates.Atomic()

	called := 0
	o := NewComplete[int](func() {
		called++
	})

	o = Create(o, gate)

	o.Complete()
	o.Complete()

	assertlite.True(t, gate.Closed())
	assertlite.True(t, 1 == called)
}

func TestCreate_ErrorPreventsComplete(t *testing.T) {
	gate := gates.Atomic()

	errCalled := 0
	cptCalled := 0

	o := New[int](
		nil,
		func(error) { errCalled++ },
		func() { cptCalled++ },
	)

	o = Create(o, gate)

	o.Error(TestError)
	o.Complete()

	assertlite.True(t, 1 == errCalled)
	assertlite.True(t, 0 == cptCalled)
}

func TestCreate_CompletePreventsError(t *testing.T) {
	gate := gates.Atomic()

	errCalled := 0
	cptCalled := 0

	o := New[int](
		nil,
		func(error) { errCalled++ },
		func() { cptCalled++ },
	)

	o = Create(o, gate)

	o.Complete()
	o.Error(TestError)

	assertlite.True(t, 0 == errCalled)
	assertlite.True(t, 1 == cptCalled)
}

func TestCreate_NilHandlersAreSafe(t *testing.T) {
	gate := gates.Atomic()

	o := &Observer[int]{}

	o = Create(o, gate)

	assertlite.NotNil(t, o.Next)
	assertlite.NotNil(t, o.Error)
	assertlite.NotNil(t, o.Complete)

	// Should not panic
	o.Next(1)
	o.Error(TestError)
	o.Complete()
}
