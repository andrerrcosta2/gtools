// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"sync/atomic"
)

func Data[V any](supplier functions.Supplier2[V, error]) data.Async[V] {
	return &asyncData[V]{
		sync: supplier,
	}
}

type asyncData[V any] struct {
	data V
	ok   atomic.Bool
	sync functions.Supplier2[V, error]
}

func (d *asyncData[V]) Value() (V, bool) {
	if d.ok.Load() {
		return d.data, true
	}
	return d.data, false
}

func (d *asyncData[V]) Sync() (err error) {
	if d.data, err = d.sync(); err == nil {
		d.ok.Store(true)
	}
	return
}
