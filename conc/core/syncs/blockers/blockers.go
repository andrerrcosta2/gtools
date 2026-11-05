// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package blockers

import "github.com/andrerrcosta2/gtools/conc/core/channel"

type once struct {
	ch chan struct{}
}

func (b *once) Done() channel.Signal {
	return b.ch
}
