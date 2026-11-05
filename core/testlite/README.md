## gtools/core/testlite
___
Lightweight helpers for Go tests:
- assertion helpers (`assertlite`)
- panic and resource cleanup helpers (`testlite`)
- quick string diffs (`differlite`)
- seed data helpers for tests (`testseed`)

```shell
go get github.com/andrerrcosta2/gtools/core/testlite
```

## Packages
___
- `testlite`: interfaces and helpers around panic handling and close/cleanup behavior.
- `assertlite`: assertion functions that fail the test on mismatch.
- `differlite`: tiny human-readable string diff helper.
- `testseed`: reusable test fixtures and seed values.

## Quick Start
___
```go
package mypkg_test

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestUserName(t *testing.T) {
	user := struct{ Name string }{Name: "andre"}

	assertlite.NotEmpty(t, user.Name)
	assertlite.Equals(t, user.Name, "andre")
}
```

## Core Helpers (`testlite`)
___
- `LogOnPanic(t, f, msg, args...)`: runs `f`, fails test with `Fatalf` if `f` panics.
- `CloseOrLog(t, c, name...)`: calls `Close() error`, logs failures without failing test.
- `CloseOrFail(t, c, name...)`: calls close and marks test failed on close/shutdown timeout.
  - Accepts either `Close() error` or `Close(context.Context) error`.
  - Waits up to 5 seconds for optional `Done() <-chan struct{}` to close.
- `TrackClose(t, c, name...)`: registers `CloseOrFail` with `t.Cleanup`.

Example:
```go
import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite"
)

func TestResource(t *testing.T) {
	r := newResource()
	testlite.TrackClose(t, r, "db-connection")
	// test body...
}
```

## Assertions (`assertlite`)
___
All assertions return `bool` (except `IsErrorOf`) and call `t.Fatal`/`t.Fatalf` on failure.

Common groups:
- Equality:
  - `Equals`, `NotEquals`
  - `EqualSlices` (order-sensitive)
  - `ElementsMatch` (order-insensitive, multiplicity-aware)
  - `EqualMaps`
- Nil/empty:
  - `IsNil`, `NotNil`, `AreNotNil`
  - `IsEmpty`, `NotEmpty`
  - `NoNilFields`, `AllFieldsAreNil`, `NoNilNonInterfaceFields`
- Boolean/predicate:
  - `True`, `False`
  - `AreTrue`, `AreFalse`
- Types/references/panics:
  - `IsTypeOf[T]`, `NotTypeOf[T]`
  - `Same`, `NotSame`
  - `Panics`, `NoPanics`
  - `Eventually`
- Other:
  - `NoError`, `IsErrorOf`, `Contains`, `Len`

Example:
```go
import (
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestEventually(t *testing.T) {
	ready := false
	go func() { ready = true }()

	assertlite.Eventually(t, 500*time.Millisecond, func() bool { return ready })
}
```

## Message Override Behavior
___
Assertions accept `msgAndArgs ...any`.
- If first value is a `string`, it is treated as a `Printf`-style format.
- Remaining values are used as format args.
- If no message is provided, a default assertion message is used.

## `differlite`
___
`differlite.Quick(received, expected)` returns a compact string diff around the first mismatch.

```go
diff := differlite.Quick("abcXYZ", "abc123")
// "Difference at index 3: ..."
```

## Notes
___
- `EqualSlices` treats `nil` and empty slices as different.
- `Same`/`NotSame` are for reference types (pointers, maps, chans, funcs, interfaces).
- `NoNilFields` and related field checks are shallow (no deep recursive traversal).
- `testseed` is intended for tests and examples, not production runtime use.
