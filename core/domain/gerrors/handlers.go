// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"unsafe"
)

// FlattenError takes an error and returns a slice of errors that represent the stack trace of the error,
// as well as the last error in the stack trace.
//
// Parameters:
// - err: The error to flatten.
//
// Returns:
// - []error: a slice of errors representing the stack trace of the error.
// - defaultErr: The last error in the stack trace.
func FlattenError(err error) []error {
	var out []error
	var dfs func(error, []error)

	dfs = func(cur error, path []error) {
		if cur == nil {
			return
		}

		// cycle detection: only in current path
		for _, p := range path {
			if Same(p, cur) {
				return
			}
		}
		path = append(path, cur)

		// JOIN: do not include node; recurse into children
		if u, ok := cur.(interface{ Unwrap() []error }); ok {
			for _, ch := range u.Unwrap() {
				dfs(ch, path)
			}
			return
		}

		// WRAP: include wrapper and STOP recursion
		if _, ok := cur.(interface{ Unwrap() error }); ok {
			out = append(out, cur)
			return
		}

		// LEAF
		out = append(out, cur)
	}

	dfs(err, nil)
	return out
}

// ReadTrace takes a slice of errors and returns a formatted string representing the stack trace of these errors.
//
// Parameters:
// - stack: a slice of errors representing the stack trace.
//
// Returns:
// - string: a formatted string representing the stack trace.
func ReadTrace(stack []error) string {
	// Since Join returns nil if every value in errs is nil,
	// stack cannot be empty.
	if len(stack) == 1 {
		return stack[0].Error()
	}

	b := []byte(stack[0].Error())
	for _, err := range stack[1:] {
		b = append(b, '\n')
		b = append(b, err.Error()...)
	}
	// At this point, b has at least one byte '\n'.
	return unsafe.String(&b[0], len(b))
}

func Unwrap(err error) (stk []error) {
	if ss, ok := err.(interface {
		Unwrap() []error
	}); ok {
		return ss.Unwrap()
	} else if sss, ok := err.(interface {
		Unwrap() error
	}); ok {
		return []error{sss.Unwrap()}
	}
	return []error{}
}
