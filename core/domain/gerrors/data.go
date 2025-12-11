// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import "fmt"

func newStack(e ...error) stack {
	return stack{
		stk: e,
	}
}

type stack struct {
	stk []error
}

func (s stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s stack) Len() int {
	return len(s.stk)
}

func (s stack) Push(v ...error) stack {
	return stack{stk: append(s.stk, v...)}
}

func (s stack) Pop() (stack, error) {
	l := s.Len()
	if l == 0 {
		return s, nil
	}
	return stack{
		stk: s.stk[:l-1],
	}, s.stk[l-1]
}

func (s stack) Peek() error {
	l := s.Len()
	if l == 0 {
		return nil
	}
	return s.stk[l-1]
}

func (s stack) String() string {
	return fmt.Sprintf("stack{\n\tstk: %v\n}", s.stk)
}
