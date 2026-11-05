// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import "errors"

func newStackFromError(origin Stackable, from error) stack {
	if from == nil {
		return newStack(origin.Unwrap()...)
	}
	flat := FlattenError(from)
	if origin.IsEmpty() {
		return newStack(flat...)

	}
	return newStack(flat...).Push(origin.Unwrap()...)
}

func isStack(e Stackable, target error) bool {
	if target == nil {
		return e.IsEmpty()
	}
	for _, err := range e.Unwrap() {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
