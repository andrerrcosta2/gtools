package limiters

import (
	"fmt"
	"sync/atomic"
)

func Intolerant(limit int32, errorMessage ...any) *intolerant {
	errMsg := buildMessage("concurrency limit exceeded", errorMessage...)
	return &intolerant{limit: limit, errorMessage: errMsg}
}

func buildMessage(defaultMsg string, msgAndArgs ...any) string {
	if len(msgAndArgs) > 0 {
		if format, ok := msgAndArgs[0].(string); ok {
			return fmt.Sprintf(format, msgAndArgs[1:]...)
		}
	}
	return defaultMsg
}

type intolerant struct {
	limit        int32
	inUse        atomic.Int32
	errorMessage string
}

func (s *intolerant) Acq() {
	if s.inUse.Add(1) > s.limit {
		s.inUse.Add(-1)
		panic(s.errorMessage)
	}
}

func (s *intolerant) Rls() {
	if s.inUse.Add(-1) < 0 {
		panic("unbalanced Release")
	}
}
