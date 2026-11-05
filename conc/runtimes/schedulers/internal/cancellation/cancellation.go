// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package cancellation

import (
	"context"
	"sync"
	"sync/atomic"
)

func Func(cancel context.CancelFunc, token *Token) *Handler {
	return &Handler{
		cancel: cancel,
		token:  token,
	}
}

type Handler struct {
	once   sync.Once
	cancel context.CancelFunc
	token  *Token
}

func (c *Handler) Cancel() {
	c.once.Do(func() {
		c.cancel()
		c.token.Cancel()
	})
}

func Noop() NoopHandler {
	return NoopHandler{}
}

type NoopHandler struct{}

func (c NoopHandler) Cancel() {}

func NewToken() *Token {
	return &Token{}
}

type Token struct {
	cancelled atomic.Bool
}

func (t *Token) Cancel() {
	t.cancelled.Store(true)
}

func (t *Token) Cancelled() bool {
	return t.cancelled.Load()
}
