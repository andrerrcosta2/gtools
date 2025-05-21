// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

func NotEquals(message, diff, received, expected string) Difference {
	return Difference{
		Message:  message,
		Diff:     diff,
		Received: received,
		Expected: expected,
		Equals:   false,
	}
}

func DiffError(message string, err error) Difference {
	return Difference{
		Message: message,
		Equals:  false,
		Err:     err,
	}
}

func Equals(message, received, expected string) Difference {
	return Difference{
		Message:  message,
		Received: received,
		Expected: expected,
		Equals:   true,
	}
}

type Difference struct {
	Message  string
	Diff     string
	Received string
	Expected string
	Equals   bool
	Err      error
}
