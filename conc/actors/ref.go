// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package actors

type Ref[M any] struct {
	mailbox chan<- M
}

func (r Ref[M]) Tell(msg M) {
	r.mailbox <- msg
}
