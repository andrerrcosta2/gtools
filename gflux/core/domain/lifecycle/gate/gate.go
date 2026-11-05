package gate

import (
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type (
	Closeable interface {
		Close()
		Closed() bool
	}

	Hookable interface {
		Closeable
		OnClose(function.Runnable)
	}

	RefCount interface {
		Acq()
		Rel()
	}
)
